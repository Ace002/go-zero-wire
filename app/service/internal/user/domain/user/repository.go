package user

import "context"

type Repository interface {
	Create(context.Context, *User) error
	Find(context.Context, int64) (*User, error)
}
