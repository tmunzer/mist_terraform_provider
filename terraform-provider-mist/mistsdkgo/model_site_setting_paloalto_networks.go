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

// checks if the SiteSettingPaloaltoNetworks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingPaloaltoNetworks{}

// SiteSettingPaloaltoNetworks struct for SiteSettingPaloaltoNetworks
type SiteSettingPaloaltoNetworks struct {
	Gateways []SiteSettingPaloaltoNetworkGateway `json:"gateways,omitempty"`
	SendMistNacUserInfo *bool `json:"send_mist_nac_user_info,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingPaloaltoNetworks SiteSettingPaloaltoNetworks

// NewSiteSettingPaloaltoNetworks instantiates a new SiteSettingPaloaltoNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingPaloaltoNetworks() *SiteSettingPaloaltoNetworks {
	this := SiteSettingPaloaltoNetworks{}
	var sendMistNacUserInfo bool = false
	this.SendMistNacUserInfo = &sendMistNacUserInfo
	return &this
}

// NewSiteSettingPaloaltoNetworksWithDefaults instantiates a new SiteSettingPaloaltoNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingPaloaltoNetworksWithDefaults() *SiteSettingPaloaltoNetworks {
	this := SiteSettingPaloaltoNetworks{}
	var sendMistNacUserInfo bool = false
	this.SendMistNacUserInfo = &sendMistNacUserInfo
	return &this
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *SiteSettingPaloaltoNetworks) GetGateways() []SiteSettingPaloaltoNetworkGateway {
	if o == nil || IsNil(o.Gateways) {
		var ret []SiteSettingPaloaltoNetworkGateway
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingPaloaltoNetworks) GetGatewaysOk() ([]SiteSettingPaloaltoNetworkGateway, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *SiteSettingPaloaltoNetworks) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []SiteSettingPaloaltoNetworkGateway and assigns it to the Gateways field.
func (o *SiteSettingPaloaltoNetworks) SetGateways(v []SiteSettingPaloaltoNetworkGateway) {
	o.Gateways = v
}

// GetSendMistNacUserInfo returns the SendMistNacUserInfo field value if set, zero value otherwise.
func (o *SiteSettingPaloaltoNetworks) GetSendMistNacUserInfo() bool {
	if o == nil || IsNil(o.SendMistNacUserInfo) {
		var ret bool
		return ret
	}
	return *o.SendMistNacUserInfo
}

// GetSendMistNacUserInfoOk returns a tuple with the SendMistNacUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingPaloaltoNetworks) GetSendMistNacUserInfoOk() (*bool, bool) {
	if o == nil || IsNil(o.SendMistNacUserInfo) {
		return nil, false
	}
	return o.SendMistNacUserInfo, true
}

// HasSendMistNacUserInfo returns a boolean if a field has been set.
func (o *SiteSettingPaloaltoNetworks) HasSendMistNacUserInfo() bool {
	if o != nil && !IsNil(o.SendMistNacUserInfo) {
		return true
	}

	return false
}

// SetSendMistNacUserInfo gets a reference to the given bool and assigns it to the SendMistNacUserInfo field.
func (o *SiteSettingPaloaltoNetworks) SetSendMistNacUserInfo(v bool) {
	o.SendMistNacUserInfo = &v
}

func (o SiteSettingPaloaltoNetworks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingPaloaltoNetworks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	if !IsNil(o.SendMistNacUserInfo) {
		toSerialize["send_mist_nac_user_info"] = o.SendMistNacUserInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingPaloaltoNetworks) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingPaloaltoNetworks := _SiteSettingPaloaltoNetworks{}

	err = json.Unmarshal(data, &varSiteSettingPaloaltoNetworks)

	if err != nil {
		return err
	}

	*o = SiteSettingPaloaltoNetworks(varSiteSettingPaloaltoNetworks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gateways")
		delete(additionalProperties, "send_mist_nac_user_info")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingPaloaltoNetworks struct {
	value *SiteSettingPaloaltoNetworks
	isSet bool
}

func (v NullableSiteSettingPaloaltoNetworks) Get() *SiteSettingPaloaltoNetworks {
	return v.value
}

func (v *NullableSiteSettingPaloaltoNetworks) Set(val *SiteSettingPaloaltoNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingPaloaltoNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingPaloaltoNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingPaloaltoNetworks(val *SiteSettingPaloaltoNetworks) *NullableSiteSettingPaloaltoNetworks {
	return &NullableSiteSettingPaloaltoNetworks{value: val, isSet: true}
}

func (v NullableSiteSettingPaloaltoNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingPaloaltoNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

