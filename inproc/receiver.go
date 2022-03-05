package inproc

import "context"

type Receiver[M any] chan M

func NewReceiver[M any](c chan M) Receiver[M] {
	return c
}

func (r Receiver[M]) Receive(ctx context.Context) (M, error) {
	var message M
	select {
	case <-ctx.Done():
		return message, ctx.Err()
	case message = <-r:
		return message, nil
	}
}
