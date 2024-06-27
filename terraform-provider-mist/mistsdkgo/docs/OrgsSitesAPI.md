# \OrgsSitesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSite**](OrgsSitesAPI.md#CreateOrgSite) | **Post** /api/v1/orgs/{org_id}/sites | createOrgSite
[**ListOrgSites**](OrgsSitesAPI.md#ListOrgSites) | **Get** /api/v1/orgs/{org_id}/sites | listOrgSites



## CreateOrgSite

> Site CreateOrgSite(ctx, orgId).Site(site).Execute()

createOrgSite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	site := *openapiclient.NewSite("Mist Office") // Site | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSitesAPI.CreateOrgSite(context.Background(), orgId).Site(site).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitesAPI.CreateOrgSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgSite`: Site
	fmt.Fprintf(os.Stdout, "Response from `OrgsSitesAPI.CreateOrgSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **site** | [**Site**](Site.md) | Request Body | 

### Return type

[**Site**](Site.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgSites

> []Site ListOrgSites(ctx, orgId).Limit(limit).Page(page).Execute()

listOrgSites



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSitesAPI.ListOrgSites(context.Background(), orgId).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitesAPI.ListOrgSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSites`: []Site
	fmt.Fprintf(os.Stdout, "Response from `OrgsSitesAPI.ListOrgSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**[]Site**](Site.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

