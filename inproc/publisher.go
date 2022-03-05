package inproc

type Publisher[M any] Sender[M]

func NewPublisher[M any](ch Sender[M]) Publisher[M] {
	return Publisher[M](ch)
}
