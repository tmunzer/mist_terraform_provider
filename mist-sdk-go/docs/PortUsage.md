# PortUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllNetworks** | Pointer to **bool** | if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;trunk&#x60;, whether to trunk all network/vlans | [optional] [default to false]
**AllowDhcpd** | Pointer to **bool** | if DHCP snooping is enabled, whether DHCP server is allowed on the interfaces with. All the interfaces from port configs using this port usage are effected. Please notice that allow_dhcpd is a tri_state.  When it is not defined, it means using the system’s default setting which depends on whether the port is a access or trunk port. | [optional] 
**AllowMultipleSupplicants** | Pointer to **bool** |  | [optional] [default to false]
**BypassAuthWhenServerDown** | Pointer to **bool** | if &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60;, bypass auth for known clients if set to true when RADIUS server is down | [optional] [default to false]
**BypassAuthWhenServerDownForUnkonwnClient** | Pointer to **bool** | if &#x60;port_auth&#x60;&#x3D;&#x60;dot1x&#x60;, bypass auth for all (including unknown clients) if set to true when RADIUS server is down | [optional] [default to false]
**Description** | Pointer to **string** | description | [optional] 
**DisableAutoneg** | Pointer to **bool** | if speed and duplex are specified, whether to disable autonegotiation | [optional] [default to false]
**Disabled** | Pointer to **bool** | whether the port is disabled | [optional] [default to false]
**Duplex** | Pointer to [**PortUsageDuplex**](PortUsageDuplex.md) |  | [optional] [default to PORTUSAGEDUPLEX_AUTO]
**DynamicVlanNetworks** | Pointer to **[]string** | if dynamic vlan is used, specify the possible networks/vlans RADIUS can return | [optional] 
**EnableMacAuth** | Pointer to **bool** | if &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60;, whether to enable MAC Auth | [optional] [default to false]
**EnableQos** | Pointer to **bool** |  | [optional] [default to false]
**GuestNetwork** | Pointer to **NullableString** | if &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60;, which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed) | [optional] 
**InterSwitchLink** | Pointer to **bool** | inter_switch_link is used together with \&quot;isolation\&quot; under networks NOTE: inter_switch_link works only between Juniper device. This has to be applied to both ports connected together | [optional] [default to false]
**MacAuthOnly** | Pointer to **bool** | only effect once &#x60;enable_mac_auth&#x60;&#x3D;&#x3D;&#x60;true&#x60; | [optional] 
**MacAuthProtocol** | Pointer to [**PortUsageMacAuthProtocol**](PortUsageMacAuthProtocol.md) |  | [optional] [default to PORTUSAGEMACAUTHPROTOCOL_EAP_MD5]
**MacLimit** | Pointer to **int32** | max number of mac addresses, default is 0 for unlimited, otherwise range is 1 or higher, with upper bound constrained by platform | [optional] [default to 0]
**Mode** | Pointer to [**PortUsageMode**](PortUsageMode.md) |  | [optional] 
**Mtu** | Pointer to **int32** | media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514. | [optional] 
**Networks** | Pointer to **[]string** | if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;trunk&#x60;, the list of network/vlans | [optional] 
**PersistMac** | Pointer to **bool** | if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;access&#x60; and &#x60;port_auth&#x60;!&#x3D;&#x60;dot1x&#x60;, whether the port should retain dynamically learned MAC addresses | [optional] [default to false]
**PoeDisabled** | Pointer to **bool** | whether PoE capabilities are disabled for a port | [optional] [default to false]
**PortAuth** | Pointer to **string** | if dot1x is desired, set to dot1x | [optional] 
**PortNetwork** | Pointer to **string** | native network/vlan for untagged traffic | [optional] 
**ReauthInterval** | Pointer to **int32** | if &#x60;port_auth&#x60;&#x3D;&#x60;dot1x&#x60;, reauthentication interval range | [optional] [default to 3600]
**RejectedNetwork** | Pointer to **NullableString** | if &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60;, when radius server reject / fails | [optional] 
**Speed** | Pointer to **string** | speed, default is auto to automatically negotiate speed | [optional] 
**StormControl** | Pointer to [**PortUsageStormControl**](PortUsageStormControl.md) |  | [optional] 
**StpEdge** | Pointer to **bool** | when enabled, the port is not expected to receive BPDU frames | [optional] [default to false]
**VoipNetwork** | Pointer to **string** | network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth | [optional] 

## Methods

### NewPortUsage

`func NewPortUsage() *PortUsage`

NewPortUsage instantiates a new PortUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortUsageWithDefaults

`func NewPortUsageWithDefaults() *PortUsage`

NewPortUsageWithDefaults instantiates a new PortUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllNetworks

`func (o *PortUsage) GetAllNetworks() bool`

GetAllNetworks returns the AllNetworks field if non-nil, zero value otherwise.

### GetAllNetworksOk

`func (o *PortUsage) GetAllNetworksOk() (*bool, bool)`

GetAllNetworksOk returns a tuple with the AllNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllNetworks

`func (o *PortUsage) SetAllNetworks(v bool)`

SetAllNetworks sets AllNetworks field to given value.

### HasAllNetworks

`func (o *PortUsage) HasAllNetworks() bool`

HasAllNetworks returns a boolean if a field has been set.

### GetAllowDhcpd

`func (o *PortUsage) GetAllowDhcpd() bool`

GetAllowDhcpd returns the AllowDhcpd field if non-nil, zero value otherwise.

### GetAllowDhcpdOk

`func (o *PortUsage) GetAllowDhcpdOk() (*bool, bool)`

GetAllowDhcpdOk returns a tuple with the AllowDhcpd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDhcpd

`func (o *PortUsage) SetAllowDhcpd(v bool)`

SetAllowDhcpd sets AllowDhcpd field to given value.

### HasAllowDhcpd

`func (o *PortUsage) HasAllowDhcpd() bool`

HasAllowDhcpd returns a boolean if a field has been set.

### GetAllowMultipleSupplicants

`func (o *PortUsage) GetAllowMultipleSupplicants() bool`

GetAllowMultipleSupplicants returns the AllowMultipleSupplicants field if non-nil, zero value otherwise.

### GetAllowMultipleSupplicantsOk

`func (o *PortUsage) GetAllowMultipleSupplicantsOk() (*bool, bool)`

GetAllowMultipleSupplicantsOk returns a tuple with the AllowMultipleSupplicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSupplicants

`func (o *PortUsage) SetAllowMultipleSupplicants(v bool)`

SetAllowMultipleSupplicants sets AllowMultipleSupplicants field to given value.

### HasAllowMultipleSupplicants

`func (o *PortUsage) HasAllowMultipleSupplicants() bool`

HasAllowMultipleSupplicants returns a boolean if a field has been set.

### GetBypassAuthWhenServerDown

`func (o *PortUsage) GetBypassAuthWhenServerDown() bool`

GetBypassAuthWhenServerDown returns the BypassAuthWhenServerDown field if non-nil, zero value otherwise.

### GetBypassAuthWhenServerDownOk

`func (o *PortUsage) GetBypassAuthWhenServerDownOk() (*bool, bool)`

GetBypassAuthWhenServerDownOk returns a tuple with the BypassAuthWhenServerDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAuthWhenServerDown

`func (o *PortUsage) SetBypassAuthWhenServerDown(v bool)`

SetBypassAuthWhenServerDown sets BypassAuthWhenServerDown field to given value.

### HasBypassAuthWhenServerDown

`func (o *PortUsage) HasBypassAuthWhenServerDown() bool`

HasBypassAuthWhenServerDown returns a boolean if a field has been set.

### GetBypassAuthWhenServerDownForUnkonwnClient

`func (o *PortUsage) GetBypassAuthWhenServerDownForUnkonwnClient() bool`

GetBypassAuthWhenServerDownForUnkonwnClient returns the BypassAuthWhenServerDownForUnkonwnClient field if non-nil, zero value otherwise.

### GetBypassAuthWhenServerDownForUnkonwnClientOk

`func (o *PortUsage) GetBypassAuthWhenServerDownForUnkonwnClientOk() (*bool, bool)`

GetBypassAuthWhenServerDownForUnkonwnClientOk returns a tuple with the BypassAuthWhenServerDownForUnkonwnClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAuthWhenServerDownForUnkonwnClient

`func (o *PortUsage) SetBypassAuthWhenServerDownForUnkonwnClient(v bool)`

SetBypassAuthWhenServerDownForUnkonwnClient sets BypassAuthWhenServerDownForUnkonwnClient field to given value.

### HasBypassAuthWhenServerDownForUnkonwnClient

`func (o *PortUsage) HasBypassAuthWhenServerDownForUnkonwnClient() bool`

HasBypassAuthWhenServerDownForUnkonwnClient returns a boolean if a field has been set.

### GetDescription

`func (o *PortUsage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PortUsage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PortUsage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PortUsage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableAutoneg

`func (o *PortUsage) GetDisableAutoneg() bool`

GetDisableAutoneg returns the DisableAutoneg field if non-nil, zero value otherwise.

### GetDisableAutonegOk

`func (o *PortUsage) GetDisableAutonegOk() (*bool, bool)`

GetDisableAutonegOk returns a tuple with the DisableAutoneg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoneg

`func (o *PortUsage) SetDisableAutoneg(v bool)`

SetDisableAutoneg sets DisableAutoneg field to given value.

### HasDisableAutoneg

`func (o *PortUsage) HasDisableAutoneg() bool`

HasDisableAutoneg returns a boolean if a field has been set.

### GetDisabled

`func (o *PortUsage) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PortUsage) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PortUsage) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PortUsage) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDuplex

`func (o *PortUsage) GetDuplex() PortUsageDuplex`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *PortUsage) GetDuplexOk() (*PortUsageDuplex, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *PortUsage) SetDuplex(v PortUsageDuplex)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *PortUsage) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetDynamicVlanNetworks

`func (o *PortUsage) GetDynamicVlanNetworks() []string`

GetDynamicVlanNetworks returns the DynamicVlanNetworks field if non-nil, zero value otherwise.

### GetDynamicVlanNetworksOk

`func (o *PortUsage) GetDynamicVlanNetworksOk() (*[]string, bool)`

GetDynamicVlanNetworksOk returns a tuple with the DynamicVlanNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVlanNetworks

`func (o *PortUsage) SetDynamicVlanNetworks(v []string)`

SetDynamicVlanNetworks sets DynamicVlanNetworks field to given value.

### HasDynamicVlanNetworks

`func (o *PortUsage) HasDynamicVlanNetworks() bool`

HasDynamicVlanNetworks returns a boolean if a field has been set.

### GetEnableMacAuth

`func (o *PortUsage) GetEnableMacAuth() bool`

GetEnableMacAuth returns the EnableMacAuth field if non-nil, zero value otherwise.

### GetEnableMacAuthOk

`func (o *PortUsage) GetEnableMacAuthOk() (*bool, bool)`

GetEnableMacAuthOk returns a tuple with the EnableMacAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMacAuth

`func (o *PortUsage) SetEnableMacAuth(v bool)`

SetEnableMacAuth sets EnableMacAuth field to given value.

### HasEnableMacAuth

`func (o *PortUsage) HasEnableMacAuth() bool`

HasEnableMacAuth returns a boolean if a field has been set.

### GetEnableQos

`func (o *PortUsage) GetEnableQos() bool`

GetEnableQos returns the EnableQos field if non-nil, zero value otherwise.

### GetEnableQosOk

`func (o *PortUsage) GetEnableQosOk() (*bool, bool)`

GetEnableQosOk returns a tuple with the EnableQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQos

`func (o *PortUsage) SetEnableQos(v bool)`

SetEnableQos sets EnableQos field to given value.

### HasEnableQos

`func (o *PortUsage) HasEnableQos() bool`

HasEnableQos returns a boolean if a field has been set.

### GetGuestNetwork

`func (o *PortUsage) GetGuestNetwork() string`

GetGuestNetwork returns the GuestNetwork field if non-nil, zero value otherwise.

### GetGuestNetworkOk

`func (o *PortUsage) GetGuestNetworkOk() (*string, bool)`

GetGuestNetworkOk returns a tuple with the GuestNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNetwork

`func (o *PortUsage) SetGuestNetwork(v string)`

SetGuestNetwork sets GuestNetwork field to given value.

### HasGuestNetwork

`func (o *PortUsage) HasGuestNetwork() bool`

HasGuestNetwork returns a boolean if a field has been set.

### SetGuestNetworkNil

`func (o *PortUsage) SetGuestNetworkNil(b bool)`

 SetGuestNetworkNil sets the value for GuestNetwork to be an explicit nil

### UnsetGuestNetwork
`func (o *PortUsage) UnsetGuestNetwork()`

UnsetGuestNetwork ensures that no value is present for GuestNetwork, not even an explicit nil
### GetInterSwitchLink

`func (o *PortUsage) GetInterSwitchLink() bool`

GetInterSwitchLink returns the InterSwitchLink field if non-nil, zero value otherwise.

### GetInterSwitchLinkOk

`func (o *PortUsage) GetInterSwitchLinkOk() (*bool, bool)`

GetInterSwitchLinkOk returns a tuple with the InterSwitchLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterSwitchLink

`func (o *PortUsage) SetInterSwitchLink(v bool)`

SetInterSwitchLink sets InterSwitchLink field to given value.

### HasInterSwitchLink

`func (o *PortUsage) HasInterSwitchLink() bool`

HasInterSwitchLink returns a boolean if a field has been set.

### GetMacAuthOnly

`func (o *PortUsage) GetMacAuthOnly() bool`

GetMacAuthOnly returns the MacAuthOnly field if non-nil, zero value otherwise.

### GetMacAuthOnlyOk

`func (o *PortUsage) GetMacAuthOnlyOk() (*bool, bool)`

GetMacAuthOnlyOk returns a tuple with the MacAuthOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAuthOnly

`func (o *PortUsage) SetMacAuthOnly(v bool)`

SetMacAuthOnly sets MacAuthOnly field to given value.

### HasMacAuthOnly

`func (o *PortUsage) HasMacAuthOnly() bool`

HasMacAuthOnly returns a boolean if a field has been set.

### GetMacAuthProtocol

`func (o *PortUsage) GetMacAuthProtocol() PortUsageMacAuthProtocol`

GetMacAuthProtocol returns the MacAuthProtocol field if non-nil, zero value otherwise.

### GetMacAuthProtocolOk

`func (o *PortUsage) GetMacAuthProtocolOk() (*PortUsageMacAuthProtocol, bool)`

GetMacAuthProtocolOk returns a tuple with the MacAuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAuthProtocol

`func (o *PortUsage) SetMacAuthProtocol(v PortUsageMacAuthProtocol)`

SetMacAuthProtocol sets MacAuthProtocol field to given value.

### HasMacAuthProtocol

`func (o *PortUsage) HasMacAuthProtocol() bool`

HasMacAuthProtocol returns a boolean if a field has been set.

### GetMacLimit

`func (o *PortUsage) GetMacLimit() int32`

GetMacLimit returns the MacLimit field if non-nil, zero value otherwise.

### GetMacLimitOk

`func (o *PortUsage) GetMacLimitOk() (*int32, bool)`

GetMacLimitOk returns a tuple with the MacLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLimit

`func (o *PortUsage) SetMacLimit(v int32)`

SetMacLimit sets MacLimit field to given value.

### HasMacLimit

`func (o *PortUsage) HasMacLimit() bool`

HasMacLimit returns a boolean if a field has been set.

### GetMode

`func (o *PortUsage) GetMode() PortUsageMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PortUsage) GetModeOk() (*PortUsageMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PortUsage) SetMode(v PortUsageMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PortUsage) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMtu

`func (o *PortUsage) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PortUsage) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PortUsage) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *PortUsage) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNetworks

`func (o *PortUsage) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *PortUsage) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *PortUsage) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *PortUsage) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetPersistMac

`func (o *PortUsage) GetPersistMac() bool`

GetPersistMac returns the PersistMac field if non-nil, zero value otherwise.

### GetPersistMacOk

`func (o *PortUsage) GetPersistMacOk() (*bool, bool)`

GetPersistMacOk returns a tuple with the PersistMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistMac

`func (o *PortUsage) SetPersistMac(v bool)`

SetPersistMac sets PersistMac field to given value.

### HasPersistMac

`func (o *PortUsage) HasPersistMac() bool`

HasPersistMac returns a boolean if a field has been set.

### GetPoeDisabled

`func (o *PortUsage) GetPoeDisabled() bool`

GetPoeDisabled returns the PoeDisabled field if non-nil, zero value otherwise.

### GetPoeDisabledOk

`func (o *PortUsage) GetPoeDisabledOk() (*bool, bool)`

GetPoeDisabledOk returns a tuple with the PoeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisabled

`func (o *PortUsage) SetPoeDisabled(v bool)`

SetPoeDisabled sets PoeDisabled field to given value.

### HasPoeDisabled

`func (o *PortUsage) HasPoeDisabled() bool`

HasPoeDisabled returns a boolean if a field has been set.

### GetPortAuth

`func (o *PortUsage) GetPortAuth() string`

GetPortAuth returns the PortAuth field if non-nil, zero value otherwise.

### GetPortAuthOk

`func (o *PortUsage) GetPortAuthOk() (*string, bool)`

GetPortAuthOk returns a tuple with the PortAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAuth

`func (o *PortUsage) SetPortAuth(v string)`

SetPortAuth sets PortAuth field to given value.

### HasPortAuth

`func (o *PortUsage) HasPortAuth() bool`

HasPortAuth returns a boolean if a field has been set.

### GetPortNetwork

`func (o *PortUsage) GetPortNetwork() string`

GetPortNetwork returns the PortNetwork field if non-nil, zero value otherwise.

### GetPortNetworkOk

`func (o *PortUsage) GetPortNetworkOk() (*string, bool)`

GetPortNetworkOk returns a tuple with the PortNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNetwork

`func (o *PortUsage) SetPortNetwork(v string)`

SetPortNetwork sets PortNetwork field to given value.

### HasPortNetwork

`func (o *PortUsage) HasPortNetwork() bool`

HasPortNetwork returns a boolean if a field has been set.

### GetReauthInterval

`func (o *PortUsage) GetReauthInterval() int32`

GetReauthInterval returns the ReauthInterval field if non-nil, zero value otherwise.

### GetReauthIntervalOk

`func (o *PortUsage) GetReauthIntervalOk() (*int32, bool)`

GetReauthIntervalOk returns a tuple with the ReauthInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthInterval

`func (o *PortUsage) SetReauthInterval(v int32)`

SetReauthInterval sets ReauthInterval field to given value.

### HasReauthInterval

`func (o *PortUsage) HasReauthInterval() bool`

HasReauthInterval returns a boolean if a field has been set.

### GetRejectedNetwork

`func (o *PortUsage) GetRejectedNetwork() string`

GetRejectedNetwork returns the RejectedNetwork field if non-nil, zero value otherwise.

### GetRejectedNetworkOk

`func (o *PortUsage) GetRejectedNetworkOk() (*string, bool)`

GetRejectedNetworkOk returns a tuple with the RejectedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNetwork

`func (o *PortUsage) SetRejectedNetwork(v string)`

SetRejectedNetwork sets RejectedNetwork field to given value.

### HasRejectedNetwork

`func (o *PortUsage) HasRejectedNetwork() bool`

HasRejectedNetwork returns a boolean if a field has been set.

### SetRejectedNetworkNil

`func (o *PortUsage) SetRejectedNetworkNil(b bool)`

 SetRejectedNetworkNil sets the value for RejectedNetwork to be an explicit nil

### UnsetRejectedNetwork
`func (o *PortUsage) UnsetRejectedNetwork()`

UnsetRejectedNetwork ensures that no value is present for RejectedNetwork, not even an explicit nil
### GetSpeed

`func (o *PortUsage) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *PortUsage) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *PortUsage) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *PortUsage) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStormControl

`func (o *PortUsage) GetStormControl() PortUsageStormControl`

GetStormControl returns the StormControl field if non-nil, zero value otherwise.

### GetStormControlOk

`func (o *PortUsage) GetStormControlOk() (*PortUsageStormControl, bool)`

GetStormControlOk returns a tuple with the StormControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormControl

`func (o *PortUsage) SetStormControl(v PortUsageStormControl)`

SetStormControl sets StormControl field to given value.

### HasStormControl

`func (o *PortUsage) HasStormControl() bool`

HasStormControl returns a boolean if a field has been set.

### GetStpEdge

`func (o *PortUsage) GetStpEdge() bool`

GetStpEdge returns the StpEdge field if non-nil, zero value otherwise.

### GetStpEdgeOk

`func (o *PortUsage) GetStpEdgeOk() (*bool, bool)`

GetStpEdgeOk returns a tuple with the StpEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpEdge

`func (o *PortUsage) SetStpEdge(v bool)`

SetStpEdge sets StpEdge field to given value.

### HasStpEdge

`func (o *PortUsage) HasStpEdge() bool`

HasStpEdge returns a boolean if a field has been set.

### GetVoipNetwork

`func (o *PortUsage) GetVoipNetwork() string`

GetVoipNetwork returns the VoipNetwork field if non-nil, zero value otherwise.

### GetVoipNetworkOk

`func (o *PortUsage) GetVoipNetworkOk() (*string, bool)`

GetVoipNetworkOk returns a tuple with the VoipNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipNetwork

`func (o *PortUsage) SetVoipNetwork(v string)`

SetVoipNetwork sets VoipNetwork field to given value.

### HasVoipNetwork

`func (o *PortUsage) HasVoipNetwork() bool`

HasVoipNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


