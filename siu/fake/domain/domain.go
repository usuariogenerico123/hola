package domain

import (
	//"fake/funcs"
	"fake/IPs"
	"fake/funcs"
	"fake/style"


	//"fmt"
	"net"
	"time"
	
)


type SubDomains struct{
	DomainName string 
	SubDomains []string
}




type Domain struct{
	Name string
	Ip  []net.IP
	//Subdomains []string 
	Cdn []string

}

func (d *Domain) CheckNs(){
	//fmt.Println(d.Ip)
	if (len(d.Ip) == 0){
		//fmt.Println("No host")
		d.Cdn = append(d.Cdn, style.RED,"Not a host", style.END)
		
	}

	domainIps := d.Ip
	

		//fmt.Println("Check cloudflare")
	for _, v := range(domainIps){
			//fmt.Println(v)
			time.Sleep(50 * time.Millisecond)
			isCloudflare := funcs.CheckCdn(v, IPs.CLOUDFLARE)
			//fmt.Println(isCloudflare)
			if(isCloudflare == true){
				d.Cdn = append(d.Cdn, style.YELLOW, style.SUB,"Cloudflare", style.END)  
			}else{
				d.Cdn = append(d.Cdn, style.RED + "Cloudflare" + style.END) 
			}

		}

	//fmt.Println("Check cloudfront")
	ips:=IPs.GetIps("./IPs/front.txt")
	for _, x := range(domainIps){
			
			isCloudFront := funcs.CheckCdn(x, ips)
			if(isCloudFront){
				d.Cdn = append(d.Cdn, style.GREEN + "Cloudfront" + style.END)
			}else{
				d.Cdn = append(d.Cdn, style.RED + "Cloudfront" + style.END)
			}

		}


		//fmt.Println("Check fastly")
	
	for _, z := range(domainIps){
			
			isCloudFront := funcs.CheckCdn(z, IPs.FASTLY)
			if(isCloudFront){
				d.Cdn = append(d.Cdn, style.GREEN + "Fastly" + style.END)
			}else{
				d.Cdn = append(d.Cdn, style.RED + "Fastly" + style.END)
			}

		}

	//fmt.Println("Check Akamai")

	ipsAkamain := IPs.GetIps("./IPs/akamai.txt")
	for _, j := range(domainIps){
			
			isCloudFront := funcs.CheckCdn(j, ipsAkamain)
			if(isCloudFront){
				d.Cdn = append(d.Cdn, style.GREEN + "Akamai" + style.END)
			}else{
				d.Cdn = append(d.Cdn, style.RED + "Akamai" + style.END)
			}

		}

	//
	ipsGoogle := IPs.GetIps("./IPs/googl.txt")
	for _, j := range(domainIps){
			
			isGoogleCloud := funcs.CheckCdn(j, ipsGoogle)
			if(isGoogleCloud){
				d.Cdn = append(d.Cdn, style.GREEN + "Google" + style.END)
			}else{
				d.Cdn = append(d.Cdn, style.RED + "Google" + style.END)
			}

		}





	
	}



	

	



// func (c *Domain) GetIp() []net.IP{
// 	ip := funcs.CheckIp(c.Name)
// 	return ip
// }

// func (d *Domain) CheckSubdomains(){
// 	d.subdomains = append(d.subdomains, )
// }