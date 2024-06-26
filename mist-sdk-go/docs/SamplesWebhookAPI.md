# \SamplesWebhookAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Alarms**](SamplesWebhookAPI.md#Alarms) | **Post** /webhook_example/_alarm_ | alarms
[**AssetRaw**](SamplesWebhookAPI.md#AssetRaw) | **Post** /webhook_example/_asset_raw_ | assetRaw
[**Audits**](SamplesWebhookAPI.md#Audits) | **Post** /webhook_example/_audit_ | audits
[**ClientInfo**](SamplesWebhookAPI.md#ClientInfo) | **Post** /webhook_example/_client_info_ | clientJoin
[**ClientJoin**](SamplesWebhookAPI.md#ClientJoin) | **Post** /webhook_example/_client_join_ | clientJoin
[**ClientLatency**](SamplesWebhookAPI.md#ClientLatency) | **Post** /webhook_example/_client_latency_ | alarms
[**ClientSessions**](SamplesWebhookAPI.md#ClientSessions) | **Post** /webhook_example/_client_sessions_ | clientSessions
[**DeviceEvents**](SamplesWebhookAPI.md#DeviceEvents) | **Post** /webhook_example/_device_events_ | deviceEvents
[**DeviceUpDown**](SamplesWebhookAPI.md#DeviceUpDown) | **Post** /webhook_example/_device_updowns_ | deviceUpDown
[**DiscoveredRawRssi**](SamplesWebhookAPI.md#DiscoveredRawRssi) | **Post** /webhook_example/_discovered_raw_rssi_ | discovered-raw-rssi
[**Location**](SamplesWebhookAPI.md#Location) | **Post** /webhook_example/_location_ | location
[**LocationAsset**](SamplesWebhookAPI.md#LocationAsset) | **Post** /webhook_example/_location_asset_ | location
[**LocationCentrak**](SamplesWebhookAPI.md#LocationCentrak) | **Post** /webhook_example/_location_centrak_ | alarms
[**LocationClient**](SamplesWebhookAPI.md#LocationClient) | **Post** /webhook_example/_location_client_ | location
[**LocationSdk**](SamplesWebhookAPI.md#LocationSdk) | **Post** /webhook_example/_location_sdk_ | location
[**LocationUnclient**](SamplesWebhookAPI.md#LocationUnclient) | **Post** /webhook_example/_location_unclient_ | location
[**NacAccounting**](SamplesWebhookAPI.md#NacAccounting) | **Post** /webhook_example/_nac_accounting_ | nacAccounting
[**NacEvents**](SamplesWebhookAPI.md#NacEvents) | **Post** /webhook_example/_nac_events_ | nac_events
[**OccupancyAlerts**](SamplesWebhookAPI.md#OccupancyAlerts) | **Post** /webhook_example/_occupancy_alerts_ | occupancyAlerts
[**Ping**](SamplesWebhookAPI.md#Ping) | **Post** /webhook_example/_ping_ | ping
[**SdkclientScanData**](SamplesWebhookAPI.md#SdkclientScanData) | **Post** /webhook_example/_sdkclient_scan_data | sdkclientScanData
[**SiteSle**](SamplesWebhookAPI.md#SiteSle) | **Post** /webhook_example/_site_sle_ | site_sle
[**Zone**](SamplesWebhookAPI.md#Zone) | **Post** /webhook_example/_zone_ | zone



## Alarms

> Alarms(ctx).WebhookAlarms(webhookAlarms).Execute()

alarms



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
	webhookAlarms := *openapiclient.NewWebhookAlarms([]openapiclient.WebhookAlarmEvent{*openapiclient.NewWebhookAlarmEvent("Id_example", float32(123), "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61", "441a1214-6928-442a-8e92-e1d34b8ec6a6", int32(123), "Type_example")}, "Topic_example") // WebhookAlarms | **N.B.**: Fields like `aps`, `bssids`, `ssids` are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose `details` to include more event specific details.  Events specific fields for other alarm event type can be found with API https://api.mist.com/api/v1/const/alarm_defs, under “fields” array of /alarm_defs response object. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.Alarms(context.Background()).WebhookAlarms(webhookAlarms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.Alarms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookAlarms** | [**WebhookAlarms**](WebhookAlarms.md) | **N.B.**: Fields like &#x60;aps&#x60;, &#x60;bssids&#x60;, &#x60;ssids&#x60; are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose &#x60;details&#x60; to include more event specific details.  Events specific fields for other alarm event type can be found with API https://api.mist.com/api/v1/const/alarm_defs, under “fields” array of /alarm_defs response object. | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetRaw

> AssetRaw(ctx).WebhookAssetRaw(webhookAssetRaw).Execute()

assetRaw



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
	webhookAssetRaw := *openapiclient.NewWebhookAssetRaw([]openapiclient.WebhookAssetRawEvent{*openapiclient.NewWebhookAssetRawEvent("AssetId_example", int32(123), "DeviceId_example", "Mac_example", "MapId_example", float32(123), "MfgData_example", float32(123), "441a1214-6928-442a-8e92-e1d34b8ec6a6", float32(123))}, "Topic_example") // WebhookAssetRaw |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.AssetRaw(context.Background()).WebhookAssetRaw(webhookAssetRaw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.AssetRaw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookAssetRaw** | [**WebhookAssetRaw**](WebhookAssetRaw.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Audits

> Audits(ctx).WebhookAudits(webhookAudits).Execute()

audits



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
	webhookAudits := *openapiclient.NewWebhookAudits([]openapiclient.WebhookAuditEvent{*openapiclient.NewWebhookAuditEvent("AdminName_example", "DeviceId_example", "Id_example", "Message_example", "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61", "441a1214-6928-442a-8e92-e1d34b8ec6a6", "SrcIp_example", float32(123))}, "Topic_example") // WebhookAudits |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.Audits(context.Background()).WebhookAudits(webhookAudits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.Audits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookAudits** | [**WebhookAudits**](WebhookAudits.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientInfo

> ClientInfo(ctx).WebhookClientInfo(webhookClientInfo).Execute()

clientJoin



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
	webhookClientInfo := *openapiclient.NewWebhookClientInfo() // WebhookClientInfo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.ClientInfo(context.Background()).WebhookClientInfo(webhookClientInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.ClientInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookClientInfo** | [**WebhookClientInfo**](WebhookClientInfo.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientJoin

> ClientJoin(ctx).WebhookClientJoin(webhookClientJoin).Execute()

clientJoin



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
	webhookClientJoin := *openapiclient.NewWebhookClientJoin([]openapiclient.WebhookClientJoinEvent{*openapiclient.NewWebhookClientJoinEvent("Ap_example", "ApName_example", "Band_example", "Bssid_example", int32(123), float32(123), "Mac_example", "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61", float32(123), "441a1214-6928-442a-8e92-e1d34b8ec6a6", "SiteName_example", "Ssid_example", float32(123), float32(123), "WlanId_example")}, "Topic_example") // WebhookClientJoin |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.ClientJoin(context.Background()).WebhookClientJoin(webhookClientJoin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.ClientJoin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookClientJoin** | [**WebhookClientJoin**](WebhookClientJoin.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientLatency

> ClientLatency(ctx).WebhookClientLatency(webhookClientLatency).Execute()

alarms



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
	webhookClientLatency := *openapiclient.NewWebhookClientLatency() // WebhookClientLatency |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.ClientLatency(context.Background()).WebhookClientLatency(webhookClientLatency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.ClientLatency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientLatencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookClientLatency** | [**WebhookClientLatency**](WebhookClientLatency.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientSessions

> ClientSessions(ctx).WebhookClientSessions(webhookClientSessions).Execute()

clientSessions



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
	webhookClientSessions := *openapiclient.NewWebhookClientSessions([]openapiclient.WebhookClientSessionsEvent{*openapiclient.NewWebhookClientSessionsEvent("Ap_example", "ApName_example", "Band_example", "Bssid_example", "ClientFamily_example", "ClientManufacture_example", "ClientModel_example", "ClientOs_example", int32(123), float32(123), int32(123), float32(123), int32(123), "Mac_example", "NextAp_example", "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61", float32(123), "441a1214-6928-442a-8e92-e1d34b8ec6a6", "SiteName_example", "Ssid_example", int32(123), float32(123), float32(123), "WlanId_example")}, "Topic_example") // WebhookClientSessions |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.ClientSessions(context.Background()).WebhookClientSessions(webhookClientSessions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.ClientSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookClientSessions** | [**WebhookClientSessions**](WebhookClientSessions.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceEvents

> DeviceEvents(ctx).WebhookDeviceEvents(webhookDeviceEvents).Execute()

deviceEvents



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
	webhookDeviceEvents := *openapiclient.NewWebhookDeviceEvents([]openapiclient.WebhookDeviceEventsEvent{*openapiclient.NewWebhookDeviceEventsEvent("DeviceName_example", openapiclient.webhook_device_events_event_device_type(""), openapiclient.webhook_device_events_event_ev_type(""), "Mac_example", "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61", int32(123), "Type_example")}, "Topic_example") // WebhookDeviceEvents |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.DeviceEvents(context.Background()).WebhookDeviceEvents(webhookDeviceEvents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.DeviceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDeviceEvents** | [**WebhookDeviceEvents**](WebhookDeviceEvents.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceUpDown

> DeviceUpDown(ctx).WebhookDeviceUpdowns(webhookDeviceUpdowns).Execute()

deviceUpDown



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
	webhookDeviceUpdowns := *openapiclient.NewWebhookDeviceUpdowns([]openapiclient.WebhookDeviceUpdownsEvent{*openapiclient.NewWebhookDeviceUpdownsEvent("Ap_example", "ApName_example", "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61", "441a1214-6928-442a-8e92-e1d34b8ec6a6", "SiteName_example", float32(123), "Type_example")}, "Topic_example") // WebhookDeviceUpdowns |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.DeviceUpDown(context.Background()).WebhookDeviceUpdowns(webhookDeviceUpdowns).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.DeviceUpDown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceUpDownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDeviceUpdowns** | [**WebhookDeviceUpdowns**](WebhookDeviceUpdowns.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveredRawRssi

> DiscoveredRawRssi(ctx).WebhookDiscoveredRawRssi(webhookDiscoveredRawRssi).Execute()

discovered-raw-rssi



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
	webhookDiscoveredRawRssi := *openapiclient.NewWebhookDiscoveredRawRssi("Topic_example") // WebhookDiscoveredRawRssi |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.DiscoveredRawRssi(context.Background()).WebhookDiscoveredRawRssi(webhookDiscoveredRawRssi).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.DiscoveredRawRssi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveredRawRssiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDiscoveredRawRssi** | [**WebhookDiscoveredRawRssi**](WebhookDiscoveredRawRssi.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Location

> Location(ctx).WebhookLocation(webhookLocation).Execute()

location



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
	webhookLocation := *openapiclient.NewWebhookLocation([]openapiclient.WebhookLocationEvent{*openapiclient.NewWebhookLocationEvent("Id_example", "MapId_example", "441a1214-6928-442a-8e92-e1d34b8ec6a6", int32(123), "Type_example", int32(123), int32(123))}, "Topic_example") // WebhookLocation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.Location(context.Background()).WebhookLocation(webhookLocation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.Location``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookLocation** | [**WebhookLocation**](WebhookLocation.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationAsset

> LocationAsset(ctx).WebhookLocationAsset(webhookLocationAsset).Execute()

location



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
	webhookLocationAsset := *openapiclient.NewWebhookLocationAsset([]openapiclient.WebhookLocationAssetEventsInner{*openapiclient.NewWebhookLocationAssetEventsInner()}, "Topic_example") // WebhookLocationAsset |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.LocationAsset(context.Background()).WebhookLocationAsset(webhookLocationAsset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.LocationAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookLocationAsset** | [**WebhookLocationAsset**](WebhookLocationAsset.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationCentrak

> LocationCentrak(ctx).WebhookLocationCentrak(webhookLocationCentrak).Execute()

alarms



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
	webhookLocationCentrak := *openapiclient.NewWebhookLocationCentrak([]openapiclient.WebhookLocationCentrakEvent{*openapiclient.NewWebhookLocationCentrakEvent()}, "Topic_example") // WebhookLocationCentrak | **N.B.**: Fields like `aps`, `bssids`, `ssids` are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose `details` to include more event specific details.  Events specific fields for other alarm event type can be found with API https://api.mist.com/api/v1/const/alarm_defs, under “fields” array of /alarm_defs response object. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.LocationCentrak(context.Background()).WebhookLocationCentrak(webhookLocationCentrak).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.LocationCentrak``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationCentrakRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookLocationCentrak** | [**WebhookLocationCentrak**](WebhookLocationCentrak.md) | **N.B.**: Fields like &#x60;aps&#x60;, &#x60;bssids&#x60;, &#x60;ssids&#x60; are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose &#x60;details&#x60; to include more event specific details.  Events specific fields for other alarm event type can be found with API https://api.mist.com/api/v1/const/alarm_defs, under “fields” array of /alarm_defs response object. | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationClient

> LocationClient(ctx).WebhookLocationClient(webhookLocationClient).Execute()

location



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
	webhookLocationClient := *openapiclient.NewWebhookLocationClient([]openapiclient.WebhookLocationClientEventsInner{*openapiclient.NewWebhookLocationClientEventsInner()}, "Topic_example") // WebhookLocationClient |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.LocationClient(context.Background()).WebhookLocationClient(webhookLocationClient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.LocationClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookLocationClient** | [**WebhookLocationClient**](WebhookLocationClient.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationSdk

> LocationSdk(ctx).WebhookLocationSdk(webhookLocationSdk).Execute()

location



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
	webhookLocationSdk := *openapiclient.NewWebhookLocationSdk([]openapiclient.WebhookLocationSdkEventsInner{*openapiclient.NewWebhookLocationSdkEventsInner()}, "Topic_example") // WebhookLocationSdk |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.LocationSdk(context.Background()).WebhookLocationSdk(webhookLocationSdk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.LocationSdk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationSdkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookLocationSdk** | [**WebhookLocationSdk**](WebhookLocationSdk.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationUnclient

> LocationUnclient(ctx).WebhookLocationUnclient(webhookLocationUnclient).Execute()

location



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
	webhookLocationUnclient := *openapiclient.NewWebhookLocationUnclient([]openapiclient.WebhookLocationClientEventsInner{*openapiclient.NewWebhookLocationClientEventsInner()}, "Topic_example") // WebhookLocationUnclient |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.LocationUnclient(context.Background()).WebhookLocationUnclient(webhookLocationUnclient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.LocationUnclient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationUnclientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookLocationUnclient** | [**WebhookLocationUnclient**](WebhookLocationUnclient.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NacAccounting

> NacAccounting(ctx).WebhookNacAccounting(webhookNacAccounting).Execute()

nacAccounting



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
	webhookNacAccounting := *openapiclient.NewWebhookNacAccounting() // WebhookNacAccounting |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.NacAccounting(context.Background()).WebhookNacAccounting(webhookNacAccounting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.NacAccounting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNacAccountingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookNacAccounting** | [**WebhookNacAccounting**](WebhookNacAccounting.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NacEvents

> NacEvents(ctx).WebhookNacEvents(webhookNacEvents).Execute()

nac_events



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
	webhookNacEvents := *openapiclient.NewWebhookNacEvents() // WebhookNacEvents |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.NacEvents(context.Background()).WebhookNacEvents(webhookNacEvents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.NacEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNacEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookNacEvents** | [**WebhookNacEvents**](WebhookNacEvents.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OccupancyAlerts

> OccupancyAlerts(ctx).WebhookOccupancyAlerts(webhookOccupancyAlerts).Execute()

occupancyAlerts



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
	webhookOccupancyAlerts := *openapiclient.NewWebhookOccupancyAlerts([]openapiclient.WebhookOccupancyAlertsEvent{*openapiclient.NewWebhookOccupancyAlertsEvent("441a1214-6928-442a-8e92-e1d34b8ec6a6", "SiteName_example")}, "Topic_example") // WebhookOccupancyAlerts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.OccupancyAlerts(context.Background()).WebhookOccupancyAlerts(webhookOccupancyAlerts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.OccupancyAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOccupancyAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookOccupancyAlerts** | [**WebhookOccupancyAlerts**](WebhookOccupancyAlerts.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> Ping(ctx).WebhookPing(webhookPing).Execute()

ping



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
	webhookPing := *openapiclient.NewWebhookPing([]openapiclient.WebhookPingEvent{*openapiclient.NewWebhookPingEvent("Id_example", "Name_example", "43e9c864-a7e4-4310-8031-d9817d2c5a43", float32(123))}, "Topic_example") // WebhookPing |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.Ping(context.Background()).WebhookPing(webhookPing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.Ping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPing** | [**WebhookPing**](WebhookPing.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SdkclientScanData

> SdkclientScanData(ctx).WebhookSdkclientScanData(webhookSdkclientScanData).Execute()

sdkclientScanData



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
	webhookSdkclientScanData := *openapiclient.NewWebhookSdkclientScanData([]openapiclient.WebhookSdkclientScanDataEvent{*openapiclient.NewWebhookSdkclientScanDataEvent("ConnectionAp_example", "ConnectionBand_example", "ConnectionBssid_example", int32(123), float32(123), float32(123), "Mac_example", "03e9c864-a7e4-4310-8031-d9817d2c5a43")}, openapiclient.webhook_sdkclient_scan_data_topic("")) // WebhookSdkclientScanData |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.SdkclientScanData(context.Background()).WebhookSdkclientScanData(webhookSdkclientScanData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.SdkclientScanData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSdkclientScanDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookSdkclientScanData** | [**WebhookSdkclientScanData**](WebhookSdkclientScanData.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteSle

> SiteSle(ctx).WebhookSiteSle(webhookSiteSle).Execute()

site_sle



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
	webhookSiteSle := *openapiclient.NewWebhookSiteSle() // WebhookSiteSle |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.SiteSle(context.Background()).WebhookSiteSle(webhookSiteSle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.SiteSle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteSleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookSiteSle** | [**WebhookSiteSle**](WebhookSiteSle.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Zone

> Zone(ctx).WebhookZone(webhookZone).Execute()

zone



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
	webhookZone := *openapiclient.NewWebhookZone([]openapiclient.WebhookZoneEvent{*openapiclient.NewWebhookZoneEvent("Id_example", "MapId_example", "03e9c864-a7e4-4310-8031-d9817d2c5a43", int32(123), openapiclient.webhook_zone_event_trigger(""), "Type_example", "ZoneId_example")}, "Topic_example") // WebhookZone |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamplesWebhookAPI.Zone(context.Background()).WebhookZone(webhookZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamplesWebhookAPI.Zone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookZone** | [**WebhookZone**](WebhookZone.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

