# ApRadioBand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRrmDisable** | Pointer to **bool** |  | [optional] 
**AntGain** | Pointer to **NullableInt32** |  | [optional] 
**AntennaMode** | Pointer to [**ApRadioBandAntennaMode**](ApRadioBandAntennaMode.md) |  | [optional] [default to APRADIOBANDANTENNAMODE_DEFAULT]
**Bandwidth** | Pointer to [**Dot11Bandwidth**](Dot11Bandwidth.md) |  | [optional] [default to DOT11BANDWIDTH__20]
**Channel** | Pointer to **NullableInt32** | For Device. (primary) channel for the band, 0 means using the Site Setting | [optional] 
**Channels** | Pointer to **[]int32** | For RFTemplates. List of channels, null or empty array means auto | [optional] 
**Disabled** | Pointer to **bool** | whether to disable the radio | [optional] 
**Power** | Pointer to **NullableInt32** | TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / … | [optional] 
**PowerMax** | Pointer to **NullableInt32** | when power&#x3D;0, max tx power to use, HW-specific values will be used if not set | [optional] 
**PowerMin** | Pointer to **NullableInt32** | when power&#x3D;0, min tx power to use, HW-specific values will be used if not set | [optional] 
**Preamble** | Pointer to [**ApRadioBandPreamble**](ApRadioBandPreamble.md) |  | [optional] [default to APRADIOBANDPREAMBLE_SHORT]
**StandardPower** | Pointer to **bool** | for 6GHz Only, standard-power operation, AFC (Automatic Frequency Coordination) will be performed and we&#39;ll fallback to Low Power Indoor if AFC failed | [optional] [default to false]
**Usage** | Pointer to [**ApRadioUsage**](ApRadioUsage.md) |  | [optional] [default to APRADIOUSAGE__24]

## Methods

### NewApRadioBand

`func NewApRadioBand() *ApRadioBand`

NewApRadioBand instantiates a new ApRadioBand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRadioBandWithDefaults

`func NewApRadioBandWithDefaults() *ApRadioBand`

NewApRadioBandWithDefaults instantiates a new ApRadioBand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRrmDisable

`func (o *ApRadioBand) GetAllowRrmDisable() bool`

GetAllowRrmDisable returns the AllowRrmDisable field if non-nil, zero value otherwise.

### GetAllowRrmDisableOk

`func (o *ApRadioBand) GetAllowRrmDisableOk() (*bool, bool)`

GetAllowRrmDisableOk returns a tuple with the AllowRrmDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRrmDisable

`func (o *ApRadioBand) SetAllowRrmDisable(v bool)`

SetAllowRrmDisable sets AllowRrmDisable field to given value.

### HasAllowRrmDisable

`func (o *ApRadioBand) HasAllowRrmDisable() bool`

HasAllowRrmDisable returns a boolean if a field has been set.

### GetAntGain

`func (o *ApRadioBand) GetAntGain() int32`

GetAntGain returns the AntGain field if non-nil, zero value otherwise.

### GetAntGainOk

`func (o *ApRadioBand) GetAntGainOk() (*int32, bool)`

GetAntGainOk returns a tuple with the AntGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain

`func (o *ApRadioBand) SetAntGain(v int32)`

SetAntGain sets AntGain field to given value.

### HasAntGain

`func (o *ApRadioBand) HasAntGain() bool`

HasAntGain returns a boolean if a field has been set.

### SetAntGainNil

`func (o *ApRadioBand) SetAntGainNil(b bool)`

 SetAntGainNil sets the value for AntGain to be an explicit nil

### UnsetAntGain
`func (o *ApRadioBand) UnsetAntGain()`

UnsetAntGain ensures that no value is present for AntGain, not even an explicit nil
### GetAntennaMode

`func (o *ApRadioBand) GetAntennaMode() ApRadioBandAntennaMode`

GetAntennaMode returns the AntennaMode field if non-nil, zero value otherwise.

### GetAntennaModeOk

`func (o *ApRadioBand) GetAntennaModeOk() (*ApRadioBandAntennaMode, bool)`

GetAntennaModeOk returns a tuple with the AntennaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntennaMode

`func (o *ApRadioBand) SetAntennaMode(v ApRadioBandAntennaMode)`

SetAntennaMode sets AntennaMode field to given value.

### HasAntennaMode

`func (o *ApRadioBand) HasAntennaMode() bool`

HasAntennaMode returns a boolean if a field has been set.

### GetBandwidth

`func (o *ApRadioBand) GetBandwidth() Dot11Bandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ApRadioBand) GetBandwidthOk() (*Dot11Bandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ApRadioBand) SetBandwidth(v Dot11Bandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *ApRadioBand) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetChannel

`func (o *ApRadioBand) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApRadioBand) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApRadioBand) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApRadioBand) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *ApRadioBand) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *ApRadioBand) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetChannels

`func (o *ApRadioBand) GetChannels() []int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ApRadioBand) GetChannelsOk() (*[]int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ApRadioBand) SetChannels(v []int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ApRadioBand) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *ApRadioBand) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *ApRadioBand) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetDisabled

`func (o *ApRadioBand) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ApRadioBand) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ApRadioBand) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ApRadioBand) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetPower

`func (o *ApRadioBand) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *ApRadioBand) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *ApRadioBand) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *ApRadioBand) HasPower() bool`

HasPower returns a boolean if a field has been set.

### SetPowerNil

`func (o *ApRadioBand) SetPowerNil(b bool)`

 SetPowerNil sets the value for Power to be an explicit nil

### UnsetPower
`func (o *ApRadioBand) UnsetPower()`

UnsetPower ensures that no value is present for Power, not even an explicit nil
### GetPowerMax

`func (o *ApRadioBand) GetPowerMax() int32`

GetPowerMax returns the PowerMax field if non-nil, zero value otherwise.

### GetPowerMaxOk

`func (o *ApRadioBand) GetPowerMaxOk() (*int32, bool)`

GetPowerMaxOk returns a tuple with the PowerMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMax

`func (o *ApRadioBand) SetPowerMax(v int32)`

SetPowerMax sets PowerMax field to given value.

### HasPowerMax

`func (o *ApRadioBand) HasPowerMax() bool`

HasPowerMax returns a boolean if a field has been set.

### SetPowerMaxNil

`func (o *ApRadioBand) SetPowerMaxNil(b bool)`

 SetPowerMaxNil sets the value for PowerMax to be an explicit nil

### UnsetPowerMax
`func (o *ApRadioBand) UnsetPowerMax()`

UnsetPowerMax ensures that no value is present for PowerMax, not even an explicit nil
### GetPowerMin

`func (o *ApRadioBand) GetPowerMin() int32`

GetPowerMin returns the PowerMin field if non-nil, zero value otherwise.

### GetPowerMinOk

`func (o *ApRadioBand) GetPowerMinOk() (*int32, bool)`

GetPowerMinOk returns a tuple with the PowerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMin

`func (o *ApRadioBand) SetPowerMin(v int32)`

SetPowerMin sets PowerMin field to given value.

### HasPowerMin

`func (o *ApRadioBand) HasPowerMin() bool`

HasPowerMin returns a boolean if a field has been set.

### SetPowerMinNil

`func (o *ApRadioBand) SetPowerMinNil(b bool)`

 SetPowerMinNil sets the value for PowerMin to be an explicit nil

### UnsetPowerMin
`func (o *ApRadioBand) UnsetPowerMin()`

UnsetPowerMin ensures that no value is present for PowerMin, not even an explicit nil
### GetPreamble

`func (o *ApRadioBand) GetPreamble() ApRadioBandPreamble`

GetPreamble returns the Preamble field if non-nil, zero value otherwise.

### GetPreambleOk

`func (o *ApRadioBand) GetPreambleOk() (*ApRadioBandPreamble, bool)`

GetPreambleOk returns a tuple with the Preamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreamble

`func (o *ApRadioBand) SetPreamble(v ApRadioBandPreamble)`

SetPreamble sets Preamble field to given value.

### HasPreamble

`func (o *ApRadioBand) HasPreamble() bool`

HasPreamble returns a boolean if a field has been set.

### GetStandardPower

`func (o *ApRadioBand) GetStandardPower() bool`

GetStandardPower returns the StandardPower field if non-nil, zero value otherwise.

### GetStandardPowerOk

`func (o *ApRadioBand) GetStandardPowerOk() (*bool, bool)`

GetStandardPowerOk returns a tuple with the StandardPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPower

`func (o *ApRadioBand) SetStandardPower(v bool)`

SetStandardPower sets StandardPower field to given value.

### HasStandardPower

`func (o *ApRadioBand) HasStandardPower() bool`

HasStandardPower returns a boolean if a field has been set.

### GetUsage

`func (o *ApRadioBand) GetUsage() ApRadioUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ApRadioBand) GetUsageOk() (*ApRadioUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ApRadioBand) SetUsage(v ApRadioUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ApRadioBand) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


