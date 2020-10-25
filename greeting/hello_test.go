package main

import "testing"

func TestHello(t *testing.T) {
	emptyResult := hello("")
	if emptyResult != "Hello Dude !" {
		t.Errorf("Hello failed")
	}
	result := hello("Mike")
	if result != "Hello Mike!" {
		t.Errorf("Hello failed.")
	}
}
