package internal

import (
	"github.com/intrsokx/Go-000/Week04/config"
	"github.com/intrsokx/Go-000/Week04/internal/di"
	"github.com/intrsokx/Go-000/Week04/internal/service"
	"io/ioutil"
)

func initDbCfg() {
	f, err := ioutil.ReadFile("../resource/db.json")
	if err != nil {
		panic(err)
	}
	config.InitDCfg(f)
}

func initEmailCfg() {
	f, err := ioutil.ReadFile("../resource/email.json")
	if err != nil {
		panic(err)
	}
	config.InitECfg(f)
}

func init() {
	initDbCfg()
	initEmailCfg()
}

type app struct {
	service *service.Service
}

func NewApp() *app {
	return &app{}
}

func (a *app) Run() error {
	s, err := di.NewService()
	if err != nil {
		return err
	}
	s.UserRegister()
	return nil
}
