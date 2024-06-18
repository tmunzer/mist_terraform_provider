# JunosNetworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **[]string** |  | [optional] 
**DnsSuffix** | Pointer to **[]string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Isolation** | Pointer to **bool** | whether to stop clients to talk to each other, default is false (when enabled, a unique isolation_vlan_id is required) NOTE: this features requires uplink device to also a be Juniper device and &#x60;inter_switch_link&#x60; to be set | [optional] [default to false]
**IsolationVlanId** | Pointer to **string** |  | [optional] 
**OspfInterfaceType** | Pointer to **string** |  | [optional] 
**Subnet** | Pointer to **string** | optional for pure switching, required when L3 / routing features are used | [optional] 
**VlanId** | **int32** |  | 
**Zone** | Pointer to **string** | used for gateway | [optional] 

## Methods

### NewJunosNetworks

`func NewJunosNetworks(vlanId int32, ) *JunosNetworks`

NewJunosNetworks instantiates a new JunosNetworks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosNetworksWithDefaults

`func NewJunosNetworksWithDefaults() *JunosNetworks`

NewJunosNetworksWithDefaults instantiates a new JunosNetworks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *JunosNetworks) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *JunosNetworks) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *JunosNetworks) SetDns(v []string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *JunosNetworks) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *JunosNetworks) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *JunosNetworks) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *JunosNetworks) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *JunosNetworks) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetGateway

`func (o *JunosNetworks) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *JunosNetworks) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *JunosNetworks) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *JunosNetworks) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIsolation

`func (o *JunosNetworks) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *JunosNetworks) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *JunosNetworks) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *JunosNetworks) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetIsolationVlanId

`func (o *JunosNetworks) GetIsolationVlanId() string`

GetIsolationVlanId returns the IsolationVlanId field if non-nil, zero value otherwise.

### GetIsolationVlanIdOk

`func (o *JunosNetworks) GetIsolationVlanIdOk() (*string, bool)`

GetIsolationVlanIdOk returns a tuple with the IsolationVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationVlanId

`func (o *JunosNetworks) SetIsolationVlanId(v string)`

SetIsolationVlanId sets IsolationVlanId field to given value.

### HasIsolationVlanId

`func (o *JunosNetworks) HasIsolationVlanId() bool`

HasIsolationVlanId returns a boolean if a field has been set.

### GetOspfInterfaceType

`func (o *JunosNetworks) GetOspfInterfaceType() string`

GetOspfInterfaceType returns the OspfInterfaceType field if non-nil, zero value otherwise.

### GetOspfInterfaceTypeOk

`func (o *JunosNetworks) GetOspfInterfaceTypeOk() (*string, bool)`

GetOspfInterfaceTypeOk returns a tuple with the OspfInterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfInterfaceType

`func (o *JunosNetworks) SetOspfInterfaceType(v string)`

SetOspfInterfaceType sets OspfInterfaceType field to given value.

### HasOspfInterfaceType

`func (o *JunosNetworks) HasOspfInterfaceType() bool`

HasOspfInterfaceType returns a boolean if a field has been set.

### GetSubnet

`func (o *JunosNetworks) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *JunosNetworks) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *JunosNetworks) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *JunosNetworks) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetVlanId

`func (o *JunosNetworks) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *JunosNetworks) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *JunosNetworks) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### GetZone

`func (o *JunosNetworks) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *JunosNetworks) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *JunosNetworks) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *JunosNetworks) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


