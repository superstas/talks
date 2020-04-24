package minimock

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// 0 OMIT
// Sender represents any message sender.
type Sender interface {
	Send(ctx context.Context, to string, msg []byte) error
}

// END 0 OMIT

// 1 OMIT
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

// END 1 OMIT
