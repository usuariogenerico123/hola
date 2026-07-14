package services

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func Server(port string, path string, ok chan bool){
	
	cancel := make(chan os.Signal, 1)
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(path)))
	server := &http.Server{
		Addr: ":"+port,
		Handler: mux,
	}
	var serverErro error
	go func(){
		serverErro = server.ListenAndServe()
		if( serverErro!= nil ){
			fmt.Println(serverErro)
			
			return 
		}
	}()
	time.Sleep(5 * time.Second)
	if(serverErro != nil){
		ok <- false
	}
	ok <-true
	
	signal.Notify(cancel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	ctx, canc := context.WithTimeout(context.Background(), 7 * time.Second)
	defer canc()
	<- cancel
	fmt.Println("Servidor cerrado")
	server.Shutdown(ctx)

}


