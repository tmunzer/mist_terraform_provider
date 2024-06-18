# JunosOspfConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Areas** | Pointer to [**map[string]JunosOspfConfigAreasValue**](JunosOspfConfigAreasValue.md) | OSPF areas to run on this device and the corresponding per-area-specific configs. Property key is the area | [optional] 
**Enabled** | Pointer to **bool** | whether to rung OSPF on this device | [optional] 
**ReferenceBandwidth** | Pointer to **string** | Bandwidth for calculating metric defaults (9600..4000000000000) | [optional] [default to "100M"]

## Methods

### NewJunosOspfConfig

`func NewJunosOspfConfig() *JunosOspfConfig`

NewJunosOspfConfig instantiates a new JunosOspfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosOspfConfigWithDefaults

`func NewJunosOspfConfigWithDefaults() *JunosOspfConfig`

NewJunosOspfConfigWithDefaults instantiates a new JunosOspfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreas

`func (o *JunosOspfConfig) GetAreas() map[string]JunosOspfConfigAreasValue`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *JunosOspfConfig) GetAreasOk() (*map[string]JunosOspfConfigAreasValue, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *JunosOspfConfig) SetAreas(v map[string]JunosOspfConfigAreasValue)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *JunosOspfConfig) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetEnabled

`func (o *JunosOspfConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JunosOspfConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JunosOspfConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JunosOspfConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReferenceBandwidth

`func (o *JunosOspfConfig) GetReferenceBandwidth() string`

GetReferenceBandwidth returns the ReferenceBandwidth field if non-nil, zero value otherwise.

### GetReferenceBandwidthOk

`func (o *JunosOspfConfig) GetReferenceBandwidthOk() (*string, bool)`

GetReferenceBandwidthOk returns a tuple with the ReferenceBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceBandwidth

`func (o *JunosOspfConfig) SetReferenceBandwidth(v string)`

SetReferenceBandwidth sets ReferenceBandwidth field to given value.

### HasReferenceBandwidth

`func (o *JunosOspfConfig) HasReferenceBandwidth() bool`

HasReferenceBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


