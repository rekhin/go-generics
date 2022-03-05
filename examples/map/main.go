package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	go_generics "github.com/rekhin/go-generics"
	"github.com/rekhin/go-generics/inproc"
)

func main() {
	var m go_generics.SafeMap[string, int] = go_generics.NewMutexMap(map[string]int{"": 0})
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

	fmt.Println([]string{"Hello, ", "playground\n"})

	// channel
	c := make(chan int)
	s := inproc.NewSender(c)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	send := func(message int) {
		wg.Add(1)
		go func() {
			if err := s.Send(ctx, message); err != nil {
				log.Printf("send failed: %s", err)
			}
			time.Sleep(10 * time.Millisecond) // даём время прочитать из канала
			wg.Done()
		}()
	}

	r := inproc.NewReceiver(c)
	var got []int
	go func() {
		for {
			i, err := r.Receive(ctx)
			if err != nil {
				log.Printf("receive failed: %s", err)
				break
			}
			got = append(got, i)
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
