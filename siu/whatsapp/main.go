package main

import (
	"fmt"
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
		fmt.Print("Espera.")
		for{

			time.Sleep(500 * time.Millisecond)
			fmt.Print(".")
		}
	}()
	
	go Hola(c)
	
	s :=<- c 
	fmt.Print(s)

}

func Send(c src.Notification){
	c.SendMessage()
}


func Hola(c chan string)string{

	time.Sleep(4 * time.Second)

	c <- "acabe"
	return "asd"
}





