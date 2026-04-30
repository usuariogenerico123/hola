package main

import (
	//"encoding/json"
	"fake/dnsmikis"
	"fake/dnsmikis/cert"
	"fake/dnsmikis/hacktarget"

	//"fake/dnsmikis/hacktarget"
	"fake/domain"
	"fake/funcs"
	"fake/style"
	"os"

	//"strings"
	"time"

	//"fake/funcs"
	"fmt"
)





func main(){

	
	
	dominio := os.Args[1]
	urlTrg := fmt.Sprintf("https://crt.sh/?q=%s&output=json", dominio)
	urlHtarget := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", dominio)
	
	
	
	crtSh := &cert.CrtSh{NameService:"crt.sh", Domain:dominio, Url: urlTrg}
	crt, err := ScanSubdomain(crtSh)
	if(err != nil){
		fmt.Println(err)
		//return
	}

	
	hackTarget := &hacktarget.Htarget{NameService:"hackertarget", Domain: dominio, Url: urlHtarget}
	ht, errt := ScanSubdomain(hackTarget)
	if(errt != nil){
		fmt.Println(errt.Error())
		//return
	}


	

	// subDomainHtarge, err := hacktarget.CheckSubdomain(urlHtarget)
	// if(err != nil){
	// 	fmt.Println(err)
	// 	return
	// }

	
	var subdomains []domain.Domain
	
	start := time.Now()
	fmt.Println(style.Banner)
	fmt.Println("Iniciando")
	
	

	subdomainsStrings := append(crt.SubDomains, ht.SubDomains...)
	listClean := funcs.DeleteRepeat(subdomainsStrings)
	

	
	//Dominio padre
	dominioPadre := domain.Domain{Name: dominio, Ip: funcs.CheckIp(dominio, true)}
	dominioPadre.CheckNs()
	subdomains = append(subdomains, dominioPadre)
	//Subdominios


	
	subdomains = Init(listClean)

	// for _, x := range listClean{
	// 	//fmt.Print(x)
	// 	time.Sleep(50 * time.Millisecond)

	// 	domaiin := domain.Domain{Name: x, Ip: funcs.CheckIp(x, true) }
	// 	domaiin.CheckNs()
	// 	subdomains = append(subdomains, domaiin)

	// }

	fmt.Println("\nMostrando resultados ")
	for _, v := range subdomains{
		time.Sleep(100 * time.Millisecond)
		//fmt.Println("-------------------------------------------------------------------------------------------")
		fmt.Println(style.YELLOW, v.Name, style.END ,style.GREEN, v.Ip, style.END ,"\n", v.Cdn)
		fmt.Println("--------------------------------------------------------------------------------------------------")
	}


	end := time.Since(start)
	fmt.Println("Terminado en :", end)
	//fmt.Println(funcs.CheckIp(dominio, false))

	
}




func Init(lista []string)[]domain.Domain{
	subdomains := []domain.Domain{}

	dmain := make(chan domain.Domain, 10)
	limitElements := 30
	numThreads := 5



	if(len(lista) > limitElements){
		fmt.Print("Acelerando....")
		time.Sleep(1 * time.Second)
		chunksSubdomains := funcs.SplitArray(lista, numThreads)

		for _,list := range chunksSubdomains{
			go func(){
				
				for _, x := range list{
					//fmt.Print(x)
				
					time.Sleep(50 * time.Millisecond)	
					domaiin := domain.Domain{Name: x, Ip: funcs.CheckIp(x, true) }
					domaiin.CheckNs()
					//subdomains = append(subdomains, domaiin)
					dmain <- domaiin
					}
				//fmt.Printf(".")
			}()

		}

		for range(len(lista)){
			info := <- dmain
			subdomains = append(subdomains, info)
		}
			
		return subdomains


	}

	for _, x := range lista{
		//fmt.Print(x)
		time.Sleep(50 * time.Millisecond)
		
		domaiin := domain.Domain{Name: x, Ip: funcs.CheckIp(x, true) }
		domaiin.CheckNs()
		subdomains = append(subdomains, domaiin)

	}
	return subdomains

}


func ScanSubdomain(s dnsmikis.Scan)(domain.SubDomains, error){

	resp, err := s.CheckSubdomain()
	if(err != nil){
		
		fmt.Println(err.Error(), s.ServiceName())
		return resp, err
	}
	return resp, nil

}