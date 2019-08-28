package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var conf Config = Config{
	Viper:       viper.New(),
	Mysql:       nil,
	Redis:       nil,
	Mail:        nil,
	Server:      nil,
	DingTalk:    nil,
	Logs:        nil,
	Application: nil,
	Parameters:  nil,
	Mq:          nil,
	OSS:         nil,
}

type Config struct {
	Viper       *viper.Viper
	Mysql       *MysqlConfig
	Redis       *RedisConfig
	Mail        *MailConfig
	Server      *ServerConfig
	DingTalk    *DingTalkSDKConfig
	Logs        *map[string]LogConfig
	Application *ApplicationConfig
	Parameters  *map[string]interface{}
	Mq          *MQConfig
	OSS         *OSSConfig
}

type MysqlConfig struct {
	Host     string
	Port     int
	Usr      string
	Pwd      string
	Database string
}

type RedisConfig struct {
	Host           string
	Port           int
	Pwd            string
	Database       int
	MaxIdle        int
	MaxActive      int
	MaxIdleTimeout int
}

type OSSConfig struct {
	BucketName      string
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
}

type MailConfig struct {
	Usr  string
	Pwd  string
	Host string
	Port int
}

type ServerConfig struct {
	Port int
	Name string
}

type DingTalkSDKConfig struct {
	SuiteKey    string
	SuiteSecret string
	Token       string
	AesKey      string
	AppId       int64
}

type LogConfig struct {
	LogPath      string
	Level        string
	FileSize     int64
	FileNum      int
	IsConsoleOut bool
}

type ApplicationConfig struct {
	RunMode   int
	CacheMode string
}

type MQConfig struct {
	Kafka  *KafkaMQConfig
}

type RocketMQConfig struct {
	GroupID    string
	NameServer string
	Log        *LogConfig
}

//KafKa MQ Config
type KafkaMQConfig struct {
	NameServers string
}

func GetMysqlConfig() *MysqlConfig {
	return conf.Mysql
}

func GetKafkaConfig() *KafkaMQConfig {
	if conf.Mq == nil {
		panic(errors.New("mq configuration is nil!"))
	}
	if conf.Mq.Kafka == nil {
		panic(errors.New("kafka configuration is nil!"))
	}
	return conf.Mq.Kafka
}

func GetOSSConfig() *OSSConfig {
	return conf.OSS
}

func GetRedisConfig() *RedisConfig {
	return conf.Redis
}

func GetConfig() Config {
	return conf
}

func GetMailConfig() *MailConfig {
	return conf.Mail
}

func GetServerConfig() *ServerConfig {
	return conf.Server
}

func GetDingTalkSdkConfig() *DingTalkSDKConfig {
	return conf.DingTalk
}

func GetApplication() *ApplicationConfig {
	return conf.Application
}

func GetLogConfig(name string) *LogConfig {
	c := (*conf.Logs)[name]
	return &c
}

func GetMQ() *MQConfig {
	return conf.Mq
}

func GetParameters() *map[string]interface{} {
	return conf.Parameters
}

func GetParameter(key string) interface{} {
	key = strings.ToLower(key)
	if conf.Parameters == nil {
		panic(errors.New("Parameters configuration is nil!"))
	}
	ps := *conf.Parameters
	if ps[key] == nil {
		panic(errors.Errorf("Parameter %s Not configured!", key))
	}
	return ps[key]
}

func LoadConfig(dir string, config string) error {
	return LoadEnvConfig(dir, config, "")
}

func LoadEnvConfig(dir string, config string, env string) error {
	//if env != "" {
	//	config += env
	//}
	conf.Viper.SetConfigName(config)
	conf.Viper.AddConfigPath(dir)
	conf.Viper.SetConfigType("yaml")
	if err := conf.Viper.ReadInConfig(); err != nil {
		return err
	}

	if env == "" {
		if err := conf.Viper.Unmarshal(&conf); err != nil {
			return err
		}

		return nil
	}

	configs := conf.Viper.AllSettings()
	viper2 := viper.New()

	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		viper2.SetDefault(k, v)
	}

	viper2.SetConfigName(config + "." + env)
	viper2.AddConfigPath(dir)
	viper2.SetConfigType("yaml")
	if err := viper2.ReadInConfig(); err != nil {
		return err
	}
	conf.Viper = viper2
	if err := conf.Viper.Unmarshal(&conf); err != nil {
		return err
	}

	return nil
}
