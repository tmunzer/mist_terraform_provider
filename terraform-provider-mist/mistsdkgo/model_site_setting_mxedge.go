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

// checks if the SiteSettingMxedge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingMxedge{}

// SiteSettingMxedge site mist edges form a cluster of radsecproxy servers
type SiteSettingMxedge struct {
	MistDas *MxedgeDas `json:"mist_das,omitempty"`
	Radsec *MxclusterRadsec `json:"radsec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingMxedge SiteSettingMxedge

// NewSiteSettingMxedge instantiates a new SiteSettingMxedge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingMxedge() *SiteSettingMxedge {
	this := SiteSettingMxedge{}
	return &this
}

// NewSiteSettingMxedgeWithDefaults instantiates a new SiteSettingMxedge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingMxedgeWithDefaults() *SiteSettingMxedge {
	this := SiteSettingMxedge{}
	return &this
}

// GetMistDas returns the MistDas field value if set, zero value otherwise.
func (o *SiteSettingMxedge) GetMistDas() MxedgeDas {
	if o == nil || IsNil(o.MistDas) {
		var ret MxedgeDas
		return ret
	}
	return *o.MistDas
}

// GetMistDasOk returns a tuple with the MistDas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingMxedge) GetMistDasOk() (*MxedgeDas, bool) {
	if o == nil || IsNil(o.MistDas) {
		return nil, false
	}
	return o.MistDas, true
}

// HasMistDas returns a boolean if a field has been set.
func (o *SiteSettingMxedge) HasMistDas() bool {
	if o != nil && !IsNil(o.MistDas) {
		return true
	}

	return false
}

// SetMistDas gets a reference to the given MxedgeDas and assigns it to the MistDas field.
func (o *SiteSettingMxedge) SetMistDas(v MxedgeDas) {
	o.MistDas = &v
}

// GetRadsec returns the Radsec field value if set, zero value otherwise.
func (o *SiteSettingMxedge) GetRadsec() MxclusterRadsec {
	if o == nil || IsNil(o.Radsec) {
		var ret MxclusterRadsec
		return ret
	}
	return *o.Radsec
}

// GetRadsecOk returns a tuple with the Radsec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingMxedge) GetRadsecOk() (*MxclusterRadsec, bool) {
	if o == nil || IsNil(o.Radsec) {
		return nil, false
	}
	return o.Radsec, true
}

// HasRadsec returns a boolean if a field has been set.
func (o *SiteSettingMxedge) HasRadsec() bool {
	if o != nil && !IsNil(o.Radsec) {
		return true
	}

	return false
}

// SetRadsec gets a reference to the given MxclusterRadsec and assigns it to the Radsec field.
func (o *SiteSettingMxedge) SetRadsec(v MxclusterRadsec) {
	o.Radsec = &v
}

func (o SiteSettingMxedge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingMxedge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MistDas) {
		toSerialize["mist_das"] = o.MistDas
	}
	if !IsNil(o.Radsec) {
		toSerialize["radsec"] = o.Radsec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingMxedge) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingMxedge := _SiteSettingMxedge{}

	err = json.Unmarshal(data, &varSiteSettingMxedge)

	if err != nil {
		return err
	}

	*o = SiteSettingMxedge(varSiteSettingMxedge)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mist_das")
		delete(additionalProperties, "radsec")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingMxedge struct {
	value *SiteSettingMxedge
	isSet bool
}

func (v NullableSiteSettingMxedge) Get() *SiteSettingMxedge {
	return v.value
}

func (v *NullableSiteSettingMxedge) Set(val *SiteSettingMxedge) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingMxedge) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingMxedge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingMxedge(val *SiteSettingMxedge) *NullableSiteSettingMxedge {
	return &NullableSiteSettingMxedge{value: val, isSet: true}
}

func (v NullableSiteSettingMxedge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingMxedge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

