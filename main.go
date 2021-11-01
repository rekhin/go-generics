package main

import (
	"fmt"
)

// go run -gcflags=-G=3 main.go map.go

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

	print([]string{"Hello, ", "playground\n"})
}
