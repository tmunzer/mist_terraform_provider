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
	"fmt"
)

// checks if the SiteEngagementDwellTagNames type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteEngagementDwellTagNames{}

// SiteEngagementDwellTagNames struct for SiteEngagementDwellTagNames
type SiteEngagementDwellTagNames struct {
	Bounce string `json:"bounce"`
	Engaged string `json:"engaged"`
	Passerby string `json:"passerby"`
	Stationed string `json:"stationed"`
	AdditionalProperties map[string]interface{}
}

type _SiteEngagementDwellTagNames SiteEngagementDwellTagNames

// NewSiteEngagementDwellTagNames instantiates a new SiteEngagementDwellTagNames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteEngagementDwellTagNames(bounce string, engaged string, passerby string, stationed string) *SiteEngagementDwellTagNames {
	this := SiteEngagementDwellTagNames{}
	this.Bounce = bounce
	this.Engaged = engaged
	this.Passerby = passerby
	this.Stationed = stationed
	return &this
}

// NewSiteEngagementDwellTagNamesWithDefaults instantiates a new SiteEngagementDwellTagNames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteEngagementDwellTagNamesWithDefaults() *SiteEngagementDwellTagNames {
	this := SiteEngagementDwellTagNames{}
	return &this
}

// GetBounce returns the Bounce field value
func (o *SiteEngagementDwellTagNames) GetBounce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bounce
}

// GetBounceOk returns a tuple with the Bounce field value
// and a boolean to check if the value has been set.
func (o *SiteEngagementDwellTagNames) GetBounceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bounce, true
}

// SetBounce sets field value
func (o *SiteEngagementDwellTagNames) SetBounce(v string) {
	o.Bounce = v
}

// GetEngaged returns the Engaged field value
func (o *SiteEngagementDwellTagNames) GetEngaged() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Engaged
}

// GetEngagedOk returns a tuple with the Engaged field value
// and a boolean to check if the value has been set.
func (o *SiteEngagementDwellTagNames) GetEngagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engaged, true
}

// SetEngaged sets field value
func (o *SiteEngagementDwellTagNames) SetEngaged(v string) {
	o.Engaged = v
}

// GetPasserby returns the Passerby field value
func (o *SiteEngagementDwellTagNames) GetPasserby() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Passerby
}

// GetPasserbyOk returns a tuple with the Passerby field value
// and a boolean to check if the value has been set.
func (o *SiteEngagementDwellTagNames) GetPasserbyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Passerby, true
}

// SetPasserby sets field value
func (o *SiteEngagementDwellTagNames) SetPasserby(v string) {
	o.Passerby = v
}

// GetStationed returns the Stationed field value
func (o *SiteEngagementDwellTagNames) GetStationed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stationed
}

// GetStationedOk returns a tuple with the Stationed field value
// and a boolean to check if the value has been set.
func (o *SiteEngagementDwellTagNames) GetStationedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stationed, true
}

// SetStationed sets field value
func (o *SiteEngagementDwellTagNames) SetStationed(v string) {
	o.Stationed = v
}

func (o SiteEngagementDwellTagNames) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteEngagementDwellTagNames) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bounce"] = o.Bounce
	toSerialize["engaged"] = o.Engaged
	toSerialize["passerby"] = o.Passerby
	toSerialize["stationed"] = o.Stationed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteEngagementDwellTagNames) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bounce",
		"engaged",
		"passerby",
		"stationed",
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

	varSiteEngagementDwellTagNames := _SiteEngagementDwellTagNames{}

	err = json.Unmarshal(data, &varSiteEngagementDwellTagNames)

	if err != nil {
		return err
	}

	*o = SiteEngagementDwellTagNames(varSiteEngagementDwellTagNames)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bounce")
		delete(additionalProperties, "engaged")
		delete(additionalProperties, "passerby")
		delete(additionalProperties, "stationed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteEngagementDwellTagNames struct {
	value *SiteEngagementDwellTagNames
	isSet bool
}

func (v NullableSiteEngagementDwellTagNames) Get() *SiteEngagementDwellTagNames {
	return v.value
}

func (v *NullableSiteEngagementDwellTagNames) Set(val *SiteEngagementDwellTagNames) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteEngagementDwellTagNames) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteEngagementDwellTagNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteEngagementDwellTagNames(val *SiteEngagementDwellTagNames) *NullableSiteEngagementDwellTagNames {
	return &NullableSiteEngagementDwellTagNames{value: val, isSet: true}
}

func (v NullableSiteEngagementDwellTagNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteEngagementDwellTagNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

