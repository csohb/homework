package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	stdin :=bufio.NewReader(os.Stdin)

	var a int
	var b int

	_,err := fmt.Scanln(&a,&b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 한줄바꿈 나올때까지 읽어라
	} else {
		fmt.Print(a,b)
	}

	_,err = fmt.Scanln(&a,&b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 한줄바꿈 나올때까지 읽어라
	} else {
		fmt.Print(a,b)
	}
}
