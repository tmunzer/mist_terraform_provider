# WebhookAudits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]WebhookAuditEvent**](WebhookAuditEvent.md) |  | 
**Topic** | **string** |  | [default to "audits"]

## Methods

### NewWebhookAudits

`func NewWebhookAudits(events []WebhookAuditEvent, topic string, ) *WebhookAudits`

NewWebhookAudits instantiates a new WebhookAudits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuditsWithDefaults

`func NewWebhookAuditsWithDefaults() *WebhookAudits`

NewWebhookAuditsWithDefaults instantiates a new WebhookAudits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookAudits) GetEvents() []WebhookAuditEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookAudits) GetEventsOk() (*[]WebhookAuditEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookAudits) SetEvents(v []WebhookAuditEvent)`

SetEvents sets Events field to given value.


### GetTopic

`func (o *WebhookAudits) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookAudits) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookAudits) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


