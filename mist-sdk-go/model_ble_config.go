/*
Mist API

> Version: **2406.1.10** > > Date: **July 1, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.11
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistsdkgo

import (
	"encoding/json"
)

// checks if the BleConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BleConfig{}

// BleConfig BLE AP settings
type BleConfig struct {
	// whether Mist beacons is enabled
	BeaconEnabled *bool `json:"beacon_enabled,omitempty"`
	// required if `beacon_rate_mode`==`custom`, 1-10, in number-beacons-per-second
	BeaconRate *int32 `json:"beacon_rate,omitempty"`
	BeaconRateMode *BleConfigBeaconRateMode `json:"beacon_rate_mode,omitempty"`
	// list of AP BLE location beam numbers (1-8) which should be disabled at the AP and not transmit location information (where beam 1 is oriented at the top the AP, growing counter-clock-wise, with 9 being the omni BLE beam)
	BeamDisabled []int32 `json:"beam_disabled,omitempty"`
	// can be enabled if `beacon_enabled`==`true`, whether to send custom packet
	CustomBlePacketEnabled *bool `json:"custom_ble_packet_enabled,omitempty"`
	// The custom frame to be sent out in this beacon. The frame must be a hexstring
	CustomBlePacketFrame *string `json:"custom_ble_packet_frame,omitempty"`
	// Frequency (msec) of data emitted by custom ble beacon
	CustomBlePacketFreqMsec *int32 `json:"custom_ble_packet_freq_msec,omitempty"`
	// advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
	EddystoneUidAdvPower *int32 `json:"eddystone_uid_adv_power,omitempty"`
	EddystoneUidBeams *string `json:"eddystone_uid_beams,omitempty"`
	// only if `beacon_enabled`==`false`, Whether Eddystone-UID beacon is enabled
	EddystoneUidEnabled *bool `json:"eddystone_uid_enabled,omitempty"`
	// Frequency (msec) of data emmit by Eddystone-UID beacon
	EddystoneUidFreqMsec *int32 `json:"eddystone_uid_freq_msec,omitempty"`
	// Eddystone-UID instance for the device
	EddystoneUidInstance *string `json:"eddystone_uid_instance,omitempty"`
	// Eddystone-UID namespace
	EddystoneUidNamespace *string `json:"eddystone_uid_namespace,omitempty"`
	// advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
	EddystoneUrlAdvPower *int32 `json:"eddystone_url_adv_power,omitempty"`
	EddystoneUrlBeams *string `json:"eddystone_url_beams,omitempty"`
	// only if `beacon_enabled`==`false`, Whether Eddystone-URL beacon is enabled
	EddystoneUrlEnabled *bool `json:"eddystone_url_enabled,omitempty"`
	// Frequency (msec) of data emit by Eddystone-UID beacon
	EddystoneUrlFreqMsec *int32 `json:"eddystone_url_freq_msec,omitempty"`
	// URL pointed by Eddystone-URL beacon
	EddystoneUrlUrl *string `json:"eddystone_url_url,omitempty"`
	// advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
	IbeaconAdvPower *int32 `json:"ibeacon_adv_power,omitempty"`
	IbeaconBeams *string `json:"ibeacon_beams,omitempty"`
	// can be enabled if `beacon_enabled`==`true`, whether to send iBeacon
	IbeaconEnabled *bool `json:"ibeacon_enabled,omitempty"`
	// Frequency (msec) of data emmit for iBeacon
	IbeaconFreqMsec *int32 `json:"ibeacon_freq_msec,omitempty"`
	// Major number for iBeacon
	IbeaconMajor *int32 `json:"ibeacon_major,omitempty"`
	// Minor number for iBeacon
	IbeaconMinor *int32 `json:"ibeacon_minor,omitempty"`
	// optional, if not specified, the same UUID as the beacon will be used
	IbeaconUuid *string `json:"ibeacon_uuid,omitempty"`
	// required if `power_mode`==`custom`
	Power *int32 `json:"power,omitempty"`
	PowerMode *BleConfigPowerMode `json:"power_mode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BleConfig BleConfig

// NewBleConfig instantiates a new BleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBleConfig() *BleConfig {
	this := BleConfig{}
	var beaconEnabled bool = false
	this.BeaconEnabled = &beaconEnabled
	var beaconRateMode BleConfigBeaconRateMode = BLECONFIGBEACONRATEMODE_DEFAULT
	this.BeaconRateMode = &beaconRateMode
	var customBlePacketEnabled bool = false
	this.CustomBlePacketEnabled = &customBlePacketEnabled
	var eddystoneUidEnabled bool = false
	this.EddystoneUidEnabled = &eddystoneUidEnabled
	var ibeaconEnabled bool = false
	this.IbeaconEnabled = &ibeaconEnabled
	var power int32 = 9
	this.Power = &power
	var powerMode BleConfigPowerMode = BLECONFIGPOWERMODE_DEFAULT
	this.PowerMode = &powerMode
	return &this
}

// NewBleConfigWithDefaults instantiates a new BleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBleConfigWithDefaults() *BleConfig {
	this := BleConfig{}
	var beaconEnabled bool = false
	this.BeaconEnabled = &beaconEnabled
	var beaconRateMode BleConfigBeaconRateMode = BLECONFIGBEACONRATEMODE_DEFAULT
	this.BeaconRateMode = &beaconRateMode
	var customBlePacketEnabled bool = false
	this.CustomBlePacketEnabled = &customBlePacketEnabled
	var eddystoneUidEnabled bool = false
	this.EddystoneUidEnabled = &eddystoneUidEnabled
	var ibeaconEnabled bool = false
	this.IbeaconEnabled = &ibeaconEnabled
	var power int32 = 9
	this.Power = &power
	var powerMode BleConfigPowerMode = BLECONFIGPOWERMODE_DEFAULT
	this.PowerMode = &powerMode
	return &this
}

// GetBeaconEnabled returns the BeaconEnabled field value if set, zero value otherwise.
func (o *BleConfig) GetBeaconEnabled() bool {
	if o == nil || IsNil(o.BeaconEnabled) {
		var ret bool
		return ret
	}
	return *o.BeaconEnabled
}

// GetBeaconEnabledOk returns a tuple with the BeaconEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetBeaconEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BeaconEnabled) {
		return nil, false
	}
	return o.BeaconEnabled, true
}

// HasBeaconEnabled returns a boolean if a field has been set.
func (o *BleConfig) HasBeaconEnabled() bool {
	if o != nil && !IsNil(o.BeaconEnabled) {
		return true
	}

	return false
}

// SetBeaconEnabled gets a reference to the given bool and assigns it to the BeaconEnabled field.
func (o *BleConfig) SetBeaconEnabled(v bool) {
	o.BeaconEnabled = &v
}

// GetBeaconRate returns the BeaconRate field value if set, zero value otherwise.
func (o *BleConfig) GetBeaconRate() int32 {
	if o == nil || IsNil(o.BeaconRate) {
		var ret int32
		return ret
	}
	return *o.BeaconRate
}

// GetBeaconRateOk returns a tuple with the BeaconRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetBeaconRateOk() (*int32, bool) {
	if o == nil || IsNil(o.BeaconRate) {
		return nil, false
	}
	return o.BeaconRate, true
}

// HasBeaconRate returns a boolean if a field has been set.
func (o *BleConfig) HasBeaconRate() bool {
	if o != nil && !IsNil(o.BeaconRate) {
		return true
	}

	return false
}

// SetBeaconRate gets a reference to the given int32 and assigns it to the BeaconRate field.
func (o *BleConfig) SetBeaconRate(v int32) {
	o.BeaconRate = &v
}

// GetBeaconRateMode returns the BeaconRateMode field value if set, zero value otherwise.
func (o *BleConfig) GetBeaconRateMode() BleConfigBeaconRateMode {
	if o == nil || IsNil(o.BeaconRateMode) {
		var ret BleConfigBeaconRateMode
		return ret
	}
	return *o.BeaconRateMode
}

// GetBeaconRateModeOk returns a tuple with the BeaconRateMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetBeaconRateModeOk() (*BleConfigBeaconRateMode, bool) {
	if o == nil || IsNil(o.BeaconRateMode) {
		return nil, false
	}
	return o.BeaconRateMode, true
}

// HasBeaconRateMode returns a boolean if a field has been set.
func (o *BleConfig) HasBeaconRateMode() bool {
	if o != nil && !IsNil(o.BeaconRateMode) {
		return true
	}

	return false
}

// SetBeaconRateMode gets a reference to the given BleConfigBeaconRateMode and assigns it to the BeaconRateMode field.
func (o *BleConfig) SetBeaconRateMode(v BleConfigBeaconRateMode) {
	o.BeaconRateMode = &v
}

// GetBeamDisabled returns the BeamDisabled field value if set, zero value otherwise.
func (o *BleConfig) GetBeamDisabled() []int32 {
	if o == nil || IsNil(o.BeamDisabled) {
		var ret []int32
		return ret
	}
	return o.BeamDisabled
}

// GetBeamDisabledOk returns a tuple with the BeamDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetBeamDisabledOk() ([]int32, bool) {
	if o == nil || IsNil(o.BeamDisabled) {
		return nil, false
	}
	return o.BeamDisabled, true
}

// HasBeamDisabled returns a boolean if a field has been set.
func (o *BleConfig) HasBeamDisabled() bool {
	if o != nil && !IsNil(o.BeamDisabled) {
		return true
	}

	return false
}

// SetBeamDisabled gets a reference to the given []int32 and assigns it to the BeamDisabled field.
func (o *BleConfig) SetBeamDisabled(v []int32) {
	o.BeamDisabled = v
}

// GetCustomBlePacketEnabled returns the CustomBlePacketEnabled field value if set, zero value otherwise.
func (o *BleConfig) GetCustomBlePacketEnabled() bool {
	if o == nil || IsNil(o.CustomBlePacketEnabled) {
		var ret bool
		return ret
	}
	return *o.CustomBlePacketEnabled
}

// GetCustomBlePacketEnabledOk returns a tuple with the CustomBlePacketEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetCustomBlePacketEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CustomBlePacketEnabled) {
		return nil, false
	}
	return o.CustomBlePacketEnabled, true
}

// HasCustomBlePacketEnabled returns a boolean if a field has been set.
func (o *BleConfig) HasCustomBlePacketEnabled() bool {
	if o != nil && !IsNil(o.CustomBlePacketEnabled) {
		return true
	}

	return false
}

// SetCustomBlePacketEnabled gets a reference to the given bool and assigns it to the CustomBlePacketEnabled field.
func (o *BleConfig) SetCustomBlePacketEnabled(v bool) {
	o.CustomBlePacketEnabled = &v
}

// GetCustomBlePacketFrame returns the CustomBlePacketFrame field value if set, zero value otherwise.
func (o *BleConfig) GetCustomBlePacketFrame() string {
	if o == nil || IsNil(o.CustomBlePacketFrame) {
		var ret string
		return ret
	}
	return *o.CustomBlePacketFrame
}

// GetCustomBlePacketFrameOk returns a tuple with the CustomBlePacketFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetCustomBlePacketFrameOk() (*string, bool) {
	if o == nil || IsNil(o.CustomBlePacketFrame) {
		return nil, false
	}
	return o.CustomBlePacketFrame, true
}

// HasCustomBlePacketFrame returns a boolean if a field has been set.
func (o *BleConfig) HasCustomBlePacketFrame() bool {
	if o != nil && !IsNil(o.CustomBlePacketFrame) {
		return true
	}

	return false
}

// SetCustomBlePacketFrame gets a reference to the given string and assigns it to the CustomBlePacketFrame field.
func (o *BleConfig) SetCustomBlePacketFrame(v string) {
	o.CustomBlePacketFrame = &v
}

// GetCustomBlePacketFreqMsec returns the CustomBlePacketFreqMsec field value if set, zero value otherwise.
func (o *BleConfig) GetCustomBlePacketFreqMsec() int32 {
	if o == nil || IsNil(o.CustomBlePacketFreqMsec) {
		var ret int32
		return ret
	}
	return *o.CustomBlePacketFreqMsec
}

// GetCustomBlePacketFreqMsecOk returns a tuple with the CustomBlePacketFreqMsec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetCustomBlePacketFreqMsecOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomBlePacketFreqMsec) {
		return nil, false
	}
	return o.CustomBlePacketFreqMsec, true
}

// HasCustomBlePacketFreqMsec returns a boolean if a field has been set.
func (o *BleConfig) HasCustomBlePacketFreqMsec() bool {
	if o != nil && !IsNil(o.CustomBlePacketFreqMsec) {
		return true
	}

	return false
}

// SetCustomBlePacketFreqMsec gets a reference to the given int32 and assigns it to the CustomBlePacketFreqMsec field.
func (o *BleConfig) SetCustomBlePacketFreqMsec(v int32) {
	o.CustomBlePacketFreqMsec = &v
}

// GetEddystoneUidAdvPower returns the EddystoneUidAdvPower field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUidAdvPower() int32 {
	if o == nil || IsNil(o.EddystoneUidAdvPower) {
		var ret int32
		return ret
	}
	return *o.EddystoneUidAdvPower
}

// GetEddystoneUidAdvPowerOk returns a tuple with the EddystoneUidAdvPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUidAdvPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.EddystoneUidAdvPower) {
		return nil, false
	}
	return o.EddystoneUidAdvPower, true
}

// HasEddystoneUidAdvPower returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUidAdvPower() bool {
	if o != nil && !IsNil(o.EddystoneUidAdvPower) {
		return true
	}

	return false
}

// SetEddystoneUidAdvPower gets a reference to the given int32 and assigns it to the EddystoneUidAdvPower field.
func (o *BleConfig) SetEddystoneUidAdvPower(v int32) {
	o.EddystoneUidAdvPower = &v
}

// GetEddystoneUidBeams returns the EddystoneUidBeams field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUidBeams() string {
	if o == nil || IsNil(o.EddystoneUidBeams) {
		var ret string
		return ret
	}
	return *o.EddystoneUidBeams
}

// GetEddystoneUidBeamsOk returns a tuple with the EddystoneUidBeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUidBeamsOk() (*string, bool) {
	if o == nil || IsNil(o.EddystoneUidBeams) {
		return nil, false
	}
	return o.EddystoneUidBeams, true
}

// HasEddystoneUidBeams returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUidBeams() bool {
	if o != nil && !IsNil(o.EddystoneUidBeams) {
		return true
	}

	return false
}

// SetEddystoneUidBeams gets a reference to the given string and assigns it to the EddystoneUidBeams field.
func (o *BleConfig) SetEddystoneUidBeams(v string) {
	o.EddystoneUidBeams = &v
}

// GetEddystoneUidEnabled returns the EddystoneUidEnabled field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUidEnabled() bool {
	if o == nil || IsNil(o.EddystoneUidEnabled) {
		var ret bool
		return ret
	}
	return *o.EddystoneUidEnabled
}

// GetEddystoneUidEnabledOk returns a tuple with the EddystoneUidEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUidEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EddystoneUidEnabled) {
		return nil, false
	}
	return o.EddystoneUidEnabled, true
}

// HasEddystoneUidEnabled returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUidEnabled() bool {
	if o != nil && !IsNil(o.EddystoneUidEnabled) {
		return true
	}

	return false
}

// SetEddystoneUidEnabled gets a reference to the given bool and assigns it to the EddystoneUidEnabled field.
func (o *BleConfig) SetEddystoneUidEnabled(v bool) {
	o.EddystoneUidEnabled = &v
}

// GetEddystoneUidFreqMsec returns the EddystoneUidFreqMsec field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUidFreqMsec() int32 {
	if o == nil || IsNil(o.EddystoneUidFreqMsec) {
		var ret int32
		return ret
	}
	return *o.EddystoneUidFreqMsec
}

// GetEddystoneUidFreqMsecOk returns a tuple with the EddystoneUidFreqMsec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUidFreqMsecOk() (*int32, bool) {
	if o == nil || IsNil(o.EddystoneUidFreqMsec) {
		return nil, false
	}
	return o.EddystoneUidFreqMsec, true
}

// HasEddystoneUidFreqMsec returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUidFreqMsec() bool {
	if o != nil && !IsNil(o.EddystoneUidFreqMsec) {
		return true
	}

	return false
}

// SetEddystoneUidFreqMsec gets a reference to the given int32 and assigns it to the EddystoneUidFreqMsec field.
func (o *BleConfig) SetEddystoneUidFreqMsec(v int32) {
	o.EddystoneUidFreqMsec = &v
}

// GetEddystoneUidInstance returns the EddystoneUidInstance field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUidInstance() string {
	if o == nil || IsNil(o.EddystoneUidInstance) {
		var ret string
		return ret
	}
	return *o.EddystoneUidInstance
}

// GetEddystoneUidInstanceOk returns a tuple with the EddystoneUidInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUidInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.EddystoneUidInstance) {
		return nil, false
	}
	return o.EddystoneUidInstance, true
}

// HasEddystoneUidInstance returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUidInstance() bool {
	if o != nil && !IsNil(o.EddystoneUidInstance) {
		return true
	}

	return false
}

// SetEddystoneUidInstance gets a reference to the given string and assigns it to the EddystoneUidInstance field.
func (o *BleConfig) SetEddystoneUidInstance(v string) {
	o.EddystoneUidInstance = &v
}

// GetEddystoneUidNamespace returns the EddystoneUidNamespace field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUidNamespace() string {
	if o == nil || IsNil(o.EddystoneUidNamespace) {
		var ret string
		return ret
	}
	return *o.EddystoneUidNamespace
}

// GetEddystoneUidNamespaceOk returns a tuple with the EddystoneUidNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUidNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.EddystoneUidNamespace) {
		return nil, false
	}
	return o.EddystoneUidNamespace, true
}

// HasEddystoneUidNamespace returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUidNamespace() bool {
	if o != nil && !IsNil(o.EddystoneUidNamespace) {
		return true
	}

	return false
}

// SetEddystoneUidNamespace gets a reference to the given string and assigns it to the EddystoneUidNamespace field.
func (o *BleConfig) SetEddystoneUidNamespace(v string) {
	o.EddystoneUidNamespace = &v
}

// GetEddystoneUrlAdvPower returns the EddystoneUrlAdvPower field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUrlAdvPower() int32 {
	if o == nil || IsNil(o.EddystoneUrlAdvPower) {
		var ret int32
		return ret
	}
	return *o.EddystoneUrlAdvPower
}

// GetEddystoneUrlAdvPowerOk returns a tuple with the EddystoneUrlAdvPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUrlAdvPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.EddystoneUrlAdvPower) {
		return nil, false
	}
	return o.EddystoneUrlAdvPower, true
}

// HasEddystoneUrlAdvPower returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUrlAdvPower() bool {
	if o != nil && !IsNil(o.EddystoneUrlAdvPower) {
		return true
	}

	return false
}

// SetEddystoneUrlAdvPower gets a reference to the given int32 and assigns it to the EddystoneUrlAdvPower field.
func (o *BleConfig) SetEddystoneUrlAdvPower(v int32) {
	o.EddystoneUrlAdvPower = &v
}

// GetEddystoneUrlBeams returns the EddystoneUrlBeams field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUrlBeams() string {
	if o == nil || IsNil(o.EddystoneUrlBeams) {
		var ret string
		return ret
	}
	return *o.EddystoneUrlBeams
}

// GetEddystoneUrlBeamsOk returns a tuple with the EddystoneUrlBeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUrlBeamsOk() (*string, bool) {
	if o == nil || IsNil(o.EddystoneUrlBeams) {
		return nil, false
	}
	return o.EddystoneUrlBeams, true
}

// HasEddystoneUrlBeams returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUrlBeams() bool {
	if o != nil && !IsNil(o.EddystoneUrlBeams) {
		return true
	}

	return false
}

// SetEddystoneUrlBeams gets a reference to the given string and assigns it to the EddystoneUrlBeams field.
func (o *BleConfig) SetEddystoneUrlBeams(v string) {
	o.EddystoneUrlBeams = &v
}

// GetEddystoneUrlEnabled returns the EddystoneUrlEnabled field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUrlEnabled() bool {
	if o == nil || IsNil(o.EddystoneUrlEnabled) {
		var ret bool
		return ret
	}
	return *o.EddystoneUrlEnabled
}

// GetEddystoneUrlEnabledOk returns a tuple with the EddystoneUrlEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUrlEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EddystoneUrlEnabled) {
		return nil, false
	}
	return o.EddystoneUrlEnabled, true
}

// HasEddystoneUrlEnabled returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUrlEnabled() bool {
	if o != nil && !IsNil(o.EddystoneUrlEnabled) {
		return true
	}

	return false
}

// SetEddystoneUrlEnabled gets a reference to the given bool and assigns it to the EddystoneUrlEnabled field.
func (o *BleConfig) SetEddystoneUrlEnabled(v bool) {
	o.EddystoneUrlEnabled = &v
}

// GetEddystoneUrlFreqMsec returns the EddystoneUrlFreqMsec field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUrlFreqMsec() int32 {
	if o == nil || IsNil(o.EddystoneUrlFreqMsec) {
		var ret int32
		return ret
	}
	return *o.EddystoneUrlFreqMsec
}

// GetEddystoneUrlFreqMsecOk returns a tuple with the EddystoneUrlFreqMsec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUrlFreqMsecOk() (*int32, bool) {
	if o == nil || IsNil(o.EddystoneUrlFreqMsec) {
		return nil, false
	}
	return o.EddystoneUrlFreqMsec, true
}

// HasEddystoneUrlFreqMsec returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUrlFreqMsec() bool {
	if o != nil && !IsNil(o.EddystoneUrlFreqMsec) {
		return true
	}

	return false
}

// SetEddystoneUrlFreqMsec gets a reference to the given int32 and assigns it to the EddystoneUrlFreqMsec field.
func (o *BleConfig) SetEddystoneUrlFreqMsec(v int32) {
	o.EddystoneUrlFreqMsec = &v
}

// GetEddystoneUrlUrl returns the EddystoneUrlUrl field value if set, zero value otherwise.
func (o *BleConfig) GetEddystoneUrlUrl() string {
	if o == nil || IsNil(o.EddystoneUrlUrl) {
		var ret string
		return ret
	}
	return *o.EddystoneUrlUrl
}

// GetEddystoneUrlUrlOk returns a tuple with the EddystoneUrlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetEddystoneUrlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EddystoneUrlUrl) {
		return nil, false
	}
	return o.EddystoneUrlUrl, true
}

// HasEddystoneUrlUrl returns a boolean if a field has been set.
func (o *BleConfig) HasEddystoneUrlUrl() bool {
	if o != nil && !IsNil(o.EddystoneUrlUrl) {
		return true
	}

	return false
}

// SetEddystoneUrlUrl gets a reference to the given string and assigns it to the EddystoneUrlUrl field.
func (o *BleConfig) SetEddystoneUrlUrl(v string) {
	o.EddystoneUrlUrl = &v
}

// GetIbeaconAdvPower returns the IbeaconAdvPower field value if set, zero value otherwise.
func (o *BleConfig) GetIbeaconAdvPower() int32 {
	if o == nil || IsNil(o.IbeaconAdvPower) {
		var ret int32
		return ret
	}
	return *o.IbeaconAdvPower
}

// GetIbeaconAdvPowerOk returns a tuple with the IbeaconAdvPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetIbeaconAdvPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.IbeaconAdvPower) {
		return nil, false
	}
	return o.IbeaconAdvPower, true
}

// HasIbeaconAdvPower returns a boolean if a field has been set.
func (o *BleConfig) HasIbeaconAdvPower() bool {
	if o != nil && !IsNil(o.IbeaconAdvPower) {
		return true
	}

	return false
}

// SetIbeaconAdvPower gets a reference to the given int32 and assigns it to the IbeaconAdvPower field.
func (o *BleConfig) SetIbeaconAdvPower(v int32) {
	o.IbeaconAdvPower = &v
}

// GetIbeaconBeams returns the IbeaconBeams field value if set, zero value otherwise.
func (o *BleConfig) GetIbeaconBeams() string {
	if o == nil || IsNil(o.IbeaconBeams) {
		var ret string
		return ret
	}
	return *o.IbeaconBeams
}

// GetIbeaconBeamsOk returns a tuple with the IbeaconBeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetIbeaconBeamsOk() (*string, bool) {
	if o == nil || IsNil(o.IbeaconBeams) {
		return nil, false
	}
	return o.IbeaconBeams, true
}

// HasIbeaconBeams returns a boolean if a field has been set.
func (o *BleConfig) HasIbeaconBeams() bool {
	if o != nil && !IsNil(o.IbeaconBeams) {
		return true
	}

	return false
}

// SetIbeaconBeams gets a reference to the given string and assigns it to the IbeaconBeams field.
func (o *BleConfig) SetIbeaconBeams(v string) {
	o.IbeaconBeams = &v
}

// GetIbeaconEnabled returns the IbeaconEnabled field value if set, zero value otherwise.
func (o *BleConfig) GetIbeaconEnabled() bool {
	if o == nil || IsNil(o.IbeaconEnabled) {
		var ret bool
		return ret
	}
	return *o.IbeaconEnabled
}

// GetIbeaconEnabledOk returns a tuple with the IbeaconEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetIbeaconEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IbeaconEnabled) {
		return nil, false
	}
	return o.IbeaconEnabled, true
}

// HasIbeaconEnabled returns a boolean if a field has been set.
func (o *BleConfig) HasIbeaconEnabled() bool {
	if o != nil && !IsNil(o.IbeaconEnabled) {
		return true
	}

	return false
}

// SetIbeaconEnabled gets a reference to the given bool and assigns it to the IbeaconEnabled field.
func (o *BleConfig) SetIbeaconEnabled(v bool) {
	o.IbeaconEnabled = &v
}

// GetIbeaconFreqMsec returns the IbeaconFreqMsec field value if set, zero value otherwise.
func (o *BleConfig) GetIbeaconFreqMsec() int32 {
	if o == nil || IsNil(o.IbeaconFreqMsec) {
		var ret int32
		return ret
	}
	return *o.IbeaconFreqMsec
}

// GetIbeaconFreqMsecOk returns a tuple with the IbeaconFreqMsec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetIbeaconFreqMsecOk() (*int32, bool) {
	if o == nil || IsNil(o.IbeaconFreqMsec) {
		return nil, false
	}
	return o.IbeaconFreqMsec, true
}

// HasIbeaconFreqMsec returns a boolean if a field has been set.
func (o *BleConfig) HasIbeaconFreqMsec() bool {
	if o != nil && !IsNil(o.IbeaconFreqMsec) {
		return true
	}

	return false
}

// SetIbeaconFreqMsec gets a reference to the given int32 and assigns it to the IbeaconFreqMsec field.
func (o *BleConfig) SetIbeaconFreqMsec(v int32) {
	o.IbeaconFreqMsec = &v
}

// GetIbeaconMajor returns the IbeaconMajor field value if set, zero value otherwise.
func (o *BleConfig) GetIbeaconMajor() int32 {
	if o == nil || IsNil(o.IbeaconMajor) {
		var ret int32
		return ret
	}
	return *o.IbeaconMajor
}

// GetIbeaconMajorOk returns a tuple with the IbeaconMajor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetIbeaconMajorOk() (*int32, bool) {
	if o == nil || IsNil(o.IbeaconMajor) {
		return nil, false
	}
	return o.IbeaconMajor, true
}

// HasIbeaconMajor returns a boolean if a field has been set.
func (o *BleConfig) HasIbeaconMajor() bool {
	if o != nil && !IsNil(o.IbeaconMajor) {
		return true
	}

	return false
}

// SetIbeaconMajor gets a reference to the given int32 and assigns it to the IbeaconMajor field.
func (o *BleConfig) SetIbeaconMajor(v int32) {
	o.IbeaconMajor = &v
}

// GetIbeaconMinor returns the IbeaconMinor field value if set, zero value otherwise.
func (o *BleConfig) GetIbeaconMinor() int32 {
	if o == nil || IsNil(o.IbeaconMinor) {
		var ret int32
		return ret
	}
	return *o.IbeaconMinor
}

// GetIbeaconMinorOk returns a tuple with the IbeaconMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetIbeaconMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.IbeaconMinor) {
		return nil, false
	}
	return o.IbeaconMinor, true
}

// HasIbeaconMinor returns a boolean if a field has been set.
func (o *BleConfig) HasIbeaconMinor() bool {
	if o != nil && !IsNil(o.IbeaconMinor) {
		return true
	}

	return false
}

// SetIbeaconMinor gets a reference to the given int32 and assigns it to the IbeaconMinor field.
func (o *BleConfig) SetIbeaconMinor(v int32) {
	o.IbeaconMinor = &v
}

// GetIbeaconUuid returns the IbeaconUuid field value if set, zero value otherwise.
func (o *BleConfig) GetIbeaconUuid() string {
	if o == nil || IsNil(o.IbeaconUuid) {
		var ret string
		return ret
	}
	return *o.IbeaconUuid
}

// GetIbeaconUuidOk returns a tuple with the IbeaconUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetIbeaconUuidOk() (*string, bool) {
	if o == nil || IsNil(o.IbeaconUuid) {
		return nil, false
	}
	return o.IbeaconUuid, true
}

// HasIbeaconUuid returns a boolean if a field has been set.
func (o *BleConfig) HasIbeaconUuid() bool {
	if o != nil && !IsNil(o.IbeaconUuid) {
		return true
	}

	return false
}

// SetIbeaconUuid gets a reference to the given string and assigns it to the IbeaconUuid field.
func (o *BleConfig) SetIbeaconUuid(v string) {
	o.IbeaconUuid = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *BleConfig) GetPower() int32 {
	if o == nil || IsNil(o.Power) {
		var ret int32
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.Power) {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *BleConfig) HasPower() bool {
	if o != nil && !IsNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given int32 and assigns it to the Power field.
func (o *BleConfig) SetPower(v int32) {
	o.Power = &v
}

// GetPowerMode returns the PowerMode field value if set, zero value otherwise.
func (o *BleConfig) GetPowerMode() BleConfigPowerMode {
	if o == nil || IsNil(o.PowerMode) {
		var ret BleConfigPowerMode
		return ret
	}
	return *o.PowerMode
}

// GetPowerModeOk returns a tuple with the PowerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BleConfig) GetPowerModeOk() (*BleConfigPowerMode, bool) {
	if o == nil || IsNil(o.PowerMode) {
		return nil, false
	}
	return o.PowerMode, true
}

// HasPowerMode returns a boolean if a field has been set.
func (o *BleConfig) HasPowerMode() bool {
	if o != nil && !IsNil(o.PowerMode) {
		return true
	}

	return false
}

// SetPowerMode gets a reference to the given BleConfigPowerMode and assigns it to the PowerMode field.
func (o *BleConfig) SetPowerMode(v BleConfigPowerMode) {
	o.PowerMode = &v
}

func (o BleConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BleConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeaconEnabled) {
		toSerialize["beacon_enabled"] = o.BeaconEnabled
	}
	if !IsNil(o.BeaconRate) {
		toSerialize["beacon_rate"] = o.BeaconRate
	}
	if !IsNil(o.BeaconRateMode) {
		toSerialize["beacon_rate_mode"] = o.BeaconRateMode
	}
	if !IsNil(o.BeamDisabled) {
		toSerialize["beam_disabled"] = o.BeamDisabled
	}
	if !IsNil(o.CustomBlePacketEnabled) {
		toSerialize["custom_ble_packet_enabled"] = o.CustomBlePacketEnabled
	}
	if !IsNil(o.CustomBlePacketFrame) {
		toSerialize["custom_ble_packet_frame"] = o.CustomBlePacketFrame
	}
	if !IsNil(o.CustomBlePacketFreqMsec) {
		toSerialize["custom_ble_packet_freq_msec"] = o.CustomBlePacketFreqMsec
	}
	if !IsNil(o.EddystoneUidAdvPower) {
		toSerialize["eddystone_uid_adv_power"] = o.EddystoneUidAdvPower
	}
	if !IsNil(o.EddystoneUidBeams) {
		toSerialize["eddystone_uid_beams"] = o.EddystoneUidBeams
	}
	if !IsNil(o.EddystoneUidEnabled) {
		toSerialize["eddystone_uid_enabled"] = o.EddystoneUidEnabled
	}
	if !IsNil(o.EddystoneUidFreqMsec) {
		toSerialize["eddystone_uid_freq_msec"] = o.EddystoneUidFreqMsec
	}
	if !IsNil(o.EddystoneUidInstance) {
		toSerialize["eddystone_uid_instance"] = o.EddystoneUidInstance
	}
	if !IsNil(o.EddystoneUidNamespace) {
		toSerialize["eddystone_uid_namespace"] = o.EddystoneUidNamespace
	}
	if !IsNil(o.EddystoneUrlAdvPower) {
		toSerialize["eddystone_url_adv_power"] = o.EddystoneUrlAdvPower
	}
	if !IsNil(o.EddystoneUrlBeams) {
		toSerialize["eddystone_url_beams"] = o.EddystoneUrlBeams
	}
	if !IsNil(o.EddystoneUrlEnabled) {
		toSerialize["eddystone_url_enabled"] = o.EddystoneUrlEnabled
	}
	if !IsNil(o.EddystoneUrlFreqMsec) {
		toSerialize["eddystone_url_freq_msec"] = o.EddystoneUrlFreqMsec
	}
	if !IsNil(o.EddystoneUrlUrl) {
		toSerialize["eddystone_url_url"] = o.EddystoneUrlUrl
	}
	if !IsNil(o.IbeaconAdvPower) {
		toSerialize["ibeacon_adv_power"] = o.IbeaconAdvPower
	}
	if !IsNil(o.IbeaconBeams) {
		toSerialize["ibeacon_beams"] = o.IbeaconBeams
	}
	if !IsNil(o.IbeaconEnabled) {
		toSerialize["ibeacon_enabled"] = o.IbeaconEnabled
	}
	if !IsNil(o.IbeaconFreqMsec) {
		toSerialize["ibeacon_freq_msec"] = o.IbeaconFreqMsec
	}
	if !IsNil(o.IbeaconMajor) {
		toSerialize["ibeacon_major"] = o.IbeaconMajor
	}
	if !IsNil(o.IbeaconMinor) {
		toSerialize["ibeacon_minor"] = o.IbeaconMinor
	}
	if !IsNil(o.IbeaconUuid) {
		toSerialize["ibeacon_uuid"] = o.IbeaconUuid
	}
	if !IsNil(o.Power) {
		toSerialize["power"] = o.Power
	}
	if !IsNil(o.PowerMode) {
		toSerialize["power_mode"] = o.PowerMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BleConfig) UnmarshalJSON(data []byte) (err error) {
	varBleConfig := _BleConfig{}

	err = json.Unmarshal(data, &varBleConfig)

	if err != nil {
		return err
	}

	*o = BleConfig(varBleConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "beacon_enabled")
		delete(additionalProperties, "beacon_rate")
		delete(additionalProperties, "beacon_rate_mode")
		delete(additionalProperties, "beam_disabled")
		delete(additionalProperties, "custom_ble_packet_enabled")
		delete(additionalProperties, "custom_ble_packet_frame")
		delete(additionalProperties, "custom_ble_packet_freq_msec")
		delete(additionalProperties, "eddystone_uid_adv_power")
		delete(additionalProperties, "eddystone_uid_beams")
		delete(additionalProperties, "eddystone_uid_enabled")
		delete(additionalProperties, "eddystone_uid_freq_msec")
		delete(additionalProperties, "eddystone_uid_instance")
		delete(additionalProperties, "eddystone_uid_namespace")
		delete(additionalProperties, "eddystone_url_adv_power")
		delete(additionalProperties, "eddystone_url_beams")
		delete(additionalProperties, "eddystone_url_enabled")
		delete(additionalProperties, "eddystone_url_freq_msec")
		delete(additionalProperties, "eddystone_url_url")
		delete(additionalProperties, "ibeacon_adv_power")
		delete(additionalProperties, "ibeacon_beams")
		delete(additionalProperties, "ibeacon_enabled")
		delete(additionalProperties, "ibeacon_freq_msec")
		delete(additionalProperties, "ibeacon_major")
		delete(additionalProperties, "ibeacon_minor")
		delete(additionalProperties, "ibeacon_uuid")
		delete(additionalProperties, "power")
		delete(additionalProperties, "power_mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBleConfig struct {
	value *BleConfig
	isSet bool
}

func (v NullableBleConfig) Get() *BleConfig {
	return v.value
}

func (v *NullableBleConfig) Set(val *BleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBleConfig(val *BleConfig) *NullableBleConfig {
	return &NullableBleConfig{value: val, isSet: true}
}

func (v NullableBleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

