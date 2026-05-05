package main

import (
	"fmt"
	"net/http"
	"time"
	"wp/src"
)



func main(){

	nombre := "Pepe"
	var saludo string = func(user string)string{return "hola "+user}(nombre)

	fmt.Println(saludo)
	c := make(chan string)

	wp := &src.WhatsApp{}
	wp.Message = "whatsapp"

	te := &src.Telegram{}
	te.Message = "telegram"

	Send(wp)
	Send(te)

	go func(){
		fmt.Printf("Tiempo de ejecucion S: %d", 1)
		cont := 0
		for{

			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("\rTiempo de ejecucion S: %d", cont)
			cont ++
		}
	}()
	
	go Hola(c)
	
	s :=<- c 
	fmt.Print(s)



	serv := &http.Server{
		Addr: ":3006",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	fmt.Println("Servidor iniciado")
	serv.ListenAndServe()
}

func Send(c src.Notification){
	c.SendMessage()
}


func Hola(c chan string)string{

	time.Sleep(4 * time.Second)

	c <- "acabe"
	return "asd"
}





