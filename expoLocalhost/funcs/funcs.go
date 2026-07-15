package funcs

import (
	"expo/config"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

func VerifyDataType(from any, to any)bool{
	return  reflect.TypeOf(from) == reflect.TypeOf(to)
}

func VerifyNumber(port string)bool{
	_, err := strconv.Atoi(port)
	if(err != nil){
		return false
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

