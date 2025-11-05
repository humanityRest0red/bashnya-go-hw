package myunique

import (
	"reflect"
	"testing"
)

func TestUniqueLines_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		options  Options
		input    []string
		expected []string
	}{
		{
			tc_name: "No flags, No duplicates",
			options: Options{},
			input: []string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
			expected: []string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		{
			tc_name: "No flags, Any duplicates",
			options: Options{},
			input: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			expected: []string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		{
			tc_name: "-c flag",
			options: Options{c: true},
			input: []string{
				"Do a barrel roll.",
				"Wake up, Mr. Freeman.",
				"Do a barrel roll.",
				"",
				"I am the one who knocks.",
				"Judgement!",
				"Judgement!",
				"Foolishness, machine. Foolishness.",
				"Judgement!",
			},
			expected: []string{
				"2 Do a barrel roll.",
				"1 Wake up, Mr. Freeman.",
				"1 ",
				"1 I am the one who knocks.",
				"3 Judgement!",
				"1 Foolishness, machine. Foolishness.",
			},
		},
		{
			tc_name: "-d flag",
			options: Options{d: true},
			input: []string{
				"Do a barrel roll.",
				"Wake up, Mr. Freeman.",
				"Do a barrel roll.",
				"",
				"I am the one who knocks.",
				"Judgement!",
				"Judgement!",
				"Foolishness, machine. Foolishness.",
				"Judgement!",
			},
			expected: []string{
				"Do a barrel roll.",
				"Judgement!",
			},
		},
		{
			tc_name: "-u flag",
			options: Options{u: true},
			input: []string{
				"Do a barrel roll.",
				"Wake up, Mr. Freeman.",
				"Do a barrel roll.",
				"",
				"I am the one who knocks.",
				"Judgement!",
				"Judgement!",
				"Foolishness, machine. Foolishness.",
				"Judgement!",
			},
			expected: []string{
				"Wake up, Mr. Freeman.",
				"",
				"I am the one who knocks.",
				"Foolishness, machine. Foolishness.",
			},
		},
		{
			tc_name: "-i flag",
			options: Options{i: true},
			input: []string{
				"Do a barrel roll.",
				"Wake up, Mr. Freeman.",
				"Do a Barrel roll.",
				"",
				"JuDGement!",
				"Judgement!",
				"Judgement!!!!",
			},
			expected: []string{
				"Do a barrel roll.",
				"Wake up, Mr. Freeman.",
				"",
				"JuDGement!",
				"Judgement!!!!",
			},
		},
		{
			tc_name: "-s flag",
			options: Options{s: 1},
			input: []string{
				"Do a barrel roll.",
				"Wake up, Mr. Freeman.",
				"To a barrel roll.",
				"",
				"Judgement!",
				"Mudgement!",
				"Midgement!",
			},
			expected: []string{
				"Do a barrel roll.",
				"Wake up, Mr. Freeman.",
				"",
				"Judgement!",
				"Midgement!",
			},
		},
		{
			tc_name: "-f flag",
			options: Options{f: 1},
			input: []string{
				"I a barrel roll.",
				"Wake up, Mr. Freeman.",
				"WE a barrel roll.",
				"",
				"Judgement! a",
				"Midgement! a",
				"Judgement!!!! a",
			},
			expected: []string{
				"I a barrel roll.",
				"Wake up, Mr. Freeman.",
				"",
				"Judgement! a",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := uniqueLines(tc.input, tc.options)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Errorf("\nExpected %v\n     got %v", tc.expected, result)
			}
		})
	}
}

func TestIsColision_Table(t *testing.T) {
	tests := []struct {
		tc_name  string
		options  Options
		expected error
	}{
		{
			tc_name:  "Only -c",
			options:  Options{c: true},
			expected: nil,
		},
		{
			tc_name:  "Only -d",
			options:  Options{d: true},
			expected: nil,
		},
		{
			tc_name:  "Only -u",
			options:  Options{u: true},
			expected: nil,
		},
		{
			tc_name:  "-c -d",
			options:  Options{c: true, d: true},
			expected: ErrCDUFlags,
		},
		{
			tc_name:  "-c -u",
			options:  Options{c: true, u: true},
			expected: ErrCDUFlags,
		},
		{
			tc_name:  "-d -u",
			options:  Options{d: true, u: true},
			expected: ErrCDUFlags,
		},
		{
			tc_name:  "-c -d -u",
			options:  Options{c: true, d: true, u: true},
			expected: ErrCDUFlags,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tc_name, func(t *testing.T) {
			result := tc.options.isColision()
			if tc.expected != result {
				t.Errorf("\nExpected %v\n     got %v", tc.expected, result)
			}
		})
	}
}
