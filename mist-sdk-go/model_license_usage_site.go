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

// checks if the LicenseUsageSite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseUsageSite{}

// LicenseUsageSite struct for LicenseUsageSite
type LicenseUsageSite struct {
	// license entitlement for the entire org
	OrgEntitled map[string]int32 `json:"org_entitled"`
	// eligibility for the Switch SLE
	SvnaEnabled bool `json:"svna_enabled"`
	TrialEnabled bool `json:"trial_enabled"`
	// subscriptions and their quantities
	Usages map[string]int32 `json:"usages"`
	// eligibility for the AP/Client SLE
	VnaEligible bool `json:"vna_eligible"`
	// if True, Conversational Assistant and Marvis Action available
	VnaUi bool `json:"vna_ui"`
	// eligibility for the WAN SLE
	WvnaEligible bool `json:"wvna_eligible"`
	AdditionalProperties map[string]interface{}
}

type _LicenseUsageSite LicenseUsageSite

// NewLicenseUsageSite instantiates a new LicenseUsageSite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseUsageSite(orgEntitled map[string]int32, svnaEnabled bool, trialEnabled bool, usages map[string]int32, vnaEligible bool, vnaUi bool, wvnaEligible bool) *LicenseUsageSite {
	this := LicenseUsageSite{}
	this.OrgEntitled = orgEntitled
	this.SvnaEnabled = svnaEnabled
	this.TrialEnabled = trialEnabled
	this.Usages = usages
	this.VnaEligible = vnaEligible
	this.VnaUi = vnaUi
	this.WvnaEligible = wvnaEligible
	return &this
}

// NewLicenseUsageSiteWithDefaults instantiates a new LicenseUsageSite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseUsageSiteWithDefaults() *LicenseUsageSite {
	this := LicenseUsageSite{}
	return &this
}

// GetOrgEntitled returns the OrgEntitled field value
func (o *LicenseUsageSite) GetOrgEntitled() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.OrgEntitled
}

// GetOrgEntitledOk returns a tuple with the OrgEntitled field value
// and a boolean to check if the value has been set.
func (o *LicenseUsageSite) GetOrgEntitledOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgEntitled, true
}

// SetOrgEntitled sets field value
func (o *LicenseUsageSite) SetOrgEntitled(v map[string]int32) {
	o.OrgEntitled = v
}

// GetSvnaEnabled returns the SvnaEnabled field value
func (o *LicenseUsageSite) GetSvnaEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SvnaEnabled
}

// GetSvnaEnabledOk returns a tuple with the SvnaEnabled field value
// and a boolean to check if the value has been set.
func (o *LicenseUsageSite) GetSvnaEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SvnaEnabled, true
}

// SetSvnaEnabled sets field value
func (o *LicenseUsageSite) SetSvnaEnabled(v bool) {
	o.SvnaEnabled = v
}

// GetTrialEnabled returns the TrialEnabled field value
func (o *LicenseUsageSite) GetTrialEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TrialEnabled
}

// GetTrialEnabledOk returns a tuple with the TrialEnabled field value
// and a boolean to check if the value has been set.
func (o *LicenseUsageSite) GetTrialEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrialEnabled, true
}

// SetTrialEnabled sets field value
func (o *LicenseUsageSite) SetTrialEnabled(v bool) {
	o.TrialEnabled = v
}

// GetUsages returns the Usages field value
func (o *LicenseUsageSite) GetUsages() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value
// and a boolean to check if the value has been set.
func (o *LicenseUsageSite) GetUsagesOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usages, true
}

// SetUsages sets field value
func (o *LicenseUsageSite) SetUsages(v map[string]int32) {
	o.Usages = v
}

// GetVnaEligible returns the VnaEligible field value
func (o *LicenseUsageSite) GetVnaEligible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VnaEligible
}

// GetVnaEligibleOk returns a tuple with the VnaEligible field value
// and a boolean to check if the value has been set.
func (o *LicenseUsageSite) GetVnaEligibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VnaEligible, true
}

// SetVnaEligible sets field value
func (o *LicenseUsageSite) SetVnaEligible(v bool) {
	o.VnaEligible = v
}

// GetVnaUi returns the VnaUi field value
func (o *LicenseUsageSite) GetVnaUi() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VnaUi
}

// GetVnaUiOk returns a tuple with the VnaUi field value
// and a boolean to check if the value has been set.
func (o *LicenseUsageSite) GetVnaUiOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VnaUi, true
}

// SetVnaUi sets field value
func (o *LicenseUsageSite) SetVnaUi(v bool) {
	o.VnaUi = v
}

// GetWvnaEligible returns the WvnaEligible field value
func (o *LicenseUsageSite) GetWvnaEligible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WvnaEligible
}

// GetWvnaEligibleOk returns a tuple with the WvnaEligible field value
// and a boolean to check if the value has been set.
func (o *LicenseUsageSite) GetWvnaEligibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WvnaEligible, true
}

// SetWvnaEligible sets field value
func (o *LicenseUsageSite) SetWvnaEligible(v bool) {
	o.WvnaEligible = v
}

func (o LicenseUsageSite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseUsageSite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_entitled"] = o.OrgEntitled
	toSerialize["svna_enabled"] = o.SvnaEnabled
	toSerialize["trial_enabled"] = o.TrialEnabled
	toSerialize["usages"] = o.Usages
	toSerialize["vna_eligible"] = o.VnaEligible
	toSerialize["vna_ui"] = o.VnaUi
	toSerialize["wvna_eligible"] = o.WvnaEligible

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LicenseUsageSite) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"org_entitled",
		"svna_enabled",
		"trial_enabled",
		"usages",
		"vna_eligible",
		"vna_ui",
		"wvna_eligible",
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

	varLicenseUsageSite := _LicenseUsageSite{}

	err = json.Unmarshal(data, &varLicenseUsageSite)

	if err != nil {
		return err
	}

	*o = LicenseUsageSite(varLicenseUsageSite)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "org_entitled")
		delete(additionalProperties, "svna_enabled")
		delete(additionalProperties, "trial_enabled")
		delete(additionalProperties, "usages")
		delete(additionalProperties, "vna_eligible")
		delete(additionalProperties, "vna_ui")
		delete(additionalProperties, "wvna_eligible")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseUsageSite struct {
	value *LicenseUsageSite
	isSet bool
}

func (v NullableLicenseUsageSite) Get() *LicenseUsageSite {
	return v.value
}

func (v *NullableLicenseUsageSite) Set(val *LicenseUsageSite) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseUsageSite) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseUsageSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseUsageSite(val *LicenseUsageSite) *NullableLicenseUsageSite {
	return &NullableLicenseUsageSite{value: val, isSet: true}
}

func (v NullableLicenseUsageSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseUsageSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

