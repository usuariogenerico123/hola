package services

import (
	"expo/config"
	"fmt"
	"net/http"

	"time"
)


func Server(port string, path string, serv chan *http.Server ){
	
	
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(path)))

	filter := config.Filter(mux)

	server := &http.Server{
		Addr: ":"+port,
		Handler: filter,
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


