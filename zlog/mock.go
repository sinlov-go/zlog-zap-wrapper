package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// MockZapLoggerInit
// for unit test
func MockZapLoggerInit() {
	//logger, _ := zap.NewDevelopment()

	var coreArr []zapcore.Core

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	levelAt := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	stdoutCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), levelAt)
	coreArr = append(coreArr, stdoutCore)
	logger := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
	logSugared = logger.Sugar()
}
