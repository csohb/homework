package main

import (
	"fmt"
	"math/rand"
	"time"
)



func main() {
	rand.Seed(time.Now().UnixNano())
	// 무작위 숫자 세개 만들기
	numbers:=MakeNumbers()

	cnt:=0
	// 무한루프
	for {
		cnt++
		// 사용자의 입력을 받는다.
		inputNumbers := InputNumbers()

		// 결과를 비교한다.
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)
		// 3S인가
		if IsGameEnd(result) {
			// 게임끝
			break
		}
	}

	// 몇번 만에 맞췄는지
	fmt.Printf("%d번 만에 맞췄습니다\n", cnt)
}

func MakeNumbers() [3]int {
	// 0~9사이의 겹치지 않는 무작위 숫자 3개를 반환
	var rst [3]int

	for i:=0; i<3; i++{
		for {
			n:=rand.Intn(10)
			duplicated:=false
			for j:=0; j<i; j++{
				if rst[j] == n {
					// 겹치면
					// 다시 뽑게 돌아감
					duplicated=true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}
	}
	fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int{
	// 사용자로 부터 0~9 사이의 겹치지 않는 3개의 숫자를 입력받음
	var rst [3]int

	for {
		var no int
		_, err := fmt.Scanf("%d",&no)
		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}
		idx:=0
		for no > 0 {
			n:=no%10
			no =no/10
			rst[idx] = n
			idx++
		}
	}

	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) bool{
	return true
}

func PrintResult(result bool){
	fmt.Println(result)
}

func IsGameEnd(result bool) bool {
	return result
}