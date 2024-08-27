package equations

type MockEquation struct {
	// normally I would have the mock be able to set the expected result but in this case hardcoding is fine
	Error error
}

func NewMockEquation() *MockEquation {
	return &MockEquation{}
}

func (m *MockEquation) Min() ([]int, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	return []int{1}, nil
}

func (m *MockEquation) Max() ([]int, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	return []int{8}, nil
}

func (m *MockEquation) Average() (float64, error) {
	if m.Error != nil {
		return 0, m.Error
	}

	return 1.5, nil
}

func (m *MockEquation) Median() (int, error) {
	if m.Error != nil {
		return 0, m.Error
	}

	return 5, nil
}

func (m *MockEquation) Percentile() (int, error) {
	if m.Error != nil {
		return 0, m.Error
	}

	return 10, nil
}
