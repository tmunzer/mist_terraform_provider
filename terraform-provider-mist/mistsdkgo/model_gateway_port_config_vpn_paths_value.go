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

// checks if the GatewayPortConfigVpnPathsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayPortConfigVpnPathsValue{}

// GatewayPortConfigVpnPathsValue struct for GatewayPortConfigVpnPathsValue
type GatewayPortConfigVpnPathsValue struct {
	BfdProfile *string `json:"bfd_profile,omitempty"`
	// whether to use tunnel mode. SSR only
	BfdUseTunnelMode *bool `json:"bfd_use_tunnel_mode,omitempty"`
	Role *string `json:"role,omitempty"`
	TrafficShaping *GatewayPortConfigVpnPathsValueTrafficShaping `json:"traffic_shaping,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayPortConfigVpnPathsValue GatewayPortConfigVpnPathsValue

// NewGatewayPortConfigVpnPathsValue instantiates a new GatewayPortConfigVpnPathsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayPortConfigVpnPathsValue() *GatewayPortConfigVpnPathsValue {
	this := GatewayPortConfigVpnPathsValue{}
	var bfdProfile string = "broadband"
	this.BfdProfile = &bfdProfile
	var bfdUseTunnelMode bool = false
	this.BfdUseTunnelMode = &bfdUseTunnelMode
	var role string = "spoke"
	this.Role = &role
	return &this
}

// NewGatewayPortConfigVpnPathsValueWithDefaults instantiates a new GatewayPortConfigVpnPathsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayPortConfigVpnPathsValueWithDefaults() *GatewayPortConfigVpnPathsValue {
	this := GatewayPortConfigVpnPathsValue{}
	var bfdProfile string = "broadband"
	this.BfdProfile = &bfdProfile
	var bfdUseTunnelMode bool = false
	this.BfdUseTunnelMode = &bfdUseTunnelMode
	var role string = "spoke"
	this.Role = &role
	return &this
}

// GetBfdProfile returns the BfdProfile field value if set, zero value otherwise.
func (o *GatewayPortConfigVpnPathsValue) GetBfdProfile() string {
	if o == nil || IsNil(o.BfdProfile) {
		var ret string
		return ret
	}
	return *o.BfdProfile
}

// GetBfdProfileOk returns a tuple with the BfdProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortConfigVpnPathsValue) GetBfdProfileOk() (*string, bool) {
	if o == nil || IsNil(o.BfdProfile) {
		return nil, false
	}
	return o.BfdProfile, true
}

// HasBfdProfile returns a boolean if a field has been set.
func (o *GatewayPortConfigVpnPathsValue) HasBfdProfile() bool {
	if o != nil && !IsNil(o.BfdProfile) {
		return true
	}

	return false
}

// SetBfdProfile gets a reference to the given string and assigns it to the BfdProfile field.
func (o *GatewayPortConfigVpnPathsValue) SetBfdProfile(v string) {
	o.BfdProfile = &v
}

// GetBfdUseTunnelMode returns the BfdUseTunnelMode field value if set, zero value otherwise.
func (o *GatewayPortConfigVpnPathsValue) GetBfdUseTunnelMode() bool {
	if o == nil || IsNil(o.BfdUseTunnelMode) {
		var ret bool
		return ret
	}
	return *o.BfdUseTunnelMode
}

// GetBfdUseTunnelModeOk returns a tuple with the BfdUseTunnelMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortConfigVpnPathsValue) GetBfdUseTunnelModeOk() (*bool, bool) {
	if o == nil || IsNil(o.BfdUseTunnelMode) {
		return nil, false
	}
	return o.BfdUseTunnelMode, true
}

// HasBfdUseTunnelMode returns a boolean if a field has been set.
func (o *GatewayPortConfigVpnPathsValue) HasBfdUseTunnelMode() bool {
	if o != nil && !IsNil(o.BfdUseTunnelMode) {
		return true
	}

	return false
}

// SetBfdUseTunnelMode gets a reference to the given bool and assigns it to the BfdUseTunnelMode field.
func (o *GatewayPortConfigVpnPathsValue) SetBfdUseTunnelMode(v bool) {
	o.BfdUseTunnelMode = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *GatewayPortConfigVpnPathsValue) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortConfigVpnPathsValue) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *GatewayPortConfigVpnPathsValue) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *GatewayPortConfigVpnPathsValue) SetRole(v string) {
	o.Role = &v
}

// GetTrafficShaping returns the TrafficShaping field value if set, zero value otherwise.
func (o *GatewayPortConfigVpnPathsValue) GetTrafficShaping() GatewayPortConfigVpnPathsValueTrafficShaping {
	if o == nil || IsNil(o.TrafficShaping) {
		var ret GatewayPortConfigVpnPathsValueTrafficShaping
		return ret
	}
	return *o.TrafficShaping
}

// GetTrafficShapingOk returns a tuple with the TrafficShaping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortConfigVpnPathsValue) GetTrafficShapingOk() (*GatewayPortConfigVpnPathsValueTrafficShaping, bool) {
	if o == nil || IsNil(o.TrafficShaping) {
		return nil, false
	}
	return o.TrafficShaping, true
}

// HasTrafficShaping returns a boolean if a field has been set.
func (o *GatewayPortConfigVpnPathsValue) HasTrafficShaping() bool {
	if o != nil && !IsNil(o.TrafficShaping) {
		return true
	}

	return false
}

// SetTrafficShaping gets a reference to the given GatewayPortConfigVpnPathsValueTrafficShaping and assigns it to the TrafficShaping field.
func (o *GatewayPortConfigVpnPathsValue) SetTrafficShaping(v GatewayPortConfigVpnPathsValueTrafficShaping) {
	o.TrafficShaping = &v
}

func (o GatewayPortConfigVpnPathsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayPortConfigVpnPathsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BfdProfile) {
		toSerialize["bfd_profile"] = o.BfdProfile
	}
	if !IsNil(o.BfdUseTunnelMode) {
		toSerialize["bfd_use_tunnel_mode"] = o.BfdUseTunnelMode
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.TrafficShaping) {
		toSerialize["traffic_shaping"] = o.TrafficShaping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayPortConfigVpnPathsValue) UnmarshalJSON(data []byte) (err error) {
	varGatewayPortConfigVpnPathsValue := _GatewayPortConfigVpnPathsValue{}

	err = json.Unmarshal(data, &varGatewayPortConfigVpnPathsValue)

	if err != nil {
		return err
	}

	*o = GatewayPortConfigVpnPathsValue(varGatewayPortConfigVpnPathsValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bfd_profile")
		delete(additionalProperties, "bfd_use_tunnel_mode")
		delete(additionalProperties, "role")
		delete(additionalProperties, "traffic_shaping")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayPortConfigVpnPathsValue struct {
	value *GatewayPortConfigVpnPathsValue
	isSet bool
}

func (v NullableGatewayPortConfigVpnPathsValue) Get() *GatewayPortConfigVpnPathsValue {
	return v.value
}

func (v *NullableGatewayPortConfigVpnPathsValue) Set(val *GatewayPortConfigVpnPathsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPortConfigVpnPathsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPortConfigVpnPathsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPortConfigVpnPathsValue(val *GatewayPortConfigVpnPathsValue) *NullableGatewayPortConfigVpnPathsValue {
	return &NullableGatewayPortConfigVpnPathsValue{value: val, isSet: true}
}

func (v NullableGatewayPortConfigVpnPathsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPortConfigVpnPathsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

