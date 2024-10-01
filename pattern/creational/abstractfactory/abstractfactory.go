// Package abstractfactory implements abstract factory pattern.
package abstractfactory

// ProductX represents object X created by abstract factory.
type ProductX interface{ Action() }

// SpecificProductX1 implements [ProductX] with specific method.
// It belongs to the object family #1.
type SpecificProductX1 struct{}

func (p SpecificProductX1) Action() {}

// SpecificProductX2 implements [ProductX] with specific method.
// It belongs to the object family #2.
type SpecificProductX2 struct{}

func (p SpecificProductX2) Action() {}

// ProductY represents object Y created by abstract factory.
type ProductY interface{ Action() }

// SpecificProductY1 implements [ProductY] with specific method.
// It belongs to the object family #1.
type SpecificProductY1 struct{}

func (p SpecificProductY1) Action() {}

// SpecificProductY2 implements [ProductY] with specific method.
// It belongs to the object family #2.
type SpecificProductY2 struct{}

func (p SpecificProductY2) Action() {}

// Factory represents abstract factory that creates objects X and Y.
type Factory interface {
	CreateProductX() ProductX
	CreateProductY() ProductY
}

// SpecificFactory1 implements [Factory] and creates objects of family #1.
type SpecificFactory1 struct{}

// CreateProductX creates and returns [ProductX] of family #1.
func (f SpecificFactory1) CreateProductX() ProductX { return SpecificProductX1{} }

// CreateProductY creates and returns [ProductY] of family #1.
func (f SpecificFactory1) CreateProductY() ProductY { return SpecificProductY1{} }

// SpecificFactory2 implements [Factory] and creates objects of family #2.
type SpecificFactory2 struct{}

// CreateProductX creates and returns [ProductX] of family #2.
func (f SpecificFactory2) CreateProductX() ProductX { return SpecificProductX2{} }

// CreateProductY creates and returns [ProductY] of family #2.
func (f SpecificFactory2) CreateProductY() ProductY { return SpecificProductY2{} }
