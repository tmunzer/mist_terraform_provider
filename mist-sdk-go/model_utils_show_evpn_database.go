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

// checks if the UtilsShowEvpnDatabase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsShowEvpnDatabase{}

// UtilsShowEvpnDatabase struct for UtilsShowEvpnDatabase
type UtilsShowEvpnDatabase struct {
	// client mac filter
	Mac *string `json:"mac,omitempty"`
	// interface name
	PortId *string `json:"port_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UtilsShowEvpnDatabase UtilsShowEvpnDatabase

// NewUtilsShowEvpnDatabase instantiates a new UtilsShowEvpnDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsShowEvpnDatabase() *UtilsShowEvpnDatabase {
	this := UtilsShowEvpnDatabase{}
	return &this
}

// NewUtilsShowEvpnDatabaseWithDefaults instantiates a new UtilsShowEvpnDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsShowEvpnDatabaseWithDefaults() *UtilsShowEvpnDatabase {
	this := UtilsShowEvpnDatabase{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *UtilsShowEvpnDatabase) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowEvpnDatabase) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *UtilsShowEvpnDatabase) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *UtilsShowEvpnDatabase) SetMac(v string) {
	o.Mac = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *UtilsShowEvpnDatabase) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowEvpnDatabase) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *UtilsShowEvpnDatabase) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *UtilsShowEvpnDatabase) SetPortId(v string) {
	o.PortId = &v
}

func (o UtilsShowEvpnDatabase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsShowEvpnDatabase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.PortId) {
		toSerialize["port_id"] = o.PortId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UtilsShowEvpnDatabase) UnmarshalJSON(data []byte) (err error) {
	varUtilsShowEvpnDatabase := _UtilsShowEvpnDatabase{}

	err = json.Unmarshal(data, &varUtilsShowEvpnDatabase)

	if err != nil {
		return err
	}

	*o = UtilsShowEvpnDatabase(varUtilsShowEvpnDatabase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mac")
		delete(additionalProperties, "port_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUtilsShowEvpnDatabase struct {
	value *UtilsShowEvpnDatabase
	isSet bool
}

func (v NullableUtilsShowEvpnDatabase) Get() *UtilsShowEvpnDatabase {
	return v.value
}

func (v *NullableUtilsShowEvpnDatabase) Set(val *UtilsShowEvpnDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsShowEvpnDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsShowEvpnDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsShowEvpnDatabase(val *UtilsShowEvpnDatabase) *NullableUtilsShowEvpnDatabase {
	return &NullableUtilsShowEvpnDatabase{value: val, isSet: true}
}

func (v NullableUtilsShowEvpnDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsShowEvpnDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

