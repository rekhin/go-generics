package buffer

import "context"

type Receiver[M any] []M

func NewReceiver[M any](messages []M) Receiver[M] {
	return messages
}

func (r Receiver[M]) Receive(ctx context.Context) (M, error) {
	var message M
	// select {
	// case <-ctx.Done():
	// 	return message, ctx.Err()
	// case message = <-r:
	// 	return message, nil
	// }
	return message, nil
}
