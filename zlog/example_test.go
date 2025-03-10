package zlog

// To use zlog, must call InitLogger before use
// the SetOutput function when your application starts.
// Default logger use Log to use.
func Example_zLogInitLogger() {
	cfg := LogsConfigDefault()
	flavorsBar := LogsConfigFlavors{}
	errDeepCopy := flavorsBar.DeepCopyFromConfig("bar", cfg)
	if errDeepCopy != nil {
		panic(errDeepCopy)
	}

	errInit := InitLogger(cfg)
	if errInit != nil {
		panic(errInit)
	}

	Log().Info("init success")
}

// To use zlog, must call InitLogger before use
// the SetOutput function when your application starts.
// Flavors can be used to customize the logger.
// If flavors logger path same with default logger path, will append name after default logger.
// flavors get from GetFlavorsSugaredLogger to use.
func Example_zLogInitWithFlavors() {
	// init config
	cfg := LogsConfigDefault()
	// init flavors
	flavorsBar := LogsConfigFlavors{}
	errDeepCopy := flavorsBar.DeepCopyFromConfig("bar", cfg)
	if errDeepCopy != nil {
		panic(errDeepCopy)
	}
	configFlavors := []LogsConfigFlavors{flavorsBar}

	// do zlog init
	errInit := InitLogger(cfg, configFlavors...)
	if errInit != nil {
		panic(errInit)
	}

	// use to print log
	Log().Info("init success")
	GetFlavorsSugaredLogger("bar").Info("bar flavors info")
}
