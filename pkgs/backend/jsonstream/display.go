package jsonstream

import (
	"context"
	"io"

	"github.com/moby/moby/api/types/jsonstream"
	"github.com/moby/moby/client/pkg/jsonmessage"
)

type (
	JSONError    = jsonstream.Error
	JSONMessage  = jsonstream.Message
	JSONProgress = jsonstream.Progress
)

type ctxReader struct {
	err chan error
	r   io.Reader
}

func (r *ctxReader) Read(p []byte) (n int, err error) {
	select {
	case err = <-r.err:
		return 0, err
	default:
		return r.r.Read(p)
	}
}

type Options func(*options)

type options struct {
	AuxCallback func(JSONMessage)
}

func WithAuxCallback(cb func(JSONMessage)) Options {
	return func(o *options) {
		o.AuxCallback = cb
	}
}

func Display(ctx context.Context, in io.Reader, out io.Writer, opts ...Options) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	reader := &ctxReader{err: make(chan error, 1), r: in}
	stopFunc := context.AfterFunc(ctx, func() { reader.err <- ctx.Err() })
	defer stopFunc()

	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	if err := jsonmessage.DisplayJSONMessagesStream(reader, out, 0, false, o.AuxCallback); err != nil {
		return err
	}

	return ctx.Err()
}
