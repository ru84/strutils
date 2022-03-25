package strutils

func RotateArray(arr []string) []string {
	lastItem := arr[len(arr)-1]

	var temp string = arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i+1], temp = temp, arr[i+1]
	}
	arr[0] = lastItem
	return arr
}

func InverseArray(input []int) []int {
	var output []int
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}
