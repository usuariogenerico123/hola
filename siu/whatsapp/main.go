package main

import (
	
	"wp/src"
)



func main(){
	wp := &src.WhatsApp{}
	wp.Message = "whatsapp"

	te := &src.Telegram{}
	te.Message = "telegram"

	Send(wp)
	Send(te)

	

}

func Send(c src.Notification){
	c.SendMessage()
}






