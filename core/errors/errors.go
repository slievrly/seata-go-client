/**
2 * @Author: Nico
3 * @Date: 2019/9/7 19:06
4 */
package errors

import "fmt"

type SEATAError struct {
	Code    int64
	Message string
	Err     error
}

type ErrorInfo struct {
	Code    int64
	Message string
}

func (e SEATAError) Error() string {
	return fmt.Sprintf("seata error code [%d], error msg [%s], error stack [%v]", e.Code, e.Message, e.Err)
}

func NewErrorInfo(code int64, msg string) ErrorInfo {
	return ErrorInfo{
		Code:    code,
		Message: msg,
	}
}

func Error(errorInfo ErrorInfo, errs ...error) SEATAError {
	var err error = nil
	if len(errs) > 0 {
		err = errs[0]
	}
	return SEATAError{
		Code:    errorInfo.Code,
		Message: errorInfo.Message,
		Err:     err,
	}
}

var (
	NotSupportSqlTypeError = NewErrorInfo(10001, "不支持的sql类型")
	SQLParserError         = NewErrorInfo(10002, "sql解析失败")
)
