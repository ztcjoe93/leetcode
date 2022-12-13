package main

func missingNumber(nums []int) int {
    var target = len(nums)
    
    for idx, num := range nums {
        target += idx - num
    }
    
    return target
}
