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

// checks if the RrmBand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RrmBand{}

// RrmBand struct for RrmBand
type RrmBand struct {
	Bandwidth *Dot11Bandwidth `json:"bandwidth,omitempty"`
	// proposed channel
	Channel *int32 `json:"channel,omitempty"`
	CurrBandwidht *Dot11Bandwidth `json:"curr_bandwidht,omitempty"`
	// current channel
	CurrChannel *int32 `json:"curr_channel,omitempty"`
	// current tx power
	CurrPower *int32 `json:"curr_power,omitempty"`
	// current radio band
	CurrUsage *string `json:"curr_usage,omitempty"`
	// proposed tx power
	Power *int32 `json:"power,omitempty"`
	// proposed radio band
	Usage *string `json:"usage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RrmBand RrmBand

// NewRrmBand instantiates a new RrmBand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRrmBand() *RrmBand {
	this := RrmBand{}
	return &this
}

// NewRrmBandWithDefaults instantiates a new RrmBand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRrmBandWithDefaults() *RrmBand {
	this := RrmBand{}
	return &this
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *RrmBand) GetBandwidth() Dot11Bandwidth {
	if o == nil || IsNil(o.Bandwidth) {
		var ret Dot11Bandwidth
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmBand) GetBandwidthOk() (*Dot11Bandwidth, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *RrmBand) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given Dot11Bandwidth and assigns it to the Bandwidth field.
func (o *RrmBand) SetBandwidth(v Dot11Bandwidth) {
	o.Bandwidth = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *RrmBand) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmBand) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *RrmBand) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *RrmBand) SetChannel(v int32) {
	o.Channel = &v
}

// GetCurrBandwidht returns the CurrBandwidht field value if set, zero value otherwise.
func (o *RrmBand) GetCurrBandwidht() Dot11Bandwidth {
	if o == nil || IsNil(o.CurrBandwidht) {
		var ret Dot11Bandwidth
		return ret
	}
	return *o.CurrBandwidht
}

// GetCurrBandwidhtOk returns a tuple with the CurrBandwidht field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmBand) GetCurrBandwidhtOk() (*Dot11Bandwidth, bool) {
	if o == nil || IsNil(o.CurrBandwidht) {
		return nil, false
	}
	return o.CurrBandwidht, true
}

// HasCurrBandwidht returns a boolean if a field has been set.
func (o *RrmBand) HasCurrBandwidht() bool {
	if o != nil && !IsNil(o.CurrBandwidht) {
		return true
	}

	return false
}

// SetCurrBandwidht gets a reference to the given Dot11Bandwidth and assigns it to the CurrBandwidht field.
func (o *RrmBand) SetCurrBandwidht(v Dot11Bandwidth) {
	o.CurrBandwidht = &v
}

// GetCurrChannel returns the CurrChannel field value if set, zero value otherwise.
func (o *RrmBand) GetCurrChannel() int32 {
	if o == nil || IsNil(o.CurrChannel) {
		var ret int32
		return ret
	}
	return *o.CurrChannel
}

// GetCurrChannelOk returns a tuple with the CurrChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmBand) GetCurrChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrChannel) {
		return nil, false
	}
	return o.CurrChannel, true
}

// HasCurrChannel returns a boolean if a field has been set.
func (o *RrmBand) HasCurrChannel() bool {
	if o != nil && !IsNil(o.CurrChannel) {
		return true
	}

	return false
}

// SetCurrChannel gets a reference to the given int32 and assigns it to the CurrChannel field.
func (o *RrmBand) SetCurrChannel(v int32) {
	o.CurrChannel = &v
}

// GetCurrPower returns the CurrPower field value if set, zero value otherwise.
func (o *RrmBand) GetCurrPower() int32 {
	if o == nil || IsNil(o.CurrPower) {
		var ret int32
		return ret
	}
	return *o.CurrPower
}

// GetCurrPowerOk returns a tuple with the CurrPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmBand) GetCurrPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrPower) {
		return nil, false
	}
	return o.CurrPower, true
}

// HasCurrPower returns a boolean if a field has been set.
func (o *RrmBand) HasCurrPower() bool {
	if o != nil && !IsNil(o.CurrPower) {
		return true
	}

	return false
}

// SetCurrPower gets a reference to the given int32 and assigns it to the CurrPower field.
func (o *RrmBand) SetCurrPower(v int32) {
	o.CurrPower = &v
}

// GetCurrUsage returns the CurrUsage field value if set, zero value otherwise.
func (o *RrmBand) GetCurrUsage() string {
	if o == nil || IsNil(o.CurrUsage) {
		var ret string
		return ret
	}
	return *o.CurrUsage
}

// GetCurrUsageOk returns a tuple with the CurrUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmBand) GetCurrUsageOk() (*string, bool) {
	if o == nil || IsNil(o.CurrUsage) {
		return nil, false
	}
	return o.CurrUsage, true
}

// HasCurrUsage returns a boolean if a field has been set.
func (o *RrmBand) HasCurrUsage() bool {
	if o != nil && !IsNil(o.CurrUsage) {
		return true
	}

	return false
}

// SetCurrUsage gets a reference to the given string and assigns it to the CurrUsage field.
func (o *RrmBand) SetCurrUsage(v string) {
	o.CurrUsage = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *RrmBand) GetPower() int32 {
	if o == nil || IsNil(o.Power) {
		var ret int32
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmBand) GetPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.Power) {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *RrmBand) HasPower() bool {
	if o != nil && !IsNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given int32 and assigns it to the Power field.
func (o *RrmBand) SetPower(v int32) {
	o.Power = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *RrmBand) GetUsage() string {
	if o == nil || IsNil(o.Usage) {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmBand) GetUsageOk() (*string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *RrmBand) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *RrmBand) SetUsage(v string) {
	o.Usage = &v
}

func (o RrmBand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RrmBand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.CurrBandwidht) {
		toSerialize["curr_bandwidht"] = o.CurrBandwidht
	}
	if !IsNil(o.CurrChannel) {
		toSerialize["curr_channel"] = o.CurrChannel
	}
	if !IsNil(o.CurrPower) {
		toSerialize["curr_power"] = o.CurrPower
	}
	if !IsNil(o.CurrUsage) {
		toSerialize["curr_usage"] = o.CurrUsage
	}
	if !IsNil(o.Power) {
		toSerialize["power"] = o.Power
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RrmBand) UnmarshalJSON(data []byte) (err error) {
	varRrmBand := _RrmBand{}

	err = json.Unmarshal(data, &varRrmBand)

	if err != nil {
		return err
	}

	*o = RrmBand(varRrmBand)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "channel")
		delete(additionalProperties, "curr_bandwidht")
		delete(additionalProperties, "curr_channel")
		delete(additionalProperties, "curr_power")
		delete(additionalProperties, "curr_usage")
		delete(additionalProperties, "power")
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRrmBand struct {
	value *RrmBand
	isSet bool
}

func (v NullableRrmBand) Get() *RrmBand {
	return v.value
}

func (v *NullableRrmBand) Set(val *RrmBand) {
	v.value = val
	v.isSet = true
}

func (v NullableRrmBand) IsSet() bool {
	return v.isSet
}

func (v *NullableRrmBand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRrmBand(val *RrmBand) *NullableRrmBand {
	return &NullableRrmBand{value: val, isSet: true}
}

func (v NullableRrmBand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRrmBand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

