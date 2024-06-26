# WebhookOccupancyAlertsEventAlertEventsItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentOccupancy** | **int32** |  | 
**MapId** | **string** |  | 
**OccupancyLimit** | **int32** |  | 
**OrgId** | **string** |  | 
**Timestamp** | **float32** |  | 
**Type** | [**WebhookOccupancyAlertType**](WebhookOccupancyAlertType.md) |  | 
**ZoneId** | **string** |  | 
**ZoneName** | **string** |  | 

## Methods

### NewWebhookOccupancyAlertsEventAlertEventsItems

`func NewWebhookOccupancyAlertsEventAlertEventsItems(currentOccupancy int32, mapId string, occupancyLimit int32, orgId string, timestamp float32, type_ WebhookOccupancyAlertType, zoneId string, zoneName string, ) *WebhookOccupancyAlertsEventAlertEventsItems`

NewWebhookOccupancyAlertsEventAlertEventsItems instantiates a new WebhookOccupancyAlertsEventAlertEventsItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOccupancyAlertsEventAlertEventsItemsWithDefaults

`func NewWebhookOccupancyAlertsEventAlertEventsItemsWithDefaults() *WebhookOccupancyAlertsEventAlertEventsItems`

NewWebhookOccupancyAlertsEventAlertEventsItemsWithDefaults instantiates a new WebhookOccupancyAlertsEventAlertEventsItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentOccupancy

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetCurrentOccupancy() int32`

GetCurrentOccupancy returns the CurrentOccupancy field if non-nil, zero value otherwise.

### GetCurrentOccupancyOk

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetCurrentOccupancyOk() (*int32, bool)`

GetCurrentOccupancyOk returns a tuple with the CurrentOccupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOccupancy

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) SetCurrentOccupancy(v int32)`

SetCurrentOccupancy sets CurrentOccupancy field to given value.


### GetMapId

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetOccupancyLimit

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetOccupancyLimit() int32`

GetOccupancyLimit returns the OccupancyLimit field if non-nil, zero value otherwise.

### GetOccupancyLimitOk

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetOccupancyLimitOk() (*int32, bool)`

GetOccupancyLimitOk returns a tuple with the OccupancyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancyLimit

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) SetOccupancyLimit(v int32)`

SetOccupancyLimit sets OccupancyLimit field to given value.


### GetOrgId

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetTimestamp

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetType() WebhookOccupancyAlertType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetTypeOk() (*WebhookOccupancyAlertType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) SetType(v WebhookOccupancyAlertType)`

SetType sets Type field to given value.


### GetZoneId

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetZoneName

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *WebhookOccupancyAlertsEventAlertEventsItems) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


