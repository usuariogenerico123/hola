package domain

import (
	"fake/IPs"
	"fake/funcs"
	"fake/style"
	"net"
)


type SubDomains struct{
	DomainName string 
	SubDomains []string
}




type Domain struct{
	Name string
	Ip  []net.IP
	//Subdomains []string 
	Cdns []string

}

func (d *Domain) FindCdn( cdn *[]IPs.Cdn){
	//fmt.Println(d.Ip)
	if (len(d.Ip) == 0){
		//fmt.Println("No host")
		d.Cdns = append(d.Cdns, style.RED,"Not a host", style.END)
		
	}

	
	for _, cdn := range *cdn{
		d.scanCdn(cdn.GetName(), cdn.GetIps())
	}
	

	//domainIps := d.Ip

	// 	//fmt.Println("Check cloudflare")
	// for _, v := range(domainIps){
	// 		//fmt.Println(v)
	// 		time.Sleep(50 * time.Millisecond)
	// 		isCloudflare := funcs.CheckCdn(v, IPs.CLOUDFLARE)
	// 		//fmt.Println(isCloudflare)
	// 		if(isCloudflare == true){
	// 			d.Cdns = append(d.Cdns, style.YELLOW, style.SUB,"Cloudflare", style.END)  
	// 		}else{
	// 			d.Cdns = append(d.Cdns, style.RED + "Cloudflare" + style.END) 
	// 		}

	// 	}

	
	}



func (d *Domain) scanCdn(cdnName string, cdnRange[]string){
	//ipsGoogle := IPs.GetIps("./IPs/googl.txt")
	for _, ip := range(d.Ip){
			
			isCdn:= funcs.CheckCdn(ip, cdnRange)
			if(isCdn){
				d.Cdns = append(d.Cdns, style.GREEN + cdnName + style.END)}
			// }else{
			// 	d.Cdns = append(d.Cdns, style.RED + cdnName + style.END)
			// }

		}
}

	



// func (c *Domain) GetIp() []net.IP{
// 	ip := funcs.CheckIp(c.Name)
// 	return ip
// }

// func (d *Domain) CheckSubdomains(){
// 	d.subdomains = append(d.subdomains, )
// }