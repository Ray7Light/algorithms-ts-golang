package main 

func BinarySearchList (array []int, searchedValue int) bool {
	lo := 0
	hi := len(array)

	for hi > lo {
		// First lo is the starting point of search in the array
		// hi - lo calculates the length of the array that is left for searching
		// this is then divided by 2 to denote the middle of the new length of array
		// offset is then added to determine the middle of the original array.
		// So in the case of low = 11 hi = 17.
		// 17 - 11 = 6
		// 6 / 2 = 3
		// 11 + 3 = 14
		// The new middle is 14.
		m := lo + (hi - lo) / 2
		v := array[m]
		
		if v == searchedValue {
			return true 
		} else if v > searchedValue {
			hi = m
		} else {
			lo = m + 1
		}
	}

	return false
}
