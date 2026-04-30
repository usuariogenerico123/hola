package requests
import(
	"fmt"
	"fake/dnsmikis/agents"
	"net/http"
	"time"
)

func Get(url string)(*http.Response, error){
	UserAgent := agents.GetRandomUa()
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



	// resp, err := http.Get(url)
	// if(err != nil){
	// 	fmt.Println("Error get")
	// 	fmt.Println(err.Error())
		
	// }
	
	return resp, nil


}
