# WebhookLocationEvent

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
**Id** | **string** | unique id of the client (a client would have different id for different org) | 
**Mac** | Pointer to **string** |  | [optional] 
**MapId** | **string** | map id | 
**MfgCompanyId** | Pointer to **int32** | optional, BLE manufacturing company ID | [optional] 
**MfgData** | Pointer to **string** | optional, BLE manufacturing data in hex byte-string format (ie \&quot;112233AABBCC\&quot;) | [optional] 
**Name** | Pointer to **string** | name of the client, may be empty | [optional] 
**SiteId** | **string** |  | [readonly] 
**Timestamp** | **int32** | timestamp of the event, epoch | 
**Type** | **string** |  | 
**WifiBeaconExtendedInfo** | Pointer to [**[]WebhookLocationEventWifiBeaconExtendedInfoItems**](WebhookLocationEventWifiBeaconExtendedInfoItems.md) | optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload | [optional] 
**X** | **int32** | x, in meter | 
**Y** | **int32** | y, in meter | 

## Methods

### NewWebhookLocationEvent

`func NewWebhookLocationEvent(id string, mapId string, siteId string, timestamp int32, type_ string, x int32, y int32, ) *WebhookLocationEvent`

NewWebhookLocationEvent instantiates a new WebhookLocationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLocationEventWithDefaults

`func NewWebhookLocationEventWithDefaults() *WebhookLocationEvent`

NewWebhookLocationEventWithDefaults instantiates a new WebhookLocationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatteryVoltage

`func (o *WebhookLocationEvent) GetBatteryVoltage() int32`

GetBatteryVoltage returns the BatteryVoltage field if non-nil, zero value otherwise.

### GetBatteryVoltageOk

`func (o *WebhookLocationEvent) GetBatteryVoltageOk() (*int32, bool)`

GetBatteryVoltageOk returns a tuple with the BatteryVoltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryVoltage

`func (o *WebhookLocationEvent) SetBatteryVoltage(v int32)`

SetBatteryVoltage sets BatteryVoltage field to given value.

### HasBatteryVoltage

`func (o *WebhookLocationEvent) HasBatteryVoltage() bool`

HasBatteryVoltage returns a boolean if a field has been set.

### GetEddystoneUidInstance

`func (o *WebhookLocationEvent) GetEddystoneUidInstance() string`

GetEddystoneUidInstance returns the EddystoneUidInstance field if non-nil, zero value otherwise.

### GetEddystoneUidInstanceOk

`func (o *WebhookLocationEvent) GetEddystoneUidInstanceOk() (*string, bool)`

GetEddystoneUidInstanceOk returns a tuple with the EddystoneUidInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidInstance

`func (o *WebhookLocationEvent) SetEddystoneUidInstance(v string)`

SetEddystoneUidInstance sets EddystoneUidInstance field to given value.

### HasEddystoneUidInstance

`func (o *WebhookLocationEvent) HasEddystoneUidInstance() bool`

HasEddystoneUidInstance returns a boolean if a field has been set.

### GetEddystoneUidNamespace

`func (o *WebhookLocationEvent) GetEddystoneUidNamespace() string`

GetEddystoneUidNamespace returns the EddystoneUidNamespace field if non-nil, zero value otherwise.

### GetEddystoneUidNamespaceOk

`func (o *WebhookLocationEvent) GetEddystoneUidNamespaceOk() (*string, bool)`

GetEddystoneUidNamespaceOk returns a tuple with the EddystoneUidNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidNamespace

`func (o *WebhookLocationEvent) SetEddystoneUidNamespace(v string)`

SetEddystoneUidNamespace sets EddystoneUidNamespace field to given value.

### HasEddystoneUidNamespace

`func (o *WebhookLocationEvent) HasEddystoneUidNamespace() bool`

HasEddystoneUidNamespace returns a boolean if a field has been set.

### GetEddystoneUrlUrl

`func (o *WebhookLocationEvent) GetEddystoneUrlUrl() string`

GetEddystoneUrlUrl returns the EddystoneUrlUrl field if non-nil, zero value otherwise.

### GetEddystoneUrlUrlOk

`func (o *WebhookLocationEvent) GetEddystoneUrlUrlOk() (*string, bool)`

GetEddystoneUrlUrlOk returns a tuple with the EddystoneUrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlUrl

`func (o *WebhookLocationEvent) SetEddystoneUrlUrl(v string)`

SetEddystoneUrlUrl sets EddystoneUrlUrl field to given value.

### HasEddystoneUrlUrl

`func (o *WebhookLocationEvent) HasEddystoneUrlUrl() bool`

HasEddystoneUrlUrl returns a boolean if a field has been set.

### GetIbeaconMajor

`func (o *WebhookLocationEvent) GetIbeaconMajor() int32`

GetIbeaconMajor returns the IbeaconMajor field if non-nil, zero value otherwise.

### GetIbeaconMajorOk

`func (o *WebhookLocationEvent) GetIbeaconMajorOk() (*int32, bool)`

GetIbeaconMajorOk returns a tuple with the IbeaconMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMajor

`func (o *WebhookLocationEvent) SetIbeaconMajor(v int32)`

SetIbeaconMajor sets IbeaconMajor field to given value.

### HasIbeaconMajor

`func (o *WebhookLocationEvent) HasIbeaconMajor() bool`

HasIbeaconMajor returns a boolean if a field has been set.

### GetIbeaconMinor

`func (o *WebhookLocationEvent) GetIbeaconMinor() int32`

GetIbeaconMinor returns the IbeaconMinor field if non-nil, zero value otherwise.

### GetIbeaconMinorOk

`func (o *WebhookLocationEvent) GetIbeaconMinorOk() (*int32, bool)`

GetIbeaconMinorOk returns a tuple with the IbeaconMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMinor

`func (o *WebhookLocationEvent) SetIbeaconMinor(v int32)`

SetIbeaconMinor sets IbeaconMinor field to given value.

### HasIbeaconMinor

`func (o *WebhookLocationEvent) HasIbeaconMinor() bool`

HasIbeaconMinor returns a boolean if a field has been set.

### GetIbeaconUuid

`func (o *WebhookLocationEvent) GetIbeaconUuid() string`

GetIbeaconUuid returns the IbeaconUuid field if non-nil, zero value otherwise.

### GetIbeaconUuidOk

`func (o *WebhookLocationEvent) GetIbeaconUuidOk() (*string, bool)`

GetIbeaconUuidOk returns a tuple with the IbeaconUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconUuid

`func (o *WebhookLocationEvent) SetIbeaconUuid(v string)`

SetIbeaconUuid sets IbeaconUuid field to given value.

### HasIbeaconUuid

`func (o *WebhookLocationEvent) HasIbeaconUuid() bool`

HasIbeaconUuid returns a boolean if a field has been set.

### GetId

`func (o *WebhookLocationEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookLocationEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookLocationEvent) SetId(v string)`

SetId sets Id field to given value.


### GetMac

`func (o *WebhookLocationEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookLocationEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookLocationEvent) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *WebhookLocationEvent) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMapId

`func (o *WebhookLocationEvent) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *WebhookLocationEvent) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *WebhookLocationEvent) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetMfgCompanyId

`func (o *WebhookLocationEvent) GetMfgCompanyId() int32`

GetMfgCompanyId returns the MfgCompanyId field if non-nil, zero value otherwise.

### GetMfgCompanyIdOk

`func (o *WebhookLocationEvent) GetMfgCompanyIdOk() (*int32, bool)`

GetMfgCompanyIdOk returns a tuple with the MfgCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgCompanyId

`func (o *WebhookLocationEvent) SetMfgCompanyId(v int32)`

SetMfgCompanyId sets MfgCompanyId field to given value.

### HasMfgCompanyId

`func (o *WebhookLocationEvent) HasMfgCompanyId() bool`

HasMfgCompanyId returns a boolean if a field has been set.

### GetMfgData

`func (o *WebhookLocationEvent) GetMfgData() string`

GetMfgData returns the MfgData field if non-nil, zero value otherwise.

### GetMfgDataOk

`func (o *WebhookLocationEvent) GetMfgDataOk() (*string, bool)`

GetMfgDataOk returns a tuple with the MfgData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgData

`func (o *WebhookLocationEvent) SetMfgData(v string)`

SetMfgData sets MfgData field to given value.

### HasMfgData

`func (o *WebhookLocationEvent) HasMfgData() bool`

HasMfgData returns a boolean if a field has been set.

### GetName

`func (o *WebhookLocationEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookLocationEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookLocationEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookLocationEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookLocationEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookLocationEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookLocationEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetTimestamp

`func (o *WebhookLocationEvent) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookLocationEvent) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookLocationEvent) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *WebhookLocationEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookLocationEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookLocationEvent) SetType(v string)`

SetType sets Type field to given value.


### GetWifiBeaconExtendedInfo

`func (o *WebhookLocationEvent) GetWifiBeaconExtendedInfo() []WebhookLocationEventWifiBeaconExtendedInfoItems`

GetWifiBeaconExtendedInfo returns the WifiBeaconExtendedInfo field if non-nil, zero value otherwise.

### GetWifiBeaconExtendedInfoOk

`func (o *WebhookLocationEvent) GetWifiBeaconExtendedInfoOk() (*[]WebhookLocationEventWifiBeaconExtendedInfoItems, bool)`

GetWifiBeaconExtendedInfoOk returns a tuple with the WifiBeaconExtendedInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiBeaconExtendedInfo

`func (o *WebhookLocationEvent) SetWifiBeaconExtendedInfo(v []WebhookLocationEventWifiBeaconExtendedInfoItems)`

SetWifiBeaconExtendedInfo sets WifiBeaconExtendedInfo field to given value.

### HasWifiBeaconExtendedInfo

`func (o *WebhookLocationEvent) HasWifiBeaconExtendedInfo() bool`

HasWifiBeaconExtendedInfo returns a boolean if a field has been set.

### GetX

`func (o *WebhookLocationEvent) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *WebhookLocationEvent) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *WebhookLocationEvent) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *WebhookLocationEvent) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *WebhookLocationEvent) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *WebhookLocationEvent) SetY(v int32)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


