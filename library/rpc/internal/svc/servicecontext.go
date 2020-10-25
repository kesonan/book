package svc

import (
	"book/library/rpc/internal/config"
	"book/library/rpc/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c            config.Config
	LibraryModel *model.LibraryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		c:            c,
		LibraryModel: model.NewLibraryModel(conn),
	}
}
