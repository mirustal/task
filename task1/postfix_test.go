package fmtpostfix

import "testing"

func TestFmtPostfix(t *testing.T) {
	testTable := []struct {
		number   int
		expected string
	}{
		{
			number:   1,
			expected: "1 компьютер",
		},
		{
			number:   2,
			expected: "2 компьютера",
		},
		{
			number:   5,
			expected: "5 компьютеров",
		},
		{
			number:   11,
			expected: "11 компьютеров",
		},
		{
			number:   25,
			expected: "25 компьютеров",
		},
		{
			number:   41,
			expected: "41 компьютер",
		},
		{
			number:   1048,
			expected: "1048 компьютеров",
		},
		
	}

	for _, testCase := range testTable {
		result := fmtPostfix(testCase.number)
		if testCase.expected != result {
			t.Errorf("failed result for %d. Expected %s, got %s", testCase.number, testCase.expected, result)
		}
	}
}