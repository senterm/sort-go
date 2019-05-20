package sort

import (
"testing"
)

func Test_Bubble_MaxAtHead(t *testing.T) {
	source := []Element{58, 1, 3, 6, 23, 7, 2, 4, 9, 13}
	expected := []Element{1, 2, 3, 4, 6, 7, 9, 13, 23, 58}
	rule := func(a Element, b Element) bool {
		return a.(int) >= b.(int)
	}

	sorted := BubbleSort(source, rule)
	t.Log(sorted)

	for k, v := range sorted {
		if v != expected[k] {
			t.Errorf("fail to sort \n, source: %v \n, expected: %v \n, current: %v", source, expected, sorted)
			break
		}
	}
}

func Test_Bubble_MinAtHead(t *testing.T) {
	source := []Element{1, 3, 6, 23, 7, 2, 4, 9, 13}
	expected := []Element{1, 2, 3, 4, 6, 7, 9, 13, 23}
	rule := func(a Element, b Element) bool {
		return a.(int) >= b.(int)
	}

	sorted := BubbleSort(source, rule)
	t.Log(sorted)

	for k, v := range sorted {
		if v != expected[k] {
			t.Errorf("fail to sort \n, source: %v \n, expected: %v \n, current: %v", source, expected, sorted)
			break
		}
	}
}

func Test_Bubble_Null(t *testing.T) {
	source := []Element{}
	expected := []Element{}
	rule := func(a Element, b Element) bool {
		return a.(int) >= b.(int)
	}

	sorted := BubbleSort(source, rule)
	t.Log(sorted)

	for k, v := range sorted {
		if v != expected[k] {
			t.Errorf("fail to sort \n, source: %v \n, expected: %v \n, current: %v", source, expected, sorted)
			break
		}
	}
}

