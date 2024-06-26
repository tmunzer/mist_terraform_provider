# WebhookDeviceEventsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to **string** | (will be deprecated soon; please use mac instead) ap mac | [optional] 
**ApName** | Pointer to **string** | (will be deprecated soon; please use device_name instead) ap name | [optional] 
**AuditId** | Pointer to **string** | (optional) audit id | [optional] 
**DeviceName** | **string** | device name | 
**DeviceType** | [**WebhookDeviceEventsEventDeviceType**](WebhookDeviceEventsEventDeviceType.md) |  | 
**EvType** | [**WebhookDeviceEventsEventEvType**](WebhookDeviceEventsEventEvType.md) |  | 
**Mac** | **string** | device mac | 
**OrgId** | **string** |  | [readonly] 
**Reason** | Pointer to **string** | (optional) event reason | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SiteName** | Pointer to **string** | site name | [optional] 
**Text** | Pointer to **string** | (optional) event description | [optional] 
**Timestamp** | **int32** | time the event occurred e.g. 1565987313 | 
**Type** | **string** | event type | 

## Methods

### NewWebhookDeviceEventsEvent

`func NewWebhookDeviceEventsEvent(deviceName string, deviceType WebhookDeviceEventsEventDeviceType, evType WebhookDeviceEventsEventEvType, mac string, orgId string, timestamp int32, type_ string, ) *WebhookDeviceEventsEvent`

NewWebhookDeviceEventsEvent instantiates a new WebhookDeviceEventsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeviceEventsEventWithDefaults

`func NewWebhookDeviceEventsEventWithDefaults() *WebhookDeviceEventsEvent`

NewWebhookDeviceEventsEventWithDefaults instantiates a new WebhookDeviceEventsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *WebhookDeviceEventsEvent) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *WebhookDeviceEventsEvent) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *WebhookDeviceEventsEvent) SetAp(v string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *WebhookDeviceEventsEvent) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetApName

`func (o *WebhookDeviceEventsEvent) GetApName() string`

GetApName returns the ApName field if non-nil, zero value otherwise.

### GetApNameOk

`func (o *WebhookDeviceEventsEvent) GetApNameOk() (*string, bool)`

GetApNameOk returns a tuple with the ApName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApName

`func (o *WebhookDeviceEventsEvent) SetApName(v string)`

SetApName sets ApName field to given value.

### HasApName

`func (o *WebhookDeviceEventsEvent) HasApName() bool`

HasApName returns a boolean if a field has been set.

### GetAuditId

`func (o *WebhookDeviceEventsEvent) GetAuditId() string`

GetAuditId returns the AuditId field if non-nil, zero value otherwise.

### GetAuditIdOk

`func (o *WebhookDeviceEventsEvent) GetAuditIdOk() (*string, bool)`

GetAuditIdOk returns a tuple with the AuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditId

`func (o *WebhookDeviceEventsEvent) SetAuditId(v string)`

SetAuditId sets AuditId field to given value.

### HasAuditId

`func (o *WebhookDeviceEventsEvent) HasAuditId() bool`

HasAuditId returns a boolean if a field has been set.

### GetDeviceName

`func (o *WebhookDeviceEventsEvent) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *WebhookDeviceEventsEvent) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *WebhookDeviceEventsEvent) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetDeviceType

`func (o *WebhookDeviceEventsEvent) GetDeviceType() WebhookDeviceEventsEventDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WebhookDeviceEventsEvent) GetDeviceTypeOk() (*WebhookDeviceEventsEventDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WebhookDeviceEventsEvent) SetDeviceType(v WebhookDeviceEventsEventDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetEvType

`func (o *WebhookDeviceEventsEvent) GetEvType() WebhookDeviceEventsEventEvType`

GetEvType returns the EvType field if non-nil, zero value otherwise.

### GetEvTypeOk

`func (o *WebhookDeviceEventsEvent) GetEvTypeOk() (*WebhookDeviceEventsEventEvType, bool)`

GetEvTypeOk returns a tuple with the EvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvType

`func (o *WebhookDeviceEventsEvent) SetEvType(v WebhookDeviceEventsEventEvType)`

SetEvType sets EvType field to given value.


### GetMac

`func (o *WebhookDeviceEventsEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookDeviceEventsEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookDeviceEventsEvent) SetMac(v string)`

SetMac sets Mac field to given value.


### GetOrgId

`func (o *WebhookDeviceEventsEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookDeviceEventsEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookDeviceEventsEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetReason

`func (o *WebhookDeviceEventsEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WebhookDeviceEventsEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WebhookDeviceEventsEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WebhookDeviceEventsEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookDeviceEventsEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookDeviceEventsEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookDeviceEventsEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WebhookDeviceEventsEvent) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *WebhookDeviceEventsEvent) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *WebhookDeviceEventsEvent) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *WebhookDeviceEventsEvent) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *WebhookDeviceEventsEvent) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetText

`func (o *WebhookDeviceEventsEvent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WebhookDeviceEventsEvent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WebhookDeviceEventsEvent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WebhookDeviceEventsEvent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookDeviceEventsEvent) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookDeviceEventsEvent) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookDeviceEventsEvent) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *WebhookDeviceEventsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookDeviceEventsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookDeviceEventsEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


