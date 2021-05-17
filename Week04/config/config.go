package config

import "encoding/json"

type BaseConfig struct {
	Debug bool `json:"debug"`
}

var BaseCfg *BaseConfig

func InitBaseCfg(cfg []byte) {
	BaseCfg = &BaseConfig{}
	if e := json.Unmarshal(cfg, BaseCfg); e != nil {
		panic(e)
	}
}

func NewBaseCfg() *BaseConfig {
	return BaseCfg
}
