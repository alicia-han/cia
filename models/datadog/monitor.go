package datadog

import (
	"context"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func AllMonitorDetails()  ([]datadog.Monitor, *http.Response, error){
	ctx := datadog.NewDefaultContext(context.Background())

	groupStates := "alert" // string | When specified, shows additional information about the group states. Choose one or more from `all`, `alert`, `warn`, and `no data`. (optional)
	//name := "name_example" // string | A string to filter monitors by name. (optional)
	//tags := "host:host0" // string | A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. For example, `host:host0`. (optional)
	//monitorTags := "service:my-app" // string | A comma separated list indicating what service and/or custom tags, if any, should be used to filter the list of monitors. Tags created in the Datadog UI automatically have the service key prepended. For example, `service:my-app`. (optional)
	//withDowntimes := true // bool | If this argument is set to true, then the returned data includes all current downtimes for each monitor. (optional)
	//idOffset := int64(789) // int64 | Monitor ID offset. (optional)
	//page := int64(0) // int64 | The page to start paginating from. If this argument is not specified, the request returns all monitors without pagination. (optional)
	//pageSize := int32(20) // int32 | The number of monitors to return per page. If the page argument is not specified, the default behavior returns all monitors without a `page_size` limit. However, if page is specified and `page_size` is not, the argument defaults to 100. (optional)
	optionalParams := datadog.ListMonitorsOptionalParameters{
		GroupStates: &groupStates,
		//Name: &name,
		//Tags: &tags,
		//MonitorTags: &monitorTags,
		//WithDowntimes: &withDowntimes,
		//IdOffset: &idOffset,
		//Page: &page,
		//PageSize: &pageSize,
	}

	configuration := datadog.NewConfiguration()

	apiClient := datadog.NewAPIClient(configuration)
	return apiClient.MonitorsApi.ListMonitors(ctx, optionalParams)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ListMonitors`: %v\n", err)
	//	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	//}
	// response from `ListMonitors`: []Monitor
	//responseContent, _ := json.MarshalIndent(resp, "", "  ")
	//fmt.Fprintf(os.Stdout, "Response from MonitorsApi.ListMonitors:\n%s\n", responseContent)
}

func GetActiveAlertsFromDataDog() (report []string) {
	monitors, resp, err := AllMonitorDetails()
	log.Info("Data Dog Monitor API visited")
	if err != nil || resp.StatusCode != 200 {
		log.Warn("err: ", err)
		log.Warn("response code: ", resp.StatusCode)
		return []string {
			"Error Getting Monitor Info From DataDog",
		}
	}

	for _, monitor := range monitors {
		if *monitor.OverallState != "OK" {
			report = append(report, *monitor.Name)
		}
	}
	return
}
