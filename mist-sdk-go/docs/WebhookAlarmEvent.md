# WebhookAlarmEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aps** | Pointer to **[]string** |  | [optional] [readonly] 
**Bssids** | Pointer to **[]string** |  | [optional] [readonly] 
**Count** | Pointer to **int32** | If present, represents number of events of given type occurred in current interval, default&#x3D;1 | [optional] [readonly] 
**EventId** | Pointer to **string** | event id | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | **string** |  | [readonly] 
**LastSeen** | **float32** |  | [readonly] 
**Node** | Pointer to [**HaClusterNodeEnum**](HaClusterNodeEnum.md) |  | [optional] 
**OrgId** | **string** |  | [readonly] 
**SiteId** | **string** |  | [readonly] 
**Ssids** | Pointer to **[]string** |  | [optional] [readonly] 
**Timestamp** | **int32** |  | [readonly] 
**Type** | **string** | event type | [readonly] 
**Update** | Pointer to **bool** | If presents, represents that this is an update to event with given id sent earlier. default&#x3D;false | [optional] [readonly] 

## Methods

### NewWebhookAlarmEvent

`func NewWebhookAlarmEvent(id string, lastSeen float32, orgId string, siteId string, timestamp int32, type_ string, ) *WebhookAlarmEvent`

NewWebhookAlarmEvent instantiates a new WebhookAlarmEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAlarmEventWithDefaults

`func NewWebhookAlarmEventWithDefaults() *WebhookAlarmEvent`

NewWebhookAlarmEventWithDefaults instantiates a new WebhookAlarmEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAps

`func (o *WebhookAlarmEvent) GetAps() []string`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *WebhookAlarmEvent) GetApsOk() (*[]string, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *WebhookAlarmEvent) SetAps(v []string)`

SetAps sets Aps field to given value.

### HasAps

`func (o *WebhookAlarmEvent) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetBssids

`func (o *WebhookAlarmEvent) GetBssids() []string`

GetBssids returns the Bssids field if non-nil, zero value otherwise.

### GetBssidsOk

`func (o *WebhookAlarmEvent) GetBssidsOk() (*[]string, bool)`

GetBssidsOk returns a tuple with the Bssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssids

`func (o *WebhookAlarmEvent) SetBssids(v []string)`

SetBssids sets Bssids field to given value.

### HasBssids

`func (o *WebhookAlarmEvent) HasBssids() bool`

HasBssids returns a boolean if a field has been set.

### GetCount

`func (o *WebhookAlarmEvent) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WebhookAlarmEvent) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WebhookAlarmEvent) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *WebhookAlarmEvent) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetEventId

`func (o *WebhookAlarmEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookAlarmEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookAlarmEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *WebhookAlarmEvent) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetForSite

`func (o *WebhookAlarmEvent) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *WebhookAlarmEvent) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *WebhookAlarmEvent) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *WebhookAlarmEvent) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *WebhookAlarmEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookAlarmEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookAlarmEvent) SetId(v string)`

SetId sets Id field to given value.


### GetLastSeen

`func (o *WebhookAlarmEvent) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *WebhookAlarmEvent) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *WebhookAlarmEvent) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.


### GetNode

`func (o *WebhookAlarmEvent) GetNode() HaClusterNodeEnum`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *WebhookAlarmEvent) GetNodeOk() (*HaClusterNodeEnum, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *WebhookAlarmEvent) SetNode(v HaClusterNodeEnum)`

SetNode sets Node field to given value.

### HasNode

`func (o *WebhookAlarmEvent) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetOrgId

`func (o *WebhookAlarmEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookAlarmEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookAlarmEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetSiteId

`func (o *WebhookAlarmEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookAlarmEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookAlarmEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSsids

`func (o *WebhookAlarmEvent) GetSsids() []string`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *WebhookAlarmEvent) GetSsidsOk() (*[]string, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *WebhookAlarmEvent) SetSsids(v []string)`

SetSsids sets Ssids field to given value.

### HasSsids

`func (o *WebhookAlarmEvent) HasSsids() bool`

HasSsids returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookAlarmEvent) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookAlarmEvent) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookAlarmEvent) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *WebhookAlarmEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAlarmEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAlarmEvent) SetType(v string)`

SetType sets Type field to given value.


### GetUpdate

`func (o *WebhookAlarmEvent) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *WebhookAlarmEvent) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *WebhookAlarmEvent) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *WebhookAlarmEvent) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


