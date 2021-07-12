package main

import "fmt"

func main()  {
	// 슬라이스
	items := []int{2, 3, 8}

	// foreach 반복문
	for key, value := range items {
		fmt.Println("key:", key, " value:", value)
	}
}
