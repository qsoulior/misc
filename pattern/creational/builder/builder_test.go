package builder

import "testing"

func TestBuilder(t *testing.T) {
	var builder SpecificBuilder
	var _ Product = builder.SetX(10).SetY(20).SetZ(30).Build()

	var director Director
	director.SetBuilder(&builder)
	var _ Product = director.GetProduct()
}
