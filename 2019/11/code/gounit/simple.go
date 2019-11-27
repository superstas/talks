package gounit

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID    int
	Name  string
	Email string
	Age   string
}

// 0 OMIT
type UserStorage interface {
	UserByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, ID int) error
}

type Service struct {
	userStorage UserStorage
}

// END 0 OMIT

func NewService(u UserStorage) *Service {
	return &Service{
		userStorage: u,
	}
}

// 1 OMIT
// AddUser checks if the user exists and creates it if not.
func (s *Service) AddUser(ctx context.Context, u *User) error {
	_, err := s.userStorage.UserByEmail(ctx, u.Email) // HL
	if err != nil && err != ErrUserNotFound {
		return err
	}

	return s.userStorage.Create(ctx, u) // HL
}

// END 1 OMIT
