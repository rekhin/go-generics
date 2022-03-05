package buffer

import (
	"context"
)

type Sender[M any] []M

func NewSender[M any](messages []M) Sender[M] {
	return messages
}

func (s Sender[M]) Send(_ context.Context, message M) error {
	s = append(s, message)
	return nil
}
