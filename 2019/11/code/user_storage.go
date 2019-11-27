package code

import "context"

type User struct {
	ID   int
	Name string
	Age  string
}

// 0 OMIT
type UserStorage interface {
	User(ctx context.Context, ID int) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, ID int) error
}

// END 0 OMIT

// 1 OMIT
type Service struct {
	// ...
	userStorage UserStorage // HL
	// ...
}

func NewService(u UserStorage) *Service {
	return &Service{
		userStorage: u, // HL
	}
}

// END 1 OMIT
