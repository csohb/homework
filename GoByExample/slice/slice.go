package main

import "fmt"
func RemoveBack(a []int) []int{
	return a[:len(a)-1]
}
func main()  {

	/*s:= make([]string,3)
	fmt.Println("emp:",s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:",s)
	fmt.Println("get:",s[2])

	fmt.Println("len:",len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:",s)

	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy:",c)

	l := s[2:5]
	fmt.Println("sl1:",l)
	fmt.Println("sl2:",s[2:])
	fmt.Println("sl3:",s[:5])

	t:= []string{"g","h","i"}
	fmt.Println("dcl:",t)

	twoD:=make([][]int,3)
	for i:=0; i<3; i++{
		innerLen:=i+1
		twoD[i]=make([]int, innerLen)
		for j:=0; j<innerLen; j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d: ",twoD)

	a:= make([]int,2,4)
	b:= append(a,3)

	fmt.Println("a:",a,"b:",b)

	b[0]=4
	b[1]=3
	fmt.Println("a:",a,"b:",b)*/
	a:=[]int{1,2,3,4,5,6,7,8,9,10}

	for i:=0; i<5; i++{
		a=RemoveBack(a)
	}

	fmt.Println(a)
}
