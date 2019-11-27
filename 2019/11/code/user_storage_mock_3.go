package code

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
)

type User struct{}

// 5 OMIT
type UserStorageMock struct {
	t               *testing.T
	UserMock        func(ctx context.Context, id int) (*User, error)
	UserMockCounter uint64 // HL
	//...
}

func (m UserStorageMock) User(ctx context.Context, id int) (u *User, err error) {
	if m.UserMock != nil {
		atomic.AddUint64(&m.UserMockCounter, 1) // HL
		return m.UserMock(ctx, id)
	}

	m.t.Fatal("Unexpected call of UserStorageMock.User")
	return
}

// END 5 OMIT

func Test_Something(t *testing.T) {

	// 6 OMIT
	userStorageMock := UserStorageMock{
		t: t,
		UserMock: func(ctx context.Context, id int) (*User, error) { // HL
			// ...
			return nil, errors.New("something is wrong")
		},
	}

	service := NewService(userStorageMock)
	// ...
	if userStorageMock.UserMockCounter != 1 {
		t.Errorf("UserStorageMock.User invoked %d times; expected: %d", userStorageMock.UserMockCounter, 1)
	}
	// END 6 OMIT
}
