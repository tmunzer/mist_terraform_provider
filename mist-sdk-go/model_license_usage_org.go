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

// checks if the LicenseUsageOrg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseUsageOrg{}

// LicenseUsageOrg struct for LicenseUsageOrg
type LicenseUsageOrg struct {
	ForSite *bool `json:"for_site,omitempty"`
	// Property key is the service name (e.g. \"SUB-MAN\")
	FullyLoaded *map[string]int32 `json:"fully_loaded,omitempty"`
	NumDevices int32 `json:"num_devices"`
	SiteId string `json:"site_id"`
	// subscriptions and their quantities. Property key is the service name (e.g. \"SUB-MAN\")
	Usages map[string]int32 `json:"usages"`
	AdditionalProperties map[string]interface{}
}

type _LicenseUsageOrg LicenseUsageOrg

// NewLicenseUsageOrg instantiates a new LicenseUsageOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseUsageOrg(numDevices int32, siteId string, usages map[string]int32) *LicenseUsageOrg {
	this := LicenseUsageOrg{}
	this.NumDevices = numDevices
	this.SiteId = siteId
	this.Usages = usages
	return &this
}

// NewLicenseUsageOrgWithDefaults instantiates a new LicenseUsageOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseUsageOrgWithDefaults() *LicenseUsageOrg {
	this := LicenseUsageOrg{}
	return &this
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *LicenseUsageOrg) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseUsageOrg) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *LicenseUsageOrg) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *LicenseUsageOrg) SetForSite(v bool) {
	o.ForSite = &v
}

// GetFullyLoaded returns the FullyLoaded field value if set, zero value otherwise.
func (o *LicenseUsageOrg) GetFullyLoaded() map[string]int32 {
	if o == nil || IsNil(o.FullyLoaded) {
		var ret map[string]int32
		return ret
	}
	return *o.FullyLoaded
}

// GetFullyLoadedOk returns a tuple with the FullyLoaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseUsageOrg) GetFullyLoadedOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.FullyLoaded) {
		return nil, false
	}
	return o.FullyLoaded, true
}

// HasFullyLoaded returns a boolean if a field has been set.
func (o *LicenseUsageOrg) HasFullyLoaded() bool {
	if o != nil && !IsNil(o.FullyLoaded) {
		return true
	}

	return false
}

// SetFullyLoaded gets a reference to the given map[string]int32 and assigns it to the FullyLoaded field.
func (o *LicenseUsageOrg) SetFullyLoaded(v map[string]int32) {
	o.FullyLoaded = &v
}

// GetNumDevices returns the NumDevices field value
func (o *LicenseUsageOrg) GetNumDevices() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumDevices
}

// GetNumDevicesOk returns a tuple with the NumDevices field value
// and a boolean to check if the value has been set.
func (o *LicenseUsageOrg) GetNumDevicesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumDevices, true
}

// SetNumDevices sets field value
func (o *LicenseUsageOrg) SetNumDevices(v int32) {
	o.NumDevices = v
}

// GetSiteId returns the SiteId field value
func (o *LicenseUsageOrg) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *LicenseUsageOrg) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *LicenseUsageOrg) SetSiteId(v string) {
	o.SiteId = v
}

// GetUsages returns the Usages field value
func (o *LicenseUsageOrg) GetUsages() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value
// and a boolean to check if the value has been set.
func (o *LicenseUsageOrg) GetUsagesOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usages, true
}

// SetUsages sets field value
func (o *LicenseUsageOrg) SetUsages(v map[string]int32) {
	o.Usages = v
}

func (o LicenseUsageOrg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseUsageOrg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if !IsNil(o.FullyLoaded) {
		toSerialize["fully_loaded"] = o.FullyLoaded
	}
	toSerialize["num_devices"] = o.NumDevices
	toSerialize["site_id"] = o.SiteId
	toSerialize["usages"] = o.Usages

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LicenseUsageOrg) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"num_devices",
		"site_id",
		"usages",
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

	varLicenseUsageOrg := _LicenseUsageOrg{}

	err = json.Unmarshal(data, &varLicenseUsageOrg)

	if err != nil {
		return err
	}

	*o = LicenseUsageOrg(varLicenseUsageOrg)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "fully_loaded")
		delete(additionalProperties, "num_devices")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "usages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseUsageOrg struct {
	value *LicenseUsageOrg
	isSet bool
}

func (v NullableLicenseUsageOrg) Get() *LicenseUsageOrg {
	return v.value
}

func (v *NullableLicenseUsageOrg) Set(val *LicenseUsageOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseUsageOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseUsageOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseUsageOrg(val *LicenseUsageOrg) *NullableLicenseUsageOrg {
	return &NullableLicenseUsageOrg{value: val, isSet: true}
}

func (v NullableLicenseUsageOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseUsageOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

