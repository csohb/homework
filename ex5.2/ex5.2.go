package main

import "fmt"

func main()  {
	i:=1
	for i <=3 {
		//fmt.Println(i)
		i+=1
	}

	for j:=7; j<=9; j++{
		//fmt.Println(j)
	}

	if num:=9; num<0{
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
