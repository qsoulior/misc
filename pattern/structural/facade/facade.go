// Package facade implements facade pattern.
package facade

// Service1 implements some service.
type Service1 struct{}

func (s Service1) Action() {}

// Service2 implements some service.
type Service2 struct{}

func (s Service2) Action() {}

// Service3 implements some service.
type Service3 struct{}

func (s Service3) Action() {}

// Facade provides access to functionality of services.
type Facade struct {
	service1 *Service1
	service2 *Service2
	service3 *Service3
}

// ComplexAction does some work with service actions.
func (s Facade) ComplexAction() {
	s.service1.Action()
	s.service2.Action()
	s.service3.Action()
}
