package zlog

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// DynamicCallerEncoder
// custom encoder: dynamically hide INFO-level call paths
type DynamicCallerEncoder struct {
	zapcore.Encoder // 内嵌原始编码器（如JSON或Console）
}

func (enc *DynamicCallerEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	switch ent.Level { // 清空调用路径
	default:
	// do nothing
	case zapcore.DebugLevel:
		ent.Caller = zapcore.EntryCaller{}
	case zapcore.InfoLevel:
		ent.Caller = zapcore.EntryCaller{}
	case zapcore.WarnLevel:
		ent.Caller = zapcore.EntryCaller{}
	}
	return enc.Encoder.EncodeEntry(ent, fields)
}
