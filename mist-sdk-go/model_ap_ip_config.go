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

// checks if the ApIpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApIpConfig{}

// ApIpConfig IP AP settings
type ApIpConfig struct {
	// if `type`==`static`
	Dns []string `json:"dns,omitempty"`
	// required if `type`==`static`
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// required if `type`==`static`
	Gateway *string `json:"gateway,omitempty"`
	Gateway6 *string `json:"gateway6,omitempty"`
	// required if `type`==`static`
	Ip *string `json:"ip,omitempty"`
	Ip6 *string `json:"ip6,omitempty"`
	Mtu *int32 `json:"mtu,omitempty"`
	// required if `type`==`static`
	Netmask *string `json:"netmask,omitempty"`
	Netmask6 *string `json:"netmask6,omitempty"`
	Type *IpType `json:"type,omitempty"`
	Type6 *IpType6 `json:"type6,omitempty"`
	// management vlan id, default is 1 (untagged)
	VlanId *int32 `json:"vlan_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApIpConfig ApIpConfig

// NewApIpConfig instantiates a new ApIpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApIpConfig() *ApIpConfig {
	this := ApIpConfig{}
	var type_ IpType = IPTYPE_DHCP
	this.Type = &type_
	var type6 IpType6 = IPTYPE6_DISABLED
	this.Type6 = &type6
	var vlanId int32 = 1
	this.VlanId = &vlanId
	return &this
}

// NewApIpConfigWithDefaults instantiates a new ApIpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApIpConfigWithDefaults() *ApIpConfig {
	this := ApIpConfig{}
	var type_ IpType = IPTYPE_DHCP
	this.Type = &type_
	var type6 IpType6 = IPTYPE6_DISABLED
	this.Type6 = &type6
	var vlanId int32 = 1
	this.VlanId = &vlanId
	return &this
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *ApIpConfig) GetDns() []string {
	if o == nil || IsNil(o.Dns) {
		var ret []string
		return ret
	}
	return o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetDnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Dns) {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *ApIpConfig) HasDns() bool {
	if o != nil && !IsNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given []string and assigns it to the Dns field.
func (o *ApIpConfig) SetDns(v []string) {
	o.Dns = v
}

// GetDnsSuffix returns the DnsSuffix field value if set, zero value otherwise.
func (o *ApIpConfig) GetDnsSuffix() []string {
	if o == nil || IsNil(o.DnsSuffix) {
		var ret []string
		return ret
	}
	return o.DnsSuffix
}

// GetDnsSuffixOk returns a tuple with the DnsSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetDnsSuffixOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsSuffix) {
		return nil, false
	}
	return o.DnsSuffix, true
}

// HasDnsSuffix returns a boolean if a field has been set.
func (o *ApIpConfig) HasDnsSuffix() bool {
	if o != nil && !IsNil(o.DnsSuffix) {
		return true
	}

	return false
}

// SetDnsSuffix gets a reference to the given []string and assigns it to the DnsSuffix field.
func (o *ApIpConfig) SetDnsSuffix(v []string) {
	o.DnsSuffix = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *ApIpConfig) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *ApIpConfig) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *ApIpConfig) SetGateway(v string) {
	o.Gateway = &v
}

// GetGateway6 returns the Gateway6 field value if set, zero value otherwise.
func (o *ApIpConfig) GetGateway6() string {
	if o == nil || IsNil(o.Gateway6) {
		var ret string
		return ret
	}
	return *o.Gateway6
}

// GetGateway6Ok returns a tuple with the Gateway6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetGateway6Ok() (*string, bool) {
	if o == nil || IsNil(o.Gateway6) {
		return nil, false
	}
	return o.Gateway6, true
}

// HasGateway6 returns a boolean if a field has been set.
func (o *ApIpConfig) HasGateway6() bool {
	if o != nil && !IsNil(o.Gateway6) {
		return true
	}

	return false
}

// SetGateway6 gets a reference to the given string and assigns it to the Gateway6 field.
func (o *ApIpConfig) SetGateway6(v string) {
	o.Gateway6 = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ApIpConfig) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ApIpConfig) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *ApIpConfig) SetIp(v string) {
	o.Ip = &v
}

// GetIp6 returns the Ip6 field value if set, zero value otherwise.
func (o *ApIpConfig) GetIp6() string {
	if o == nil || IsNil(o.Ip6) {
		var ret string
		return ret
	}
	return *o.Ip6
}

// GetIp6Ok returns a tuple with the Ip6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetIp6Ok() (*string, bool) {
	if o == nil || IsNil(o.Ip6) {
		return nil, false
	}
	return o.Ip6, true
}

// HasIp6 returns a boolean if a field has been set.
func (o *ApIpConfig) HasIp6() bool {
	if o != nil && !IsNil(o.Ip6) {
		return true
	}

	return false
}

// SetIp6 gets a reference to the given string and assigns it to the Ip6 field.
func (o *ApIpConfig) SetIp6(v string) {
	o.Ip6 = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *ApIpConfig) GetMtu() int32 {
	if o == nil || IsNil(o.Mtu) {
		var ret int32
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetMtuOk() (*int32, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *ApIpConfig) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int32 and assigns it to the Mtu field.
func (o *ApIpConfig) SetMtu(v int32) {
	o.Mtu = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *ApIpConfig) GetNetmask() string {
	if o == nil || IsNil(o.Netmask) {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetNetmaskOk() (*string, bool) {
	if o == nil || IsNil(o.Netmask) {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *ApIpConfig) HasNetmask() bool {
	if o != nil && !IsNil(o.Netmask) {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *ApIpConfig) SetNetmask(v string) {
	o.Netmask = &v
}

// GetNetmask6 returns the Netmask6 field value if set, zero value otherwise.
func (o *ApIpConfig) GetNetmask6() string {
	if o == nil || IsNil(o.Netmask6) {
		var ret string
		return ret
	}
	return *o.Netmask6
}

// GetNetmask6Ok returns a tuple with the Netmask6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetNetmask6Ok() (*string, bool) {
	if o == nil || IsNil(o.Netmask6) {
		return nil, false
	}
	return o.Netmask6, true
}

// HasNetmask6 returns a boolean if a field has been set.
func (o *ApIpConfig) HasNetmask6() bool {
	if o != nil && !IsNil(o.Netmask6) {
		return true
	}

	return false
}

// SetNetmask6 gets a reference to the given string and assigns it to the Netmask6 field.
func (o *ApIpConfig) SetNetmask6(v string) {
	o.Netmask6 = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApIpConfig) GetType() IpType {
	if o == nil || IsNil(o.Type) {
		var ret IpType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetTypeOk() (*IpType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApIpConfig) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given IpType and assigns it to the Type field.
func (o *ApIpConfig) SetType(v IpType) {
	o.Type = &v
}

// GetType6 returns the Type6 field value if set, zero value otherwise.
func (o *ApIpConfig) GetType6() IpType6 {
	if o == nil || IsNil(o.Type6) {
		var ret IpType6
		return ret
	}
	return *o.Type6
}

// GetType6Ok returns a tuple with the Type6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetType6Ok() (*IpType6, bool) {
	if o == nil || IsNil(o.Type6) {
		return nil, false
	}
	return o.Type6, true
}

// HasType6 returns a boolean if a field has been set.
func (o *ApIpConfig) HasType6() bool {
	if o != nil && !IsNil(o.Type6) {
		return true
	}

	return false
}

// SetType6 gets a reference to the given IpType6 and assigns it to the Type6 field.
func (o *ApIpConfig) SetType6(v IpType6) {
	o.Type6 = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *ApIpConfig) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIpConfig) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *ApIpConfig) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *ApIpConfig) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o ApIpConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApIpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !IsNil(o.DnsSuffix) {
		toSerialize["dns_suffix"] = o.DnsSuffix
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Gateway6) {
		toSerialize["gateway6"] = o.Gateway6
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Ip6) {
		toSerialize["ip6"] = o.Ip6
	}
	if !IsNil(o.Mtu) {
		toSerialize["mtu"] = o.Mtu
	}
	if !IsNil(o.Netmask) {
		toSerialize["netmask"] = o.Netmask
	}
	if !IsNil(o.Netmask6) {
		toSerialize["netmask6"] = o.Netmask6
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Type6) {
		toSerialize["type6"] = o.Type6
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlan_id"] = o.VlanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApIpConfig) UnmarshalJSON(data []byte) (err error) {
	varApIpConfig := _ApIpConfig{}

	err = json.Unmarshal(data, &varApIpConfig)

	if err != nil {
		return err
	}

	*o = ApIpConfig(varApIpConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dns")
		delete(additionalProperties, "dns_suffix")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "gateway6")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "ip6")
		delete(additionalProperties, "mtu")
		delete(additionalProperties, "netmask")
		delete(additionalProperties, "netmask6")
		delete(additionalProperties, "type")
		delete(additionalProperties, "type6")
		delete(additionalProperties, "vlan_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApIpConfig struct {
	value *ApIpConfig
	isSet bool
}

func (v NullableApIpConfig) Get() *ApIpConfig {
	return v.value
}

func (v *NullableApIpConfig) Set(val *ApIpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApIpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApIpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApIpConfig(val *ApIpConfig) *NullableApIpConfig {
	return &NullableApIpConfig{value: val, isSet: true}
}

func (v NullableApIpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApIpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

