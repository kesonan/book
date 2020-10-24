package svc

import (
	"book/user/api/internal/config"
	"book/user/api/internal/middleware"
	"book/user/api/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserModel
	UserCheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	um := model.NewUserModel(conn)
	return &ServiceContext{
		Config:    c,
		UserModel: um,
		UserCheck: middleware.NewUserCheckMiddleware().Handle,
	}
}
