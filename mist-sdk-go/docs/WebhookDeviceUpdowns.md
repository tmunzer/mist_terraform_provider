# WebhookDeviceUpdowns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]WebhookDeviceUpdownsEvent**](WebhookDeviceUpdownsEvent.md) |  | 
**Topic** | **string** |  | [default to "device-updowns"]

## Methods

### NewWebhookDeviceUpdowns

`func NewWebhookDeviceUpdowns(events []WebhookDeviceUpdownsEvent, topic string, ) *WebhookDeviceUpdowns`

NewWebhookDeviceUpdowns instantiates a new WebhookDeviceUpdowns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeviceUpdownsWithDefaults

`func NewWebhookDeviceUpdownsWithDefaults() *WebhookDeviceUpdowns`

NewWebhookDeviceUpdownsWithDefaults instantiates a new WebhookDeviceUpdowns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookDeviceUpdowns) GetEvents() []WebhookDeviceUpdownsEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookDeviceUpdowns) GetEventsOk() (*[]WebhookDeviceUpdownsEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookDeviceUpdowns) SetEvents(v []WebhookDeviceUpdownsEvent)`

SetEvents sets Events field to given value.


### GetTopic

`func (o *WebhookDeviceUpdowns) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookDeviceUpdowns) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookDeviceUpdowns) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


