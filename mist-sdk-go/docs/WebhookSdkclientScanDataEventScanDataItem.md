# WebhookSdkclientScanDataEventScanDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** | mac address of the AP associated with the BSSID scanned | 
**Band** | [**WebhookSdkclientScanDataEventScanDataItemBand**](WebhookSdkclientScanDataEventScanDataItemBand.md) |  | 
**Bssid** | **string** | BSSID found during client’s background scan for wifi | 
**Channel** | **int32** | channel of the band found in the scan | 
**Rssi** | **float32** | client’s RSSI relative to the BSSID scanned | 
**Ssid** | **string** | ESSID containing the BSSID scanned | 
**Timestamp** | **float32** | time the scan of the particular BSSID occurred | 

## Methods

### NewWebhookSdkclientScanDataEventScanDataItem

`func NewWebhookSdkclientScanDataEventScanDataItem(ap string, band WebhookSdkclientScanDataEventScanDataItemBand, bssid string, channel int32, rssi float32, ssid string, timestamp float32, ) *WebhookSdkclientScanDataEventScanDataItem`

NewWebhookSdkclientScanDataEventScanDataItem instantiates a new WebhookSdkclientScanDataEventScanDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSdkclientScanDataEventScanDataItemWithDefaults

`func NewWebhookSdkclientScanDataEventScanDataItemWithDefaults() *WebhookSdkclientScanDataEventScanDataItem`

NewWebhookSdkclientScanDataEventScanDataItemWithDefaults instantiates a new WebhookSdkclientScanDataEventScanDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *WebhookSdkclientScanDataEventScanDataItem) SetAp(v string)`

SetAp sets Ap field to given value.


### GetBand

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetBand() WebhookSdkclientScanDataEventScanDataItemBand`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetBandOk() (*WebhookSdkclientScanDataEventScanDataItemBand, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *WebhookSdkclientScanDataEventScanDataItem) SetBand(v WebhookSdkclientScanDataEventScanDataItemBand)`

SetBand sets Band field to given value.


### GetBssid

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *WebhookSdkclientScanDataEventScanDataItem) SetBssid(v string)`

SetBssid sets Bssid field to given value.


### GetChannel

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WebhookSdkclientScanDataEventScanDataItem) SetChannel(v int32)`

SetChannel sets Channel field to given value.


### GetRssi

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetRssi() float32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetRssiOk() (*float32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *WebhookSdkclientScanDataEventScanDataItem) SetRssi(v float32)`

SetRssi sets Rssi field to given value.


### GetSsid

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WebhookSdkclientScanDataEventScanDataItem) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetTimestamp

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookSdkclientScanDataEventScanDataItem) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookSdkclientScanDataEventScanDataItem) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


