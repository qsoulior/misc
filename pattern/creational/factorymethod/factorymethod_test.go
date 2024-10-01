package factorymethod

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	var factory1 Factory = ConcreteFactory1{}
	product := factory1.CreateProduct()
	product.Action()

	var factory2 Factory = ConcreteFactory2{}
	product = factory2.CreateProduct()
	product.Action()
}
