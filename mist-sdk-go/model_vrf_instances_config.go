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

// checks if the VrfInstancesConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfInstancesConfig{}

// VrfInstancesConfig struct for VrfInstancesConfig
type VrfInstancesConfig struct {
	// Property key is the destination CIDR (e.g. \"10.0.0.0/8\")
	ExtraRoutes *map[string]VrfExtraRoutesValue `json:"extra_routes,omitempty"`
	Networks []string `json:"networks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrfInstancesConfig VrfInstancesConfig

// NewVrfInstancesConfig instantiates a new VrfInstancesConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfInstancesConfig() *VrfInstancesConfig {
	this := VrfInstancesConfig{}
	return &this
}

// NewVrfInstancesConfigWithDefaults instantiates a new VrfInstancesConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfInstancesConfigWithDefaults() *VrfInstancesConfig {
	this := VrfInstancesConfig{}
	return &this
}

// GetExtraRoutes returns the ExtraRoutes field value if set, zero value otherwise.
func (o *VrfInstancesConfig) GetExtraRoutes() map[string]VrfExtraRoutesValue {
	if o == nil || IsNil(o.ExtraRoutes) {
		var ret map[string]VrfExtraRoutesValue
		return ret
	}
	return *o.ExtraRoutes
}

// GetExtraRoutesOk returns a tuple with the ExtraRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfInstancesConfig) GetExtraRoutesOk() (*map[string]VrfExtraRoutesValue, bool) {
	if o == nil || IsNil(o.ExtraRoutes) {
		return nil, false
	}
	return o.ExtraRoutes, true
}

// HasExtraRoutes returns a boolean if a field has been set.
func (o *VrfInstancesConfig) HasExtraRoutes() bool {
	if o != nil && !IsNil(o.ExtraRoutes) {
		return true
	}

	return false
}

// SetExtraRoutes gets a reference to the given map[string]VrfExtraRoutesValue and assigns it to the ExtraRoutes field.
func (o *VrfInstancesConfig) SetExtraRoutes(v map[string]VrfExtraRoutesValue) {
	o.ExtraRoutes = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *VrfInstancesConfig) GetNetworks() []string {
	if o == nil || IsNil(o.Networks) {
		var ret []string
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfInstancesConfig) GetNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *VrfInstancesConfig) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *VrfInstancesConfig) SetNetworks(v []string) {
	o.Networks = v
}

func (o VrfInstancesConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfInstancesConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtraRoutes) {
		toSerialize["extra_routes"] = o.ExtraRoutes
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfInstancesConfig) UnmarshalJSON(data []byte) (err error) {
	varVrfInstancesConfig := _VrfInstancesConfig{}

	err = json.Unmarshal(data, &varVrfInstancesConfig)

	if err != nil {
		return err
	}

	*o = VrfInstancesConfig(varVrfInstancesConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "extra_routes")
		delete(additionalProperties, "networks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfInstancesConfig struct {
	value *VrfInstancesConfig
	isSet bool
}

func (v NullableVrfInstancesConfig) Get() *VrfInstancesConfig {
	return v.value
}

func (v *NullableVrfInstancesConfig) Set(val *VrfInstancesConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfInstancesConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfInstancesConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfInstancesConfig(val *VrfInstancesConfig) *NullableVrfInstancesConfig {
	return &NullableVrfInstancesConfig{value: val, isSet: true}
}

func (v NullableVrfInstancesConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfInstancesConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

