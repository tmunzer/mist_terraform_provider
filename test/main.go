package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "test/github.com/tmunzer/mist-sdk-go"
)

func main() {
	orgId := "d0f2ae46-b6f9-4f5b-b8ba-6d9a28556035" // string |
	//ctx := context.WithValue(context.Background(), openapiclient.ContextServerIndex, 0)

	site := *openapiclient.NewSite("test1234")
	site.SetAddress("123 rue de ici")

	configuration := openapiclient.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", "Token ")
	apiClient := openapiclient.NewAPIClient(configuration)

	//resp, r, err := apiClient.OrgsAPI.CreateOrg(context.Background()).Org(org).Execute()
	resp, r, err := apiClient.OrgsSitesAPI.CreateOrgSite(context.Background(), orgId).Site(site).Execute()
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
