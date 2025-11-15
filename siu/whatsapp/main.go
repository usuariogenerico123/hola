package main

import (
	
	"wp/banners"
	"fmt"
	"os/exec"
)



func main(){
	command("clear")

	banner := banners.MainBanner()

	fmt.Println(banner)

}




func command(comando string){
	com := exec.Command(comando)
	c, err := com.Output()
	if(err != nil){
		fmt.Println(string(err.Error()))
	}
	fmt.Println(string(c))
}