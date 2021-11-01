package main

import (
	"fmt"
	"sync"
)

type mutexMap[K comparable, V any] struct {
	mtx sync.RWMutex
	m 	map[K]V
}

var a safeMap[int, int] = new(mutexMap[int, int])

func newMutexMap[K comparable, V any](m map[K]V) *mutexMap[K,V] {
	return &mutexMap[K,V]{m: m}
}

func (m *mutexMap[K,V]) Get(key K) (value V, ok bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	value, ok = m.m[key]
	return value, ok
}

func (m *mutexMap[K,V]) Set(key K, value V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.m[key] = value
}

func (m *mutexMap[K,V]) SetMany(from map[K]V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for key, value := range from {
		m.m[key] = value
	}
}

func (m *mutexMap[K,V]) Delete(key K) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	delete(m.m, key)
}

func (m *mutexMap[K,V]) Lenght() int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return len(m.m)
}

func (m *mutexMap[K,V]) Copy() map[K]V {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	to := make(map[K]V, len(m.m))
	for k, v := range m.m {
		to[k] = v // TODO возможно тут надо копировать значения
	}
	return to
}

func (m *mutexMap[K,V]) Range(f func(K, V) bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	for key, value := range m.m {
		if ok := f(key, value); !ok {
			return
		}
	}
}

func (m *mutexMap[K,V]) String() string {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return fmt.Sprintf("%+v", m.m)
}
