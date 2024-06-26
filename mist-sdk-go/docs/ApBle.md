# ApBle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeaconEnabled** | Pointer to **bool** | whether Mist beacons is enabled | [optional] [default to false]
**BeaconRate** | Pointer to **int32** | required if &#x60;beacon_rate_mode&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, 1-10, in number-beacons-per-second | [optional] 
**BeaconRateMode** | Pointer to [**ApBleBeaconRateMode**](ApBleBeaconRateMode.md) |  | [optional] [default to APBLEBEACONRATEMODE_DEFAULT]
**BeamDisabled** | Pointer to **[]int32** | list of AP BLE location beam numbers (1-8) which should be disabled at the AP and not transmit location information (where beam 1 is oriented at the top the AP, growing counter-clock-wise, with 9 being the omni BLE beam) | [optional] 
**CustomBlePacketEnabled** | Pointer to **bool** | can be enabled if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;, whether to send custom packet | [optional] [default to false]
**CustomBlePacketFrame** | Pointer to **string** | The custom frame to be sent out in this beacon. The frame must be a hexstring | [optional] 
**CustomBlePacketFreqMsec** | Pointer to **int32** | Frequency (msec) of data emitted by custom ble beacon | [optional] 
**EddystoneUidAdvPower** | Pointer to **int32** | advertised TX Power, -100 to 20 (dBm), omit this attribute to use default | [optional] 
**EddystoneUidBeams** | Pointer to **string** |  | [optional] 
**EddystoneUidEnabled** | Pointer to **bool** | only if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;false&#x60;, Whether Eddystone-UID beacon is enabled | [optional] [default to false]
**EddystoneUidFreqMsec** | Pointer to **int32** | Frequency (msec) of data emmit by Eddystone-UID beacon | [optional] 
**EddystoneUidInstance** | Pointer to **string** | Eddystone-UID instance for the device | [optional] 
**EddystoneUidNamespace** | Pointer to **string** | Eddystone-UID namespace | [optional] 
**EddystoneUrlAdvPower** | Pointer to **int32** | advertised TX Power, -100 to 20 (dBm), omit this attribute to use default | [optional] 
**EddystoneUrlBeams** | Pointer to **string** |  | [optional] 
**EddystoneUrlEnabled** | Pointer to **bool** | only if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;false&#x60;, Whether Eddystone-URL beacon is enabled | [optional] 
**EddystoneUrlFreqMsec** | Pointer to **int32** | Frequency (msec) of data emit by Eddystone-UID beacon | [optional] 
**EddystoneUrlUrl** | Pointer to **string** | URL pointed by Eddystone-URL beacon | [optional] 
**IbeaconAdvPower** | Pointer to **int32** | advertised TX Power, -100 to 20 (dBm), omit this attribute to use default | [optional] 
**IbeaconBeams** | Pointer to **string** |  | [optional] 
**IbeaconEnabled** | Pointer to **bool** | can be enabled if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;, whether to send iBeacon | [optional] [default to false]
**IbeaconFreqMsec** | Pointer to **int32** | Frequency (msec) of data emmit for iBeacon | [optional] 
**IbeaconMajor** | Pointer to **int32** | Major number for iBeacon | [optional] 
**IbeaconMinor** | Pointer to **int32** | Minor number for iBeacon | [optional] 
**IbeaconUuid** | Pointer to **string** | optional, if not specified, the same UUID as the beacon will be used | [optional] 
**Power** | Pointer to **int32** | required if &#x60;power_mode&#x60;&#x3D;&#x3D;&#x60;custom&#x60; | [optional] [default to 9]
**PowerMode** | Pointer to [**ApBlePowerMode**](ApBlePowerMode.md) |  | [optional] [default to APBLEPOWERMODE_DEFAULT]

## Methods

### NewApBle

`func NewApBle() *ApBle`

NewApBle instantiates a new ApBle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApBleWithDefaults

`func NewApBleWithDefaults() *ApBle`

NewApBleWithDefaults instantiates a new ApBle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeaconEnabled

`func (o *ApBle) GetBeaconEnabled() bool`

GetBeaconEnabled returns the BeaconEnabled field if non-nil, zero value otherwise.

### GetBeaconEnabledOk

`func (o *ApBle) GetBeaconEnabledOk() (*bool, bool)`

GetBeaconEnabledOk returns a tuple with the BeaconEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconEnabled

`func (o *ApBle) SetBeaconEnabled(v bool)`

SetBeaconEnabled sets BeaconEnabled field to given value.

### HasBeaconEnabled

`func (o *ApBle) HasBeaconEnabled() bool`

HasBeaconEnabled returns a boolean if a field has been set.

### GetBeaconRate

`func (o *ApBle) GetBeaconRate() int32`

GetBeaconRate returns the BeaconRate field if non-nil, zero value otherwise.

### GetBeaconRateOk

`func (o *ApBle) GetBeaconRateOk() (*int32, bool)`

GetBeaconRateOk returns a tuple with the BeaconRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconRate

`func (o *ApBle) SetBeaconRate(v int32)`

SetBeaconRate sets BeaconRate field to given value.

### HasBeaconRate

`func (o *ApBle) HasBeaconRate() bool`

HasBeaconRate returns a boolean if a field has been set.

### GetBeaconRateMode

`func (o *ApBle) GetBeaconRateMode() ApBleBeaconRateMode`

GetBeaconRateMode returns the BeaconRateMode field if non-nil, zero value otherwise.

### GetBeaconRateModeOk

`func (o *ApBle) GetBeaconRateModeOk() (*ApBleBeaconRateMode, bool)`

GetBeaconRateModeOk returns a tuple with the BeaconRateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconRateMode

`func (o *ApBle) SetBeaconRateMode(v ApBleBeaconRateMode)`

SetBeaconRateMode sets BeaconRateMode field to given value.

### HasBeaconRateMode

`func (o *ApBle) HasBeaconRateMode() bool`

HasBeaconRateMode returns a boolean if a field has been set.

### GetBeamDisabled

`func (o *ApBle) GetBeamDisabled() []int32`

GetBeamDisabled returns the BeamDisabled field if non-nil, zero value otherwise.

### GetBeamDisabledOk

`func (o *ApBle) GetBeamDisabledOk() (*[]int32, bool)`

GetBeamDisabledOk returns a tuple with the BeamDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamDisabled

`func (o *ApBle) SetBeamDisabled(v []int32)`

SetBeamDisabled sets BeamDisabled field to given value.

### HasBeamDisabled

`func (o *ApBle) HasBeamDisabled() bool`

HasBeamDisabled returns a boolean if a field has been set.

### GetCustomBlePacketEnabled

`func (o *ApBle) GetCustomBlePacketEnabled() bool`

GetCustomBlePacketEnabled returns the CustomBlePacketEnabled field if non-nil, zero value otherwise.

### GetCustomBlePacketEnabledOk

`func (o *ApBle) GetCustomBlePacketEnabledOk() (*bool, bool)`

GetCustomBlePacketEnabledOk returns a tuple with the CustomBlePacketEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBlePacketEnabled

`func (o *ApBle) SetCustomBlePacketEnabled(v bool)`

SetCustomBlePacketEnabled sets CustomBlePacketEnabled field to given value.

### HasCustomBlePacketEnabled

`func (o *ApBle) HasCustomBlePacketEnabled() bool`

HasCustomBlePacketEnabled returns a boolean if a field has been set.

### GetCustomBlePacketFrame

`func (o *ApBle) GetCustomBlePacketFrame() string`

GetCustomBlePacketFrame returns the CustomBlePacketFrame field if non-nil, zero value otherwise.

### GetCustomBlePacketFrameOk

`func (o *ApBle) GetCustomBlePacketFrameOk() (*string, bool)`

GetCustomBlePacketFrameOk returns a tuple with the CustomBlePacketFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBlePacketFrame

`func (o *ApBle) SetCustomBlePacketFrame(v string)`

SetCustomBlePacketFrame sets CustomBlePacketFrame field to given value.

### HasCustomBlePacketFrame

`func (o *ApBle) HasCustomBlePacketFrame() bool`

HasCustomBlePacketFrame returns a boolean if a field has been set.

### GetCustomBlePacketFreqMsec

`func (o *ApBle) GetCustomBlePacketFreqMsec() int32`

GetCustomBlePacketFreqMsec returns the CustomBlePacketFreqMsec field if non-nil, zero value otherwise.

### GetCustomBlePacketFreqMsecOk

`func (o *ApBle) GetCustomBlePacketFreqMsecOk() (*int32, bool)`

GetCustomBlePacketFreqMsecOk returns a tuple with the CustomBlePacketFreqMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBlePacketFreqMsec

`func (o *ApBle) SetCustomBlePacketFreqMsec(v int32)`

SetCustomBlePacketFreqMsec sets CustomBlePacketFreqMsec field to given value.

### HasCustomBlePacketFreqMsec

`func (o *ApBle) HasCustomBlePacketFreqMsec() bool`

HasCustomBlePacketFreqMsec returns a boolean if a field has been set.

### GetEddystoneUidAdvPower

`func (o *ApBle) GetEddystoneUidAdvPower() int32`

GetEddystoneUidAdvPower returns the EddystoneUidAdvPower field if non-nil, zero value otherwise.

### GetEddystoneUidAdvPowerOk

`func (o *ApBle) GetEddystoneUidAdvPowerOk() (*int32, bool)`

GetEddystoneUidAdvPowerOk returns a tuple with the EddystoneUidAdvPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidAdvPower

`func (o *ApBle) SetEddystoneUidAdvPower(v int32)`

SetEddystoneUidAdvPower sets EddystoneUidAdvPower field to given value.

### HasEddystoneUidAdvPower

`func (o *ApBle) HasEddystoneUidAdvPower() bool`

HasEddystoneUidAdvPower returns a boolean if a field has been set.

### GetEddystoneUidBeams

`func (o *ApBle) GetEddystoneUidBeams() string`

GetEddystoneUidBeams returns the EddystoneUidBeams field if non-nil, zero value otherwise.

### GetEddystoneUidBeamsOk

`func (o *ApBle) GetEddystoneUidBeamsOk() (*string, bool)`

GetEddystoneUidBeamsOk returns a tuple with the EddystoneUidBeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidBeams

`func (o *ApBle) SetEddystoneUidBeams(v string)`

SetEddystoneUidBeams sets EddystoneUidBeams field to given value.

### HasEddystoneUidBeams

`func (o *ApBle) HasEddystoneUidBeams() bool`

HasEddystoneUidBeams returns a boolean if a field has been set.

### GetEddystoneUidEnabled

`func (o *ApBle) GetEddystoneUidEnabled() bool`

GetEddystoneUidEnabled returns the EddystoneUidEnabled field if non-nil, zero value otherwise.

### GetEddystoneUidEnabledOk

`func (o *ApBle) GetEddystoneUidEnabledOk() (*bool, bool)`

GetEddystoneUidEnabledOk returns a tuple with the EddystoneUidEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidEnabled

`func (o *ApBle) SetEddystoneUidEnabled(v bool)`

SetEddystoneUidEnabled sets EddystoneUidEnabled field to given value.

### HasEddystoneUidEnabled

`func (o *ApBle) HasEddystoneUidEnabled() bool`

HasEddystoneUidEnabled returns a boolean if a field has been set.

### GetEddystoneUidFreqMsec

`func (o *ApBle) GetEddystoneUidFreqMsec() int32`

GetEddystoneUidFreqMsec returns the EddystoneUidFreqMsec field if non-nil, zero value otherwise.

### GetEddystoneUidFreqMsecOk

`func (o *ApBle) GetEddystoneUidFreqMsecOk() (*int32, bool)`

GetEddystoneUidFreqMsecOk returns a tuple with the EddystoneUidFreqMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidFreqMsec

`func (o *ApBle) SetEddystoneUidFreqMsec(v int32)`

SetEddystoneUidFreqMsec sets EddystoneUidFreqMsec field to given value.

### HasEddystoneUidFreqMsec

`func (o *ApBle) HasEddystoneUidFreqMsec() bool`

HasEddystoneUidFreqMsec returns a boolean if a field has been set.

### GetEddystoneUidInstance

`func (o *ApBle) GetEddystoneUidInstance() string`

GetEddystoneUidInstance returns the EddystoneUidInstance field if non-nil, zero value otherwise.

### GetEddystoneUidInstanceOk

`func (o *ApBle) GetEddystoneUidInstanceOk() (*string, bool)`

GetEddystoneUidInstanceOk returns a tuple with the EddystoneUidInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidInstance

`func (o *ApBle) SetEddystoneUidInstance(v string)`

SetEddystoneUidInstance sets EddystoneUidInstance field to given value.

### HasEddystoneUidInstance

`func (o *ApBle) HasEddystoneUidInstance() bool`

HasEddystoneUidInstance returns a boolean if a field has been set.

### GetEddystoneUidNamespace

`func (o *ApBle) GetEddystoneUidNamespace() string`

GetEddystoneUidNamespace returns the EddystoneUidNamespace field if non-nil, zero value otherwise.

### GetEddystoneUidNamespaceOk

`func (o *ApBle) GetEddystoneUidNamespaceOk() (*string, bool)`

GetEddystoneUidNamespaceOk returns a tuple with the EddystoneUidNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidNamespace

`func (o *ApBle) SetEddystoneUidNamespace(v string)`

SetEddystoneUidNamespace sets EddystoneUidNamespace field to given value.

### HasEddystoneUidNamespace

`func (o *ApBle) HasEddystoneUidNamespace() bool`

HasEddystoneUidNamespace returns a boolean if a field has been set.

### GetEddystoneUrlAdvPower

`func (o *ApBle) GetEddystoneUrlAdvPower() int32`

GetEddystoneUrlAdvPower returns the EddystoneUrlAdvPower field if non-nil, zero value otherwise.

### GetEddystoneUrlAdvPowerOk

`func (o *ApBle) GetEddystoneUrlAdvPowerOk() (*int32, bool)`

GetEddystoneUrlAdvPowerOk returns a tuple with the EddystoneUrlAdvPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlAdvPower

`func (o *ApBle) SetEddystoneUrlAdvPower(v int32)`

SetEddystoneUrlAdvPower sets EddystoneUrlAdvPower field to given value.

### HasEddystoneUrlAdvPower

`func (o *ApBle) HasEddystoneUrlAdvPower() bool`

HasEddystoneUrlAdvPower returns a boolean if a field has been set.

### GetEddystoneUrlBeams

`func (o *ApBle) GetEddystoneUrlBeams() string`

GetEddystoneUrlBeams returns the EddystoneUrlBeams field if non-nil, zero value otherwise.

### GetEddystoneUrlBeamsOk

`func (o *ApBle) GetEddystoneUrlBeamsOk() (*string, bool)`

GetEddystoneUrlBeamsOk returns a tuple with the EddystoneUrlBeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlBeams

`func (o *ApBle) SetEddystoneUrlBeams(v string)`

SetEddystoneUrlBeams sets EddystoneUrlBeams field to given value.

### HasEddystoneUrlBeams

`func (o *ApBle) HasEddystoneUrlBeams() bool`

HasEddystoneUrlBeams returns a boolean if a field has been set.

### GetEddystoneUrlEnabled

`func (o *ApBle) GetEddystoneUrlEnabled() bool`

GetEddystoneUrlEnabled returns the EddystoneUrlEnabled field if non-nil, zero value otherwise.

### GetEddystoneUrlEnabledOk

`func (o *ApBle) GetEddystoneUrlEnabledOk() (*bool, bool)`

GetEddystoneUrlEnabledOk returns a tuple with the EddystoneUrlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlEnabled

`func (o *ApBle) SetEddystoneUrlEnabled(v bool)`

SetEddystoneUrlEnabled sets EddystoneUrlEnabled field to given value.

### HasEddystoneUrlEnabled

`func (o *ApBle) HasEddystoneUrlEnabled() bool`

HasEddystoneUrlEnabled returns a boolean if a field has been set.

### GetEddystoneUrlFreqMsec

`func (o *ApBle) GetEddystoneUrlFreqMsec() int32`

GetEddystoneUrlFreqMsec returns the EddystoneUrlFreqMsec field if non-nil, zero value otherwise.

### GetEddystoneUrlFreqMsecOk

`func (o *ApBle) GetEddystoneUrlFreqMsecOk() (*int32, bool)`

GetEddystoneUrlFreqMsecOk returns a tuple with the EddystoneUrlFreqMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlFreqMsec

`func (o *ApBle) SetEddystoneUrlFreqMsec(v int32)`

SetEddystoneUrlFreqMsec sets EddystoneUrlFreqMsec field to given value.

### HasEddystoneUrlFreqMsec

`func (o *ApBle) HasEddystoneUrlFreqMsec() bool`

HasEddystoneUrlFreqMsec returns a boolean if a field has been set.

### GetEddystoneUrlUrl

`func (o *ApBle) GetEddystoneUrlUrl() string`

GetEddystoneUrlUrl returns the EddystoneUrlUrl field if non-nil, zero value otherwise.

### GetEddystoneUrlUrlOk

`func (o *ApBle) GetEddystoneUrlUrlOk() (*string, bool)`

GetEddystoneUrlUrlOk returns a tuple with the EddystoneUrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlUrl

`func (o *ApBle) SetEddystoneUrlUrl(v string)`

SetEddystoneUrlUrl sets EddystoneUrlUrl field to given value.

### HasEddystoneUrlUrl

`func (o *ApBle) HasEddystoneUrlUrl() bool`

HasEddystoneUrlUrl returns a boolean if a field has been set.

### GetIbeaconAdvPower

`func (o *ApBle) GetIbeaconAdvPower() int32`

GetIbeaconAdvPower returns the IbeaconAdvPower field if non-nil, zero value otherwise.

### GetIbeaconAdvPowerOk

`func (o *ApBle) GetIbeaconAdvPowerOk() (*int32, bool)`

GetIbeaconAdvPowerOk returns a tuple with the IbeaconAdvPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconAdvPower

`func (o *ApBle) SetIbeaconAdvPower(v int32)`

SetIbeaconAdvPower sets IbeaconAdvPower field to given value.

### HasIbeaconAdvPower

`func (o *ApBle) HasIbeaconAdvPower() bool`

HasIbeaconAdvPower returns a boolean if a field has been set.

### GetIbeaconBeams

`func (o *ApBle) GetIbeaconBeams() string`

GetIbeaconBeams returns the IbeaconBeams field if non-nil, zero value otherwise.

### GetIbeaconBeamsOk

`func (o *ApBle) GetIbeaconBeamsOk() (*string, bool)`

GetIbeaconBeamsOk returns a tuple with the IbeaconBeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconBeams

`func (o *ApBle) SetIbeaconBeams(v string)`

SetIbeaconBeams sets IbeaconBeams field to given value.

### HasIbeaconBeams

`func (o *ApBle) HasIbeaconBeams() bool`

HasIbeaconBeams returns a boolean if a field has been set.

### GetIbeaconEnabled

`func (o *ApBle) GetIbeaconEnabled() bool`

GetIbeaconEnabled returns the IbeaconEnabled field if non-nil, zero value otherwise.

### GetIbeaconEnabledOk

`func (o *ApBle) GetIbeaconEnabledOk() (*bool, bool)`

GetIbeaconEnabledOk returns a tuple with the IbeaconEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconEnabled

`func (o *ApBle) SetIbeaconEnabled(v bool)`

SetIbeaconEnabled sets IbeaconEnabled field to given value.

### HasIbeaconEnabled

`func (o *ApBle) HasIbeaconEnabled() bool`

HasIbeaconEnabled returns a boolean if a field has been set.

### GetIbeaconFreqMsec

`func (o *ApBle) GetIbeaconFreqMsec() int32`

GetIbeaconFreqMsec returns the IbeaconFreqMsec field if non-nil, zero value otherwise.

### GetIbeaconFreqMsecOk

`func (o *ApBle) GetIbeaconFreqMsecOk() (*int32, bool)`

GetIbeaconFreqMsecOk returns a tuple with the IbeaconFreqMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconFreqMsec

`func (o *ApBle) SetIbeaconFreqMsec(v int32)`

SetIbeaconFreqMsec sets IbeaconFreqMsec field to given value.

### HasIbeaconFreqMsec

`func (o *ApBle) HasIbeaconFreqMsec() bool`

HasIbeaconFreqMsec returns a boolean if a field has been set.

### GetIbeaconMajor

`func (o *ApBle) GetIbeaconMajor() int32`

GetIbeaconMajor returns the IbeaconMajor field if non-nil, zero value otherwise.

### GetIbeaconMajorOk

`func (o *ApBle) GetIbeaconMajorOk() (*int32, bool)`

GetIbeaconMajorOk returns a tuple with the IbeaconMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMajor

`func (o *ApBle) SetIbeaconMajor(v int32)`

SetIbeaconMajor sets IbeaconMajor field to given value.

### HasIbeaconMajor

`func (o *ApBle) HasIbeaconMajor() bool`

HasIbeaconMajor returns a boolean if a field has been set.

### GetIbeaconMinor

`func (o *ApBle) GetIbeaconMinor() int32`

GetIbeaconMinor returns the IbeaconMinor field if non-nil, zero value otherwise.

### GetIbeaconMinorOk

`func (o *ApBle) GetIbeaconMinorOk() (*int32, bool)`

GetIbeaconMinorOk returns a tuple with the IbeaconMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMinor

`func (o *ApBle) SetIbeaconMinor(v int32)`

SetIbeaconMinor sets IbeaconMinor field to given value.

### HasIbeaconMinor

`func (o *ApBle) HasIbeaconMinor() bool`

HasIbeaconMinor returns a boolean if a field has been set.

### GetIbeaconUuid

`func (o *ApBle) GetIbeaconUuid() string`

GetIbeaconUuid returns the IbeaconUuid field if non-nil, zero value otherwise.

### GetIbeaconUuidOk

`func (o *ApBle) GetIbeaconUuidOk() (*string, bool)`

GetIbeaconUuidOk returns a tuple with the IbeaconUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconUuid

`func (o *ApBle) SetIbeaconUuid(v string)`

SetIbeaconUuid sets IbeaconUuid field to given value.

### HasIbeaconUuid

`func (o *ApBle) HasIbeaconUuid() bool`

HasIbeaconUuid returns a boolean if a field has been set.

### GetPower

`func (o *ApBle) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *ApBle) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *ApBle) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *ApBle) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPowerMode

`func (o *ApBle) GetPowerMode() ApBlePowerMode`

GetPowerMode returns the PowerMode field if non-nil, zero value otherwise.

### GetPowerModeOk

`func (o *ApBle) GetPowerModeOk() (*ApBlePowerMode, bool)`

GetPowerModeOk returns a tuple with the PowerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMode

`func (o *ApBle) SetPowerMode(v ApBlePowerMode)`

SetPowerMode sets PowerMode field to given value.

### HasPowerMode

`func (o *ApBle) HasPowerMode() bool`

HasPowerMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


