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

// checks if the ConstApplicationDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstApplicationDefinition{}

// ConstApplicationDefinition struct for ConstApplicationDefinition
type ConstApplicationDefinition struct {
	AppId *bool `json:"app_id,omitempty"`
	AppImageUrl *string `json:"app_image_url,omitempty"`
	AppProbe *bool `json:"app_probe,omitempty"`
	Category *string `json:"category,omitempty"`
	Group *string `json:"group,omitempty"`
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	SignatureBased *bool `json:"signature_based,omitempty"`
	SsrAppId *bool `json:"ssr_app_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstApplicationDefinition ConstApplicationDefinition

// NewConstApplicationDefinition instantiates a new ConstApplicationDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstApplicationDefinition() *ConstApplicationDefinition {
	this := ConstApplicationDefinition{}
	return &this
}

// NewConstApplicationDefinitionWithDefaults instantiates a new ConstApplicationDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstApplicationDefinitionWithDefaults() *ConstApplicationDefinition {
	this := ConstApplicationDefinition{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ConstApplicationDefinition) GetAppId() bool {
	if o == nil || IsNil(o.AppId) {
		var ret bool
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstApplicationDefinition) GetAppIdOk() (*bool, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ConstApplicationDefinition) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given bool and assigns it to the AppId field.
func (o *ConstApplicationDefinition) SetAppId(v bool) {
	o.AppId = &v
}

// GetAppImageUrl returns the AppImageUrl field value if set, zero value otherwise.
func (o *ConstApplicationDefinition) GetAppImageUrl() string {
	if o == nil || IsNil(o.AppImageUrl) {
		var ret string
		return ret
	}
	return *o.AppImageUrl
}

// GetAppImageUrlOk returns a tuple with the AppImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstApplicationDefinition) GetAppImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AppImageUrl) {
		return nil, false
	}
	return o.AppImageUrl, true
}

// HasAppImageUrl returns a boolean if a field has been set.
func (o *ConstApplicationDefinition) HasAppImageUrl() bool {
	if o != nil && !IsNil(o.AppImageUrl) {
		return true
	}

	return false
}

// SetAppImageUrl gets a reference to the given string and assigns it to the AppImageUrl field.
func (o *ConstApplicationDefinition) SetAppImageUrl(v string) {
	o.AppImageUrl = &v
}

// GetAppProbe returns the AppProbe field value if set, zero value otherwise.
func (o *ConstApplicationDefinition) GetAppProbe() bool {
	if o == nil || IsNil(o.AppProbe) {
		var ret bool
		return ret
	}
	return *o.AppProbe
}

// GetAppProbeOk returns a tuple with the AppProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstApplicationDefinition) GetAppProbeOk() (*bool, bool) {
	if o == nil || IsNil(o.AppProbe) {
		return nil, false
	}
	return o.AppProbe, true
}

// HasAppProbe returns a boolean if a field has been set.
func (o *ConstApplicationDefinition) HasAppProbe() bool {
	if o != nil && !IsNil(o.AppProbe) {
		return true
	}

	return false
}

// SetAppProbe gets a reference to the given bool and assigns it to the AppProbe field.
func (o *ConstApplicationDefinition) SetAppProbe(v bool) {
	o.AppProbe = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ConstApplicationDefinition) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstApplicationDefinition) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ConstApplicationDefinition) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ConstApplicationDefinition) SetCategory(v string) {
	o.Category = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ConstApplicationDefinition) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstApplicationDefinition) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ConstApplicationDefinition) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ConstApplicationDefinition) SetGroup(v string) {
	o.Group = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ConstApplicationDefinition) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstApplicationDefinition) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ConstApplicationDefinition) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ConstApplicationDefinition) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConstApplicationDefinition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstApplicationDefinition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConstApplicationDefinition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConstApplicationDefinition) SetName(v string) {
	o.Name = &v
}

// GetSignatureBased returns the SignatureBased field value if set, zero value otherwise.
func (o *ConstApplicationDefinition) GetSignatureBased() bool {
	if o == nil || IsNil(o.SignatureBased) {
		var ret bool
		return ret
	}
	return *o.SignatureBased
}

// GetSignatureBasedOk returns a tuple with the SignatureBased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstApplicationDefinition) GetSignatureBasedOk() (*bool, bool) {
	if o == nil || IsNil(o.SignatureBased) {
		return nil, false
	}
	return o.SignatureBased, true
}

// HasSignatureBased returns a boolean if a field has been set.
func (o *ConstApplicationDefinition) HasSignatureBased() bool {
	if o != nil && !IsNil(o.SignatureBased) {
		return true
	}

	return false
}

// SetSignatureBased gets a reference to the given bool and assigns it to the SignatureBased field.
func (o *ConstApplicationDefinition) SetSignatureBased(v bool) {
	o.SignatureBased = &v
}

// GetSsrAppId returns the SsrAppId field value if set, zero value otherwise.
func (o *ConstApplicationDefinition) GetSsrAppId() bool {
	if o == nil || IsNil(o.SsrAppId) {
		var ret bool
		return ret
	}
	return *o.SsrAppId
}

// GetSsrAppIdOk returns a tuple with the SsrAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstApplicationDefinition) GetSsrAppIdOk() (*bool, bool) {
	if o == nil || IsNil(o.SsrAppId) {
		return nil, false
	}
	return o.SsrAppId, true
}

// HasSsrAppId returns a boolean if a field has been set.
func (o *ConstApplicationDefinition) HasSsrAppId() bool {
	if o != nil && !IsNil(o.SsrAppId) {
		return true
	}

	return false
}

// SetSsrAppId gets a reference to the given bool and assigns it to the SsrAppId field.
func (o *ConstApplicationDefinition) SetSsrAppId(v bool) {
	o.SsrAppId = &v
}

func (o ConstApplicationDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstApplicationDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !IsNil(o.AppImageUrl) {
		toSerialize["app_image_url"] = o.AppImageUrl
	}
	if !IsNil(o.AppProbe) {
		toSerialize["app_probe"] = o.AppProbe
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SignatureBased) {
		toSerialize["signature_based"] = o.SignatureBased
	}
	if !IsNil(o.SsrAppId) {
		toSerialize["ssr_app_id"] = o.SsrAppId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstApplicationDefinition) UnmarshalJSON(data []byte) (err error) {
	varConstApplicationDefinition := _ConstApplicationDefinition{}

	err = json.Unmarshal(data, &varConstApplicationDefinition)

	if err != nil {
		return err
	}

	*o = ConstApplicationDefinition(varConstApplicationDefinition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app_id")
		delete(additionalProperties, "app_image_url")
		delete(additionalProperties, "app_probe")
		delete(additionalProperties, "category")
		delete(additionalProperties, "group")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "signature_based")
		delete(additionalProperties, "ssr_app_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstApplicationDefinition struct {
	value *ConstApplicationDefinition
	isSet bool
}

func (v NullableConstApplicationDefinition) Get() *ConstApplicationDefinition {
	return v.value
}

func (v *NullableConstApplicationDefinition) Set(val *ConstApplicationDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableConstApplicationDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableConstApplicationDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstApplicationDefinition(val *ConstApplicationDefinition) *NullableConstApplicationDefinition {
	return &NullableConstApplicationDefinition{value: val, isSet: true}
}

func (v NullableConstApplicationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstApplicationDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

