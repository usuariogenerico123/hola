package dnsmikis

import "fake/domain"






type Scan interface{
	
	ServiceName()string
	CheckSubdomain() (domain.SubDomains, error)
}