# WebhookNacAccounting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]WebhookNacAccountingEvent**](WebhookNacAccountingEvent.md) |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] [default to "nac-accounting"]

## Methods

### NewWebhookNacAccounting

`func NewWebhookNacAccounting() *WebhookNacAccounting`

NewWebhookNacAccounting instantiates a new WebhookNacAccounting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookNacAccountingWithDefaults

`func NewWebhookNacAccountingWithDefaults() *WebhookNacAccounting`

NewWebhookNacAccountingWithDefaults instantiates a new WebhookNacAccounting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookNacAccounting) GetEvents() []WebhookNacAccountingEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookNacAccounting) GetEventsOk() (*[]WebhookNacAccountingEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookNacAccounting) SetEvents(v []WebhookNacAccountingEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookNacAccounting) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTopic

`func (o *WebhookNacAccounting) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookNacAccounting) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookNacAccounting) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *WebhookNacAccounting) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


