# WebhookClientLatencyEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgAuth** | Pointer to **float32** |  | [optional] 
**AvgDhcp** | Pointer to **float32** |  | [optional] 
**AvgDns** | Pointer to **float32** |  | [optional] 
**MaxAuth** | Pointer to **float32** |  | [optional] 
**MaxDhcp** | Pointer to **float32** |  | [optional] 
**MaxDns** | Pointer to **float32** |  | [optional] 
**MinAuth** | Pointer to **float32** |  | [optional] 
**MinDhcp** | Pointer to **float32** |  | [optional] 
**MinDns** | Pointer to **float32** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **int32** |  | [optional] 

## Methods

### NewWebhookClientLatencyEvent

`func NewWebhookClientLatencyEvent() *WebhookClientLatencyEvent`

NewWebhookClientLatencyEvent instantiates a new WebhookClientLatencyEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookClientLatencyEventWithDefaults

`func NewWebhookClientLatencyEventWithDefaults() *WebhookClientLatencyEvent`

NewWebhookClientLatencyEventWithDefaults instantiates a new WebhookClientLatencyEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgAuth

`func (o *WebhookClientLatencyEvent) GetAvgAuth() float32`

GetAvgAuth returns the AvgAuth field if non-nil, zero value otherwise.

### GetAvgAuthOk

`func (o *WebhookClientLatencyEvent) GetAvgAuthOk() (*float32, bool)`

GetAvgAuthOk returns a tuple with the AvgAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgAuth

`func (o *WebhookClientLatencyEvent) SetAvgAuth(v float32)`

SetAvgAuth sets AvgAuth field to given value.

### HasAvgAuth

`func (o *WebhookClientLatencyEvent) HasAvgAuth() bool`

HasAvgAuth returns a boolean if a field has been set.

### GetAvgDhcp

`func (o *WebhookClientLatencyEvent) GetAvgDhcp() float32`

GetAvgDhcp returns the AvgDhcp field if non-nil, zero value otherwise.

### GetAvgDhcpOk

`func (o *WebhookClientLatencyEvent) GetAvgDhcpOk() (*float32, bool)`

GetAvgDhcpOk returns a tuple with the AvgDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDhcp

`func (o *WebhookClientLatencyEvent) SetAvgDhcp(v float32)`

SetAvgDhcp sets AvgDhcp field to given value.

### HasAvgDhcp

`func (o *WebhookClientLatencyEvent) HasAvgDhcp() bool`

HasAvgDhcp returns a boolean if a field has been set.

### GetAvgDns

`func (o *WebhookClientLatencyEvent) GetAvgDns() float32`

GetAvgDns returns the AvgDns field if non-nil, zero value otherwise.

### GetAvgDnsOk

`func (o *WebhookClientLatencyEvent) GetAvgDnsOk() (*float32, bool)`

GetAvgDnsOk returns a tuple with the AvgDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDns

`func (o *WebhookClientLatencyEvent) SetAvgDns(v float32)`

SetAvgDns sets AvgDns field to given value.

### HasAvgDns

`func (o *WebhookClientLatencyEvent) HasAvgDns() bool`

HasAvgDns returns a boolean if a field has been set.

### GetMaxAuth

`func (o *WebhookClientLatencyEvent) GetMaxAuth() float32`

GetMaxAuth returns the MaxAuth field if non-nil, zero value otherwise.

### GetMaxAuthOk

`func (o *WebhookClientLatencyEvent) GetMaxAuthOk() (*float32, bool)`

GetMaxAuthOk returns a tuple with the MaxAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAuth

`func (o *WebhookClientLatencyEvent) SetMaxAuth(v float32)`

SetMaxAuth sets MaxAuth field to given value.

### HasMaxAuth

`func (o *WebhookClientLatencyEvent) HasMaxAuth() bool`

HasMaxAuth returns a boolean if a field has been set.

### GetMaxDhcp

`func (o *WebhookClientLatencyEvent) GetMaxDhcp() float32`

GetMaxDhcp returns the MaxDhcp field if non-nil, zero value otherwise.

### GetMaxDhcpOk

`func (o *WebhookClientLatencyEvent) GetMaxDhcpOk() (*float32, bool)`

GetMaxDhcpOk returns a tuple with the MaxDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDhcp

`func (o *WebhookClientLatencyEvent) SetMaxDhcp(v float32)`

SetMaxDhcp sets MaxDhcp field to given value.

### HasMaxDhcp

`func (o *WebhookClientLatencyEvent) HasMaxDhcp() bool`

HasMaxDhcp returns a boolean if a field has been set.

### GetMaxDns

`func (o *WebhookClientLatencyEvent) GetMaxDns() float32`

GetMaxDns returns the MaxDns field if non-nil, zero value otherwise.

### GetMaxDnsOk

`func (o *WebhookClientLatencyEvent) GetMaxDnsOk() (*float32, bool)`

GetMaxDnsOk returns a tuple with the MaxDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDns

`func (o *WebhookClientLatencyEvent) SetMaxDns(v float32)`

SetMaxDns sets MaxDns field to given value.

### HasMaxDns

`func (o *WebhookClientLatencyEvent) HasMaxDns() bool`

HasMaxDns returns a boolean if a field has been set.

### GetMinAuth

`func (o *WebhookClientLatencyEvent) GetMinAuth() float32`

GetMinAuth returns the MinAuth field if non-nil, zero value otherwise.

### GetMinAuthOk

`func (o *WebhookClientLatencyEvent) GetMinAuthOk() (*float32, bool)`

GetMinAuthOk returns a tuple with the MinAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAuth

`func (o *WebhookClientLatencyEvent) SetMinAuth(v float32)`

SetMinAuth sets MinAuth field to given value.

### HasMinAuth

`func (o *WebhookClientLatencyEvent) HasMinAuth() bool`

HasMinAuth returns a boolean if a field has been set.

### GetMinDhcp

`func (o *WebhookClientLatencyEvent) GetMinDhcp() float32`

GetMinDhcp returns the MinDhcp field if non-nil, zero value otherwise.

### GetMinDhcpOk

`func (o *WebhookClientLatencyEvent) GetMinDhcpOk() (*float32, bool)`

GetMinDhcpOk returns a tuple with the MinDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDhcp

`func (o *WebhookClientLatencyEvent) SetMinDhcp(v float32)`

SetMinDhcp sets MinDhcp field to given value.

### HasMinDhcp

`func (o *WebhookClientLatencyEvent) HasMinDhcp() bool`

HasMinDhcp returns a boolean if a field has been set.

### GetMinDns

`func (o *WebhookClientLatencyEvent) GetMinDns() float32`

GetMinDns returns the MinDns field if non-nil, zero value otherwise.

### GetMinDnsOk

`func (o *WebhookClientLatencyEvent) GetMinDnsOk() (*float32, bool)`

GetMinDnsOk returns a tuple with the MinDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDns

`func (o *WebhookClientLatencyEvent) SetMinDns(v float32)`

SetMinDns sets MinDns field to given value.

### HasMinDns

`func (o *WebhookClientLatencyEvent) HasMinDns() bool`

HasMinDns returns a boolean if a field has been set.

### GetOrgId

`func (o *WebhookClientLatencyEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookClientLatencyEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookClientLatencyEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WebhookClientLatencyEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookClientLatencyEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookClientLatencyEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookClientLatencyEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WebhookClientLatencyEvent) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookClientLatencyEvent) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookClientLatencyEvent) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookClientLatencyEvent) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebhookClientLatencyEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


