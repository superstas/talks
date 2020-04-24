package gomock

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// Sender represents any message sender.
type Sender interface {
	Send(ctx context.Context, to string, msg []byte) error
}

// MultiSender duplicates a message to all provided senders.
type MultiSender struct {
	senders []Sender
}

// NewMultiSender creates MultiSender.
func NewMultiSender(s ...Sender) *MultiSender {
	return &MultiSender{s}
}

// Send sends a message to a recipient.
func (s *MultiSender) Send(ctx context.Context, to string, msg []byte) error {
	var errg errgroup.Group
	for _, sender := range s.senders {
		sender := sender
		errg.Go(func() error {
			return sender.Send(ctx, to, msg) // HL
		})
	}
	return errg.Wait()
}
