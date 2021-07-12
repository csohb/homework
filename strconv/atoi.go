package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main()  {
	/*age:="13"
	fmt.Printf("Datatype of age before conversion: %T\n", age)
	fmt.Println("After Conversion:")
	if sv, err:=strconv.Atoi(age); err==nil{
		fmt.Printf("%T, %v\n", sv,sv)
	}*/
	x:="123"
	fmt.Println("Value of x before conversion:",x)
	fmt.Println("Datatype of x before conversion:", reflect.TypeOf(x))
	intStr, _:=strconv.ParseInt(x,10,64)
	fmt.Println("Value after conversion:",intStr)
	fmt.Println("Datatype after conversion:", reflect.TypeOf(intStr))

}
