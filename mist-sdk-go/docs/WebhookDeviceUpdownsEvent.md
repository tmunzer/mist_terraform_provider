# WebhookDeviceUpdownsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** |  | [readonly] 
**ApName** | **string** |  | [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**OrgId** | **string** |  | [readonly] 
**SiteId** | **string** |  | [readonly] 
**SiteName** | **string** |  | [readonly] 
**Timestamp** | **float32** |  | 
**Type** | **string** |  | [readonly] 

## Methods

### NewWebhookDeviceUpdownsEvent

`func NewWebhookDeviceUpdownsEvent(ap string, apName string, orgId string, siteId string, siteName string, timestamp float32, type_ string, ) *WebhookDeviceUpdownsEvent`

NewWebhookDeviceUpdownsEvent instantiates a new WebhookDeviceUpdownsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeviceUpdownsEventWithDefaults

`func NewWebhookDeviceUpdownsEventWithDefaults() *WebhookDeviceUpdownsEvent`

NewWebhookDeviceUpdownsEventWithDefaults instantiates a new WebhookDeviceUpdownsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *WebhookDeviceUpdownsEvent) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *WebhookDeviceUpdownsEvent) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *WebhookDeviceUpdownsEvent) SetAp(v string)`

SetAp sets Ap field to given value.


### GetApName

`func (o *WebhookDeviceUpdownsEvent) GetApName() string`

GetApName returns the ApName field if non-nil, zero value otherwise.

### GetApNameOk

`func (o *WebhookDeviceUpdownsEvent) GetApNameOk() (*string, bool)`

GetApNameOk returns a tuple with the ApName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApName

`func (o *WebhookDeviceUpdownsEvent) SetApName(v string)`

SetApName sets ApName field to given value.


### GetForSite

`func (o *WebhookDeviceUpdownsEvent) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *WebhookDeviceUpdownsEvent) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *WebhookDeviceUpdownsEvent) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *WebhookDeviceUpdownsEvent) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetOrgId

`func (o *WebhookDeviceUpdownsEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookDeviceUpdownsEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookDeviceUpdownsEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetSiteId

`func (o *WebhookDeviceUpdownsEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookDeviceUpdownsEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookDeviceUpdownsEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSiteName

`func (o *WebhookDeviceUpdownsEvent) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *WebhookDeviceUpdownsEvent) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *WebhookDeviceUpdownsEvent) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### GetTimestamp

`func (o *WebhookDeviceUpdownsEvent) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookDeviceUpdownsEvent) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookDeviceUpdownsEvent) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *WebhookDeviceUpdownsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookDeviceUpdownsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookDeviceUpdownsEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


