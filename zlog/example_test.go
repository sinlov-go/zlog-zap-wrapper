package zlog_test

import "github.com/sinlov-go/zlog-zap-wrapper/zlog"

// To use zlog, must call InitLogger before use
// the SetOutput function when your application starts.
// Default logger use Log to use.
func Example_zLogInitLogger() {
	cfg := zlog.LogsConfigDefault()
	flavorsBar := zlog.LogsConfigFlavors{}
	errDeepCopy := flavorsBar.DeepCopyFromConfig("bar", cfg)
	if errDeepCopy != nil {
		panic(errDeepCopy)
	}

	errInit := zlog.InitLogger(cfg)
	if errInit != nil {
		panic(errInit)
	}

	zlog.Log().Info("init success")
}

// To use zlog, must call InitLogger before use
// the SetOutput function when your application starts.
// Flavors can be used to customize the logger.
// If flavors logger path same with default logger path, will append name after default logger.
// flavors get from GetFlavorsSugaredLogger to use.
func Example_zLogInitWithFlavors() {
	// init config
	cfg := zlog.LogsConfigDefault()
	// init flavors
	flavorsBar := zlog.LogsConfigFlavors{}
	errDeepCopy := flavorsBar.DeepCopyFromConfig("bar", cfg)
	if errDeepCopy != nil {
		panic(errDeepCopy)
	}
	configFlavors := []zlog.LogsConfigFlavors{flavorsBar}

	// do zlog init
	errInit := zlog.InitLogger(cfg, configFlavors...)
	if errInit != nil {
		panic(errInit)
	}

	// use to print log
	zlog.Log().Info("init success")
	zlog.GetFlavorsSugaredLogger("bar").Info("bar flavors info")
}
