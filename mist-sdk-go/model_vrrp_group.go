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

// checks if the VrrpGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrrpGroup{}

// VrrpGroup Junos VRRP group
type VrrpGroup struct {
	// if `auth_type`==`md5`
	AuthKey *string `json:"auth_key,omitempty"`
	// if `auth_type`==`simple`
	AuthPassword *string `json:"auth_password,omitempty"`
	AuthType *VrrpGroupAuthType `json:"auth_type,omitempty"`
	// Property key is the network name
	Networks *map[string]VrrpGroupNetwork `json:"networks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrrpGroup VrrpGroup

// NewVrrpGroup instantiates a new VrrpGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrrpGroup() *VrrpGroup {
	this := VrrpGroup{}
	var authType VrrpGroupAuthType = VRRPGROUPAUTHTYPE_MD5
	this.AuthType = &authType
	return &this
}

// NewVrrpGroupWithDefaults instantiates a new VrrpGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrrpGroupWithDefaults() *VrrpGroup {
	this := VrrpGroup{}
	var authType VrrpGroupAuthType = VRRPGROUPAUTHTYPE_MD5
	this.AuthType = &authType
	return &this
}

// GetAuthKey returns the AuthKey field value if set, zero value otherwise.
func (o *VrrpGroup) GetAuthKey() string {
	if o == nil || IsNil(o.AuthKey) {
		var ret string
		return ret
	}
	return *o.AuthKey
}

// GetAuthKeyOk returns a tuple with the AuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrrpGroup) GetAuthKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AuthKey) {
		return nil, false
	}
	return o.AuthKey, true
}

// HasAuthKey returns a boolean if a field has been set.
func (o *VrrpGroup) HasAuthKey() bool {
	if o != nil && !IsNil(o.AuthKey) {
		return true
	}

	return false
}

// SetAuthKey gets a reference to the given string and assigns it to the AuthKey field.
func (o *VrrpGroup) SetAuthKey(v string) {
	o.AuthKey = &v
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *VrrpGroup) GetAuthPassword() string {
	if o == nil || IsNil(o.AuthPassword) {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrrpGroup) GetAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPassword) {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *VrrpGroup) HasAuthPassword() bool {
	if o != nil && !IsNil(o.AuthPassword) {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *VrrpGroup) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *VrrpGroup) GetAuthType() VrrpGroupAuthType {
	if o == nil || IsNil(o.AuthType) {
		var ret VrrpGroupAuthType
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrrpGroup) GetAuthTypeOk() (*VrrpGroupAuthType, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *VrrpGroup) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given VrrpGroupAuthType and assigns it to the AuthType field.
func (o *VrrpGroup) SetAuthType(v VrrpGroupAuthType) {
	o.AuthType = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *VrrpGroup) GetNetworks() map[string]VrrpGroupNetwork {
	if o == nil || IsNil(o.Networks) {
		var ret map[string]VrrpGroupNetwork
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrrpGroup) GetNetworksOk() (*map[string]VrrpGroupNetwork, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *VrrpGroup) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given map[string]VrrpGroupNetwork and assigns it to the Networks field.
func (o *VrrpGroup) SetNetworks(v map[string]VrrpGroupNetwork) {
	o.Networks = &v
}

func (o VrrpGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrrpGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthKey) {
		toSerialize["auth_key"] = o.AuthKey
	}
	if !IsNil(o.AuthPassword) {
		toSerialize["auth_password"] = o.AuthPassword
	}
	if !IsNil(o.AuthType) {
		toSerialize["auth_type"] = o.AuthType
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrrpGroup) UnmarshalJSON(data []byte) (err error) {
	varVrrpGroup := _VrrpGroup{}

	err = json.Unmarshal(data, &varVrrpGroup)

	if err != nil {
		return err
	}

	*o = VrrpGroup(varVrrpGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auth_key")
		delete(additionalProperties, "auth_password")
		delete(additionalProperties, "auth_type")
		delete(additionalProperties, "networks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrrpGroup struct {
	value *VrrpGroup
	isSet bool
}

func (v NullableVrrpGroup) Get() *VrrpGroup {
	return v.value
}

func (v *NullableVrrpGroup) Set(val *VrrpGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableVrrpGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableVrrpGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrrpGroup(val *VrrpGroup) *NullableVrrpGroup {
	return &NullableVrrpGroup{value: val, isSet: true}
}

func (v NullableVrrpGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrrpGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

