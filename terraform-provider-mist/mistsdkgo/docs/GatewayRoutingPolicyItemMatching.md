# GatewayRoutingPolicyItemMatching

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsPath** | Pointer to **[]string** | takes regular expression | [optional] 
**Community** | Pointer to **[]string** |  | [optional] 
**Network** | Pointer to **[]string** |  | [optional] [default to []]
**Prefix** | Pointer to **[]string** | zero or more criteria/filter can be specified to match the term, all criteria have to be met | [optional] 
**Protocol** | Pointer to **[]string** | &#x60;direct&#x60;, &#x60;bgp&#x60;, &#x60;osp&#x60;, ... | [optional] 
**RouteExists** | Pointer to [**GatewayRoutingPolicyItemMatchingRouteExists**](GatewayRoutingPolicyItemMatchingRouteExists.md) |  | [optional] 
**VpnNeighborMac** | Pointer to **[]string** | overlay-facing criteria (used for bgp_config where via&#x3D;vpn) | [optional] 
**VpnPath** | Pointer to **[]string** | overlay-facing criteria (used for bgp_config where via&#x3D;vpn) ordered- | [optional] 
**VpnPathSla** | Pointer to [**GatewayRoutingPolicyItemMatchingVpnPathSla**](GatewayRoutingPolicyItemMatchingVpnPathSla.md) |  | [optional] 

## Methods

### NewGatewayRoutingPolicyItemMatching

`func NewGatewayRoutingPolicyItemMatching() *GatewayRoutingPolicyItemMatching`

NewGatewayRoutingPolicyItemMatching instantiates a new GatewayRoutingPolicyItemMatching object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRoutingPolicyItemMatchingWithDefaults

`func NewGatewayRoutingPolicyItemMatchingWithDefaults() *GatewayRoutingPolicyItemMatching`

NewGatewayRoutingPolicyItemMatchingWithDefaults instantiates a new GatewayRoutingPolicyItemMatching object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsPath

`func (o *GatewayRoutingPolicyItemMatching) GetAsPath() []string`

GetAsPath returns the AsPath field if non-nil, zero value otherwise.

### GetAsPathOk

`func (o *GatewayRoutingPolicyItemMatching) GetAsPathOk() (*[]string, bool)`

GetAsPathOk returns a tuple with the AsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsPath

`func (o *GatewayRoutingPolicyItemMatching) SetAsPath(v []string)`

SetAsPath sets AsPath field to given value.

### HasAsPath

`func (o *GatewayRoutingPolicyItemMatching) HasAsPath() bool`

HasAsPath returns a boolean if a field has been set.

### GetCommunity

`func (o *GatewayRoutingPolicyItemMatching) GetCommunity() []string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *GatewayRoutingPolicyItemMatching) GetCommunityOk() (*[]string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *GatewayRoutingPolicyItemMatching) SetCommunity(v []string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *GatewayRoutingPolicyItemMatching) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetNetwork

`func (o *GatewayRoutingPolicyItemMatching) GetNetwork() []string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayRoutingPolicyItemMatching) GetNetworkOk() (*[]string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayRoutingPolicyItemMatching) SetNetwork(v []string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GatewayRoutingPolicyItemMatching) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPrefix

`func (o *GatewayRoutingPolicyItemMatching) GetPrefix() []string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GatewayRoutingPolicyItemMatching) GetPrefixOk() (*[]string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GatewayRoutingPolicyItemMatching) SetPrefix(v []string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GatewayRoutingPolicyItemMatching) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProtocol

`func (o *GatewayRoutingPolicyItemMatching) GetProtocol() []string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayRoutingPolicyItemMatching) GetProtocolOk() (*[]string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayRoutingPolicyItemMatching) SetProtocol(v []string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GatewayRoutingPolicyItemMatching) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRouteExists

`func (o *GatewayRoutingPolicyItemMatching) GetRouteExists() GatewayRoutingPolicyItemMatchingRouteExists`

GetRouteExists returns the RouteExists field if non-nil, zero value otherwise.

### GetRouteExistsOk

`func (o *GatewayRoutingPolicyItemMatching) GetRouteExistsOk() (*GatewayRoutingPolicyItemMatchingRouteExists, bool)`

GetRouteExistsOk returns a tuple with the RouteExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteExists

`func (o *GatewayRoutingPolicyItemMatching) SetRouteExists(v GatewayRoutingPolicyItemMatchingRouteExists)`

SetRouteExists sets RouteExists field to given value.

### HasRouteExists

`func (o *GatewayRoutingPolicyItemMatching) HasRouteExists() bool`

HasRouteExists returns a boolean if a field has been set.

### GetVpnNeighborMac

`func (o *GatewayRoutingPolicyItemMatching) GetVpnNeighborMac() []string`

GetVpnNeighborMac returns the VpnNeighborMac field if non-nil, zero value otherwise.

### GetVpnNeighborMacOk

`func (o *GatewayRoutingPolicyItemMatching) GetVpnNeighborMacOk() (*[]string, bool)`

GetVpnNeighborMacOk returns a tuple with the VpnNeighborMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnNeighborMac

`func (o *GatewayRoutingPolicyItemMatching) SetVpnNeighborMac(v []string)`

SetVpnNeighborMac sets VpnNeighborMac field to given value.

### HasVpnNeighborMac

`func (o *GatewayRoutingPolicyItemMatching) HasVpnNeighborMac() bool`

HasVpnNeighborMac returns a boolean if a field has been set.

### GetVpnPath

`func (o *GatewayRoutingPolicyItemMatching) GetVpnPath() []string`

GetVpnPath returns the VpnPath field if non-nil, zero value otherwise.

### GetVpnPathOk

`func (o *GatewayRoutingPolicyItemMatching) GetVpnPathOk() (*[]string, bool)`

GetVpnPathOk returns a tuple with the VpnPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPath

`func (o *GatewayRoutingPolicyItemMatching) SetVpnPath(v []string)`

SetVpnPath sets VpnPath field to given value.

### HasVpnPath

`func (o *GatewayRoutingPolicyItemMatching) HasVpnPath() bool`

HasVpnPath returns a boolean if a field has been set.

### GetVpnPathSla

`func (o *GatewayRoutingPolicyItemMatching) GetVpnPathSla() GatewayRoutingPolicyItemMatchingVpnPathSla`

GetVpnPathSla returns the VpnPathSla field if non-nil, zero value otherwise.

### GetVpnPathSlaOk

`func (o *GatewayRoutingPolicyItemMatching) GetVpnPathSlaOk() (*GatewayRoutingPolicyItemMatchingVpnPathSla, bool)`

GetVpnPathSlaOk returns a tuple with the VpnPathSla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPathSla

`func (o *GatewayRoutingPolicyItemMatching) SetVpnPathSla(v GatewayRoutingPolicyItemMatchingVpnPathSla)`

SetVpnPathSla sets VpnPathSla field to given value.

### HasVpnPathSla

`func (o *GatewayRoutingPolicyItemMatching) HasVpnPathSla() bool`

HasVpnPathSla returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


