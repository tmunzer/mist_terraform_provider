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

// checks if the SwitchExtraRoutesValueNextQualifiedValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchExtraRoutesValueNextQualifiedValue{}

// SwitchExtraRoutesValueNextQualifiedValue struct for SwitchExtraRoutesValueNextQualifiedValue
type SwitchExtraRoutesValueNextQualifiedValue struct {
	Metric NullableInt32 `json:"metric,omitempty"`
	Preference NullableInt32 `json:"preference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwitchExtraRoutesValueNextQualifiedValue SwitchExtraRoutesValueNextQualifiedValue

// NewSwitchExtraRoutesValueNextQualifiedValue instantiates a new SwitchExtraRoutesValueNextQualifiedValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchExtraRoutesValueNextQualifiedValue() *SwitchExtraRoutesValueNextQualifiedValue {
	this := SwitchExtraRoutesValueNextQualifiedValue{}
	return &this
}

// NewSwitchExtraRoutesValueNextQualifiedValueWithDefaults instantiates a new SwitchExtraRoutesValueNextQualifiedValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchExtraRoutesValueNextQualifiedValueWithDefaults() *SwitchExtraRoutesValueNextQualifiedValue {
	this := SwitchExtraRoutesValueNextQualifiedValue{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SwitchExtraRoutesValueNextQualifiedValue) GetMetric() int32 {
	if o == nil || IsNil(o.Metric.Get()) {
		var ret int32
		return ret
	}
	return *o.Metric.Get()
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SwitchExtraRoutesValueNextQualifiedValue) GetMetricOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metric.Get(), o.Metric.IsSet()
}

// HasMetric returns a boolean if a field has been set.
func (o *SwitchExtraRoutesValueNextQualifiedValue) HasMetric() bool {
	if o != nil && o.Metric.IsSet() {
		return true
	}

	return false
}

// SetMetric gets a reference to the given NullableInt32 and assigns it to the Metric field.
func (o *SwitchExtraRoutesValueNextQualifiedValue) SetMetric(v int32) {
	o.Metric.Set(&v)
}
// SetMetricNil sets the value for Metric to be an explicit nil
func (o *SwitchExtraRoutesValueNextQualifiedValue) SetMetricNil() {
	o.Metric.Set(nil)
}

// UnsetMetric ensures that no value is present for Metric, not even an explicit nil
func (o *SwitchExtraRoutesValueNextQualifiedValue) UnsetMetric() {
	o.Metric.Unset()
}

// GetPreference returns the Preference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SwitchExtraRoutesValueNextQualifiedValue) GetPreference() int32 {
	if o == nil || IsNil(o.Preference.Get()) {
		var ret int32
		return ret
	}
	return *o.Preference.Get()
}

// GetPreferenceOk returns a tuple with the Preference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SwitchExtraRoutesValueNextQualifiedValue) GetPreferenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preference.Get(), o.Preference.IsSet()
}

// HasPreference returns a boolean if a field has been set.
func (o *SwitchExtraRoutesValueNextQualifiedValue) HasPreference() bool {
	if o != nil && o.Preference.IsSet() {
		return true
	}

	return false
}

// SetPreference gets a reference to the given NullableInt32 and assigns it to the Preference field.
func (o *SwitchExtraRoutesValueNextQualifiedValue) SetPreference(v int32) {
	o.Preference.Set(&v)
}
// SetPreferenceNil sets the value for Preference to be an explicit nil
func (o *SwitchExtraRoutesValueNextQualifiedValue) SetPreferenceNil() {
	o.Preference.Set(nil)
}

// UnsetPreference ensures that no value is present for Preference, not even an explicit nil
func (o *SwitchExtraRoutesValueNextQualifiedValue) UnsetPreference() {
	o.Preference.Unset()
}

func (o SwitchExtraRoutesValueNextQualifiedValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchExtraRoutesValueNextQualifiedValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Metric.IsSet() {
		toSerialize["metric"] = o.Metric.Get()
	}
	if o.Preference.IsSet() {
		toSerialize["preference"] = o.Preference.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwitchExtraRoutesValueNextQualifiedValue) UnmarshalJSON(data []byte) (err error) {
	varSwitchExtraRoutesValueNextQualifiedValue := _SwitchExtraRoutesValueNextQualifiedValue{}

	err = json.Unmarshal(data, &varSwitchExtraRoutesValueNextQualifiedValue)

	if err != nil {
		return err
	}

	*o = SwitchExtraRoutesValueNextQualifiedValue(varSwitchExtraRoutesValueNextQualifiedValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "metric")
		delete(additionalProperties, "preference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwitchExtraRoutesValueNextQualifiedValue struct {
	value *SwitchExtraRoutesValueNextQualifiedValue
	isSet bool
}

func (v NullableSwitchExtraRoutesValueNextQualifiedValue) Get() *SwitchExtraRoutesValueNextQualifiedValue {
	return v.value
}

func (v *NullableSwitchExtraRoutesValueNextQualifiedValue) Set(val *SwitchExtraRoutesValueNextQualifiedValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchExtraRoutesValueNextQualifiedValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchExtraRoutesValueNextQualifiedValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchExtraRoutesValueNextQualifiedValue(val *SwitchExtraRoutesValueNextQualifiedValue) *NullableSwitchExtraRoutesValueNextQualifiedValue {
	return &NullableSwitchExtraRoutesValueNextQualifiedValue{value: val, isSet: true}
}

func (v NullableSwitchExtraRoutesValueNextQualifiedValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchExtraRoutesValueNextQualifiedValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

