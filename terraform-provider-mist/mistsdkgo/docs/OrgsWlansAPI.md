# \OrgsWlansAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWlan**](OrgsWlansAPI.md#CreateOrgWlan) | **Post** /api/v1/orgs/{org_id}/wlans | createOrgWlan
[**DeleteOrgWlan**](OrgsWlansAPI.md#DeleteOrgWlan) | **Delete** /api/v1/orgs/{org_id}/wlans/{wlan_id} | deleteOrgWlan
[**GetOrgWLAN**](OrgsWlansAPI.md#GetOrgWLAN) | **Get** /api/v1/orgs/{org_id}/wlans/{wlan_id} | getOrgWLAN
[**ListOrgWlans**](OrgsWlansAPI.md#ListOrgWlans) | **Get** /api/v1/orgs/{org_id}/wlans | listOrgWlans
[**UpdateOrgWlan**](OrgsWlansAPI.md#UpdateOrgWlan) | **Put** /api/v1/orgs/{org_id}/wlans/{wlan_id} | updateOrgWlan



## CreateOrgWlan

> Wlan CreateOrgWlan(ctx, orgId).Wlan(wlan).Execute()

createOrgWlan



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
	wlan := *openapiclient.NewWlan("corporate") // Wlan | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWlansAPI.CreateOrgWlan(context.Background(), orgId).Wlan(wlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWlansAPI.CreateOrgWlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgWlan`: Wlan
	fmt.Fprintf(os.Stdout, "Response from `OrgsWlansAPI.CreateOrgWlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgWlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wlan** | [**Wlan**](Wlan.md) | Request Body | 

### Return type

[**Wlan**](Wlan.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgWlan

> DeleteOrgWlan(ctx, orgId, wlanId).Execute()

deleteOrgWlan



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
	wlanId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsWlansAPI.DeleteOrgWlan(context.Background(), orgId, wlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWlansAPI.DeleteOrgWlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgWlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgWLAN

> []Wlan GetOrgWLAN(ctx, orgId, wlanId).Execute()

getOrgWLAN



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
	wlanId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWlansAPI.GetOrgWLAN(context.Background(), orgId, wlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWlansAPI.GetOrgWLAN``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgWLAN`: []Wlan
	fmt.Fprintf(os.Stdout, "Response from `OrgsWlansAPI.GetOrgWLAN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgWLANRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Wlan**](Wlan.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgWlans

> []Wlan ListOrgWlans(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgWlans



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWlansAPI.ListOrgWlans(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWlansAPI.ListOrgWlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgWlans`: []Wlan
	fmt.Fprintf(os.Stdout, "Response from `OrgsWlansAPI.ListOrgWlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Wlan**](Wlan.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgWlan

> Wlan UpdateOrgWlan(ctx, orgId, wlanId).Wlan(wlan).Execute()

updateOrgWlan



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
	wlanId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wlan := *openapiclient.NewWlan("corporate") // Wlan | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWlansAPI.UpdateOrgWlan(context.Background(), orgId, wlanId).Wlan(wlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWlansAPI.UpdateOrgWlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgWlan`: Wlan
	fmt.Fprintf(os.Stdout, "Response from `OrgsWlansAPI.UpdateOrgWlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgWlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wlan** | [**Wlan**](Wlan.md) | Request Body | 

### Return type

[**Wlan**](Wlan.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

