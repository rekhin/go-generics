package main

import (
	"fmt"

	go_generics "github.com/rekhin/go-generics"
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
}
