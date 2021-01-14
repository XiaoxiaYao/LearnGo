package main

import "testing"

func TestAdd(t *testing.T) {
	testingData := []struct{ a, b, c int }{
		{3, 4, 7},
	}

	for _, item := range testingData {
		if result := add(item.a, item.b); result != item.c {
			t.Errorf("add(%d, %d); "+"got %d; expected %d", item.a, item.b, result, item.c)
		}
	}
}
