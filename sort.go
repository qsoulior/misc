package alg

import "math/rand"

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

// Рекурсивная быстрая сортировка, работающая за O(n*log(n))
func QuickSort[S ~[]E, E any](arr S, cmp func(E, E) int) {
	n := len(arr) // длина массива
	if n < 2 {
		// Если массив содержит 0 или 1 элемент, он отсортирован.
		return
	}

	left, right := 0, n-1 // границы массива (индексы первого и последнего элементов)
	pivot := rand.Intn(n) // индекс опорного элемента

	// Меняем местами опорный и последний элементы.
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := 0; i < right; i++ {
		// Если элемент меньше опорного,
		// отправляем его в начало после других отправленных в начало элементов.
		if cmp(arr[i], arr[right]) < 0 {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Меняем местами опорный элемент и первый элемент,
	// который больше или равен опорному.
	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left], cmp)
	QuickSort(arr[left+1:], cmp)
}
