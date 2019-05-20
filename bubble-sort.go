package sort

func BubbleSort(data []Element, rule func(Element, Element) bool) []Element {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if rule(data[i], data[j]) {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
