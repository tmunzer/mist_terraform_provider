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

// checks if the Recover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recover{}

// Recover struct for Recover
type Recover struct {
	Email string `json:"email"`
	// see https://www.google.com/recaptcha/
	Recaptcha *string `json:"recaptcha,omitempty"`
	RecaptchaFlavor *RecaptchaFlavor `json:"recaptcha_flavor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Recover Recover

// NewRecover instantiates a new Recover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecover(email string) *Recover {
	this := Recover{}
	this.Email = email
	var recaptchaFlavor RecaptchaFlavor = RECAPTCHAFLAVOR_GOOGLE
	this.RecaptchaFlavor = &recaptchaFlavor
	return &this
}

// NewRecoverWithDefaults instantiates a new Recover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverWithDefaults() *Recover {
	this := Recover{}
	var recaptchaFlavor RecaptchaFlavor = RECAPTCHAFLAVOR_GOOGLE
	this.RecaptchaFlavor = &recaptchaFlavor
	return &this
}

// GetEmail returns the Email field value
func (o *Recover) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Recover) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Recover) SetEmail(v string) {
	o.Email = v
}

// GetRecaptcha returns the Recaptcha field value if set, zero value otherwise.
func (o *Recover) GetRecaptcha() string {
	if o == nil || IsNil(o.Recaptcha) {
		var ret string
		return ret
	}
	return *o.Recaptcha
}

// GetRecaptchaOk returns a tuple with the Recaptcha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recover) GetRecaptchaOk() (*string, bool) {
	if o == nil || IsNil(o.Recaptcha) {
		return nil, false
	}
	return o.Recaptcha, true
}

// HasRecaptcha returns a boolean if a field has been set.
func (o *Recover) HasRecaptcha() bool {
	if o != nil && !IsNil(o.Recaptcha) {
		return true
	}

	return false
}

// SetRecaptcha gets a reference to the given string and assigns it to the Recaptcha field.
func (o *Recover) SetRecaptcha(v string) {
	o.Recaptcha = &v
}

// GetRecaptchaFlavor returns the RecaptchaFlavor field value if set, zero value otherwise.
func (o *Recover) GetRecaptchaFlavor() RecaptchaFlavor {
	if o == nil || IsNil(o.RecaptchaFlavor) {
		var ret RecaptchaFlavor
		return ret
	}
	return *o.RecaptchaFlavor
}

// GetRecaptchaFlavorOk returns a tuple with the RecaptchaFlavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recover) GetRecaptchaFlavorOk() (*RecaptchaFlavor, bool) {
	if o == nil || IsNil(o.RecaptchaFlavor) {
		return nil, false
	}
	return o.RecaptchaFlavor, true
}

// HasRecaptchaFlavor returns a boolean if a field has been set.
func (o *Recover) HasRecaptchaFlavor() bool {
	if o != nil && !IsNil(o.RecaptchaFlavor) {
		return true
	}

	return false
}

// SetRecaptchaFlavor gets a reference to the given RecaptchaFlavor and assigns it to the RecaptchaFlavor field.
func (o *Recover) SetRecaptchaFlavor(v RecaptchaFlavor) {
	o.RecaptchaFlavor = &v
}

func (o Recover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Recover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.Recaptcha) {
		toSerialize["recaptcha"] = o.Recaptcha
	}
	if !IsNil(o.RecaptchaFlavor) {
		toSerialize["recaptcha_flavor"] = o.RecaptchaFlavor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Recover) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varRecover := _Recover{}

	err = json.Unmarshal(data, &varRecover)

	if err != nil {
		return err
	}

	*o = Recover(varRecover)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "recaptcha")
		delete(additionalProperties, "recaptcha_flavor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecover struct {
	value *Recover
	isSet bool
}

func (v NullableRecover) Get() *Recover {
	return v.value
}

func (v *NullableRecover) Set(val *Recover) {
	v.value = val
	v.isSet = true
}

func (v NullableRecover) IsSet() bool {
	return v.isSet
}

func (v *NullableRecover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecover(val *Recover) *NullableRecover {
	return &NullableRecover{value: val, isSet: true}
}

func (v NullableRecover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

