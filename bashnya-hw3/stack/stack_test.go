package stack

import "testing"

func TestStackPush_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Stack[int]
		elem     int
		expected *Stack[int]
	}{
		{
			tc_name:  "Was empty",
			s:        New[int](),
			elem:     -1,
			expected: New(-1),
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			tc.s.Push(tc.elem)
			if !tc.s.Eq(tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.s)
			}
		})
	}
}

func TestStackPop_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Stack[int]
		elem     int
		expected *Stack[int]
		err      error
	}{
		{
			tc_name:  "Was empty",
			s:        New[int](),
			elem:     0,
			expected: New[int](),
			err:      ErrEmptyStack,
		},
		{
			tc_name:  "Size = 1 elem",
			s:        New(42),
			elem:     42,
			expected: New[int](),
			err:      nil,
		},
		{
			tc_name:  "Many elems",
			s:        New(42, 3, 1, 0, 2, 0),
			elem:     0,
			expected: New(42, 3, 1, 0, 2),
			err:      nil,
		},
		{
			tc_name:  "Many elems",
			s:        New(42, 3, 1, 0, 2, 0),
			elem:     0,
			expected: New(42, 3, 1, 0, 2),
			err:      nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result, err := tc.s.Pop()

			if err != tc.err {
				t.Errorf("Expected %v, but got %v", tc.err, err)
			}
			if err == nil && !tc.s.Eq(tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.s)
			}
			if err == nil && tc.elem != result {
				t.Errorf("Expected %v, but got %v", tc.elem, result)
			}
		})
	}
}

func TestStackPeek_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Stack[int]
		elem     int
		expected *Stack[int]
		err      error
	}{
		{
			tc_name:  "Was empty",
			s:        New[int](),
			elem:     0,
			expected: New[int](),
			err:      ErrEmptyStack,
		},
		{
			tc_name:  "Size = 1 elem",
			s:        New(42),
			elem:     42,
			expected: New(42),
			err:      nil,
		},
		{
			tc_name:  "Many elems",
			s:        New(42, 3, 1, 0, 2, 0),
			elem:     0,
			expected: New(42, 3, 1, 0, 2, 0),
			err:      nil,
		},
		{
			tc_name:  "Many elems",
			s:        New(42, 3, 1, 0, 2, 0),
			elem:     0,
			expected: New(42, 3, 1, 0, 2, 0),
			err:      nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result, err := tc.s.Peek()

			if err != tc.err {
				t.Errorf("Expected %v, but got %v", tc.err, err)
			}
			if err == nil && !tc.s.Eq(tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.s)
			}
			if err == nil && tc.elem != result {
				t.Errorf("Expected %v, but got %v", tc.elem, result)
			}
		})
	}
}

func TestStackSize_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Stack[int]
		expected uint
	}{
		{
			tc_name:  "Empty",
			s:        New[int](),
			expected: 0,
		},
		{
			tc_name:  "1 elem",
			s:        New(1),
			expected: 1,
		},
		{
			tc_name:  "Many elems",
			s:        New(1, 5, 10, -1, 1, 2, 3),
			expected: 7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := tc.s.Size()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestStackIsEmpty_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Stack[int]
		expected bool
	}{
		{
			tc_name:  "Empty",
			s:        New[int](),
			expected: true,
		},
		{
			tc_name:  "1 elem",
			s:        New(1),
			expected: false,
		},
		{
			tc_name:  "Many elems",
			s:        New(1, 5, 10, -1, 1, 2, 3),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := tc.s.IsEmpty()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestStackClear_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Stack[int]
		expected bool
	}{
		{
			tc_name:  "Empty",
			s:        New[int](),
			expected: true,
		},
		{
			tc_name:  "1 elem",
			s:        New(1),
			expected: true,
		},
		{
			tc_name:  "Many elems",
			s:        New(1, 5, 10, -1, 1, 2, 3),
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			tc.s.Clear()
			result := tc.s.IsEmpty()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestStackEq_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s1       *Stack[rune]
		s2       *Stack[rune]
		expected bool
	}{
		{
			tc_name:  "Empty",
			s1:       New[rune](),
			s2:       New[rune](),
			expected: true,
		},
		{
			tc_name:  "1 elem, Eq",
			s1:       New('q'),
			s2:       New('q'),
			expected: true,
		},
		{
			tc_name:  "1 elem, Not Eq",
			s1:       New('q'),
			s2:       New('w'),
			expected: false,
		},
		{
			tc_name:  "One is empty and other is not",
			s1:       New('h', 'e', 'l', 'l'),
			s2:       New[rune](),
			expected: false,
		},
		{
			tc_name:  "Len is eq, but elems are not",
			s1:       New('h', 'e', 'l', 'l'),
			s2:       New('w', 'o', 'r', 'd'),
			expected: false,
		},
		{
			tc_name:  "Are almost equal",
			s1:       New('h', 'e', 'l', 'l'),
			s2:       New('h', 'e', 'l', 'l', 'o'),
			expected: false,
		},
		{
			tc_name:  "Are Equal",
			s1:       New('h', 'e', 'l', 'l', 'o'),
			s2:       New('h', 'e', 'l', 'l', 'o'),
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := tc.s1.Eq(tc.s2)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
