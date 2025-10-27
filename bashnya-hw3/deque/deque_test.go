package deque

import "testing"

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
