package main

import (
	"fake/IPs"
	"fake/menu"
	"fake/style"
	"fmt"
	"os"
	//
)




func main(){
	
	fmt.Println(style.Banner)
	
	SaveinfFile := false
	option := os.Args
	if(len(option) < 2){
		fmt.Println()
		fmt.Println("Invalid option, please type --help")
		fmt.Println()

		return
	}
	for _, v := range option{
		if(v == "--save" || v == "-s"){
			SaveinfFile = true
		}
	}

	ips := &IPs.IpRanges{IPsPath: "./IPs"}
	ips.Load()
	cdnList := ips.GetListCdn()

	switch option[1]{
	
	case "--ip":
		
		if(len(option) < 3){
			fmt.Println()
			fmt.Println("Invalid <argument>, please type --help")
			fmt.Println()
			return
		}
		menu.CheckCdnOnly(&cdnList, option[2])	

		
	case "--domain":
		if(len(option) < 3){
			fmt.Println()
			fmt.Println("Invalid <argument>, please type --help")
			fmt.Println()
			return
		}
		menu.CheckAllSubdomain(&cdnList, os.Args[2], SaveinfFile)
		
	case "--help":
		fmt.Println(menu.Help())
		

	default:
		fmt.Println("\nInvalid option, please type --help")
		fmt.Println()
	}


	
	

	
}







