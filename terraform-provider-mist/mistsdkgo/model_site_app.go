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

// checks if the SiteApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteApp{}

// SiteApp struct for SiteApp
type SiteApp struct {
	Group string `json:"group"`
	Key string `json:"key"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _SiteApp SiteApp

// NewSiteApp instantiates a new SiteApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteApp(group string, key string, name string) *SiteApp {
	this := SiteApp{}
	this.Group = group
	this.Key = key
	this.Name = name
	return &this
}

// NewSiteAppWithDefaults instantiates a new SiteApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAppWithDefaults() *SiteApp {
	this := SiteApp{}
	return &this
}

// GetGroup returns the Group field value
func (o *SiteApp) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SiteApp) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SiteApp) SetGroup(v string) {
	o.Group = v
}

// GetKey returns the Key field value
func (o *SiteApp) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SiteApp) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SiteApp) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *SiteApp) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteApp) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteApp) SetName(v string) {
	o.Name = v
}

func (o SiteApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteApp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group",
		"key",
		"name",
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

	varSiteApp := _SiteApp{}

	err = json.Unmarshal(data, &varSiteApp)

	if err != nil {
		return err
	}

	*o = SiteApp(varSiteApp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteApp struct {
	value *SiteApp
	isSet bool
}

func (v NullableSiteApp) Get() *SiteApp {
	return v.value
}

func (v *NullableSiteApp) Set(val *SiteApp) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteApp) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteApp(val *SiteApp) *NullableSiteApp {
	return &NullableSiteApp{value: val, isSet: true}
}

func (v NullableSiteApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

