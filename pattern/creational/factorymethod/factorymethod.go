// Package factorymethod implements factory method pattern.
package factorymethod

// Product represents abstract product created by factory method.
type Product interface{ Action() }

// ConcreteProduct1 implements [Product] with specific method.
type ConcreteProduct1 struct{}

func (p ConcreteProduct1) Action() {}

// ConcreteProduct2 implements [Product] with specific method.
type ConcreteProduct2 struct{}

func (p ConcreteProduct2) Action() {}

// Factory represents abstract entity that can create [Product].
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactory1 implements [Factory] and can create [ConcreteProduct1].
type ConcreteFactory1 struct{}

func (f ConcreteFactory1) CreateProduct() Product { return ConcreteProduct1{} }

// ConcreteFactory2 implements [Factory] and can create [ConcreteProduct2].
type ConcreteFactory2 struct{}

func (f ConcreteFactory2) CreateProduct() Product { return ConcreteProduct2{} }
