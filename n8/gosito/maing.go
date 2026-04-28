package main

import (
	//"net/http"
	"bytes"
	"encoding/json"
	"face/fbok"
	"fmt"
	"io"
	"net/http"
	"time"
	
)

//https://www.facebook.com/share/p/1DjsYzDYV8/

//ssh -p 443 -R0:127.0.0.1:4008 free.pinggy.io
func main(){
	fmt.Println("Server iniciado.....");

	serv := &http.Server{
		Addr: ":4008",
	}
	
	http.HandleFunc("/webhook", handle)

	serv.ListenAndServe()


	



}



type Response struct{
	Message string `json:"message"`
	PageToken string `json:"access_token"`
}



func handle(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	
	fmt.Println(r.Method)

	if( r.Method == "POST"  ){
		

		e := r.Body
		defer e.Close()
		y, _ := io.ReadAll(e)
		body := fbok.Unm(string(y))

		resp := Response{Message: "Hola", PageToken: PAGE_ACCCESS_TOKEN}

		
		time.Sleep(1 * time.Second)
		if(body.Entry[0].Changes[0].Value.From.Id != "1105873972611788" && body.Entry[0].Changes[0].Field == "feed"){
			fmt.Println("Enviando mensaje....")
			Post(body.Entry[0].Changes[0].Value.CommentId, resp)
			return
		}
		return
		
		
	}
	if( r.Method == "GET"){
		
		fmt.Println(r.RequestURI)
		fmt.Println(r.FormValue("hub_challenge"))

		hub := r.FormValue("hub_challenge")

		fmt.Fprintf(w, hub)
	}

	


}




func Post(commentId string, data Response)bool{

	var url string =fmt.Sprintf("https://graph.facebook.com/v25.0/%s/comments", commentId)

	resp, err := json.Marshal(data)
	fmt.Println("Marshal ejecutando..")
	if(err != nil){
		fmt.Println(err.Error())
		return false
	}

	respp, errr := http.Post(url, "application/json", bytes.NewBuffer(resp))
	fmt.Println("Post ejecutando..")
	if(errr != nil){
		fmt.Println(err.Error())
		return false
	}

	fmt.Println("HECHO")
	fmt.Println(respp.StatusCode)
	return true


}







const (
	TOKEN_USUARIO string = "EAASUYEyTrZCQBRc38nfPRKpN7TN01m52v3bdJo4jgSJvRxQqqRLbAqKNR6iHvHZCfRVYVaIGXx8SF63FeGGc1xpxJCJwfWOZBkmEVbDp0Ee2QEyoswFjeh5hQ8eAse8EIBP94p01CoivbqyKt1st8oak23FAy6LJ3eNu3ckFlgWRpXCX4PZAguY7dyYbeoEO"
	PAGE_ACCCESS_TOKEN string = "EAASUYEyTrZCQBRQVLQUPEuaZAFdRZCs0vCeF5ZBLyZBnKtqLxz6TnZAiTxcBqVx1HohDtZCJmPvskGZCZBfKOyqS9eZC5n5CZCPymzlR3GqLucwZAb91wQTNuTNQAsiwzQK4ZB8p6QyBKuO39WDAH5ZCclMFOWqGa3PfYpsyDlwujc1H8VGcZARRskEUkXsxWWgRJOYwU3uwJHP"	

)



//1105873972611788/subscribed_apps
//EAASUYEyTrZCQBRWjN6jLeOx60m8BOKEtSfPGqPMJnAZBkAcFMkycCPWIJMPRg5IrawkoWQjR3ZB2sKIzws3YYVJaZBK08JigVlZCgIqhVTxoVZCXe2ZBHUp0YKbCRX04ro4lkhHP48qptviTAwrlXU8yFw99IrVCrR5V5au3hudE2ObuoGgidAqRMuvwPQzZCf9Fp5mjFJqC