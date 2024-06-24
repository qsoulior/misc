package alg

// Сортировка выбором, работающая за O(n^2)
func SelectionSort[S ~[]E, E any](arr S, cmp func(E, E) int) {
	n := len(arr) // длина массива
	for i := 0; i < n; i++ {
		// Ищем индекс наименьшего элемента в неотсортированной части массива.
		min := i
		for j := i + 1; j < n; j++ {
			if cmp(arr[j], arr[min]) < 0 {
				min = j
			}
		}

		// Меняем местами текущий элемент и наименьший элемент неотсортированной части.
		arr[i], arr[min] = arr[min], arr[i]
	}
}
