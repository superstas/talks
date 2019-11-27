package code

import (
	"context"
	"errors"
	"testing"
)

type User struct{}

// 3 OMIT
type UserStorageMock struct {
	t        *testing.T
	UserMock func(ctx context.Context, id int) (*User, error) // HL
	//...
}

func (m UserStorageMock) User(ctx context.Context, id int) (u *User, err error) {
	if m.UserMock != nil {
		return m.UserMock(ctx, id) // HL
	}

	m.t.Fatal("Unexpected call of UserStorageMock.User")
	return
}

// END 3 OMIT

func Test_Something(t *testing.T) {

	// 4 OMIT
	userStorageMock := UserStorageMock{
		UserMock: func(ctx context.Context, id int) (*User, error) { // HL
			// ...
			return nil, errors.New("something is wrong")
		},
	}

	service := NewService(userStorageMock)
	// ...
	// END 4 OMIT
}
