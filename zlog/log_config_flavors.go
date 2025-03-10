package zlog

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type LogsConfigFlavors struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	LogsConfig
	ConfigFlavorsDeepCopy
}

type ConfigFlavorsDeepCopy interface {
	DeepCopyFromConfig(name string, src LogsConfig) error
	DeepCopyToConfig() (cfg *LogsConfig, err error)

	DeepCopyFrom(src LogsConfigFlavors) error
	DeepCopyNew() (*LogsConfigFlavors, error)
}

//	deep copy from LogsConfig
//
// @return error
func (l *LogsConfigFlavors) DeepCopyFromConfig(name string, src LogsConfig) error {
	if name == "" {
		return fmt.Errorf("new flavors name is empty")
	}
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}

	errDecoder := gob.NewDecoder(&buffer).Decode(l)
	if errDecoder != nil {
		return errDecoder
	}
	l.Name = name
	return nil
}

func (l *LogsConfigFlavors) DeepCopyToConfig() (cfg *LogsConfig, err error) {
	var buffer bytes.Buffer
	if errNewEncoder := gob.NewEncoder(&buffer).Encode(l.LogsConfig); errNewEncoder != nil {
		err = errNewEncoder
		return
	}
	var newLogsConfig LogsConfig
	if errNewDecoder := gob.NewDecoder(&buffer).Decode(&newLogsConfig); errNewDecoder != nil {
		err = errNewDecoder
		return
	}
	cfg = &newLogsConfig
	return
}

//	deep copy from src
//
// @return error
func (l *LogsConfigFlavors) DeepCopyFrom(src LogsConfigFlavors) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buffer).Decode(l)
}

//	deep copy new instance
//
// @return *LogsConfig, error
func (l *LogsConfigFlavors) DeepCopyNew() (*LogsConfigFlavors, error) {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(l); err != nil {
		return nil, err
	}
	var newLogsConfig LogsConfigFlavors
	if err := gob.NewDecoder(&buffer).Decode(&newLogsConfig); err != nil {
		return nil, err
	}
	return &newLogsConfig, nil
}
