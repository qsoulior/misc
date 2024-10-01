package prototype

import (
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	var prototype Prototype = ConcretePrototype1{x: 1, y: 2}
	clone1 := prototype.Clone()
	if !reflect.DeepEqual(prototype, clone1) {
		t.Errorf("prototype = %v, clone = %v", prototype, clone1)
		return
	}

	prototype = ConcretePrototype2{a: 1, b: 2}
	clone2 := prototype.Clone()
	if !reflect.DeepEqual(prototype, clone2) {
		t.Errorf("prototype = %v, clone = %v", prototype, clone2)
		return
	}
}
