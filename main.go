package main

import (
	"fmt"
	"goproject/config"
	"io/ioutil"
)

func main() {

	// byte로 파일 읽어옴
	data, err := ioutil.ReadFile("config.ini")
	if err != nil {
		panic(err)
	}

	// byte data 형변환 필요
	newData := string(data)

	newCon := config.NewUnixConfig(newData)

	err = newCon.LoadConfig(newData)
	if err != nil {
		panic(err)
	}

	newCon.Printconfig()

	val, err := newCon.GetConfigBoolean("Server", "bool")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(val)
	}

}
