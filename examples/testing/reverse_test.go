package testing

import (
	"testing"
)

func TestReverse(t *testing.T) {
	value := Reverse("ecnerual")
	if value != "laurence" {
		t.Errorf("Reverse() = %v, want %v", value, "laurence")
	}
}