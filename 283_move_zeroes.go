package main

func moveZeroes(nums []int)  {
    cur := 0
    for cur < len(nums) {
        if nums[cur] == 0 {
            backCur := cur+1
            for backCur < len(nums) {
                if nums[backCur] != 0 {
                    tmp := nums[backCur]
                    nums[backCur] = nums[cur]
                    nums[cur] = tmp
                    break
                }
                backCur++
            }
        }
        cur++
    }
}
