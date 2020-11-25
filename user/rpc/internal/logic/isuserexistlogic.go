package logic

import (
	"book/user/model"
	"book/user/rpc/user"
	"context"

	"book/user/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsUserExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsUserExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUserExistLogic {
	return &IsUserExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 判断用户是否存在
func (l *IsUserExistLogic) IsUserExist(in *user.UserExistReq) (*user.UserExistReply, error) {
	_, err := l.svcCtx.UserModel.FindOne(in.Id)
	switch err {
	case nil:
		return &user.UserExistReply{Exists: true}, nil
	case model.ErrNotFound:
		return &user.UserExistReply{}, nil
	default:
		return nil, err
	}

}
