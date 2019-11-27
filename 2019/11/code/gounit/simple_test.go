package gounit

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

// 0 OMIT
func TestService_AddUser(t *testing.T) {
	type args struct {
		ctx context.Context
		u   *User
	}
	tests := []struct {
		name       string
		init       func(t minimock.Tester) *Service
		args       func(t minimock.Tester) args
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation
	}{
		//TODO: Add test cases // HL
	}
	//...
	// END 0 OMIT
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			defer mc.Wait(time.Second)

			tArgs := tt.args(mc)
			receiver := tt.init(mc)

			err := receiver.AddUser(tArgs.ctx, tArgs.u)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if tt.wantErr {
				if assert.Error(t, err) && tt.inspectErr != nil {
					tt.inspectErr(err, t)
				}
			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestService_AddUser1(t *testing.T) {
	type args struct {
		ctx context.Context
		u   *User
	}
	tests := []struct {
		name    string
		init    func(t minimock.Tester) *Service
		inspect func(r *Service, t *testing.T) //inspects *Service after execution of AddUser

		args func(t minimock.Tester) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation
	}{

		// 1 OMIT
		{
			name: "OK case",
			init: func(t minimock.Tester) *Service {
				return NewService(
					NewUserStorageMock(t).
						UserByEmailMock.
						Expect(context.Background(), "test@test.com"). // HL
						Return(nil, ErrUserNotFound).                  // HL

						CreateMock.Return(nil), // HL
				)
			},
			args: func(t minimock.Tester) args {
				return args{context.Background(), &User{Email: "test@test.com"}} // HL
			},
			wantErr: false, // HL
		},
		// END 1 OMIT
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			defer mc.Wait(time.Second)

			tArgs := tt.args(mc)
			receiver := tt.init(mc)

			err := receiver.AddUser(tArgs.ctx, tArgs.u)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if tt.wantErr {
				if assert.Error(t, err) && tt.inspectErr != nil {
					tt.inspectErr(err, t)
				}
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
