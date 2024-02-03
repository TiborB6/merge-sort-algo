package main

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	firstHalf := arr[:len(arr)/2]
	secondHalf := arr[len(arr)/2:]

	sortedFirstHalf := mergeSort(firstHalf)
	sortedSecondHalf := mergeSort(secondHalf)

	arr = []int{}

	i := 0
	j := 0

	for i < len(sortedFirstHalf) && j < len(sortedSecondHalf) {
		if sortedFirstHalf[i] <= sortedSecondHalf[j] {
			arr = append(arr, sortedFirstHalf[i])
			i++
		} else {
			arr = append(arr, sortedSecondHalf[j])
			j++
		}
	}

	for i < len(sortedFirstHalf) {
		arr = append(arr, sortedFirstHalf[i])
		i++
	}

	for j < len(sortedSecondHalf) {
		arr = append(arr, sortedSecondHalf[j])
		j++
	}

	return arr
}
