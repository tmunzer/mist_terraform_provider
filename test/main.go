package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "test/github.com/tmunzer/mist-sdk-go"
)

func main() {
	//orgId := "203d3d02-dbc0-4c1b-9f41-76896a3330f4" // string |
	//ctx := context.WithValue(context.Background(), openapiclient.ContextServerIndex, 0)

	org := *openapiclient.NewOrg("test1234")

	configuration := openapiclient.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", "Token ")
	apiClient := openapiclient.NewAPIClient(configuration)

	resp, r, err := apiClient.OrgsAPI.CreateOrg(context.Background()).Org(org).Execute()

	//fmt.Fprintf(os.Stdout, "%v\n\n", apiClient.GetConfig())
	//resp, r, err := apiClient.OrgsAPI.GetOrg(ctx, orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrg`: Org
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetOrg`: %v\n", resp)
	//fmt.Fprint(os.Stdout, "\n\n", resp.Name)
	fmt.Fprint(os.Stdout, "\n\n", resp.GetId())
}
