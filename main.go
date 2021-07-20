package main

import (
	"fmt"
	"goproject/config"
	"io/ioutil"
)

func main() {

	// byte로 파일 읽어옴
	// 이 처리 자체를 (readFile) NewUnixConfig에서 할 것
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
