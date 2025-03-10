## code

```go
package config

func InitLogger(zLogsConfig *zlog.LogsConfig)  (error) {
	errInitZlog := zlog.InitLogger(*zLogsConfig)
	if errInitZlog != nil {
		return errInitZlog
	}
}
```

## load by viper

### viper libs 

- go.mod

```go
package foo

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/sinlov-go/zlog-zap-wrapper/zlog"
)

func exampleInit() {
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
```

### config yml

```yml
logs:
  # 日志等级(-1:debug, 0:info, 1:warn, 2:error, 3:dpanic, 4:panic, 5:fatal, -1<=level<=5, 参照 zap.level 源码)
  level: info # 注意: debug 会对序列化为全大写字符串，并添加颜色，其他级别不被影响
  caller-hide: false # 是否完全隐藏文件和行号
  caller-disable-dynamic: false # 禁用动态级别行号显示, 默认会隐藏 debug info warn 的行号
  caller-full-path: false # 是否显示完整文件和行号
  # 是否开启 stdout 输出，只影响 2:Error zapcore.ErrorLevel 之前的输出, 正式建议关闭，全写在文件中
  stdout-enable: true
  # 日志文件 base 路径, 如果为空则 只输出到 stdout， 不为空，则自动区分 不同等级放置日志
  path-base: logs/log
  # 是否日志文件根目录使用执行程序所在的目录，默认使用当前运行的目录，开启后会拼接 path-base 的 base dir
  path-use-executable: false
  # 日志文件最大大小, 单位 M
  max-size: 50
  # 日志备份数
  max-backups: 7
  # 日志存放时间, 单位 天
  max-age: 30
  # 日志是否压缩
  compress: false
```