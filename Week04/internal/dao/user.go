package dao

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/intrsokx/Go-000/Week04/config"
)

type UserRepository interface {
	AddUser()
	DeleteUser(id int)
}

type UserRepo struct {
	*sql.DB
	cfg *config.DbCfg
}

func (u *UserRepo) AddUser() {
	panic("implement me")
}

func (u *UserRepo) DeleteUser(id int) {
	panic("implement me")
}

func NewUserRepo(cfg *config.DbCfg) *UserRepo {
	return &UserRepo{cfg: cfg}
}

//TODO 确认含义
//声明NewUserRepo的返回值是UserRepository接口类型
var UserSet = wire.NewSet(NewUserRepo, wire.Bind(new(UserRepository), new(*UserRepo)))
