package sort

import (
	"sync"
)

func QuickSort(data []Element, rule func(Element, Element) bool) []Element {
	if len(data) <= 1 {
		return data
	}

	current := data[0]
	var left, right []Element

	for k, v := range data {
		if k == 0 {
			continue
		}

		if rule(current, v) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		left = QuickSort(left, rule)
		wg.Done()
	}()

	go func() {
		right = QuickSort(right, rule)
		wg.Done()
	}()
	wg.Wait()

	return append(append(left, current), right...)
}
