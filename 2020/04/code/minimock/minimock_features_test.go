package minimock

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func Test_OneSender(t *testing.T) {
	// 0 OMIT
	senderMock := NewSenderMock(t).SendMock.Return(nil) // HL
	defer senderMock.MinimockFinish()

	err := NewMultiSender(senderMock).Send(context.Background(), "to", []byte{})
	assert.NoError(t, err)
	// END 0 OMIT
}

func Test_Sender(t *testing.T) {
	// 1 OMIT
	tcases := []struct {
		name, to   string
		msg        []byte
		senderMock *SenderMock
		wantErr    error
	}{
		{
			senderMock: NewSenderMock(t).SendMock.Return(nil),
			wantErr:    nil,
		},
		{
			senderMock: NewSenderMock(t).SendMock.Return(io.EOF),
			wantErr:    io.EOF,
		},
	}

	for _, tc := range tcases {
		// snipped
		// END 1 OMIT
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			err := NewMultiSender(tc.senderMock).Send(ctx, tc.to, tc.msg)

			if tc.wantErr != nil {
				assert.EqualError(t, tc.wantErr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			tc.senderMock.MinimockFinish()
		})
	}
}

func Test_OneSender_Expect(t *testing.T) {
	ctx := context.Background()
	to, msg := "to", []byte{}
	// 2 OMIT

	senderMock := NewSenderMock(t).SendMock.Expect(ctx, to, msg).Return(nil) // HL

	err := NewMultiSender(senderMock).Send(ctx, to, msg)
	// SenderMock.Send got unexpected parameters...

	assert.NoError(t, err)
	assert.EqualValues(t, 1, senderMock.SendAfterCounter())
	// END 2 OMIT
}

func Test_TwoSenders(t *testing.T) {
	to, msg := "123", []byte("msg")
	wantErr := errors.New("something happened")
	ctx := context.Background()

	// 3 OMIT
	mc := minimock.NewController(t) // HL
	defer mc.Finish()               // HL

	sender1 := NewSenderMock(mc).SendMock.Expect(ctx, to, msg).Return(nil)
	sender2 := NewSenderMock(mc).SendMock.Expect(ctx, to, msg).Return(wantErr)

	err := NewMultiSender(sender1, sender2).Send(ctx, to, msg)
	// END 3 OMIT
	assert.EqualError(t, wantErr, err.Error())

}

func Test_WantThen(t *testing.T) {
	wantErr := errors.New("something happened")
	ctx := context.Background()
	msg := []byte{}

	// 4 OMIT
	mockSender := NewSenderMock(t).
		SendMock.When(ctx, "1", msg).Then(nil).
		SendMock.When(ctx, "2", msg).Then(wantErr)

	err := NewMultiSender(mockSender).Send(ctx, "1", msg)
	assert.NoError(t, err)

	err = NewMultiSender(mockSender).Send(ctx, "2", msg)
	assert.EqualError(t, wantErr, err.Error())
	// END 4 OMIT

	assert.EqualValues(t, 2, mockSender.SendAfterCounter())

}
