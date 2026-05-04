package IPs

import (
	"fmt"
	"os"
	"strings"
)




var(

	CLOUDFLARE []string = []string{
		"173.245.48.0/20",
		"103.21.244.0/22",
		"103.22.200.0/22",
		"103.31.4.0/22",
		"141.101.64.0/18",	
		"108.162.192.0/18",
		"190.93.240.0/20",
		"188.114.96.0/20",
		"197.234.240.0/22",
		"198.41.128.0/17",
		"162.158.0.0/15",
		"104.16.0.0/13",
		"104.24.0.0/14",
		"172.64.0.0/13",
		"131.0.72.0/22",
	}

	CLOUDFRONT []string = []string{}

	FASTLY []string = []string{"23.235.32.0/20","43.249.72.0/22","103.244.50.0/24","103.245.222.0/23","103.245.224.0/24","104.156.80.0/20","140.248.64.0/18","140.248.128.0/17","146.75.0.0/17","151.101.0.0/16","157.52.64.0/18","167.82.0.0/17","167.82.128.0/20","167.82.160.0/20","167.82.224.0/20","172.111.64.0/18","185.31.16.0/22","199.27.72.0/21","199.232.0.0/16"}


	//AKAMAI []string = []string{}

)


type Cloudflare struct{
	Name string
	IpRange []string
}

type Google struct{
	Name string 
	IpRange []string
}

type Akamai struct {
	Name string 
	IpRange []string
}

type Fastly struct{
	Name string 
	IpRange []string
}

type Aws struct{
	Name string 
	IpRange []string
}



type IpRanges struct{
	Cloudflare  
	Akamai 
	Fastly 
	Google
	Aws

	IpsPath string
}

func (i *IpRanges) GetIps()[]string{
	data, err := os.ReadFile(i.IpsPath)
	if(err != nil){
		fmt.Println("Error lectura .txt", err)
		return []string{""}
	}
	return strings.Fields(string(data))
}
func (i *IpRanges)Load(){
	i.Cloudflare.Name = "Cloudflare"
	i.Fastly.Name = "Fastly"
	i.Akamai.Name = "Akamai"
	i.Aws.Name = "Cloudfront"
	i.Google.Name = "Google"

	i.Cloudflare.IpRange = []string{"173.245.48.0/20","103.21.244.0/22","103.22.200.0/22","103.31.4.0/22","141.101.64.0/18","108.162.192.0/18","190.93.240.0/20","188.114.96.0/20","197.234.240.0/22","198.41.128.0/17","162.158.0.0/15","104.16.0.0/13","104.24.0.0/14","172.64.0.0/13","131.0.72.0/22",}
	i.Fastly.IpRange = []string{"23.235.32.0/20","43.249.72.0/22","103.244.50.0/24","103.245.222.0/23","103.245.224.0/24","104.156.80.0/20","140.248.64.0/18","140.248.128.0/17","146.75.0.0/17","151.101.0.0/16","157.52.64.0/18","167.82.0.0/17","167.82.128.0/20","167.82.160.0/20","167.82.224.0/20","172.111.64.0/18","185.31.16.0/22","199.27.72.0/21","199.232.0.0/16"}
	i.Akamai.IpRange = i.GetIps()
	i.Aws.IpRange = i.GetIps()
	i.Google.IpRange = i.GetIps()

}