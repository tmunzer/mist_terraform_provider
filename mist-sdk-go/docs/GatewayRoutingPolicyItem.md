# GatewayRoutingPolicyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**GatewayRoutingPolicyItemAction**](GatewayRoutingPolicyItemAction.md) |  | [optional] 
**Matching** | Pointer to [**GatewayRoutingPolicyItemMatching**](GatewayRoutingPolicyItemMatching.md) |  | [optional] 

## Methods

### NewGatewayRoutingPolicyItem

`func NewGatewayRoutingPolicyItem() *GatewayRoutingPolicyItem`

NewGatewayRoutingPolicyItem instantiates a new GatewayRoutingPolicyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRoutingPolicyItemWithDefaults

`func NewGatewayRoutingPolicyItemWithDefaults() *GatewayRoutingPolicyItem`

NewGatewayRoutingPolicyItemWithDefaults instantiates a new GatewayRoutingPolicyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GatewayRoutingPolicyItem) GetAction() GatewayRoutingPolicyItemAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GatewayRoutingPolicyItem) GetActionOk() (*GatewayRoutingPolicyItemAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GatewayRoutingPolicyItem) SetAction(v GatewayRoutingPolicyItemAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *GatewayRoutingPolicyItem) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMatching

`func (o *GatewayRoutingPolicyItem) GetMatching() GatewayRoutingPolicyItemMatching`

GetMatching returns the Matching field if non-nil, zero value otherwise.

### GetMatchingOk

`func (o *GatewayRoutingPolicyItem) GetMatchingOk() (*GatewayRoutingPolicyItemMatching, bool)`

GetMatchingOk returns a tuple with the Matching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatching

`func (o *GatewayRoutingPolicyItem) SetMatching(v GatewayRoutingPolicyItemMatching)`

SetMatching sets Matching field to given value.

### HasMatching

`func (o *GatewayRoutingPolicyItem) HasMatching() bool`

HasMatching returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


