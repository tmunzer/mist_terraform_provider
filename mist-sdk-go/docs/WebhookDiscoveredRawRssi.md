# WebhookDiscoveredRawRssi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]WebhookDiscoveredRawRssiEvent**](WebhookDiscoveredRawRssiEvent.md) |  | [optional] 
**Topic** | **string** |  | 

## Methods

### NewWebhookDiscoveredRawRssi

`func NewWebhookDiscoveredRawRssi(topic string, ) *WebhookDiscoveredRawRssi`

NewWebhookDiscoveredRawRssi instantiates a new WebhookDiscoveredRawRssi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDiscoveredRawRssiWithDefaults

`func NewWebhookDiscoveredRawRssiWithDefaults() *WebhookDiscoveredRawRssi`

NewWebhookDiscoveredRawRssiWithDefaults instantiates a new WebhookDiscoveredRawRssi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookDiscoveredRawRssi) GetEvents() []WebhookDiscoveredRawRssiEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookDiscoveredRawRssi) GetEventsOk() (*[]WebhookDiscoveredRawRssiEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookDiscoveredRawRssi) SetEvents(v []WebhookDiscoveredRawRssiEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookDiscoveredRawRssi) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTopic

`func (o *WebhookDiscoveredRawRssi) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookDiscoveredRawRssi) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookDiscoveredRawRssi) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


