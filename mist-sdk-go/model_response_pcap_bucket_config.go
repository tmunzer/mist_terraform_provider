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

// checks if the ResponsePcapBucketConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsePcapBucketConfig{}

// ResponsePcapBucketConfig struct for ResponsePcapBucketConfig
type ResponsePcapBucketConfig struct {
	Bucket *string `json:"bucket,omitempty"`
	Detail *string `json:"detail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponsePcapBucketConfig ResponsePcapBucketConfig

// NewResponsePcapBucketConfig instantiates a new ResponsePcapBucketConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsePcapBucketConfig() *ResponsePcapBucketConfig {
	this := ResponsePcapBucketConfig{}
	return &this
}

// NewResponsePcapBucketConfigWithDefaults instantiates a new ResponsePcapBucketConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsePcapBucketConfigWithDefaults() *ResponsePcapBucketConfig {
	this := ResponsePcapBucketConfig{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *ResponsePcapBucketConfig) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapBucketConfig) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *ResponsePcapBucketConfig) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *ResponsePcapBucketConfig) SetBucket(v string) {
	o.Bucket = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponsePcapBucketConfig) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapBucketConfig) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponsePcapBucketConfig) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponsePcapBucketConfig) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponsePcapBucketConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsePcapBucketConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponsePcapBucketConfig) UnmarshalJSON(data []byte) (err error) {
	varResponsePcapBucketConfig := _ResponsePcapBucketConfig{}

	err = json.Unmarshal(data, &varResponsePcapBucketConfig)

	if err != nil {
		return err
	}

	*o = ResponsePcapBucketConfig(varResponsePcapBucketConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bucket")
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponsePcapBucketConfig struct {
	value *ResponsePcapBucketConfig
	isSet bool
}

func (v NullableResponsePcapBucketConfig) Get() *ResponsePcapBucketConfig {
	return v.value
}

func (v *NullableResponsePcapBucketConfig) Set(val *ResponsePcapBucketConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePcapBucketConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePcapBucketConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePcapBucketConfig(val *ResponsePcapBucketConfig) *NullableResponsePcapBucketConfig {
	return &NullableResponsePcapBucketConfig{value: val, isSet: true}
}

func (v NullableResponsePcapBucketConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePcapBucketConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

