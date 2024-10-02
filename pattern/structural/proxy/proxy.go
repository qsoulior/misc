// Package proxy implements proxy pattern.
package proxy

// Service represents abstract service.
type Service interface{ Action() }

// ConcreteService implements [Service] with specified method.
type ConcreteService struct{}

func (s ConcreteService) Action() {}

// Proxy implements [Service] and wraps [ConcreteService] to do additional work.
type Proxy struct {
	service *ConcreteService
}

func (p Proxy) Action() {
	// additional work
	p.service.Action()
	// additional work
}
