package user

import "errors"

type User struct {
	ID     int64 `json:"-"`
	Number int64 `json:"number"`
}

var (
	// ErrInvalidNumber ...
	ErrInvalidNumber = errors.New("Number must be greater than 0")
)

// NewUser ...
func NewUser(number int64) (*User, error) {
	if number <= 0 {
		return nil, ErrInvalidNumber
	}

	return &User{
		Number: number,
	}, nil
}
