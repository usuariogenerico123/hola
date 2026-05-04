package IPs

import (
	"fmt"
	"os"
	"strings"
)

type Cdn interface{
	GetName()string
	GetIps()[]string
}


type Cloudflare struct{
	Name string 
	IpRange []string
}
func (c *Cloudflare)GetName()string{
	return c.Name
}
func (c *Cloudflare)GetIps()[]string{
	return c.IpRange
}
type Google struct{
	Name string 
	IpRange []string
}
func (g *Google)GetName()string{
	return g.Name
}
func (g *Google)GetIps()[]string{
	return g.IpRange
}



type Akamai struct {
	Name string 
	IpRange []string
}
func (a *Akamai)GetName()string{
	return a.Name
}
func (a *Akamai)GetIps()[]string{
	return a.IpRange
}

type Fastly struct{
	Name string 
	IpRange []string
}
func (f *Fastly)GetName()string{
	return f.Name
}
func (f *Fastly)GetIps()[]string{
	return f.IpRange
}


type Aws struct{
	Name string 
	IpRange []string
}
func (a *Aws)GetName()string{
	return a.Name
}
func (a *Aws)GetIps()[]string{
	return a.IpRange
}





type IpRanges struct{
	
	List []Cdn
	IPsPath string
}

func (i *IpRanges) getIpsFromFile(fileName string)[]string{
	data, err := os.ReadFile(i.IPsPath + "/"+fileName)
	if(err != nil){
		fmt.Println("Error lectura .txt", err)
		return []string{""}
	}
	return strings.Fields(string(data))
}

func (i *IpRanges)Load(){
	cloudflare := &Cloudflare{Name: "cloudflare", IpRange: []string{"173.245.48.0/20","103.21.244.0/22","103.22.200.0/22","103.31.4.0/22","141.101.64.0/18","108.162.192.0/18","190.93.240.0/20","188.114.96.0/20","197.234.240.0/22","198.41.128.0/17","162.158.0.0/15","104.16.0.0/13","104.24.0.0/14","172.64.0.0/13","131.0.72.0/22",}}
	
	fastly := &Fastly{Name: "fastly", IpRange: []string{"23.235.32.0/20","43.249.72.0/22","103.244.50.0/24","103.245.222.0/23","103.245.224.0/24","104.156.80.0/20","140.248.64.0/18","140.248.128.0/17","146.75.0.0/17","151.101.0.0/16","157.52.64.0/18","167.82.0.0/17","167.82.128.0/20","167.82.160.0/20","167.82.224.0/20","172.111.64.0/18","185.31.16.0/22","199.27.72.0/21","199.232.0.0/16"}}
	
	akamai := &Akamai{Name:"akamai"}
	akamai.IpRange = i.getIpsFromFile(akamai.GetName() + ".txt" )

	cloudfront := &Aws{Name:"cloudfront"}
	cloudfront.IpRange = i.getIpsFromFile(cloudfront.GetName() + ".txt")

	google :=  &Google{Name:"google"}
	google.IpRange =i.getIpsFromFile( google.GetName() + ".txt")


	i.List = []Cdn{cloudflare, fastly, akamai, cloudfront, google}
	
}

func (i *IpRanges)GetListCdn() []Cdn{
	return i.List

}