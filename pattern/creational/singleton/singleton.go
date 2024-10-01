// Package singleton implements singleton pattern.
package singleton

import "sync"

var (
	once     sync.Once
	instance *singleton
)

type singleton struct{}

func New() *singleton {
	if instance == nil {
		once.Do(func() { instance = &singleton{} })
	}
	return instance
}
