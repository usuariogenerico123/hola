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

	if (len(d.Ip) == 0){
		//fmt.Println("No host")
		d.Cdns = append(d.Cdns, style.RED,"Not a host", style.END)	
	}

	
	for _, cdns := range *cdn{
		d.scanCdn(cdns.GetName(), cdns.GetIps())
	}

	
	}



func (d *Domain) scanCdn(cdnName string, cdnRange []string){
	
	for _, ip := range(d.Ip){
			
			isBunnyCdn := funcs.CheckBunnyCDN(ip, cdnRange)
			if(isBunnyCdn){
				d.Cdns = append(d.Cdns, style.SUB, style.GREEN + cdnName + style.END)
				return
			}

			isCdn:= funcs.CheckCdn(cdnName, ip, cdnRange)
			if(isCdn){
				d.Cdns = append(d.Cdns, style.SUB, style.GREEN + cdnName + style.END)
				return
			}
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