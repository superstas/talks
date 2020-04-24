package code

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
)

// 5 OMIT
type UserStorageMock struct {
	t                 *testing.T
	UserMock          func(ctx context.Context, id int) (*User, error)
	UserMockCounter   uint64
	CreateMock        func(context.Context, *User) error
	CreateMockCounter uint64
	UpdateMock        func(context.Context, *User) error
	UpdateMockCounter uint64
	DeleteMock        func(ctx context.Context, id int) error
	DeleteMockCounter uint64
}

func (m UserStorageMock) User(ctx context.Context, id int) (u *User, err error) {
	if m.UserMock != nil {
		atomic.AddUint64(&m.UserMockCounter, 1) // HL
		return m.UserMock(ctx, id)              // HL
	}

	m.t.Fatal("Unexpected call of UserStorageMock.User")
	return
}
func (m UserStorageMock) Create(ctx context.Context, u *User) error { /* snipped */ }
func (m UserStorageMock) Update(ctx context.Context, u *User) error { /* snipped */ }
func (m UserStorageMock) Delete(ctx context.Context, id int) error  { /* snipped */ }

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
