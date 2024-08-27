package equations

import (
	"math"
	"sort"
)

type ErrBadRequest string

func (e ErrBadRequest) Error() string {
	return string(e)
}

// Not a fan of this naming but I am leaving it from when I did it originally because naming is hard.
type Equation interface {
	Min() ([]int, error)
	Max() ([]int, error)
	Average() (float64, error)
	Median() (int, error)
	Percentile() (int, error)
}

type equation struct {
	List       []int
	Quantifier int
}

func NewEquationWithQuantifier(list []int, quantifier int) Equation {
	return &equation{
		List:       list,
		Quantifier: quantifier,
	}
}

func NewEquation(list []int) Equation {
	return &equation{
		List: list,
	}
}

// Min returns a list the size of quantifier containing the smallest numbers in the list
// If I was to redo this whole thing it would be nice to have all these functions be generic
// and the endpoints can then use differnent types
func (e *equation) Min() ([]int, error) {
	// I go back and forth on putting these checks in their own function.
	// Min and max are the same check but percentile is slightly different.
	// I could do it in its own function even with the small difference and probably
	// would in a real system that is more complex.
	if e.Quantifier > len(e.List) {
		return nil, ErrBadRequest("quantifier must be less than or equal to the length of the list")
	}
	if e.Quantifier == 0 {
		return nil, ErrBadRequest("quantifier must be greater than 0")
	}

	sort.Ints(e.List)
	return e.List[:e.Quantifier], nil
}

// Max returns a list the size of quantifier containing the largest numbers in the list
func (e *equation) Max() ([]int, error) {
	if e.Quantifier > len(e.List) {
		return nil, ErrBadRequest("quantifier must be less than or equal to the length of the list")
	}
	if e.Quantifier == 0 {
		return nil, ErrBadRequest("quantifier must be greater than 0")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(e.List)))
	return e.List[:e.Quantifier], nil
}

// Average returns the average of the numbers in the list
func (e *equation) Average() (float64, error) {
	sum := 0
	for _, x := range e.List {
		sum += x
	}

	avg := (float64(sum)) / (float64(len(e.List)))
	truncated := float64(int(avg*100)) / 100
	return truncated, nil
}

// Median returns the median of the numbers in the list
func (e *equation) Median() (int, error) {
	sort.Ints(e.List)

	middle := len(e.List) / 2
	medianValue := 0
	if len(e.List)%2 == 1 {
		medianValue = e.List[middle]
	} else {
		medianValue = (e.List[middle-1] + e.List[middle]) / 2
	}

	return medianValue, nil
}

// Percentile returns the percentile of the list given the provided quantifier
func (e *equation) Percentile() (int, error) {
	if e.Quantifier > 100 {
		return 0, ErrBadRequest("quantifier must be less than or equal to 100")
	}
	if e.Quantifier == 0 {
		return 0, ErrBadRequest("quantifier must be greater than 0")
	}

	sort.Ints(e.List)

	// if the quantifier is 100 then return the last value
	if e.Quantifier == 100 {
		return e.List[len(e.List)-1], nil
	}

	ordinal := (float64(e.Quantifier) / 100) * 5
	round := int(math.Ceil(ordinal))

	return e.List[round-1], nil
}
