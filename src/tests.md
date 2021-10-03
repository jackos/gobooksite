```go
package main

import "testing"

// Loop tests
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, test := range tests {
		got, err := divide(test.dividend, test.divisor)
		if test.isErr {
			if err == nil {
				t.Errorf("[%s]: Expected an error but did not get one", test.name)
			}
		} else {
			if err != nil {
				t.Error("unexpected error:", err.Error())
			}
		}

		if got != test.expected {
			t.Errorf("failed %s: expected %f but got %f", test.name, test.expected, got)
		}

	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when you should have")
	}
}
```