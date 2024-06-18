# JunosVrrpGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKey** | Pointer to **string** | if &#x60;auth_type&#x60;&#x3D;&#x3D;&#x60;md5&#x60; | [optional] 
**AuthPassword** | Pointer to **string** | if &#x60;auth_type&#x60;&#x3D;&#x3D;&#x60;simple&#x60; | [optional] 
**AuthType** | Pointer to **string** |  | [optional] [default to "md5"]
**Networks** | Pointer to [**map[string]JunosVrrpGroupNetworksValue**](JunosVrrpGroupNetworksValue.md) | Property key is the network name | [optional] 

## Methods

### NewJunosVrrpGroup

`func NewJunosVrrpGroup() *JunosVrrpGroup`

NewJunosVrrpGroup instantiates a new JunosVrrpGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosVrrpGroupWithDefaults

`func NewJunosVrrpGroupWithDefaults() *JunosVrrpGroup`

NewJunosVrrpGroupWithDefaults instantiates a new JunosVrrpGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKey

`func (o *JunosVrrpGroup) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *JunosVrrpGroup) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *JunosVrrpGroup) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *JunosVrrpGroup) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetAuthPassword

`func (o *JunosVrrpGroup) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *JunosVrrpGroup) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *JunosVrrpGroup) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *JunosVrrpGroup) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthType

`func (o *JunosVrrpGroup) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *JunosVrrpGroup) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *JunosVrrpGroup) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *JunosVrrpGroup) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetNetworks

`func (o *JunosVrrpGroup) GetNetworks() map[string]JunosVrrpGroupNetworksValue`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *JunosVrrpGroup) GetNetworksOk() (*map[string]JunosVrrpGroupNetworksValue, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *JunosVrrpGroup) SetNetworks(v map[string]JunosVrrpGroupNetworksValue)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *JunosVrrpGroup) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


