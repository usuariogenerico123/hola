package agents

import (
	
	"math/rand"
)




var UserAgents [2]string = [2]string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",

}



func GetRandomUa()string{
	ua := rand.Intn(len(UserAgents)-1)
	return UserAgents[ua]
}