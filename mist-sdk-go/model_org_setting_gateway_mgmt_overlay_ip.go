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

// checks if the OrgSettingGatewayMgmtOverlayIp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingGatewayMgmtOverlayIp{}

// OrgSettingGatewayMgmtOverlayIp struct for OrgSettingGatewayMgmtOverlayIp
type OrgSettingGatewayMgmtOverlayIp struct {
	// when it's going overlay, a routable IP to overlay will be required
	Ip *string `json:"ip,omitempty"`
	// for SSR HA cluster, another IP for node1 will be required, too
	Node1Ip *string `json:"node1_ip,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingGatewayMgmtOverlayIp OrgSettingGatewayMgmtOverlayIp

// NewOrgSettingGatewayMgmtOverlayIp instantiates a new OrgSettingGatewayMgmtOverlayIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingGatewayMgmtOverlayIp() *OrgSettingGatewayMgmtOverlayIp {
	this := OrgSettingGatewayMgmtOverlayIp{}
	return &this
}

// NewOrgSettingGatewayMgmtOverlayIpWithDefaults instantiates a new OrgSettingGatewayMgmtOverlayIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingGatewayMgmtOverlayIpWithDefaults() *OrgSettingGatewayMgmtOverlayIp {
	this := OrgSettingGatewayMgmtOverlayIp{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *OrgSettingGatewayMgmtOverlayIp) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingGatewayMgmtOverlayIp) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *OrgSettingGatewayMgmtOverlayIp) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *OrgSettingGatewayMgmtOverlayIp) SetIp(v string) {
	o.Ip = &v
}

// GetNode1Ip returns the Node1Ip field value if set, zero value otherwise.
func (o *OrgSettingGatewayMgmtOverlayIp) GetNode1Ip() string {
	if o == nil || IsNil(o.Node1Ip) {
		var ret string
		return ret
	}
	return *o.Node1Ip
}

// GetNode1IpOk returns a tuple with the Node1Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingGatewayMgmtOverlayIp) GetNode1IpOk() (*string, bool) {
	if o == nil || IsNil(o.Node1Ip) {
		return nil, false
	}
	return o.Node1Ip, true
}

// HasNode1Ip returns a boolean if a field has been set.
func (o *OrgSettingGatewayMgmtOverlayIp) HasNode1Ip() bool {
	if o != nil && !IsNil(o.Node1Ip) {
		return true
	}

	return false
}

// SetNode1Ip gets a reference to the given string and assigns it to the Node1Ip field.
func (o *OrgSettingGatewayMgmtOverlayIp) SetNode1Ip(v string) {
	o.Node1Ip = &v
}

func (o OrgSettingGatewayMgmtOverlayIp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingGatewayMgmtOverlayIp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Node1Ip) {
		toSerialize["node1_ip"] = o.Node1Ip
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingGatewayMgmtOverlayIp) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingGatewayMgmtOverlayIp := _OrgSettingGatewayMgmtOverlayIp{}

	err = json.Unmarshal(data, &varOrgSettingGatewayMgmtOverlayIp)

	if err != nil {
		return err
	}

	*o = OrgSettingGatewayMgmtOverlayIp(varOrgSettingGatewayMgmtOverlayIp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ip")
		delete(additionalProperties, "node1_ip")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingGatewayMgmtOverlayIp struct {
	value *OrgSettingGatewayMgmtOverlayIp
	isSet bool
}

func (v NullableOrgSettingGatewayMgmtOverlayIp) Get() *OrgSettingGatewayMgmtOverlayIp {
	return v.value
}

func (v *NullableOrgSettingGatewayMgmtOverlayIp) Set(val *OrgSettingGatewayMgmtOverlayIp) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingGatewayMgmtOverlayIp) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingGatewayMgmtOverlayIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingGatewayMgmtOverlayIp(val *OrgSettingGatewayMgmtOverlayIp) *NullableOrgSettingGatewayMgmtOverlayIp {
	return &NullableOrgSettingGatewayMgmtOverlayIp{value: val, isSet: true}
}

func (v NullableOrgSettingGatewayMgmtOverlayIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingGatewayMgmtOverlayIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

