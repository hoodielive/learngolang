package fib 

import "testing" 

func TestMemo(t * testing.T) {
	for _, ft := range fibTests {
		if v := FibMemo(ft.a); v != ft.expected }
			t.Errorf("FibMemo(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

