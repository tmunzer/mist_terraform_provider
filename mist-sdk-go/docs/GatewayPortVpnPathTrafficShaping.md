# GatewayPortVpnPathTrafficShaping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassPercentage** | Pointer to **[]int32** | percentages for differet class of traffic: high / medium / low / best-effort sum must be equal to 100 | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MaxTxKbps** | Pointer to **int32** |  | [optional] 

## Methods

### NewGatewayPortVpnPathTrafficShaping

`func NewGatewayPortVpnPathTrafficShaping() *GatewayPortVpnPathTrafficShaping`

NewGatewayPortVpnPathTrafficShaping instantiates a new GatewayPortVpnPathTrafficShaping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPortVpnPathTrafficShapingWithDefaults

`func NewGatewayPortVpnPathTrafficShapingWithDefaults() *GatewayPortVpnPathTrafficShaping`

NewGatewayPortVpnPathTrafficShapingWithDefaults instantiates a new GatewayPortVpnPathTrafficShaping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassPercentage

`func (o *GatewayPortVpnPathTrafficShaping) GetClassPercentage() []int32`

GetClassPercentage returns the ClassPercentage field if non-nil, zero value otherwise.

### GetClassPercentageOk

`func (o *GatewayPortVpnPathTrafficShaping) GetClassPercentageOk() (*[]int32, bool)`

GetClassPercentageOk returns a tuple with the ClassPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassPercentage

`func (o *GatewayPortVpnPathTrafficShaping) SetClassPercentage(v []int32)`

SetClassPercentage sets ClassPercentage field to given value.

### HasClassPercentage

`func (o *GatewayPortVpnPathTrafficShaping) HasClassPercentage() bool`

HasClassPercentage returns a boolean if a field has been set.

### GetEnabled

`func (o *GatewayPortVpnPathTrafficShaping) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayPortVpnPathTrafficShaping) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayPortVpnPathTrafficShaping) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GatewayPortVpnPathTrafficShaping) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxTxKbps

`func (o *GatewayPortVpnPathTrafficShaping) GetMaxTxKbps() int32`

GetMaxTxKbps returns the MaxTxKbps field if non-nil, zero value otherwise.

### GetMaxTxKbpsOk

`func (o *GatewayPortVpnPathTrafficShaping) GetMaxTxKbpsOk() (*int32, bool)`

GetMaxTxKbpsOk returns a tuple with the MaxTxKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTxKbps

`func (o *GatewayPortVpnPathTrafficShaping) SetMaxTxKbps(v int32)`

SetMaxTxKbps sets MaxTxKbps field to given value.

### HasMaxTxKbps

`func (o *GatewayPortVpnPathTrafficShaping) HasMaxTxKbps() bool`

HasMaxTxKbps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


