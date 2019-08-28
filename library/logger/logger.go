package logger

import (
	"github.com/slievrly/seata-go-client/library/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"os"
	"sync"
)

var (
	_lock   sync.Mutex
	_logger = map[string]*SysLogger{}
)

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

type SysLogger struct {
	log          *zap.SugaredLogger
	name         string
	path         string
	level        zapcore.Level
	isConsoleOut bool
}

func (s *SysLogger) Info(msg string) {
	s.Init()
	s.log.Info(msg)
}

func (s *SysLogger) InfoF(template string, args ...interface{}) {
	s.Init()
	s.log.Infof(template, args...)
}

func (s *SysLogger) Error(msg string) {
	s.Init()
	s.log.Error(msg)
}

func (s *SysLogger) ErrorF(template string, args ...interface{}) {
	s.Init()
	s.log.Errorf(template, args...)
}

func (s *SysLogger) Debug(msg string) {
	s.Init()
	s.log.Debug(msg)
}

func (s *SysLogger) DebugF(template string, args ...interface{}) {
	s.Init()
	s.log.Debugf(template, args...)
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// 获得默认Logger对象
func GetDefaultLog() *SysLogger {
	return CreateLogger("default")
}

func CreateLogger(name string) *SysLogger {
	return &SysLogger{
		name: name,
	}
}

// 获得Logger对象
func (s *SysLogger) Init() {
	if s.log == nil {
		_lock.Lock()
		defer _lock.Unlock()
		if s.log == nil {
			s.InitLogger()
		}
	}
}

func (s *SysLogger) InitLogger() *SysLogger {
	logConfig := config.GetLogConfig(s.name)
	if logConfig == nil {
		panic(s.name + " log config not exist!")
	}

	level := getLoggerLevel(logConfig.Level)
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logConfig.LogPath, // ⽇志⽂件路径
		MaxSize:    1024,              // megabytes
		MaxBackups: 20,                //最多保留20个备份
		LocalTime:  true,
		Compress:   true, // 是否压缩 disabled by default
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	//consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	var allCore []zapcore.Core

	if logConfig.IsConsoleOut == true {
		allCore = append(allCore, zapcore.NewCore(zapcore.NewJSONEncoder(encoder), consoleDebugging, zap.NewAtomicLevelAt(level)))
	}

	allCore = append(allCore, zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level)))

	core := zapcore.NewTee(allCore...)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	s.log = logger.Sugar()
	s.path = logConfig.LogPath
	s.level = level
	s.isConsoleOut = logConfig.IsConsoleOut
	_logger[logConfig.LogPath] = s

	return _logger[logConfig.LogPath]
}
