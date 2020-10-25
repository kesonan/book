package logic

import (
	"book/borrow/rpc/model"
	"book/shared"
	"context"
	"time"

	"book/borrow/rpc/internal/svc"
	borrow "book/borrow/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type BorrowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBorrowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BorrowLogic {
	return &BorrowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 借书
func (l *BorrowLogic) Borrow(in *borrow.BorrowReq) (*borrow.BaseReply, error) {
	_, err := l.svcCtx.BorrowSystemModel.FindOneByBookNoAndStatus(in.BookNo, model.Borrowing)
	switch err {
	case nil:
		return nil, shared.NewDefaultGRPCError("该图书已被其他人借阅")
	case model.ErrNotFound:
		_, err := l.svcCtx.BorrowSystemModel.Insert(model.BorrowSystem{
			BookNo:         in.BookNo,
			UserId:         in.UserId,
			Status:         model.Borrowing,
			ReturnPlanDate: time.Unix(in.ReturnPlanDate, 0),
		})
		if err != nil {
			return nil, shared.NewGRPCErrorFromError(err)
		}
		return &borrow.BaseReply{}, nil
	default:
		return nil, shared.NewGRPCErrorFromError(err)
	}
}
