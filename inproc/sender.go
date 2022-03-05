package inproc

import (
	"context"
	"time"
)

type Sender[M any] chan M

func NewSender[M any](c chan M) Sender[M] {
	return c
}

func (s Sender[M]) Send(ctx context.Context, message M) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s <- message:
		return nil
	}
}

func (c Sender[M]) SendWithTimeout(ctx context.Context, message M, timeout time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case c <- message:
		return nil
	}
}
