package funcs

import (
	"errors"
	"expo/config"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

func VerifyDataType(from any, to any)bool{
	return  reflect.TypeOf(from) == reflect.TypeOf(to)
}

func VerifyNumber(port string)bool{
	InvalidPorts := []string{"22", "443", "25", "21", "53", "445"}
	_, err := strconv.Atoi(port)
	if(err != nil){
		return false
	}
	for _, v := range(InvalidPorts){
		if(v == port){
			return false
		}
	}
	return true
}

func VerifyNgrokToken(fileName string)bool{
	path := GetPathToken()
	pathToken := filepath.Join(path, fileName)
	_ , err := os.Stat(pathToken)
	if(err != nil){
		return false
	}
	return true

}
func VerifyConfigPath(appName string)bool{
	path, err := os.UserConfigDir()
	if(err != nil){
		fmt.Println("Ocurrio un error inesperado: ", err.Error())
		return false
	}
	pathConfig := filepath.Join(path, appName)
	_, er := os.Stat(pathConfig)
	if(er != nil){
		fmt.Println("No existe el archivo de configuracion")
		return false
	}
	return true

}


func CreateTokenFile(file string, token string)bool{
	path := GetPathToken()
	pathFile := filepath.Join(path, file)
	err := os.WriteFile(pathFile, []byte(token), 0644)
	if(err != nil){
		fmt.Println("error al escribir token, erro: "+err.Error())
		return false
	}
	return true

}

func SetupAppConfigPath(appName string)bool{
	configPath,_ := os.UserConfigDir()
	absolutePath := filepath.Join(configPath, appName)
	err := os.MkdirAll(absolutePath, 0777)
	if(err != nil){
		fmt.Println("No se pudo crear la ruta de configuracion: ", err.Error())
		return false
	}
	return true
}


func GetPathToken()string{
	resp, _ := os.UserConfigDir()
	path := filepath.Join(resp, config.AppName)
	return path

}

func VerifyIndexPage(path string)bool{
	_, resp := os.Stat(path+"/index.html")
	if(resp != nil){
		return false
	}
	return true
}

func VerifyYesNo(data string)(bool, error){
	dat := strings.ToLower(data)
	words := []string{"yes", "y", "no", "n"}
	inList:=false
	for _,v :=range words{
		
		if(v == dat){
			inList=true
			break
		}
	}
	if(!inList){
		return false, errors.New("Error N|Y")
	}

	if(dat == "y" || dat == "yes" ){
		return true, nil
	}
	
	return false, nil
}