package zlog

import "encoding/gob"

func init() {
	gob.Register(LogsConfig{})
	gob.Register(LogsConfigFlavors{})
}
