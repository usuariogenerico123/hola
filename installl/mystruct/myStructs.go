package mystruct





//---------------------------------------------
type Files struct{
	Paths []string `json:"paths"`
}
type Directory struct{
	Paths []string `json:"paths"`
}

type Userr struct {
	Active bool `json:"active"`
	TelegramId int `json:"telegramId"`
	TelegramToken string `json:"telegramToken"`
}

type Docs struct{
	User Userr `json:"user"`
	File Files `json:"file"`
	Dir  Directory `json:"dir"`


}
//------------------------------------


type SendMessage struct{
	ChatId int `json:"chat_id"`
	Text string `json:"text"`
}


type Data struct{
	Dirs []string
	Files []string 
	TotalDirs int 
	TotalFiles int
}





