package main

import "fmt"

func main() {
	a:=[]int{1,2,3,4,5}
	clone:=make([]int,len(a))

	for i:=0; i<5; i++{
		clone[i]=a[len(a)-1-i]
	}

	fmt.Println(clone)

}
