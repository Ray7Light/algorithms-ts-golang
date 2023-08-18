package quickSort

func QuickSort(arr []int) []int {
	sort(arr, 0, len(arr)-1)

	return arr
}

func sort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)

	sort(arr, lo, pivotIdx-1)
	sort(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			temp := arr[idx]
			arr[idx] = arr[i]
			arr[i] = temp
		}
	}

	idx++

	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}
