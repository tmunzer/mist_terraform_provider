# GatewayTemplateTunnelSecondary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** |  | [optional] 
**InternalIps** | Pointer to **[]string** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**ProbeIps** | Pointer to **[]string** |  | [optional] [default to []]
**RemoteIds** | Pointer to **[]string** | Only if:  * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**WanNames** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewGatewayTemplateTunnelSecondary

`func NewGatewayTemplateTunnelSecondary() *GatewayTemplateTunnelSecondary`

NewGatewayTemplateTunnelSecondary instantiates a new GatewayTemplateTunnelSecondary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelSecondaryWithDefaults

`func NewGatewayTemplateTunnelSecondaryWithDefaults() *GatewayTemplateTunnelSecondary`

NewGatewayTemplateTunnelSecondaryWithDefaults instantiates a new GatewayTemplateTunnelSecondary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *GatewayTemplateTunnelSecondary) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *GatewayTemplateTunnelSecondary) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *GatewayTemplateTunnelSecondary) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *GatewayTemplateTunnelSecondary) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetInternalIps

`func (o *GatewayTemplateTunnelSecondary) GetInternalIps() []string`

GetInternalIps returns the InternalIps field if non-nil, zero value otherwise.

### GetInternalIpsOk

`func (o *GatewayTemplateTunnelSecondary) GetInternalIpsOk() (*[]string, bool)`

GetInternalIpsOk returns a tuple with the InternalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIps

`func (o *GatewayTemplateTunnelSecondary) SetInternalIps(v []string)`

SetInternalIps sets InternalIps field to given value.

### HasInternalIps

`func (o *GatewayTemplateTunnelSecondary) HasInternalIps() bool`

HasInternalIps returns a boolean if a field has been set.

### GetProbeIps

`func (o *GatewayTemplateTunnelSecondary) GetProbeIps() []string`

GetProbeIps returns the ProbeIps field if non-nil, zero value otherwise.

### GetProbeIpsOk

`func (o *GatewayTemplateTunnelSecondary) GetProbeIpsOk() (*[]string, bool)`

GetProbeIpsOk returns a tuple with the ProbeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeIps

`func (o *GatewayTemplateTunnelSecondary) SetProbeIps(v []string)`

SetProbeIps sets ProbeIps field to given value.

### HasProbeIps

`func (o *GatewayTemplateTunnelSecondary) HasProbeIps() bool`

HasProbeIps returns a boolean if a field has been set.

### GetRemoteIds

`func (o *GatewayTemplateTunnelSecondary) GetRemoteIds() []string`

GetRemoteIds returns the RemoteIds field if non-nil, zero value otherwise.

### GetRemoteIdsOk

`func (o *GatewayTemplateTunnelSecondary) GetRemoteIdsOk() (*[]string, bool)`

GetRemoteIdsOk returns a tuple with the RemoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIds

`func (o *GatewayTemplateTunnelSecondary) SetRemoteIds(v []string)`

SetRemoteIds sets RemoteIds field to given value.

### HasRemoteIds

`func (o *GatewayTemplateTunnelSecondary) HasRemoteIds() bool`

HasRemoteIds returns a boolean if a field has been set.

### GetWanNames

`func (o *GatewayTemplateTunnelSecondary) GetWanNames() []string`

GetWanNames returns the WanNames field if non-nil, zero value otherwise.

### GetWanNamesOk

`func (o *GatewayTemplateTunnelSecondary) GetWanNamesOk() (*[]string, bool)`

GetWanNamesOk returns a tuple with the WanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanNames

`func (o *GatewayTemplateTunnelSecondary) SetWanNames(v []string)`

SetWanNames sets WanNames field to given value.

### HasWanNames

`func (o *GatewayTemplateTunnelSecondary) HasWanNames() bool`

HasWanNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


