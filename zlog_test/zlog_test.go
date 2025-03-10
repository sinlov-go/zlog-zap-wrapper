package zlog_test

import (
	"github.com/sinlov-go/zlog-zap-wrapper/zlog"
	"testing"
)

func TestLogInfoPrint(t *testing.T) {

	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test info",
			args: args{
				format: "test info",
				v:      nil,
			},
		},
		{
			name: "testBaseFolderPath",
			args: args{
				format: "testBaseFolderPath.GetTestDataFolderFullPath %s",
				v:      []interface{}{testGoldenKit.GetTestDataFolderFullPath()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zlog.Log().Infof(tt.args.format, tt.args.v...)
		})
	}
}
