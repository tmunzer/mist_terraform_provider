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

// checks if the ResponseCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCertificate{}

// ResponseCertificate struct for ResponseCertificate
type ResponseCertificate struct {
	Cert string `json:"cert"`
	AdditionalProperties map[string]interface{}
}

type _ResponseCertificate ResponseCertificate

// NewResponseCertificate instantiates a new ResponseCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCertificate(cert string) *ResponseCertificate {
	this := ResponseCertificate{}
	this.Cert = cert
	return &this
}

// NewResponseCertificateWithDefaults instantiates a new ResponseCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCertificateWithDefaults() *ResponseCertificate {
	this := ResponseCertificate{}
	return &this
}

// GetCert returns the Cert field value
func (o *ResponseCertificate) GetCert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cert
}

// GetCertOk returns a tuple with the Cert field value
// and a boolean to check if the value has been set.
func (o *ResponseCertificate) GetCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cert, true
}

// SetCert sets field value
func (o *ResponseCertificate) SetCert(v string) {
	o.Cert = v
}

func (o ResponseCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cert"] = o.Cert

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseCertificate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cert",
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

	varResponseCertificate := _ResponseCertificate{}

	err = json.Unmarshal(data, &varResponseCertificate)

	if err != nil {
		return err
	}

	*o = ResponseCertificate(varResponseCertificate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cert")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseCertificate struct {
	value *ResponseCertificate
	isSet bool
}

func (v NullableResponseCertificate) Get() *ResponseCertificate {
	return v.value
}

func (v *NullableResponseCertificate) Set(val *ResponseCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCertificate(val *ResponseCertificate) *NullableResponseCertificate {
	return &NullableResponseCertificate{value: val, isSet: true}
}

func (v NullableResponseCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

