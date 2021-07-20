package config

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type UnixConfigHandler interface {
	GetConfigInteger(section, key string) (int, error)
	GetConfigString(section, key string) (string, error)
	GetConfigBoolean(section, key string) (bool, error)
	LoadConfig(file string) error
}

type unixConfig struct {
	File   string
	config map[string]map[string]interface{}
}

const (
	T = true
	F = false
)

// 여기에서 section, key, value 분류?
func (u *unixConfig) LoadConfig(File string) error {
	// 1. file check
	if len(File) != 0 {
		// 2. file read
		slice := strings.Split(File, "\n")
		var section string
		m := make(map[string]map[string]interface{})
		kv := make(map[string]interface{})
		// 3. data parsing
		for _, str := range slice {
			str = strings.TrimSpace(str)
			if len(str) != 0 {
				//  section
				if str[0] == '[' {
					// if section is valid
					if _, err := checkValidSection(str); err == nil {
						if section != "" {
							m[section] = kv
							//fmt.Println(m)
						}
						section = str[1 : len(str)-1]
						kv = make(map[string]interface{})
					} else {
						return err
					}
				} else if strings.Contains(str, "=") {
					// 구분자 index
					deli := strings.Index(str, "=")
					key := str[:deli]
					value := str[deli+1:]
					if _, err := checkValidKey(key); err == nil {
						if val, err := checkValueType(value); err == nil {
							kv[key[1:len(key)-1]] = val
						} else {
							return err
						}
					} else {
						return err
						continue
					}
				}
			}
		}
		//
		if len(section) > 0 {
			if len(kv) > 0 {
				m[section] = kv
			}
		}
		u.config = m
		fmt.Println(u.config)
	} else {
		return fmt.Errorf("config file is empty")
	}
	return nil
}

func (u *unixConfig) GetConfigInteger(section, key string) (int, error) {
	// if section 존재, key에 해당하는 value 존재할때 -> value를 int로 형 변환하여 return value, nil
	// error 있으면 _, error
	// 1. 사용자가 원한 섹션을 가지고 있어?
	// 2. 사용자가 원하는 섹션에 키를 가지고 있어?
	// 3. 사용자가 원하는 값의 타입이 맞아?
	var iVal int
	if kv, has := u.config[section]; has == true {
		if val, has := kv[key]; has == true {
			switch val.(type) {
			case int:
				iVal = val.(int)
				return iVal, nil
			default:
				return iVal, errors.New("Value is not int")
			}
		} else {
			return iVal, errors.New("This Section has not the key")
		}
	}
	return iVal, errors.New("Section does not exist")
}

func (u *unixConfig) GetConfigString(section, key string) (string, error) {
	var sVal string
	if kv, has := u.config[section]; has == true {
		if val, has := kv[key]; has == true {
			switch val.(type) {
			case string:
				sVal = val.(string)
				return sVal, nil
			default:
				return sVal, errors.New("Value is not string")
			}
		} else {
			return sVal, errors.New("This Section has not the key")
		}
	}
	return sVal, errors.New("Section does not exist")
}

func (u *unixConfig) GetConfigBoolean(section, key string) (bool, error) {
	var bVal bool
	if kv, has := u.config[section]; has == true {
		if val, has := kv[key]; has == true {
			switch val.(type) {
			case bool:
				bVal = val.(bool)
				return bVal, nil
			default:
				return bVal, errors.New("Value is not Boolean")
			}
		} else {
			return bVal, errors.New("This Section has not the key")
		}
	}
	return bVal, errors.New("Section does not exist")
}

// 실행 함수
func NewUnixConfig(filename string) UnixConfigHandler {
	// filename을 받아서 새로운 unixConfig 구조체 인스턴스를 생성하고, 그 생성된 인스턴스의 주소값을 return
	return &unixConfig{
		File: filename,
	}
}

// section 유효성 체크
func checkValidSection(section string) (bool, error) {
	if section[len(section)-1] == ']' {
		return true, nil
	}
	return false, errors.New("Section is not valid")
}

// key 유효성 체크
func checkValidKey(key string) (bool, error) {
	bKey := []byte(key)
	if bKey[0] == '"' && bKey[len(bKey)-1] == '"' {
		return true, nil
	}
	return false, errors.New("Key is not valid")
}

// value 유효성, 타입 체크
func checkValueType(value string) (interface{}, error) {
	if value[0] == '"' {
		if value[len(value)-1] == '"' {
			value = value[1 : len(value)-1]
			return value, nil
		} else {
			return value, errors.New("Not Valid String")
		}
	} else if value == "true" || value == "false" {
		switch value {
		case "true":
			return T, nil
		case "false":
			return F, nil
		default:
			return value, errors.New("Not Valid Bool")
		}
	} else if val, err := strconv.Atoi(value); err == nil {
		return val, nil
	} else {
		return value, errors.New("Not Valid Value")
	}
}
