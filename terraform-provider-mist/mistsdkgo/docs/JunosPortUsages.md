# JunosPortUsages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllNetworks** | Pointer to **bool** | if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;trunk&#x60;, whether to trunk all network/vlans | [optional] [default to false]
**AllowDhcpd** | Pointer to **bool** | if DHCP snooping is enabled, whether DHCP server is allowed on the interfaces with. All the interfaces from port configs using this port usage are effected. Please notice that allow_dhcpd is a tri-state.  When it is not defined, it means using the system’s default setting which depends on whether the port is a access or trunk port. | [optional] 
**AllowMultipleSupplicants** | Pointer to **bool** |  | [optional] [default to false]
**BypassAuthWhenServerDown** | Pointer to **bool** | if &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60;, bypass auth for known clients if set to true when RADIUS server is down | [optional] [default to false]
**BypassAuthWhenServerDownForUnkonwnClient** | Pointer to **bool** | if &#x60;port_auth&#x60;&#x3D;&#x60;dot1x&#x60;, bypass auth for all (including unknown clients) if set to true when RADIUS server is down | [optional] [default to false]
**Description** | Pointer to **string** | description | [optional] 
**DisableAutoneg** | Pointer to **bool** | if speed and duplex are specified, whether to disable autonegotiation | [optional] [default to false]
**Disabled** | Pointer to **bool** | whether the port is disabled | [optional] [default to false]
**Duplex** | Pointer to **string** | link connection mode, choices are auto (default), full, and half | [optional] [default to "auto"]
**DynamicVlanNetworks** | Pointer to **[]string** | if dynamic vlan is used, specify the possible networks/vlans RADIUS can return | [optional] 
**EnableMacAuth** | Pointer to **bool** | if &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60;, whether to enable MAC Auth | [optional] [default to false]
**EnableQos** | Pointer to **bool** |  | [optional] [default to false]
**GuestNetwork** | Pointer to **NullableString** | if &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60;, which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed) | [optional] 
**MacAuthOnly** | Pointer to **bool** | only effect once &#x60;enable_mac_auth&#x60;&#x3D;&#x3D;&#x60;true&#x60; | [optional] 
**MacAuthProtocol** | Pointer to **string** | if &#x60;enable_mac_auth&#x60; &#x3D;&#x3D;&#x60;true&#x60;. This type is ignored if mist_nac is enabled. | [optional] [default to "eap-md5"]
**MacLimit** | Pointer to **int32** | max number of mac addresses, default is 0 for unlimited, otherwise range is 1 or higher, with upper bound constrained by platform | [optional] [default to 0]
**Mode** | Pointer to **string** | access (default) / trunk | [optional] 
**Mtu** | Pointer to **int32** | media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514. | [optional] 
**Networks** | Pointer to **[]string** | if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;trunk&#x60;, the list of network/vlans | [optional] 
**PersistMac** | Pointer to **bool** | if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;access&#x60; and &#x60;port_auth&#x60;!&#x3D;&#x60;dot1x&#x60;, whether the port should retain dynamically learned MAC addresses | [optional] [default to false]
**PoeDisabled** | Pointer to **bool** | whether PoE capabilities are disabled for a port | [optional] [default to false]
**PortAuth** | Pointer to **string** | if dot1x is desired, set to dot1x | [optional] 
**PortNetwork** | Pointer to **string** | native network/vlan for untagged traffic | [optional] 
**ReauthInterval** | Pointer to **int32** | if &#x60;port_auth&#x60;&#x3D;&#x60;dot1x&#x60;, reauthentication interval range | [optional] [default to 3600]
**RejectedNetwork** | Pointer to **NullableString** | if &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60;, when radius server reject / fails | [optional] 
**Speed** | Pointer to **string** | speed, default is auto to automatically negotiate speed | [optional] 
**StormControl** | Pointer to [**JunosStormControl**](JunosStormControl.md) |  | [optional] 
**StpEdge** | Pointer to **bool** | when enabled, the port is not expected to receive BPDU frames | [optional] [default to false]
**VoipNetwork** | Pointer to **string** | network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth | [optional] 

## Methods

### NewJunosPortUsages

`func NewJunosPortUsages() *JunosPortUsages`

NewJunosPortUsages instantiates a new JunosPortUsages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosPortUsagesWithDefaults

`func NewJunosPortUsagesWithDefaults() *JunosPortUsages`

NewJunosPortUsagesWithDefaults instantiates a new JunosPortUsages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllNetworks

`func (o *JunosPortUsages) GetAllNetworks() bool`

GetAllNetworks returns the AllNetworks field if non-nil, zero value otherwise.

### GetAllNetworksOk

`func (o *JunosPortUsages) GetAllNetworksOk() (*bool, bool)`

GetAllNetworksOk returns a tuple with the AllNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllNetworks

`func (o *JunosPortUsages) SetAllNetworks(v bool)`

SetAllNetworks sets AllNetworks field to given value.

### HasAllNetworks

`func (o *JunosPortUsages) HasAllNetworks() bool`

HasAllNetworks returns a boolean if a field has been set.

### GetAllowDhcpd

`func (o *JunosPortUsages) GetAllowDhcpd() bool`

GetAllowDhcpd returns the AllowDhcpd field if non-nil, zero value otherwise.

### GetAllowDhcpdOk

`func (o *JunosPortUsages) GetAllowDhcpdOk() (*bool, bool)`

GetAllowDhcpdOk returns a tuple with the AllowDhcpd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDhcpd

`func (o *JunosPortUsages) SetAllowDhcpd(v bool)`

SetAllowDhcpd sets AllowDhcpd field to given value.

### HasAllowDhcpd

`func (o *JunosPortUsages) HasAllowDhcpd() bool`

HasAllowDhcpd returns a boolean if a field has been set.

### GetAllowMultipleSupplicants

`func (o *JunosPortUsages) GetAllowMultipleSupplicants() bool`

GetAllowMultipleSupplicants returns the AllowMultipleSupplicants field if non-nil, zero value otherwise.

### GetAllowMultipleSupplicantsOk

`func (o *JunosPortUsages) GetAllowMultipleSupplicantsOk() (*bool, bool)`

GetAllowMultipleSupplicantsOk returns a tuple with the AllowMultipleSupplicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSupplicants

`func (o *JunosPortUsages) SetAllowMultipleSupplicants(v bool)`

SetAllowMultipleSupplicants sets AllowMultipleSupplicants field to given value.

### HasAllowMultipleSupplicants

`func (o *JunosPortUsages) HasAllowMultipleSupplicants() bool`

HasAllowMultipleSupplicants returns a boolean if a field has been set.

### GetBypassAuthWhenServerDown

`func (o *JunosPortUsages) GetBypassAuthWhenServerDown() bool`

GetBypassAuthWhenServerDown returns the BypassAuthWhenServerDown field if non-nil, zero value otherwise.

### GetBypassAuthWhenServerDownOk

`func (o *JunosPortUsages) GetBypassAuthWhenServerDownOk() (*bool, bool)`

GetBypassAuthWhenServerDownOk returns a tuple with the BypassAuthWhenServerDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAuthWhenServerDown

`func (o *JunosPortUsages) SetBypassAuthWhenServerDown(v bool)`

SetBypassAuthWhenServerDown sets BypassAuthWhenServerDown field to given value.

### HasBypassAuthWhenServerDown

`func (o *JunosPortUsages) HasBypassAuthWhenServerDown() bool`

HasBypassAuthWhenServerDown returns a boolean if a field has been set.

### GetBypassAuthWhenServerDownForUnkonwnClient

`func (o *JunosPortUsages) GetBypassAuthWhenServerDownForUnkonwnClient() bool`

GetBypassAuthWhenServerDownForUnkonwnClient returns the BypassAuthWhenServerDownForUnkonwnClient field if non-nil, zero value otherwise.

### GetBypassAuthWhenServerDownForUnkonwnClientOk

`func (o *JunosPortUsages) GetBypassAuthWhenServerDownForUnkonwnClientOk() (*bool, bool)`

GetBypassAuthWhenServerDownForUnkonwnClientOk returns a tuple with the BypassAuthWhenServerDownForUnkonwnClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAuthWhenServerDownForUnkonwnClient

`func (o *JunosPortUsages) SetBypassAuthWhenServerDownForUnkonwnClient(v bool)`

SetBypassAuthWhenServerDownForUnkonwnClient sets BypassAuthWhenServerDownForUnkonwnClient field to given value.

### HasBypassAuthWhenServerDownForUnkonwnClient

`func (o *JunosPortUsages) HasBypassAuthWhenServerDownForUnkonwnClient() bool`

HasBypassAuthWhenServerDownForUnkonwnClient returns a boolean if a field has been set.

### GetDescription

`func (o *JunosPortUsages) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JunosPortUsages) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JunosPortUsages) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JunosPortUsages) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableAutoneg

`func (o *JunosPortUsages) GetDisableAutoneg() bool`

GetDisableAutoneg returns the DisableAutoneg field if non-nil, zero value otherwise.

### GetDisableAutonegOk

`func (o *JunosPortUsages) GetDisableAutonegOk() (*bool, bool)`

GetDisableAutonegOk returns a tuple with the DisableAutoneg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoneg

`func (o *JunosPortUsages) SetDisableAutoneg(v bool)`

SetDisableAutoneg sets DisableAutoneg field to given value.

### HasDisableAutoneg

`func (o *JunosPortUsages) HasDisableAutoneg() bool`

HasDisableAutoneg returns a boolean if a field has been set.

### GetDisabled

`func (o *JunosPortUsages) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *JunosPortUsages) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *JunosPortUsages) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *JunosPortUsages) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDuplex

`func (o *JunosPortUsages) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *JunosPortUsages) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *JunosPortUsages) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *JunosPortUsages) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetDynamicVlanNetworks

`func (o *JunosPortUsages) GetDynamicVlanNetworks() []string`

GetDynamicVlanNetworks returns the DynamicVlanNetworks field if non-nil, zero value otherwise.

### GetDynamicVlanNetworksOk

`func (o *JunosPortUsages) GetDynamicVlanNetworksOk() (*[]string, bool)`

GetDynamicVlanNetworksOk returns a tuple with the DynamicVlanNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVlanNetworks

`func (o *JunosPortUsages) SetDynamicVlanNetworks(v []string)`

SetDynamicVlanNetworks sets DynamicVlanNetworks field to given value.

### HasDynamicVlanNetworks

`func (o *JunosPortUsages) HasDynamicVlanNetworks() bool`

HasDynamicVlanNetworks returns a boolean if a field has been set.

### GetEnableMacAuth

`func (o *JunosPortUsages) GetEnableMacAuth() bool`

GetEnableMacAuth returns the EnableMacAuth field if non-nil, zero value otherwise.

### GetEnableMacAuthOk

`func (o *JunosPortUsages) GetEnableMacAuthOk() (*bool, bool)`

GetEnableMacAuthOk returns a tuple with the EnableMacAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMacAuth

`func (o *JunosPortUsages) SetEnableMacAuth(v bool)`

SetEnableMacAuth sets EnableMacAuth field to given value.

### HasEnableMacAuth

`func (o *JunosPortUsages) HasEnableMacAuth() bool`

HasEnableMacAuth returns a boolean if a field has been set.

### GetEnableQos

`func (o *JunosPortUsages) GetEnableQos() bool`

GetEnableQos returns the EnableQos field if non-nil, zero value otherwise.

### GetEnableQosOk

`func (o *JunosPortUsages) GetEnableQosOk() (*bool, bool)`

GetEnableQosOk returns a tuple with the EnableQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQos

`func (o *JunosPortUsages) SetEnableQos(v bool)`

SetEnableQos sets EnableQos field to given value.

### HasEnableQos

`func (o *JunosPortUsages) HasEnableQos() bool`

HasEnableQos returns a boolean if a field has been set.

### GetGuestNetwork

`func (o *JunosPortUsages) GetGuestNetwork() string`

GetGuestNetwork returns the GuestNetwork field if non-nil, zero value otherwise.

### GetGuestNetworkOk

`func (o *JunosPortUsages) GetGuestNetworkOk() (*string, bool)`

GetGuestNetworkOk returns a tuple with the GuestNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNetwork

`func (o *JunosPortUsages) SetGuestNetwork(v string)`

SetGuestNetwork sets GuestNetwork field to given value.

### HasGuestNetwork

`func (o *JunosPortUsages) HasGuestNetwork() bool`

HasGuestNetwork returns a boolean if a field has been set.

### SetGuestNetworkNil

`func (o *JunosPortUsages) SetGuestNetworkNil(b bool)`

 SetGuestNetworkNil sets the value for GuestNetwork to be an explicit nil

### UnsetGuestNetwork
`func (o *JunosPortUsages) UnsetGuestNetwork()`

UnsetGuestNetwork ensures that no value is present for GuestNetwork, not even an explicit nil
### GetMacAuthOnly

`func (o *JunosPortUsages) GetMacAuthOnly() bool`

GetMacAuthOnly returns the MacAuthOnly field if non-nil, zero value otherwise.

### GetMacAuthOnlyOk

`func (o *JunosPortUsages) GetMacAuthOnlyOk() (*bool, bool)`

GetMacAuthOnlyOk returns a tuple with the MacAuthOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAuthOnly

`func (o *JunosPortUsages) SetMacAuthOnly(v bool)`

SetMacAuthOnly sets MacAuthOnly field to given value.

### HasMacAuthOnly

`func (o *JunosPortUsages) HasMacAuthOnly() bool`

HasMacAuthOnly returns a boolean if a field has been set.

### GetMacAuthProtocol

`func (o *JunosPortUsages) GetMacAuthProtocol() string`

GetMacAuthProtocol returns the MacAuthProtocol field if non-nil, zero value otherwise.

### GetMacAuthProtocolOk

`func (o *JunosPortUsages) GetMacAuthProtocolOk() (*string, bool)`

GetMacAuthProtocolOk returns a tuple with the MacAuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAuthProtocol

`func (o *JunosPortUsages) SetMacAuthProtocol(v string)`

SetMacAuthProtocol sets MacAuthProtocol field to given value.

### HasMacAuthProtocol

`func (o *JunosPortUsages) HasMacAuthProtocol() bool`

HasMacAuthProtocol returns a boolean if a field has been set.

### GetMacLimit

`func (o *JunosPortUsages) GetMacLimit() int32`

GetMacLimit returns the MacLimit field if non-nil, zero value otherwise.

### GetMacLimitOk

`func (o *JunosPortUsages) GetMacLimitOk() (*int32, bool)`

GetMacLimitOk returns a tuple with the MacLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLimit

`func (o *JunosPortUsages) SetMacLimit(v int32)`

SetMacLimit sets MacLimit field to given value.

### HasMacLimit

`func (o *JunosPortUsages) HasMacLimit() bool`

HasMacLimit returns a boolean if a field has been set.

### GetMode

`func (o *JunosPortUsages) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *JunosPortUsages) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *JunosPortUsages) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *JunosPortUsages) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMtu

`func (o *JunosPortUsages) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *JunosPortUsages) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *JunosPortUsages) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *JunosPortUsages) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNetworks

`func (o *JunosPortUsages) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *JunosPortUsages) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *JunosPortUsages) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *JunosPortUsages) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetPersistMac

`func (o *JunosPortUsages) GetPersistMac() bool`

GetPersistMac returns the PersistMac field if non-nil, zero value otherwise.

### GetPersistMacOk

`func (o *JunosPortUsages) GetPersistMacOk() (*bool, bool)`

GetPersistMacOk returns a tuple with the PersistMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistMac

`func (o *JunosPortUsages) SetPersistMac(v bool)`

SetPersistMac sets PersistMac field to given value.

### HasPersistMac

`func (o *JunosPortUsages) HasPersistMac() bool`

HasPersistMac returns a boolean if a field has been set.

### GetPoeDisabled

`func (o *JunosPortUsages) GetPoeDisabled() bool`

GetPoeDisabled returns the PoeDisabled field if non-nil, zero value otherwise.

### GetPoeDisabledOk

`func (o *JunosPortUsages) GetPoeDisabledOk() (*bool, bool)`

GetPoeDisabledOk returns a tuple with the PoeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisabled

`func (o *JunosPortUsages) SetPoeDisabled(v bool)`

SetPoeDisabled sets PoeDisabled field to given value.

### HasPoeDisabled

`func (o *JunosPortUsages) HasPoeDisabled() bool`

HasPoeDisabled returns a boolean if a field has been set.

### GetPortAuth

`func (o *JunosPortUsages) GetPortAuth() string`

GetPortAuth returns the PortAuth field if non-nil, zero value otherwise.

### GetPortAuthOk

`func (o *JunosPortUsages) GetPortAuthOk() (*string, bool)`

GetPortAuthOk returns a tuple with the PortAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAuth

`func (o *JunosPortUsages) SetPortAuth(v string)`

SetPortAuth sets PortAuth field to given value.

### HasPortAuth

`func (o *JunosPortUsages) HasPortAuth() bool`

HasPortAuth returns a boolean if a field has been set.

### GetPortNetwork

`func (o *JunosPortUsages) GetPortNetwork() string`

GetPortNetwork returns the PortNetwork field if non-nil, zero value otherwise.

### GetPortNetworkOk

`func (o *JunosPortUsages) GetPortNetworkOk() (*string, bool)`

GetPortNetworkOk returns a tuple with the PortNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNetwork

`func (o *JunosPortUsages) SetPortNetwork(v string)`

SetPortNetwork sets PortNetwork field to given value.

### HasPortNetwork

`func (o *JunosPortUsages) HasPortNetwork() bool`

HasPortNetwork returns a boolean if a field has been set.

### GetReauthInterval

`func (o *JunosPortUsages) GetReauthInterval() int32`

GetReauthInterval returns the ReauthInterval field if non-nil, zero value otherwise.

### GetReauthIntervalOk

`func (o *JunosPortUsages) GetReauthIntervalOk() (*int32, bool)`

GetReauthIntervalOk returns a tuple with the ReauthInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthInterval

`func (o *JunosPortUsages) SetReauthInterval(v int32)`

SetReauthInterval sets ReauthInterval field to given value.

### HasReauthInterval

`func (o *JunosPortUsages) HasReauthInterval() bool`

HasReauthInterval returns a boolean if a field has been set.

### GetRejectedNetwork

`func (o *JunosPortUsages) GetRejectedNetwork() string`

GetRejectedNetwork returns the RejectedNetwork field if non-nil, zero value otherwise.

### GetRejectedNetworkOk

`func (o *JunosPortUsages) GetRejectedNetworkOk() (*string, bool)`

GetRejectedNetworkOk returns a tuple with the RejectedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNetwork

`func (o *JunosPortUsages) SetRejectedNetwork(v string)`

SetRejectedNetwork sets RejectedNetwork field to given value.

### HasRejectedNetwork

`func (o *JunosPortUsages) HasRejectedNetwork() bool`

HasRejectedNetwork returns a boolean if a field has been set.

### SetRejectedNetworkNil

`func (o *JunosPortUsages) SetRejectedNetworkNil(b bool)`

 SetRejectedNetworkNil sets the value for RejectedNetwork to be an explicit nil

### UnsetRejectedNetwork
`func (o *JunosPortUsages) UnsetRejectedNetwork()`

UnsetRejectedNetwork ensures that no value is present for RejectedNetwork, not even an explicit nil
### GetSpeed

`func (o *JunosPortUsages) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *JunosPortUsages) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *JunosPortUsages) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *JunosPortUsages) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStormControl

`func (o *JunosPortUsages) GetStormControl() JunosStormControl`

GetStormControl returns the StormControl field if non-nil, zero value otherwise.

### GetStormControlOk

`func (o *JunosPortUsages) GetStormControlOk() (*JunosStormControl, bool)`

GetStormControlOk returns a tuple with the StormControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormControl

`func (o *JunosPortUsages) SetStormControl(v JunosStormControl)`

SetStormControl sets StormControl field to given value.

### HasStormControl

`func (o *JunosPortUsages) HasStormControl() bool`

HasStormControl returns a boolean if a field has been set.

### GetStpEdge

`func (o *JunosPortUsages) GetStpEdge() bool`

GetStpEdge returns the StpEdge field if non-nil, zero value otherwise.

### GetStpEdgeOk

`func (o *JunosPortUsages) GetStpEdgeOk() (*bool, bool)`

GetStpEdgeOk returns a tuple with the StpEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpEdge

`func (o *JunosPortUsages) SetStpEdge(v bool)`

SetStpEdge sets StpEdge field to given value.

### HasStpEdge

`func (o *JunosPortUsages) HasStpEdge() bool`

HasStpEdge returns a boolean if a field has been set.

### GetVoipNetwork

`func (o *JunosPortUsages) GetVoipNetwork() string`

GetVoipNetwork returns the VoipNetwork field if non-nil, zero value otherwise.

### GetVoipNetworkOk

`func (o *JunosPortUsages) GetVoipNetworkOk() (*string, bool)`

GetVoipNetworkOk returns a tuple with the VoipNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipNetwork

`func (o *JunosPortUsages) SetVoipNetwork(v string)`

SetVoipNetwork sets VoipNetwork field to given value.

### HasVoipNetwork

`func (o *JunosPortUsages) HasVoipNetwork() bool`

HasVoipNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

