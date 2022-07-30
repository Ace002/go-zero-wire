package logic

import (
	"context"

	"service/internal/user/domain/user"
)

type Logic interface {
	Create(context.Context, int64) error
	Find(context.Context, int64) (*user.User, error)
}

type LogicImpl struct {
	Repo user.Repository
}

func NewLogic(repo user.Repository) Logic {
	return &LogicImpl{
		Repo: repo,
	}
}

// Create ...
func (l *LogicImpl) Create(ctx context.Context, num int64) (err error) {
	user, err := user.NewUser(num)
	if err != nil {
		return
	}

	err = l.Repo.Create(ctx, user)
	return
}

// Find ...
func (l *LogicImpl) Find(ctx context.Context, num int64) (usr *user.User, err error) {
	usr, err = l.Repo.Find(ctx, num)
	return
}
