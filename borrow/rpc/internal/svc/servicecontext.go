package svc

import (
	"book/borrow/rpc/internal/config"
	"book/borrow/rpc/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c                 config.Config
	BorrowSystemModel *model.BorrowSystemModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		c:                 c,
		BorrowSystemModel: model.NewBorrowSystemModel(conn),
	}
}
