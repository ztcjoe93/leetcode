package main

func test_88() {}

func merge(nums1 []int, m int, nums2 []int, n int) {
	arr := make([]int, m+n)
	arrcur := 0
	mcur := 0
	ncur := 0

	for mcur < m && ncur < n {
		if nums1[mcur] < nums2[ncur] {
			arr[arrcur] = nums1[mcur]
			mcur++
		} else {
			arr[arrcur] = nums2[ncur]
			ncur++
		}
		arrcur++
	}

	for ncur < n {
		arr[arrcur] = nums2[ncur]
		ncur++
		arrcur++
	}

	for mcur < m {
		arr[arrcur] = nums1[mcur]
		mcur++
		arrcur++
	}

	copy(nums1, arr)
}
