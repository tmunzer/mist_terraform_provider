# GatewayTemplateTunnelIkeProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthAlgo** | Pointer to **string** |  | [optional] 
**DhGroup** | Pointer to **string** | * 1  * 2 (1024-bit)  * 5 * 14 (default, 2048-bit) * 15 (3072-bit)  * 16 (4096-bit) * 19 (256-bit ECP) * 20 (384-bit ECP) * 21 (521-bit ECP)  * 24 (2048-bit ECP) | [optional] [default to "14"]
**EncAlgo** | Pointer to **NullableString** |  | [optional] [default to "aes256"]

## Methods

### NewGatewayTemplateTunnelIkeProposal

`func NewGatewayTemplateTunnelIkeProposal() *GatewayTemplateTunnelIkeProposal`

NewGatewayTemplateTunnelIkeProposal instantiates a new GatewayTemplateTunnelIkeProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelIkeProposalWithDefaults

`func NewGatewayTemplateTunnelIkeProposalWithDefaults() *GatewayTemplateTunnelIkeProposal`

NewGatewayTemplateTunnelIkeProposalWithDefaults instantiates a new GatewayTemplateTunnelIkeProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthAlgo

`func (o *GatewayTemplateTunnelIkeProposal) GetAuthAlgo() string`

GetAuthAlgo returns the AuthAlgo field if non-nil, zero value otherwise.

### GetAuthAlgoOk

`func (o *GatewayTemplateTunnelIkeProposal) GetAuthAlgoOk() (*string, bool)`

GetAuthAlgoOk returns a tuple with the AuthAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthAlgo

`func (o *GatewayTemplateTunnelIkeProposal) SetAuthAlgo(v string)`

SetAuthAlgo sets AuthAlgo field to given value.

### HasAuthAlgo

`func (o *GatewayTemplateTunnelIkeProposal) HasAuthAlgo() bool`

HasAuthAlgo returns a boolean if a field has been set.

### GetDhGroup

`func (o *GatewayTemplateTunnelIkeProposal) GetDhGroup() string`

GetDhGroup returns the DhGroup field if non-nil, zero value otherwise.

### GetDhGroupOk

`func (o *GatewayTemplateTunnelIkeProposal) GetDhGroupOk() (*string, bool)`

GetDhGroupOk returns a tuple with the DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhGroup

`func (o *GatewayTemplateTunnelIkeProposal) SetDhGroup(v string)`

SetDhGroup sets DhGroup field to given value.

### HasDhGroup

`func (o *GatewayTemplateTunnelIkeProposal) HasDhGroup() bool`

HasDhGroup returns a boolean if a field has been set.

### GetEncAlgo

`func (o *GatewayTemplateTunnelIkeProposal) GetEncAlgo() string`

GetEncAlgo returns the EncAlgo field if non-nil, zero value otherwise.

### GetEncAlgoOk

`func (o *GatewayTemplateTunnelIkeProposal) GetEncAlgoOk() (*string, bool)`

GetEncAlgoOk returns a tuple with the EncAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncAlgo

`func (o *GatewayTemplateTunnelIkeProposal) SetEncAlgo(v string)`

SetEncAlgo sets EncAlgo field to given value.

### HasEncAlgo

`func (o *GatewayTemplateTunnelIkeProposal) HasEncAlgo() bool`

HasEncAlgo returns a boolean if a field has been set.

### SetEncAlgoNil

`func (o *GatewayTemplateTunnelIkeProposal) SetEncAlgoNil(b bool)`

 SetEncAlgoNil sets the value for EncAlgo to be an explicit nil

### UnsetEncAlgo
`func (o *GatewayTemplateTunnelIkeProposal) UnsetEncAlgo()`

UnsetEncAlgo ensures that no value is present for EncAlgo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


