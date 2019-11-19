package core

import (
	"context"
	"log"
	"net/http"
	"os"
)

type ErrorHandler interface {
	// ServeError 处理回调的错误, 比如 xml 解码出错, return_code != "SUCCESS", result_code != "SUCCESS", ...
	ServeError(http.ResponseWriter, *http.Request, error, context.Context)
}

var DefaultErrorHandler ErrorHandler = ErrorHandlerFunc(defaultErrorHandlerFunc)

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request, error, context.Context)

func (fn ErrorHandlerFunc) ServeError(w http.ResponseWriter, r *http.Request, err error, ctx context.Context) {
	fn(w, r, err, ctx)
}

var errorLogger = log.New(os.Stderr, "[WECHAT_ERROR] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

func defaultErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error, ctx context.Context) {
	errorLogger.Output(3, err.Error())
}
