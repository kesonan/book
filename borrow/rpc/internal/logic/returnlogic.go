package logic

import (
	"book/borrow/rpc/model"
	"book/shared"
	"context"

	"book/borrow/rpc/internal/svc"
	borrow "book/borrow/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnLogic {
	return &ReturnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 还书
func (l *ReturnLogic) Return(in *borrow.ReturnReq) (*borrow.BaseReply, error) {
	info, err := l.svcCtx.BorrowSystemModel.FindOneByBookNo(in.BookNo)
	switch err {
	case nil:
		switch info.Status {
		case model.Borrowing:
			info.Status = model.Return
			_, err = l.svcCtx.BorrowSystemModel.Update(*info)
			if err != nil {
				return nil, shared.NewGRPCErrorFromError(err)
			}
			return &borrow.BaseReply{}, nil
		default:
			return nil, shared.NewDefaultGRPCError("该图书已归还")
		}
	case model.ErrNotFound:
		return nil, shared.NewDefaultGRPCError("该图书未被借阅")
	default:
		return nil, shared.NewGRPCErrorFromError(err)
	}
}
