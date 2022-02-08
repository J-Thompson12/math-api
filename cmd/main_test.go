package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompare(t *testing.T) {
	cases := []struct {
		v1       string
		v2       string
		expected int
		testCase string
	}{
		{"1.0.1", "2.1.1", -1, "v2 > v1"},
		{"1.2.1", "1.2.0", 1, "v1 > v2"},
		{"1.2.1", "1.2.1", 0, "v1 == v2"},
	}

	for _, c := range cases {
		result := compare(c.v1, c.v2)
		require.Equal(t, c.expected, result, c.testCase)
	}
}

func TestMatchArrSize(t *testing.T) {
	len3Case := []string{"0", "1", "3"}
	len5Case := []string{"0", "1", "1", "3", "4"}

	matchArrSize(&len3Case, &len5Case)
	require.Equal(t, len(len3Case), 5)
	require.Equal(t, len(len5Case), 5)

	len3Case = []string{"0", "1", "3"}
	len5Case = []string{"0", "1", "1", "3", "4"}

	matchArrSize(&len5Case, &len3Case)
	require.Equal(t, len(len3Case), 5)
	require.Equal(t, len(len5Case), 5)
}
