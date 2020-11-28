package stack

import "testing"

func TestSumFunc(t *testing.T) {
	type test struct {
		a      int
		b      int
		answer int
	}
	tests := []test{
		test{21, 21, 42},
		test{5, 6, 11},
		test{7, 10, 17},
	}
	for _, v := range tests {
		got := sumFunc(v.a, v.b)
		want := v.answer
		if got != want {
			t.Error("Expected", want, "Got", got)
		}
	}
	got := sumFunc(2, 3)
	want := 5
	if got != want {
		t.Error("Expected", want, "Got", got)
	}
}

func TestIsEmpty(t *testing.T) {
	type test struct {
		s      stack
		answer bool
	}

	var s1 stack
	var s2 stack = []interface{}{1, 2, 3}

	tests := []test{
		test{s1, true},
		test{s2, false},
	}

	for _, v := range tests {
		got := v.s.isEmpty()
		want := v.answer
		if got != want {
			t.Error("Expected", want, "Got", got)
		}
	}
}
