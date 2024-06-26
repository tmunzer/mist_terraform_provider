# WebhookLocationAssetEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatteryVoltage** | Pointer to **int32** |  | [optional] 
**EddystoneUidInstance** | Pointer to **string** |  | [optional] 
**EddystoneUidNamespace** | Pointer to **string** |  | [optional] 
**EddystoneUrlUrl** | Pointer to **string** |  | [optional] 
**IbeaconMajor** | Pointer to **int32** |  | [optional] 
**IbeaconMinor** | Pointer to **int32** |  | [optional] 
**IbeaconUuid** | Pointer to **string** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**MapId** | Pointer to **string** |  | [optional] 
**MfgCompanyId** | Pointer to **int32** | optional, BLE manufacturing company ID | [optional] 
**MfgData** | Pointer to **string** | optional, BLE manufacturing data in hex byte-string format (ie: “112233AABBCC”) | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "asset"]
**X** | Pointer to **float32** | x, in meter | [optional] 
**Y** | Pointer to **float32** | y, in meter | [optional] 

## Methods

### NewWebhookLocationAssetEventsInner

`func NewWebhookLocationAssetEventsInner() *WebhookLocationAssetEventsInner`

NewWebhookLocationAssetEventsInner instantiates a new WebhookLocationAssetEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLocationAssetEventsInnerWithDefaults

`func NewWebhookLocationAssetEventsInnerWithDefaults() *WebhookLocationAssetEventsInner`

NewWebhookLocationAssetEventsInnerWithDefaults instantiates a new WebhookLocationAssetEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatteryVoltage

`func (o *WebhookLocationAssetEventsInner) GetBatteryVoltage() int32`

GetBatteryVoltage returns the BatteryVoltage field if non-nil, zero value otherwise.

### GetBatteryVoltageOk

`func (o *WebhookLocationAssetEventsInner) GetBatteryVoltageOk() (*int32, bool)`

GetBatteryVoltageOk returns a tuple with the BatteryVoltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryVoltage

`func (o *WebhookLocationAssetEventsInner) SetBatteryVoltage(v int32)`

SetBatteryVoltage sets BatteryVoltage field to given value.

### HasBatteryVoltage

`func (o *WebhookLocationAssetEventsInner) HasBatteryVoltage() bool`

HasBatteryVoltage returns a boolean if a field has been set.

### GetEddystoneUidInstance

`func (o *WebhookLocationAssetEventsInner) GetEddystoneUidInstance() string`

GetEddystoneUidInstance returns the EddystoneUidInstance field if non-nil, zero value otherwise.

### GetEddystoneUidInstanceOk

`func (o *WebhookLocationAssetEventsInner) GetEddystoneUidInstanceOk() (*string, bool)`

GetEddystoneUidInstanceOk returns a tuple with the EddystoneUidInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidInstance

`func (o *WebhookLocationAssetEventsInner) SetEddystoneUidInstance(v string)`

SetEddystoneUidInstance sets EddystoneUidInstance field to given value.

### HasEddystoneUidInstance

`func (o *WebhookLocationAssetEventsInner) HasEddystoneUidInstance() bool`

HasEddystoneUidInstance returns a boolean if a field has been set.

### GetEddystoneUidNamespace

`func (o *WebhookLocationAssetEventsInner) GetEddystoneUidNamespace() string`

GetEddystoneUidNamespace returns the EddystoneUidNamespace field if non-nil, zero value otherwise.

### GetEddystoneUidNamespaceOk

`func (o *WebhookLocationAssetEventsInner) GetEddystoneUidNamespaceOk() (*string, bool)`

GetEddystoneUidNamespaceOk returns a tuple with the EddystoneUidNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidNamespace

`func (o *WebhookLocationAssetEventsInner) SetEddystoneUidNamespace(v string)`

SetEddystoneUidNamespace sets EddystoneUidNamespace field to given value.

### HasEddystoneUidNamespace

`func (o *WebhookLocationAssetEventsInner) HasEddystoneUidNamespace() bool`

HasEddystoneUidNamespace returns a boolean if a field has been set.

### GetEddystoneUrlUrl

`func (o *WebhookLocationAssetEventsInner) GetEddystoneUrlUrl() string`

GetEddystoneUrlUrl returns the EddystoneUrlUrl field if non-nil, zero value otherwise.

### GetEddystoneUrlUrlOk

`func (o *WebhookLocationAssetEventsInner) GetEddystoneUrlUrlOk() (*string, bool)`

GetEddystoneUrlUrlOk returns a tuple with the EddystoneUrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlUrl

`func (o *WebhookLocationAssetEventsInner) SetEddystoneUrlUrl(v string)`

SetEddystoneUrlUrl sets EddystoneUrlUrl field to given value.

### HasEddystoneUrlUrl

`func (o *WebhookLocationAssetEventsInner) HasEddystoneUrlUrl() bool`

HasEddystoneUrlUrl returns a boolean if a field has been set.

### GetIbeaconMajor

`func (o *WebhookLocationAssetEventsInner) GetIbeaconMajor() int32`

GetIbeaconMajor returns the IbeaconMajor field if non-nil, zero value otherwise.

### GetIbeaconMajorOk

`func (o *WebhookLocationAssetEventsInner) GetIbeaconMajorOk() (*int32, bool)`

GetIbeaconMajorOk returns a tuple with the IbeaconMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMajor

`func (o *WebhookLocationAssetEventsInner) SetIbeaconMajor(v int32)`

SetIbeaconMajor sets IbeaconMajor field to given value.

### HasIbeaconMajor

`func (o *WebhookLocationAssetEventsInner) HasIbeaconMajor() bool`

HasIbeaconMajor returns a boolean if a field has been set.

### GetIbeaconMinor

`func (o *WebhookLocationAssetEventsInner) GetIbeaconMinor() int32`

GetIbeaconMinor returns the IbeaconMinor field if non-nil, zero value otherwise.

### GetIbeaconMinorOk

`func (o *WebhookLocationAssetEventsInner) GetIbeaconMinorOk() (*int32, bool)`

GetIbeaconMinorOk returns a tuple with the IbeaconMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMinor

`func (o *WebhookLocationAssetEventsInner) SetIbeaconMinor(v int32)`

SetIbeaconMinor sets IbeaconMinor field to given value.

### HasIbeaconMinor

`func (o *WebhookLocationAssetEventsInner) HasIbeaconMinor() bool`

HasIbeaconMinor returns a boolean if a field has been set.

### GetIbeaconUuid

`func (o *WebhookLocationAssetEventsInner) GetIbeaconUuid() string`

GetIbeaconUuid returns the IbeaconUuid field if non-nil, zero value otherwise.

### GetIbeaconUuidOk

`func (o *WebhookLocationAssetEventsInner) GetIbeaconUuidOk() (*string, bool)`

GetIbeaconUuidOk returns a tuple with the IbeaconUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconUuid

`func (o *WebhookLocationAssetEventsInner) SetIbeaconUuid(v string)`

SetIbeaconUuid sets IbeaconUuid field to given value.

### HasIbeaconUuid

`func (o *WebhookLocationAssetEventsInner) HasIbeaconUuid() bool`

HasIbeaconUuid returns a boolean if a field has been set.

### GetMac

`func (o *WebhookLocationAssetEventsInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookLocationAssetEventsInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookLocationAssetEventsInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *WebhookLocationAssetEventsInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMapId

`func (o *WebhookLocationAssetEventsInner) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *WebhookLocationAssetEventsInner) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *WebhookLocationAssetEventsInner) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *WebhookLocationAssetEventsInner) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetMfgCompanyId

`func (o *WebhookLocationAssetEventsInner) GetMfgCompanyId() int32`

GetMfgCompanyId returns the MfgCompanyId field if non-nil, zero value otherwise.

### GetMfgCompanyIdOk

`func (o *WebhookLocationAssetEventsInner) GetMfgCompanyIdOk() (*int32, bool)`

GetMfgCompanyIdOk returns a tuple with the MfgCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgCompanyId

`func (o *WebhookLocationAssetEventsInner) SetMfgCompanyId(v int32)`

SetMfgCompanyId sets MfgCompanyId field to given value.

### HasMfgCompanyId

`func (o *WebhookLocationAssetEventsInner) HasMfgCompanyId() bool`

HasMfgCompanyId returns a boolean if a field has been set.

### GetMfgData

`func (o *WebhookLocationAssetEventsInner) GetMfgData() string`

GetMfgData returns the MfgData field if non-nil, zero value otherwise.

### GetMfgDataOk

`func (o *WebhookLocationAssetEventsInner) GetMfgDataOk() (*string, bool)`

GetMfgDataOk returns a tuple with the MfgData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgData

`func (o *WebhookLocationAssetEventsInner) SetMfgData(v string)`

SetMfgData sets MfgData field to given value.

### HasMfgData

`func (o *WebhookLocationAssetEventsInner) HasMfgData() bool`

HasMfgData returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookLocationAssetEventsInner) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookLocationAssetEventsInner) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookLocationAssetEventsInner) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WebhookLocationAssetEventsInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookLocationAssetEventsInner) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookLocationAssetEventsInner) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookLocationAssetEventsInner) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebhookLocationAssetEventsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *WebhookLocationAssetEventsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookLocationAssetEventsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookLocationAssetEventsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WebhookLocationAssetEventsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetX

`func (o *WebhookLocationAssetEventsInner) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *WebhookLocationAssetEventsInner) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *WebhookLocationAssetEventsInner) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *WebhookLocationAssetEventsInner) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *WebhookLocationAssetEventsInner) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *WebhookLocationAssetEventsInner) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *WebhookLocationAssetEventsInner) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *WebhookLocationAssetEventsInner) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


