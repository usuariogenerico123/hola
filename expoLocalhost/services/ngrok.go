package services

import (
	"context"
	"fmt"
	"os"
	"golang.ngrok.com/ngrok/v2"
)


func CreateTunnel(port string, system string){
	token := os.Getenv("NGROK_TOKEN")
	ctx := context.Background()
	agent, err := ngrok.NewAgent(ngrok.WithAuthtoken(token))
	if (err != nil) {
		fmt.Println(err)
		return
	}

	ln, errr := agent.Forward(ctx, ngrok.WithUpstream("http://127.0.0.1:"+port))
	if(errr != nil){
		fmt.Println(errr)
		return
	}
	fmt.Println("Link: "+"\033[0;34m"+ln.URL().String()+"\033[0m")
	<- ln.Done()
	
}

func ExposeServer(port string, path string){

}