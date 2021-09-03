package http

import (
	"io/ioutil"
	"net/http"
)

func GetWithAuth(url, apiToken string) (int,[]byte,error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0,nil,err
	}

	if apiToken != "" {
		req.Header.Add("Authorization", "Bearer " + apiToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0,nil,err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode,respBytes,err
}