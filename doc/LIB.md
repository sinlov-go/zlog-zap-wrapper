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
)
```

### config yml

```yml
logs:
  # 日志等级(-1:Debug, 0:Info, 1:Warn, 2:Error, 3:DPanic, 4:Panic, 5:Fatal, -1<=level<=5, 参照 zap. level 源码)
  level: 0
  # 是否开启 stdout 输出，只影响 Error 之前的输出, 正式建议关闭，全写在文件中
  stdout-enable: true
  # 日志文件 base 路径, 如果为空则 只输出到 stdout， 不为空，则自动区分 不同等级放置日志
  path-base: logs/log
  # 是否日志文件根目录使用执行程序所在的目录，默认使用当前运行的目录
  path-use-executable: false
  # 日志文件最大大小, M
  max-size: 50
  # 日志备份数
  max-backups: 7
  # 日志存放时间, 天
  max-age: 30
  # 日志是否压缩
  compress: false
```