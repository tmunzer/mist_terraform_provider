# WebhookClientLatency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]WebhookClientLatencyEvent**](WebhookClientLatencyEvent.md) |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookClientLatency

`func NewWebhookClientLatency() *WebhookClientLatency`

NewWebhookClientLatency instantiates a new WebhookClientLatency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookClientLatencyWithDefaults

`func NewWebhookClientLatencyWithDefaults() *WebhookClientLatency`

NewWebhookClientLatencyWithDefaults instantiates a new WebhookClientLatency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookClientLatency) GetEvents() []WebhookClientLatencyEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookClientLatency) GetEventsOk() (*[]WebhookClientLatencyEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookClientLatency) SetEvents(v []WebhookClientLatencyEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookClientLatency) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTopic

`func (o *WebhookClientLatency) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookClientLatency) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookClientLatency) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *WebhookClientLatency) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


