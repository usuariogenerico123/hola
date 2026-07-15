package menu

import (
	"context"
	"expo/config"
	"expo/funcs"
	"expo/services"
	"fmt"
	"runtime"
)


func Menu(){
	
	var option string
	var port string
	var path string
	systemName := runtime.GOOS

	fmt.Println("Os: "+systemName)
	Loop:
	for {
		fmt.Println(Options())
		fmt.Print("Elige una opcion: ")
		_, err := fmt.Scanln(&option)
		if(err != nil){
			//fmt.Println("erro: scan ",err)
			
			continue
		}

		switch option {
		case "1":
			fmt.Print("Escribe el numero de puerto ejem(3001): ")
			_, err := fmt.Scanln(&port)
			if(err != nil){
				fmt.Println("Algo salio mal")
				continue
			}
			if(!funcs.VerifyNumber(port)){
				fmt.Println("\nPuerto invalido")
				continue
			}
			fmt.Println("Iniciando..")
			services.CreateTunnel(port, systemName, context.Background())

		case "2":
			fmt.Print("Escribe la ruta a tu proyecto (/index.html): ")
			_, err := fmt.Scanln(&path)
			if(err != nil){
				fmt.Println("Algo salio mal")
				continue
			}
			fmt.Print("Escribe el numero de puerto que quieres abrir ejem(3007): ")
			_, err = fmt.Scanln(&port)
			if(err != nil){fmt.Println("Algo salio mal")
				continue
			}
			if(!funcs.VerifyNumber(port)){
				fmt.Println("\nPuerto invalido")
				continue
			}
			services.ExposeServer(port, path, systemName)
			fmt.Println("Exit")
			break Loop
		
		case "3":
			Config()
		
		case "0":
			fmt.Println("Exit")
			break Loop
		default:
			fmt.Println("Opcion incorrecta")
		}
	
		
	}
	

}




func Options()string{
	options := fmt.Sprintf(`
| %sEXPONE TU LOCALHOST A INTERNET %s|

Opciones:
1: Crear url de acceso  (tu servidor ya debe estar encendido)
2: Crear servidor + url de acceso (crea un servidor para tu proyecto)
3: Configurar token de ngrok
0: Salir
	`, Green, End)

	return options
}


var (
	End  = "\033[0m"
	Red    = "\033[0;31m"
	Green  = "\033[0;32m"
	Yellow = "\033[0;33m"
	Blue   = "\033[0;34m"
	Purple = "\033[0;35m"
	Cyan   = "\033[0;36m"
	White  = "\033[0;37m"
	Black  = "\033[0;30m"
)


func Config(){
	var token string
	fmt.Println("Configurar token")
	fmt.Print("Inserta tu token ngrok: ")
	_, er := fmt.Scanln(&token)
	if(er != nil){
		fmt.Println(Red+"Datos incorrectos"+End)
		return 
	}
	fmt.Println("Configurando token...")
	if(!funcs.CreateTokenFile(config.NameTokenFile, token)){
		return
	}
	fmt.Println(Green+"Token configurado con exito."+End)
}