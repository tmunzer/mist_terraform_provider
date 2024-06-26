# WebhookClientSessionsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** | mac address of the AP the client roamed or disconnected from | 
**ApName** | **string** | user-friendly name of the AP the client roamed or disconnected from. | 
**Band** | **string** | 5GHz or 2.4GHz band | 
**Bssid** | **string** |  | 
**ClientFamily** | **string** | device family E.g. “Mac”, “iPhone”, “Apple watch” | 
**ClientManufacture** | **string** | device manufacturer E.g. “Apple” | 
**ClientModel** | **string** | device model E.g. “8+”, “XS” | 
**ClientOs** | **string** | device operating system E.g. “Mojave”, “Windows 10”, “Linux” | 
**Connect** | **int32** | time when the user connects | 
**ConnectFloat** | **float32** | floating point connect timestamp with millisecond precision | 
**Disconnect** | **int32** | time when the user disconnects | 
**DisconnectFloat** | **float32** | floating point disconnect timestamp with millisecond precision | 
**Duration** | **int32** | the duration of the roamed or complete session indicated by termination_reason field. | 
**Mac** | **string** | the client’s mac | 
**NextAp** | **string** | the AP the client has roamed to. | 
**OrgId** | **string** |  | [readonly] 
**Rssi** | **float32** | latest average RSSI before the user disconnects | 
**SiteId** | **string** |  | [readonly] 
**SiteName** | **string** |  | 
**Ssid** | **string** |  | 
**TerminationReason** | **int32** | 1 disassociate - when the client disassociates. 2 inactive - when the client is timeout. 3 roamed - when the client is roamed between APs | 
**Timestamp** | **float32** |  | 
**Version** | **float32** | schema version of this message | 
**WlanId** | **string** |  | 

## Methods

### NewWebhookClientSessionsEvent

`func NewWebhookClientSessionsEvent(ap string, apName string, band string, bssid string, clientFamily string, clientManufacture string, clientModel string, clientOs string, connect int32, connectFloat float32, disconnect int32, disconnectFloat float32, duration int32, mac string, nextAp string, orgId string, rssi float32, siteId string, siteName string, ssid string, terminationReason int32, timestamp float32, version float32, wlanId string, ) *WebhookClientSessionsEvent`

NewWebhookClientSessionsEvent instantiates a new WebhookClientSessionsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookClientSessionsEventWithDefaults

`func NewWebhookClientSessionsEventWithDefaults() *WebhookClientSessionsEvent`

NewWebhookClientSessionsEventWithDefaults instantiates a new WebhookClientSessionsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *WebhookClientSessionsEvent) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *WebhookClientSessionsEvent) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *WebhookClientSessionsEvent) SetAp(v string)`

SetAp sets Ap field to given value.


### GetApName

`func (o *WebhookClientSessionsEvent) GetApName() string`

GetApName returns the ApName field if non-nil, zero value otherwise.

### GetApNameOk

`func (o *WebhookClientSessionsEvent) GetApNameOk() (*string, bool)`

GetApNameOk returns a tuple with the ApName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApName

`func (o *WebhookClientSessionsEvent) SetApName(v string)`

SetApName sets ApName field to given value.


### GetBand

`func (o *WebhookClientSessionsEvent) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *WebhookClientSessionsEvent) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *WebhookClientSessionsEvent) SetBand(v string)`

SetBand sets Band field to given value.


### GetBssid

`func (o *WebhookClientSessionsEvent) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *WebhookClientSessionsEvent) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *WebhookClientSessionsEvent) SetBssid(v string)`

SetBssid sets Bssid field to given value.


### GetClientFamily

`func (o *WebhookClientSessionsEvent) GetClientFamily() string`

GetClientFamily returns the ClientFamily field if non-nil, zero value otherwise.

### GetClientFamilyOk

`func (o *WebhookClientSessionsEvent) GetClientFamilyOk() (*string, bool)`

GetClientFamilyOk returns a tuple with the ClientFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFamily

`func (o *WebhookClientSessionsEvent) SetClientFamily(v string)`

SetClientFamily sets ClientFamily field to given value.


### GetClientManufacture

`func (o *WebhookClientSessionsEvent) GetClientManufacture() string`

GetClientManufacture returns the ClientManufacture field if non-nil, zero value otherwise.

### GetClientManufactureOk

`func (o *WebhookClientSessionsEvent) GetClientManufactureOk() (*string, bool)`

GetClientManufactureOk returns a tuple with the ClientManufacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientManufacture

`func (o *WebhookClientSessionsEvent) SetClientManufacture(v string)`

SetClientManufacture sets ClientManufacture field to given value.


### GetClientModel

`func (o *WebhookClientSessionsEvent) GetClientModel() string`

GetClientModel returns the ClientModel field if non-nil, zero value otherwise.

### GetClientModelOk

`func (o *WebhookClientSessionsEvent) GetClientModelOk() (*string, bool)`

GetClientModelOk returns a tuple with the ClientModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientModel

`func (o *WebhookClientSessionsEvent) SetClientModel(v string)`

SetClientModel sets ClientModel field to given value.


### GetClientOs

`func (o *WebhookClientSessionsEvent) GetClientOs() string`

GetClientOs returns the ClientOs field if non-nil, zero value otherwise.

### GetClientOsOk

`func (o *WebhookClientSessionsEvent) GetClientOsOk() (*string, bool)`

GetClientOsOk returns a tuple with the ClientOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOs

`func (o *WebhookClientSessionsEvent) SetClientOs(v string)`

SetClientOs sets ClientOs field to given value.


### GetConnect

`func (o *WebhookClientSessionsEvent) GetConnect() int32`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *WebhookClientSessionsEvent) GetConnectOk() (*int32, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *WebhookClientSessionsEvent) SetConnect(v int32)`

SetConnect sets Connect field to given value.


### GetConnectFloat

`func (o *WebhookClientSessionsEvent) GetConnectFloat() float32`

GetConnectFloat returns the ConnectFloat field if non-nil, zero value otherwise.

### GetConnectFloatOk

`func (o *WebhookClientSessionsEvent) GetConnectFloatOk() (*float32, bool)`

GetConnectFloatOk returns a tuple with the ConnectFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectFloat

`func (o *WebhookClientSessionsEvent) SetConnectFloat(v float32)`

SetConnectFloat sets ConnectFloat field to given value.


### GetDisconnect

`func (o *WebhookClientSessionsEvent) GetDisconnect() int32`

GetDisconnect returns the Disconnect field if non-nil, zero value otherwise.

### GetDisconnectOk

`func (o *WebhookClientSessionsEvent) GetDisconnectOk() (*int32, bool)`

GetDisconnectOk returns a tuple with the Disconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnect

`func (o *WebhookClientSessionsEvent) SetDisconnect(v int32)`

SetDisconnect sets Disconnect field to given value.


### GetDisconnectFloat

`func (o *WebhookClientSessionsEvent) GetDisconnectFloat() float32`

GetDisconnectFloat returns the DisconnectFloat field if non-nil, zero value otherwise.

### GetDisconnectFloatOk

`func (o *WebhookClientSessionsEvent) GetDisconnectFloatOk() (*float32, bool)`

GetDisconnectFloatOk returns a tuple with the DisconnectFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectFloat

`func (o *WebhookClientSessionsEvent) SetDisconnectFloat(v float32)`

SetDisconnectFloat sets DisconnectFloat field to given value.


### GetDuration

`func (o *WebhookClientSessionsEvent) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WebhookClientSessionsEvent) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WebhookClientSessionsEvent) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetMac

`func (o *WebhookClientSessionsEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookClientSessionsEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookClientSessionsEvent) SetMac(v string)`

SetMac sets Mac field to given value.


### GetNextAp

`func (o *WebhookClientSessionsEvent) GetNextAp() string`

GetNextAp returns the NextAp field if non-nil, zero value otherwise.

### GetNextApOk

`func (o *WebhookClientSessionsEvent) GetNextApOk() (*string, bool)`

GetNextApOk returns a tuple with the NextAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAp

`func (o *WebhookClientSessionsEvent) SetNextAp(v string)`

SetNextAp sets NextAp field to given value.


### GetOrgId

`func (o *WebhookClientSessionsEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookClientSessionsEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookClientSessionsEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRssi

`func (o *WebhookClientSessionsEvent) GetRssi() float32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *WebhookClientSessionsEvent) GetRssiOk() (*float32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *WebhookClientSessionsEvent) SetRssi(v float32)`

SetRssi sets Rssi field to given value.


### GetSiteId

`func (o *WebhookClientSessionsEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookClientSessionsEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookClientSessionsEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSiteName

`func (o *WebhookClientSessionsEvent) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *WebhookClientSessionsEvent) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *WebhookClientSessionsEvent) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### GetSsid

`func (o *WebhookClientSessionsEvent) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WebhookClientSessionsEvent) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WebhookClientSessionsEvent) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetTerminationReason

`func (o *WebhookClientSessionsEvent) GetTerminationReason() int32`

GetTerminationReason returns the TerminationReason field if non-nil, zero value otherwise.

### GetTerminationReasonOk

`func (o *WebhookClientSessionsEvent) GetTerminationReasonOk() (*int32, bool)`

GetTerminationReasonOk returns a tuple with the TerminationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationReason

`func (o *WebhookClientSessionsEvent) SetTerminationReason(v int32)`

SetTerminationReason sets TerminationReason field to given value.


### GetTimestamp

`func (o *WebhookClientSessionsEvent) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookClientSessionsEvent) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookClientSessionsEvent) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetVersion

`func (o *WebhookClientSessionsEvent) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebhookClientSessionsEvent) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebhookClientSessionsEvent) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetWlanId

`func (o *WebhookClientSessionsEvent) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *WebhookClientSessionsEvent) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *WebhookClientSessionsEvent) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


