package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	singleton1 := New()
	singleton2 := New()

	if singleton1 != singleton2 {
		t.Errorf("singleton1 = %v, singleton2 = %v", singleton1, singleton2)
	}
}
