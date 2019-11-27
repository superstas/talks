package code

import "context"

// 2 OMIT
type UserStorageMock struct{}

func (m UserStorageMock) User(_ context.Context, _ int) (*User, error) {
	return &User{}, nil
}

func (m UserStorageMock) Create(_ context.Context, _ *User) error {
	return nil
}

func (m UserStorageMock) Update(_ context.Context, _ *User) error {
	return nil
}

func (m UserStorageMock) Delete(_ context.Context, _ int) error {
	return nil
}

// END 2 OMIT
