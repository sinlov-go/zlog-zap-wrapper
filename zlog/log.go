package zlog

import (
	"fmt"
	"go.uber.org/zap"
)

// logSugared
// 全局日志变量
// var logSugared *zap.Logger
var logSugared *zap.SugaredLogger

func Log() *zap.SugaredLogger {
	if logSugared == nil {
		panic(fmt.Errorf("please check zlog init then use"))
	}
	return logSugared
}

// accessLoggerGlobal
// 全局 accessLogger
var accessLoggerGlobal *zap.Logger

func AccessLogger() *zap.Logger {
	if accessLoggerGlobal == nil {
		panic(fmt.Errorf("please check zlog init then use"))
	}
	return accessLoggerGlobal
}

// cronLoggerGlobal
// 全局 cronLogger
var cronLoggerGlobal *zap.Logger

func CronLogger() *zap.Logger {
	if cronLoggerGlobal == nil {
		panic(fmt.Errorf("please check zlog init then use"))
	}
	return cronLoggerGlobal
}
