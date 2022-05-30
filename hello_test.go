package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Cuong")
	want := "Hello, Cuong"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}