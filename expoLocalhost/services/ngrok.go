package services

import (
	"context"
	"expo/config"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"path/filepath"

	"golang.ngrok.com/ngrok/v2"
)


func CreateTunnel(port string, system string, ctx context.Context){
	
	
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
	
	fmt.Println("\033[0;32mTUNEL CREADO CON EXITO EN EL PUERTO: "+port+"\033[0m")
	fmt.Println("Link: "+"\033[0;33m"+ln.URL().String()+"\033[0m")
	fmt.Println("Presiona CTRL+C para detener")
	<- ln.Done()
	
}

func ExposeServer(port string, path string, system string){
	ctx:= context.Background()
	
	
	stop := make(chan os.Signal, 1)

	infoServer := make(chan *http.Server)
	go Server(port, path, infoServer)
	server := <- infoServer
	
	fmt.Println("Servidor iniciado en el puerto:"+port)
	fmt.Println("Servidor:\033[0;32m http://127.0.0.1"+server.Addr+"\033[0m")



	CreateTunnel(port, system, ctx)


	//close server 
	signal.Notify(stop, syscall.SIGINT, os.Interrupt)
	<- stop 
	fmt.Println("Server stoped")
	server.Shutdown(ctx)


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