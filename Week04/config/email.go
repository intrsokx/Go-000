package config

import "encoding/json"

type EmailCfg struct {
	Receiver []MailUser `json:"receiver_list"`
}

type MailUser struct {
	Address string `json:"address"`
}

var ECfg *EmailCfg

func InitECfg(db []byte) error {
	return json.Unmarshal(db, ECfg)
}

func NewEmailCfg() *EmailCfg {
	return ECfg
}
