package random_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dmitrymomot/random"
)

func TestGetRandomWithProbabilities(t *testing.T) {
	t.Parallel()

	t.Run("single non-zero probability", func(t *testing.T) {
		t.Parallel()

		items := []string{"a", "b", "c"}
		probabilities := []float64{0.0, 0.2, 0.0}

		got := random.GetRandomWithProbabilities(items, probabilities)
		require.Equal(t, "b", got)
	})

	t.Run("empty slices", func(t *testing.T) {
		t.Parallel()

		got := random.GetRandomWithProbabilities([]string{}, []float64{})
		require.Equal(t, "", got)
	})

	t.Run("mismatched lengths", func(t *testing.T) {
		t.Parallel()

		items := []string{"a", "b"}
		probabilities := []float64{0.5, 0.3, 0.2}

		got := random.GetRandomWithProbabilities(items, probabilities)
		require.Equal(t, "", got)
	})

	t.Run("negative probability", func(t *testing.T) {
		t.Parallel()

		items := []string{"a", "b", "c"}
		probabilities := []float64{0.5, -0.1, 0.4}

		got := random.GetRandomWithProbabilities(items, probabilities)
		require.Equal(t, "", got)
	})

	t.Run("all zero probabilities", func(t *testing.T) {
		t.Parallel()

		items := []string{"a", "b", "c"}
		probabilities := []float64{0.0, 0.0, 0.0}

		got := random.GetRandomWithProbabilities(items, probabilities)
		require.Equal(t, "", got)
	})

	t.Run("distribution check", func(t *testing.T) {
		t.Parallel()

		items := []string{"a", "b", "c"}
		probabilities := []float64{0.7, 0.2, 0.1}

		counts := make(map[string]int)
		for i := 0; i < 1000; i++ {
			result := random.GetRandomWithProbabilities(items, probabilities)
			counts[result]++
		}

		// a should appear most frequently
		require.Greater(t, counts["a"], counts["b"])
		require.Greater(t, counts["b"], counts["c"])
	})
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
	t.Parallel()

	t.Run("single non-zero probability", func(t *testing.T) {
		t.Parallel()

		a := testStruct{Field1: "a", Field2: 1, Probability: 0.0}
		b := testStruct{Field1: "b", Field2: 2, Probability: 0.2}
		c := testStruct{Field1: "c", Field2: 3, Probability: 0.0}

		items := []testStruct{a, b, c}
		got := random.GetRandomStructWithProbabilities(items)
		require.Equal(t, b, got)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()

		got := random.GetRandomStructWithProbabilities([]testStruct{})
		require.Equal(t, testStruct{}, got)
	})

	t.Run("negative probability", func(t *testing.T) {
		t.Parallel()

		a := testStruct{Field1: "a", Field2: 1, Probability: 0.5}
		b := testStruct{Field1: "b", Field2: 2, Probability: -0.1}

		items := []testStruct{a, b}
		got := random.GetRandomStructWithProbabilities(items)
		require.Equal(t, testStruct{}, got)
	})

	t.Run("all zero probabilities", func(t *testing.T) {
		t.Parallel()

		a := testStruct{Field1: "a", Field2: 1, Probability: 0.0}
		b := testStruct{Field1: "b", Field2: 2, Probability: 0.0}

		items := []testStruct{a, b}
		got := random.GetRandomStructWithProbabilities(items)
		require.Equal(t, testStruct{}, got)
	})

	t.Run("distribution check", func(t *testing.T) {
		t.Parallel()

		a := testStruct{Field1: "a", Field2: 1, Probability: 0.7}
		b := testStruct{Field1: "b", Field2: 2, Probability: 0.2}
		c := testStruct{Field1: "c", Field2: 3, Probability: 0.1}

		items := []testStruct{a, b, c}

		counts := make(map[string]int)
		for i := 0; i < 1000; i++ {
			result := random.GetRandomStructWithProbabilities(items)
			counts[result.Field1]++
		}

		require.Greater(t, counts["a"], counts["b"])
		require.Greater(t, counts["b"], counts["c"])
	})
}

func TestGetRandomMapItemWithProbabilities(t *testing.T) {
	t.Parallel()

	t.Run("single non-zero probability", func(t *testing.T) {
		t.Parallel()

		items := map[string]float64{
			"a": 0.0,
			"b": 0.1,
			"c": 0.0,
		}

		got := random.GetRandomMapItemWithProbabilities(items)
		require.Equal(t, "b", got)
	})

	t.Run("empty map", func(t *testing.T) {
		t.Parallel()

		got := random.GetRandomMapItemWithProbabilities(map[string]float64{})
		require.Equal(t, "", got)
	})

	t.Run("negative probability", func(t *testing.T) {
		t.Parallel()

		items := map[string]float64{
			"a": 0.5,
			"b": -0.1,
			"c": 0.4,
		}

		got := random.GetRandomMapItemWithProbabilities(items)
		require.Equal(t, "", got)
	})

	t.Run("all zero probabilities", func(t *testing.T) {
		t.Parallel()

		items := map[string]float64{
			"a": 0.0,
			"b": 0.0,
			"c": 0.0,
		}

		got := random.GetRandomMapItemWithProbabilities(items)
		require.Equal(t, "", got)
	})

	t.Run("distribution check", func(t *testing.T) {
		t.Parallel()

		items := map[string]float64{
			"a": 0.7,
			"b": 0.2,
			"c": 0.1,
		}

		counts := make(map[string]int)
		for i := 0; i < 1000; i++ {
			result := random.GetRandomMapItemWithProbabilities(items)
			counts[result]++
		}

		require.Greater(t, counts["a"], counts["b"])
		require.Greater(t, counts["b"], counts["c"])
	})
}

func TestGetRandomMapItemWithPercent(t *testing.T) {
	t.Parallel()

	t.Run("empty map", func(t *testing.T) {
		t.Parallel()

		got := random.GetRandomMapItemWithPercent(map[string]float64{})
		require.Equal(t, "", got)
	})

	t.Run("single item", func(t *testing.T) {
		t.Parallel()

		items := map[string]float64{
			"a": 100.0,
		}

		got := random.GetRandomMapItemWithPercent(items)
		require.Equal(t, "a", got)
	})

	t.Run("negative probability", func(t *testing.T) {
		t.Parallel()

		items := map[string]float64{
			"a": 50.0,
			"b": -10.0,
			"c": 30.0,
		}

		got := random.GetRandomMapItemWithPercent(items)
		require.Equal(t, "", got)
	})

	t.Run("all zero probabilities", func(t *testing.T) {
		t.Parallel()

		items := map[string]float64{
			"a": 0.0,
			"b": 0.0,
			"c": 0.0,
		}

		got := random.GetRandomMapItemWithPercent(items)
		require.Equal(t, "", got)
	})

	t.Run("distribution check", func(t *testing.T) {
		t.Parallel()

		items := map[string]float64{
			"a": 50.0,
			"b": 30.0,
			"c": 12.0,
			"d": 7.5,
			"e": 0.1,
		}

		counts := make(map[string]int)
		for i := 0; i < 1000; i++ {
			result := random.GetRandomMapItemWithPercent(items)
			counts[result]++
		}

		require.Greater(t, counts["a"], counts["b"])
		require.Greater(t, counts["b"], counts["c"])
		require.Greater(t, counts["c"], counts["d"])
		require.Greater(t, counts["d"], counts["e"])
	})
}

func Test_randomMapItem(t *testing.T) {
	t.Parallel()

	var (
		a, b, c, d, e int
	)
	for i := 0; i < 1000; i++ {
		switch random.GetRandomMapItemWithPercent(map[string]float64{
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
