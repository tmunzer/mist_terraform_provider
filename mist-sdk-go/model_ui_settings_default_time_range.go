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

// checks if the UiSettingsDefaultTimeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiSettingsDefaultTimeRange{}

// UiSettingsDefaultTimeRange struct for UiSettingsDefaultTimeRange
type UiSettingsDefaultTimeRange struct {
	End *int32 `json:"end,omitempty"`
	EndDate *string `json:"endDate,omitempty"`
	Interval *string `json:"interval,omitempty"`
	Name *string `json:"name,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
	Start *int32 `json:"start,omitempty"`
	UsePreset *bool `json:"usePreset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UiSettingsDefaultTimeRange UiSettingsDefaultTimeRange

// NewUiSettingsDefaultTimeRange instantiates a new UiSettingsDefaultTimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiSettingsDefaultTimeRange() *UiSettingsDefaultTimeRange {
	this := UiSettingsDefaultTimeRange{}
	return &this
}

// NewUiSettingsDefaultTimeRangeWithDefaults instantiates a new UiSettingsDefaultTimeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiSettingsDefaultTimeRangeWithDefaults() *UiSettingsDefaultTimeRange {
	this := UiSettingsDefaultTimeRange{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *UiSettingsDefaultTimeRange) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsDefaultTimeRange) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *UiSettingsDefaultTimeRange) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *UiSettingsDefaultTimeRange) SetEnd(v int32) {
	o.End = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *UiSettingsDefaultTimeRange) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsDefaultTimeRange) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *UiSettingsDefaultTimeRange) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *UiSettingsDefaultTimeRange) SetEndDate(v string) {
	o.EndDate = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *UiSettingsDefaultTimeRange) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsDefaultTimeRange) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *UiSettingsDefaultTimeRange) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *UiSettingsDefaultTimeRange) SetInterval(v string) {
	o.Interval = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UiSettingsDefaultTimeRange) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsDefaultTimeRange) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UiSettingsDefaultTimeRange) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UiSettingsDefaultTimeRange) SetName(v string) {
	o.Name = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *UiSettingsDefaultTimeRange) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsDefaultTimeRange) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *UiSettingsDefaultTimeRange) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *UiSettingsDefaultTimeRange) SetShortName(v string) {
	o.ShortName = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *UiSettingsDefaultTimeRange) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsDefaultTimeRange) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *UiSettingsDefaultTimeRange) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *UiSettingsDefaultTimeRange) SetStart(v int32) {
	o.Start = &v
}

// GetUsePreset returns the UsePreset field value if set, zero value otherwise.
func (o *UiSettingsDefaultTimeRange) GetUsePreset() bool {
	if o == nil || IsNil(o.UsePreset) {
		var ret bool
		return ret
	}
	return *o.UsePreset
}

// GetUsePresetOk returns a tuple with the UsePreset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsDefaultTimeRange) GetUsePresetOk() (*bool, bool) {
	if o == nil || IsNil(o.UsePreset) {
		return nil, false
	}
	return o.UsePreset, true
}

// HasUsePreset returns a boolean if a field has been set.
func (o *UiSettingsDefaultTimeRange) HasUsePreset() bool {
	if o != nil && !IsNil(o.UsePreset) {
		return true
	}

	return false
}

// SetUsePreset gets a reference to the given bool and assigns it to the UsePreset field.
func (o *UiSettingsDefaultTimeRange) SetUsePreset(v bool) {
	o.UsePreset = &v
}

func (o UiSettingsDefaultTimeRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiSettingsDefaultTimeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.UsePreset) {
		toSerialize["usePreset"] = o.UsePreset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UiSettingsDefaultTimeRange) UnmarshalJSON(data []byte) (err error) {
	varUiSettingsDefaultTimeRange := _UiSettingsDefaultTimeRange{}

	err = json.Unmarshal(data, &varUiSettingsDefaultTimeRange)

	if err != nil {
		return err
	}

	*o = UiSettingsDefaultTimeRange(varUiSettingsDefaultTimeRange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "name")
		delete(additionalProperties, "shortName")
		delete(additionalProperties, "start")
		delete(additionalProperties, "usePreset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUiSettingsDefaultTimeRange struct {
	value *UiSettingsDefaultTimeRange
	isSet bool
}

func (v NullableUiSettingsDefaultTimeRange) Get() *UiSettingsDefaultTimeRange {
	return v.value
}

func (v *NullableUiSettingsDefaultTimeRange) Set(val *UiSettingsDefaultTimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableUiSettingsDefaultTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableUiSettingsDefaultTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiSettingsDefaultTimeRange(val *UiSettingsDefaultTimeRange) *NullableUiSettingsDefaultTimeRange {
	return &NullableUiSettingsDefaultTimeRange{value: val, isSet: true}
}

func (v NullableUiSettingsDefaultTimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiSettingsDefaultTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

