package logic

import (
	"context"

	"service/api/rest/internal/types"
	userLogic "service/internal/user/logic"
)

type GetUserLogic struct {
	// logx.Logger
	// ctx       context.Context
	// svcCtx    *svc.ServiceContext
	userLogic userLogic.Logic
}

func NewGetUserLogic( /*ctx context.Context,  svcCtx *svc.ServiceContext, */ userLogic userLogic.Logic) *GetUserLogic {
	return &GetUserLogic{
		// Logger:    logx.WithContext(ctx),
		// ctx:       ctx,
		// svcCtx:    svcCtx,
		userLogic: userLogic,
	}
}

func (l *GetUserLogic) GetUser(ctx context.Context, req *types.GetReq) (resp *types.GetResp, err error) {
	user, err := l.userLogic.Find(ctx, req.Id)
	resp = &types.GetResp{
		Numbers: user.Number,
	}
	return
}
