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

// checks if the GatewayPortConfigTrafficShaping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayPortConfigTrafficShaping{}

// GatewayPortConfigTrafficShaping struct for GatewayPortConfigTrafficShaping
type GatewayPortConfigTrafficShaping struct {
	// percentages for differet class of traffic: high / medium / low / best-effort sum must be equal to 100
	ClassPercentages []int32 `json:"class_percentages,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayPortConfigTrafficShaping GatewayPortConfigTrafficShaping

// NewGatewayPortConfigTrafficShaping instantiates a new GatewayPortConfigTrafficShaping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayPortConfigTrafficShaping() *GatewayPortConfigTrafficShaping {
	this := GatewayPortConfigTrafficShaping{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewGatewayPortConfigTrafficShapingWithDefaults instantiates a new GatewayPortConfigTrafficShaping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayPortConfigTrafficShapingWithDefaults() *GatewayPortConfigTrafficShaping {
	this := GatewayPortConfigTrafficShaping{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetClassPercentages returns the ClassPercentages field value if set, zero value otherwise.
func (o *GatewayPortConfigTrafficShaping) GetClassPercentages() []int32 {
	if o == nil || IsNil(o.ClassPercentages) {
		var ret []int32
		return ret
	}
	return o.ClassPercentages
}

// GetClassPercentagesOk returns a tuple with the ClassPercentages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortConfigTrafficShaping) GetClassPercentagesOk() ([]int32, bool) {
	if o == nil || IsNil(o.ClassPercentages) {
		return nil, false
	}
	return o.ClassPercentages, true
}

// HasClassPercentages returns a boolean if a field has been set.
func (o *GatewayPortConfigTrafficShaping) HasClassPercentages() bool {
	if o != nil && !IsNil(o.ClassPercentages) {
		return true
	}

	return false
}

// SetClassPercentages gets a reference to the given []int32 and assigns it to the ClassPercentages field.
func (o *GatewayPortConfigTrafficShaping) SetClassPercentages(v []int32) {
	o.ClassPercentages = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GatewayPortConfigTrafficShaping) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortConfigTrafficShaping) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GatewayPortConfigTrafficShaping) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GatewayPortConfigTrafficShaping) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GatewayPortConfigTrafficShaping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayPortConfigTrafficShaping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClassPercentages) {
		toSerialize["class_percentages"] = o.ClassPercentages
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayPortConfigTrafficShaping) UnmarshalJSON(data []byte) (err error) {
	varGatewayPortConfigTrafficShaping := _GatewayPortConfigTrafficShaping{}

	err = json.Unmarshal(data, &varGatewayPortConfigTrafficShaping)

	if err != nil {
		return err
	}

	*o = GatewayPortConfigTrafficShaping(varGatewayPortConfigTrafficShaping)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "class_percentages")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayPortConfigTrafficShaping struct {
	value *GatewayPortConfigTrafficShaping
	isSet bool
}

func (v NullableGatewayPortConfigTrafficShaping) Get() *GatewayPortConfigTrafficShaping {
	return v.value
}

func (v *NullableGatewayPortConfigTrafficShaping) Set(val *GatewayPortConfigTrafficShaping) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPortConfigTrafficShaping) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPortConfigTrafficShaping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPortConfigTrafficShaping(val *GatewayPortConfigTrafficShaping) *NullableGatewayPortConfigTrafficShaping {
	return &NullableGatewayPortConfigTrafficShaping{value: val, isSet: true}
}

func (v NullableGatewayPortConfigTrafficShaping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPortConfigTrafficShaping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

