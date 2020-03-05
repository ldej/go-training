package email

import (
	"fmt"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "", want: false},
		{input: "@", want: false},
		{input: "@xebia", want: false},
		{input: "l/dejong@xebia..com", want: false},
		{input: "ldejong@xebia..com", want: false},
		{input: "ldejong@xebia", want: true},
		{input: "ldejong@xebia.com", want: true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Testcase: %s", tc.input), func(t *testing.T) {
			got := IsValidEmailAddress(tc.input)
			if got != tc.want {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
