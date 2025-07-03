package main

import (
	"testing"
)

type TestCase struct {
	name     string
	a        int
	b        int
	expected int
}

func TestAdd(t *testing.T) {
	tests := []TestCase{
		{"add 1 + 1 numbers", 1, 1, 2},
		{"add 2 + 5 numbers", 2, 5, 7},
	}

	for _, values := range tests {
		t.Run(values.name, func(t *testing.T) {
			result := Add(values.a, values.b)

			if values.expected != result {
				t.Errorf("Result %d for expression (%d + %d) is not expected %d", result, values.a, values.b, values.expected)
			}
		})
	}
}

func TestConvertNumberOrPanic(t *testing.T) {
	t.Run("valid number", func(t *testing.T) {
		got := ConvertNumber("42")
		if got != 42 {
			t.Errorf("expected 42, got %d", got)
		}
	})

	t.Run("invalid number should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic but got none")
			} else {
				expected := "failed to convert number"
				if r != expected {
					t.Errorf("expected panic %q, got %q", expected, r)
				}
			}
		}()

		ConvertNumber("not-a-number")
	})
}

func TestCalculator(t *testing.T) {
	type Scenario struct {
		TestCase
		operation   string
		shouldError bool
	}

	tests := []Scenario{
		{
			TestCase: TestCase{
				name:     "Add 1 + 1 equals 2",
				a:        1,
				b:        1,
				expected: 2,
			},
			operation: "+",
		},
		{
			TestCase: TestCase{
				name:     "Add 7 + 3 equals 10",
				a:        7,
				b:        3,
				expected: 10,
			},
			operation: "+",
		},
		{
			TestCase: TestCase{
				name:     "Subtract 7 - 3 equals 4",
				a:        7,
				b:        3,
				expected: 4,
			},
			operation: "-",
		},
		{
			TestCase: TestCase{
				name:     "Subtract 7 - (-3) equals 10",
				a:        7,
				b:        -3,
				expected: 10,
			},
			operation: "-",
		},
		{
			TestCase: TestCase{
				name:     "Multiply 7 * 3 equals 21",
				a:        7,
				b:        3,
				expected: 21,
			},
			operation: "*",
		},
		{
			TestCase: TestCase{
				name:     "Divide 7 / 7",
				a:        7,
				b:        7,
				expected: 1,
			},
			operation: "/",
		},
		{
			TestCase: TestCase{
				name: "Divide by 0",
				a:    7,
				b:    0,
			},
			operation:   "/",
			shouldError: true,
		},
		{
			TestCase: TestCase{
				name:     "should return invalid operation",
				a:        1,
				b:        2,
				expected: 2,
			},
			operation:   "**",
			shouldError: true,
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			// Panic safety wrapper
			defer func() {
				if r := recover(); r != nil {
					if !scenario.shouldError {
						t.Errorf("unexpected panic: %v", r)
					}
				}
			}()

			result := Calculator(scenario.a, scenario.b, scenario.operation)

			if result != scenario.expected {
				t.Errorf("expected %d, got %v", scenario.expected, result)
			}
		})
	}
}
