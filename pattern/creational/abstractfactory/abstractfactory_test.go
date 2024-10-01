package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	var factory1 Factory = SpecificFactory1{}
	productX := factory1.CreateProductX()
	productY := factory1.CreateProductY()
	productX.Action()
	productY.Action()

	var factory2 Factory = SpecificFactory2{}
	productX = factory2.CreateProductX()
	productY = factory2.CreateProductY()
	productX.Action()
	productY.Action()
}
