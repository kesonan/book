package logic

import (
	"book/borrow/api/internal/svc"
	"book/borrow/api/internal/types"
	"book/borrow/model"
	"book/library/rpc/library"
	"book/shared"
	"book/user/rpc/user"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type BorrowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBorrowLogic(ctx context.Context, svcCtx *svc.ServiceContext) BorrowLogic {
	return BorrowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BorrowLogic) Borrow(userId string, req types.BorrowReq) error {
	userInt, err := strconv.ParseInt(fmt.Sprintf("%v", userId), 10, 64)
	if err != nil {
		return err
	}

	if req.ReturnPlan < time.Now().Unix() {
		return errInvalidParam
	}

	reply, err := l.svcCtx.UserRpc.IsUserExist(l.ctx, &user.UserExistReq{Id: userInt})
	if err != nil { // code error
		// 这里判断not found是为了有些业务场景需要使用到not found,然后进行数据更新
		// 当前业务其实可以直接返回error
		if shared.IsGRPCNotFound(err) {
			return errUserNotFound
		}
		return err
	}

	if !reply.Exists {
		return errUserNotFound
	}

	book, err := l.svcCtx.LibraryRpc.FindBookByName(l.ctx, &library.FindBookReq{Name: req.BookName})
	if err != nil { // code error
		if shared.IsGRPCNotFound(err) {
			return errBookNotFound
		}
		return err
	}

	_, err = l.svcCtx.BorrowSystemModel.FindOneBookNo(book.No, model.Borrowing)
	switch err {
	case nil:
		return errBookBorrowed
	case model.ErrNotFound:
		_, err = l.svcCtx.BorrowSystemModel.Insert(model.BorrowSystem{
			BookNo:         book.No,
			UserId:         userInt,
			Status:         model.Borrowing,
			ReturnPlanDate: time.Unix(req.ReturnPlan, 0),
		})
		return err
	default:
		return err
	}
}
