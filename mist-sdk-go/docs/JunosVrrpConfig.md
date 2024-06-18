# JunosVrrpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Groups** | Pointer to [**JunosVrrpConfigGroups**](JunosVrrpConfigGroups.md) |  | [optional] 

## Methods

### NewJunosVrrpConfig

`func NewJunosVrrpConfig() *JunosVrrpConfig`

NewJunosVrrpConfig instantiates a new JunosVrrpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosVrrpConfigWithDefaults

`func NewJunosVrrpConfigWithDefaults() *JunosVrrpConfig`

NewJunosVrrpConfigWithDefaults instantiates a new JunosVrrpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *JunosVrrpConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JunosVrrpConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JunosVrrpConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JunosVrrpConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroups

`func (o *JunosVrrpConfig) GetGroups() JunosVrrpConfigGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *JunosVrrpConfig) GetGroupsOk() (*JunosVrrpConfigGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *JunosVrrpConfig) SetGroups(v JunosVrrpConfigGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *JunosVrrpConfig) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


