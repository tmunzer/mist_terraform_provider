# SnmpV3ConfigTargetAddressInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AddressMask** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] [default to 161]
**TagList** | Pointer to **string** | &lt;refer to notify tag, can be multiple with blank | [optional] 
**TargetAddressName** | Pointer to **string** |  | [optional] 
**TargetParameters** | Pointer to **string** | refer to notify target parameters name | [optional] 

## Methods

### NewSnmpV3ConfigTargetAddressInner

`func NewSnmpV3ConfigTargetAddressInner() *SnmpV3ConfigTargetAddressInner`

NewSnmpV3ConfigTargetAddressInner instantiates a new SnmpV3ConfigTargetAddressInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpV3ConfigTargetAddressInnerWithDefaults

`func NewSnmpV3ConfigTargetAddressInnerWithDefaults() *SnmpV3ConfigTargetAddressInner`

NewSnmpV3ConfigTargetAddressInnerWithDefaults instantiates a new SnmpV3ConfigTargetAddressInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SnmpV3ConfigTargetAddressInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SnmpV3ConfigTargetAddressInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SnmpV3ConfigTargetAddressInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SnmpV3ConfigTargetAddressInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressMask

`func (o *SnmpV3ConfigTargetAddressInner) GetAddressMask() string`

GetAddressMask returns the AddressMask field if non-nil, zero value otherwise.

### GetAddressMaskOk

`func (o *SnmpV3ConfigTargetAddressInner) GetAddressMaskOk() (*string, bool)`

GetAddressMaskOk returns a tuple with the AddressMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMask

`func (o *SnmpV3ConfigTargetAddressInner) SetAddressMask(v string)`

SetAddressMask sets AddressMask field to given value.

### HasAddressMask

`func (o *SnmpV3ConfigTargetAddressInner) HasAddressMask() bool`

HasAddressMask returns a boolean if a field has been set.

### GetPort

`func (o *SnmpV3ConfigTargetAddressInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SnmpV3ConfigTargetAddressInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SnmpV3ConfigTargetAddressInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SnmpV3ConfigTargetAddressInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTagList

`func (o *SnmpV3ConfigTargetAddressInner) GetTagList() string`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *SnmpV3ConfigTargetAddressInner) GetTagListOk() (*string, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *SnmpV3ConfigTargetAddressInner) SetTagList(v string)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *SnmpV3ConfigTargetAddressInner) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetTargetAddressName

`func (o *SnmpV3ConfigTargetAddressInner) GetTargetAddressName() string`

GetTargetAddressName returns the TargetAddressName field if non-nil, zero value otherwise.

### GetTargetAddressNameOk

`func (o *SnmpV3ConfigTargetAddressInner) GetTargetAddressNameOk() (*string, bool)`

GetTargetAddressNameOk returns a tuple with the TargetAddressName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAddressName

`func (o *SnmpV3ConfigTargetAddressInner) SetTargetAddressName(v string)`

SetTargetAddressName sets TargetAddressName field to given value.

### HasTargetAddressName

`func (o *SnmpV3ConfigTargetAddressInner) HasTargetAddressName() bool`

HasTargetAddressName returns a boolean if a field has been set.

### GetTargetParameters

`func (o *SnmpV3ConfigTargetAddressInner) GetTargetParameters() string`

GetTargetParameters returns the TargetParameters field if non-nil, zero value otherwise.

### GetTargetParametersOk

`func (o *SnmpV3ConfigTargetAddressInner) GetTargetParametersOk() (*string, bool)`

GetTargetParametersOk returns a tuple with the TargetParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetParameters

`func (o *SnmpV3ConfigTargetAddressInner) SetTargetParameters(v string)`

SetTargetParameters sets TargetParameters field to given value.

### HasTargetParameters

`func (o *SnmpV3ConfigTargetAddressInner) HasTargetParameters() bool`

HasTargetParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


