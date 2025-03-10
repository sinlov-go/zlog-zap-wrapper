package zlog

import (
	"fmt"
	"go.uber.org/zap"
)

// DestructorInit
// destructor init
// warning, this will let panic(), make sure what you want to do.
// Most use before InitLogger()
func DestructorInit() error {
	rwLockInit.Lock()

	logSugared = nil
	savedLoggerConfig = nil

	loggerFlavorsMap = nil
	logSugaredFlavorsMap = nil

	defer rwLockInit.Unlock()
	return nil
}

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

var (
	loggerFlavorsMap     map[string]*zap.Logger
	logSugaredFlavorsMap map[string]*zap.SugaredLogger
)

func GetFlavorsLogger(name string) *zap.Logger {
	if loggerFlavorsMap == nil {
		panic(fmt.Errorf("please check zlog init then use"))
	}
	logger, ok := loggerFlavorsMap[name]
	if !ok {
		panic(fmt.Errorf("not found logger flaver by %s", name))
	}
	return logger
}

func GetFlavorsSugaredLogger(name string) *zap.SugaredLogger {
	if logSugaredFlavorsMap == nil {
		panic(fmt.Errorf("please check zlog init then use"))
	}
	logFlavorsSugared, ok := logSugaredFlavorsMap[name]
	if !ok {
		panic(fmt.Errorf("not found Sugared logger flaver by %s", name))
	}
	return logFlavorsSugared
}
