package composite

import "testing"

func TestComposite(t *testing.T) {
	var component Component = Leaf{}
	component.Execute()

	component = Composite{children: []Component{
		Composite{children: []Component{Leaf{}}},
		Leaf{},
		Leaf{},
	}}
	component.Execute()
}
