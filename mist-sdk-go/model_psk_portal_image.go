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
	"os"
)

// checks if the PskPortalImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PskPortalImage{}

// PskPortalImage struct for PskPortalImage
type PskPortalImage struct {
	// Binary file
	File **os.File `json:"file,omitempty"`
	// JSON string describing the upload
	Json *string `json:"json,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PskPortalImage PskPortalImage

// NewPskPortalImage instantiates a new PskPortalImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPskPortalImage() *PskPortalImage {
	this := PskPortalImage{}
	return &this
}

// NewPskPortalImageWithDefaults instantiates a new PskPortalImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPskPortalImageWithDefaults() *PskPortalImage {
	this := PskPortalImage{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *PskPortalImage) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PskPortalImage) GetFileOk() (**os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *PskPortalImage) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *PskPortalImage) SetFile(v *os.File) {
	o.File = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *PskPortalImage) GetJson() string {
	if o == nil || IsNil(o.Json) {
		var ret string
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PskPortalImage) GetJsonOk() (*string, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *PskPortalImage) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given string and assigns it to the Json field.
func (o *PskPortalImage) SetJson(v string) {
	o.Json = &v
}

func (o PskPortalImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PskPortalImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PskPortalImage) UnmarshalJSON(data []byte) (err error) {
	varPskPortalImage := _PskPortalImage{}

	err = json.Unmarshal(data, &varPskPortalImage)

	if err != nil {
		return err
	}

	*o = PskPortalImage(varPskPortalImage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		delete(additionalProperties, "json")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePskPortalImage struct {
	value *PskPortalImage
	isSet bool
}

func (v NullablePskPortalImage) Get() *PskPortalImage {
	return v.value
}

func (v *NullablePskPortalImage) Set(val *PskPortalImage) {
	v.value = val
	v.isSet = true
}

func (v NullablePskPortalImage) IsSet() bool {
	return v.isSet
}

func (v *NullablePskPortalImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePskPortalImage(val *PskPortalImage) *NullablePskPortalImage {
	return &NullablePskPortalImage{value: val, isSet: true}
}

func (v NullablePskPortalImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePskPortalImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

