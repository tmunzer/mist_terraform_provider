# \SitesRFTemplatesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteRfTemplateDerived**](SitesRFTemplatesAPI.md#GetSiteRfTemplateDerived) | **Get** /api/v1/sites/{site_id}/rftemplates/derived | getSiteRfTemplateDerived



## GetSiteRfTemplateDerived

> RfTemplate GetSiteRfTemplateDerived(ctx, siteId).Resolve(resolve).Execute()

getSiteRfTemplateDerived



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/mistsdkgo"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	resolve := true // bool | whether resolve the site variables (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRFTemplatesAPI.GetSiteRfTemplateDerived(context.Background(), siteId).Resolve(resolve).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRFTemplatesAPI.GetSiteRfTemplateDerived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteRfTemplateDerived`: RfTemplate
	fmt.Fprintf(os.Stdout, "Response from `SitesRFTemplatesAPI.GetSiteRfTemplateDerived`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRfTemplateDerivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolve** | **bool** | whether resolve the site variables | 

### Return type

[**RfTemplate**](RfTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
