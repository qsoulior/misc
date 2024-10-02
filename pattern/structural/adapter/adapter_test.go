package adapter

import "testing"

func TestAdapter(t *testing.T) {
	adaptee := Adaptee{}
	adaptee.Do()
	var client Client = Adapter{adaptee}
	client.Action()
}
