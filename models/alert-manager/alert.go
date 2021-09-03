package alert_manager

type Alert struct {
	Annotations struct{
		Description 	string `json:"description"`
		RunbookUrl 		string `json:"runbook_url"`
		Summary 		string `json:"summary"`
	} `json:"annotations"`
	StartsAt string `json:"starts_at"`
	Status struct{
		State string `json:"state"`
	} `json:"status"`
	Labels struct {
		Alertname 	string `json:"alertname"`
		Container 	string `json:"container"`
		Namespace 	string `json:"namespace"`
		Pod 		string `json:"pod"`
		Severity	string `json:"severity"`
	} `json:"labels"`
}

	
