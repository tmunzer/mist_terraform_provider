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

// checks if the LastTrouble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastTrouble{}

// LastTrouble last trouble code of switch
type LastTrouble struct {
	// Code definitions list at /api/v1/consts/ap_led_status
	Code *string `json:"code,omitempty"`
	Timestamp *int32 `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LastTrouble LastTrouble

// NewLastTrouble instantiates a new LastTrouble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastTrouble() *LastTrouble {
	this := LastTrouble{}
	return &this
}

// NewLastTroubleWithDefaults instantiates a new LastTrouble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastTroubleWithDefaults() *LastTrouble {
	this := LastTrouble{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *LastTrouble) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastTrouble) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *LastTrouble) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *LastTrouble) SetCode(v string) {
	o.Code = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LastTrouble) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastTrouble) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LastTrouble) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *LastTrouble) SetTimestamp(v int32) {
	o.Timestamp = &v
}

func (o LastTrouble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastTrouble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LastTrouble) UnmarshalJSON(data []byte) (err error) {
	varLastTrouble := _LastTrouble{}

	err = json.Unmarshal(data, &varLastTrouble)

	if err != nil {
		return err
	}

	*o = LastTrouble(varLastTrouble)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLastTrouble struct {
	value *LastTrouble
	isSet bool
}

func (v NullableLastTrouble) Get() *LastTrouble {
	return v.value
}

func (v *NullableLastTrouble) Set(val *LastTrouble) {
	v.value = val
	v.isSet = true
}

func (v NullableLastTrouble) IsSet() bool {
	return v.isSet
}

func (v *NullableLastTrouble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastTrouble(val *LastTrouble) *NullableLastTrouble {
	return &NullableLastTrouble{value: val, isSet: true}
}

func (v NullableLastTrouble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastTrouble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

