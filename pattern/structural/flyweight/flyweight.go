// Package flyweight implements flyweight pattern.
package flyweight

// InternalState implements repeating state.
type InternalState struct{}

// ExternalState implements unique state.
type ExternalState struct{}

// Flyweight implements some structure with internal state.
type Flyweight struct{ state InternalState }

func (f Flyweight) Action(state ExternalState) {}

// Context implements some structure that is paired with [Flyweight] and has external state.
type Context struct {
	flyweight *Flyweight
	state     ExternalState
}

func (c Context) Action() { c.flyweight.Action(c.state) }

// FlyweightFactory implements pool of existing flyweights.
type FlyweightFactory map[InternalState]*Flyweight

func (f FlyweightFactory) Get(state InternalState) *Flyweight {
	if f[state] == nil {
		f[state] = &Flyweight{state}
	}
	return f[state]
}
