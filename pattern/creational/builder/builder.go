// Package builder implements builder pattern.
package builder

// Product implements product that is being built.
type Product struct{ x, y, z int }

// Builder represents abstract builder that builds and creates [Product].
type Builder interface {
	Reset() Builder
	SetX(x int) Builder
	SetY(y int) Builder
	SetZ(z int) Builder
	Build() Product
}

// SpecificBuilder implements [Builder] with simple logic.
type SpecificBuilder struct{ x, y, z int }

// Reset sets initial state for [Product] object.
func (b *SpecificBuilder) Reset() Builder {
	b.x = 0
	b.y = 0
	b.z = 0
	return b
}

// SetX sets x for state and returns [SpecificBuilder].
func (b *SpecificBuilder) SetX(x int) Builder {
	b.x = x
	return b
}

// SetX sets y for state and returns [SpecificBuilder].
func (b *SpecificBuilder) SetY(y int) Builder {
	b.y = y
	return b
}

// SetZ sets z for state and returns [SpecificBuilder].
func (b *SpecificBuilder) SetZ(z int) Builder {
	b.z = z
	return b
}

// Build creates and returns built [Product].
func (b SpecificBuilder) Build() Product { return Product{b.x, b.y, b.z} }

// Director defines the order in which to call build steps.
type Director struct{ builder Builder }

// SetBuilder sets active [Builder].
func (d *Director) SetBuilder(builder Builder) { d.builder = builder }

// GetProduct builds and returns [Product].
func (d Director) GetProduct() Product { return d.builder.Reset().SetX(10).SetY(20).Build() }
