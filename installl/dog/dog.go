package dog

import (
	"bytes"
	"dog/config"
	"dog/funcs"
	"dog/mystruct"
	"dog/telegram"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)





func Analize(file string){
	
	fil, _ := os.Stat(file)
	fileName := file
	last := fil.ModTime()
	

	for {
		time.Sleep(5 * time.Second)
		fil , err := os.Stat(fileName)
		if(err != nil){
			
			msg:= fmt.Sprintf("Archivo > %s eliminado o removido, %s", fileName, time.Now().Format("2006-01-02 15:04:05"))
			telegram.Send(msg)
			//log.Println("No pude encontrar el archivo", err)
			
			fmt.Println(msg)
			funcs.WriteLog(msg)
			break
		}else{
			if (last != fil.ModTime()){
				//log.Printf("Se modifico el archivo..%s", string(fileName))
				msg := fmt.Sprintf("%s Se modifico >> %s", string(fil.ModTime().Format("2006-01-02 15:04:05")),string(fileName))
				fmt.Println(msg)
				telegram.Send(msg)
				last = fil.ModTime()
				funcs.WriteLog(msg)
			}
		}
	}
}



//Leemos los archivos que estan dentro de directorios
func GetFilesFromDir(PathDir []string)(mystruct.Data, error){
	data := PathDir
	
	var listFiles []string
	var out bytes.Buffer

	var finalList []string
	//
	// var totalDirs int
	
	var InfoFiles mystruct.Data = mystruct.Data{} 

	if(len(data) == 0){
		return InfoFiles, errors.New("No hay directorios")
	}

	for i:=0; i<len(data); i++{
		time.Sleep(200 * time.Millisecond)
		info, er := os.Stat(data[i])
		fmt.Println("Carpetas encontradas: ", len(data))
		if(er != nil){
			continue
		}else if( info.IsDir()){ //verifica si es una carpeta real
			//totalDirs += 1

			
			out.Reset() //limpiamos 
			
			InfoFiles.Dirs = append(InfoFiles.Dirs, data[i])
			
			fmt.Println("Verificando carpeta..", data[i])
			time.Sleep(3*time.Second)
			cmd := exec.Command("ls", data[i])
			cmd.Stdout = &out
			cmd.Run()
			//fmt.Println(out.String())


			//convertimos en lista y cada ruta la Anadimos a la lista principal
			list := funcs.SplitString(out.String())
			list = funcs.DeleteLastElement(list)
			//fmt.Println(len(list))
			for fi:=0; fi<len(list); fi++{
				//fmt.Println(list[fi])
				ruta := data[i] + list[fi] //ruta
				listFiles = append(listFiles, ruta)
			}
			

		}
		
	}


	//limpiar lista
	
	fmt.Println("limpiando lista")
	//fmt.Println(len(listFiles))
	//fmt.Println(listFiles)
	for i:=0; i<len(listFiles); i++{
		time.Sleep(100*time.Millisecond)
		resp, err := funcs.VerifyFile(listFiles[i]) //verificamos los archivos 
		if(err == false){
			continue
		}
		finalList = append(finalList, resp) 

		InfoFiles.Files = append(InfoFiles.Files, resp)
	}

	InfoFiles.TotalDirs = len(InfoFiles.Dirs)
	InfoFiles.TotalFiles = len(InfoFiles.Files)

	//devolvemos la cantidad de carpetas validas, lista de todos los archivos, nil
	return InfoFiles, nil
	

}


//Leemos los archivos sueltos sin carpetas
func GetFiles(paths[]string)([]string, error){
	var files []string
	data := paths

	if(len(data) == 0){
		return []string{}, errors.New("Archivos vacios o no validos: files{paths:[]}")
	}
	
	for i:=0; i<len(data); i++{
		time.Sleep(50 * time.Millisecond)
		f, err := funcs.VerifyFile(data[i])

		if(err == false){
			continue
		}
		files = append(files, f)
		
	}
	if(len(files)==0){
		return []string{}, errors.New("Se encontraron archivos pero no son validos: files{paths:['no existe']}") 
	}
	return files, nil
}





//Lee el info.json y lo parsea en el struct Docs para ser procesado

func ReadDocs()mystruct.Docs{

	pathinfojson := config.INFOJSONFILE

	files, err := os.ReadFile(pathinfojson)
	if(err != nil){
		panic("No existe el archivo "+pathinfojson + "\n"+ err.Error())
	}


	var inf mystruct.Docs
	er := json.Unmarshal([]byte(string(files)), &inf)
	
	if(er != nil){
		panic("Las rutas de los archivos deben estar en formato string \n"+er.Error())
		
	}

	return inf

}









