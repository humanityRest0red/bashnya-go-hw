package stack

import (
	"testing"
)

func stackEqSlice[T comparable](stack *Stack[T], slice []T) bool {
	if uint(len(slice)) != stack.Size() {
		return false
	}

	for _, v := range slice {
		elem, _ := stack.Pop()
		if v != elem {
			return false
		}
	}

	return true
}

func TestStackPush_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		elem     int
		expected []int
	}{
		{
			tc_name:  "Valid input",
			elem:     -1,
			expected: []int{-1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			s := Stack[int]{}
			s.Push(tc.elem)
			if !stackEqSlice(&s, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, s)
			}
		})
	}
}

// func TestStackPush_Table(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    string
// 		expected []int
// 		is_err   bool
// 	}{
// 		{
// 			name:     "Valid input",
// 			input:    "1 2 -3 4 0",
// 			expected: []int{1, 2, -3, 4, 0},
// 			is_err:   false,
// 		},
// 		{
// 			name:     "Invalid float input",
// 			input:    "1 2.3",
// 			expected: []int{},
// 			is_err:   true,
// 		},
// 		{
// 			name:     "Invalid char input",
// 			input:    "1 2 -3 4 q",
// 			expected: []int{},
// 			is_err:   true,
// 		},
// 		{
// 			name:     "Empty input",
// 			input:    "",
// 			expected: []int{},
// 			is_err:   true,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			r := strings.NewReader(tc.input)

// 			nums, err := inputNums(r)

// 			if (err != nil) != tc.is_err {
// 				t.Errorf("Expected %v, but got %v", tc.is_err, err)
// 			}

// 			if err == nil && !reflect.DeepEqual(nums, tc.expected) {
// 				t.Errorf("Expected %v, but got %v", tc.expected, nums)
// 			}
// 		})
// 	}
// }
