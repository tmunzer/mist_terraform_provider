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

// checks if the ResponseAutoPlacementInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseAutoPlacementInfo{}

// ResponseAutoPlacementInfo struct for ResponseAutoPlacementInfo
type ResponseAutoPlacementInfo struct {
	// time when autoplacement completed or was manually stopped
	EndTime *float32 `json:"end_time,omitempty"`
	// (Only when inprogress) estimate of the time to completion
	EstTimeLeft *float32 `json:"est_time_left,omitempty"`
	// time when autoplacement process was last queued for this map
	StartTime *float32 `json:"start_time,omitempty"`
	Status *AutoPlacementInfoStatus `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseAutoPlacementInfo ResponseAutoPlacementInfo

// NewResponseAutoPlacementInfo instantiates a new ResponseAutoPlacementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAutoPlacementInfo() *ResponseAutoPlacementInfo {
	this := ResponseAutoPlacementInfo{}
	return &this
}

// NewResponseAutoPlacementInfoWithDefaults instantiates a new ResponseAutoPlacementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAutoPlacementInfoWithDefaults() *ResponseAutoPlacementInfo {
	this := ResponseAutoPlacementInfo{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ResponseAutoPlacementInfo) GetEndTime() float32 {
	if o == nil || IsNil(o.EndTime) {
		var ret float32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAutoPlacementInfo) GetEndTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ResponseAutoPlacementInfo) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given float32 and assigns it to the EndTime field.
func (o *ResponseAutoPlacementInfo) SetEndTime(v float32) {
	o.EndTime = &v
}

// GetEstTimeLeft returns the EstTimeLeft field value if set, zero value otherwise.
func (o *ResponseAutoPlacementInfo) GetEstTimeLeft() float32 {
	if o == nil || IsNil(o.EstTimeLeft) {
		var ret float32
		return ret
	}
	return *o.EstTimeLeft
}

// GetEstTimeLeftOk returns a tuple with the EstTimeLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAutoPlacementInfo) GetEstTimeLeftOk() (*float32, bool) {
	if o == nil || IsNil(o.EstTimeLeft) {
		return nil, false
	}
	return o.EstTimeLeft, true
}

// HasEstTimeLeft returns a boolean if a field has been set.
func (o *ResponseAutoPlacementInfo) HasEstTimeLeft() bool {
	if o != nil && !IsNil(o.EstTimeLeft) {
		return true
	}

	return false
}

// SetEstTimeLeft gets a reference to the given float32 and assigns it to the EstTimeLeft field.
func (o *ResponseAutoPlacementInfo) SetEstTimeLeft(v float32) {
	o.EstTimeLeft = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ResponseAutoPlacementInfo) GetStartTime() float32 {
	if o == nil || IsNil(o.StartTime) {
		var ret float32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAutoPlacementInfo) GetStartTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ResponseAutoPlacementInfo) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given float32 and assigns it to the StartTime field.
func (o *ResponseAutoPlacementInfo) SetStartTime(v float32) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseAutoPlacementInfo) GetStatus() AutoPlacementInfoStatus {
	if o == nil || IsNil(o.Status) {
		var ret AutoPlacementInfoStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAutoPlacementInfo) GetStatusOk() (*AutoPlacementInfoStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseAutoPlacementInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AutoPlacementInfoStatus and assigns it to the Status field.
func (o *ResponseAutoPlacementInfo) SetStatus(v AutoPlacementInfoStatus) {
	o.Status = &v
}

func (o ResponseAutoPlacementInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAutoPlacementInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.EstTimeLeft) {
		toSerialize["est_time_left"] = o.EstTimeLeft
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseAutoPlacementInfo) UnmarshalJSON(data []byte) (err error) {
	varResponseAutoPlacementInfo := _ResponseAutoPlacementInfo{}

	err = json.Unmarshal(data, &varResponseAutoPlacementInfo)

	if err != nil {
		return err
	}

	*o = ResponseAutoPlacementInfo(varResponseAutoPlacementInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end_time")
		delete(additionalProperties, "est_time_left")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseAutoPlacementInfo struct {
	value *ResponseAutoPlacementInfo
	isSet bool
}

func (v NullableResponseAutoPlacementInfo) Get() *ResponseAutoPlacementInfo {
	return v.value
}

func (v *NullableResponseAutoPlacementInfo) Set(val *ResponseAutoPlacementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAutoPlacementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAutoPlacementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAutoPlacementInfo(val *ResponseAutoPlacementInfo) *NullableResponseAutoPlacementInfo {
	return &NullableResponseAutoPlacementInfo{value: val, isSet: true}
}

func (v NullableResponseAutoPlacementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAutoPlacementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

