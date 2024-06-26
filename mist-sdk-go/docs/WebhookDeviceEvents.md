# WebhookDeviceEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]WebhookDeviceEventsEvent**](WebhookDeviceEventsEvent.md) | list of events | 
**Topic** | **string** | topic subscribed to | [default to "device_events"]

## Methods

### NewWebhookDeviceEvents

`func NewWebhookDeviceEvents(events []WebhookDeviceEventsEvent, topic string, ) *WebhookDeviceEvents`

NewWebhookDeviceEvents instantiates a new WebhookDeviceEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeviceEventsWithDefaults

`func NewWebhookDeviceEventsWithDefaults() *WebhookDeviceEvents`

NewWebhookDeviceEventsWithDefaults instantiates a new WebhookDeviceEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookDeviceEvents) GetEvents() []WebhookDeviceEventsEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookDeviceEvents) GetEventsOk() (*[]WebhookDeviceEventsEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookDeviceEvents) SetEvents(v []WebhookDeviceEventsEvent)`

SetEvents sets Events field to given value.


### GetTopic

`func (o *WebhookDeviceEvents) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookDeviceEvents) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookDeviceEvents) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


