# GatewayPortConfigVpnPathsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BfdProfile** | Pointer to **string** |  | [optional] [default to "broadband"]
**BfdUseTunnelMode** | Pointer to **bool** | whether to use tunnel mode. SSR only | [optional] [default to false]
**Role** | Pointer to **string** |  | [optional] [default to "spoke"]
**TrafficShaping** | Pointer to [**GatewayPortConfigVpnPathsValueTrafficShaping**](GatewayPortConfigVpnPathsValueTrafficShaping.md) |  | [optional] 

## Methods

### NewGatewayPortConfigVpnPathsValue

`func NewGatewayPortConfigVpnPathsValue() *GatewayPortConfigVpnPathsValue`

NewGatewayPortConfigVpnPathsValue instantiates a new GatewayPortConfigVpnPathsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPortConfigVpnPathsValueWithDefaults

`func NewGatewayPortConfigVpnPathsValueWithDefaults() *GatewayPortConfigVpnPathsValue`

NewGatewayPortConfigVpnPathsValueWithDefaults instantiates a new GatewayPortConfigVpnPathsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBfdProfile

`func (o *GatewayPortConfigVpnPathsValue) GetBfdProfile() string`

GetBfdProfile returns the BfdProfile field if non-nil, zero value otherwise.

### GetBfdProfileOk

`func (o *GatewayPortConfigVpnPathsValue) GetBfdProfileOk() (*string, bool)`

GetBfdProfileOk returns a tuple with the BfdProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdProfile

`func (o *GatewayPortConfigVpnPathsValue) SetBfdProfile(v string)`

SetBfdProfile sets BfdProfile field to given value.

### HasBfdProfile

`func (o *GatewayPortConfigVpnPathsValue) HasBfdProfile() bool`

HasBfdProfile returns a boolean if a field has been set.

### GetBfdUseTunnelMode

`func (o *GatewayPortConfigVpnPathsValue) GetBfdUseTunnelMode() bool`

GetBfdUseTunnelMode returns the BfdUseTunnelMode field if non-nil, zero value otherwise.

### GetBfdUseTunnelModeOk

`func (o *GatewayPortConfigVpnPathsValue) GetBfdUseTunnelModeOk() (*bool, bool)`

GetBfdUseTunnelModeOk returns a tuple with the BfdUseTunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdUseTunnelMode

`func (o *GatewayPortConfigVpnPathsValue) SetBfdUseTunnelMode(v bool)`

SetBfdUseTunnelMode sets BfdUseTunnelMode field to given value.

### HasBfdUseTunnelMode

`func (o *GatewayPortConfigVpnPathsValue) HasBfdUseTunnelMode() bool`

HasBfdUseTunnelMode returns a boolean if a field has been set.

### GetRole

`func (o *GatewayPortConfigVpnPathsValue) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GatewayPortConfigVpnPathsValue) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GatewayPortConfigVpnPathsValue) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GatewayPortConfigVpnPathsValue) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTrafficShaping

`func (o *GatewayPortConfigVpnPathsValue) GetTrafficShaping() GatewayPortConfigVpnPathsValueTrafficShaping`

GetTrafficShaping returns the TrafficShaping field if non-nil, zero value otherwise.

### GetTrafficShapingOk

`func (o *GatewayPortConfigVpnPathsValue) GetTrafficShapingOk() (*GatewayPortConfigVpnPathsValueTrafficShaping, bool)`

GetTrafficShapingOk returns a tuple with the TrafficShaping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficShaping

`func (o *GatewayPortConfigVpnPathsValue) SetTrafficShaping(v GatewayPortConfigVpnPathsValueTrafficShaping)`

SetTrafficShaping sets TrafficShaping field to given value.

### HasTrafficShaping

`func (o *GatewayPortConfigVpnPathsValue) HasTrafficShaping() bool`

HasTrafficShaping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


