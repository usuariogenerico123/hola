package main

import (
	//"encoding/json"
	"fake/dnsmikis"
	"fake/dnsmikis/cert"
	"fake/dnsmikis/hacktarget"
	"fake/dnsmikis/urlscan"
	"sync/atomic"

	"net"
	"sync"

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
	fmt.Println(style.Banner)
	dominio := input("Escribe el nombre de tu dominio ->: ")
	
	//dominio := os.Args[1]
	if(len(funcs.CheckNs(dominio)) == 0){
		fmt.Printf(style.RED + "Dominio: %s no existe\n"+style.END , dominio  )
		return
	}

	urlCrt := fmt.Sprintf("https://crt.sh/?q=%s&output=json", dominio)
	urlHtarget := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", dominio)
	urlUrlScanio := fmt.Sprintf("https://urlscan.io/api/v1/search/?q=domain:%s", dominio)
	
	
	crtSh := &cert.CrtSh{NameService:"crt.sh", Domain:dominio, Url: urlCrt}
	crt, err := ScanSubdomain(crtSh)
	if(err != nil){
		fmt.Println("iNTENATO")
		intentos := 5
		var ok atomic.Bool
		ok.Store(false)
		respCrt := make(chan domain.SubDomains)
		go func(){fmt.Print("Espera.");for{time.Sleep(2000 * time.Millisecond);if(ok.Load() == true){break};fmt.Print(".")}}()
		go func(){
			for range intentos-1{
				fmt.Print("|")
				time.Sleep(5 * time.Second)
				crt, err = ScanSubdomain(crtSh)
				if(err == nil){

					ok.Store(true)
					respCrt <- crt
					break
				}
			}
			respCrt <- domain.SubDomains{}
			ok.Store(true)
		}()
		crt = <- respCrt
		//fmt.Println(err)
	}
	
	
	hackTarget := &hacktarget.Htarget{NameService:"hackertarget", Domain: dominio, Url: urlHtarget}
	ht, errt := ScanSubdomain(hackTarget)
	if(errt != nil){
		fmt.Println(errt.Error())
		ht = domain.SubDomains{}
	}
	

	scanIo := &urlscan.UrlScan{NameService: "urlscan.io", Domain: dominio, Url: urlUrlScanio}
	sci, errs := ScanSubdomain(scanIo)
	if(errs != nil){
		fmt.Println(errs.Error())
	}




	
	var subdomains []domain.Domain
	
	start := time.Now()
	
	fmt.Println("Iniciando")
	
	

	subdomainsStrings := append(crt.SubDomains, ht.SubDomains...)
	subdomainsStrings = append(subdomainsStrings, sci.SubDomains...)
	listClean := funcs.DeleteRepeat(subdomainsStrings)
	

	
	//Dominio padre
	dominioPadre := domain.Domain{Name: dominio, Ip: func()[]net.IP{ r, _ := funcs.CheckIp(dominio, true);return r}()}
	dominioPadre.CheckNs()
	subdomains = append(subdomains, dominioPadre)
	//Subdominios


	
	subdomains = Init(listClean)


	fmt.Println("\nResultados: ")
	for _, v := range subdomains{
		time.Sleep(100 * time.Millisecond)
		//fmt.Println("-------------------------------------------------------------------------------------------")
		fmt.Println(style.YELLOW, v.Name, style.END ,style.GREEN, v.Ip, style.END ,"\n", v.Cdn)
		fmt.Println("--------------------------------------------------------------------------------------------------")
	}


	end := time.Since(start)
	fmt.Println("Tiempo de ejecucion :", end)
	os.Exit(1)
	//fmt.Println(funcs.CheckIp(dominio, false))
	
	
}




func Init(lista []string)[]domain.Domain{
	subdomains := []domain.Domain{}

	dmain := make(chan *domain.Domain, 10)
	var wg sync.WaitGroup
	limitElements := 100
	numThreads := 7



	if(len(lista) > limitElements){
		fmt.Print("Acelerando....")
		time.Sleep(1 * time.Second)
		chunksSubdomains := funcs.SplitArray(lista, numThreads)

		for _,list := range chunksSubdomains{
			wg.Add(1)
			go func(lista []string){
				defer wg.Done()
				for _, x := range lista{
					
					ip, err := funcs.CheckIp(x, true)
					if(err != nil){
						fmt.Printf("\r%s", style.RED + err.Error() + style.END)
						dmain <- nil
						continue
					}
					domaiin := &domain.Domain{Name: x, Ip: ip }
					domaiin.CheckNs()
					//subdomains = append(subdomains, domaiin)
					dmain <- domaiin
					}
				
				
			}(list)

		}

		go func(){
			wg.Wait()
			close(dmain)
		}()

		for info := range dmain{
			
			if(info != nil){
				subdomains = append(subdomains, *info)
			}
		
			
		}
			
		return subdomains


	}

	for _, x := range lista{
		ip, _ := funcs.CheckIp(x, true)
		domaiin := domain.Domain{Name: x, Ip: ip }
		domaiin.CheckNs()
		subdomains = append(subdomains, domaiin)

	}
	return subdomains

}


func ScanSubdomain(s dnsmikis.Scan)(domain.SubDomains, error){

	resp, err := s.CheckSubdomain()
	if(err != nil){
		
		//fmt.Println(s.ServiceName())
		return resp, err
	}
	return resp, nil

}


func input(message string)string{
	var dmnio string
	fmt.Print(message)
	_, err := fmt.Scanln(&dmnio)
	if(err != nil){
		return ""
	}
	
	return dmnio

}