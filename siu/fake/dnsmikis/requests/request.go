package requests
import(
	"fmt"
	"net/http"
	"time"
	"math/rand"
)

func Get(url string)(*http.Response, error){
	UserAgent := GetRandomUa()
	cli := &http.Client{
		Timeout: 50 * time.Second,
	}


	req, err := http.NewRequest("GET", url, nil)
	if(err != nil){
		//fmt.Println("error Get:", err.Error())
		return nil, fmt.Errorf("Error http NewRquests: %s", err.Error())
	}
	
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "applicattion/json")
	

	resp, err := cli.Do(req)
	if(err != nil){

		return nil, fmt.Errorf("Error get: %s", err.Error())
		
	}
	if(resp.StatusCode != 200){
		return nil , fmt.Errorf("Error status code: %d", resp.StatusCode)
	}


	
	return resp, nil


}





var UserAgents [2]string = [2]string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",

}



func GetRandomUa()string{
	ua := rand.Intn(len(UserAgents)-1)
	return UserAgents[ua]
}

