# GatewayTemplateTunnelConfigsProbe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **int32** | how often to trigger the probe | [optional] 
**Threshold** | Pointer to **int32** | number of consecutive misses before declaring the tunnel down | [optional] 
**Timeout** | Pointer to **int32** | time within which to complete the connectivity check | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "icmp"]

## Methods

### NewGatewayTemplateTunnelConfigsProbe

`func NewGatewayTemplateTunnelConfigsProbe() *GatewayTemplateTunnelConfigsProbe`

NewGatewayTemplateTunnelConfigsProbe instantiates a new GatewayTemplateTunnelConfigsProbe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelConfigsProbeWithDefaults

`func NewGatewayTemplateTunnelConfigsProbeWithDefaults() *GatewayTemplateTunnelConfigsProbe`

NewGatewayTemplateTunnelConfigsProbeWithDefaults instantiates a new GatewayTemplateTunnelConfigsProbe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *GatewayTemplateTunnelConfigsProbe) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GatewayTemplateTunnelConfigsProbe) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GatewayTemplateTunnelConfigsProbe) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GatewayTemplateTunnelConfigsProbe) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetThreshold

`func (o *GatewayTemplateTunnelConfigsProbe) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GatewayTemplateTunnelConfigsProbe) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GatewayTemplateTunnelConfigsProbe) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *GatewayTemplateTunnelConfigsProbe) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTimeout

`func (o *GatewayTemplateTunnelConfigsProbe) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GatewayTemplateTunnelConfigsProbe) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GatewayTemplateTunnelConfigsProbe) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GatewayTemplateTunnelConfigsProbe) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *GatewayTemplateTunnelConfigsProbe) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayTemplateTunnelConfigsProbe) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayTemplateTunnelConfigsProbe) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayTemplateTunnelConfigsProbe) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


