# WebhookOccupancyAlerts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]WebhookOccupancyAlertsEvent**](WebhookOccupancyAlertsEvent.md) |  | 
**Topic** | **string** |  | [default to "occupancy-alerts"]

## Methods

### NewWebhookOccupancyAlerts

`func NewWebhookOccupancyAlerts(events []WebhookOccupancyAlertsEvent, topic string, ) *WebhookOccupancyAlerts`

NewWebhookOccupancyAlerts instantiates a new WebhookOccupancyAlerts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOccupancyAlertsWithDefaults

`func NewWebhookOccupancyAlertsWithDefaults() *WebhookOccupancyAlerts`

NewWebhookOccupancyAlertsWithDefaults instantiates a new WebhookOccupancyAlerts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookOccupancyAlerts) GetEvents() []WebhookOccupancyAlertsEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookOccupancyAlerts) GetEventsOk() (*[]WebhookOccupancyAlertsEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookOccupancyAlerts) SetEvents(v []WebhookOccupancyAlertsEvent)`

SetEvents sets Events field to given value.


### GetTopic

`func (o *WebhookOccupancyAlerts) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookOccupancyAlerts) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookOccupancyAlerts) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


