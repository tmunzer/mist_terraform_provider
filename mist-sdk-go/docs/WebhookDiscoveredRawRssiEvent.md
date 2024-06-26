# WebhookDiscoveredRawRssiEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApLoc** | Pointer to **[]float32** | coordinates (if any) of reporting AP (updated once in 60s per client) | [optional] 
**Beam** | **int32** | antenna index, from 1-8, clock-wise starting from the LED | 
**DeviceId** | **string** | device id of the reporting AP | 
**IbeaconMajor** | Pointer to **int32** |  | [optional] 
**IbeaconMinor** | Pointer to **int32** |  | [optional] 
**IbeaconUuid** | Pointer to **string** |  | [optional] 
**IsAsset** | Pointer to **bool** |  | [optional] 
**Mac** | **string** | MAC of the asset/ beacon | 
**MapId** | **string** |  | 
**MfgCompanyId** | Pointer to **string** | BLE manufacturing company ID | [optional] 
**MfgData** | Pointer to **string** | BLE manufacturing data in hex byte-string format (ie: “112233AABBCC”) | [optional] 
**OrgId** | **string** |  | [readonly] 
**Rssi** | **float32** | signal strength | 
**ServicePackets** | Pointer to [**[]ServicePacket**](ServicePacket.md) | list of service data packets heard from the asset/ beacon | [optional] 
**SiteId** | **string** |  | [readonly] 
**Timestamp** | Pointer to **int32** |  | [optional] 

## Methods

### NewWebhookDiscoveredRawRssiEvent

`func NewWebhookDiscoveredRawRssiEvent(beam int32, deviceId string, mac string, mapId string, orgId string, rssi float32, siteId string, ) *WebhookDiscoveredRawRssiEvent`

NewWebhookDiscoveredRawRssiEvent instantiates a new WebhookDiscoveredRawRssiEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDiscoveredRawRssiEventWithDefaults

`func NewWebhookDiscoveredRawRssiEventWithDefaults() *WebhookDiscoveredRawRssiEvent`

NewWebhookDiscoveredRawRssiEventWithDefaults instantiates a new WebhookDiscoveredRawRssiEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApLoc

`func (o *WebhookDiscoveredRawRssiEvent) GetApLoc() []float32`

GetApLoc returns the ApLoc field if non-nil, zero value otherwise.

### GetApLocOk

`func (o *WebhookDiscoveredRawRssiEvent) GetApLocOk() (*[]float32, bool)`

GetApLocOk returns a tuple with the ApLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApLoc

`func (o *WebhookDiscoveredRawRssiEvent) SetApLoc(v []float32)`

SetApLoc sets ApLoc field to given value.

### HasApLoc

`func (o *WebhookDiscoveredRawRssiEvent) HasApLoc() bool`

HasApLoc returns a boolean if a field has been set.

### GetBeam

`func (o *WebhookDiscoveredRawRssiEvent) GetBeam() int32`

GetBeam returns the Beam field if non-nil, zero value otherwise.

### GetBeamOk

`func (o *WebhookDiscoveredRawRssiEvent) GetBeamOk() (*int32, bool)`

GetBeamOk returns a tuple with the Beam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeam

`func (o *WebhookDiscoveredRawRssiEvent) SetBeam(v int32)`

SetBeam sets Beam field to given value.


### GetDeviceId

`func (o *WebhookDiscoveredRawRssiEvent) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *WebhookDiscoveredRawRssiEvent) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *WebhookDiscoveredRawRssiEvent) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetIbeaconMajor

`func (o *WebhookDiscoveredRawRssiEvent) GetIbeaconMajor() int32`

GetIbeaconMajor returns the IbeaconMajor field if non-nil, zero value otherwise.

### GetIbeaconMajorOk

`func (o *WebhookDiscoveredRawRssiEvent) GetIbeaconMajorOk() (*int32, bool)`

GetIbeaconMajorOk returns a tuple with the IbeaconMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMajor

`func (o *WebhookDiscoveredRawRssiEvent) SetIbeaconMajor(v int32)`

SetIbeaconMajor sets IbeaconMajor field to given value.

### HasIbeaconMajor

`func (o *WebhookDiscoveredRawRssiEvent) HasIbeaconMajor() bool`

HasIbeaconMajor returns a boolean if a field has been set.

### GetIbeaconMinor

`func (o *WebhookDiscoveredRawRssiEvent) GetIbeaconMinor() int32`

GetIbeaconMinor returns the IbeaconMinor field if non-nil, zero value otherwise.

### GetIbeaconMinorOk

`func (o *WebhookDiscoveredRawRssiEvent) GetIbeaconMinorOk() (*int32, bool)`

GetIbeaconMinorOk returns a tuple with the IbeaconMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMinor

`func (o *WebhookDiscoveredRawRssiEvent) SetIbeaconMinor(v int32)`

SetIbeaconMinor sets IbeaconMinor field to given value.

### HasIbeaconMinor

`func (o *WebhookDiscoveredRawRssiEvent) HasIbeaconMinor() bool`

HasIbeaconMinor returns a boolean if a field has been set.

### GetIbeaconUuid

`func (o *WebhookDiscoveredRawRssiEvent) GetIbeaconUuid() string`

GetIbeaconUuid returns the IbeaconUuid field if non-nil, zero value otherwise.

### GetIbeaconUuidOk

`func (o *WebhookDiscoveredRawRssiEvent) GetIbeaconUuidOk() (*string, bool)`

GetIbeaconUuidOk returns a tuple with the IbeaconUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconUuid

`func (o *WebhookDiscoveredRawRssiEvent) SetIbeaconUuid(v string)`

SetIbeaconUuid sets IbeaconUuid field to given value.

### HasIbeaconUuid

`func (o *WebhookDiscoveredRawRssiEvent) HasIbeaconUuid() bool`

HasIbeaconUuid returns a boolean if a field has been set.

### GetIsAsset

`func (o *WebhookDiscoveredRawRssiEvent) GetIsAsset() bool`

GetIsAsset returns the IsAsset field if non-nil, zero value otherwise.

### GetIsAssetOk

`func (o *WebhookDiscoveredRawRssiEvent) GetIsAssetOk() (*bool, bool)`

GetIsAssetOk returns a tuple with the IsAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsset

`func (o *WebhookDiscoveredRawRssiEvent) SetIsAsset(v bool)`

SetIsAsset sets IsAsset field to given value.

### HasIsAsset

`func (o *WebhookDiscoveredRawRssiEvent) HasIsAsset() bool`

HasIsAsset returns a boolean if a field has been set.

### GetMac

`func (o *WebhookDiscoveredRawRssiEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookDiscoveredRawRssiEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookDiscoveredRawRssiEvent) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMapId

`func (o *WebhookDiscoveredRawRssiEvent) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *WebhookDiscoveredRawRssiEvent) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *WebhookDiscoveredRawRssiEvent) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetMfgCompanyId

`func (o *WebhookDiscoveredRawRssiEvent) GetMfgCompanyId() string`

GetMfgCompanyId returns the MfgCompanyId field if non-nil, zero value otherwise.

### GetMfgCompanyIdOk

`func (o *WebhookDiscoveredRawRssiEvent) GetMfgCompanyIdOk() (*string, bool)`

GetMfgCompanyIdOk returns a tuple with the MfgCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgCompanyId

`func (o *WebhookDiscoveredRawRssiEvent) SetMfgCompanyId(v string)`

SetMfgCompanyId sets MfgCompanyId field to given value.

### HasMfgCompanyId

`func (o *WebhookDiscoveredRawRssiEvent) HasMfgCompanyId() bool`

HasMfgCompanyId returns a boolean if a field has been set.

### GetMfgData

`func (o *WebhookDiscoveredRawRssiEvent) GetMfgData() string`

GetMfgData returns the MfgData field if non-nil, zero value otherwise.

### GetMfgDataOk

`func (o *WebhookDiscoveredRawRssiEvent) GetMfgDataOk() (*string, bool)`

GetMfgDataOk returns a tuple with the MfgData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgData

`func (o *WebhookDiscoveredRawRssiEvent) SetMfgData(v string)`

SetMfgData sets MfgData field to given value.

### HasMfgData

`func (o *WebhookDiscoveredRawRssiEvent) HasMfgData() bool`

HasMfgData returns a boolean if a field has been set.

### GetOrgId

`func (o *WebhookDiscoveredRawRssiEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookDiscoveredRawRssiEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookDiscoveredRawRssiEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRssi

`func (o *WebhookDiscoveredRawRssiEvent) GetRssi() float32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *WebhookDiscoveredRawRssiEvent) GetRssiOk() (*float32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *WebhookDiscoveredRawRssiEvent) SetRssi(v float32)`

SetRssi sets Rssi field to given value.


### GetServicePackets

`func (o *WebhookDiscoveredRawRssiEvent) GetServicePackets() []ServicePacket`

GetServicePackets returns the ServicePackets field if non-nil, zero value otherwise.

### GetServicePacketsOk

`func (o *WebhookDiscoveredRawRssiEvent) GetServicePacketsOk() (*[]ServicePacket, bool)`

GetServicePacketsOk returns a tuple with the ServicePackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePackets

`func (o *WebhookDiscoveredRawRssiEvent) SetServicePackets(v []ServicePacket)`

SetServicePackets sets ServicePackets field to given value.

### HasServicePackets

`func (o *WebhookDiscoveredRawRssiEvent) HasServicePackets() bool`

HasServicePackets returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookDiscoveredRawRssiEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookDiscoveredRawRssiEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookDiscoveredRawRssiEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetTimestamp

`func (o *WebhookDiscoveredRawRssiEvent) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookDiscoveredRawRssiEvent) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookDiscoveredRawRssiEvent) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebhookDiscoveredRawRssiEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


