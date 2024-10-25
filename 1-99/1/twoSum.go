package main 

func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)

	for key, value := range nums {
		found := target - value

		if found, ok := hashMap[found]; ok {
			return []int {found, value}
		} else {
			hashMap[value] = key
		}
	}
	return []int {-1, -1}
}