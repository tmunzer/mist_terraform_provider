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

// checks if the RemoteSyslogContentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteSyslogContentItem{}

// RemoteSyslogContentItem struct for RemoteSyslogContentItem
type RemoteSyslogContentItem struct {
	Facility *RemoteSyslogFacility `json:"facility,omitempty"`
	Severity *RemoteSyslogSeverity `json:"severity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RemoteSyslogContentItem RemoteSyslogContentItem

// NewRemoteSyslogContentItem instantiates a new RemoteSyslogContentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteSyslogContentItem() *RemoteSyslogContentItem {
	this := RemoteSyslogContentItem{}
	return &this
}

// NewRemoteSyslogContentItemWithDefaults instantiates a new RemoteSyslogContentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteSyslogContentItemWithDefaults() *RemoteSyslogContentItem {
	this := RemoteSyslogContentItem{}
	return &this
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *RemoteSyslogContentItem) GetFacility() RemoteSyslogFacility {
	if o == nil || IsNil(o.Facility) {
		var ret RemoteSyslogFacility
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteSyslogContentItem) GetFacilityOk() (*RemoteSyslogFacility, bool) {
	if o == nil || IsNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *RemoteSyslogContentItem) HasFacility() bool {
	if o != nil && !IsNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given RemoteSyslogFacility and assigns it to the Facility field.
func (o *RemoteSyslogContentItem) SetFacility(v RemoteSyslogFacility) {
	o.Facility = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *RemoteSyslogContentItem) GetSeverity() RemoteSyslogSeverity {
	if o == nil || IsNil(o.Severity) {
		var ret RemoteSyslogSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteSyslogContentItem) GetSeverityOk() (*RemoteSyslogSeverity, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *RemoteSyslogContentItem) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given RemoteSyslogSeverity and assigns it to the Severity field.
func (o *RemoteSyslogContentItem) SetSeverity(v RemoteSyslogSeverity) {
	o.Severity = &v
}

func (o RemoteSyslogContentItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteSyslogContentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Facility) {
		toSerialize["facility"] = o.Facility
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RemoteSyslogContentItem) UnmarshalJSON(data []byte) (err error) {
	varRemoteSyslogContentItem := _RemoteSyslogContentItem{}

	err = json.Unmarshal(data, &varRemoteSyslogContentItem)

	if err != nil {
		return err
	}

	*o = RemoteSyslogContentItem(varRemoteSyslogContentItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "facility")
		delete(additionalProperties, "severity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemoteSyslogContentItem struct {
	value *RemoteSyslogContentItem
	isSet bool
}

func (v NullableRemoteSyslogContentItem) Get() *RemoteSyslogContentItem {
	return v.value
}

func (v *NullableRemoteSyslogContentItem) Set(val *RemoteSyslogContentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteSyslogContentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteSyslogContentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteSyslogContentItem(val *RemoteSyslogContentItem) *NullableRemoteSyslogContentItem {
	return &NullableRemoteSyslogContentItem{value: val, isSet: true}
}

func (v NullableRemoteSyslogContentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteSyslogContentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

