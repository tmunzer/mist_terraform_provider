# JunosEvpnOptionsUnderlay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsBase** | Pointer to **int32** |  | [optional] 
**RoutedIdPrefix** | Pointer to **string** |  | [optional] 
**Subnet** | Pointer to **string** | underlay subnet, by default, &#x60;10.255.240.0/20&#x60;, or &#x60;fd31:5700::/64&#x60; for ipv6 | [optional] 
**UseIpv6** | Pointer to **bool** | if v6 is desired for underlay | [optional] [default to false]

## Methods

### NewJunosEvpnOptionsUnderlay

`func NewJunosEvpnOptionsUnderlay() *JunosEvpnOptionsUnderlay`

NewJunosEvpnOptionsUnderlay instantiates a new JunosEvpnOptionsUnderlay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosEvpnOptionsUnderlayWithDefaults

`func NewJunosEvpnOptionsUnderlayWithDefaults() *JunosEvpnOptionsUnderlay`

NewJunosEvpnOptionsUnderlayWithDefaults instantiates a new JunosEvpnOptionsUnderlay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsBase

`func (o *JunosEvpnOptionsUnderlay) GetAsBase() int32`

GetAsBase returns the AsBase field if non-nil, zero value otherwise.

### GetAsBaseOk

`func (o *JunosEvpnOptionsUnderlay) GetAsBaseOk() (*int32, bool)`

GetAsBaseOk returns a tuple with the AsBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsBase

`func (o *JunosEvpnOptionsUnderlay) SetAsBase(v int32)`

SetAsBase sets AsBase field to given value.

### HasAsBase

`func (o *JunosEvpnOptionsUnderlay) HasAsBase() bool`

HasAsBase returns a boolean if a field has been set.

### GetRoutedIdPrefix

`func (o *JunosEvpnOptionsUnderlay) GetRoutedIdPrefix() string`

GetRoutedIdPrefix returns the RoutedIdPrefix field if non-nil, zero value otherwise.

### GetRoutedIdPrefixOk

`func (o *JunosEvpnOptionsUnderlay) GetRoutedIdPrefixOk() (*string, bool)`

GetRoutedIdPrefixOk returns a tuple with the RoutedIdPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutedIdPrefix

`func (o *JunosEvpnOptionsUnderlay) SetRoutedIdPrefix(v string)`

SetRoutedIdPrefix sets RoutedIdPrefix field to given value.

### HasRoutedIdPrefix

`func (o *JunosEvpnOptionsUnderlay) HasRoutedIdPrefix() bool`

HasRoutedIdPrefix returns a boolean if a field has been set.

### GetSubnet

`func (o *JunosEvpnOptionsUnderlay) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *JunosEvpnOptionsUnderlay) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *JunosEvpnOptionsUnderlay) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *JunosEvpnOptionsUnderlay) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetUseIpv6

`func (o *JunosEvpnOptionsUnderlay) GetUseIpv6() bool`

GetUseIpv6 returns the UseIpv6 field if non-nil, zero value otherwise.

### GetUseIpv6Ok

`func (o *JunosEvpnOptionsUnderlay) GetUseIpv6Ok() (*bool, bool)`

GetUseIpv6Ok returns a tuple with the UseIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpv6

`func (o *JunosEvpnOptionsUnderlay) SetUseIpv6(v bool)`

SetUseIpv6 sets UseIpv6 field to given value.

### HasUseIpv6

`func (o *JunosEvpnOptionsUnderlay) HasUseIpv6() bool`

HasUseIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

