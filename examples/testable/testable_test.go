package testable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result, err := Add(1, 2)
	assert.Nil(t, err)

	if result != 4 {
		t.Error("result not 4")
	}
}

func TestDivide2(t *testing.T) {
	_, err := Divide(1, 0)
	if err != ErrDivideByZero {
		t.Error("expected an error")
	}
	result, err := Divide(10, 10)
	if err != nil {
		t.Error("did not expect an error")
	}
	if result != 1 {
		t.Error("expected 1")
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "success", args: args{a: 10, b: 10}, want: 1, wantErr: false},
		{name: "by 0", args: args{a: 1, b: 0}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Divide() got = %v, want %v", got, tt.want)
			}
		})
	}
}
