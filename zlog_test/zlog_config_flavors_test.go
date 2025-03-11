package zlog_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/sinlov-go/zlog-zap-wrapper/zlog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogsConfigFlavors(t *testing.T) {
	// mock LogsConfigFlavors
	type args struct {
		config zlog.LogsConfig
		name   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				config: zlog.LogsConfigDefault(),
			},
			wantErr: true,
		},
		{
			name: "default",
			args: args{
				name:   "default",
				config: zlog.LogsConfigDefault(),
			},
		},
		{
			name: "production",
			args: args{
				name:   "production",
				config: zlog.LogsConfigProduction(),
			},
		},
		{
			name: "debug",
			args: args{
				name:   "debug",
				config: zlog.LogsConfigDebug(),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do LogsConfigFlavors
			logsCfgFlavors := zlog.LogsConfigFlavors{}
			errCopyFromConfig := logsCfgFlavors.DeepCopyFromConfig(tc.args.name, tc.args.config)

			// verify LogsConfigFlavors
			assert.Equal(t, tc.wantErr, errCopyFromConfig != nil)
			if tc.wantErr {
				t.Logf("err DeepCopyFromConfig %v", errCopyFromConfig)
				return
			}
			assert.Equal(t, tc.name, logsCfgFlavors.Name)
			assert.Equal(t, tc.args.config, logsCfgFlavors.LogsConfig)

			copyCfg, errDeepCopyToConfig := logsCfgFlavors.DeepCopyToConfig()
			assert.Equal(t, tc.wantErr, errDeepCopyToConfig != nil)
			if tc.wantErr {
				t.Logf("err DeepCopyToConfig %v", errDeepCopyToConfig)
			}
			assert.Equal(t, tc.args.config, *copyCfg)

			copyNew, errDeepCopyNew := logsCfgFlavors.DeepCopyNew()
			assert.Equal(t, tc.wantErr, errDeepCopyNew != nil)
			if tc.wantErr {
				t.Logf("err DeepCopyNew %v", errDeepCopyNew)
			}
			assert.Equal(t, logsCfgFlavors, *copyNew)

			logsNewCopyFlavors := zlog.LogsConfigFlavors{}
			errDeepCopyFrom := logsNewCopyFlavors.DeepCopyFrom(logsCfgFlavors)
			assert.Equal(t, tc.wantErr, errDeepCopyFrom != nil)
			assert.Equal(t, logsCfgFlavors, *copyNew)
		})
	}
}

func TestLogsConfigFlavorsFormat(t *testing.T) {
	// mock LogsConfigFlavorsFormat
	type args struct {
		name   string
		config zlog.LogsConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				config: zlog.LogsConfigDefault(),
			},
			wantErr: true,
		},
		{
			name: "default",
			args: args{
				name:   "default",
				config: zlog.LogsConfigDefault(),
			},
		},
		{
			name: "production",
			args: args{
				name:   "production",
				config: zlog.LogsConfigProduction(),
			},
		},
		{
			name: "debug",
			args: args{
				name:   "debug",
				config: zlog.LogsConfigDebug(),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			logsCfgFlavors := zlog.LogsConfigFlavors{}
			errCopyFromConfig := logsCfgFlavors.DeepCopyFromConfig(tc.args.name, tc.args.config)
			assert.Equal(t, tc.wantErr, errCopyFromConfig != nil)
			if tc.wantErr {
				t.Logf("err DeepCopyFromConfig %v", errCopyFromConfig)
				return
			}

			// do LogsConfigFlavorsFormat
			// verify LogsConfigFlavorsFormat
			g.AssertJson(t, t.Name(), logsCfgFlavors)
		})
	}
}
