// Package bridge implements bridge pattern.
package bridge

// Implementation represents layer to which abstraction layer delegates the work.
type Implementation interface{ Action() }

// ConcreteImplementation1 implements [Implementation] with specified action.
type ConcreteImplementation1 struct{}

func (i ConcreteImplementation1) Action() {}

// ConcreteImplementation2 implements [Implementation] with specified action.
type ConcreteImplementation2 struct{}

func (i ConcreteImplementation2) Action() {}

// Abstraction represents high-level layer to work with some entity.
// It should delegate the work to the implementation layer.
type Abstraction interface {
	SetImplementation(implementation Implementation)
	Action()
}

// ConcreteAbstraction1 implements [Abstraction] with specified implementation.
type ConcreteAbstraction1 struct {
	implementation Implementation // bridge
}

func (a *ConcreteAbstraction1) SetImplementation(implementation Implementation) {
	a.implementation = implementation
}

func (a ConcreteAbstraction1) Action() { a.implementation.Action() }

// ConcreteAbstraction2 implements [Abstraction] with specified implementation.
type ConcreteAbstraction2 struct {
	implementation Implementation // bridge
}

func (a *ConcreteAbstraction2) SetImplementation(implementation Implementation) {
	a.implementation = implementation
}

func (a ConcreteAbstraction2) Action() { a.implementation.Action() }
