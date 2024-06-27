# GatewayTemplateTunnelIpsecProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthAlgo** | Pointer to [**GatewayTemplateTunnelAuthAlgo**](GatewayTemplateTunnelAuthAlgo.md) |  | [optional] 
**DhGroup** | Pointer to [**GatewayTemplateTunnelDhGroup**](GatewayTemplateTunnelDhGroup.md) |  | [optional] [default to GATEWAYTEMPLATETUNNELDHGROUP__14]
**EncAlgo** | Pointer to [**NullableGatewayTemplateTunnelEncAlgo**](GatewayTemplateTunnelEncAlgo.md) |  | [optional] [default to GATEWAYTEMPLATETUNNELENCALGO_AES256]

## Methods

### NewGatewayTemplateTunnelIpsecProposal

`func NewGatewayTemplateTunnelIpsecProposal() *GatewayTemplateTunnelIpsecProposal`

NewGatewayTemplateTunnelIpsecProposal instantiates a new GatewayTemplateTunnelIpsecProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelIpsecProposalWithDefaults

`func NewGatewayTemplateTunnelIpsecProposalWithDefaults() *GatewayTemplateTunnelIpsecProposal`

NewGatewayTemplateTunnelIpsecProposalWithDefaults instantiates a new GatewayTemplateTunnelIpsecProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) GetAuthAlgo() GatewayTemplateTunnelAuthAlgo`

GetAuthAlgo returns the AuthAlgo field if non-nil, zero value otherwise.

### GetAuthAlgoOk

`func (o *GatewayTemplateTunnelIpsecProposal) GetAuthAlgoOk() (*GatewayTemplateTunnelAuthAlgo, bool)`

GetAuthAlgoOk returns a tuple with the AuthAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) SetAuthAlgo(v GatewayTemplateTunnelAuthAlgo)`

SetAuthAlgo sets AuthAlgo field to given value.

### HasAuthAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) HasAuthAlgo() bool`

HasAuthAlgo returns a boolean if a field has been set.

### GetDhGroup

`func (o *GatewayTemplateTunnelIpsecProposal) GetDhGroup() GatewayTemplateTunnelDhGroup`

GetDhGroup returns the DhGroup field if non-nil, zero value otherwise.

### GetDhGroupOk

`func (o *GatewayTemplateTunnelIpsecProposal) GetDhGroupOk() (*GatewayTemplateTunnelDhGroup, bool)`

GetDhGroupOk returns a tuple with the DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhGroup

`func (o *GatewayTemplateTunnelIpsecProposal) SetDhGroup(v GatewayTemplateTunnelDhGroup)`

SetDhGroup sets DhGroup field to given value.

### HasDhGroup

`func (o *GatewayTemplateTunnelIpsecProposal) HasDhGroup() bool`

HasDhGroup returns a boolean if a field has been set.

### GetEncAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) GetEncAlgo() GatewayTemplateTunnelEncAlgo`

GetEncAlgo returns the EncAlgo field if non-nil, zero value otherwise.

### GetEncAlgoOk

`func (o *GatewayTemplateTunnelIpsecProposal) GetEncAlgoOk() (*GatewayTemplateTunnelEncAlgo, bool)`

GetEncAlgoOk returns a tuple with the EncAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) SetEncAlgo(v GatewayTemplateTunnelEncAlgo)`

SetEncAlgo sets EncAlgo field to given value.

### HasEncAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) HasEncAlgo() bool`

HasEncAlgo returns a boolean if a field has been set.

### SetEncAlgoNil

`func (o *GatewayTemplateTunnelIpsecProposal) SetEncAlgoNil(b bool)`

 SetEncAlgoNil sets the value for EncAlgo to be an explicit nil

### UnsetEncAlgo
`func (o *GatewayTemplateTunnelIpsecProposal) UnsetEncAlgo()`

UnsetEncAlgo ensures that no value is present for EncAlgo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


