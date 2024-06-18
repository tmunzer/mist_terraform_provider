# \SitesSettingAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteSetting**](SitesSettingAPI.md#GetSiteSetting) | **Get** /api/v1/sites/{site_id}/setting | getSiteSetting
[**UpdateSiteSettings**](SitesSettingAPI.md#UpdateSiteSettings) | **Put** /api/v1/sites/{site_id}/setting | updateSiteSettings



## GetSiteSetting

> SiteSetting GetSiteSetting(ctx, siteId).Execute()

getSiteSetting



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSettingAPI.GetSiteSetting(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSettingAPI.GetSiteSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSetting`: SiteSetting
	fmt.Fprintf(os.Stdout, "Response from `SitesSettingAPI.GetSiteSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiteSetting**](SiteSetting.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteSettings

> SiteSetting UpdateSiteSettings(ctx, siteId).SiteSetting(siteSetting).Execute()

updateSiteSettings



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	siteSetting := *openapiclient.NewSiteSetting() // SiteSetting | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSettingAPI.UpdateSiteSettings(context.Background(), siteId).SiteSetting(siteSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSettingAPI.UpdateSiteSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteSettings`: SiteSetting
	fmt.Fprintf(os.Stdout, "Response from `SitesSettingAPI.UpdateSiteSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteSetting** | [**SiteSetting**](SiteSetting.md) | Request Body | 

### Return type

[**SiteSetting**](SiteSetting.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

