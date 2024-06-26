# WebhookLocationCentrakEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapId** | Pointer to **string** | map id | [optional] 
**MfgCompanyId** | Pointer to **int32** | optional, BLE manufacturing company ID | [optional] 
**MfgData** | Pointer to **string** | optional, BLE manufacturing data in hex byte-string format (i.e. “112233AABBCC”) | [optional] 
**Timestamp** | Pointer to **int32** | timestamp of the event, epoch | [optional] 
**WifiBeaconExtendedInfo** | Pointer to [**[]WebhookLocationEventWifiBeaconExtendedInfoItems**](WebhookLocationEventWifiBeaconExtendedInfoItems.md) | optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload | [optional] 
**X** | Pointer to **float64** | x, in meter | [optional] 
**Y** | Pointer to **float64** | y, in meter | [optional] 

## Methods

### NewWebhookLocationCentrakEvent

`func NewWebhookLocationCentrakEvent() *WebhookLocationCentrakEvent`

NewWebhookLocationCentrakEvent instantiates a new WebhookLocationCentrakEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLocationCentrakEventWithDefaults

`func NewWebhookLocationCentrakEventWithDefaults() *WebhookLocationCentrakEvent`

NewWebhookLocationCentrakEventWithDefaults instantiates a new WebhookLocationCentrakEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapId

`func (o *WebhookLocationCentrakEvent) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *WebhookLocationCentrakEvent) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *WebhookLocationCentrakEvent) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *WebhookLocationCentrakEvent) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetMfgCompanyId

`func (o *WebhookLocationCentrakEvent) GetMfgCompanyId() int32`

GetMfgCompanyId returns the MfgCompanyId field if non-nil, zero value otherwise.

### GetMfgCompanyIdOk

`func (o *WebhookLocationCentrakEvent) GetMfgCompanyIdOk() (*int32, bool)`

GetMfgCompanyIdOk returns a tuple with the MfgCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgCompanyId

`func (o *WebhookLocationCentrakEvent) SetMfgCompanyId(v int32)`

SetMfgCompanyId sets MfgCompanyId field to given value.

### HasMfgCompanyId

`func (o *WebhookLocationCentrakEvent) HasMfgCompanyId() bool`

HasMfgCompanyId returns a boolean if a field has been set.

### GetMfgData

`func (o *WebhookLocationCentrakEvent) GetMfgData() string`

GetMfgData returns the MfgData field if non-nil, zero value otherwise.

### GetMfgDataOk

`func (o *WebhookLocationCentrakEvent) GetMfgDataOk() (*string, bool)`

GetMfgDataOk returns a tuple with the MfgData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgData

`func (o *WebhookLocationCentrakEvent) SetMfgData(v string)`

SetMfgData sets MfgData field to given value.

### HasMfgData

`func (o *WebhookLocationCentrakEvent) HasMfgData() bool`

HasMfgData returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookLocationCentrakEvent) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookLocationCentrakEvent) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookLocationCentrakEvent) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebhookLocationCentrakEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetWifiBeaconExtendedInfo

`func (o *WebhookLocationCentrakEvent) GetWifiBeaconExtendedInfo() []WebhookLocationEventWifiBeaconExtendedInfoItems`

GetWifiBeaconExtendedInfo returns the WifiBeaconExtendedInfo field if non-nil, zero value otherwise.

### GetWifiBeaconExtendedInfoOk

`func (o *WebhookLocationCentrakEvent) GetWifiBeaconExtendedInfoOk() (*[]WebhookLocationEventWifiBeaconExtendedInfoItems, bool)`

GetWifiBeaconExtendedInfoOk returns a tuple with the WifiBeaconExtendedInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiBeaconExtendedInfo

`func (o *WebhookLocationCentrakEvent) SetWifiBeaconExtendedInfo(v []WebhookLocationEventWifiBeaconExtendedInfoItems)`

SetWifiBeaconExtendedInfo sets WifiBeaconExtendedInfo field to given value.

### HasWifiBeaconExtendedInfo

`func (o *WebhookLocationCentrakEvent) HasWifiBeaconExtendedInfo() bool`

HasWifiBeaconExtendedInfo returns a boolean if a field has been set.

### GetX

`func (o *WebhookLocationCentrakEvent) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *WebhookLocationCentrakEvent) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *WebhookLocationCentrakEvent) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *WebhookLocationCentrakEvent) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *WebhookLocationCentrakEvent) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *WebhookLocationCentrakEvent) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *WebhookLocationCentrakEvent) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *WebhookLocationCentrakEvent) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


