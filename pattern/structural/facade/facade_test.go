package facade

import "testing"

func TestFacade(t *testing.T) {
	service1 := new(Service1)
	service2 := new(Service2)
	service3 := new(Service3)

	facade := Facade{service1, service2, service3}
	facade.ComplexAction()
}
