# GatewayTemplateTunnelConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoProvision** | Pointer to [**GatewayTemplateTunnelConfigsAutoProvision**](GatewayTemplateTunnelConfigsAutoProvision.md) |  | [optional] 
**IkeLifetime** | Pointer to **int32** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**IkeMode** | Pointer to **string** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] [default to "main"]
**IkeProposals** | Pointer to [**[]GatewayTemplateTunnelIkeProposal**](GatewayTemplateTunnelIkeProposal.md) | if &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**IpsecLifetime** | Pointer to **int32** | if &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**IpsecProposals** | Pointer to [**[]GatewayTemplateTunnelIpsecProposal**](GatewayTemplateTunnelIpsecProposal.md) | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**LocalId** | Pointer to **string** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;zscaler-ipsec&#x60; * &#x60;provider&#x60;&#x3D;&#x3D;&#x60;jse-ipsec&#x60; * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**Mode** | Pointer to **string** |  | [optional] [default to "active-standby"]
**Primary** | Pointer to [**GatewayTemplateTunnelPrimary**](GatewayTemplateTunnelPrimary.md) |  | [optional] 
**Probe** | Pointer to [**GatewayTemplateTunnelConfigsProbe**](GatewayTemplateTunnelConfigsProbe.md) |  | [optional] 
**Protocol** | Pointer to **string** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Psk** | Pointer to **string** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;zscaler-ipsec&#x60; * &#x60;provider&#x60;&#x3D;&#x3D;&#x60;jse-ipsec&#x60; * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**Secondary** | Pointer to [**GatewayTemplateTunnelSecondary**](GatewayTemplateTunnelSecondary.md) |  | [optional] 
**Version** | Pointer to **string** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-gre&#x60;  * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] [default to "2"]

## Methods

### NewGatewayTemplateTunnelConfigs

`func NewGatewayTemplateTunnelConfigs() *GatewayTemplateTunnelConfigs`

NewGatewayTemplateTunnelConfigs instantiates a new GatewayTemplateTunnelConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelConfigsWithDefaults

`func NewGatewayTemplateTunnelConfigsWithDefaults() *GatewayTemplateTunnelConfigs`

NewGatewayTemplateTunnelConfigsWithDefaults instantiates a new GatewayTemplateTunnelConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoProvision

`func (o *GatewayTemplateTunnelConfigs) GetAutoProvision() GatewayTemplateTunnelConfigsAutoProvision`

GetAutoProvision returns the AutoProvision field if non-nil, zero value otherwise.

### GetAutoProvisionOk

`func (o *GatewayTemplateTunnelConfigs) GetAutoProvisionOk() (*GatewayTemplateTunnelConfigsAutoProvision, bool)`

GetAutoProvisionOk returns a tuple with the AutoProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoProvision

`func (o *GatewayTemplateTunnelConfigs) SetAutoProvision(v GatewayTemplateTunnelConfigsAutoProvision)`

SetAutoProvision sets AutoProvision field to given value.

### HasAutoProvision

`func (o *GatewayTemplateTunnelConfigs) HasAutoProvision() bool`

HasAutoProvision returns a boolean if a field has been set.

### GetIkeLifetime

`func (o *GatewayTemplateTunnelConfigs) GetIkeLifetime() int32`

GetIkeLifetime returns the IkeLifetime field if non-nil, zero value otherwise.

### GetIkeLifetimeOk

`func (o *GatewayTemplateTunnelConfigs) GetIkeLifetimeOk() (*int32, bool)`

GetIkeLifetimeOk returns a tuple with the IkeLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifetime

`func (o *GatewayTemplateTunnelConfigs) SetIkeLifetime(v int32)`

SetIkeLifetime sets IkeLifetime field to given value.

### HasIkeLifetime

`func (o *GatewayTemplateTunnelConfigs) HasIkeLifetime() bool`

HasIkeLifetime returns a boolean if a field has been set.

### GetIkeMode

`func (o *GatewayTemplateTunnelConfigs) GetIkeMode() string`

GetIkeMode returns the IkeMode field if non-nil, zero value otherwise.

### GetIkeModeOk

`func (o *GatewayTemplateTunnelConfigs) GetIkeModeOk() (*string, bool)`

GetIkeModeOk returns a tuple with the IkeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeMode

`func (o *GatewayTemplateTunnelConfigs) SetIkeMode(v string)`

SetIkeMode sets IkeMode field to given value.

### HasIkeMode

`func (o *GatewayTemplateTunnelConfigs) HasIkeMode() bool`

HasIkeMode returns a boolean if a field has been set.

### GetIkeProposals

`func (o *GatewayTemplateTunnelConfigs) GetIkeProposals() []GatewayTemplateTunnelIkeProposal`

GetIkeProposals returns the IkeProposals field if non-nil, zero value otherwise.

### GetIkeProposalsOk

`func (o *GatewayTemplateTunnelConfigs) GetIkeProposalsOk() (*[]GatewayTemplateTunnelIkeProposal, bool)`

GetIkeProposalsOk returns a tuple with the IkeProposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeProposals

`func (o *GatewayTemplateTunnelConfigs) SetIkeProposals(v []GatewayTemplateTunnelIkeProposal)`

SetIkeProposals sets IkeProposals field to given value.

### HasIkeProposals

`func (o *GatewayTemplateTunnelConfigs) HasIkeProposals() bool`

HasIkeProposals returns a boolean if a field has been set.

### GetIpsecLifetime

`func (o *GatewayTemplateTunnelConfigs) GetIpsecLifetime() int32`

GetIpsecLifetime returns the IpsecLifetime field if non-nil, zero value otherwise.

### GetIpsecLifetimeOk

`func (o *GatewayTemplateTunnelConfigs) GetIpsecLifetimeOk() (*int32, bool)`

GetIpsecLifetimeOk returns a tuple with the IpsecLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecLifetime

`func (o *GatewayTemplateTunnelConfigs) SetIpsecLifetime(v int32)`

SetIpsecLifetime sets IpsecLifetime field to given value.

### HasIpsecLifetime

`func (o *GatewayTemplateTunnelConfigs) HasIpsecLifetime() bool`

HasIpsecLifetime returns a boolean if a field has been set.

### GetIpsecProposals

`func (o *GatewayTemplateTunnelConfigs) GetIpsecProposals() []GatewayTemplateTunnelIpsecProposal`

GetIpsecProposals returns the IpsecProposals field if non-nil, zero value otherwise.

### GetIpsecProposalsOk

`func (o *GatewayTemplateTunnelConfigs) GetIpsecProposalsOk() (*[]GatewayTemplateTunnelIpsecProposal, bool)`

GetIpsecProposalsOk returns a tuple with the IpsecProposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecProposals

`func (o *GatewayTemplateTunnelConfigs) SetIpsecProposals(v []GatewayTemplateTunnelIpsecProposal)`

SetIpsecProposals sets IpsecProposals field to given value.

### HasIpsecProposals

`func (o *GatewayTemplateTunnelConfigs) HasIpsecProposals() bool`

HasIpsecProposals returns a boolean if a field has been set.

### GetLocalId

`func (o *GatewayTemplateTunnelConfigs) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *GatewayTemplateTunnelConfigs) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *GatewayTemplateTunnelConfigs) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *GatewayTemplateTunnelConfigs) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetMode

`func (o *GatewayTemplateTunnelConfigs) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GatewayTemplateTunnelConfigs) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GatewayTemplateTunnelConfigs) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GatewayTemplateTunnelConfigs) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPrimary

`func (o *GatewayTemplateTunnelConfigs) GetPrimary() GatewayTemplateTunnelPrimary`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *GatewayTemplateTunnelConfigs) GetPrimaryOk() (*GatewayTemplateTunnelPrimary, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *GatewayTemplateTunnelConfigs) SetPrimary(v GatewayTemplateTunnelPrimary)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *GatewayTemplateTunnelConfigs) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetProbe

`func (o *GatewayTemplateTunnelConfigs) GetProbe() GatewayTemplateTunnelConfigsProbe`

GetProbe returns the Probe field if non-nil, zero value otherwise.

### GetProbeOk

`func (o *GatewayTemplateTunnelConfigs) GetProbeOk() (*GatewayTemplateTunnelConfigsProbe, bool)`

GetProbeOk returns a tuple with the Probe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbe

`func (o *GatewayTemplateTunnelConfigs) SetProbe(v GatewayTemplateTunnelConfigsProbe)`

SetProbe sets Probe field to given value.

### HasProbe

`func (o *GatewayTemplateTunnelConfigs) HasProbe() bool`

HasProbe returns a boolean if a field has been set.

### GetProtocol

`func (o *GatewayTemplateTunnelConfigs) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayTemplateTunnelConfigs) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayTemplateTunnelConfigs) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GatewayTemplateTunnelConfigs) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProvider

`func (o *GatewayTemplateTunnelConfigs) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GatewayTemplateTunnelConfigs) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GatewayTemplateTunnelConfigs) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GatewayTemplateTunnelConfigs) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPsk

`func (o *GatewayTemplateTunnelConfigs) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *GatewayTemplateTunnelConfigs) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *GatewayTemplateTunnelConfigs) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *GatewayTemplateTunnelConfigs) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetSecondary

`func (o *GatewayTemplateTunnelConfigs) GetSecondary() GatewayTemplateTunnelSecondary`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *GatewayTemplateTunnelConfigs) GetSecondaryOk() (*GatewayTemplateTunnelSecondary, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *GatewayTemplateTunnelConfigs) SetSecondary(v GatewayTemplateTunnelSecondary)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *GatewayTemplateTunnelConfigs) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetVersion

`func (o *GatewayTemplateTunnelConfigs) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GatewayTemplateTunnelConfigs) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GatewayTemplateTunnelConfigs) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GatewayTemplateTunnelConfigs) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


