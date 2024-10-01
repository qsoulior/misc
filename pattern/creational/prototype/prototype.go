// Package prototype implements prototype pattern.
package prototype

// Prototype represents abstract entity that can be cloned.
type Prototype interface {
	Clone() Prototype
}

// ConcretePrototype1 implements [Prototype] and can be cloned.
type ConcretePrototype1 struct{ x, y int }

func (p ConcretePrototype1) Clone() Prototype { return ConcretePrototype1{p.x, p.y} }

// ConcretePrototype2 implements [Prototype] and can be cloned.
type ConcretePrototype2 struct{ a, b int }

func (p ConcretePrototype2) Clone() Prototype { return ConcretePrototype2{p.a, p.b} }
