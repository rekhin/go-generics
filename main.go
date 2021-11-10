package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// go run -gcflags=-G=3 main.go map.go safe.go

func main() {
	var m safeMap[string, int] = newMutexMap(map[string]int{"": 0})
	m.Set("hello", 11)
	m.Set("world", 22)
	m.SetMany(map[string]int{
		"33": 33,
		"44": 44,
	})
	m.Range(func(key string, value int) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
	fmt.Println(m)
	fmt.Println("lenght:", m.Lenght())

	// channel
	ch := newChannel(make(chan int))
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	send := func(message int) {
		wg.Add(1)
		go func() {
			ch.Send(ctx, message)
			time.Sleep(10 * time.Millisecond) // даём время прочитать из канала
			wg.Done()
		}()
	}

	var got []int
	go func() {
		for {
			message, ok := ch.Receive(ctx)
			if !ok {
				return
			}
			got = append(got, message)
		}
	}()

	send(1)
	wg.Wait()
	// assert.Equal(t, []int{1}, got)
	fmt.Println("got:", got)

	send(2)
	wg.Wait()
	// assert.Equal(t, []int{1, 2}, got)
	fmt.Println("got:", got)

	cancel()
	send(3) // отправка при отсутсвии читающего не проходит
	wg.Wait()

	// assert.Equal(t, []int{1, 2}, got)
	fmt.Println("got:", got)
}
