package example

import (
	"github.com/sinlov-go/zlog-zap-wrapper/zlog"
	"github.com/spf13/viper"
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
