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

// checks if the SleSummarySle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleSummarySle{}

// SleSummarySle struct for SleSummarySle
type SleSummarySle struct {
	Interval float32 `json:"interval"`
	Name string `json:"name"`
	Samples SleSummarySleSamples `json:"samples"`
	XLabel string `json:"x_label"`
	YLabel string `json:"y_label"`
	AdditionalProperties map[string]interface{}
}

type _SleSummarySle SleSummarySle

// NewSleSummarySle instantiates a new SleSummarySle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleSummarySle(interval float32, name string, samples SleSummarySleSamples, xLabel string, yLabel string) *SleSummarySle {
	this := SleSummarySle{}
	this.Interval = interval
	this.Name = name
	this.Samples = samples
	this.XLabel = xLabel
	this.YLabel = yLabel
	return &this
}

// NewSleSummarySleWithDefaults instantiates a new SleSummarySle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleSummarySleWithDefaults() *SleSummarySle {
	this := SleSummarySle{}
	return &this
}

// GetInterval returns the Interval field value
func (o *SleSummarySle) GetInterval() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *SleSummarySle) GetIntervalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *SleSummarySle) SetInterval(v float32) {
	o.Interval = v
}

// GetName returns the Name field value
func (o *SleSummarySle) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SleSummarySle) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SleSummarySle) SetName(v string) {
	o.Name = v
}

// GetSamples returns the Samples field value
func (o *SleSummarySle) GetSamples() SleSummarySleSamples {
	if o == nil {
		var ret SleSummarySleSamples
		return ret
	}

	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value
// and a boolean to check if the value has been set.
func (o *SleSummarySle) GetSamplesOk() (*SleSummarySleSamples, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Samples, true
}

// SetSamples sets field value
func (o *SleSummarySle) SetSamples(v SleSummarySleSamples) {
	o.Samples = v
}

// GetXLabel returns the XLabel field value
func (o *SleSummarySle) GetXLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.XLabel
}

// GetXLabelOk returns a tuple with the XLabel field value
// and a boolean to check if the value has been set.
func (o *SleSummarySle) GetXLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XLabel, true
}

// SetXLabel sets field value
func (o *SleSummarySle) SetXLabel(v string) {
	o.XLabel = v
}

// GetYLabel returns the YLabel field value
func (o *SleSummarySle) GetYLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YLabel
}

// GetYLabelOk returns a tuple with the YLabel field value
// and a boolean to check if the value has been set.
func (o *SleSummarySle) GetYLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YLabel, true
}

// SetYLabel sets field value
func (o *SleSummarySle) SetYLabel(v string) {
	o.YLabel = v
}

func (o SleSummarySle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleSummarySle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interval"] = o.Interval
	toSerialize["name"] = o.Name
	toSerialize["samples"] = o.Samples
	toSerialize["x_label"] = o.XLabel
	toSerialize["y_label"] = o.YLabel

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleSummarySle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interval",
		"name",
		"samples",
		"x_label",
		"y_label",
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

	varSleSummarySle := _SleSummarySle{}

	err = json.Unmarshal(data, &varSleSummarySle)

	if err != nil {
		return err
	}

	*o = SleSummarySle(varSleSummarySle)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "interval")
		delete(additionalProperties, "name")
		delete(additionalProperties, "samples")
		delete(additionalProperties, "x_label")
		delete(additionalProperties, "y_label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleSummarySle struct {
	value *SleSummarySle
	isSet bool
}

func (v NullableSleSummarySle) Get() *SleSummarySle {
	return v.value
}

func (v *NullableSleSummarySle) Set(val *SleSummarySle) {
	v.value = val
	v.isSet = true
}

func (v NullableSleSummarySle) IsSet() bool {
	return v.isSet
}

func (v *NullableSleSummarySle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleSummarySle(val *SleSummarySle) *NullableSleSummarySle {
	return &NullableSleSummarySle{value: val, isSet: true}
}

func (v NullableSleSummarySle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleSummarySle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

