package main

type safeMap[K comparable, V any] interface{
	Get(key K) (value V, ok bool)
	Set(key K, value V)
	SetMany(from map[K]V)
	Delete(key K)
	Lenght() int
	Copy() map[K]V
	Range(f func(K, V) bool)
	String() string
}
