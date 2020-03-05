package testing

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  []string
	}{
		{input: "abc", sep: "/", want: []string{"abc"}},
		{input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Split(tt.input, tt.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
