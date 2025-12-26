package main

import "testing"

func TestProcess(t *testing.T) {
	tests := []struct {
		tc_name  string
		num      int
		expected int
		is_err   bool
	}{
		{
			tc_name:  "-1: service error",
			num:      -1,
			expected: 36_082,
			is_err:   true,
		},
		{
			tc_name:  "1: service error",
			num:      1,
			expected: 36_082,
			is_err:   true,
		},
		{
			tc_name:  "-2",
			num:      -2,
			expected: 36_082,
			is_err:   false,
		},
		{
			tc_name:  "2",
			num:      2,
			expected: 36_082,
			is_err:   false,
		},
		{
			tc_name:  "3",
			num:      3,
			expected: 14_212,
			is_err:   false,
		},
		{
			tc_name:  "7: [%7==0]",
			num:      7,
			expected: 96649,
			is_err:   false,
		},
		{
			tc_name:  "9: [%9==0]",
			num:      9,
			expected: 0,
			is_err:   true,
		},
		{
			tc_name:  "10: service error",
			num:      10,
			expected: 0,
			is_err:   true,
		},
		{
			tc_name:  "339: service error",
			num:      10,
			expected: 0,
			is_err:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result, err := Process(tc.num)

			if (err != nil) != tc.is_err {
				t.Errorf("Expected %v, but got %v", tc.is_err, err)
			}

			if err == nil && result != tc.expected {
				t.Errorf("\nExpected: %v\nBut Got:  %v", tc.expected, result)
			}
		})
	}
}

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
		{
			tc_name:  "500_000_000",
			num:      500_000_000,
			expected: "Пятьсот миллионов",
		},
		{
			tc_name:  "12 тысяч",
			num:      12_345,
			expected: "Двенадцать тысяч триста сорок пять",
		},
		{
			tc_name:  "BIG",
			num:      1_234_567_890,
			expected: "Один миллиард двести тридцать четыре миллиона пятьсот шестьдесят семь тысяч восемьсот девяносто",
		},
	}
	// 12345679,

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := NumIntoWords(tc.num)
			if result != tc.expected {
				t.Errorf("\nExpected: %v|\nBut Got:  %v|", tc.expected, result)
			}
		})
	}
}
