package zlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitLogger(t *testing.T) {
	// mock InitLogger
	type args struct {
		//
		config  LogsConfig
		flavors []LogsConfigFlavors
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "debug",
			args: args{
				config: LogsConfigDebug(),
			},
		},
		{
			name: "production",
			args: args{
				config: LogsConfigProduction(),
			},
		},
		{
			name: "default",
			args: args{
				config: LogsConfigDefault(),
			},
		},
		{
			name: "default-with-flavor",
			args: args{
				config: LogsConfigDefault(),
				flavors: []LogsConfigFlavors{
					{
						Name:       "foo",
						LogsConfig: LogsConfigDefault(),
					},
					{
						Name:       "bar",
						LogsConfig: LogsConfigDefault(),
					},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do InitLogger
			logsConfigFlavors := tc.args.flavors
			gotErr := InitLogger(tc.args.config, logsConfigFlavors...)

			// verify InitLogger
			assert.Equal(t, tc.wantErr, gotErr != nil)
			if tc.wantErr {
				t.Logf("want init error %v", gotErr)
				return
			}
			getConfig := GetLoggerConfig()
			assert.NotNil(t, getConfig)
			Log().Infof("init zlog success by mode name: %s", tc.name)

			assert.Equal(t, tc.args.config.Level, getConfig.Level)

			if len(logsConfigFlavors) > 0 {
				for _, flavor := range logsConfigFlavors {
					flavorsLogger := GetFlavorsLogger(flavor.Name)
					assert.NotNil(t, flavorsLogger)
					assert.Equal(t, flavor.Level, flavorsLogger.Level())

					flavorsSugaredLogger := GetFlavorsSugaredLogger(flavor.Name)
					assert.NotNil(t, flavorsSugaredLogger)
					assert.Equal(t, flavor.Level, flavorsSugaredLogger.Level())
					flavorsSugaredLogger.Infof("just use flavors as: %s", flavor.Name)
				}
			}

			pruneLogFolder, gotErrPruneLogs := getConfig.PruneLogs()
			assert.Nil(t, gotErrPruneLogs)
			t.Logf("prune Logs at folder: %s", pruneLogFolder)

			errRemoveInit := removeInit()
			assert.Nil(t, errRemoveInit)
		})
	}
}
