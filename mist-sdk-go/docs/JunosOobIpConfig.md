# JunosOobIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **[]string** |  | [optional] [default to []]
**DnsSuffix** | Pointer to **[]string** |  | [optional] [default to []]
**Gateway** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] 
**Network** | Pointer to **string** | optional, the network to be used for mgmt | [optional] 
**Node1** | Pointer to [**JunosOobIpConfigNode1**](JunosOobIpConfigNode1.md) |  | [optional] 
**Type** | Pointer to [**IpConfigType**](IpConfigType.md) |  | [optional] [default to IPCONFIGTYPE_DYNAMIC]
**UseMgmtVrf** | Pointer to **bool** | if supported on the platform. If enabled, DNS will be using this routing-instance, too | [optional] [default to true]
**UseMgmtVrfForHostOut** | Pointer to **bool** | whether to use &#x60;mgmt_junos&#x60; for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired | [optional] [default to false]
**VlanId** | Pointer to **int32** |  | [optional] 

## Methods

### NewJunosOobIpConfig

`func NewJunosOobIpConfig() *JunosOobIpConfig`

NewJunosOobIpConfig instantiates a new JunosOobIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosOobIpConfigWithDefaults

`func NewJunosOobIpConfigWithDefaults() *JunosOobIpConfig`

NewJunosOobIpConfigWithDefaults instantiates a new JunosOobIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *JunosOobIpConfig) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *JunosOobIpConfig) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *JunosOobIpConfig) SetDns(v []string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *JunosOobIpConfig) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *JunosOobIpConfig) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *JunosOobIpConfig) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *JunosOobIpConfig) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *JunosOobIpConfig) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetGateway

`func (o *JunosOobIpConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *JunosOobIpConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *JunosOobIpConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *JunosOobIpConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIp

`func (o *JunosOobIpConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *JunosOobIpConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *JunosOobIpConfig) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *JunosOobIpConfig) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *JunosOobIpConfig) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *JunosOobIpConfig) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *JunosOobIpConfig) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *JunosOobIpConfig) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *JunosOobIpConfig) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *JunosOobIpConfig) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *JunosOobIpConfig) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *JunosOobIpConfig) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNode1

`func (o *JunosOobIpConfig) GetNode1() JunosOobIpConfigNode1`

GetNode1 returns the Node1 field if non-nil, zero value otherwise.

### GetNode1Ok

`func (o *JunosOobIpConfig) GetNode1Ok() (*JunosOobIpConfigNode1, bool)`

GetNode1Ok returns a tuple with the Node1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode1

`func (o *JunosOobIpConfig) SetNode1(v JunosOobIpConfigNode1)`

SetNode1 sets Node1 field to given value.

### HasNode1

`func (o *JunosOobIpConfig) HasNode1() bool`

HasNode1 returns a boolean if a field has been set.

### GetType

`func (o *JunosOobIpConfig) GetType() IpConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosOobIpConfig) GetTypeOk() (*IpConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosOobIpConfig) SetType(v IpConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *JunosOobIpConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUseMgmtVrf

`func (o *JunosOobIpConfig) GetUseMgmtVrf() bool`

GetUseMgmtVrf returns the UseMgmtVrf field if non-nil, zero value otherwise.

### GetUseMgmtVrfOk

`func (o *JunosOobIpConfig) GetUseMgmtVrfOk() (*bool, bool)`

GetUseMgmtVrfOk returns a tuple with the UseMgmtVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrf

`func (o *JunosOobIpConfig) SetUseMgmtVrf(v bool)`

SetUseMgmtVrf sets UseMgmtVrf field to given value.

### HasUseMgmtVrf

`func (o *JunosOobIpConfig) HasUseMgmtVrf() bool`

HasUseMgmtVrf returns a boolean if a field has been set.

### GetUseMgmtVrfForHostOut

`func (o *JunosOobIpConfig) GetUseMgmtVrfForHostOut() bool`

GetUseMgmtVrfForHostOut returns the UseMgmtVrfForHostOut field if non-nil, zero value otherwise.

### GetUseMgmtVrfForHostOutOk

`func (o *JunosOobIpConfig) GetUseMgmtVrfForHostOutOk() (*bool, bool)`

GetUseMgmtVrfForHostOutOk returns a tuple with the UseMgmtVrfForHostOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrfForHostOut

`func (o *JunosOobIpConfig) SetUseMgmtVrfForHostOut(v bool)`

SetUseMgmtVrfForHostOut sets UseMgmtVrfForHostOut field to given value.

### HasUseMgmtVrfForHostOut

`func (o *JunosOobIpConfig) HasUseMgmtVrfForHostOut() bool`

HasUseMgmtVrfForHostOut returns a boolean if a field has been set.

### GetVlanId

`func (o *JunosOobIpConfig) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *JunosOobIpConfig) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *JunosOobIpConfig) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *JunosOobIpConfig) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


