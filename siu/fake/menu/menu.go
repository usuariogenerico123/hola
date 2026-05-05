package menu

import (
	"fake/IPs"
	"fake/funcs"
	"fmt"
	"net"
)


func CheckCdnOnly(ip string, ips *[]IPs.Cdn){
	fmt.Println("Check ip")
	for _,v := range *ips{
		b := funcs.CheckBunnyCDN(net.ParseIP(ip), v.GetIps())
		if(b){
			fmt.Println("bunnycdn")
			return
		}
		

		cdn := funcs.CheckCdn(v.GetName(), net.ParseIP(ip), v.GetIps())
		if(cdn){
			fmt.Println(v.GetName())
			return
		}
	}
	fmt.Println("Not found cdn")
}