package svc

import (
	"book/borrow/api/internal/config"
	"book/borrow/model"
	"book/library/rpc/library"
	"book/user/rpc/user"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	BorrowSystemModel *model.BorrowSystemModel
	UserRpc           user.User
	LibraryRpc        library.Library
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	ur := user.NewUser(zrpc.MustNewClient(c.UserRpc))
	lr := library.NewLibrary(zrpc.MustNewClient(c.LibraryRpc))
	return &ServiceContext{
		Config:            c,
		BorrowSystemModel: model.NewBorrowSystemModel(conn),
		UserRpc:           ur,
		LibraryRpc:        lr,
	}
}
