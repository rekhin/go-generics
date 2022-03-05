package go_generics

import "context"

type Channel[M any] chan M

func NewChannel[M any](c chan M) Channel[M] {
	return c
}

func (c Channel[M]) Send(ctx context.Context, message M) bool {
	select {
	case <-ctx.Done():
		return false
	case c <- message:
		return true
	}

}

func (c Channel[M]) Receive(ctx context.Context) (message M, ok bool) {
	select {
	case <-ctx.Done():
		return message, false
	case message = <-c:
		return message, true
	}
}

type Publisher[M any] Channel[M]

func NewPublisher[M any](ch Channel[M]) Publisher[M] {
	return Publisher[M](ch)
}

type Subscriber[M any] Channel[M]

func NewSubscriber[M any](ch Channel[M]) Subscriber[M] {
	return Subscriber[M](ch)
}

func (s Subscriber[M]) Subscribe(ctx context.Context) (message M, ok bool) {
	return Channel[M](s).Receive(ctx)
}
