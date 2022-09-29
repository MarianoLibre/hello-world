package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, Mariano!"
    if got := Hello("Mariano"); got != want {
        t.Errorf("Hello() = %q, want %q\n", got, want)
    }
}
