package Pingdom


type PingdomAgent struct {
	url 		string
	apiToken 	string
	checks 		Checks
}


func (pa PingdomAgent) Checks() Checks {
	return pa.checks
}


func NewPingdomAgent(url string, apiToken string) *PingdomAgent {
	return &PingdomAgent{url: url, apiToken: apiToken}
}



