package equations

import (
	"errors"
	"math"
	"sort"
)

func Min(list []int, quantifier int) ([]int, error) {
	if len(list) == 0 {
		return nil, errors.New("Cannot detect a min value in an empty slice")
	}

	sort.Ints(list)

	return list[:quantifier], nil
}

func Max(list []int, quantifier int) ([]int, error) {
	if len(list) == 0 {
		return nil, errors.New("Cannot detect a max value in an empty slice")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(list)))

	return list[:quantifier], nil
}

func Average(list []int) (float64, error) {
	if len(list) == 0 {
		return 0, errors.New("Cannot detect an averge value in an empty slice")
	}

	sum := 0
	for _, x := range list {
		sum += x
	}

	avg := (float64(sum)) / (float64(len(list)))
	truncated := float64(int(avg*100)) / 100
	return truncated, nil
}

func Median(list []int) (int, error) {
	if len(list) == 0 {
		return 0, errors.New("Cannot detect a median value in an empty slice")
	}

	sort.Ints(list)

	middle := len(list) / 2
	medianValue := 0
	if len(list)%2 == 1 {
		medianValue = list[middle]
	} else {
		medianValue = (list[middle-1] + list[middle]) / 2
	}

	return medianValue, nil
}

func Percentile(list []int, quantifier int) (int, error) {
	if len(list) == 0 {
		return 0, errors.New("Cannot detect a percentile value in an empty slice")
	}

	sort.Ints(list)
	ordinal := (float64(quantifier) / 100) * 5
	round := int(math.Ceil(ordinal))

	return list[round-1], nil
}
