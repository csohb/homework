package main

import (
	"fmt"
	config "goproject/config"
)

func main() {
	File := `
	[Server] 
	"port"=20
	"ip"="1.1.2.3"
	"except=555`

	newCon := config.NewUnixConfig(File)
	fmt.Println(newCon)

	err := newCon.LoadConfig(File)
	if err != nil {
		panic(err)
	}

}
