package sort

// Insertion sort implementation
func Insertion(inputs []int)[]int {
	for i := 1; i < len(inputs); i++ {
		val := inputs[i]
		var j int
		// 1 23 2 3 4 
		for j = i-1; j >= 0 && inputs[j] > val; j-- {
			inputs[j+1] = inputs[j]
		}
		inputs[j+1] = val
	}
	return inputs
}