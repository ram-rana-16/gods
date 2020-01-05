package sort


// SelectionsSort implementation
func SelectionsSort(inputs []int) []int {
	for i :=0; i < len(inputs); i++ {
		currentMin := inputs[i]
		minPos := i
		for j := i+1; j < len(inputs); j++ {
			if inputs[j] < currentMin {
				currentMin = inputs[j]
				minPos = j
			}
		}
		inputs[i], inputs[minPos] = inputs[minPos], inputs[i]
	}
	return inputs	
}