package zlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"path/filepath"
	"time"
)

func GetLoggerConfig() LogsConfig {
	if savedLoggerConfig == nil {
		panic(fmt.Errorf("please zlog.InitLogger then use"))
	}
	return *savedLoggerConfig
}

var savedLoggerConfig *LogsConfig

// InitLogger 初始化日志
//
//	logs:
//	  # 日志等级(-1:debug, 0:info, 1:warn, 2:error, 3:dpanic, 4:panic, 5:fatal, -1<=level<=5, 参照 zap.level 源码)
//	  level: info # 注意: debug 会对序列化为全大写字符串，并添加颜色，其他级别不被影响
//	  caller-hide: false # 是否完全隐藏文件和行号
//	  caller-disable-dynamic: false # 禁用动态级别行号显示, 默认会隐藏 debug info warn 的行号
//	  caller-full-path: false # 是否显示完整文件和行号
//	  # 是否开启 stdout 输出，只影响 2:Error zapcore.ErrorLevel 之前的输出, 正式建议关闭，全写在文件中
//	  stdout-enable: true
//	  # 日志文件 base 路径, 如果为空则 只输出到 stdout， 不为空，则自动区分 不同等级放置日志
//	  path-base: logs/log
//	  # 是否日志文件根目录使用执行程序所在的目录，默认使用当前运行的目录，开启后会拼接 path-base 的 base dir
//	  path-use-executable: false
//	  # 日志文件最大大小, 单位 M
//	  max-size: 50
//	  # 日志备份数
//	  max-backups: 7
//	  # 日志存放时间, 单位 天
//	  max-age: 30
//	  # 日志是否压缩
//	  compress: false
func InitLogger(config LogsConfig) error {

	copyNewConfig, errDeepCopyNew := config.DeepCopyNew()
	if errDeepCopyNew != nil {
		return errDeepCopyNew
	}

	logPathBase := config.PathBase
	if config.PathUseExecutable {
		execPath, errExecutable := os.Executable()
		if errExecutable == nil {
			nowConfigPathBase := filepath.Base(logPathBase)
			logPathBase = path.Join(path.Dir(execPath), nowConfigPathBase)
		}
	}
	copyNewConfig.PathBase = logPathBase

	var coreArr []zapcore.Core
	var accessArr []zapcore.Core
	var cronArr []zapcore.Core

	//var encoderConfig zapcore.EncoderConfig
	//if copyNewConfig.CallerHide {
	//	encoderConfig = zap.NewDevelopmentEncoderConfig()
	//} else {
	//	encoderConfig = zap.NewProductionEncoderConfig()
	//	if !copyNewConfig.CallerHide { // 单独处理
	//		encoderConfig.CallerKey = "" // 不显示 日志调用 路径
	//	}
	//}

	// 获取编码器 默认产品级编码器
	encoderConfig := zap.NewProductionEncoderConfig()

	if copyNewConfig.CallFullPath {
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder // 显示完整文件路径
	} else {
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // 显示短文件路径
	}
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // ，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 指定时间格式
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	// 打开 debug 时 使用调试编码器
	if copyNewConfig.Level == zapcore.DebugLevel {
		encoderConfig = zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			TimeKey:       "time",
			NameKey:       "name",
			CallerKey:     "file",
			FunctionKey:   "func",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalColorLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			//EncodeTime: zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式
			//EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			//	enc.AppendInt64(int64(d) / 1000000)
			//},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			//EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeCaller: zapcore.FullCallerEncoder, // 显示完整文件路径
			//EncodeName:       nil,
			//ConsoleSeparator: "",
		}
	}

	encoderConsole := zapcore.NewConsoleEncoder(encoderConfig)
	if !config.CallDisableDynamic {
		dynamicEncoder := &DynamicCallerEncoder{
			Encoder: encoderConsole,
		}
		encoderConsole = dynamicEncoder
	}

	levelAt := zap.NewAtomicLevelAt(copyNewConfig.Level)
	if logPathBase == "" { // only stdout
		stdoutCore := zapcore.NewCore(encoderConsole, zapcore.AddSync(os.Stdout), levelAt)
		coreArr = append(coreArr, stdoutCore)
		accessArr = append(accessArr, stdoutCore)
		cronArr = append(cronArr, stdoutCore)
	} else {
		now := time.Now()
		coreArr = coreLogArrInit(coreArr, logPathBase, *copyNewConfig, encoderConsole, now)
		accessArr = accessLogArrInit(accessArr, logPathBase, *copyNewConfig, encoderConsole, now)
		cronArr = cronLogArrInit(accessArr, logPathBase, *copyNewConfig, encoderConsole, now)
	}

	callerEnableOption := zap.WithCaller(!config.CallerHide)

	accessLogger := zap.New(zapcore.NewTee(accessArr...), zap.AddCaller(), callerEnableOption)
	accessLoggerGlobal = accessLogger

	cronLogger := zap.New(zapcore.NewTee(cronArr...), zap.AddCaller(), callerEnableOption)
	cronLoggerGlobal = cronLogger

	logger := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(), callerEnableOption)

	logSugared = logger.Sugar()

	savedLoggerConfig = copyNewConfig

	logSugared.Infof("initialize zap log complete, log path at: %s", savedLoggerConfig.PathBase)
	//logSugared.Errorf("initialize zap error case")

	return nil
}

func accessLogArrInit(logArr []zapcore.Core, logPathBase string, config LogsConfig, encoder zapcore.Encoder, now time.Time) []zapcore.Core {
	return zapLogArrInit(logArr, logPathBase, config, encoder, now, "access")
}

func cronLogArrInit(logArr []zapcore.Core, logPathBase string, config LogsConfig, encoder zapcore.Encoder, now time.Time) []zapcore.Core {
	return zapLogArrInit(logArr, logPathBase, config, encoder, now, "cron")
}

func zapLogArrInit(logArr []zapcore.Core, logPathBase string, config LogsConfig, encoder zapcore.Encoder, now time.Time, dirName string) []zapcore.Core {
	fileNameLogFileName := fmt.Sprintf("%s/%s/%s-%04d-%02d-%02d.log", logPathBase, dirName, dirName, now.Year(), now.Month(), now.Day())
	logFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileNameLogFileName, //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    config.MaxSize,      //文件大小限制,单位MB
		MaxAge:     config.MaxAge,       //日志文件保留天数
		MaxBackups: config.MaxBackups,   //最大保留日志文件数量
		LocalTime:  false,
		Compress:   config.Compress, //是否压缩处理
	})
	logWriteSyncer := zapcore.NewMultiWriteSyncer(logFileWriteSyncer)
	if config.StdoutEnable {
		logWriteSyncer = zapcore.NewMultiWriteSyncer(logFileWriteSyncer, zapcore.AddSync(os.Stdout))
	}
	logPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return true
	})
	loggerFileCore := zapcore.NewCore(encoder, logWriteSyncer, logPriority)
	logArr = append(logArr, loggerFileCore)
	return logArr
}

func coreLogArrInit(logArr []zapcore.Core, logPathBase string, config LogsConfig, encoder zapcore.Encoder, now time.Time) []zapcore.Core {
	infoLogFileName := fmt.Sprintf("%s/info/%04d-%02d-%02d.log", logPathBase, now.Year(), now.Month(), now.Day())
	errorLogFileName := fmt.Sprintf("%s/error/%04d-%02d-%02d.log", config.PathBase, now.Year(), now.Month(), now.Day())

	// 高日志级别
	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	})
	// 低日志级别
	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		//return level < zap.ErrorLevel && level >= zap.DebugLevel
		return level < zap.ErrorLevel && level >= config.Level
	})

	// 当yml配置中的等级大于Error时，lowPriority级别日志停止记录
	if config.Level >= 2 {
		lowPriority = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return false
		})
	}

	// info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   infoLogFileName,   //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    config.MaxSize,    //文件大小限制,单位MB
		MaxAge:     config.MaxAge,     //日志文件保留天数
		MaxBackups: config.MaxBackups, //最大保留日志文件数量
		LocalTime:  false,
		Compress:   config.Compress, //是否压缩处理
	})
	// 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	infoWriteSyncer := zapcore.NewMultiWriteSyncer(infoFileWriteSyncer)
	if config.StdoutEnable {
		infoWriteSyncer = zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout))
	}

	infoFileCore := zapcore.NewCore(encoder, infoWriteSyncer, lowPriority)

	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   errorLogFileName,  //日志文件存放目录
		MaxSize:    config.MaxSize,    //文件大小限制,单位MB
		MaxAge:     config.MaxAge,     //日志文件保留天数
		MaxBackups: config.MaxBackups, //最大保留日志文件数量
		LocalTime:  false,
		Compress:   config.Compress, //是否压缩处理
	})

	// 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	logArr = append(logArr, infoFileCore)
	logArr = append(logArr, errorFileCore)

	return logArr
}
