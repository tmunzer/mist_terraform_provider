# JunosAclTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GbpTag** | Pointer to **float32** | required if - &#x60;type&#x60;&#x3D;&#x3D;&#x60;dynamic_gbp&#x60; (gbp_tag received from RADIUS) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; (applying gbp tag against matching conditions) | [optional] 
**Macs** | Pointer to **[]string** | required if  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;mac&#x60; - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching mac | [optional] 
**Network** | Pointer to **string** | if: - &#x60;type&#x60;&#x3D;&#x3D;&#x60;mac&#x60; (optional. default is &#x60;any&#x60;) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;subnet&#x60; (optional. default is &#x60;any&#x60;) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;network&#x60; - &#x60;type&#x60;&#x3D;&#x3D;&#x60;resource&#x60; (optional. default is &#x60;any&#x60;) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching network (vlan) | [optional] 
**RadiusGroup** | Pointer to **string** | required if  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius_group&#x60;  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching radius_group | [optional] 
**Specs** | Pointer to [**[]JunosAclTagsSpecsInner**](JunosAclTagsSpecsInner.md) | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;resource&#x60; empty means unrestricted, i.e. any | [optional] 
**Subnets** | Pointer to **[]string** | if  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;subnet&#x60;  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;resource&#x60; (optional. default is &#x60;any&#x60;) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching subnet | [optional] 
**Type** | [**JunosAclTagType**](JunosAclTagType.md) |  | 

## Methods

### NewJunosAclTags

`func NewJunosAclTags(type_ JunosAclTagType, ) *JunosAclTags`

NewJunosAclTags instantiates a new JunosAclTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosAclTagsWithDefaults

`func NewJunosAclTagsWithDefaults() *JunosAclTags`

NewJunosAclTagsWithDefaults instantiates a new JunosAclTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGbpTag

`func (o *JunosAclTags) GetGbpTag() float32`

GetGbpTag returns the GbpTag field if non-nil, zero value otherwise.

### GetGbpTagOk

`func (o *JunosAclTags) GetGbpTagOk() (*float32, bool)`

GetGbpTagOk returns a tuple with the GbpTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbpTag

`func (o *JunosAclTags) SetGbpTag(v float32)`

SetGbpTag sets GbpTag field to given value.

### HasGbpTag

`func (o *JunosAclTags) HasGbpTag() bool`

HasGbpTag returns a boolean if a field has been set.

### GetMacs

`func (o *JunosAclTags) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *JunosAclTags) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *JunosAclTags) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *JunosAclTags) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetNetwork

`func (o *JunosAclTags) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *JunosAclTags) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *JunosAclTags) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *JunosAclTags) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRadiusGroup

`func (o *JunosAclTags) GetRadiusGroup() string`

GetRadiusGroup returns the RadiusGroup field if non-nil, zero value otherwise.

### GetRadiusGroupOk

`func (o *JunosAclTags) GetRadiusGroupOk() (*string, bool)`

GetRadiusGroupOk returns a tuple with the RadiusGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroup

`func (o *JunosAclTags) SetRadiusGroup(v string)`

SetRadiusGroup sets RadiusGroup field to given value.

### HasRadiusGroup

`func (o *JunosAclTags) HasRadiusGroup() bool`

HasRadiusGroup returns a boolean if a field has been set.

### GetSpecs

`func (o *JunosAclTags) GetSpecs() []JunosAclTagsSpecsInner`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *JunosAclTags) GetSpecsOk() (*[]JunosAclTagsSpecsInner, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *JunosAclTags) SetSpecs(v []JunosAclTagsSpecsInner)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *JunosAclTags) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetSubnets

`func (o *JunosAclTags) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *JunosAclTags) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *JunosAclTags) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *JunosAclTags) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetType

`func (o *JunosAclTags) GetType() JunosAclTagType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosAclTags) GetTypeOk() (*JunosAclTagType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosAclTags) SetType(v JunosAclTagType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


