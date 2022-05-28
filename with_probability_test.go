package random

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

type testStruct struct {
	Field1      string
	Field2      int
	Probability float64
}

func (t testStruct) GetProbability() float64 {
	return t.Probability
}

func TestGetRandomStructWithProbabilities(t *testing.T) {
	type args struct {
		items []interface{ GetProbability() float64 }
	}

	a := testStruct{
		Field1:      "a",
		Field2:      1,
		Probability: 0.0,
	}
	b := testStruct{
		Field1:      "b",
		Field2:      2,
		Probability: 0.2,
	}
	c := testStruct{
		Field1:      "c",
		Field2:      3,
		Probability: 0.0,
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "test1",
			args: args{
				items: []interface{ GetProbability() float64 }{a, b, c},
			},
			want: b,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomStructWithProbabilities(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomStructWithProbabilities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandomMapItemWithProbabilities(t *testing.T) {
	type args struct {
		items map[string]float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				items: map[string]float64{
					"a": 0.0,
					"b": 0.1,
					"c": 0.0,
				},
			},
			want: "b",
		},
		{
			name: "test2",
			args: args{
				items: map[string]float64{
					"a": 55.1,
					"b": 12.1,
					"c": 0.2,
				},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomMapItemWithProbabilities(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomMapItemWithProbabilities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandomMapItemWithPrecent(t *testing.T) {
	type args struct {
		items map[string]float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				items: map[string]float64{
					"a": 0.0,
					"b": 0.5,
					"c": 0.0,
				},
			},
			want: "b",
		},
		{
			name: "test2",
			args: args{
				items: map[string]float64{
					"a": 56.7,
					"b": 12.3,
					"c": 0.2,
				},
			},
			want: "a",
		},
		{
			name: "test3",
			args: args{
				items: map[string]float64{
					"a": 0.0,
					"b": 0.0,
					"c": 0.0,
					"d": 0.1,
				},
			},
			want: "d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomMapItemWithPrecent(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomMapItemWithPrecent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomMapItem(t *testing.T) {
	var (
		a, b, c, d, e int
	)
	for i := 0; i < 1000; i++ {
		switch GetRandomMapItemWithPrecent(map[string]float64{
			"a": 50.0,
			"b": 30.0,
			"c": 12.0,
			"d": 7.5,
			"e": 0.1,
		}) {
		case "a":
			a++
		case "b":
			b++
		case "c":
			c++
		case "d":
			d++
		case "e":
			e++
		}
	}

	assert.Less(t, e, d)
	assert.Less(t, d, c)
	assert.Less(t, c, b)
	assert.Less(t, b, a)

	t.Logf("a: %d, b: %d, c: %d, d: %d, e: %d", a, b, c, d, e)
}
