package funcs

import (
	"fake/IPs"
	"fmt"
	"net"
	"testing"
)




func TestCheckIp(t *testing.T){
	url := "money.tigo.com.bo"
	r, _ := CheckIp(url, true)
	fmt.Println("IP: ",r)
	fmt.Println("NS: ",CheckNs(url))

}


func TestCheckCdn(t *testing.T){
	ji := CheckCdn(net.IPv4(162,159,138,52), IPs.CLOUDFLARE)
	fmt.Println("Is cloudflare")
	fmt.Println(ji)

}



func TestSplitArray(t *testing.T){
	//lista := []int{23, 2,23,4,23,1,43,76,98,34,00,4, 23, 656, 12, 24, 56, 12, 67, 12, 7, 12, 7, 23, 56, 23, 5, 3, 43, 12, 34, 56, 76, 3, 45, 2, 6, 2}
	//s:=SplitArray(lista, 2)

	//fmt.Println(s)
}