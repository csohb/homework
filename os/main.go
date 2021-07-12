package main

import (
	"fmt"
	"os"
)

func main()  {
	/*le, err:=os.Create("myfile")
	if err!=nil{
		panic(err)
	}
	data:=[]byte("Hello World\n")
	file.Write(data)
	file.WriteString("This is a String")
	file.Close()
	*/
	file,err:=os.Open("myfile")
	if err!=nil{
		panic(err)
	}
	data:=make([]byte,5)
	file.Seek(6,0) // seek the file pointer to the 6th byte of the data
	/*
		seek has two arguments, first one is the position of file pointer to read data
		From where in the file to read from 0 - beginning, 1-at current position, 2-from end
	 */
	file.Read(data) // read the data from file to the byte array 'data'
	fmt.Println(string(data))
	file.Close()
}
