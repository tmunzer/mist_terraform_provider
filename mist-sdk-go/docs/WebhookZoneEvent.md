# WebhookZoneEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | Pointer to **string** | uuid of named asset | [optional] 
**Id** | **string** | uuid of SDK-client | 
**Mac** | Pointer to **string** | mac address of wifi client or asset | [optional] 
**MapId** | **string** | map id | 
**Name** | Pointer to **string** | name of the client, may be empty | [optional] 
**SiteId** | **string** |  | 
**Timestamp** | **int32** | timestamp of the event, epoch | 
**Trigger** | [**WebhookZoneEventTrigger**](WebhookZoneEventTrigger.md) |  | 
**Type** | **string** |  | 
**ZoneId** | **string** | zone id | 

## Methods

### NewWebhookZoneEvent

`func NewWebhookZoneEvent(id string, mapId string, siteId string, timestamp int32, trigger WebhookZoneEventTrigger, type_ string, zoneId string, ) *WebhookZoneEvent`

NewWebhookZoneEvent instantiates a new WebhookZoneEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookZoneEventWithDefaults

`func NewWebhookZoneEventWithDefaults() *WebhookZoneEvent`

NewWebhookZoneEventWithDefaults instantiates a new WebhookZoneEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *WebhookZoneEvent) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *WebhookZoneEvent) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *WebhookZoneEvent) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *WebhookZoneEvent) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetId

`func (o *WebhookZoneEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookZoneEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookZoneEvent) SetId(v string)`

SetId sets Id field to given value.


### GetMac

`func (o *WebhookZoneEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WebhookZoneEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WebhookZoneEvent) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *WebhookZoneEvent) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMapId

`func (o *WebhookZoneEvent) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *WebhookZoneEvent) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *WebhookZoneEvent) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetName

`func (o *WebhookZoneEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookZoneEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookZoneEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookZoneEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookZoneEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookZoneEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookZoneEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetTimestamp

`func (o *WebhookZoneEvent) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookZoneEvent) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookZoneEvent) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTrigger

`func (o *WebhookZoneEvent) GetTrigger() WebhookZoneEventTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *WebhookZoneEvent) GetTriggerOk() (*WebhookZoneEventTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *WebhookZoneEvent) SetTrigger(v WebhookZoneEventTrigger)`

SetTrigger sets Trigger field to given value.


### GetType

`func (o *WebhookZoneEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookZoneEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookZoneEvent) SetType(v string)`

SetType sets Type field to given value.


### GetZoneId

`func (o *WebhookZoneEvent) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *WebhookZoneEvent) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *WebhookZoneEvent) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


