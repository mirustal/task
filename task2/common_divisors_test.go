package findCommonDivisors

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindCommonDivisors(t *testing.T) {
	testTable := []struct {
		name     string  
		numbers  []int
		expected []int
	}{
		{
			name:     "Common divisors of 42, 12, and 18",
			numbers:  []int{42, 12, 18},
			expected: []int{2, 3, 6},
		},
		{
			name:     "Common divisors of 100, 1000, 10000",
			numbers:  []int{100, 1000, 10000},
			expected: []int{2, 4, 5, 10, 20, 25, 50, 100},
		},
		{
			name:     "Common divisors of 10, 20, 30",
			numbers:  []int{10, 20, 30},
			expected: []int{2, 5, 10},
		},
	}

	for _, testCase := range testTable {
		result := findCommonDivisors(testCase.numbers)
		sort.Ints(result)  
		sort.Ints(testCase.expected) 


		if !reflect.DeepEqual(testCase.expected, result) {
			t.Errorf("Test %s failed. Expected %v, got %v", testCase.name, testCase.expected, result)
		}
	}
}