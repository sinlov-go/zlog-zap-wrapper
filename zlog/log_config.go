package zlog

import (
	"bytes"
	"encoding/gob"
	"go.uber.org/zap/zapcore"
	"os"
)

func LogsConfigDefault() LogsConfig {
	return LogsConfig{
		Level:              zapcore.InfoLevel,
		CallerHide:         false,
		CallDisableDynamic: false,
		CallFullPath:       false,

		StdoutEnable:      true,
		PathBase:          "logs/log",
		PathUseExecutable: false,
		MaxSize:           10,
		MaxBackups:        7,
		MaxAge:            360,
		Compress:          true,
	}
}

func LogsConfigProduction() LogsConfig {
	return LogsConfig{
		Level:              zapcore.WarnLevel,
		CallerHide:         false,
		CallDisableDynamic: false,
		CallFullPath:       true,

		StdoutEnable:      false,
		PathBase:          "logs/log",
		PathUseExecutable: false,
		MaxSize:           10,
		MaxBackups:        7,
		MaxAge:            720,
		Compress:          true,
	}
}

func LogsConfigDebug() LogsConfig {
	return LogsConfig{
		Level:              zapcore.DebugLevel,
		CallerHide:         false,
		CallDisableDynamic: false,
		CallFullPath:       true,

		StdoutEnable:      true,
		PathBase:          "logs/log",
		PathUseExecutable: false,
		MaxSize:           10,
		MaxBackups:        7,
		MaxAge:            30,
		Compress:          false,
	}
}

type LogsConfig struct {
	ConfigDeepCopy

	// log level by zapcore.Level
	// Note: debug will serialize the string to all uppercase and add color. Other levels will not be affected.
	Level zapcore.Level `mapstructure:"level" json:"level" yaml:"level"`
	// whether to completely hide file and line numbers
	CallerHide bool `mapstructure:"caller-hide" json:"callerHide" yaml:"caller-hide"`
	// Disable the display of dynamic level line numbers. By default, the line numbers of debug info warn are hidden.
	CallDisableDynamic bool `mapstructure:"caller-disable-dynamic" json:"callDisableDynamic" yaml:"caller-disable-dynamic"`
	// Whether to enable full path output,
	CallFullPath bool `mapstructure:"caller-full-path" json:"callFullPath" yaml:"caller-full-path"`

	// Whether to enable stdout output, only affects the output before the zapcore.ErrorLevel, the formal recommendation is to disable, all written in the file
	StdoutEnable bool `mapstructure:"stdout-enable" json:"stdoutEnable" yaml:"stdout-enable"`
	// The base path of the log file. If it is empty, the log file is output to stdout. If it is not empty, the log file is automatically placed at different levels.
	PathBase string `mapstructure:"path-base" json:"pathBase" yaml:"path-base"`
	// Whether the root directory of the log file uses the directory where the execution program is located. By default, the current running directory is used. After opening, the base dir of path-base will be spliced.
	PathUseExecutable bool `mapstructure:"path-use-executable" json:"pathUseExecutable" yaml:"path-use-executable"`
	// maximum log file size in m
	MaxSize int `mapstructure:"max-size" json:"maxSize" yaml:"max-size"`
	// number of log backups
	MaxBackups int `mapstructure:"max-backups" json:"maxBackups" yaml:"max-backups"`
	// log storage time unit day
	MaxAge int `mapstructure:"max-age" json:"maxAge" yaml:"max-age"`
	// whether the log is compressed
	Compress bool `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type ConfigDeepCopy interface {
	DeepCopyFrom(src LogsConfig) error
	DeepCopyNew() (*LogsConfig, error)
	PruneLogs() (string, error)
}

//	deep copy from src
//
// @return error
func (l *LogsConfig) DeepCopyFrom(src LogsConfig) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buffer).Decode(l)
}

//	deep copy new instance
//
// @return *LogsConfig, error
func (l *LogsConfig) DeepCopyNew() (*LogsConfig, error) {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(l); err != nil {
		return nil, err
	}
	var newLogsConfig LogsConfig
	if err := gob.NewDecoder(&buffer).Decode(&newLogsConfig); err != nil {
		return nil, err
	}
	return &newLogsConfig, nil
}

//	remove logs by config
//
//	@return string, error
//
// "" and nil if no error is can not be pruned by folder not exits
func (l *LogsConfig) PruneLogs() (string, error) {
	exists, errPathExits := pathExists(l.PathBase)
	if errPathExits != nil || !exists {
		return "", nil
	}

	errRemoveAll := os.RemoveAll(l.PathBase)
	if errRemoveAll != nil {
		return l.PathBase, errRemoveAll
	}

	return l.PathBase, nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
