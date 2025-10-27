package deque

import "testing"

func TestDequePushFront_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Deque[int]
		elem     int
		expected *Deque[int]
	}{
		{
			tc_name:  "Was empty",
			s:        New[int](),
			elem:     -1,
			expected: New(-1),
		},
		{
			tc_name:  "Common",
			s:        New(1),
			elem:     42,
			expected: New(42, 1),
		},
		{
			tc_name:  "WTF",
			s:        New(7, 7, 7, 7),
			elem:     7,
			expected: New(7, 7, 7, 7, 7),
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			tc.s.PushFront(tc.elem)
			if !tc.s.Eq(tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.s)
			}
		})
	}
}

func TestDequePushBack_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Deque[int]
		elem     int
		expected *Deque[int]
	}{
		{
			tc_name:  "Was empty",
			s:        New[int](),
			elem:     -1,
			expected: New(-1),
		},
		{
			tc_name:  "Common",
			s:        New(1),
			elem:     42,
			expected: New(1, 42),
		},
		{
			tc_name:  "WTF",
			s:        New(7, 7, 7, 7),
			elem:     7,
			expected: New(7, 7, 7, 7, 7),
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			tc.s.PushBack(tc.elem)
			if !tc.s.Eq(tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.s)
			}
		})
	}
}

func TestDequePopBack_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Deque[int]
		elem     int
		expected *Deque[int]
		err      error
	}{
		{
			tc_name:  "Empty",
			s:        New[int](),
			elem:     0,
			expected: New[int](),
			err:      ErrEmptyDeque,
		},
		{
			tc_name:  "Size == 1",
			s:        New(1),
			elem:     1,
			expected: New[int](),
		},
		{
			tc_name:  "Common",
			s:        New(-3, 42, 0, 1, -5),
			elem:     -5,
			expected: New(-3, 42, 0, 1),
		},
		{
			tc_name:  "WTF",
			s:        New(7, 7, 7, 7),
			elem:     7,
			expected: New(7, 7, 7),
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result, err := tc.s.PopBack()
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

func TestDequePopFront_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Deque[int]
		elem     int
		expected *Deque[int]
		err      error
	}{
		{
			tc_name:  "Empty",
			s:        New[int](),
			elem:     0,
			expected: New[int](),
			err:      ErrEmptyDeque,
		},
		{
			tc_name:  "Size == 1",
			s:        New(1),
			elem:     1,
			expected: New[int](),
		},
		{
			tc_name:  "Common",
			s:        New(-3, 42, 0, 1, -5),
			elem:     -3,
			expected: New(42, 0, 1, -5),
		},
		{
			tc_name:  "WTF",
			s:        New(7, 7, 7, 7),
			elem:     7,
			expected: New(7, 7, 7),
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result, err := tc.s.PopFront()
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

func TestDequeSize_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Deque[int]
		expected uint
	}{
		{
			tc_name:  "Empty",
			s:        New[int](),
			expected: 0,
		},
		{
			tc_name:  "Size == 1",
			s:        New(1),
			expected: 1,
		},
		{
			tc_name:  "Common",
			s:        New(-3, 42, 0, 1, -5),
			expected: 5,
		},
		{
			tc_name:  "WTF",
			s:        New(7, 7, 7, 7),
			expected: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := tc.s.Size()
			if tc.expected != result {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestDequeIsEmpty_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Deque[int]
		expected bool
	}{
		{
			tc_name:  "Empty",
			s:        New[int](),
			expected: true,
		},
		{
			tc_name:  "Size == 1",
			s:        New(1),
			expected: false,
		},
		{
			tc_name:  "Common",
			s:        New(-3, 42, 0, 1, -5),
			expected: false,
		},
		{
			tc_name:  "WTF",
			s:        New(7, 7, 7, 7),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := tc.s.IsEmpty()
			if tc.expected != result {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestDequeClear_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		s        *Deque[float64]
		expected *Deque[float64]
	}{
		{
			tc_name:  "Empty",
			s:        New[float64](),
			expected: New[float64](),
		},
		{
			tc_name:  "Size == 1",
			s:        New(1.),
			expected: New[float64](),
		},
		{
			tc_name:  "Common",
			s:        New(-3.0, 42.1, 0, 1.9, -5),
			expected: New[float64](),
		},
		{
			tc_name:  "WTF",
			s:        New(7., 7., 7., 7.),
			expected: New[float64](),
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			tc.s.Clear()
			if !tc.s.Eq(tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.s)
			}
		})
	}
}

func TestDequeEq_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		d1       *Deque[rune]
		d2       *Deque[rune]
		expected bool
	}{
		{
			tc_name:  "Empty",
			d1:       New[rune](),
			d2:       New[rune](),
			expected: true,
		},
		{
			tc_name:  "1 elem, Eq",
			d1:       New('q'),
			d2:       New('q'),
			expected: true,
		},
		{
			tc_name:  "1 elem, Not Eq",
			d1:       New('q'),
			d2:       New('w'),
			expected: false,
		},
		{
			tc_name:  "One is empty and other is not",
			d1:       New('h', 'e', 'l', 'l'),
			d2:       New[rune](),
			expected: false,
		},
		{
			tc_name:  "Len is eq, but elems are not",
			d1:       New('h', 'e', 'l', 'l'),
			d2:       New('w', 'o', 'r', 'd'),
			expected: false,
		},
		{
			tc_name:  "Are almost equal",
			d1:       New('h', 'e', 'l', 'l'),
			d2:       New('h', 'e', 'l', 'l', 'o'),
			expected: false,
		},
		{
			tc_name:  "Are Equal",
			d1:       New('h', 'e', 'l', 'l', 'o'),
			d2:       New('h', 'e', 'l', 'l', 'o'),
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := tc.d1.Eq(tc.d2)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
