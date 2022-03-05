package main

import (
	"context"
	"fmt"
	"time"

	go_generics "github.com/rekhin/go-generics"
)

func main() {
	ctx := context.Background()

	ch := go_generics.MakeChannel(make(chan int))

	go func() {
		for {
			i, ok := ch.Receive(ctx)
			if ok {
				fmt.Println("receive: ", i)
			}
		}
	}()

	for i := 0; i < 1000; i++ {
		ch.Send(ctx, i)
	}

	time.Sleep(1000 * time.Millisecond)
}
