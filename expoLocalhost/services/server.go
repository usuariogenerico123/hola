package services

import (
	"fmt"
	"net/http"

	"time"
)


func Server(port string, path string, serv chan *http.Server ){
	
	
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(path)))
	server := &http.Server{
		Addr: ":"+port,
		Handler: mux,
		ReadTimeout: 3 * time.Second,
	}
	
	go func(){
		serverErro := server.ListenAndServe()
		if( serverErro!= nil ){
			//panic(serverErro)
			fmt.Println(serverErro)
			return
		}
	}()
	time.Sleep(10 * time.Second)
	serv <- server
		
	
}


