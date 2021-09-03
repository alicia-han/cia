package http

import (
	"testing"
)

func TestGetWithAuth(t *testing.T) {
	type args struct {
		url      string
		apiToken string
	}

	argWithoutApiToken :=  & args{
		url: "https://alertmanager.staging-fr-onroad.com/api/v2/status",
		apiToken: "",
	}

	test1,err := GetWithAuth(argWithoutApiToken.url, argWithoutApiToken.apiToken)

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(string(test1))

}
