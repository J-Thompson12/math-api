package equations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	cases := []struct {
		list       []int
		quantifier int
		expected   []int
	}{
		{[]int{2, 1, 4}, 2, []int{1, 2}},
		{[]int{2, 1, 4}, 1, []int{1}},
		{[]int{2, 1, 4, 6, 10, 3, 20}, 5, []int{1, 2, 3, 4, 6}},
	}

	_, err := Min([]int{}, 3)
	require.Error(t, err)

	for _, c := range cases {
		result, err := Min(c.list, c.quantifier)
		require.NoError(t, err)
		require.Equal(t, c.expected, result)
	}
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
