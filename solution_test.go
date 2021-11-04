package main

import (
	"testing"
)

func TestTask178b(t *testing.T) {

	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{2, 5, 15, 18, 21, 15}, 2},
		{[]int{3, 6, 14, 17, 20}, 2},
		{[]int{-3, 6, -15, 17, 20}, 2},
	}

	for _, test := range tests {
		if output := Task178b(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestTask178v(t *testing.T) {

	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{2, 5, 15, 18, 21}, 0},
		{[]int{3, 9, 14, 16, 20}, 1},
		{[]int{-3, 6, -15, 17, 20}, 0},
	}

	for _, test := range tests {
		if output := Task178v(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestTask554(t *testing.T) {

	var tests = []struct {
		input    int
		expected string
	}{
		{5, "3^2 + 4^2 = 5^2"},
		{8, "3^2 + 4^2 = 5^2"},
	}

	for _, test := range tests {
		if output := Task554(test.input); output[0][:(len(output[0])-2)] != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output[0][:(len(output[0])-2)])
		}
	}
}

func TestIsSqrt(t *testing.T) {
	var tests = []struct {
		input    int
		expected bool
	}{
		{9, true},
		{4, true},
		{5, false},
		{-1, false},
	}

	for _, test := range tests {
		if output := IsSqrt(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestConsoleReader(t *testing.T) {
	var tests = []struct {
		expected []int
		input    string
	}{
		{[]int{1, 3}, "2.4 1 3"},
		{[]int{-1, 2, 3}, "-1 2 3"},
		{[]int{3, 1}, "gugf 3 1"},
	}

	for _, test := range tests {
		output := ConsoleReader(test.input)
		errCount := 0
		for id, val := range output {
			if test.expected[id] != val {
				errCount++
			}
		}
		if errCount > 0 {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
