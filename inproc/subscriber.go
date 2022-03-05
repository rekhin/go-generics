package inproc

import "context"

type Subscriber[M any] Receiver[M]

func NewSubscriber[M any](ch Receiver[M]) Subscriber[M] {
	return Subscriber[M](ch)
}

func (s Subscriber[M]) Subscribe(ctx context.Context) (M, error) {
	return Receiver[M](s).Receive(ctx)
}
