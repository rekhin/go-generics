package main

import (
	"context"
	"time"
)

type channel[T any] struct {
	ch chan T
}

func newChannel[T any](ch chan T) *channel[T] {
	return &channel[T]{
		ch: ch,
	}
}

func (c *channel[T]) Send(ctx context.Context, message T) bool {
	select {
	case <-ctx.Done():
		return false
	case c.ch <- message:
	}
	return true
}

func (c *channel[T]) SendWithTimeout(ctx context.Context, message T, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	ok := c.Send(ctx, message)
	return ok
}

func (c *channel[T]) Receive(ctx context.Context) (T, bool) {
	var message T
	select {
	case <-ctx.Done():
		return message, false

	case message = <-c.ch:
	}
	return message, true
}

