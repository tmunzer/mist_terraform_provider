# GatewayRoutingPolicyItemMatchingRouteExists

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Route** | Pointer to **string** |  | [optional] 
**VrfName** | Pointer to **string** | name of the vrf instance it can also be the name of the VPN or wan if they | [optional] [default to "default"]

## Methods

### NewGatewayRoutingPolicyItemMatchingRouteExists

`func NewGatewayRoutingPolicyItemMatchingRouteExists() *GatewayRoutingPolicyItemMatchingRouteExists`

NewGatewayRoutingPolicyItemMatchingRouteExists instantiates a new GatewayRoutingPolicyItemMatchingRouteExists object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRoutingPolicyItemMatchingRouteExistsWithDefaults

`func NewGatewayRoutingPolicyItemMatchingRouteExistsWithDefaults() *GatewayRoutingPolicyItemMatchingRouteExists`

NewGatewayRoutingPolicyItemMatchingRouteExistsWithDefaults instantiates a new GatewayRoutingPolicyItemMatchingRouteExists object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoute

`func (o *GatewayRoutingPolicyItemMatchingRouteExists) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *GatewayRoutingPolicyItemMatchingRouteExists) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *GatewayRoutingPolicyItemMatchingRouteExists) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *GatewayRoutingPolicyItemMatchingRouteExists) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetVrfName

`func (o *GatewayRoutingPolicyItemMatchingRouteExists) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *GatewayRoutingPolicyItemMatchingRouteExists) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *GatewayRoutingPolicyItemMatchingRouteExists) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *GatewayRoutingPolicyItemMatchingRouteExists) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


