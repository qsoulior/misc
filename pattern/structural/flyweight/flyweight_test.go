package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	factory := make(FlyweightFactory, 1)
	context := Context{
		flyweight: factory.Get(InternalState{}),
		state:     ExternalState{},
	}
	context.Action()
}
