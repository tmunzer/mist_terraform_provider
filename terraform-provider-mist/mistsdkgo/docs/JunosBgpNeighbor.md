# JunosBgpNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | If true, the BGP session to this neighbor will be administratively disabled/shutdown | [optional] [default to false]
**ExportPolicy** | Pointer to **string** |  | [optional] 
**HoldTime** | Pointer to **int32** |  | [optional] [default to 90]
**ImportPolicy** | Pointer to **string** |  | [optional] 
**MultihopTtl** | Pointer to **int32** | assuming BGP neighbor is directly connected | [optional] 
**NeighborAs** | Pointer to **int32** |  | [optional] 

## Methods

### NewJunosBgpNeighbor

`func NewJunosBgpNeighbor() *JunosBgpNeighbor`

NewJunosBgpNeighbor instantiates a new JunosBgpNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosBgpNeighborWithDefaults

`func NewJunosBgpNeighborWithDefaults() *JunosBgpNeighbor`

NewJunosBgpNeighborWithDefaults instantiates a new JunosBgpNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *JunosBgpNeighbor) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *JunosBgpNeighbor) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *JunosBgpNeighbor) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *JunosBgpNeighbor) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExportPolicy

`func (o *JunosBgpNeighbor) GetExportPolicy() string`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *JunosBgpNeighbor) GetExportPolicyOk() (*string, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *JunosBgpNeighbor) SetExportPolicy(v string)`

SetExportPolicy sets ExportPolicy field to given value.

### HasExportPolicy

`func (o *JunosBgpNeighbor) HasExportPolicy() bool`

HasExportPolicy returns a boolean if a field has been set.

### GetHoldTime

`func (o *JunosBgpNeighbor) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *JunosBgpNeighbor) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *JunosBgpNeighbor) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *JunosBgpNeighbor) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.

### GetImportPolicy

`func (o *JunosBgpNeighbor) GetImportPolicy() string`

GetImportPolicy returns the ImportPolicy field if non-nil, zero value otherwise.

### GetImportPolicyOk

`func (o *JunosBgpNeighbor) GetImportPolicyOk() (*string, bool)`

GetImportPolicyOk returns a tuple with the ImportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicy

`func (o *JunosBgpNeighbor) SetImportPolicy(v string)`

SetImportPolicy sets ImportPolicy field to given value.

### HasImportPolicy

`func (o *JunosBgpNeighbor) HasImportPolicy() bool`

HasImportPolicy returns a boolean if a field has been set.

### GetMultihopTtl

`func (o *JunosBgpNeighbor) GetMultihopTtl() int32`

GetMultihopTtl returns the MultihopTtl field if non-nil, zero value otherwise.

### GetMultihopTtlOk

`func (o *JunosBgpNeighbor) GetMultihopTtlOk() (*int32, bool)`

GetMultihopTtlOk returns a tuple with the MultihopTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultihopTtl

`func (o *JunosBgpNeighbor) SetMultihopTtl(v int32)`

SetMultihopTtl sets MultihopTtl field to given value.

### HasMultihopTtl

`func (o *JunosBgpNeighbor) HasMultihopTtl() bool`

HasMultihopTtl returns a boolean if a field has been set.

### GetNeighborAs

`func (o *JunosBgpNeighbor) GetNeighborAs() int32`

GetNeighborAs returns the NeighborAs field if non-nil, zero value otherwise.

### GetNeighborAsOk

`func (o *JunosBgpNeighbor) GetNeighborAsOk() (*int32, bool)`

GetNeighborAsOk returns a tuple with the NeighborAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborAs

`func (o *JunosBgpNeighbor) SetNeighborAs(v int32)`

SetNeighborAs sets NeighborAs field to given value.

### HasNeighborAs

`func (o *JunosBgpNeighbor) HasNeighborAs() bool`

HasNeighborAs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


