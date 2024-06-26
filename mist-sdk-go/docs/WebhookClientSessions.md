# WebhookClientSessions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]WebhookClientSessionsEvent**](WebhookClientSessionsEvent.md) |  | 
**Topic** | **string** |  | [default to "client-sessions"]

## Methods

### NewWebhookClientSessions

`func NewWebhookClientSessions(events []WebhookClientSessionsEvent, topic string, ) *WebhookClientSessions`

NewWebhookClientSessions instantiates a new WebhookClientSessions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookClientSessionsWithDefaults

`func NewWebhookClientSessionsWithDefaults() *WebhookClientSessions`

NewWebhookClientSessionsWithDefaults instantiates a new WebhookClientSessions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookClientSessions) GetEvents() []WebhookClientSessionsEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookClientSessions) GetEventsOk() (*[]WebhookClientSessionsEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookClientSessions) SetEvents(v []WebhookClientSessionsEvent)`

SetEvents sets Events field to given value.


### GetTopic

`func (o *WebhookClientSessions) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookClientSessions) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookClientSessions) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


