package datadog

import (
	"context"
	"encoding/json"
	"fmt"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	log "github.com/sirupsen/logrus"
	"os"
)

func ApiClient()  {
	ctx := datadog.NewDefaultContext(context.Background())

	configuration := datadog.NewConfiguration()

	apiClient := datadog.NewAPIClient(configuration)

	resp, r, err := apiClient.AuthenticationApi.Validate(ctx)

	if err != nil {
		log.Warn("Error auth validate: ", err)
		log.Warn("Full Http Response: ", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from AuthenticationApi.Validate:\n%s\n", responseContent)
}