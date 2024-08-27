package equations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// If I did this all over I would do all the tests in this format.
// I feel its easier to read when you have lots of tests.
// I left the rest as is for comparison.
func TestMin(t *testing.T) {
	t.Run("0 quantifier", func(t *testing.T) {
		e := NewEquationWithQuantifier([]int{2, 1, 4}, 0)
		_, err := e.Min()
		require.ErrorAs(t, err, new(ErrBadRequest))
	})

	t.Run("quantifier larger than list", func(t *testing.T) {
		e := NewEquationWithQuantifier([]int{2, 1, 4}, 5)
		_, err := e.Min()
		require.ErrorAs(t, err, new(ErrBadRequest))
	})

	t.Run("1 quantifier", func(t *testing.T) {
		e := NewEquationWithQuantifier([]int{2, 1, 4}, 1)
		result, err := e.Min()
		require.NoError(t, err)
		require.Equal(t, []int{1}, result)
	})

	t.Run("2 quantifier", func(t *testing.T) {
		e := NewEquationWithQuantifier([]int{2, 1, 4}, 2)
		result, err := e.Min()
		require.NoError(t, err)
		require.Equal(t, []int{1, 2}, result)
	})
}

func TestMax(t *testing.T) {
	cases := []struct {
		e        Equation
		expected []int
	}{
		{NewEquationWithQuantifier([]int{2, 1, 4}, 2), []int{4, 2}},
		{NewEquationWithQuantifier([]int{2, 1, 4}, 1), []int{4}},
		{NewEquationWithQuantifier([]int{2, 1, 4, 6, 10, 3, 20}, 5), []int{20, 10, 6, 4, 3}},
	}

	e := NewEquationWithQuantifier([]int{}, 0)
	_, err := e.Max()
	require.Error(t, err)

	for _, c := range cases {
		result, err := c.e.Max()
		require.NoError(t, err)
		require.Equal(t, c.expected, result)
	}
}

func TestAverage(t *testing.T) {
	cases := []struct {
		e        Equation
		expected float64
	}{
		{NewEquation([]int{2, 1, 4, 4}), 2.75},
		{NewEquation([]int{2, 1, 4}), 2.33},
		{NewEquation([]int{2, 1, 4, 6, 10, 3, 20}), 6.57},
	}

	for _, c := range cases {
		result, err := c.e.Average()
		require.NoError(t, err)
		require.Equal(t, c.expected, result)
	}
}

func TestMedian(t *testing.T) {
	cases := []struct {
		e        Equation
		expected int
	}{
		{NewEquation([]int{2, 1, 4, 4}), 3},
		{NewEquation([]int{1, 2, 4, 5, 3}), 3},
		{NewEquation([]int{2, 1, 4, 6, 10, 3, 20}), 4},
	}

	for _, c := range cases {
		result, err := c.e.Median()
		require.NoError(t, err)
		require.Equal(t, c.expected, result)
	}
}

func TestPercentile(t *testing.T) {
	cases := []struct {
		e        Equation
		expected int
	}{
		{NewEquationWithQuantifier([]int{15, 20, 35, 40, 50}, 30), 20},
		{NewEquationWithQuantifier([]int{15, 20, 35, 40, 50}, 5), 15},
		{NewEquationWithQuantifier([]int{15, 20, 35, 40, 50}, 50), 35},
		{NewEquationWithQuantifier([]int{15, 20, 35, 40, 50}, 100), 50},
	}

	for _, c := range cases {
		result, err := c.e.Percentile()
		require.NoError(t, err)
		require.Equal(t, c.expected, result)
	}
}
