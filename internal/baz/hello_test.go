package baz_test

import "strings"
import "testing"
import "github.com/candlerb/sample-exe-3/internal/baz"

func TestGreeting(t *testing.T) {
	g := baz.Greeting()
	if !strings.Contains(g, "Hello") {
		t.Errorf("Greeting is not polite: %#v", g)
	}
}
