package decorator

import "testing"

func TestDecorator(t *testing.T) {
	var component Component = Decoratee{}
	component = Decorator1{component}
	component = Decorator2{component}
	component.Execute()
}
