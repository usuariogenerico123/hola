package funcs

import (
	//"fake/IPs"
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
	ranges := &IPs.IpRanges{IPsPath: "../IPs/"}
	ranges.Load()
	cdnForTest := "bunnycdn"


	var cdn []string
	for _, v := range ranges.List{
		fmt.Println(v.GetName())
		if(v.GetName() == cdnForTest){
			cdn = append(cdn, v.GetIps()...)
		}
	} 

	ji := CheckCdn(net.IPv4(79,127,213,212), cdn)
	if(ji){
		fmt.Println("is: ", cdnForTest)
		fmt.Println(ji)
	}else{
		fmt.Println("No")
	}
	
	

}

func TestCheckNs(t *testing.T){

	er := CheckNs("enteasdasdal.bo")
	fmt.Println(er)
}

func TestSplitArray(t *testing.T){
	//lista := []int{23, 2,23,4,23,1,43,76,98,34,00,4, 23, 656, 12, 24, 56, 12, 67, 12, 7, 12, 7, 23, 56, 23, 5, 3, 43, 12, 34, 56, 76, 3, 45, 2, 6, 2}
	//s:=SplitArray(lista, 2)

	//fmt.Println(s)
}