package funcs

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"net"

	"time"
)

//---------------
func DeleteRepeat(list []string)[]string{

	// fmt.Println("Limpiando lista?---")
	// fmt.Println(list)
	var newList []string

	for i:=0; i<len(list); i++{
		num := list[i]
		for x:=0; x<len(list); x++{
			if num != list[x] && elementInList(newList, num) == false{
				newList = append(newList, num)
			}
		}
	}

	return  newList

}


func elementInList(list []string, element string)bool{

	for i:=0; i<len(list); i++{
		if(element == list[i]){
			return true
		}
	}
	return false

}

//------------------


func CheckIp(url string, onlyIpv4 bool)([]net.IP, error){
	

	resp, err := net.LookupIP(url)
	if(err != nil){
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		// fmt.Printf("\r %s", err.Error())
		return []net.IP{}, err
	}
	if(onlyIpv4){
		
		var ipv4s []net.IP 

		//fmt.Println(len(resp))
		for i:=0; i<len(resp);i++ {
			if(len(resp[i]) != net.IPv6len){
				ipv4s = append(ipv4s, resp[i])
			}
		}
		return ipv4s, nil
	}

	return resp, nil
}




func CheckNs(url string)[]string{
	resp, err := net.LookupNS(url)
	if(err != nil){
		//fmt.Println(err.Error())
		return []string{}
	}
	list := []string{}
	for _, v := range resp{
		list = append(list, v.Host)
	}
	return list
}



func CheckCdn(ip net.IP, rangeIps[]string)bool{

	for _, v := range(rangeIps){

		
		_, ipnet, err := net.ParseCIDR(v)
		if(err != nil){
			fmt.Println("Error checkcnd: "+ v + err.Error())
			return false
		}

		ipp := net.ParseIP(ip.String())
		if(ipnet.Contains(ipp)){
			return true
		}


	}
	return false
	
}


//Dividimos un array en pequenos chunks
//Lo convertimos en un array bidimensional [[1,2,3], [4,5,6]]

func SplitArray(list []string, numSplit int)[][]string {
	newArray := [][]string{}


	nElementos := len(list) / numSplit
	resto := len(list) % numSplit
	
	
	//fmt.Println(list)
	indice := 0 
	final := nElementos
	for i :=1; i<=numSplit; i++{
		//time.Sleep(1 * time.Second)
		fin := final * i
		chunk := list[indice : fin]
	
		//fmt.Println("Inicio: ", i, fin)
		
		if(resto != 0 && i == numSplit){ //Si hay restantes (resto) y es el ultimo chunk se anade los elementos a este ultimo 
			for _, restante := range list[fin:]{
				chunk = append(chunk, restante)
			// 	fmt.Println(restante)
			 }
			// fmt.Println(list[fin:])
			// fmt.Println("Si hay restante")
			// fmt.Println(chunk)
		}

		newArray = append(newArray, chunk)
		//fmt.Println(chunk)

		indice = fin
	}
	return newArray

}


//Parsea el json string hacia un struct 
//[T any] es el tipo de dato T que manejara y devolvera
func Parser[T any](content string)(T, error){
	var data T
	resp := json.Unmarshal([]byte(content), &data)
	if(resp!= nil){
		fmt.Println("Error marshal: ", resp.Error())
		return data, resp
		
	}
	return data, nil
}