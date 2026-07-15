package main

import (
	"expo/config"
	"expo/funcs"
	"expo/menu"
)


func main(){
	if(!funcs.VerifyConfigPath(config.AppName)){
		if(!funcs.SetupAppConfigPath(config.AppName)){
			return
		}
	}
	if(!funcs.VerifyNgrokToken(config.NameTokenFile)){
		menu.Config()
	}

	menu.Menu()
}