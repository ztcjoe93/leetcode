package main

func test_70() {
	eval(2, 2, climbStairs(2))
	eval(3, 3, climbStairs(3))
	eval(4, 5, climbStairs(4))
	eval(5, 8, climbStairs(5))
	eval(6, 13, climbStairs(6))
	eval(7, 21, climbStairs(7))
}

func climbStairs(n int) int {
	arr := make([]int, n+1)
	arr[0] = 0

	for i := 1; i <= n; i++ {
		if i == 1 {
			arr[i] = 1
		} else if i == 2 {
			arr[i] = 2
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}

	return arr[len(arr)-1]
}
