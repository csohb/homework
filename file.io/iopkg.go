package main

import "fmt"

func main()  {
	/*r:=strings.NewReader("some io.Reader stream to be read\n")

	if _,err:=io.Copy(os.Stdout, r); err!=nil{
		log.Fatal(err)
	}*/
	var a int
	var b int

	n,err:=fmt.Scanln(&a,&b) // 주소값 넣어줌
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n,a,b)
	}
}
