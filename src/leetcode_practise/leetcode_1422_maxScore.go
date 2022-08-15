package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func maxScore(s string) int {
	total1 := 0
	for _, str := range s {
		//fmt.Printf("%c\n", str)
		if str == '1' {
			total1 += 1
		}
	}
	//fmt.Printf("%d\n", total1)

	lens := len(s)
	var (
		left0  int
		right1 int
	)
	if s[0] == '0' {
		left0 = 1
		right1 = total1
	} else {
		left0 = 0
		right1 = total1 - 1
	}
	res := left0 + right1

	//fmt.Printf("left0 = %d, right1 = %d\n", left0, right1)

	for i := 1; i < lens-1; i++ {
		if s[i] == '0' {
			left0 += 1
		} else {
			right1 -= 1
		}
		//fmt.Printf("for left0 = %d, right1 = %d\n", left0, right1)
		res = int(math.Max(float64(left0+right1), float64(res)))
	}

	return res
}

func maxScore2(s string) int {
	res := int('1'-s[0]) + strings.Count(s[1:], "1")
	score := res
	for _, value := range s[1 : len(s)-1] {
		if value == '0' {
			score++
		} else {
			score--
		}
		if score > res {
			res = score
		}
	}
	return res

}

// 随机字符串
func GetRandomString(n int) string {
	// 随机种子
	rand.Seed(time.Now().UnixNano())
	str := "01"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func main() {

	str1 := "1111"
	str1 = GetRandomString(500)
	fmt.Printf("str1 = %v\n", str1)

	res := maxScore(str1)
	fmt.Printf("maxScore res = %d\n", res)

	res = maxScore2(str1)
	fmt.Printf("maxScore2 res = %d\n", res)
}
