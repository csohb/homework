package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)


// 사용자 인터페이스
type UnixConfigAPI interface {
	GetConfigInteger(section, key string) (int, error)
	GetConfigString(section, key string) (string, error)
	GetConfigBoolean(section, key string) (bool, error)
}


// section 유효성 체크
func checkValidSection(section string) (bool, error) {
	if section[len(section)-1] == ']' {
		return true, nil
	}
	return false,errors.New("Section is not valid")
}
// key 유효성 체크
func checkValidKey(key string) (bool, error)  {
	bKey:=[]byte(key)
	if bKey[0]=='"' && bKey[len(bKey)-1]=='"'{
		return true,nil
	}
	return false,errors.New("Key is not valid")
}

// value + type 구조체
type ValueType struct {
	Value string
	Type string
}

// value 유효성, 타입 체크
func checkValueType(value string) (ValueType, error)  {
	if value[0] == '"' {
		if value[len(value)-1]=='"' {
			value=value[1:len(value)-1]
			return ValueType{value, "string"}, nil
		} else {
			return ValueType{value,"not valid"}, errors.New("Not Valid String")
		}
	} else if value == "true" || value == "false" {
		return ValueType{value,"bool"}, nil
	} else if _, err:=strconv.Atoi(value); err==nil {
		return ValueType{value,"int"}, nil
	} else {
		return ValueType{value,"not valid"}, errors.New("Not Valid String")
	}
}


func main()  {
	file:=`
	[Server] 
	"port"=true
	"ip"="1.1.2.3"
	"except"=555

	[log]
	"path"="/home/user/log"

	[test]
	"aa"="lll"

	[hi
`
	//m:=make(map[string]map[string]interface{})

	slice:=strings.Split(file, "\n")
	var section []string
	kv:=make(map[string]ValueType)


	for _, str := range slice{
		str=strings.TrimSpace(str)
		if len(str)!=0{
			//  section
			if str[0] == '[' {
				// if section is valid
				if _, err := checkValidSection(str); err == nil {
					section=append(section, str[1:len(str)-1])
				} else {
					fmt.Println(err)
					continue
				}
			}
			if strings.Contains(str, "="){
				// 구분자 index
				deli:=strings.Index(str, "=")
				key:=str[:deli]
				value:=str[deli+1:]
				if _, err:=checkValidKey(key); err == nil {
					if val, err:=checkValueType(value); err==nil {
						kv[key[1:len(key)-1]]=val
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
					continue
				}

			}
		}
	}

	fmt.Println(section)
	fmt.Println(kv)

	fmt.Println(section)
	for k,v := range kv {
		fmt.Println(k,v)
	}
	fmt.Println("---------------------------")
}
