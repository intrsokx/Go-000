package config

import "encoding/json"

type DbCfg struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DataBase string `json:"data_base"`
}

var DCfg *DbCfg

func InitDCfg(db []byte) error {
	return json.Unmarshal(db, DCfg)
}

func NewDbCfg() *DbCfg {
	return DCfg
}
