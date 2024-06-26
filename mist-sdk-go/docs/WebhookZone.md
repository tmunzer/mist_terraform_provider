# WebhookZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]WebhookZoneEvent**](WebhookZoneEvent.md) | list of events | 
**Topic** | **string** | topic subscribed to | [default to "zone"]

## Methods

### NewWebhookZone

`func NewWebhookZone(events []WebhookZoneEvent, topic string, ) *WebhookZone`

NewWebhookZone instantiates a new WebhookZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookZoneWithDefaults

`func NewWebhookZoneWithDefaults() *WebhookZone`

NewWebhookZoneWithDefaults instantiates a new WebhookZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookZone) GetEvents() []WebhookZoneEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookZone) GetEventsOk() (*[]WebhookZoneEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookZone) SetEvents(v []WebhookZoneEvent)`

SetEvents sets Events field to given value.


### GetTopic

`func (o *WebhookZone) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookZone) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookZone) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


