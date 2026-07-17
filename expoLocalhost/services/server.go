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
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5* time.Second,
		IdleTimeout: 30 * time.Second,
	}
	defer server.Close()
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


