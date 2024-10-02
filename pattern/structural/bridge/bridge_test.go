package bridge

import "testing"

func TestBridge(t *testing.T) {
	var abstraction Abstraction = new(ConcreteAbstraction1)
	var implementation Implementation = ConcreteImplementation1{}
	abstraction.SetImplementation(implementation)
	abstraction.Action()

	implementation = ConcreteImplementation2{}
	abstraction.SetImplementation(implementation)
	abstraction.Action()

	abstraction = new(ConcreteAbstraction2)
	abstraction.SetImplementation(implementation)
	abstraction.Action()
}
