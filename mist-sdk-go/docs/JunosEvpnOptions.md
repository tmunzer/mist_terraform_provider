# JunosEvpnOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLoopbackSubnet** | Pointer to **string** | optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server-id-overrides | [optional] 
**AutoLoopbackSubnet6** | Pointer to **string** | optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server-id-overrides | [optional] 
**AutoRouterIdSubnet** | Pointer to **string** | optional, this generates router_id automatically, if specified, &#x60;router_id_prefix&#x60; is ignored | [optional] 
**AutoRouterIdSubnet6** | Pointer to **string** | optional, this generates router_id automatically, if specified, &#x60;router_id_prefix&#x60; is ignored | [optional] [default to "fd31:5700:1::/64"]
**CoreAsBorder** | Pointer to **bool** | optional, for ERB or CLOS, you can either use esilag to upstream routers or to also be the virtual-gateway when &#x60;routed_at&#x60; !&#x3D; &#x60;core&#x60;, whether to do virtual-gateway at core as well | [optional] [default to false]
**Overlay** | Pointer to [**JunosEvpnOptionsOverlay**](JunosEvpnOptionsOverlay.md) |  | [optional] 
**PerVlanVgaV4Mac** | Pointer to **bool** | by default, JUNOS uses 00-00-5e-00-01-01 as the virtual-gateway-address&#39;s v4-mac if enabled, 00-00-5e-00-XX-YY will be used (where XX&#x3D;vlan_id/256, YY&#x3D;vlan_id%256) | [optional] [default to false]
**RoutedAt** | Pointer to [**JunosEvpnOptionsRoutedAt**](JunosEvpnOptionsRoutedAt.md) |  | [optional] [default to JUNOSEVPNOPTIONSROUTEDAT_EDGE]
**Underlay** | Pointer to [**JunosEvpnOptionsUnderlay**](JunosEvpnOptionsUnderlay.md) |  | [optional] 
**VsInstances** | Pointer to [**map[string]JunosEvpnOptionsVsInstancesValue**](JunosEvpnOptionsVsInstancesValue.md) | optional, for EX9200 only to seggregate virtual-switches | [optional] 

## Methods

### NewJunosEvpnOptions

`func NewJunosEvpnOptions() *JunosEvpnOptions`

NewJunosEvpnOptions instantiates a new JunosEvpnOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosEvpnOptionsWithDefaults

`func NewJunosEvpnOptionsWithDefaults() *JunosEvpnOptions`

NewJunosEvpnOptionsWithDefaults instantiates a new JunosEvpnOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoLoopbackSubnet

`func (o *JunosEvpnOptions) GetAutoLoopbackSubnet() string`

GetAutoLoopbackSubnet returns the AutoLoopbackSubnet field if non-nil, zero value otherwise.

### GetAutoLoopbackSubnetOk

`func (o *JunosEvpnOptions) GetAutoLoopbackSubnetOk() (*string, bool)`

GetAutoLoopbackSubnetOk returns a tuple with the AutoLoopbackSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLoopbackSubnet

`func (o *JunosEvpnOptions) SetAutoLoopbackSubnet(v string)`

SetAutoLoopbackSubnet sets AutoLoopbackSubnet field to given value.

### HasAutoLoopbackSubnet

`func (o *JunosEvpnOptions) HasAutoLoopbackSubnet() bool`

HasAutoLoopbackSubnet returns a boolean if a field has been set.

### GetAutoLoopbackSubnet6

`func (o *JunosEvpnOptions) GetAutoLoopbackSubnet6() string`

GetAutoLoopbackSubnet6 returns the AutoLoopbackSubnet6 field if non-nil, zero value otherwise.

### GetAutoLoopbackSubnet6Ok

`func (o *JunosEvpnOptions) GetAutoLoopbackSubnet6Ok() (*string, bool)`

GetAutoLoopbackSubnet6Ok returns a tuple with the AutoLoopbackSubnet6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLoopbackSubnet6

`func (o *JunosEvpnOptions) SetAutoLoopbackSubnet6(v string)`

SetAutoLoopbackSubnet6 sets AutoLoopbackSubnet6 field to given value.

### HasAutoLoopbackSubnet6

`func (o *JunosEvpnOptions) HasAutoLoopbackSubnet6() bool`

HasAutoLoopbackSubnet6 returns a boolean if a field has been set.

### GetAutoRouterIdSubnet

`func (o *JunosEvpnOptions) GetAutoRouterIdSubnet() string`

GetAutoRouterIdSubnet returns the AutoRouterIdSubnet field if non-nil, zero value otherwise.

### GetAutoRouterIdSubnetOk

`func (o *JunosEvpnOptions) GetAutoRouterIdSubnetOk() (*string, bool)`

GetAutoRouterIdSubnetOk returns a tuple with the AutoRouterIdSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouterIdSubnet

`func (o *JunosEvpnOptions) SetAutoRouterIdSubnet(v string)`

SetAutoRouterIdSubnet sets AutoRouterIdSubnet field to given value.

### HasAutoRouterIdSubnet

`func (o *JunosEvpnOptions) HasAutoRouterIdSubnet() bool`

HasAutoRouterIdSubnet returns a boolean if a field has been set.

### GetAutoRouterIdSubnet6

`func (o *JunosEvpnOptions) GetAutoRouterIdSubnet6() string`

GetAutoRouterIdSubnet6 returns the AutoRouterIdSubnet6 field if non-nil, zero value otherwise.

### GetAutoRouterIdSubnet6Ok

`func (o *JunosEvpnOptions) GetAutoRouterIdSubnet6Ok() (*string, bool)`

GetAutoRouterIdSubnet6Ok returns a tuple with the AutoRouterIdSubnet6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouterIdSubnet6

`func (o *JunosEvpnOptions) SetAutoRouterIdSubnet6(v string)`

SetAutoRouterIdSubnet6 sets AutoRouterIdSubnet6 field to given value.

### HasAutoRouterIdSubnet6

`func (o *JunosEvpnOptions) HasAutoRouterIdSubnet6() bool`

HasAutoRouterIdSubnet6 returns a boolean if a field has been set.

### GetCoreAsBorder

`func (o *JunosEvpnOptions) GetCoreAsBorder() bool`

GetCoreAsBorder returns the CoreAsBorder field if non-nil, zero value otherwise.

### GetCoreAsBorderOk

`func (o *JunosEvpnOptions) GetCoreAsBorderOk() (*bool, bool)`

GetCoreAsBorderOk returns a tuple with the CoreAsBorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAsBorder

`func (o *JunosEvpnOptions) SetCoreAsBorder(v bool)`

SetCoreAsBorder sets CoreAsBorder field to given value.

### HasCoreAsBorder

`func (o *JunosEvpnOptions) HasCoreAsBorder() bool`

HasCoreAsBorder returns a boolean if a field has been set.

### GetOverlay

`func (o *JunosEvpnOptions) GetOverlay() JunosEvpnOptionsOverlay`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *JunosEvpnOptions) GetOverlayOk() (*JunosEvpnOptionsOverlay, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *JunosEvpnOptions) SetOverlay(v JunosEvpnOptionsOverlay)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *JunosEvpnOptions) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### GetPerVlanVgaV4Mac

`func (o *JunosEvpnOptions) GetPerVlanVgaV4Mac() bool`

GetPerVlanVgaV4Mac returns the PerVlanVgaV4Mac field if non-nil, zero value otherwise.

### GetPerVlanVgaV4MacOk

`func (o *JunosEvpnOptions) GetPerVlanVgaV4MacOk() (*bool, bool)`

GetPerVlanVgaV4MacOk returns a tuple with the PerVlanVgaV4Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerVlanVgaV4Mac

`func (o *JunosEvpnOptions) SetPerVlanVgaV4Mac(v bool)`

SetPerVlanVgaV4Mac sets PerVlanVgaV4Mac field to given value.

### HasPerVlanVgaV4Mac

`func (o *JunosEvpnOptions) HasPerVlanVgaV4Mac() bool`

HasPerVlanVgaV4Mac returns a boolean if a field has been set.

### GetRoutedAt

`func (o *JunosEvpnOptions) GetRoutedAt() JunosEvpnOptionsRoutedAt`

GetRoutedAt returns the RoutedAt field if non-nil, zero value otherwise.

### GetRoutedAtOk

`func (o *JunosEvpnOptions) GetRoutedAtOk() (*JunosEvpnOptionsRoutedAt, bool)`

GetRoutedAtOk returns a tuple with the RoutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutedAt

`func (o *JunosEvpnOptions) SetRoutedAt(v JunosEvpnOptionsRoutedAt)`

SetRoutedAt sets RoutedAt field to given value.

### HasRoutedAt

`func (o *JunosEvpnOptions) HasRoutedAt() bool`

HasRoutedAt returns a boolean if a field has been set.

### GetUnderlay

`func (o *JunosEvpnOptions) GetUnderlay() JunosEvpnOptionsUnderlay`

GetUnderlay returns the Underlay field if non-nil, zero value otherwise.

### GetUnderlayOk

`func (o *JunosEvpnOptions) GetUnderlayOk() (*JunosEvpnOptionsUnderlay, bool)`

GetUnderlayOk returns a tuple with the Underlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlay

`func (o *JunosEvpnOptions) SetUnderlay(v JunosEvpnOptionsUnderlay)`

SetUnderlay sets Underlay field to given value.

### HasUnderlay

`func (o *JunosEvpnOptions) HasUnderlay() bool`

HasUnderlay returns a boolean if a field has been set.

### GetVsInstances

`func (o *JunosEvpnOptions) GetVsInstances() map[string]JunosEvpnOptionsVsInstancesValue`

GetVsInstances returns the VsInstances field if non-nil, zero value otherwise.

### GetVsInstancesOk

`func (o *JunosEvpnOptions) GetVsInstancesOk() (*map[string]JunosEvpnOptionsVsInstancesValue, bool)`

GetVsInstancesOk returns a tuple with the VsInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsInstances

`func (o *JunosEvpnOptions) SetVsInstances(v map[string]JunosEvpnOptionsVsInstancesValue)`

SetVsInstances sets VsInstances field to given value.

### HasVsInstances

`func (o *JunosEvpnOptions) HasVsInstances() bool`

HasVsInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


