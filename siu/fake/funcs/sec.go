package funcs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	
)



var MyId []string

func VerifyId(name string)string{
	resp , err:= http.Get("https://raw.githubusercontent.com/xosuu/gatuso/refs/heads/main/credits")
	if(err != nil){
		fmt.Println(err)
	}
	body := resp.Body
	defer body.Close()
	s, _ := io.ReadAll(body)
	
	return string(s)

}


func GenerateId()string{
	e,r := os.Hostname()
	if(r != nil){
		fmt.Println(r)
	}

	resp := e + os.Args[1]
	MyId =append(MyId, resp)
	return resp
}


