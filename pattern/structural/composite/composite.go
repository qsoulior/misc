// Package composite implements composite pattern.
package composite

// Component represents abstract element of tree.
type Component interface{ Execute() }

// Leaf implements [Component] that does not have sub-components.
type Leaf struct{}

func (l Leaf) Execute() {}

// Composite implements [Component] that has sub-components.
type Composite struct {
	children []Component
}

// Execute does some work and runs execution for each sub-component.
func (c Composite) Execute() {
	for _, child := range c.children {
		child.Execute()
	}
}
