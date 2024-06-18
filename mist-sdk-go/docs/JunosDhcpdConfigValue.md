# JunosDhcpdConfigValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | when defined in template, this allows device to override | [optional] [default to false]
**DnsServers** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; - optional, if not defined, system one will be used | [optional] 
**DnsSuffix** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; - optional, if not defined, system one will be used | [optional] 
**FixedBindings** | Pointer to [**map[string]JunosDhcpdConfigFixedBindingsValue**](JunosDhcpdConfigFixedBindingsValue.md) | Property key is the MAC Address | [optional] 
**Gateway** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; - optional, &#x60;ip&#x60; will be used if not provided | [optional] 
**IpEnd** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; | [optional] 
**IpStart** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; | [optional] 
**LeaseTime** | Pointer to **int32** | in seconds, lease time has to be between 3600 [1hr] - 604800 [1 week], default is 86400 [1 day] | [optional] [default to 86400]
**Options** | Pointer to [**map[string]JunosDhcpdConfigOptionsValue**](JunosDhcpdConfigOptionsValue.md) | Property key is the DHCP option number | [optional] 
**ServerIdOverride** | Pointer to **bool** | &#x60;server_id_override&#x60;&#x3D;&#x3D;&#x60;true&#x60; means the device, when acts as DHCP relay and forwards DHCP responses from DHCP server to clients,  should overwrite the Sever Identifier option (i.e. DHCP option 54) in DHCP responses with its own IP address. | [optional] [default to false]
**Servers** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;relay&#x60; | [optional] 
**Type** | Pointer to **string** | DHCP Server (local) or DHCP Relay (relay) | [optional] [default to "local"]
**VendorEncapulated** | Pointer to [**map[string]JunosDhcpdConfigOptionsValue**](JunosDhcpdConfigOptionsValue.md) | Property key is &lt;enterprise number&gt;:&lt;sub option code&gt;, with * enterprise number: 1-65535 (https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers) * sub option code: 1-255, sub-option code | [optional] 

## Methods

### NewJunosDhcpdConfigValue

`func NewJunosDhcpdConfigValue() *JunosDhcpdConfigValue`

NewJunosDhcpdConfigValue instantiates a new JunosDhcpdConfigValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosDhcpdConfigValueWithDefaults

`func NewJunosDhcpdConfigValueWithDefaults() *JunosDhcpdConfigValue`

NewJunosDhcpdConfigValueWithDefaults instantiates a new JunosDhcpdConfigValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *JunosDhcpdConfigValue) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *JunosDhcpdConfigValue) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *JunosDhcpdConfigValue) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *JunosDhcpdConfigValue) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDnsServers

`func (o *JunosDhcpdConfigValue) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *JunosDhcpdConfigValue) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *JunosDhcpdConfigValue) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *JunosDhcpdConfigValue) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *JunosDhcpdConfigValue) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *JunosDhcpdConfigValue) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *JunosDhcpdConfigValue) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *JunosDhcpdConfigValue) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetFixedBindings

`func (o *JunosDhcpdConfigValue) GetFixedBindings() map[string]JunosDhcpdConfigFixedBindingsValue`

GetFixedBindings returns the FixedBindings field if non-nil, zero value otherwise.

### GetFixedBindingsOk

`func (o *JunosDhcpdConfigValue) GetFixedBindingsOk() (*map[string]JunosDhcpdConfigFixedBindingsValue, bool)`

GetFixedBindingsOk returns a tuple with the FixedBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedBindings

`func (o *JunosDhcpdConfigValue) SetFixedBindings(v map[string]JunosDhcpdConfigFixedBindingsValue)`

SetFixedBindings sets FixedBindings field to given value.

### HasFixedBindings

`func (o *JunosDhcpdConfigValue) HasFixedBindings() bool`

HasFixedBindings returns a boolean if a field has been set.

### GetGateway

`func (o *JunosDhcpdConfigValue) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *JunosDhcpdConfigValue) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *JunosDhcpdConfigValue) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *JunosDhcpdConfigValue) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpEnd

`func (o *JunosDhcpdConfigValue) GetIpEnd() string`

GetIpEnd returns the IpEnd field if non-nil, zero value otherwise.

### GetIpEndOk

`func (o *JunosDhcpdConfigValue) GetIpEndOk() (*string, bool)`

GetIpEndOk returns a tuple with the IpEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpEnd

`func (o *JunosDhcpdConfigValue) SetIpEnd(v string)`

SetIpEnd sets IpEnd field to given value.

### HasIpEnd

`func (o *JunosDhcpdConfigValue) HasIpEnd() bool`

HasIpEnd returns a boolean if a field has been set.

### GetIpStart

`func (o *JunosDhcpdConfigValue) GetIpStart() string`

GetIpStart returns the IpStart field if non-nil, zero value otherwise.

### GetIpStartOk

`func (o *JunosDhcpdConfigValue) GetIpStartOk() (*string, bool)`

GetIpStartOk returns a tuple with the IpStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStart

`func (o *JunosDhcpdConfigValue) SetIpStart(v string)`

SetIpStart sets IpStart field to given value.

### HasIpStart

`func (o *JunosDhcpdConfigValue) HasIpStart() bool`

HasIpStart returns a boolean if a field has been set.

### GetLeaseTime

`func (o *JunosDhcpdConfigValue) GetLeaseTime() int32`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *JunosDhcpdConfigValue) GetLeaseTimeOk() (*int32, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *JunosDhcpdConfigValue) SetLeaseTime(v int32)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *JunosDhcpdConfigValue) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetOptions

`func (o *JunosDhcpdConfigValue) GetOptions() map[string]JunosDhcpdConfigOptionsValue`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *JunosDhcpdConfigValue) GetOptionsOk() (*map[string]JunosDhcpdConfigOptionsValue, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *JunosDhcpdConfigValue) SetOptions(v map[string]JunosDhcpdConfigOptionsValue)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *JunosDhcpdConfigValue) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetServerIdOverride

`func (o *JunosDhcpdConfigValue) GetServerIdOverride() bool`

GetServerIdOverride returns the ServerIdOverride field if non-nil, zero value otherwise.

### GetServerIdOverrideOk

`func (o *JunosDhcpdConfigValue) GetServerIdOverrideOk() (*bool, bool)`

GetServerIdOverrideOk returns a tuple with the ServerIdOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIdOverride

`func (o *JunosDhcpdConfigValue) SetServerIdOverride(v bool)`

SetServerIdOverride sets ServerIdOverride field to given value.

### HasServerIdOverride

`func (o *JunosDhcpdConfigValue) HasServerIdOverride() bool`

HasServerIdOverride returns a boolean if a field has been set.

### GetServers

`func (o *JunosDhcpdConfigValue) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *JunosDhcpdConfigValue) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *JunosDhcpdConfigValue) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *JunosDhcpdConfigValue) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetType

`func (o *JunosDhcpdConfigValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosDhcpdConfigValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosDhcpdConfigValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JunosDhcpdConfigValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendorEncapulated

`func (o *JunosDhcpdConfigValue) GetVendorEncapulated() map[string]JunosDhcpdConfigOptionsValue`

GetVendorEncapulated returns the VendorEncapulated field if non-nil, zero value otherwise.

### GetVendorEncapulatedOk

`func (o *JunosDhcpdConfigValue) GetVendorEncapulatedOk() (*map[string]JunosDhcpdConfigOptionsValue, bool)`

GetVendorEncapulatedOk returns a tuple with the VendorEncapulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorEncapulated

`func (o *JunosDhcpdConfigValue) SetVendorEncapulated(v map[string]JunosDhcpdConfigOptionsValue)`

SetVendorEncapulated sets VendorEncapulated field to given value.

### HasVendorEncapulated

`func (o *JunosDhcpdConfigValue) HasVendorEncapulated() bool`

HasVendorEncapulated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


