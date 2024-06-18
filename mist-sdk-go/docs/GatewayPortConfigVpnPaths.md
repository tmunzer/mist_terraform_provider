# GatewayPortConfigVpnPaths

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preference** | Pointer to **int32** | for a given VPN, when path_selection.strategy&#x3D;simple, the preference for a path (lower is preferred) | [optional] 

## Methods

### NewGatewayPortConfigVpnPaths

`func NewGatewayPortConfigVpnPaths() *GatewayPortConfigVpnPaths`

NewGatewayPortConfigVpnPaths instantiates a new GatewayPortConfigVpnPaths object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPortConfigVpnPathsWithDefaults

`func NewGatewayPortConfigVpnPathsWithDefaults() *GatewayPortConfigVpnPaths`

NewGatewayPortConfigVpnPathsWithDefaults instantiates a new GatewayPortConfigVpnPaths object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreference

`func (o *GatewayPortConfigVpnPaths) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *GatewayPortConfigVpnPaths) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *GatewayPortConfigVpnPaths) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *GatewayPortConfigVpnPaths) HasPreference() bool`

HasPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


