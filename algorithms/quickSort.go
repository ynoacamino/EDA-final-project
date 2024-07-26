package algorithms

func partition[V any](arr []V, low, high int, compareTo func(a, b V) int) ([]V, int) {
	pivot := arr[high]

	i := low

	for j := low; j < high; j++ {
		if compareTo(arr[j], pivot) < 0 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}

func quickSort[V any](arr []V, low, high int, compareTo func(a, b V) int) []V {
	if low < high {
		var p int
		arr, p = partition(arr, low, high, compareTo)
		arr = quickSort(arr, low, p-1, compareTo)
		arr = quickSort(arr, p+1, high, compareTo)
	}
	return arr
}

func QuickSortStart[V any](arr []V, compareTo func(a, b V) int) []V {
	return quickSort(arr, 0, len(arr)-1, compareTo)
}
