package config

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


var htmlExample string = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>
</head>
<style>
  body{
    width:100%;
    
  }
  .test{
   
    width:70%;
    margin: 3rem auto auto auto;
    border-radius:10px;
    box-shadow: 5px 5px 10px ;
    padding:1rem;
    background-color: yellow;
    font-weight:600;

  }
  h2{
    text-align:center;
  }
</style>
<body>
  <div class="test">
    <h2>
      Texto de prueba

    </h2>
    <p>
      Lorem ipsum dolor sit amet consectetur adipisicing elit. Laborum animi deserunt est iusto quis, expedita pariatur, quasi iure et omnis, a quia aspernatur inventore harum. Id exercitationem fugiat nam amet?
    </p>
  </div>
</body>
</html>

`

func TestFilter(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, htmlExample)
	})

	filter := Filter(mux)
	server := httptest.NewServer(filter)
	defer server.Close()

	cases := []struct{
		Name string 
		Path string 
		ResponseExpected int
	}{
		{"Ruta normal", "", 200},
		{"Otr ruta normal", "/path", 200},
		{"Ruta errada ", "/.mikis", 204},
		{"Otra Ruta errada", "/ruta/.file.txt", 204},
	}

	for _, v := range cases{
		t.Run(v.Name, func(t *testing.T) {
			resp, err := http.Get(server.URL+v.Path)
			if(err != nil){
				t.Fatal(err)
			}
			defer resp.Body.Close()
			_, er := io.ReadAll(resp.Body)
			if(er != nil){
				fmt.Println("er", resp.StatusCode)
				t.Fatal(er.Error(), resp.StatusCode)
			}		
			// fmt.Println(resp.StatusCode)
			fmt.Println(resp.Request.URL.Path)
			if(resp.StatusCode != v.ResponseExpected){
				fmt.Println("Se esperaba: ", v.ResponseExpected, "Se obtuvo: ", resp.StatusCode)
				t.Fatal(resp.StatusCode)
			}		
		})
	}
}