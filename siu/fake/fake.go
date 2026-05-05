package main

import (
	//"encoding/json"
	"fake/IPs"
	"fake/dnsmikis"
	"fake/dnsmikis/cert"
	"fake/dnsmikis/hacktarget"
	"fake/dnsmikis/rapiddns"
	"fake/dnsmikis/urlscan"

	//"sync/atomic"
	"fake/domain"
	"fake/funcs"
	"fake/style"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)




func main(){
	fmt.Println(style.Banner)
	//dominio := input("Escribe el nombre de tu dominio ->: ")
	dominio := os.Args[1]
	if(len(funcs.CheckNs(dominio)) == 0){
		fmt.Printf(style.RED + "Dominio: %s no existe\n"+style.END , dominio  )
		return
	}

	ips := &IPs.IpRanges{IPsPath: "./IPs"}
	ips.Load()
	cdnList := ips.GetListCdn()
	
	



	urlCrt := fmt.Sprintf("https://crt.sh/?q=%s&output=json", dominio)
	urlHtarget := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", dominio)
	urlUrlScanio := fmt.Sprintf("https://urlscan.io/api/v1/search/?q=domain:%s", dominio)
	urlRapidDns := fmt.Sprintf("https://rapiddns.io/subdomain/%s?full=1", dominio)

	
	
	crtSh := &cert.CrtSh{NameService:"crt.sh", Domain:dominio, Url: urlCrt}
	crt, err := ScanSubdomain(crtSh)
	if(err != nil){
		//fmt.Println("Intentando: 1")
		
		intentos := 10
		//var ok atomic.Bool
		//ok.Store(false)
		respCrt := make(chan domain.SubDomains)
		//go func(){fmt.Print("Espera.");for{time.Sleep(2000 * time.Millisecond);if(ok.Load() == true){break};fmt.Print(".")}}()
		go func(){
			ok := false
			for v := range intentos-1{
				fmt.Printf("\rTrying: %s%d%s", style.Randcolor() ,v ,style.END)
				time.Sleep(1 * time.Second)
				crt, err = ScanSubdomain(crtSh)
				if(len(crt.SubDomains) > 0){
					//fmt.Print(crt.SubDomains)
					//ok.Store(true)
					respCrt <- crt
					fmt.Printf("\r%s", "---- ok ----")
					ok = true
					
					break
				}
			}
			
			if(ok == false){
				fmt.Printf("\r%s", err.Error())
				respCrt <- domain.SubDomains{}

			}
			//ok.Store(true)
		}()
		crt = <- respCrt
		
	}
	
	
	hackTarget := &hacktarget.Htarget{NameService:"hackertarget", Domain: dominio, Url: urlHtarget}
	ht, errt := ScanSubdomain(hackTarget)
	if(errt != nil){
		fmt.Println("\nhtarget> ",errt.Error())
		ht = domain.SubDomains{}
	}
	

	scanIo := &urlscan.UrlScan{NameService: "urlscan.io", Domain: dominio, Url: urlUrlScanio}
	sci, errs := ScanSubdomain(scanIo)
	if(errs != nil){
		fmt.Println("urlscan > ", errs.Error())
		sci = domain.SubDomains{}
	}


	rapid := &rapiddns.RapidDns{NameService: "rapiddns", Domain: dominio, Url: urlRapidDns}
	rapidns, errp := ScanSubdomain(rapid)
	if(errp != nil){
		fmt.Println("RpDns >", errp.Error())
		rapidns = domain.SubDomains{}
	}
	




	
	var subdomains []domain.Domain
	
	start := time.Now()
	
	fmt.Println("\n","Starting:")
	
	

	subdomainsStrings := append(crt.SubDomains, ht.SubDomains...)
	subdomainsStrings = append(subdomainsStrings, sci.SubDomains...)
	subdomainsStrings = append(subdomainsStrings, rapidns.SubDomains...)
	listClean := funcs.DeleteRepeat(subdomainsStrings)
	

	
	//Dominio padre
	dominioPadre := domain.Domain{Name: dominio, Ip: func()[]net.IP{ r, _ := funcs.CheckIp(dominio, true);return r}()}
	dominioPadre.FindCdn(&cdnList)
	subdomains = append(subdomains, dominioPadre)
	//Subdominios


	
	subdomains = Init(listClean, &cdnList)


	fmt.Println("\nResults: ")
	for _, v := range subdomains{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(style.YELLOW, v.Name, style.END ,style.GREEN, v.Ip, style.END ,"\n",style.Randcolor(),"Cdn: >", style.END, v.Cdns)
		fmt.Println("--------------------------------------------------------------------------------------------------")
	}


	end := time.Since(start)
	fmt.Println("Execution time:", end)
	//os.Exit(1)
	//fmt.Println(funcs.CheckIp(dominio, false))
	
	
}




func Init(lista []string, cdnlist *[]IPs.Cdn)[]domain.Domain{
	subdomains := []domain.Domain{}

	dmain := make(chan *domain.Domain, 10)
	var wg sync.WaitGroup
	limitElements := 100
	numThreads := 7



	if(len(lista) > limitElements){
		fmt.Println("Accelerating....")
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
					domaiin.FindCdn(cdnlist)
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
		domaiin.FindCdn(cdnlist)
		subdomains = append(subdomains, domaiin)

	}
	return subdomains

}


func ScanSubdomain(s dnsmikis.Scan)(domain.SubDomains, error){

	resp, err := s.CheckSubdomain()
	if(err != nil){
		
		//fmt.Printf("\r%s",s.ServiceName())
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