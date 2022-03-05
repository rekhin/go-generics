package go_generics

import "context"

type Channel[M any] chan M

func MakeChannel[M any](ch chan M) Channel[M] {
	return ch
}

func MakeChannelWithBuffer[M any](size int) Channel[M] {
	return make(chan M, size)
}

func (ch Channel[M]) Send(ctx context.Context, message M) bool {
	select {
	case <-ctx.Done():
		return false
	case ch <- message:
		return true
	}

}

func (ch Channel[M]) Receive(ctx context.Context) (message M, ok bool) {
	select {
	case <-ctx.Done():
		return message, false

	case message = <-ch:
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
