package gowrap

import (
	"context"

	"go.opencensus.io/trace"
)

// 0 OMIT
// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using tracing_template template

//go:generate gowrap gen -p talk/code/gowrap -i Sender -t tracing_template -o sender_tracing.go

// END 0 OMIT

// 1 OMIT

// SenderWithTracing implements Sender interface instrumented with opencencus
type SenderWithTracing struct {
	_name string
	Sender
}

// NewSenderWithTracing returns SenderWithTracing
func NewSenderWithTracing(name string, base Sender) SenderWithTracing {
	return SenderWithTracing{_name: name, Sender: base}
}

// Send implements Sender
func (_d SenderWithTracing) Send(ctx context.Context, to string, payload []byte) (err error) {

	ctx, span := trace.StartSpan(ctx, _d._name+".Sender.Send", // HL
		trace.WithSpanKind(trace.SpanKindClient)) // HL
	defer span.End()

	return _d.Sender.Send(ctx, to, payload)
}

// END 1 OMIT
