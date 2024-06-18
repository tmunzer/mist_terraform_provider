# GatewayTemplateTunnelPrimary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** |  | [optional] 
**InternalIps** | Pointer to **[]string** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;zscaler-gre&#x60;  * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-gre&#x60; | [optional] 
**ProbeIps** | Pointer to **[]string** |  | [optional] 
**RemoteIds** | Pointer to **[]string** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**WanNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGatewayTemplateTunnelPrimary

`func NewGatewayTemplateTunnelPrimary() *GatewayTemplateTunnelPrimary`

NewGatewayTemplateTunnelPrimary instantiates a new GatewayTemplateTunnelPrimary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelPrimaryWithDefaults

`func NewGatewayTemplateTunnelPrimaryWithDefaults() *GatewayTemplateTunnelPrimary`

NewGatewayTemplateTunnelPrimaryWithDefaults instantiates a new GatewayTemplateTunnelPrimary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *GatewayTemplateTunnelPrimary) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *GatewayTemplateTunnelPrimary) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *GatewayTemplateTunnelPrimary) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *GatewayTemplateTunnelPrimary) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetInternalIps

`func (o *GatewayTemplateTunnelPrimary) GetInternalIps() []string`

GetInternalIps returns the InternalIps field if non-nil, zero value otherwise.

### GetInternalIpsOk

`func (o *GatewayTemplateTunnelPrimary) GetInternalIpsOk() (*[]string, bool)`

GetInternalIpsOk returns a tuple with the InternalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIps

`func (o *GatewayTemplateTunnelPrimary) SetInternalIps(v []string)`

SetInternalIps sets InternalIps field to given value.

### HasInternalIps

`func (o *GatewayTemplateTunnelPrimary) HasInternalIps() bool`

HasInternalIps returns a boolean if a field has been set.

### GetProbeIps

`func (o *GatewayTemplateTunnelPrimary) GetProbeIps() []string`

GetProbeIps returns the ProbeIps field if non-nil, zero value otherwise.

### GetProbeIpsOk

`func (o *GatewayTemplateTunnelPrimary) GetProbeIpsOk() (*[]string, bool)`

GetProbeIpsOk returns a tuple with the ProbeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeIps

`func (o *GatewayTemplateTunnelPrimary) SetProbeIps(v []string)`

SetProbeIps sets ProbeIps field to given value.

### HasProbeIps

`func (o *GatewayTemplateTunnelPrimary) HasProbeIps() bool`

HasProbeIps returns a boolean if a field has been set.

### GetRemoteIds

`func (o *GatewayTemplateTunnelPrimary) GetRemoteIds() []string`

GetRemoteIds returns the RemoteIds field if non-nil, zero value otherwise.

### GetRemoteIdsOk

`func (o *GatewayTemplateTunnelPrimary) GetRemoteIdsOk() (*[]string, bool)`

GetRemoteIdsOk returns a tuple with the RemoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIds

`func (o *GatewayTemplateTunnelPrimary) SetRemoteIds(v []string)`

SetRemoteIds sets RemoteIds field to given value.

### HasRemoteIds

`func (o *GatewayTemplateTunnelPrimary) HasRemoteIds() bool`

HasRemoteIds returns a boolean if a field has been set.

### GetWanNames

`func (o *GatewayTemplateTunnelPrimary) GetWanNames() []string`

GetWanNames returns the WanNames field if non-nil, zero value otherwise.

### GetWanNamesOk

`func (o *GatewayTemplateTunnelPrimary) GetWanNamesOk() (*[]string, bool)`

GetWanNamesOk returns a tuple with the WanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanNames

`func (o *GatewayTemplateTunnelPrimary) SetWanNames(v []string)`

SetWanNames sets WanNames field to given value.

### HasWanNames

`func (o *GatewayTemplateTunnelPrimary) HasWanNames() bool`

HasWanNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


