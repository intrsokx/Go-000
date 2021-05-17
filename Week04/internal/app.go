package internal

import (
	"github.com/intrsokx/Go-000/Week04/config"
	"github.com/intrsokx/Go-000/Week04/internal/di"
	"github.com/intrsokx/Go-000/Week04/internal/service"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

/**
我们知道golang 的相对路径是相对于执行命令时的目录。
因此在本项目主，启动项目时的目录机构应如下所示。
			.
			├── cmd
			│   └── main
			└── resource
				├── db.json
				└── email.json
*/
func initDbCfg() {
	b, e := ioutil.ReadFile("../resource/db.json")
	if e != nil {
		panic(e)
	}
	config.InitDCfg(b)
	logrus.Debugf("%#v", config.DCfg)
}

func initEmailCfg() {
	b, e := ioutil.ReadFile("../resource/email.json")
	if e != nil {
		panic(e)
	}
	config.InitECfg(b)
	logrus.Debugf("%#v", config.ECfg)
}

func initLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.SetReportCaller(true)
	if config.BaseCfg.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func loadBaseCfg() {
	b, e := ioutil.ReadFile("../resource/config.json")
	if e != nil {
		panic(e)
	}
	config.InitBaseCfg(b)
}

func init() {
	//load base config
	loadBaseCfg()
	initLog()
	initEmailCfg()
	initDbCfg()
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
