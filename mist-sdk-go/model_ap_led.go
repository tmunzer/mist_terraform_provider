/*
Mist API

> Version: **2405.1.6** > > Date: **June 6, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location-services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2405.1.6
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistsdkgo

import (
	"encoding/json"
)

// checks if the ApLed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApLed{}

// ApLed LED AP settings
type ApLed struct {
	Brightness *int32 `json:"brightness,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApLed ApLed

// NewApLed instantiates a new ApLed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApLed() *ApLed {
	this := ApLed{}
	var brightness int32 = 255
	this.Brightness = &brightness
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewApLedWithDefaults instantiates a new ApLed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApLedWithDefaults() *ApLed {
	this := ApLed{}
	var brightness int32 = 255
	this.Brightness = &brightness
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetBrightness returns the Brightness field value if set, zero value otherwise.
func (o *ApLed) GetBrightness() int32 {
	if o == nil || IsNil(o.Brightness) {
		var ret int32
		return ret
	}
	return *o.Brightness
}

// GetBrightnessOk returns a tuple with the Brightness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApLed) GetBrightnessOk() (*int32, bool) {
	if o == nil || IsNil(o.Brightness) {
		return nil, false
	}
	return o.Brightness, true
}

// HasBrightness returns a boolean if a field has been set.
func (o *ApLed) HasBrightness() bool {
	if o != nil && !IsNil(o.Brightness) {
		return true
	}

	return false
}

// SetBrightness gets a reference to the given int32 and assigns it to the Brightness field.
func (o *ApLed) SetBrightness(v int32) {
	o.Brightness = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApLed) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApLed) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApLed) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApLed) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ApLed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApLed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Brightness) {
		toSerialize["brightness"] = o.Brightness
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApLed) UnmarshalJSON(data []byte) (err error) {
	varApLed := _ApLed{}

	err = json.Unmarshal(data, &varApLed)

	if err != nil {
		return err
	}

	*o = ApLed(varApLed)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "brightness")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApLed struct {
	value *ApLed
	isSet bool
}

func (v NullableApLed) Get() *ApLed {
	return v.value
}

func (v *NullableApLed) Set(val *ApLed) {
	v.value = val
	v.isSet = true
}

func (v NullableApLed) IsSet() bool {
	return v.isSet
}

func (v *NullableApLed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApLed(val *ApLed) *NullableApLed {
	return &NullableApLed{value: val, isSet: true}
}

func (v NullableApLed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApLed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

