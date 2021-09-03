package Pingdom

type Check struct {
	Id       			int64		`json:"id"`
	Created  			int64 		`json:"created"`
	Name     			string 		`json:"name"`
	Hostname 			string 		`json:"hostname"`
	Resolution        	int      	`json:"resolution"`
	Type              	string   	`json:"type"`
	Ipv6              	bool     	`json:"ipv6"`
	VerifyCertificate 	bool     	`json:"verify_certificate"`
	Lasterrortime     	int64      	`json:"lasterrortime"`
	Lasttesttime      	int64      	`json:"lasttesttime"`
	Lastresponsetime  	int64      	`json:"lastresponsetime"`
	Lastdownstart     	int64      	`json:"lastdownstart"`
	Lastdownend       	int64      	`json:"lastdownend"`
	Status            	string   	`json:"status"`
	Maintenanceids    	[]string 	`json:"maintenanceids"`
}

type Checks struct {
	Checks []Check `json:"checks"`
	Counts struct {
		Total    int `json:"total"`
		Limited  int `json:"limited"`
		Filtered int `json:"filtered"`
	} `json:"counts"`
}