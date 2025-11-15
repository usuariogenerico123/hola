package banners

import (
	"os"
)





func MainBanner()(string){
	red := "\033[1;31m"

	fin:="\033[0m"


	whatsapp, err := os.ReadFile("./banners/whatsapp.txt")
	if(err != nil){
		return string(err.Error())
	}
	return red + string(whatsapp) + fin + "\n"
	
}


