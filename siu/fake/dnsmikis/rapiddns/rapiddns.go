package rapiddns

import (
	"fake/dnsmikis/requests"
	"fake/domain"
	"fmt"
	"io"
	"regexp"
	"strings"
)



type RapidDdns struct{
	NameService string 
	Domain string 
	Url string 
}


func (r *RapidDdns) CheckSubdomain()(domain.SubDomains, error){
	var result domain.SubDomains
	resp, err := requests.Get(r.Url)
	if(err != nil){
		return result, err
	}

	body, err:= io.ReadAll(resp.Body)
	if(err != nil){
		// fmt.Println("Error: ")
		// fmt.Println(err.Error())
		return result, err
	}
	
	//fmt.Println(string(body))
	list := []string{}

 	data :=  regexp.MustCompile(`<td>([a-zA-Z0-9._-]+\.[a-zA-Z]{2,})</td>`)
	jijo := data.FindAllSubmatch([]byte(body), -1)
	for _, v := range jijo{
		//dominio := strings.ReplaceAll(string(v[0]), "</td>", "")
		fmt.Println(strings.Replace())
		list = append(list, strings.Replace(string(v[0]), "</td>", "", 6))
	}
	//fmt.Println(list)
	return result, nil
	
}