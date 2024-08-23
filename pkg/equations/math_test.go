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
		_, err := Min([]int{2, 1, 4}, 0)
		require.ErrorAs(t, err, new(ErrBadRequest))
	})

	t.Run("quantifier larger than list", func(t *testing.T) {
		_, err := Min([]int{2, 1, 4}, 4)
		require.ErrorAs(t, err, new(ErrBadRequest))
	})

	t.Run("1 quantifier", func(t *testing.T) {
		result, err := Min([]int{2, 1, 4}, 1)
		require.NoError(t, err)
		require.Equal(t, []int{1}, result)
	})

	t.Run("2 quantifier", func(t *testing.T) {
		result, err := Min([]int{2, 1, 4}, 2)
		require.NoError(t, err)
		require.Equal(t, []int{1, 2}, result)
	})
}

func TestMax(t *testing.T) {
	cases := []struct {
		list       []int
		quantifier int
		expected   []int
	}{
		{[]int{2, 1, 4}, 2, []int{4, 2}},
		{[]int{2, 1, 4}, 1, []int{4}},
		{[]int{2, 1, 4, 6, 10, 3, 20}, 5, []int{20, 10, 6, 4, 3}},
	}

	_, err := Max([]int{}, 3)
	require.Error(t, err)

	for _, c := range cases {
		result, err := Max(c.list, c.quantifier)
		require.NoError(t, err)
		require.Equal(t, c.expected, result)
	}
}

func TestAverage(t *testing.T) {
	cases := []struct {
		list     []int
		expected float64
	}{
		{[]int{2, 1, 4, 4}, 2.75},
		{[]int{2, 1, 4}, 2.33},
		{[]int{2, 1, 4, 6, 10, 3, 20}, 6.57},
	}

	_, err := Average([]int{})
	require.Error(t, err)

	for _, c := range cases {
		result, err := Average(c.list)
		require.NoError(t, err)
		require.Equal(t, c.expected, result)
	}
}

func TestMedian(t *testing.T) {
	cases := []struct {
		list     []int
		expected int
	}{
		{[]int{2, 1, 4, 4}, 3},
		{[]int{1, 2, 4, 5, 3}, 3},
		{[]int{2, 1, 4, 6, 10, 3, 20}, 4},
	}

	_, err := Median([]int{})
	require.Error(t, err)

	for _, c := range cases {
		result, err := Median(c.list)
		require.NoError(t, err)
		require.Equal(t, c.expected, result)
	}
}

func TestPercentile(t *testing.T) {
	cases := []struct {
		list       []int
		quantifier int
		expected   int
	}{
		{[]int{15, 20, 35, 40, 50}, 30, 20},
		{[]int{15, 20, 35, 40, 50}, 5, 15},
		{[]int{15, 20, 35, 40, 50}, 50, 35},
		{[]int{15, 20, 35, 40, 50}, 100, 50},
	}

	_, err := Percentile([]int{}, 20)
	require.Error(t, err)

	for _, c := range cases {
		result, err := Percentile(c.list, c.quantifier)
		require.NoError(t, err)
		require.Equal(t, c.expected, result)
	}
}
