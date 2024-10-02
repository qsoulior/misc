package proxy

import "testing"

func TestProxy(t *testing.T) {
	service := new(ConcreteService)
	var _ Service = service
	var proxy Service = Proxy{service}
	proxy.Action()
}
