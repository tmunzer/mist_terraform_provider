# WebhookClientJoinEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** | mac address of the AP the client connected to | 
**ApName** | **string** | user-friendly name of the AP the client connected to. | 
**Band** | **string** | 5GHz or 2.4GHz band | 
**Bssid** | **string** |  | 
**Connect** | **int32** | time when the user connects | 
**ConnectFloat** | **float32** | floating point connect timestamp with millisecond precision | 
**Mac** | **string** | the client’s mac | 
**OrgId** | **string** |  | [readonly] 
**Rssi** | **float32** | RSSI when the client associated | 
**SiteId** | **string** |  | [readonly] 
**SiteName** | **string** |  | 
**Ssid** | **string** | ESSID | 
**Timestamp** | **float32** |  | 
**Version** | **float32** | schema version of this message | 
**WlanId** | **string** |  | 

## Methods

### NewWebhookClientJoinEvent

`func NewWebhookClientJoinEvent(ap string, apName string, band string, bssid string, connect int32, connectFloat float32, mac string, orgId string, rssi float32, siteId string, siteName string, ssid string, timestamp float32, version float32, wlanId string, ) *WebhookClientJoinEvent`

NewWebhookClientJoinEvent instantiates a new WebhookClientJoinEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookClientJoinEventWithDefaults

`func NewWebhookClientJoinEventWithDefaults() *WebhookClientJoinEvent`

NewWebhookClientJoinEventWithDefaults instantiates a new WebhookClientJoinEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *WebhookClientJoinEvent) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *WebhookClientJoinEvent) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *WebhookClientJoinEvent) SetAp(v string)`

SetAp sets Ap field to given value.


### GetApName

`func (o *WebhookClientJoinEvent) GetApName() string`

GetApName returns the ApName field if non-nil, zero value otherwise.

### GetApNameOk

`func (o *WebhookClientJoinEvent) GetApNameOk() (*string, bool)`

GetApNameOk returns a tuple with the ApName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApName

`func (o *WebhookClientJoinEvent) SetApName(v string)`

SetApName sets ApName field to given value.


### GetBand

`func (o *WebhookClientJoinEvent) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *WebhookClientJoinEvent) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *WebhookClientJoinEvent) SetBand(v string)`

SetBand sets Band field to given value.


### GetBssid

`func (o *WebhookClientJoinEvent) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *WebhookClientJoinEvent) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *WebhookClientJoinEvent) SetBssid(v string)`

SetBssid sets Bssid field to given value.


### GetConnect

`func (o *WebhookClientJoinEvent) GetConnect() int32`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *WebhookClientJoinEvent) GetConnectOk() (*int32, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *WebhookClientJoinEvent) SetConnect(v int32)`

SetConnect sets Connect field to given value.


### GetConnectFloat

`func (o *WebhookClientJoinEvent) GetConnectFloat() float32`

GetConnectFloat returns the ConnectFloat field if non-nil, zero value otherwise.

### GetConnectFloatOk

`func (o *WebhookClientJoinEvent) GetConnectFloatOk() (*float32, bool)`

GetConnectFloatOk returns a tuple with the ConnectFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectFloat

`func (o *WebhookClientJoinEvent) SetConnectFloat(v float32)`

SetConnectFloat sets ConnectFloat field to given value.


### GetMac

`func (o *WebhookClientJoinEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookClientJoinEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookClientJoinEvent) SetMac(v string)`

SetMac sets Mac field to given value.


### GetOrgId

`func (o *WebhookClientJoinEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookClientJoinEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookClientJoinEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRssi

`func (o *WebhookClientJoinEvent) GetRssi() float32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *WebhookClientJoinEvent) GetRssiOk() (*float32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *WebhookClientJoinEvent) SetRssi(v float32)`

SetRssi sets Rssi field to given value.


### GetSiteId

`func (o *WebhookClientJoinEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookClientJoinEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookClientJoinEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSiteName

`func (o *WebhookClientJoinEvent) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *WebhookClientJoinEvent) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *WebhookClientJoinEvent) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### GetSsid

`func (o *WebhookClientJoinEvent) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WebhookClientJoinEvent) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WebhookClientJoinEvent) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetTimestamp

`func (o *WebhookClientJoinEvent) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookClientJoinEvent) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookClientJoinEvent) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetVersion

`func (o *WebhookClientJoinEvent) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebhookClientJoinEvent) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebhookClientJoinEvent) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetWlanId

`func (o *WebhookClientJoinEvent) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *WebhookClientJoinEvent) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *WebhookClientJoinEvent) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


