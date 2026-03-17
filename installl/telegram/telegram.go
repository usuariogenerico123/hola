package telegram

import (
	// "net/http"
	// "os"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	// "dog/dog"
	"dog/config"
	"dog/funcs"
	"dog/mystruct"
)



var URL string = "https://api.telegram.org/"
var active bool
var token string
var chatId int 


func Send(msg string)(string, error){

	if(active){
		Body := mystruct.SendMessage{ChatId: chatId, Text: msg}
		respp, _ := json.Marshal(Body)
		resp , err := http.Post(URL+"bot"+ token + "/" + "sendMessage", "application/json", bytes.NewBuffer(respp))
		if(err != nil){
			fmt.Println(err.Error())
		}
		fmt.Println(resp.StatusCode)	
		return "", nil
	}
	return "", errors.New("Esta opcion no esta activa ")
	
}


func getToken()string{
	jsn,err := os.ReadFile("./info.json")
	if(err != nil){
		panic(" No existe el archivo info.json")
	}
	
	var Data mystruct.Docs
	resp := json.Unmarshal([]byte(string(jsn)), &Data)
	if(resp!=nil){
		fmt.Println(resp.Error())
	}
	return Data.User.TelegramToken

}

func verifyToken(token string)bool{
	resp, err := http.Get(URL+"bot"+token+"/getMe")
	if(err != nil){
		panic("Algo salio mal al verificar el token : " + err.Error())
	}

	// body, err := io.ReadAll(resp.Body)
	// defer resp.Body.Close()

	if(resp.StatusCode != 200){
		fmt.Println("Token incorrecto")
		return false
	}
	
	return true

}

func VerifyUserStatus(){
	var Info mystruct.Docs

	infopath := config.INFOJSONFILE

	body, er := os.ReadFile(infopath)
	if(er != nil){
		funcs.CreateInfoJson()
		fmt.Println()
	}
	
	resp := json.Unmarshal([]byte(string(body)), &Info)
	if(resp != nil){
		fmt.Println(resp)
	}
	if(Info.User.Active == true){

		fmt.Println("Configurando bot de telegram")
		if(Info.User.TelegramToken == " "){
			fmt.Println("Token vacio")
			return
		}

		if(verifyToken(Info.User.TelegramToken)){
			token = Info.User.TelegramToken
			chatId = Info.User.TelegramId
			
			active = true
			fmt.Println("Token verificado con exito")
			return
		}
	}
	//fmt.Println("No se configuro un bot de telegram..")
	active = false
	
}



