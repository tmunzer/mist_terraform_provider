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

// checks if the UtilsSendBleBeacon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsSendBleBeacon{}

// UtilsSendBleBeacon struct for UtilsSendBleBeacon
type UtilsSendBleBeacon struct {
	BeaconFrame *string `json:"beacon_frame,omitempty"`
	BeaconFreq *int32 `json:"beacon_freq,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	Macs []string `json:"macs,omitempty"`
	MapIds []string `json:"map_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UtilsSendBleBeacon UtilsSendBleBeacon

// NewUtilsSendBleBeacon instantiates a new UtilsSendBleBeacon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsSendBleBeacon() *UtilsSendBleBeacon {
	this := UtilsSendBleBeacon{}
	return &this
}

// NewUtilsSendBleBeaconWithDefaults instantiates a new UtilsSendBleBeacon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsSendBleBeaconWithDefaults() *UtilsSendBleBeacon {
	this := UtilsSendBleBeacon{}
	return &this
}

// GetBeaconFrame returns the BeaconFrame field value if set, zero value otherwise.
func (o *UtilsSendBleBeacon) GetBeaconFrame() string {
	if o == nil || IsNil(o.BeaconFrame) {
		var ret string
		return ret
	}
	return *o.BeaconFrame
}

// GetBeaconFrameOk returns a tuple with the BeaconFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsSendBleBeacon) GetBeaconFrameOk() (*string, bool) {
	if o == nil || IsNil(o.BeaconFrame) {
		return nil, false
	}
	return o.BeaconFrame, true
}

// HasBeaconFrame returns a boolean if a field has been set.
func (o *UtilsSendBleBeacon) HasBeaconFrame() bool {
	if o != nil && !IsNil(o.BeaconFrame) {
		return true
	}

	return false
}

// SetBeaconFrame gets a reference to the given string and assigns it to the BeaconFrame field.
func (o *UtilsSendBleBeacon) SetBeaconFrame(v string) {
	o.BeaconFrame = &v
}

// GetBeaconFreq returns the BeaconFreq field value if set, zero value otherwise.
func (o *UtilsSendBleBeacon) GetBeaconFreq() int32 {
	if o == nil || IsNil(o.BeaconFreq) {
		var ret int32
		return ret
	}
	return *o.BeaconFreq
}

// GetBeaconFreqOk returns a tuple with the BeaconFreq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsSendBleBeacon) GetBeaconFreqOk() (*int32, bool) {
	if o == nil || IsNil(o.BeaconFreq) {
		return nil, false
	}
	return o.BeaconFreq, true
}

// HasBeaconFreq returns a boolean if a field has been set.
func (o *UtilsSendBleBeacon) HasBeaconFreq() bool {
	if o != nil && !IsNil(o.BeaconFreq) {
		return true
	}

	return false
}

// SetBeaconFreq gets a reference to the given int32 and assigns it to the BeaconFreq field.
func (o *UtilsSendBleBeacon) SetBeaconFreq(v int32) {
	o.BeaconFreq = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *UtilsSendBleBeacon) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsSendBleBeacon) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *UtilsSendBleBeacon) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *UtilsSendBleBeacon) SetDuration(v int32) {
	o.Duration = &v
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *UtilsSendBleBeacon) GetMacs() []string {
	if o == nil || IsNil(o.Macs) {
		var ret []string
		return ret
	}
	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsSendBleBeacon) GetMacsOk() ([]string, bool) {
	if o == nil || IsNil(o.Macs) {
		return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *UtilsSendBleBeacon) HasMacs() bool {
	if o != nil && !IsNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given []string and assigns it to the Macs field.
func (o *UtilsSendBleBeacon) SetMacs(v []string) {
	o.Macs = v
}

// GetMapIds returns the MapIds field value if set, zero value otherwise.
func (o *UtilsSendBleBeacon) GetMapIds() []string {
	if o == nil || IsNil(o.MapIds) {
		var ret []string
		return ret
	}
	return o.MapIds
}

// GetMapIdsOk returns a tuple with the MapIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsSendBleBeacon) GetMapIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MapIds) {
		return nil, false
	}
	return o.MapIds, true
}

// HasMapIds returns a boolean if a field has been set.
func (o *UtilsSendBleBeacon) HasMapIds() bool {
	if o != nil && !IsNil(o.MapIds) {
		return true
	}

	return false
}

// SetMapIds gets a reference to the given []string and assigns it to the MapIds field.
func (o *UtilsSendBleBeacon) SetMapIds(v []string) {
	o.MapIds = v
}

func (o UtilsSendBleBeacon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsSendBleBeacon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeaconFrame) {
		toSerialize["beacon_frame"] = o.BeaconFrame
	}
	if !IsNil(o.BeaconFreq) {
		toSerialize["beacon_freq"] = o.BeaconFreq
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Macs) {
		toSerialize["macs"] = o.Macs
	}
	if !IsNil(o.MapIds) {
		toSerialize["map_ids"] = o.MapIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UtilsSendBleBeacon) UnmarshalJSON(data []byte) (err error) {
	varUtilsSendBleBeacon := _UtilsSendBleBeacon{}

	err = json.Unmarshal(data, &varUtilsSendBleBeacon)

	if err != nil {
		return err
	}

	*o = UtilsSendBleBeacon(varUtilsSendBleBeacon)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "beacon_frame")
		delete(additionalProperties, "beacon_freq")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "macs")
		delete(additionalProperties, "map_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUtilsSendBleBeacon struct {
	value *UtilsSendBleBeacon
	isSet bool
}

func (v NullableUtilsSendBleBeacon) Get() *UtilsSendBleBeacon {
	return v.value
}

func (v *NullableUtilsSendBleBeacon) Set(val *UtilsSendBleBeacon) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsSendBleBeacon) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsSendBleBeacon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsSendBleBeacon(val *UtilsSendBleBeacon) *NullableUtilsSendBleBeacon {
	return &NullableUtilsSendBleBeacon{value: val, isSet: true}
}

func (v NullableUtilsSendBleBeacon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsSendBleBeacon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

