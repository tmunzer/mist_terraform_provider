# GatewayRoutingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Terms** | Pointer to [**[]GatewayRoutingPolicyItem**](GatewayRoutingPolicyItem.md) | zero or more criteria/filter can be specified to match the term, all criteria have to be met | [optional] 

## Methods

### NewGatewayRoutingPolicy

`func NewGatewayRoutingPolicy() *GatewayRoutingPolicy`

NewGatewayRoutingPolicy instantiates a new GatewayRoutingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRoutingPolicyWithDefaults

`func NewGatewayRoutingPolicyWithDefaults() *GatewayRoutingPolicy`

NewGatewayRoutingPolicyWithDefaults instantiates a new GatewayRoutingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerms

`func (o *GatewayRoutingPolicy) GetTerms() []GatewayRoutingPolicyItem`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *GatewayRoutingPolicy) GetTermsOk() (*[]GatewayRoutingPolicyItem, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *GatewayRoutingPolicy) SetTerms(v []GatewayRoutingPolicyItem)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *GatewayRoutingPolicy) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


