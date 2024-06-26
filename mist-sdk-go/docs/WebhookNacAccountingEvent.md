# WebhookNacAccountingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to **string** | mac address of the AP the client roamed or disconnected from | [optional] 
**AuthType** | Pointer to **string** | radius authentication type | [optional] 
**Bssid** | Pointer to **string** | it’s the MAC physical address of the access point | [optional] 
**ClientIp** | Pointer to **string** | IP Address of client | [optional] 
**ClientType** | Pointer to **string** | client type E.g. “wired”, “wireless”, “vty” | [optional] 
**Mac** | Pointer to **string** | the client’s mac | [optional] 
**NasVendor** | Pointer to **string** | NAS Device vendor name E.g. “Juniper”, “Cisco” | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**RxPkts** | Pointer to **int32** | number of packets received | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Ssid** | Pointer to **string** | ESSID | [optional] 
**Timestamp** | Pointer to **float32** | sampling time (in epoch) | [optional] 
**TxPkts** | Pointer to **int32** | number of packets sent | [optional] 
**Type** | Pointer to **string** | type of event. E.g. “ACCOUNTING_START”, “ACCOUNTING_UPDATE”, “ACCOUNTING_STOP” | [optional] 
**Username** | Pointer to **string** | username authenticated with | [optional] 

## Methods

### NewWebhookNacAccountingEvent

`func NewWebhookNacAccountingEvent() *WebhookNacAccountingEvent`

NewWebhookNacAccountingEvent instantiates a new WebhookNacAccountingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookNacAccountingEventWithDefaults

`func NewWebhookNacAccountingEventWithDefaults() *WebhookNacAccountingEvent`

NewWebhookNacAccountingEventWithDefaults instantiates a new WebhookNacAccountingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *WebhookNacAccountingEvent) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *WebhookNacAccountingEvent) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *WebhookNacAccountingEvent) SetAp(v string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *WebhookNacAccountingEvent) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetAuthType

`func (o *WebhookNacAccountingEvent) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WebhookNacAccountingEvent) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WebhookNacAccountingEvent) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *WebhookNacAccountingEvent) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetBssid

`func (o *WebhookNacAccountingEvent) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *WebhookNacAccountingEvent) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *WebhookNacAccountingEvent) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *WebhookNacAccountingEvent) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetClientIp

`func (o *WebhookNacAccountingEvent) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *WebhookNacAccountingEvent) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *WebhookNacAccountingEvent) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *WebhookNacAccountingEvent) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientType

`func (o *WebhookNacAccountingEvent) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *WebhookNacAccountingEvent) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *WebhookNacAccountingEvent) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *WebhookNacAccountingEvent) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetMac

`func (o *WebhookNacAccountingEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookNacAccountingEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookNacAccountingEvent) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *WebhookNacAccountingEvent) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNasVendor

`func (o *WebhookNacAccountingEvent) GetNasVendor() string`

GetNasVendor returns the NasVendor field if non-nil, zero value otherwise.

### GetNasVendorOk

`func (o *WebhookNacAccountingEvent) GetNasVendorOk() (*string, bool)`

GetNasVendorOk returns a tuple with the NasVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasVendor

`func (o *WebhookNacAccountingEvent) SetNasVendor(v string)`

SetNasVendor sets NasVendor field to given value.

### HasNasVendor

`func (o *WebhookNacAccountingEvent) HasNasVendor() bool`

HasNasVendor returns a boolean if a field has been set.

### GetOrgId

`func (o *WebhookNacAccountingEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookNacAccountingEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookNacAccountingEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WebhookNacAccountingEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRxPkts

`func (o *WebhookNacAccountingEvent) GetRxPkts() int32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *WebhookNacAccountingEvent) GetRxPktsOk() (*int32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *WebhookNacAccountingEvent) SetRxPkts(v int32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *WebhookNacAccountingEvent) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookNacAccountingEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookNacAccountingEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookNacAccountingEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WebhookNacAccountingEvent) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSsid

`func (o *WebhookNacAccountingEvent) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WebhookNacAccountingEvent) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WebhookNacAccountingEvent) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *WebhookNacAccountingEvent) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookNacAccountingEvent) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookNacAccountingEvent) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookNacAccountingEvent) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebhookNacAccountingEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTxPkts

`func (o *WebhookNacAccountingEvent) GetTxPkts() int32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *WebhookNacAccountingEvent) GetTxPktsOk() (*int32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *WebhookNacAccountingEvent) SetTxPkts(v int32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *WebhookNacAccountingEvent) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetType

`func (o *WebhookNacAccountingEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookNacAccountingEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookNacAccountingEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WebhookNacAccountingEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *WebhookNacAccountingEvent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WebhookNacAccountingEvent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WebhookNacAccountingEvent) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *WebhookNacAccountingEvent) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


