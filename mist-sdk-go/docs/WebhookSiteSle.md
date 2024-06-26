# WebhookSiteSle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]WebhookSiteSleEvent**](WebhookSiteSleEvent.md) |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookSiteSle

`func NewWebhookSiteSle() *WebhookSiteSle`

NewWebhookSiteSle instantiates a new WebhookSiteSle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSiteSleWithDefaults

`func NewWebhookSiteSleWithDefaults() *WebhookSiteSle`

NewWebhookSiteSleWithDefaults instantiates a new WebhookSiteSle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookSiteSle) GetEvents() []WebhookSiteSleEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookSiteSle) GetEventsOk() (*[]WebhookSiteSleEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookSiteSle) SetEvents(v []WebhookSiteSleEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookSiteSle) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTopic

`func (o *WebhookSiteSle) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookSiteSle) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookSiteSle) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *WebhookSiteSle) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


