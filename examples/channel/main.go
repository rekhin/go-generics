package main

import (
	"context"
	"log"
	"time"

	"github.com/rekhin/go-generics/buffer"
	"github.com/rekhin/go-generics/inproc"
)

func main() {
	ctx := context.Background()

	c := make(chan int)
	r := inproc.NewReceiver(c)

	go func() {
		for {
			i, err := r.Receive(ctx)
			if err != nil {
				log.Printf("receive failed: %s", err)
				continue
			}
			log.Println(i)
		}
	}()

	s := inproc.NewSender(c)
	for i := 0; i < 1000; i++ {
		s.Send(ctx, i)
	}

	b := make([]int, 0)
	bs := buffer.NewSender(b)

	time.Sleep(1000 * time.Millisecond)
}
