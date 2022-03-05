package main

import (
	"context"
	"fmt"
	"time"

	go_generics "github.com/rekhin/go-generics"
)

func main() {
	ctx := context.Background()

	c := go_generics.NewChannel(make(chan int))

	go func() {
		for {
			i, ok := c.Receive(ctx)
			if ok {
				fmt.Println("receive: ", i)
			}
		}
	}()

	for i := 0; i < 1000; i++ {
		c.Send(ctx, i)
	}

	time.Sleep(1000 * time.Millisecond)
}
