# ApClientBridgeAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Psk** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | wpa2-AES/CCMPp is assumed when &#x60;type&#x60;&#x3D;&#x3D;&#x60;psk&#x60; | [optional] [default to "psk"]

## Methods

### NewApClientBridgeAuth

`func NewApClientBridgeAuth() *ApClientBridgeAuth`

NewApClientBridgeAuth instantiates a new ApClientBridgeAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApClientBridgeAuthWithDefaults

`func NewApClientBridgeAuthWithDefaults() *ApClientBridgeAuth`

NewApClientBridgeAuthWithDefaults instantiates a new ApClientBridgeAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPsk

`func (o *ApClientBridgeAuth) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *ApClientBridgeAuth) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *ApClientBridgeAuth) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *ApClientBridgeAuth) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetType

`func (o *ApClientBridgeAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApClientBridgeAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApClientBridgeAuth) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApClientBridgeAuth) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


