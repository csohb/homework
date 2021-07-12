package main

import "fmt"

var g int = 10 // 패키지 전역변수

func main()  {
	var m int = 10
	{
		var s int = 50
		fmt.Println(m, s, g)
	}

}
