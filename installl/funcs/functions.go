package funcs

import (
	"bytes"
	"dog/config"
	"dog/mystruct"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

)

//Verifica si un archivo existe o no
func VerifyFile(filePath string)(string, bool){
	_, err := os.Stat(filePath)

	if(err != nil){
		return "", false
	}
	return filePath, true

}



//hace string.split
func SplitString(info string)[]string{
	er := strings.ReplaceAll(info, "\n", " ")
	return strings.Split(er, " ")
	
}

//elimina el ultimo elemento
func DeleteLastElement(list []string)[]string{
	return list[:len(list)-1]
}


//limpia espacios vacios etc de una lista 
func ListCleaner(list[]string)[]string{

	var newList []string

	for i:=0; i<len(list); i++{
		if(list[i] == ""){
			continue
		}else if(list[i]== " "){
			continue
		}else{
			newList = append(newList, list[i])
		}
	}

	return newList

}




func WriteLog(content string)bool{
	logFile :=  config.LOGFILE

	 

	file, err := os.OpenFile(logFile, os.O_CREATE | os.O_RDWR  | os.O_APPEND, 0644)

	if(err != nil){
		fmt.Println(err.Error())
		return false
	}
	defer file.Close()
	file.Write([]byte("\n"+content))
	
	return true

}




func Cmd(command string){
	var outp bytes.Buffer

	cmd := exec.Command(command)
	cmd.Stdout = &outp

	cmd.Run()
	fmt.Println(outp.String())
}

func CreateInfoJson(){
	
	fileInfoJson := config.INFOJSONFILE


	Json := mystruct.Docs{
		User: mystruct.Userr{Active: false, TelegramId: 123456, TelegramToken: "telegram-token"},
		File: mystruct.Files{Paths: []string{"/rutaAbsoluta/hacia/el/archivo.txt"}},
		Dir: mystruct.Directory{Paths: []string{"/rutaAbsoluta/hacia/la/carpeta/"}},
	
	}

	
	fmt.Println("Archivo info.json no encontrado,....generando nuevo archivo info.json")
	time.Sleep(2 * time.Second)
	gson, err := json.MarshalIndent(Json, "", "   ")
	if(err != nil){
		fmt.Println(err.Error())
		
	}	
	file := os.WriteFile(fileInfoJson, []byte(string(gson)), 0666)
	if(file != nil){
		panic("No se pudo crear el archivo: Erro: " + file.Error())
		
	}
	fmt.Println("Archivo creado con exito :" + fileInfoJson )
	
	fmt.Println("Abre el archivo "+ fileInfoJson + "configura tus rutas y vuelve a iniciar el programa")
	time.Sleep(2 * time.Second)
	os.Exit(0)

}

