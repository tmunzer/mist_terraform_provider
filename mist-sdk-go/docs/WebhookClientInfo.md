# WebhookClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]WebhookClientInfoEvent**](WebhookClientInfoEvent.md) |  | [optional] 
**Topic** | Pointer to [**WebhookClientInfoTopic**](WebhookClientInfoTopic.md) |  | [optional] 

## Methods

### NewWebhookClientInfo

`func NewWebhookClientInfo() *WebhookClientInfo`

NewWebhookClientInfo instantiates a new WebhookClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookClientInfoWithDefaults

`func NewWebhookClientInfoWithDefaults() *WebhookClientInfo`

NewWebhookClientInfoWithDefaults instantiates a new WebhookClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookClientInfo) GetEvents() []WebhookClientInfoEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookClientInfo) GetEventsOk() (*[]WebhookClientInfoEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookClientInfo) SetEvents(v []WebhookClientInfoEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookClientInfo) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTopic

`func (o *WebhookClientInfo) GetTopic() WebhookClientInfoTopic`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookClientInfo) GetTopicOk() (*WebhookClientInfoTopic, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookClientInfo) SetTopic(v WebhookClientInfoTopic)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *WebhookClientInfo) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


