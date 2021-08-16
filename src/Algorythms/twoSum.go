package twoSum

// Looking for two numbers that add up to the sum
func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, v := range nums {
		j, ok := lookup[-v]
		// By storing value - target in a map, we have the value
		// required to reach the sum as the key, and i gives us
		// the index of the original number
		lookup[v-target] = i
		if ok {
			return []int{j, i}
		}
	}
	return []int{}
}
