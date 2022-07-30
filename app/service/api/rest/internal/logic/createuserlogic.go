package logic

import (
	"context"

	"service/api/rest/internal/types"
	userLogic "service/internal/user/logic"
)

type CreateUserLogic struct {
	// logx.Logger
	// ctx       context.Context
	// svcCtx    *svc.ServiceContext
	userLogic userLogic.Logic
}

func NewCreateUserLogic( /*ctx context.Context,  svcCtx *svc.ServiceContext, */ userLogic userLogic.Logic) *CreateUserLogic {
	return &CreateUserLogic{
		// Logger:    logx.WithContext(ctx),
		// ctx:       ctx,
		// svcCtx:    svcCtx,
		userLogic: userLogic,
	}
}

func (l *CreateUserLogic) CreateUser(ctx context.Context, req *types.CreateReq) error {
	return l.userLogic.Create(ctx, req.Number)
}
