package singleton

import (
	"sync"
)

// 懒汉模式(Mutex)
var mu sync.Mutex

type lazy struct {
	Name string
}

var l *lazy = nil

func GetLazy() *lazy {
	if l == nil {
		mu.Lock()
		if l == nil {
			l = new(lazy)
			l.Name = "Lazy man"
		}
		mu.Unlock()
	}
	return l
}

//懒汉模式(Once)
var once sync.Once

type lazyonce struct {
	Name string
}

var lo *lazyonce = nil

func GetOnce() *lazyonce {
	once.Do(func() {
		lo = &lazyonce{Name: "lazyOnce"}
	})
	return lo
}

//饿汉模式(init)
var h *hungry = nil

type hungry struct {
	Name string
}

func init() {
	h = new(hungry)
	h.Name = "Hungry man"
}

func GetHungry() *hungry {
	return h
}
