package funcs

import (
	"reflect"
	"strconv"
)

func VerifyDataType(from any, to any)bool{
	return  reflect.TypeOf(from) == reflect.TypeOf(to)
}

func VerifyNumber(port string)bool{
	_, err := strconv.Atoi(port)
	if(err != nil){
		return false
	}
	return true
}