package main

import "testing"

func TestHello(t *testing.T) {
  rc := Foo()
  if rc != 1 {
    t.Error("Must return 1")
  }
}

func ExampleHello() {
  Foo()
  // Output:
  // hello
}
