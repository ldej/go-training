package email

import (
	"fmt"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		input   string
		isValid bool
	}{
		{input: "", isValid: false},
		{input: "@", isValid: false},
		{input: "@xebia", isValid: false},
		{input: "l/dejong@xebia..com", isValid: false},
		{input: "ldejong@xebia..com", isValid: false},
		{input: "ldejong@xebia", isValid: true},
		{input: "ldejong@xebia.com", isValid: true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Testcase: %s", tc.input), func(t *testing.T) {
			got := IsValidEmailAddress(tc.input)
			if got != tc.isValid {
				t.Fatalf("expected: %v, got: %v", tc.isValid, got)
			}
		})
	}
}
