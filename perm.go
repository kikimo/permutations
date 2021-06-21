package main

import (
	"os"
	"fmt"
	"strconv"
)

func Perm(nums []int) [][]int {
	ret := [][]int{}

	for i := range nums {
		v := nums[i]
		others := append([]int{}, nums[:i]...)
		others = append(others, nums[i+1:]...)
		subRets := Perm(others)
		if len(subRets) == 0 {
			ret = append(ret, []int{v})
		} else {
			for j := range subRets {
				o := append(subRets[j], v)
				ret = append(ret, o)
			}
		}
	}

	return ret
}

func main() {
	ns := os.Args[1]
	n, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}

	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, i + 1)
	}
	ret := Perm(nums)
	fmt.Printf("sz: %d\n", len(ret))
}

