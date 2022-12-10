package main

func titleToNumber(columnTitle string) int {
	multiplier := 1
	sum := 0

	for i := len(columnTitle) - 1; i >= 0; i-- {
		sum += (int(columnTitle[i]) - 64) * multiplier
		multiplier *= 26
	}
	return sum
}
