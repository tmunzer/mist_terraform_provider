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

// checks if the OrgDeviceUpgradeSiteUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgDeviceUpgradeSiteUpgrade{}

// OrgDeviceUpgradeSiteUpgrade struct for OrgDeviceUpgradeSiteUpgrade
type OrgDeviceUpgradeSiteUpgrade struct {
	SiteId *string `json:"site_id,omitempty"`
	UpgradeId *string `json:"upgrade_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgDeviceUpgradeSiteUpgrade OrgDeviceUpgradeSiteUpgrade

// NewOrgDeviceUpgradeSiteUpgrade instantiates a new OrgDeviceUpgradeSiteUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgDeviceUpgradeSiteUpgrade() *OrgDeviceUpgradeSiteUpgrade {
	this := OrgDeviceUpgradeSiteUpgrade{}
	return &this
}

// NewOrgDeviceUpgradeSiteUpgradeWithDefaults instantiates a new OrgDeviceUpgradeSiteUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgDeviceUpgradeSiteUpgradeWithDefaults() *OrgDeviceUpgradeSiteUpgrade {
	this := OrgDeviceUpgradeSiteUpgrade{}
	return &this
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *OrgDeviceUpgradeSiteUpgrade) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDeviceUpgradeSiteUpgrade) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *OrgDeviceUpgradeSiteUpgrade) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *OrgDeviceUpgradeSiteUpgrade) SetSiteId(v string) {
	o.SiteId = &v
}

// GetUpgradeId returns the UpgradeId field value if set, zero value otherwise.
func (o *OrgDeviceUpgradeSiteUpgrade) GetUpgradeId() string {
	if o == nil || IsNil(o.UpgradeId) {
		var ret string
		return ret
	}
	return *o.UpgradeId
}

// GetUpgradeIdOk returns a tuple with the UpgradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDeviceUpgradeSiteUpgrade) GetUpgradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeId) {
		return nil, false
	}
	return o.UpgradeId, true
}

// HasUpgradeId returns a boolean if a field has been set.
func (o *OrgDeviceUpgradeSiteUpgrade) HasUpgradeId() bool {
	if o != nil && !IsNil(o.UpgradeId) {
		return true
	}

	return false
}

// SetUpgradeId gets a reference to the given string and assigns it to the UpgradeId field.
func (o *OrgDeviceUpgradeSiteUpgrade) SetUpgradeId(v string) {
	o.UpgradeId = &v
}

func (o OrgDeviceUpgradeSiteUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgDeviceUpgradeSiteUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.UpgradeId) {
		toSerialize["upgrade_id"] = o.UpgradeId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgDeviceUpgradeSiteUpgrade) UnmarshalJSON(data []byte) (err error) {
	varOrgDeviceUpgradeSiteUpgrade := _OrgDeviceUpgradeSiteUpgrade{}

	err = json.Unmarshal(data, &varOrgDeviceUpgradeSiteUpgrade)

	if err != nil {
		return err
	}

	*o = OrgDeviceUpgradeSiteUpgrade(varOrgDeviceUpgradeSiteUpgrade)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "upgrade_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgDeviceUpgradeSiteUpgrade struct {
	value *OrgDeviceUpgradeSiteUpgrade
	isSet bool
}

func (v NullableOrgDeviceUpgradeSiteUpgrade) Get() *OrgDeviceUpgradeSiteUpgrade {
	return v.value
}

func (v *NullableOrgDeviceUpgradeSiteUpgrade) Set(val *OrgDeviceUpgradeSiteUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgDeviceUpgradeSiteUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgDeviceUpgradeSiteUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgDeviceUpgradeSiteUpgrade(val *OrgDeviceUpgradeSiteUpgrade) *NullableOrgDeviceUpgradeSiteUpgrade {
	return &NullableOrgDeviceUpgradeSiteUpgrade{value: val, isSet: true}
}

func (v NullableOrgDeviceUpgradeSiteUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgDeviceUpgradeSiteUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

