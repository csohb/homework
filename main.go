package main

import (
	"fmt"
	"goproject/config"
)

func main() {

	newCon := config.NewUnixConfig("config.ini")

	err := newCon.LoadConfig()
	if err != nil {
		panic(err)
	}

	bVal, err := newCon.GetConfigBoolean("Server", "bool")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(bVal)
	}

	iVal, err := newCon.GetConfigInteger("Server", "port")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(iVal)
	}

	sVal, err := newCon.GetConfigString("log", "port")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(sVal)
	}
}
