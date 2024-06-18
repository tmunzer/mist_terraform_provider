# NetworkStaticNatValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalIp** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WanName** | Pointer to **string** | If not set, we configure the nat policies against all WAN ports for simplicity | [optional] 

## Methods

### NewNetworkStaticNatValue

`func NewNetworkStaticNatValue() *NetworkStaticNatValue`

NewNetworkStaticNatValue instantiates a new NetworkStaticNatValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkStaticNatValueWithDefaults

`func NewNetworkStaticNatValueWithDefaults() *NetworkStaticNatValue`

NewNetworkStaticNatValueWithDefaults instantiates a new NetworkStaticNatValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalIp

`func (o *NetworkStaticNatValue) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *NetworkStaticNatValue) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *NetworkStaticNatValue) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *NetworkStaticNatValue) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetName

`func (o *NetworkStaticNatValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkStaticNatValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkStaticNatValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkStaticNatValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWanName

`func (o *NetworkStaticNatValue) GetWanName() string`

GetWanName returns the WanName field if non-nil, zero value otherwise.

### GetWanNameOk

`func (o *NetworkStaticNatValue) GetWanNameOk() (*string, bool)`

GetWanNameOk returns a tuple with the WanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanName

`func (o *NetworkStaticNatValue) SetWanName(v string)`

SetWanName sets WanName field to given value.

### HasWanName

`func (o *NetworkStaticNatValue) HasWanName() bool`

HasWanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


