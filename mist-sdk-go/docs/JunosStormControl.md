# JunosStormControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoBroadcast** | Pointer to **bool** | whether to disable storm control on broadcast traffic | [optional] [default to false]
**NoMulticast** | Pointer to **bool** | whether to disable storm control on multicast traffic | [optional] [default to false]
**NoRegisteredMulticast** | Pointer to **bool** | whether to disable storm control on registered multicast traffic | [optional] [default to false]
**NoUnknownUnicast** | Pointer to **bool** | whether to disable storm control on unknown unicast traffic | [optional] [default to false]
**Percentage** | Pointer to **int32** | bandwidth-percentage, configures the storm control level as a percentage of the available bandwidth | [optional] [default to 80]

## Methods

### NewJunosStormControl

`func NewJunosStormControl() *JunosStormControl`

NewJunosStormControl instantiates a new JunosStormControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosStormControlWithDefaults

`func NewJunosStormControlWithDefaults() *JunosStormControl`

NewJunosStormControlWithDefaults instantiates a new JunosStormControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoBroadcast

`func (o *JunosStormControl) GetNoBroadcast() bool`

GetNoBroadcast returns the NoBroadcast field if non-nil, zero value otherwise.

### GetNoBroadcastOk

`func (o *JunosStormControl) GetNoBroadcastOk() (*bool, bool)`

GetNoBroadcastOk returns a tuple with the NoBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBroadcast

`func (o *JunosStormControl) SetNoBroadcast(v bool)`

SetNoBroadcast sets NoBroadcast field to given value.

### HasNoBroadcast

`func (o *JunosStormControl) HasNoBroadcast() bool`

HasNoBroadcast returns a boolean if a field has been set.

### GetNoMulticast

`func (o *JunosStormControl) GetNoMulticast() bool`

GetNoMulticast returns the NoMulticast field if non-nil, zero value otherwise.

### GetNoMulticastOk

`func (o *JunosStormControl) GetNoMulticastOk() (*bool, bool)`

GetNoMulticastOk returns a tuple with the NoMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMulticast

`func (o *JunosStormControl) SetNoMulticast(v bool)`

SetNoMulticast sets NoMulticast field to given value.

### HasNoMulticast

`func (o *JunosStormControl) HasNoMulticast() bool`

HasNoMulticast returns a boolean if a field has been set.

### GetNoRegisteredMulticast

`func (o *JunosStormControl) GetNoRegisteredMulticast() bool`

GetNoRegisteredMulticast returns the NoRegisteredMulticast field if non-nil, zero value otherwise.

### GetNoRegisteredMulticastOk

`func (o *JunosStormControl) GetNoRegisteredMulticastOk() (*bool, bool)`

GetNoRegisteredMulticastOk returns a tuple with the NoRegisteredMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRegisteredMulticast

`func (o *JunosStormControl) SetNoRegisteredMulticast(v bool)`

SetNoRegisteredMulticast sets NoRegisteredMulticast field to given value.

### HasNoRegisteredMulticast

`func (o *JunosStormControl) HasNoRegisteredMulticast() bool`

HasNoRegisteredMulticast returns a boolean if a field has been set.

### GetNoUnknownUnicast

`func (o *JunosStormControl) GetNoUnknownUnicast() bool`

GetNoUnknownUnicast returns the NoUnknownUnicast field if non-nil, zero value otherwise.

### GetNoUnknownUnicastOk

`func (o *JunosStormControl) GetNoUnknownUnicastOk() (*bool, bool)`

GetNoUnknownUnicastOk returns a tuple with the NoUnknownUnicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoUnknownUnicast

`func (o *JunosStormControl) SetNoUnknownUnicast(v bool)`

SetNoUnknownUnicast sets NoUnknownUnicast field to given value.

### HasNoUnknownUnicast

`func (o *JunosStormControl) HasNoUnknownUnicast() bool`

HasNoUnknownUnicast returns a boolean if a field has been set.

### GetPercentage

`func (o *JunosStormControl) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *JunosStormControl) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *JunosStormControl) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *JunosStormControl) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


