package random

import (
	"reflect"
	"testing"
)

func TestGetRandomWithProbabilities(t *testing.T) {
	type args struct {
		items         []interface{}
		probabilities []float64
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "test1",
			args: args{
				items: []interface{}{
					"a",
					"b",
					"c",
				},
				probabilities: []float64{
					0.0,
					0.2,
					0.0,
				},
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomWithProbabilities(tt.args.items, tt.args.probabilities); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomWithProbabilities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomFloat64(t *testing.T) {
	type args struct {
		max float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				max: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomFloat64(tt.args.max); got > tt.args.max {
				t.Errorf("randomFloat64() = %v, want <= %v", got, tt.args.max)
			}
		})
	}
}

func Test_sumSlice(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				values: []float64{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				values: []float64{0.1, 2, 3.5, 4.1},
			},
			want: 9.7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSlice(tt.args.values); got != tt.want {
				t.Errorf("sumSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
