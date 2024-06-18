# GatewayPortConfigTrafficShaping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassPercentages** | Pointer to **[]int32** | percentages for differet class of traffic: high / medium / low / best-effort sum must be equal to 100 | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGatewayPortConfigTrafficShaping

`func NewGatewayPortConfigTrafficShaping() *GatewayPortConfigTrafficShaping`

NewGatewayPortConfigTrafficShaping instantiates a new GatewayPortConfigTrafficShaping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPortConfigTrafficShapingWithDefaults

`func NewGatewayPortConfigTrafficShapingWithDefaults() *GatewayPortConfigTrafficShaping`

NewGatewayPortConfigTrafficShapingWithDefaults instantiates a new GatewayPortConfigTrafficShaping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassPercentages

`func (o *GatewayPortConfigTrafficShaping) GetClassPercentages() []int32`

GetClassPercentages returns the ClassPercentages field if non-nil, zero value otherwise.

### GetClassPercentagesOk

`func (o *GatewayPortConfigTrafficShaping) GetClassPercentagesOk() (*[]int32, bool)`

GetClassPercentagesOk returns a tuple with the ClassPercentages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassPercentages

`func (o *GatewayPortConfigTrafficShaping) SetClassPercentages(v []int32)`

SetClassPercentages sets ClassPercentages field to given value.

### HasClassPercentages

`func (o *GatewayPortConfigTrafficShaping) HasClassPercentages() bool`

HasClassPercentages returns a boolean if a field has been set.

### GetEnabled

`func (o *GatewayPortConfigTrafficShaping) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayPortConfigTrafficShaping) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayPortConfigTrafficShaping) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GatewayPortConfigTrafficShaping) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


