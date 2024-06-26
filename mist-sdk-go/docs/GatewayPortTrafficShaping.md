# GatewayPortTrafficShaping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassPercentages** | Pointer to **[]int32** | percentages for differet class of traffic: high / medium / low / best-effort sum must be equal to 100 | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGatewayPortTrafficShaping

`func NewGatewayPortTrafficShaping() *GatewayPortTrafficShaping`

NewGatewayPortTrafficShaping instantiates a new GatewayPortTrafficShaping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPortTrafficShapingWithDefaults

`func NewGatewayPortTrafficShapingWithDefaults() *GatewayPortTrafficShaping`

NewGatewayPortTrafficShapingWithDefaults instantiates a new GatewayPortTrafficShaping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassPercentages

`func (o *GatewayPortTrafficShaping) GetClassPercentages() []int32`

GetClassPercentages returns the ClassPercentages field if non-nil, zero value otherwise.

### GetClassPercentagesOk

`func (o *GatewayPortTrafficShaping) GetClassPercentagesOk() (*[]int32, bool)`

GetClassPercentagesOk returns a tuple with the ClassPercentages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassPercentages

`func (o *GatewayPortTrafficShaping) SetClassPercentages(v []int32)`

SetClassPercentages sets ClassPercentages field to given value.

### HasClassPercentages

`func (o *GatewayPortTrafficShaping) HasClassPercentages() bool`

HasClassPercentages returns a boolean if a field has been set.

### GetEnabled

`func (o *GatewayPortTrafficShaping) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayPortTrafficShaping) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayPortTrafficShaping) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GatewayPortTrafficShaping) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


