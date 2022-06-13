package main

import "testing"

func Test_Hello_TakesInput_ReturnsMessageWithInput(t *testing.T) {
	got := Hello("Tunde")
	expected := "Hello,Tunde"

	if got != expected {
		t.Errorf("got = %q :: and expected = %q", got, expected)
	}
}
