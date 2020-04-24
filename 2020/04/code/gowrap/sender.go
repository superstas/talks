package gowrap

import (
	"context"

	"go.opencensus.io/trace"
)

// 0 OMIT
// Sender represents any message sender.
type Sender interface {
	Send(ctx context.Context, to string, msg []byte) error
}

// END 0 OMIT

// 1 OMIT
// SenderWithTracing implements Sender interface instrumented with OpenCencus
type SenderWithTracing struct {
	name string
	Sender
}

// NewSenderWithTracing returns SenderWithTracing
func NewSenderWithTracing(s Sender, name string) SenderWithTracing {
	return SenderWithTracing{
		name:   name,
		Sender: s,
	}
}

// Send implements Sender
func (s SenderWithTracing) Send(ctx context.Context, to string, payload []byte) (err error) {
	ctx, span := trace.StartSpan(ctx, s.name+".Sender.Send", // HL
		trace.WithSpanKind(trace.SpanKindClient)) // HL
	defer span.End()

	return s.Sender.Send(ctx, to, payload)
}

// END 1 OMIT
