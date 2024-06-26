# WebhookClientInfoEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | Hostname of client | [optional] 
**Ip** | Pointer to **string** | IP address of client | [optional] 
**Mac** | Pointer to **string** | the client’s mac | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **float32** | time at which IP address was assigned E.g. 1703003956 | [optional] 

## Methods

### NewWebhookClientInfoEvent

`func NewWebhookClientInfoEvent() *WebhookClientInfoEvent`

NewWebhookClientInfoEvent instantiates a new WebhookClientInfoEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookClientInfoEventWithDefaults

`func NewWebhookClientInfoEventWithDefaults() *WebhookClientInfoEvent`

NewWebhookClientInfoEventWithDefaults instantiates a new WebhookClientInfoEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *WebhookClientInfoEvent) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *WebhookClientInfoEvent) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *WebhookClientInfoEvent) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *WebhookClientInfoEvent) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *WebhookClientInfoEvent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *WebhookClientInfoEvent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *WebhookClientInfoEvent) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *WebhookClientInfoEvent) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *WebhookClientInfoEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookClientInfoEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookClientInfoEvent) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *WebhookClientInfoEvent) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetOrgId

`func (o *WebhookClientInfoEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookClientInfoEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookClientInfoEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WebhookClientInfoEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookClientInfoEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookClientInfoEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookClientInfoEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WebhookClientInfoEvent) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookClientInfoEvent) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookClientInfoEvent) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookClientInfoEvent) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebhookClientInfoEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


