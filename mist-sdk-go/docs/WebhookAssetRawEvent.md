# WebhookAssetRawEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | asset id | 
**Beam** | **int32** | antenna index, from 1-8, clock-wise starting from the LED | 
**DeviceId** | **string** | device where the asset reading is from | 
**IbeaconMajor** | Pointer to **int32** | iBeacon major | [optional] 
**IbeaconMinor** | Pointer to **int32** | iBeacon minor | [optional] 
**IbeaconUuid** | Pointer to **string** | iBeacon UUID | [optional] 
**Mac** | **string** | MAC of the beacon | 
**MapId** | **string** | map id | 
**MfgCompanyId** | **float32** | optional, BLE manufacturing company ID | 
**MfgData** | **string** | optional, BLE manufacturing data in hex byte-string format (ie: “112233AABBCC”) | 
**Rssi** | **float32** | signal strength | 
**ServiceDataData** | Pointer to **string** | optional, data from service data | [optional] 
**ServiceDataLastRxTime** | Pointer to **int32** | optional, last data transmit time from service data | [optional] 
**ServiceDataRxCnt** | Pointer to **int32** | optional, data transmit count from service data | [optional] 
**ServiceDataUuid** | Pointer to **string** | optional, UUID from service data | [optional] 
**ServicePackets** | Pointer to [**[]ServicePacket**](ServicePacket.md) | list of service data packets heard from the asset/ beacon | [optional] 
**SiteId** | **string** |  | [readonly] 
**Timestamp** | **float32** |  | 

## Methods

### NewWebhookAssetRawEvent

`func NewWebhookAssetRawEvent(assetId string, beam int32, deviceId string, mac string, mapId string, mfgCompanyId float32, mfgData string, rssi float32, siteId string, timestamp float32, ) *WebhookAssetRawEvent`

NewWebhookAssetRawEvent instantiates a new WebhookAssetRawEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAssetRawEventWithDefaults

`func NewWebhookAssetRawEventWithDefaults() *WebhookAssetRawEvent`

NewWebhookAssetRawEventWithDefaults instantiates a new WebhookAssetRawEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *WebhookAssetRawEvent) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *WebhookAssetRawEvent) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *WebhookAssetRawEvent) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetBeam

`func (o *WebhookAssetRawEvent) GetBeam() int32`

GetBeam returns the Beam field if non-nil, zero value otherwise.

### GetBeamOk

`func (o *WebhookAssetRawEvent) GetBeamOk() (*int32, bool)`

GetBeamOk returns a tuple with the Beam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeam

`func (o *WebhookAssetRawEvent) SetBeam(v int32)`

SetBeam sets Beam field to given value.


### GetDeviceId

`func (o *WebhookAssetRawEvent) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *WebhookAssetRawEvent) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *WebhookAssetRawEvent) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetIbeaconMajor

`func (o *WebhookAssetRawEvent) GetIbeaconMajor() int32`

GetIbeaconMajor returns the IbeaconMajor field if non-nil, zero value otherwise.

### GetIbeaconMajorOk

`func (o *WebhookAssetRawEvent) GetIbeaconMajorOk() (*int32, bool)`

GetIbeaconMajorOk returns a tuple with the IbeaconMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMajor

`func (o *WebhookAssetRawEvent) SetIbeaconMajor(v int32)`

SetIbeaconMajor sets IbeaconMajor field to given value.

### HasIbeaconMajor

`func (o *WebhookAssetRawEvent) HasIbeaconMajor() bool`

HasIbeaconMajor returns a boolean if a field has been set.

### GetIbeaconMinor

`func (o *WebhookAssetRawEvent) GetIbeaconMinor() int32`

GetIbeaconMinor returns the IbeaconMinor field if non-nil, zero value otherwise.

### GetIbeaconMinorOk

`func (o *WebhookAssetRawEvent) GetIbeaconMinorOk() (*int32, bool)`

GetIbeaconMinorOk returns a tuple with the IbeaconMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMinor

`func (o *WebhookAssetRawEvent) SetIbeaconMinor(v int32)`

SetIbeaconMinor sets IbeaconMinor field to given value.

### HasIbeaconMinor

`func (o *WebhookAssetRawEvent) HasIbeaconMinor() bool`

HasIbeaconMinor returns a boolean if a field has been set.

### GetIbeaconUuid

`func (o *WebhookAssetRawEvent) GetIbeaconUuid() string`

GetIbeaconUuid returns the IbeaconUuid field if non-nil, zero value otherwise.

### GetIbeaconUuidOk

`func (o *WebhookAssetRawEvent) GetIbeaconUuidOk() (*string, bool)`

GetIbeaconUuidOk returns a tuple with the IbeaconUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconUuid

`func (o *WebhookAssetRawEvent) SetIbeaconUuid(v string)`

SetIbeaconUuid sets IbeaconUuid field to given value.

### HasIbeaconUuid

`func (o *WebhookAssetRawEvent) HasIbeaconUuid() bool`

HasIbeaconUuid returns a boolean if a field has been set.

### GetMac

`func (o *WebhookAssetRawEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookAssetRawEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookAssetRawEvent) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMapId

`func (o *WebhookAssetRawEvent) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *WebhookAssetRawEvent) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *WebhookAssetRawEvent) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetMfgCompanyId

`func (o *WebhookAssetRawEvent) GetMfgCompanyId() float32`

GetMfgCompanyId returns the MfgCompanyId field if non-nil, zero value otherwise.

### GetMfgCompanyIdOk

`func (o *WebhookAssetRawEvent) GetMfgCompanyIdOk() (*float32, bool)`

GetMfgCompanyIdOk returns a tuple with the MfgCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgCompanyId

`func (o *WebhookAssetRawEvent) SetMfgCompanyId(v float32)`

SetMfgCompanyId sets MfgCompanyId field to given value.


### GetMfgData

`func (o *WebhookAssetRawEvent) GetMfgData() string`

GetMfgData returns the MfgData field if non-nil, zero value otherwise.

### GetMfgDataOk

`func (o *WebhookAssetRawEvent) GetMfgDataOk() (*string, bool)`

GetMfgDataOk returns a tuple with the MfgData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgData

`func (o *WebhookAssetRawEvent) SetMfgData(v string)`

SetMfgData sets MfgData field to given value.


### GetRssi

`func (o *WebhookAssetRawEvent) GetRssi() float32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *WebhookAssetRawEvent) GetRssiOk() (*float32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *WebhookAssetRawEvent) SetRssi(v float32)`

SetRssi sets Rssi field to given value.


### GetServiceDataData

`func (o *WebhookAssetRawEvent) GetServiceDataData() string`

GetServiceDataData returns the ServiceDataData field if non-nil, zero value otherwise.

### GetServiceDataDataOk

`func (o *WebhookAssetRawEvent) GetServiceDataDataOk() (*string, bool)`

GetServiceDataDataOk returns a tuple with the ServiceDataData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDataData

`func (o *WebhookAssetRawEvent) SetServiceDataData(v string)`

SetServiceDataData sets ServiceDataData field to given value.

### HasServiceDataData

`func (o *WebhookAssetRawEvent) HasServiceDataData() bool`

HasServiceDataData returns a boolean if a field has been set.

### GetServiceDataLastRxTime

`func (o *WebhookAssetRawEvent) GetServiceDataLastRxTime() int32`

GetServiceDataLastRxTime returns the ServiceDataLastRxTime field if non-nil, zero value otherwise.

### GetServiceDataLastRxTimeOk

`func (o *WebhookAssetRawEvent) GetServiceDataLastRxTimeOk() (*int32, bool)`

GetServiceDataLastRxTimeOk returns a tuple with the ServiceDataLastRxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDataLastRxTime

`func (o *WebhookAssetRawEvent) SetServiceDataLastRxTime(v int32)`

SetServiceDataLastRxTime sets ServiceDataLastRxTime field to given value.

### HasServiceDataLastRxTime

`func (o *WebhookAssetRawEvent) HasServiceDataLastRxTime() bool`

HasServiceDataLastRxTime returns a boolean if a field has been set.

### GetServiceDataRxCnt

`func (o *WebhookAssetRawEvent) GetServiceDataRxCnt() int32`

GetServiceDataRxCnt returns the ServiceDataRxCnt field if non-nil, zero value otherwise.

### GetServiceDataRxCntOk

`func (o *WebhookAssetRawEvent) GetServiceDataRxCntOk() (*int32, bool)`

GetServiceDataRxCntOk returns a tuple with the ServiceDataRxCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDataRxCnt

`func (o *WebhookAssetRawEvent) SetServiceDataRxCnt(v int32)`

SetServiceDataRxCnt sets ServiceDataRxCnt field to given value.

### HasServiceDataRxCnt

`func (o *WebhookAssetRawEvent) HasServiceDataRxCnt() bool`

HasServiceDataRxCnt returns a boolean if a field has been set.

### GetServiceDataUuid

`func (o *WebhookAssetRawEvent) GetServiceDataUuid() string`

GetServiceDataUuid returns the ServiceDataUuid field if non-nil, zero value otherwise.

### GetServiceDataUuidOk

`func (o *WebhookAssetRawEvent) GetServiceDataUuidOk() (*string, bool)`

GetServiceDataUuidOk returns a tuple with the ServiceDataUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDataUuid

`func (o *WebhookAssetRawEvent) SetServiceDataUuid(v string)`

SetServiceDataUuid sets ServiceDataUuid field to given value.

### HasServiceDataUuid

`func (o *WebhookAssetRawEvent) HasServiceDataUuid() bool`

HasServiceDataUuid returns a boolean if a field has been set.

### GetServicePackets

`func (o *WebhookAssetRawEvent) GetServicePackets() []ServicePacket`

GetServicePackets returns the ServicePackets field if non-nil, zero value otherwise.

### GetServicePacketsOk

`func (o *WebhookAssetRawEvent) GetServicePacketsOk() (*[]ServicePacket, bool)`

GetServicePacketsOk returns a tuple with the ServicePackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePackets

`func (o *WebhookAssetRawEvent) SetServicePackets(v []ServicePacket)`

SetServicePackets sets ServicePackets field to given value.

### HasServicePackets

`func (o *WebhookAssetRawEvent) HasServicePackets() bool`

HasServicePackets returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookAssetRawEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookAssetRawEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookAssetRawEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetTimestamp

`func (o *WebhookAssetRawEvent) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookAssetRawEvent) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookAssetRawEvent) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


