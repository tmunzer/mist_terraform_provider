# JunosEvpnTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ConfigSwitch**](ConfigSwitch.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Overwrite** | Pointer to **bool** |  | [optional] 
**PodNames** | Pointer to **map[string]string** | Property key is the pod number | [optional] 
**Switches** | [**[]JunosEvpnTopologySwitch**](JunosEvpnTopologySwitch.md) |  | 

## Methods

### NewJunosEvpnTopology

`func NewJunosEvpnTopology(switches []JunosEvpnTopologySwitch, ) *JunosEvpnTopology`

NewJunosEvpnTopology instantiates a new JunosEvpnTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosEvpnTopologyWithDefaults

`func NewJunosEvpnTopologyWithDefaults() *JunosEvpnTopology`

NewJunosEvpnTopologyWithDefaults instantiates a new JunosEvpnTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *JunosEvpnTopology) GetConfig() ConfigSwitch`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *JunosEvpnTopology) GetConfigOk() (*ConfigSwitch, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *JunosEvpnTopology) SetConfig(v ConfigSwitch)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *JunosEvpnTopology) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetId

`func (o *JunosEvpnTopology) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JunosEvpnTopology) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JunosEvpnTopology) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JunosEvpnTopology) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *JunosEvpnTopology) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JunosEvpnTopology) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JunosEvpnTopology) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JunosEvpnTopology) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverwrite

`func (o *JunosEvpnTopology) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *JunosEvpnTopology) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *JunosEvpnTopology) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *JunosEvpnTopology) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetPodNames

`func (o *JunosEvpnTopology) GetPodNames() map[string]string`

GetPodNames returns the PodNames field if non-nil, zero value otherwise.

### GetPodNamesOk

`func (o *JunosEvpnTopology) GetPodNamesOk() (*map[string]string, bool)`

GetPodNamesOk returns a tuple with the PodNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodNames

`func (o *JunosEvpnTopology) SetPodNames(v map[string]string)`

SetPodNames sets PodNames field to given value.

### HasPodNames

`func (o *JunosEvpnTopology) HasPodNames() bool`

HasPodNames returns a boolean if a field has been set.

### GetSwitches

`func (o *JunosEvpnTopology) GetSwitches() []JunosEvpnTopologySwitch`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *JunosEvpnTopology) GetSwitchesOk() (*[]JunosEvpnTopologySwitch, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *JunosEvpnTopology) SetSwitches(v []JunosEvpnTopologySwitch)`

SetSwitches sets Switches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


