// Package adapter implements adapter pattern.
package adapter

// Client represents some interface to work with client.
type Client interface{ Action() }

// Adaptee is some structure that does not implement [Client].
type Adaptee struct{}

func (a Adaptee) Do() {}

// Adapter wraps [Adaptee] to implement [Client].
type Adapter struct {
	adaptee Adaptee
}

func (a Adapter) Action() { a.adaptee.Do() }
