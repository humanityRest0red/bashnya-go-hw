package main

import "testing"

// func TestProcess(t *testing.T) {
// 	tests := []struct {
// 		tc_name  string
// 		num      int
// 		expected int
// 	}{
// 		{
// 			tc_name:  "2141",
// 			num:      2141,
// 			expected: _,
// 		},
// 		{
// 			tc_name:  "23017",
// 			num:      23017,
// 			expected: _,
// 		},
// 		{
// 			tc_name:  "356123",
// 			num:      356_123,
// 			expected: _,
// 		},
// 		{
// 			tc_name:  "451000",
// 			num:      451_000,
// 			expected: _,
// 		},
// 		{
// 			tc_name:  "500000",
// 			num:      500_000,
// 			expected: _,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.tc_name, func(t *testing.T) {
// 			result, _ := Process(tc.num)
// 			if result != tc.expected {
// 				t.Errorf("\nExpected: %v|\nBut Got:  %v|", tc.expected, result)
// 			}
// 		})
// 	}
// }

func TestNumIntoWords(t *testing.T) {
	tests := []struct {
		tc_name  string
		num      int
		expected string
	}{
		{
			tc_name:  "2141",
			num:      2141,
			expected: "Две тысячи сто сорок один",
		},
		{
			tc_name:  "23017",
			num:      23017,
			expected: "Двадцать три тысячи семнадцать",
		},
		{
			tc_name:  "356_123",
			num:      356_123,
			expected: "Триста пятьдесят шесть тысяч сто двадцать три",
		},
		{
			tc_name:  "451_000",
			num:      451_000,
			expected: "Четыреста пятьдесят одна тысяча",
		},
		{
			tc_name:  "500_000",
			num:      500_000,
			expected: "Пятьсот тысяч",
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := NumIntoWords(tc.num)
			if result != tc.expected {
				t.Errorf("\nExpected: %v|\nBut Got:  %v|", tc.expected, result)
			}
		})
	}
}
