package lock4

import (
	"sync"
)

const LOCK4_VERSION = "2024.02.28"

type Lock4 struct {
}

var (
	instance *Lock4
	once     sync.Once
)

func GetInstance() *Lock4 {
	once.Do(func() {
		instance = NewLock3()
	})
	return instance
}

func NewLock3() *Lock4 {
	var l = &Lock4{}

	return l
}
