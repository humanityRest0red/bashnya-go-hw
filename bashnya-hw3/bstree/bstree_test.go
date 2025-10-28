package bstree

import (
	"testing"
)

func TestBSTreeInsert_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		b        *BSTree[int]
		elem     int
		expected *BSTree[int]
	}{
		{
			tc_name:  "Was empty, added 1 elem",
			b:        &BSTree[int]{},
			elem:     2,
			expected: New(2),
		},
		{
			tc_name:  "Added copy of the elem",
			b:        New(2),
			elem:     2,
			expected: New(2, 2),
		},
		{
			tc_name:  "Common test",
			b:        New(1, 5, 10, 11),
			elem:     6,
			expected: New(1, 5, 6, 10, 11),
		},
		{
			tc_name:  "Added copy of the elem",
			b:        New(5, 11, 10),
			elem:     1,
			expected: New(1, 5, 11, 10),
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			tc.b.Insert(tc.elem)
			if !tc.b.Eq(tc.expected) {
				t.Errorf("Expected %v contains %v", tc.b, tc.elem)
			}
		})
	}
}

func TestBSTreeRemove_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		b        *BSTree[int]
		elem     int
		expected *BSTree[int]
	}{
		{
			tc_name:  "Empty BST",
			b:        &BSTree[int]{},
			elem:     2,
			expected: &BSTree[int]{},
		},
		{
			tc_name:  "Removing the only one elem",
			b:        New(2),
			elem:     2,
			expected: &BSTree[int]{},
		},
		{
			tc_name:  "Common test",
			b:        New(1, 5, 10, 6, 11),
			elem:     6,
			expected: New(1, 5, 10, 11),
		},
		{
			tc_name:  "Root remove",
			b:        New(2, 1, 3),
			elem:     2,
			expected: New(1, 3),
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			tc.b.Remove(tc.elem)
			if !tc.b.Eq(tc.expected) {
				t.Errorf("Expected %#v contains %v", tc.b, tc.elem)
			}
		})
	}
}

func TestBSTreeFind_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		b        *BSTree[int]
		elem     int
		expected bool
	}{
		{
			tc_name:  "Empty",
			b:        &BSTree[int]{},
			elem:     42,
			expected: false,
		},
		{
			tc_name:  "Size = 1 and contains",
			b:        New(42),
			elem:     42,
			expected: true,
		},
		{
			tc_name:  "Size = 1 and not contains",
			b:        New(412),
			elem:     42,
			expected: false,
		},
		{
			tc_name:  "Not contains",
			b:        New(1, 2, 3, 4, 6, 7),
			elem:     5,
			expected: false,
		},
		{
			tc_name:  "Contains",
			b:        New(1, 2, 3, 4, 6, 7),
			elem:     6,
			expected: true,
		},
		{
			tc_name:  "WTF",
			b:        New(7, 7, 7, 7, 7, 7, 7, 7, 7),
			elem:     7,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := tc.b.Find(tc.elem)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestBSTreeDepth_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		b        *BSTree[int]
		expected uint
	}{
		{
			tc_name:  "Empty",
			b:        &BSTree[int]{},
			expected: 0,
		},
		{
			tc_name:  "1 elem",
			b:        New(42),
			expected: 1,
		},
		{
			tc_name:  "2 elems",
			b:        New(42, 41),
			expected: 2,
		},
		{
			tc_name:  "3 elems but Depth is 2",
			b:        New(42, 41, 43),
			expected: 2,
		},
		{
			tc_name:  "BST as a List",
			b:        New(1, 2, 3, 4, 5),
			expected: 5,
		},
		{
			tc_name:  "BST as a List reversed",
			b:        New(100, 10, 1, 0),
			expected: 4,
		},
		{
			tc_name:  "Almost List",
			b:        New(1, 2, 3, 4, 0),
			expected: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := tc.b.Depth()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
