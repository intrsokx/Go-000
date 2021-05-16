package email

import (
	"github.com/google/wire"
	"github.com/intrsokx/Go-000/Week04/config"
	"net/smtp"
)

type IEmail interface {
	Send(msg string)
}
type email struct {
	smtp.Client
	cfg *config.EmailCfg
}

func (e *email) Send(msg string) {
	//TODO send msg
}

func NewEmail(cfg *config.EmailCfg) *email {
	return &email{
		//Client: smtp.Client{},
		cfg: cfg,
	}
}

// MailSet 声明NewEmail的返回值是email接口类型
var EmailSet = wire.NewSet(NewEmail, wire.Bind(new(IEmail), new(*email)))
