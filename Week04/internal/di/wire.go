// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"github.com/intrsokx/Go-000/Week04/config"
	"github.com/intrsokx/Go-000/Week04/internal/dao"
	"github.com/intrsokx/Go-000/Week04/internal/service"
	"github.com/intrsokx/Go-000/Week04/pkg/email"
)

//di 依赖注入
//声明依赖注入文件
func NewService() (*service.Service, error) {
	wire.Build(service.NewService,
		config.NewDbCfg,
		config.NewEmailCfg,
		email.EmailSet,
		dao.UserSet)
	return &service.Service{}, nil
}
