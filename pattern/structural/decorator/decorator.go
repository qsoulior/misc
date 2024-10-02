// Package decorator implements decorator pattern.
package decorator

// Component represents some target interface.
type Component interface{ Execute() }

// Decoratee implements [Component] with specified method.
type Decoratee struct{}

func (d Decoratee) Execute() {}

// Decorator1 implements [Component] and wraps [Decoratee] to do additional work.
type Decorator1 struct{ decoratee Component }

func (d Decorator1) Execute() {
	// additional work
	d.decoratee.Execute()
}

// Decorator2 implements [Component] and wraps [Decoratee] to do additional work.
type Decorator2 struct{ decoratee Component }

func (d Decorator2) Execute() {
	d.decoratee.Execute()
	// additional work
}
