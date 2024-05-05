package findSimple

import (
	"reflect"
	"testing"
)


func TestCheckListSimple(t *testing.T) {
	testTable := []struct {
		name     string
		min      int
		max      int
		expected []int
	}{
		{
			name:     "Simples between 10 and 20",
			min:      10,
			max:      20,
			expected: []int{11, 13, 17, 19},
		},
		{
			name:     "Simples between 1 and 10",
			min:      1,
			max:      10,
			expected: []int{2, 3, 5, 7},
		},
		{
			name:     "Simples between 100 and 110",
			min:      100,
			max:      110,
			expected: []int{101, 103, 107, 109},
		},
	}

	for _, testCase := range testTable {
		result := CheckListSimple(testCase.min, testCase.max)
		if !reflect.DeepEqual(testCase.expected, result) {
			t.Errorf("Test %s failed. Expected %v, got %v", testCase.name, testCase.expected, result)
		}
	}
}
