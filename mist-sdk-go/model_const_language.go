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

// checks if the ConstLanguage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstLanguage{}

// ConstLanguage struct for ConstLanguage
type ConstLanguage struct {
	Display string `json:"display"`
	DisplayNative string `json:"display_native"`
	Key string `json:"key"`
	AdditionalProperties map[string]interface{}
}

type _ConstLanguage ConstLanguage

// NewConstLanguage instantiates a new ConstLanguage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstLanguage(display string, displayNative string, key string) *ConstLanguage {
	this := ConstLanguage{}
	this.Display = display
	this.DisplayNative = displayNative
	this.Key = key
	return &this
}

// NewConstLanguageWithDefaults instantiates a new ConstLanguage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstLanguageWithDefaults() *ConstLanguage {
	this := ConstLanguage{}
	return &this
}

// GetDisplay returns the Display field value
func (o *ConstLanguage) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConstLanguage) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConstLanguage) SetDisplay(v string) {
	o.Display = v
}

// GetDisplayNative returns the DisplayNative field value
func (o *ConstLanguage) GetDisplayNative() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayNative
}

// GetDisplayNativeOk returns a tuple with the DisplayNative field value
// and a boolean to check if the value has been set.
func (o *ConstLanguage) GetDisplayNativeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayNative, true
}

// SetDisplayNative sets field value
func (o *ConstLanguage) SetDisplayNative(v string) {
	o.DisplayNative = v
}

// GetKey returns the Key field value
func (o *ConstLanguage) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ConstLanguage) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ConstLanguage) SetKey(v string) {
	o.Key = v
}

func (o ConstLanguage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstLanguage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display"] = o.Display
	toSerialize["display_native"] = o.DisplayNative
	toSerialize["key"] = o.Key

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstLanguage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"display",
		"display_native",
		"key",
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

	varConstLanguage := _ConstLanguage{}

	err = json.Unmarshal(data, &varConstLanguage)

	if err != nil {
		return err
	}

	*o = ConstLanguage(varConstLanguage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display")
		delete(additionalProperties, "display_native")
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstLanguage struct {
	value *ConstLanguage
	isSet bool
}

func (v NullableConstLanguage) Get() *ConstLanguage {
	return v.value
}

func (v *NullableConstLanguage) Set(val *ConstLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableConstLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableConstLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstLanguage(val *ConstLanguage) *NullableConstLanguage {
	return &NullableConstLanguage{value: val, isSet: true}
}

func (v NullableConstLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

