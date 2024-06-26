# WebhookLocationClientEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** |  | [optional] 
**MapId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "wifi"]
**WifiBeaconExtendedInfo** | Pointer to [**[]WebhookLocationClientEventsInnerWifiBeaconExtendedInfoInner**](WebhookLocationClientEventsInnerWifiBeaconExtendedInfoInner.md) | optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload | [optional] 
**X** | Pointer to **float32** | x, in meter | [optional] 
**Y** | Pointer to **float32** | y, in meter | [optional] 

## Methods

### NewWebhookLocationClientEventsInner

`func NewWebhookLocationClientEventsInner() *WebhookLocationClientEventsInner`

NewWebhookLocationClientEventsInner instantiates a new WebhookLocationClientEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLocationClientEventsInnerWithDefaults

`func NewWebhookLocationClientEventsInnerWithDefaults() *WebhookLocationClientEventsInner`

NewWebhookLocationClientEventsInnerWithDefaults instantiates a new WebhookLocationClientEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *WebhookLocationClientEventsInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookLocationClientEventsInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookLocationClientEventsInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *WebhookLocationClientEventsInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMapId

`func (o *WebhookLocationClientEventsInner) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *WebhookLocationClientEventsInner) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *WebhookLocationClientEventsInner) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *WebhookLocationClientEventsInner) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookLocationClientEventsInner) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookLocationClientEventsInner) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookLocationClientEventsInner) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WebhookLocationClientEventsInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookLocationClientEventsInner) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookLocationClientEventsInner) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookLocationClientEventsInner) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebhookLocationClientEventsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *WebhookLocationClientEventsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookLocationClientEventsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookLocationClientEventsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WebhookLocationClientEventsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWifiBeaconExtendedInfo

`func (o *WebhookLocationClientEventsInner) GetWifiBeaconExtendedInfo() []WebhookLocationClientEventsInnerWifiBeaconExtendedInfoInner`

GetWifiBeaconExtendedInfo returns the WifiBeaconExtendedInfo field if non-nil, zero value otherwise.

### GetWifiBeaconExtendedInfoOk

`func (o *WebhookLocationClientEventsInner) GetWifiBeaconExtendedInfoOk() (*[]WebhookLocationClientEventsInnerWifiBeaconExtendedInfoInner, bool)`

GetWifiBeaconExtendedInfoOk returns a tuple with the WifiBeaconExtendedInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiBeaconExtendedInfo

`func (o *WebhookLocationClientEventsInner) SetWifiBeaconExtendedInfo(v []WebhookLocationClientEventsInnerWifiBeaconExtendedInfoInner)`

SetWifiBeaconExtendedInfo sets WifiBeaconExtendedInfo field to given value.

### HasWifiBeaconExtendedInfo

`func (o *WebhookLocationClientEventsInner) HasWifiBeaconExtendedInfo() bool`

HasWifiBeaconExtendedInfo returns a boolean if a field has been set.

### GetX

`func (o *WebhookLocationClientEventsInner) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *WebhookLocationClientEventsInner) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *WebhookLocationClientEventsInner) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *WebhookLocationClientEventsInner) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *WebhookLocationClientEventsInner) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *WebhookLocationClientEventsInner) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *WebhookLocationClientEventsInner) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *WebhookLocationClientEventsInner) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


