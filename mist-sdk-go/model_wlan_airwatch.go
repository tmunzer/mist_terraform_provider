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

// checks if the WlanAirwatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanAirwatch{}

// WlanAirwatch airwatch wlan settings
type WlanAirwatch struct {
	// API Key
	ApiKey *string `json:"api_key,omitempty"`
	// console URL
	ConsoleUrl *string `json:"console_url,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// password
	Password *string `json:"password,omitempty"`
	// username
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WlanAirwatch WlanAirwatch

// NewWlanAirwatch instantiates a new WlanAirwatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanAirwatch() *WlanAirwatch {
	this := WlanAirwatch{}
	return &this
}

// NewWlanAirwatchWithDefaults instantiates a new WlanAirwatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanAirwatchWithDefaults() *WlanAirwatch {
	this := WlanAirwatch{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *WlanAirwatch) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAirwatch) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *WlanAirwatch) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *WlanAirwatch) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetConsoleUrl returns the ConsoleUrl field value if set, zero value otherwise.
func (o *WlanAirwatch) GetConsoleUrl() string {
	if o == nil || IsNil(o.ConsoleUrl) {
		var ret string
		return ret
	}
	return *o.ConsoleUrl
}

// GetConsoleUrlOk returns a tuple with the ConsoleUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAirwatch) GetConsoleUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ConsoleUrl) {
		return nil, false
	}
	return o.ConsoleUrl, true
}

// HasConsoleUrl returns a boolean if a field has been set.
func (o *WlanAirwatch) HasConsoleUrl() bool {
	if o != nil && !IsNil(o.ConsoleUrl) {
		return true
	}

	return false
}

// SetConsoleUrl gets a reference to the given string and assigns it to the ConsoleUrl field.
func (o *WlanAirwatch) SetConsoleUrl(v string) {
	o.ConsoleUrl = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WlanAirwatch) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAirwatch) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WlanAirwatch) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WlanAirwatch) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *WlanAirwatch) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAirwatch) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *WlanAirwatch) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *WlanAirwatch) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *WlanAirwatch) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAirwatch) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *WlanAirwatch) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *WlanAirwatch) SetUsername(v string) {
	o.Username = &v
}

func (o WlanAirwatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanAirwatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	if !IsNil(o.ConsoleUrl) {
		toSerialize["console_url"] = o.ConsoleUrl
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WlanAirwatch) UnmarshalJSON(data []byte) (err error) {
	varWlanAirwatch := _WlanAirwatch{}

	err = json.Unmarshal(data, &varWlanAirwatch)

	if err != nil {
		return err
	}

	*o = WlanAirwatch(varWlanAirwatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_key")
		delete(additionalProperties, "console_url")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "password")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWlanAirwatch struct {
	value *WlanAirwatch
	isSet bool
}

func (v NullableWlanAirwatch) Get() *WlanAirwatch {
	return v.value
}

func (v *NullableWlanAirwatch) Set(val *WlanAirwatch) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanAirwatch) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanAirwatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanAirwatch(val *WlanAirwatch) *NullableWlanAirwatch {
	return &NullableWlanAirwatch{value: val, isSet: true}
}

func (v NullableWlanAirwatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanAirwatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

