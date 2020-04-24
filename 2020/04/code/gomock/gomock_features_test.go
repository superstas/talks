package gomock

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_OneSender(t *testing.T) {
	// 0 OMIT
	c := gomock.NewController(t)
	defer c.Finish()

	senderMock := NewMockSender(c)                                                 // HL
	senderMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil) // HL

	err := NewMultiSender(senderMock).Send(context.Background(), "to", []byte{})
	assert.NoError(t, err)
	// END 0 OMIT
}

func Test_OneSender_Expect(t *testing.T) {
	ctx := context.Background()
	to, msg := "to", []byte{}

	// 1 OMIT
	c := gomock.NewController(t)
	defer c.Finish()

	senderMock := NewMockSender(c)
	senderMock.EXPECT().Send(ctx, to, msg).Return(nil) // HL
	//senderMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	err := NewMultiSender(senderMock).Send(ctx, to, msg)
	assert.NoError(t, err)

	// END 1 OMIT
}

func Test_TwoSenders(t *testing.T) {
	to, msg := "123", []byte("msg")
	wantErr := errors.New("something happened")
	ctx := context.Background()

	// 3 OMIT
	c := gomock.NewController(t) // HL
	defer c.Finish()             // HL

	sender1 := NewMockSender(c)
	sender1.EXPECT().Send(ctx, to, msg).Return(nil)

	sender2 := NewMockSender(c)
	sender2.EXPECT().Send(ctx, to, msg).Return(wantErr)

	err := NewMultiSender(sender1, sender2).Send(ctx, to, msg)
	// END 3 OMIT
	assert.EqualError(t, wantErr, err.Error())
}

func Test_WantThen(t *testing.T) {
	wantErr := errors.New("something happened")
	ctx := context.Background()
	msg := []byte{}

	c := gomock.NewController(t)
	defer c.Finish()

	// 4 OMIT
	mockSender := NewMockSender(c)
	mockSender.EXPECT().Send(ctx, "1", msg).Return(nil).Times(1)
	mockSender.EXPECT().Send(ctx, "2", msg).Return(wantErr).Times(1)

	err := NewMultiSender(mockSender).Send(ctx, "1", msg)
	assert.NoError(t, err)

	err = NewMultiSender(mockSender).Send(ctx, "2", msg)
	assert.EqualError(t, wantErr, err.Error())
	// END 4 OMIT

}
