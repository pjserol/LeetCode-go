package main

import "log"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	log.Println(twoSumOnePass(nums, target))
	log.Println(twoSumTwoPass(nums, target))
	log.Println(twoSumBruteForce(nums, target))
}

func twoSumOnePass(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		j := m[complement]
		if j != 0 {
			return []int{m[complement] - 1, i}
		}
		m[nums[i]] = i + 1
	}
	return nil
}

func twoSumTwoPass(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		j := m[complement]
		if j != 0 && m[complement] != i {
			return []int{i, m[complement]}
		}
	}

	return nil
}

func twoSumBruteForce(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
