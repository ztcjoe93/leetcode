package main

func test_69() {
	eval(4, 2, mySqrt(4))
	eval(8, 2, mySqrt(8))
}

func mySqrt(x int) int {
	var num int = 0
	for {
		if num*num > x {
			num--
			break
		}
		num++
	}

	return num
}
