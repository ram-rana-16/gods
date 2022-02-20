package main


// https://leetcode.com/problems/two-sum/
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]

// compxity - o(n) time | O(n) space where n is lenght of array 
func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int, len(nums))
	for i, num := range nums {
		if pos, ok := mymap[target - num]; ok {
			return []int{pos, i}
		}else {
			mymap[num] = i  
		}
	}
	return nil 
}


