package zlog_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/sinlov-go/zlog-zap-wrapper/zlog"
	"testing"
)

func TestLogsConfigFormat(t *testing.T) {
	// mock LogsConfigFlavorsFormat
	type args struct {
		config zlog.LogsConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				config: zlog.LogsConfigDefault(),
			},
		},
		{
			name: "production",
			args: args{
				config: zlog.LogsConfigProduction(),
			},
		},
		{
			name: "debug",
			args: args{
				config: zlog.LogsConfigDebug(),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// verify LogsConfigFlavorsFormat
			g.AssertJson(t, t.Name(), tc.args.config)
		})
	}
}
