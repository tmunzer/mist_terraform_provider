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

// checks if the RfTemplateModelSpecificValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RfTemplateModelSpecificValue{}

// RfTemplateModelSpecificValue struct for RfTemplateModelSpecificValue
type RfTemplateModelSpecificValue struct {
	Band24 *ApRadioBand `json:"band_24,omitempty"`
	Band5 *ApRadioBand `json:"band_5,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RfTemplateModelSpecificValue RfTemplateModelSpecificValue

// NewRfTemplateModelSpecificValue instantiates a new RfTemplateModelSpecificValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRfTemplateModelSpecificValue() *RfTemplateModelSpecificValue {
	this := RfTemplateModelSpecificValue{}
	return &this
}

// NewRfTemplateModelSpecificValueWithDefaults instantiates a new RfTemplateModelSpecificValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRfTemplateModelSpecificValueWithDefaults() *RfTemplateModelSpecificValue {
	this := RfTemplateModelSpecificValue{}
	return &this
}

// GetBand24 returns the Band24 field value if set, zero value otherwise.
func (o *RfTemplateModelSpecificValue) GetBand24() ApRadioBand {
	if o == nil || IsNil(o.Band24) {
		var ret ApRadioBand
		return ret
	}
	return *o.Band24
}

// GetBand24Ok returns a tuple with the Band24 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplateModelSpecificValue) GetBand24Ok() (*ApRadioBand, bool) {
	if o == nil || IsNil(o.Band24) {
		return nil, false
	}
	return o.Band24, true
}

// HasBand24 returns a boolean if a field has been set.
func (o *RfTemplateModelSpecificValue) HasBand24() bool {
	if o != nil && !IsNil(o.Band24) {
		return true
	}

	return false
}

// SetBand24 gets a reference to the given ApRadioBand and assigns it to the Band24 field.
func (o *RfTemplateModelSpecificValue) SetBand24(v ApRadioBand) {
	o.Band24 = &v
}

// GetBand5 returns the Band5 field value if set, zero value otherwise.
func (o *RfTemplateModelSpecificValue) GetBand5() ApRadioBand {
	if o == nil || IsNil(o.Band5) {
		var ret ApRadioBand
		return ret
	}
	return *o.Band5
}

// GetBand5Ok returns a tuple with the Band5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplateModelSpecificValue) GetBand5Ok() (*ApRadioBand, bool) {
	if o == nil || IsNil(o.Band5) {
		return nil, false
	}
	return o.Band5, true
}

// HasBand5 returns a boolean if a field has been set.
func (o *RfTemplateModelSpecificValue) HasBand5() bool {
	if o != nil && !IsNil(o.Band5) {
		return true
	}

	return false
}

// SetBand5 gets a reference to the given ApRadioBand and assigns it to the Band5 field.
func (o *RfTemplateModelSpecificValue) SetBand5(v ApRadioBand) {
	o.Band5 = &v
}

func (o RfTemplateModelSpecificValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RfTemplateModelSpecificValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Band24) {
		toSerialize["band_24"] = o.Band24
	}
	if !IsNil(o.Band5) {
		toSerialize["band_5"] = o.Band5
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RfTemplateModelSpecificValue) UnmarshalJSON(data []byte) (err error) {
	varRfTemplateModelSpecificValue := _RfTemplateModelSpecificValue{}

	err = json.Unmarshal(data, &varRfTemplateModelSpecificValue)

	if err != nil {
		return err
	}

	*o = RfTemplateModelSpecificValue(varRfTemplateModelSpecificValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "band_24")
		delete(additionalProperties, "band_5")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRfTemplateModelSpecificValue struct {
	value *RfTemplateModelSpecificValue
	isSet bool
}

func (v NullableRfTemplateModelSpecificValue) Get() *RfTemplateModelSpecificValue {
	return v.value
}

func (v *NullableRfTemplateModelSpecificValue) Set(val *RfTemplateModelSpecificValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRfTemplateModelSpecificValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRfTemplateModelSpecificValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRfTemplateModelSpecificValue(val *RfTemplateModelSpecificValue) *NullableRfTemplateModelSpecificValue {
	return &NullableRfTemplateModelSpecificValue{value: val, isSet: true}
}

func (v NullableRfTemplateModelSpecificValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRfTemplateModelSpecificValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

