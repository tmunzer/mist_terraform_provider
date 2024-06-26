# WebhookLocationEventWifiBeaconExtendedInfoItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrameCtrl** | Pointer to **int32** | frame control field of 802.11 header | [optional] 
**Payload** | Pointer to **string** | Extended Info Payload associated with frame | [optional] 
**SeqCtrl** | Pointer to **int32** | sequence control field of 802.11 header | [optional] 

## Methods

### NewWebhookLocationEventWifiBeaconExtendedInfoItems

`func NewWebhookLocationEventWifiBeaconExtendedInfoItems() *WebhookLocationEventWifiBeaconExtendedInfoItems`

NewWebhookLocationEventWifiBeaconExtendedInfoItems instantiates a new WebhookLocationEventWifiBeaconExtendedInfoItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLocationEventWifiBeaconExtendedInfoItemsWithDefaults

`func NewWebhookLocationEventWifiBeaconExtendedInfoItemsWithDefaults() *WebhookLocationEventWifiBeaconExtendedInfoItems`

NewWebhookLocationEventWifiBeaconExtendedInfoItemsWithDefaults instantiates a new WebhookLocationEventWifiBeaconExtendedInfoItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrameCtrl

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) GetFrameCtrl() int32`

GetFrameCtrl returns the FrameCtrl field if non-nil, zero value otherwise.

### GetFrameCtrlOk

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) GetFrameCtrlOk() (*int32, bool)`

GetFrameCtrlOk returns a tuple with the FrameCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameCtrl

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) SetFrameCtrl(v int32)`

SetFrameCtrl sets FrameCtrl field to given value.

### HasFrameCtrl

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) HasFrameCtrl() bool`

HasFrameCtrl returns a boolean if a field has been set.

### GetPayload

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSeqCtrl

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) GetSeqCtrl() int32`

GetSeqCtrl returns the SeqCtrl field if non-nil, zero value otherwise.

### GetSeqCtrlOk

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) GetSeqCtrlOk() (*int32, bool)`

GetSeqCtrlOk returns a tuple with the SeqCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqCtrl

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) SetSeqCtrl(v int32)`

SetSeqCtrl sets SeqCtrl field to given value.

### HasSeqCtrl

`func (o *WebhookLocationEventWifiBeaconExtendedInfoItems) HasSeqCtrl() bool`

HasSeqCtrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


