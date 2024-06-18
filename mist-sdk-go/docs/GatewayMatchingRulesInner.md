# GatewayMatchingRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalConfigCmds** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PortConfig** | Pointer to [**map[string]JunosPortConfig**](JunosPortConfig.md) |  | [optional] 

## Methods

### NewGatewayMatchingRulesInner

`func NewGatewayMatchingRulesInner() *GatewayMatchingRulesInner`

NewGatewayMatchingRulesInner instantiates a new GatewayMatchingRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayMatchingRulesInnerWithDefaults

`func NewGatewayMatchingRulesInnerWithDefaults() *GatewayMatchingRulesInner`

NewGatewayMatchingRulesInnerWithDefaults instantiates a new GatewayMatchingRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalConfigCmds

`func (o *GatewayMatchingRulesInner) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *GatewayMatchingRulesInner) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *GatewayMatchingRulesInner) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *GatewayMatchingRulesInner) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetName

`func (o *GatewayMatchingRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayMatchingRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayMatchingRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayMatchingRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortConfig

`func (o *GatewayMatchingRulesInner) GetPortConfig() map[string]JunosPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *GatewayMatchingRulesInner) GetPortConfigOk() (*map[string]JunosPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *GatewayMatchingRulesInner) SetPortConfig(v map[string]JunosPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *GatewayMatchingRulesInner) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


