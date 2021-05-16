package service

import (
	"github.com/intrsokx/Go-000/Week04/internal/dao"
	"github.com/intrsokx/Go-000/Week04/pkg/email"
)

//处理业务
type Service struct {
	m email.IEmail
	d dao.UserRepository
}

func NewService(m email.IEmail, d dao.UserRepository) *Service {
	return &Service{
		m: m,
		d: d,
	}
}

func (s *Service) UserRegister() {
	//TODO 这里应该生成一个service层的用户对象，这个对象应该是全局的一个base user对象， 然后传下去，再由dao层生成特定的repo user对象
	s.d.AddUser()
	s.m.Send("用户注册成功")
}
