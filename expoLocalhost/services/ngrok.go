package services

import (
	"context"
	"expo/config"
	"fmt"
	"os"
	
	"path/filepath"
	

	"golang.ngrok.com/ngrok/v2"
)


func CreateTunnel(port string, system string){
	
	ctx := context.Background()
	token, erro := LoadToken()
	if(erro != nil){
		fmt.Println(erro)
		return
	}

	agent, err := ngrok.NewAgent(ngrok.WithAuthtoken(token))
	if (err != nil) {
		fmt.Println("Verifica tu token: "+err.Error())
		return
	}
	
	ln, errr := agent.Forward(ctx, ngrok.WithUpstream("http://127.0.0.1:"+port))
	if(errr != nil){
		fmt.Println(errr)
		return
	}
	
	fmt.Println("TUNEL CREADO CON EXITO EN EL PUERTO: "+port)
	fmt.Println("Link: "+"\033[0;33m"+ln.URL().String()+"\033[0m")
	fmt.Println("Presiona CTRL+C para detener")
	<- ln.Done()
	
}

func ExposeServer(port string, path string, system string){
	
	ok := make(chan bool)
	
	go Server(port, path, ok)
	serverOk := <- ok
	if(!serverOk){
		fmt.Println("No se pudo crear el servidor verifica si el puerto no esta ocupado")
		return
	}
	fmt.Println("Servidor iniciado en el puerto:"+port)
	fmt.Println("Servidor: http://127.0.0.1:"+port)

	CreateTunnel(port, system)



}


func LoadToken()(string, error){
	
	path, err := os.UserConfigDir()
	if(err != nil){
		panic("Ocurrio un error inesperado: "+ err.Error())
		
	}

	absolutePath := filepath.Join(path, config.AppName)
	absolutePathToken := filepath.Join(absolutePath, config.NameTokenFile)

	resp , err := os.ReadFile(absolutePathToken)
	if(err != nil){
		panic(err)
	
	}
	return string(resp), nil
}