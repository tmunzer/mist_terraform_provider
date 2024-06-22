# VrfInstancesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraRoutes** | Pointer to [**map[string]VrfExtraRoutesValue**](VrfExtraRoutesValue.md) | Property key is the destination CIDR (e.g. \&quot;10.0.0.0/8\&quot;) | [optional] 
**Networks** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewVrfInstancesConfig

`func NewVrfInstancesConfig() *VrfInstancesConfig`

NewVrfInstancesConfig instantiates a new VrfInstancesConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfInstancesConfigWithDefaults

`func NewVrfInstancesConfigWithDefaults() *VrfInstancesConfig`

NewVrfInstancesConfigWithDefaults instantiates a new VrfInstancesConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraRoutes

`func (o *VrfInstancesConfig) GetExtraRoutes() map[string]VrfExtraRoutesValue`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *VrfInstancesConfig) GetExtraRoutesOk() (*map[string]VrfExtraRoutesValue, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *VrfInstancesConfig) SetExtraRoutes(v map[string]VrfExtraRoutesValue)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *VrfInstancesConfig) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetNetworks

`func (o *VrfInstancesConfig) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VrfInstancesConfig) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VrfInstancesConfig) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VrfInstancesConfig) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


