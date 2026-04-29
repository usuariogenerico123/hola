package src

import "fmt"



type Notification interface{
	SendMessage() bool
}



type WhatsApp struct{
	Message string
}

func (w *WhatsApp) SendMessage()bool{
	fmt.Println("Enviando: ", w.Message)
	return true
}


type Telegram struct{
	Message string 
}
func (t *Telegram) SendMessage()bool{
	fmt.Println("Enviando: ", t.Message)
	return true
}

// func Send(s Notification){
// 	s.SendMessage()
// }



