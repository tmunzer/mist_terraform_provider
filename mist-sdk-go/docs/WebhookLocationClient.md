# WebhookLocationClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]WebhookLocationClientEventsInner**](WebhookLocationClientEventsInner.md) | list of events | 
**Topic** | **string** | topic subscribed to | [default to "location_client"]

## Methods

### NewWebhookLocationClient

`func NewWebhookLocationClient(events []WebhookLocationClientEventsInner, topic string, ) *WebhookLocationClient`

NewWebhookLocationClient instantiates a new WebhookLocationClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLocationClientWithDefaults

`func NewWebhookLocationClientWithDefaults() *WebhookLocationClient`

NewWebhookLocationClientWithDefaults instantiates a new WebhookLocationClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookLocationClient) GetEvents() []WebhookLocationClientEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookLocationClient) GetEventsOk() (*[]WebhookLocationClientEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookLocationClient) SetEvents(v []WebhookLocationClientEventsInner)`

SetEvents sets Events field to given value.


### GetTopic

`func (o *WebhookLocationClient) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookLocationClient) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookLocationClient) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


