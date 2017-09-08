package kotori

// Calculate the sum of an array
func Sum(arr []int) (result int) {
	result = 0
	for _, val := range arr {
		result += val
	}

	return
}

// return min element of an array
func Min(arr []float64) (result float64) {
	result = arr[0]
	for _, val := range arr {
		if val < result {
			result = val
		}
	}

	return
}

// return max element of an array
func Max(arr []float64) (result float64) {
	result = arr[0]
	for _, val := range arr {
		if val > result {
			result = val
		}
	}

	return
}
