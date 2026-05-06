package menu

import (
	"fake/IPs"
	"fake/funcs"
	"fmt"
	"sync"
	"time"
	"fake/domain"
	"fake/style"
	"fake/dnsmikis/cert"
	"fake/dnsmikis/hacktarget"
	"fake/dnsmikis/rapiddns"
	"fake/dnsmikis/urlscan"
	"fake/dnsmikis"
	"net"
)



//---------Buscar cdn de una sola ip
func CheckCdnOnly(ips *[]IPs.Cdn,ip string){
	fmt.Println("")
	fmt.Println("SEARCHING IP CDN >:",  ip)
	time.Sleep(1 * time.Second)
	for _,v := range *ips{
		b := funcs.CheckBunnyCDN(net.ParseIP(ip), v.GetIps())
		if(b){
			fmt.Println("bunnycdn")
			return
		}
		

		cdn := funcs.CheckCdn(v.GetName(), net.ParseIP(ip), v.GetIps())
		if(cdn){
			fmt.Println(v.GetName())
			return
		}
	}
	fmt.Println("")
	fmt.Println("CDN NOT FOUND")
	fmt.Println("\nIP: ", ip)
}




func CheckAllSubdomain(cdnList *[]IPs.Cdn , dominio string){
	
	if(len(funcs.CheckNs(dominio)) == 0){
		fmt.Printf(style.RED + "Domain: %s not found\n"+style.END , dominio  )
		return
	}

	//-------Urls------------

	urlCrt := fmt.Sprintf("https://crt.sh/?q=%s&output=json", dominio)
	urlHtarget := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", dominio)
	urlUrlScanio := fmt.Sprintf("https://urlscan.io/api/v1/search/?q=domain:%s", dominio)
	urlRapidDns := fmt.Sprintf("https://rapiddns.io/subdomain/%s?full=1", dominio)

	
	//-----------------CRT.SH--------------------------------
	crtSh := &cert.CrtSh{NameService:"crt.sh", Domain:dominio, Url: urlCrt}
	crt, err := ScanSubdomain(crtSh)
	if(err != nil){
		
		intentos := 10
		
		respCrt := make(chan domain.SubDomains)
		go func(){
			ok := false
			for v := range intentos-1{
				fmt.Printf("\rTrying: %s%d%s", style.Randcolor() ,v ,style.END)
				time.Sleep(1 * time.Second)
				crt, err = ScanSubdomain(crtSh)
				if(len(crt.SubDomains) > 0){
					respCrt <- crt
					fmt.Printf("\r%s", "---- ok ----")
					ok = true
					break
				}
			}
			
			if(ok == false){
				fmt.Printf("\rcrt.sh > %s", err.Error())
				time.Sleep(1 * time.Second)
				respCrt <- domain.SubDomains{}

			}
			
		}()
		crt = <- respCrt
		
	}
	
	//-----------------HACKERTARGET.COM----------------------------------------
	hackTarget := &hacktarget.Htarget{NameService:"hackertarget", Domain: dominio, Url: urlHtarget}
	ht, errt := ScanSubdomain(hackTarget)
	if(errt != nil){
		fmt.Println("\nhtarget> ",errt.Error())
		ht = domain.SubDomains{}
	}

	//-------------------URLSCAN.IO-----------------------------------------

	scanIo := &urlscan.UrlScan{NameService: "urlscan.io", Domain: dominio, Url: urlUrlScanio}
	sci, errs := ScanSubdomain(scanIo)
	if(errs != nil){
		fmt.Println("urlscan > ", errs.Error())
		sci = domain.SubDomains{}
	}

	//----------------RAPIDDNS.COM------------------------------------------

	rapid := &rapiddns.RapidDns{NameService: "rapiddns", Domain: dominio, Url: urlRapidDns}
	rapidns, errp := ScanSubdomain(rapid)
	if(errp != nil){
		fmt.Println("RpDns >", errp.Error())
		rapidns = domain.SubDomains{}
	}
	//-----------------------------------------------------------

	//-----INICIO LISTA DE struct DOMINIOS (SUBDOMINIOS)
	var subdomains []domain.Domain
	
	start := time.Now()
	//fmt.Print("\n")
	fmt.Printf("\r%s", "Starting:")
	
	
	//-----UNIENDO LISTAS DE SUBDOMINIOS-------
	subdomainsStrings := append(crt.SubDomains, ht.SubDomains...)
	subdomainsStrings = append(subdomainsStrings, sci.SubDomains...)
	subdomainsStrings = append(subdomainsStrings, rapidns.SubDomains...)
	listClean := funcs.DeleteRepeat(subdomainsStrings)
	

	
	//Dominio padre
	dominioPadre := domain.Domain{Name: dominio, Ip: func()[]net.IP{ r, _ := funcs.CheckIp(dominio, true);return r}()}
	dominioPadre.FindCdn(cdnList)

	subdomains = append(subdomains, dominioPadre)
	


	//------ INICIO ----
	subdomains = Start(listClean, cdnList)


	//-------RESULTADOS
	fmt.Printf("\r%s","------------------------Results--------------------------\n")
	for n, v := range subdomains{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(n+1,style.YELLOW, v.Name, style.END ,style.GREEN, v.Ip, style.END) 
		fmt.Println(style.Randcolor() + "   Cdn: >" + style.END, v.Cdns)
		fmt.Println("---------------------------------------------------------------------------")
	}
	end := time.Since(start)
	fmt.Println("Execution time:", end)
	

}


func Start(lista []string, cdnlist *[]IPs.Cdn)[]domain.Domain{
	subdomains := []domain.Domain{}
	dmain := make(chan *domain.Domain, 10)
	var wg sync.WaitGroup
	limitElements := 100
	numThreads := 7


	// ---- SI HAY MAS DE 100 SUBDOMINIOS CREAMOS 7 GOROUTINES PARALELAS
	if(len(lista) > limitElements){
		fmt.Printf("\r%s", "Accelerating....\n")
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

	//-----SI SON MENOS DE 100 SUBDOMIIOS SE HACE ESCANEO NORMAL---

	for _, x := range lista{
		ip, _ := funcs.CheckIp(x, true)
		domaiin := domain.Domain{Name: x, Ip: ip }
		domaiin.FindCdn(cdnlist)
		subdomains = append(subdomains, domaiin)

	}

	//-----RETORNAMOS RESULTADOS
	return subdomains

}



// ------PARA PROBAR LA INTERFACE :| 
func ScanSubdomain(s dnsmikis.Scan)(domain.SubDomains, error){

	resp, err := s.CheckSubdomain()
	if(err != nil){
		
		//fmt.Printf("\r%s",s.ServiceName())
		return resp, err
	}
	return resp, nil

}




func Help()string{
	return `

Use: app [options] <arguments>
Options:
	--ip <IP>           "Scann all CDN for this ip"
	--domain <DOMAIN>   "Scann all CDN and Subdomains for this domain"	
	--help		    "Help"
Example:
	app --ip  123.123.123.123
	app --domain  mydomain.com
`
}

