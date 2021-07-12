package main

import (
	"fmt"
	"os"
)

func main()  {
	file,err:=os.Open("myfile")
	if err!=nil{
		panic(err)
	}
	data:=make([]byte,3)
	file.Read(data) // read the data from file to the byte array 'data'
	fmt.Printf("The file data is %s\n",string(data))
	file.Close()

}
