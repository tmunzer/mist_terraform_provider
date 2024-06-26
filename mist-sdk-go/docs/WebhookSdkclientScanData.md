# WebhookSdkclientScanData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]WebhookSdkclientScanDataEvent**](WebhookSdkclientScanDataEvent.md) |  | 
**Topic** | [**WebhookSdkclientScanDataTopic**](WebhookSdkclientScanDataTopic.md) |  | [default to WEBHOOKSDKCLIENTSCANDATATOPIC_SDKCLIENT_SCAN_DATA]

## Methods

### NewWebhookSdkclientScanData

`func NewWebhookSdkclientScanData(events []WebhookSdkclientScanDataEvent, topic WebhookSdkclientScanDataTopic, ) *WebhookSdkclientScanData`

NewWebhookSdkclientScanData instantiates a new WebhookSdkclientScanData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSdkclientScanDataWithDefaults

`func NewWebhookSdkclientScanDataWithDefaults() *WebhookSdkclientScanData`

NewWebhookSdkclientScanDataWithDefaults instantiates a new WebhookSdkclientScanData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookSdkclientScanData) GetEvents() []WebhookSdkclientScanDataEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookSdkclientScanData) GetEventsOk() (*[]WebhookSdkclientScanDataEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookSdkclientScanData) SetEvents(v []WebhookSdkclientScanDataEvent)`

SetEvents sets Events field to given value.


### GetTopic

`func (o *WebhookSdkclientScanData) GetTopic() WebhookSdkclientScanDataTopic`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookSdkclientScanData) GetTopicOk() (*WebhookSdkclientScanDataTopic, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookSdkclientScanData) SetTopic(v WebhookSdkclientScanDataTopic)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


