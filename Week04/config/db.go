package config

import "encoding/json"

type DbCfg struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DataBase string `json:"data_base"`
}

var DCfg *DbCfg

func InitDCfg(db []byte) {
	DCfg = &DbCfg{}
	if e := json.Unmarshal(db, DCfg); e != nil {
		panic(e)
	}
}

func NewDbCfg() *DbCfg {
	return DCfg
}
