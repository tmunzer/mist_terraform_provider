# SwitchMatchingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalConfigCmds** | Pointer to **[]string** |  | [optional] 
**MatchRole** | Pointer to **string** | role to match | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PortConfig** | Pointer to [**map[string]JunosPortConfig**](JunosPortConfig.md) | Propery key is the interface name or interface range | [optional] 
**PortMirroring** | Pointer to [**map[string]SwitchMatchingRulePortMirroringValue**](SwitchMatchingRulePortMirroringValue.md) |  | [optional] 
**SwitchMgmt** | Pointer to [**ConfigSwitch**](ConfigSwitch.md) |  | [optional] 

## Methods

### NewSwitchMatchingRule

`func NewSwitchMatchingRule() *SwitchMatchingRule`

NewSwitchMatchingRule instantiates a new SwitchMatchingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchMatchingRuleWithDefaults

`func NewSwitchMatchingRuleWithDefaults() *SwitchMatchingRule`

NewSwitchMatchingRuleWithDefaults instantiates a new SwitchMatchingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalConfigCmds

`func (o *SwitchMatchingRule) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *SwitchMatchingRule) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *SwitchMatchingRule) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *SwitchMatchingRule) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetMatchRole

`func (o *SwitchMatchingRule) GetMatchRole() string`

GetMatchRole returns the MatchRole field if non-nil, zero value otherwise.

### GetMatchRoleOk

`func (o *SwitchMatchingRule) GetMatchRoleOk() (*string, bool)`

GetMatchRoleOk returns a tuple with the MatchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRole

`func (o *SwitchMatchingRule) SetMatchRole(v string)`

SetMatchRole sets MatchRole field to given value.

### HasMatchRole

`func (o *SwitchMatchingRule) HasMatchRole() bool`

HasMatchRole returns a boolean if a field has been set.

### GetName

`func (o *SwitchMatchingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwitchMatchingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwitchMatchingRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SwitchMatchingRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortConfig

`func (o *SwitchMatchingRule) GetPortConfig() map[string]JunosPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *SwitchMatchingRule) GetPortConfigOk() (*map[string]JunosPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *SwitchMatchingRule) SetPortConfig(v map[string]JunosPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *SwitchMatchingRule) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.

### GetPortMirroring

`func (o *SwitchMatchingRule) GetPortMirroring() map[string]SwitchMatchingRulePortMirroringValue`

GetPortMirroring returns the PortMirroring field if non-nil, zero value otherwise.

### GetPortMirroringOk

`func (o *SwitchMatchingRule) GetPortMirroringOk() (*map[string]SwitchMatchingRulePortMirroringValue, bool)`

GetPortMirroringOk returns a tuple with the PortMirroring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMirroring

`func (o *SwitchMatchingRule) SetPortMirroring(v map[string]SwitchMatchingRulePortMirroringValue)`

SetPortMirroring sets PortMirroring field to given value.

### HasPortMirroring

`func (o *SwitchMatchingRule) HasPortMirroring() bool`

HasPortMirroring returns a boolean if a field has been set.

### GetSwitchMgmt

`func (o *SwitchMatchingRule) GetSwitchMgmt() ConfigSwitch`

GetSwitchMgmt returns the SwitchMgmt field if non-nil, zero value otherwise.

### GetSwitchMgmtOk

`func (o *SwitchMatchingRule) GetSwitchMgmtOk() (*ConfigSwitch, bool)`

GetSwitchMgmtOk returns a tuple with the SwitchMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMgmt

`func (o *SwitchMatchingRule) SetSwitchMgmt(v ConfigSwitch)`

SetSwitchMgmt sets SwitchMgmt field to given value.

### HasSwitchMgmt

`func (o *SwitchMatchingRule) HasSwitchMgmt() bool`

HasSwitchMgmt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

