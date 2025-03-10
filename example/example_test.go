package example

import (
	"github.com/sinlov-go/zlog-zap-wrapper/zlog"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

// To use this lib, can load by viper
// the SetOutput function when your application starts.
func Example() {
	// load viper
	viper.SetConfigType("yml")
	viper.SetConfigFile("config.yml")

	errViperReadInConfig := viper.ReadInConfig()
	if errViperReadInConfig != nil {
		panic(errViperReadInConfig)
	}

	var logConfig zlog.LogsConfig

	errViperUnmarshal := viper.Unmarshal(&logConfig)
	if errViperUnmarshal != nil {
		panic(errViperUnmarshal)
	}

	errInitLogger := zlog.InitLogger(logConfig)
	if errInitLogger != nil {
		panic(errInitLogger)
	}

	zlog.Log().Info("hello zlog")
}

func TestPruneLogs(t *testing.T) {
	logsConfig := zlog.LogsConfigDefault()
	logsConfig.PathBase = "log/log"

	errInitLogger := zlog.InitLogger(logsConfig)
	if errInitLogger != nil {
		panic(errInitLogger)
	}

	zlog.Log().Info("hello zlog")

	getConfig := zlog.GetLoggerConfig()
	assert.NotNil(t, getConfig)
	zlog.Log().Infof("log folder: %s", getConfig.PathBase)

	// is system windows
	if runtime.GOOS == "windows" {
		t.Skip("skip test on windows, by err like &fs.PathError{Op:\"remove\", Path:\"log\\\\log\\\\info\\\\xxx.log\", Err:0x20}")
	} else {
		pruneLogFolder, gotErrPruneLogs := getConfig.PruneLogs()
		assert.Nil(t, gotErrPruneLogs)
		t.Logf("prune Logs at folder: %s", pruneLogFolder)
	}
}

func TestInitLogger(t *testing.T) {
	// mock InitLogger
	type args struct {
		//
		config  zlog.LogsConfig
		flavors []zlog.LogsConfigFlavors
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "debug",
			args: args{
				config: zlog.LogsConfigDebug(),
			},
		},
		{
			name: "production",
			args: args{
				config: zlog.LogsConfigProduction(),
			},
		},
		{
			name: "default",
			args: args{
				config: zlog.LogsConfigDefault(),
			},
		},
		{
			name: "default-with-flavor",
			args: args{
				config: zlog.LogsConfigDefault(),
				flavors: []zlog.LogsConfigFlavors{
					{
						Name:       "foo",
						LogsConfig: zlog.LogsConfigDefault(),
					},
					{
						Name:       "bar",
						LogsConfig: zlog.LogsConfigDefault(),
					},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			errRemoveInit := zlog.DestructorInit()
			assert.Nil(t, errRemoveInit)

			// do InitLogger
			logsConfigFlavors := tc.args.flavors
			gotErr := zlog.InitLogger(tc.args.config, logsConfigFlavors...)

			// verify InitLogger
			assert.Equal(t, tc.wantErr, gotErr != nil)
			if tc.wantErr {
				t.Logf("want init error %v", gotErr)
				return
			}
			getConfig := zlog.GetLoggerConfig()
			assert.NotNil(t, getConfig)
			zlog.Log().Infof("init zlog success by mode name: %s", tc.name)

			assert.Equal(t, tc.args.config.Level, getConfig.Level)

			if len(logsConfigFlavors) > 0 {
				for _, flavor := range logsConfigFlavors {
					flavorsLogger := zlog.GetFlavorsLogger(flavor.Name)
					assert.NotNil(t, flavorsLogger)
					assert.Equal(t, flavor.Level, flavorsLogger.Level())

					flavorsSugaredLogger := zlog.GetFlavorsSugaredLogger(flavor.Name)
					assert.NotNil(t, flavorsSugaredLogger)
					assert.Equal(t, flavor.Level, flavorsSugaredLogger.Level())
					flavorsSugaredLogger.Infof("just use flavors as: %s", flavor.Name)
				}
			}

			// windows will got error like
			// got: &fs.PathError{Op:"remove", Path:"logs\\log\\bar\\info\\xxxx", Err:0x20}
			//pruneLogFolder, gotErrPruneLogs := getConfig.PruneLogs()
			//assert.Nil(t, gotErrPruneLogs)
			//t.Logf("prune Logs at folder: %s", pruneLogFolder)
		})
	}
}
