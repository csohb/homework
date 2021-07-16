package main

import (
	config "goproject/config"
)

func main() {
	File := `
	[Server] 
	"port"=20
	"ip"="1.1.2.3"
	"except"=true

	[log]
	"path"="/home/user/log"

	[test]
	"aa"="lll"

	[hi]
	`

	newCon := config.NewUnixConfig(File)
	//fmt.Println(newCon)

	err := newCon.LoadConfig(File)
	if err != nil {
		panic(err)
	}

}
