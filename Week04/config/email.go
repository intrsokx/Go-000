package config

import "encoding/json"

type EmailCfg struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Sender     string   `json:"sender"`
	SenderName string   `json:"sender_name"`
	Password   string   `json:"password"`
	Receiver   []string `json:"receiver"`
}

var ECfg *EmailCfg

func InitECfg(db []byte) {
	ECfg = &EmailCfg{}
	if e := json.Unmarshal(db, ECfg); e != nil {
		panic(e)
	}
}

func NewEmailCfg() *EmailCfg {
	return ECfg
}
