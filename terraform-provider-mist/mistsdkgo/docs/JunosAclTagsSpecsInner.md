# JunosAclTagsSpecsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortRange** | Pointer to **int32** | matched dst port, \&quot;0\&quot; means any | [optional] [default to 0]
**Protocol** | Pointer to **string** | &#x60;tcp&#x60; / &#x60;udp&#x60; / &#x60;icmp&#x60; / &#x60;gre&#x60; / &#x60;any&#x60; / &#x60;:protocol_number&#x60;. &#x60;protocol_number&#x60; is between 1-254 | [optional] [default to "any"]

## Methods

### NewJunosAclTagsSpecsInner

`func NewJunosAclTagsSpecsInner() *JunosAclTagsSpecsInner`

NewJunosAclTagsSpecsInner instantiates a new JunosAclTagsSpecsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosAclTagsSpecsInnerWithDefaults

`func NewJunosAclTagsSpecsInnerWithDefaults() *JunosAclTagsSpecsInner`

NewJunosAclTagsSpecsInnerWithDefaults instantiates a new JunosAclTagsSpecsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortRange

`func (o *JunosAclTagsSpecsInner) GetPortRange() int32`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *JunosAclTagsSpecsInner) GetPortRangeOk() (*int32, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *JunosAclTagsSpecsInner) SetPortRange(v int32)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *JunosAclTagsSpecsInner) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetProtocol

`func (o *JunosAclTagsSpecsInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *JunosAclTagsSpecsInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *JunosAclTagsSpecsInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *JunosAclTagsSpecsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

