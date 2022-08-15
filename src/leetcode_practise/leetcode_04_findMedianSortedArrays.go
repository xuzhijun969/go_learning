package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// fmt.Println(len(nums1))
	// fmt.Println(len(nums2))

	var numsNew []int
	// 合并切片
	numsNew = append(nums1, nums2...)
	//fmt.Println(numsNew)

	lennumsNew := len(numsNew)
	// 切片排序
	sort.Ints(numsNew)
	//fmt.Println(numsNew)

	//fmt.Println(lennumsNew)
	var result float64

	if lennumsNew%2 != 0 {
		result = float64(numsNew[lennumsNew/2])
	} else {
		result = float64(numsNew[lennumsNew/2]+numsNew[lennumsNew/2-1]) / 2
	}
	return result
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	// 优化时间复杂度
	// fmt.Println(len(nums1))
	// fmt.Println(len(nums2))

	// 使用游标
	lenNums1 := len(nums1)
	lenNums2 := len(nums2)

	lenHalf := (lenNums1 + lenNums2) / 2
	lenMode := (lenNums1 + lenNums2) % 2

	i := 0
	j := 0
	var numsNew []int

	var nextData int
	for {
		if i+j > lenHalf {
			break
		}

		if i < lenNums1 && j < lenNums2 {
			if nums1[i] < nums2[j] {
				nextData = nums1[i]
				i++
			} else {
				nextData = nums2[j]
				j++
			}
		} else if i >= lenNums1 && j < lenNums2 {
			nextData = nums2[j]
			j++
		} else if i < lenNums1 && j >= lenNums2 {
			nextData = nums1[i]
			i++
		} else {
			break
		}
		numsNew = append(numsNew, nextData)
	}

	var result float64

	if lenMode != 0 {
		result = float64(numsNew[i+j-1])
	} else {
		result = float64(numsNew[i+j-1]+numsNew[i+j-2]) / 2
	}

	return result
}

func findMedianSortedArrays3(numspoint1 *[]int, numspoint2 *[]int) float64 {
	// 优化时间复杂度，使用指针
	// fmt.Println(len(nums1))
	// fmt.Println(len(nums2))

	// 使用游标
	nums1 := *numspoint1
	nums2 := *numspoint2
	lenNums1 := len(nums1)
	lenNums2 := len(nums2)

	lenHalf := (lenNums1 + lenNums2) / 2
	lenMode := (lenNums1 + lenNums2) % 2

	i := 0
	j := 0
	var numsNew []int

	var nextData int
	for {
		if i+j > lenHalf {
			break
		}

		if i < lenNums1 && j < lenNums2 {
			if nums1[i] < nums2[j] {
				nextData = nums1[i]
				i++
			} else {
				nextData = nums2[j]
				j++
			}
		} else if i >= lenNums1 && j < lenNums2 {
			nextData = nums2[j]
			j++
		} else if i < lenNums1 && j >= lenNums2 {
			nextData = nums1[i]
			i++
		} else {
			break
		}
		numsNew = append(numsNew, nextData)
	}

	var result float64

	if lenMode != 0 {
		result = float64(numsNew[i+j-1])
	} else {
		result = float64(numsNew[i+j-1]+numsNew[i+j-2]) / 2
	}

	return result
}

func makeSlice(nums1 []int) []int {
	time.Sleep(100)
	rand.Seed(time.Now().UnixNano())
	sliceNum := 1000
	intSegment := 2000000 / sliceNum

	for m := 0; m < sliceNum; m++ {
		nums1 = append(nums1, rand.Intn(intSegment)-1000000+intSegment*m)
	}
	return nums1
}

func main() {
	var nums1 []int
	var nums2 []int
	nums1 = makeSlice(nums1)
	nums2 = makeSlice(nums2)
	//nums1 = []int{1, 2}
	//nums2 = []int{2, 3}
	//fmt.Println(nums1)
	//fmt.Println(nums2)

	res := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("findMedianSortedArrays res = %f\n", res)

	res = findMedianSortedArrays2(nums1, nums2)
	fmt.Printf("findMedianSortedArrays2 res = %f\n", res)

	res = findMedianSortedArrays3(&nums1, &nums2)
	fmt.Printf("findMedianSortedArrays3 res = %f\n", res)

	timeBegin := time.Now().UnixNano()
	//fmt.Println(timeBegin)
	for i := 0; i < 10000; i++ {
		findMedianSortedArrays(nums1, nums2)
	}
	fmt.Printf("findMedianSortedArrays time_spend: %v\n", time.Now().UnixNano()-timeBegin)
	//fmt.Printf("findMedianSortedArrays res = %v\n", res)

	timeBegin = time.Now().UnixNano()
	for i := 0; i < 10000; i++ {
		findMedianSortedArrays2(nums1, nums2)
	}
	//fmt.Printf("findMedianSortedArrays2 res = %v\n", res)
	fmt.Printf("findMedianSortedArrays2 time_spend: %v\n", time.Now().UnixNano()-timeBegin)

}
