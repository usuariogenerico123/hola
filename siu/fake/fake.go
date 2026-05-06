package main

import (
	"fake/IPs"
	"fake/menu"
	"fake/style"
	"fmt"
	"os"
	"os/exec"
)




func main(){
	cmd, _ := exec.Command("clear").Output()
	fmt.Println(string(cmd))
	ips := &IPs.IpRanges{IPsPath: "./IPs"}
	ips.Load()
	cdnList := ips.GetListCdn()


	fmt.Println(style.Banner)
	

	option := os.Args[1]
	switch option{
	case "--ip":
		menu.CheckCdnOnly(&cdnList, os.Args[2])
		
	case "--domain":
		menu.CheckAllSubdomain(&cdnList, os.Args[2])
		
	case "--help":
		fmt.Println(menu.Help())

	default:
		fmt.Println("\nInvalid option, please type --help")

	}
	

	
}







