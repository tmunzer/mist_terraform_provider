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
	"fmt"
)

// checks if the RrmConsideration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RrmConsideration{}

// RrmConsideration struct for RrmConsideration
type RrmConsideration struct {
	Channel int32 `json:"channel"`
	Noise float32 `json:"noise"`
	// the avg RSSI heard from other APs (that does NOT belongs to the same site)
	OtherRssi *float32 `json:"other_rssi,omitempty"`
	// SSID from other AP that we heard from with the max RSSI
	OtherSsid *string `json:"other_ssid,omitempty"`
	// utilization score, 0-1, lower means less utilization (cleaner RF)
	UtilScore float32 `json:"util_score"`
	// non-wifi utilization score, 0-1, lower means less utilization (cleaner RF)
	UtilScoreNonWifi float32 `json:"util_score_non_wifi"`
	// other utilization score, 0-1, lower means less utilization (cleaner RF)
	UtilScoreOther float32 `json:"util_score_other"`
	AdditionalProperties map[string]interface{}
}

type _RrmConsideration RrmConsideration

// NewRrmConsideration instantiates a new RrmConsideration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRrmConsideration(channel int32, noise float32, utilScore float32, utilScoreNonWifi float32, utilScoreOther float32) *RrmConsideration {
	this := RrmConsideration{}
	this.Channel = channel
	this.Noise = noise
	this.UtilScore = utilScore
	this.UtilScoreNonWifi = utilScoreNonWifi
	this.UtilScoreOther = utilScoreOther
	return &this
}

// NewRrmConsiderationWithDefaults instantiates a new RrmConsideration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRrmConsiderationWithDefaults() *RrmConsideration {
	this := RrmConsideration{}
	return &this
}

// GetChannel returns the Channel field value
func (o *RrmConsideration) GetChannel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *RrmConsideration) GetChannelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *RrmConsideration) SetChannel(v int32) {
	o.Channel = v
}

// GetNoise returns the Noise field value
func (o *RrmConsideration) GetNoise() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Noise
}

// GetNoiseOk returns a tuple with the Noise field value
// and a boolean to check if the value has been set.
func (o *RrmConsideration) GetNoiseOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Noise, true
}

// SetNoise sets field value
func (o *RrmConsideration) SetNoise(v float32) {
	o.Noise = v
}

// GetOtherRssi returns the OtherRssi field value if set, zero value otherwise.
func (o *RrmConsideration) GetOtherRssi() float32 {
	if o == nil || IsNil(o.OtherRssi) {
		var ret float32
		return ret
	}
	return *o.OtherRssi
}

// GetOtherRssiOk returns a tuple with the OtherRssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmConsideration) GetOtherRssiOk() (*float32, bool) {
	if o == nil || IsNil(o.OtherRssi) {
		return nil, false
	}
	return o.OtherRssi, true
}

// HasOtherRssi returns a boolean if a field has been set.
func (o *RrmConsideration) HasOtherRssi() bool {
	if o != nil && !IsNil(o.OtherRssi) {
		return true
	}

	return false
}

// SetOtherRssi gets a reference to the given float32 and assigns it to the OtherRssi field.
func (o *RrmConsideration) SetOtherRssi(v float32) {
	o.OtherRssi = &v
}

// GetOtherSsid returns the OtherSsid field value if set, zero value otherwise.
func (o *RrmConsideration) GetOtherSsid() string {
	if o == nil || IsNil(o.OtherSsid) {
		var ret string
		return ret
	}
	return *o.OtherSsid
}

// GetOtherSsidOk returns a tuple with the OtherSsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmConsideration) GetOtherSsidOk() (*string, bool) {
	if o == nil || IsNil(o.OtherSsid) {
		return nil, false
	}
	return o.OtherSsid, true
}

// HasOtherSsid returns a boolean if a field has been set.
func (o *RrmConsideration) HasOtherSsid() bool {
	if o != nil && !IsNil(o.OtherSsid) {
		return true
	}

	return false
}

// SetOtherSsid gets a reference to the given string and assigns it to the OtherSsid field.
func (o *RrmConsideration) SetOtherSsid(v string) {
	o.OtherSsid = &v
}

// GetUtilScore returns the UtilScore field value
func (o *RrmConsideration) GetUtilScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UtilScore
}

// GetUtilScoreOk returns a tuple with the UtilScore field value
// and a boolean to check if the value has been set.
func (o *RrmConsideration) GetUtilScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UtilScore, true
}

// SetUtilScore sets field value
func (o *RrmConsideration) SetUtilScore(v float32) {
	o.UtilScore = v
}

// GetUtilScoreNonWifi returns the UtilScoreNonWifi field value
func (o *RrmConsideration) GetUtilScoreNonWifi() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UtilScoreNonWifi
}

// GetUtilScoreNonWifiOk returns a tuple with the UtilScoreNonWifi field value
// and a boolean to check if the value has been set.
func (o *RrmConsideration) GetUtilScoreNonWifiOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UtilScoreNonWifi, true
}

// SetUtilScoreNonWifi sets field value
func (o *RrmConsideration) SetUtilScoreNonWifi(v float32) {
	o.UtilScoreNonWifi = v
}

// GetUtilScoreOther returns the UtilScoreOther field value
func (o *RrmConsideration) GetUtilScoreOther() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UtilScoreOther
}

// GetUtilScoreOtherOk returns a tuple with the UtilScoreOther field value
// and a boolean to check if the value has been set.
func (o *RrmConsideration) GetUtilScoreOtherOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UtilScoreOther, true
}

// SetUtilScoreOther sets field value
func (o *RrmConsideration) SetUtilScoreOther(v float32) {
	o.UtilScoreOther = v
}

func (o RrmConsideration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RrmConsideration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["noise"] = o.Noise
	if !IsNil(o.OtherRssi) {
		toSerialize["other_rssi"] = o.OtherRssi
	}
	if !IsNil(o.OtherSsid) {
		toSerialize["other_ssid"] = o.OtherSsid
	}
	toSerialize["util_score"] = o.UtilScore
	toSerialize["util_score_non_wifi"] = o.UtilScoreNonWifi
	toSerialize["util_score_other"] = o.UtilScoreOther

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RrmConsideration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel",
		"noise",
		"util_score",
		"util_score_non_wifi",
		"util_score_other",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRrmConsideration := _RrmConsideration{}

	err = json.Unmarshal(data, &varRrmConsideration)

	if err != nil {
		return err
	}

	*o = RrmConsideration(varRrmConsideration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "channel")
		delete(additionalProperties, "noise")
		delete(additionalProperties, "other_rssi")
		delete(additionalProperties, "other_ssid")
		delete(additionalProperties, "util_score")
		delete(additionalProperties, "util_score_non_wifi")
		delete(additionalProperties, "util_score_other")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRrmConsideration struct {
	value *RrmConsideration
	isSet bool
}

func (v NullableRrmConsideration) Get() *RrmConsideration {
	return v.value
}

func (v *NullableRrmConsideration) Set(val *RrmConsideration) {
	v.value = val
	v.isSet = true
}

func (v NullableRrmConsideration) IsSet() bool {
	return v.isSet
}

func (v *NullableRrmConsideration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRrmConsideration(val *RrmConsideration) *NullableRrmConsideration {
	return &NullableRrmConsideration{value: val, isSet: true}
}

func (v NullableRrmConsideration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRrmConsideration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

