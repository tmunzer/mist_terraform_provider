# GatewayPortVpnPaths

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preference** | Pointer to **int32** | for a given VPN, when path_selection.strategy&#x3D;simple, the preference for a path (lower is preferred) | [optional] 

## Methods

### NewGatewayPortVpnPaths

`func NewGatewayPortVpnPaths() *GatewayPortVpnPaths`

NewGatewayPortVpnPaths instantiates a new GatewayPortVpnPaths object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPortVpnPathsWithDefaults

`func NewGatewayPortVpnPathsWithDefaults() *GatewayPortVpnPaths`

NewGatewayPortVpnPathsWithDefaults instantiates a new GatewayPortVpnPaths object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreference

`func (o *GatewayPortVpnPaths) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *GatewayPortVpnPaths) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *GatewayPortVpnPaths) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *GatewayPortVpnPaths) HasPreference() bool`

HasPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


