package IPs

import (
	"fmt"
	"testing"
)



func TestGetIps(t *testing.T){
	resp := GetIps("./IPs/")
	fmt.Println(resp[5])
}
