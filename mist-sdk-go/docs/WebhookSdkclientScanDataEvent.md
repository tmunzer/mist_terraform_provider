# WebhookSdkclientScanDataEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionAp** | **string** | mac address of the AP the client is connected to | 
**ConnectionBand** | **string** | 5GHz or 2.4GHz band, of the BSSID the client is connected to | 
**ConnectionBssid** | **string** | BSSID of the AP the client is connected to | 
**ConnectionChannel** | **int32** | channel of the band the client is connected to | 
**ConnectionRssi** | **float32** | RSSI of the client’s connection to the AP/BSSID | 
**LastSeen** | **float32** | time client last seen with scan data | 
**Mac** | **string** | the client’s mac | 
**ScanData** | Pointer to [**[]WebhookSdkclientScanDataEventScanDataItem**](WebhookSdkclientScanDataEventScanDataItem.md) |  | [optional] 
**SiteId** | **string** |  | 

## Methods

### NewWebhookSdkclientScanDataEvent

`func NewWebhookSdkclientScanDataEvent(connectionAp string, connectionBand string, connectionBssid string, connectionChannel int32, connectionRssi float32, lastSeen float32, mac string, siteId string, ) *WebhookSdkclientScanDataEvent`

NewWebhookSdkclientScanDataEvent instantiates a new WebhookSdkclientScanDataEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSdkclientScanDataEventWithDefaults

`func NewWebhookSdkclientScanDataEventWithDefaults() *WebhookSdkclientScanDataEvent`

NewWebhookSdkclientScanDataEventWithDefaults instantiates a new WebhookSdkclientScanDataEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionAp

`func (o *WebhookSdkclientScanDataEvent) GetConnectionAp() string`

GetConnectionAp returns the ConnectionAp field if non-nil, zero value otherwise.

### GetConnectionApOk

`func (o *WebhookSdkclientScanDataEvent) GetConnectionApOk() (*string, bool)`

GetConnectionApOk returns a tuple with the ConnectionAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionAp

`func (o *WebhookSdkclientScanDataEvent) SetConnectionAp(v string)`

SetConnectionAp sets ConnectionAp field to given value.


### GetConnectionBand

`func (o *WebhookSdkclientScanDataEvent) GetConnectionBand() string`

GetConnectionBand returns the ConnectionBand field if non-nil, zero value otherwise.

### GetConnectionBandOk

`func (o *WebhookSdkclientScanDataEvent) GetConnectionBandOk() (*string, bool)`

GetConnectionBandOk returns a tuple with the ConnectionBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionBand

`func (o *WebhookSdkclientScanDataEvent) SetConnectionBand(v string)`

SetConnectionBand sets ConnectionBand field to given value.


### GetConnectionBssid

`func (o *WebhookSdkclientScanDataEvent) GetConnectionBssid() string`

GetConnectionBssid returns the ConnectionBssid field if non-nil, zero value otherwise.

### GetConnectionBssidOk

`func (o *WebhookSdkclientScanDataEvent) GetConnectionBssidOk() (*string, bool)`

GetConnectionBssidOk returns a tuple with the ConnectionBssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionBssid

`func (o *WebhookSdkclientScanDataEvent) SetConnectionBssid(v string)`

SetConnectionBssid sets ConnectionBssid field to given value.


### GetConnectionChannel

`func (o *WebhookSdkclientScanDataEvent) GetConnectionChannel() int32`

GetConnectionChannel returns the ConnectionChannel field if non-nil, zero value otherwise.

### GetConnectionChannelOk

`func (o *WebhookSdkclientScanDataEvent) GetConnectionChannelOk() (*int32, bool)`

GetConnectionChannelOk returns a tuple with the ConnectionChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionChannel

`func (o *WebhookSdkclientScanDataEvent) SetConnectionChannel(v int32)`

SetConnectionChannel sets ConnectionChannel field to given value.


### GetConnectionRssi

`func (o *WebhookSdkclientScanDataEvent) GetConnectionRssi() float32`

GetConnectionRssi returns the ConnectionRssi field if non-nil, zero value otherwise.

### GetConnectionRssiOk

`func (o *WebhookSdkclientScanDataEvent) GetConnectionRssiOk() (*float32, bool)`

GetConnectionRssiOk returns a tuple with the ConnectionRssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionRssi

`func (o *WebhookSdkclientScanDataEvent) SetConnectionRssi(v float32)`

SetConnectionRssi sets ConnectionRssi field to given value.


### GetLastSeen

`func (o *WebhookSdkclientScanDataEvent) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *WebhookSdkclientScanDataEvent) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *WebhookSdkclientScanDataEvent) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.


### GetMac

`func (o *WebhookSdkclientScanDataEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookSdkclientScanDataEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookSdkclientScanDataEvent) SetMac(v string)`

SetMac sets Mac field to given value.


### GetScanData

`func (o *WebhookSdkclientScanDataEvent) GetScanData() []WebhookSdkclientScanDataEventScanDataItem`

GetScanData returns the ScanData field if non-nil, zero value otherwise.

### GetScanDataOk

`func (o *WebhookSdkclientScanDataEvent) GetScanDataOk() (*[]WebhookSdkclientScanDataEventScanDataItem, bool)`

GetScanDataOk returns a tuple with the ScanData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanData

`func (o *WebhookSdkclientScanDataEvent) SetScanData(v []WebhookSdkclientScanDataEventScanDataItem)`

SetScanData sets ScanData field to given value.

### HasScanData

`func (o *WebhookSdkclientScanDataEvent) HasScanData() bool`

HasScanData returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookSdkclientScanDataEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookSdkclientScanDataEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookSdkclientScanDataEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


