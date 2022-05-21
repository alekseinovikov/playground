package internal

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		params []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1+1=2",
			args{[]int{1, 1}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.params...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcatStringAndNumber(t *testing.T) {
	type args struct {
		str string
		num float32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Hello:,42.12 = Hello:42.12",
			args{"Hello:", 42.12},
			"Hello:42.12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatStringAndNumber(tt.args.str, tt.args.num); got != tt.want {
				t.Errorf("ConcatStringAndNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
