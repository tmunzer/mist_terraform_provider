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

// checks if the ConstTrafficType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstTrafficType{}

// ConstTrafficType struct for ConstTrafficType
type ConstTrafficType struct {
	Display *string `json:"display,omitempty"`
	Dscp *int32 `json:"dscp,omitempty"`
	FailoverPolicy *string `json:"failover_policy,omitempty"`
	MaxJitter *int32 `json:"max_jitter,omitempty"`
	MaxLatency *int32 `json:"max_latency,omitempty"`
	MaxLoss *int32 `json:"max_loss,omitempty"`
	Name *string `json:"name,omitempty"`
	TrafficClass *string `json:"traffic_class,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstTrafficType ConstTrafficType

// NewConstTrafficType instantiates a new ConstTrafficType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstTrafficType() *ConstTrafficType {
	this := ConstTrafficType{}
	return &this
}

// NewConstTrafficTypeWithDefaults instantiates a new ConstTrafficType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstTrafficTypeWithDefaults() *ConstTrafficType {
	this := ConstTrafficType{}
	return &this
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ConstTrafficType) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstTrafficType) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ConstTrafficType) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *ConstTrafficType) SetDisplay(v string) {
	o.Display = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *ConstTrafficType) GetDscp() int32 {
	if o == nil || IsNil(o.Dscp) {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstTrafficType) GetDscpOk() (*int32, bool) {
	if o == nil || IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *ConstTrafficType) HasDscp() bool {
	if o != nil && !IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *ConstTrafficType) SetDscp(v int32) {
	o.Dscp = &v
}

// GetFailoverPolicy returns the FailoverPolicy field value if set, zero value otherwise.
func (o *ConstTrafficType) GetFailoverPolicy() string {
	if o == nil || IsNil(o.FailoverPolicy) {
		var ret string
		return ret
	}
	return *o.FailoverPolicy
}

// GetFailoverPolicyOk returns a tuple with the FailoverPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstTrafficType) GetFailoverPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.FailoverPolicy) {
		return nil, false
	}
	return o.FailoverPolicy, true
}

// HasFailoverPolicy returns a boolean if a field has been set.
func (o *ConstTrafficType) HasFailoverPolicy() bool {
	if o != nil && !IsNil(o.FailoverPolicy) {
		return true
	}

	return false
}

// SetFailoverPolicy gets a reference to the given string and assigns it to the FailoverPolicy field.
func (o *ConstTrafficType) SetFailoverPolicy(v string) {
	o.FailoverPolicy = &v
}

// GetMaxJitter returns the MaxJitter field value if set, zero value otherwise.
func (o *ConstTrafficType) GetMaxJitter() int32 {
	if o == nil || IsNil(o.MaxJitter) {
		var ret int32
		return ret
	}
	return *o.MaxJitter
}

// GetMaxJitterOk returns a tuple with the MaxJitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstTrafficType) GetMaxJitterOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxJitter) {
		return nil, false
	}
	return o.MaxJitter, true
}

// HasMaxJitter returns a boolean if a field has been set.
func (o *ConstTrafficType) HasMaxJitter() bool {
	if o != nil && !IsNil(o.MaxJitter) {
		return true
	}

	return false
}

// SetMaxJitter gets a reference to the given int32 and assigns it to the MaxJitter field.
func (o *ConstTrafficType) SetMaxJitter(v int32) {
	o.MaxJitter = &v
}

// GetMaxLatency returns the MaxLatency field value if set, zero value otherwise.
func (o *ConstTrafficType) GetMaxLatency() int32 {
	if o == nil || IsNil(o.MaxLatency) {
		var ret int32
		return ret
	}
	return *o.MaxLatency
}

// GetMaxLatencyOk returns a tuple with the MaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstTrafficType) GetMaxLatencyOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLatency) {
		return nil, false
	}
	return o.MaxLatency, true
}

// HasMaxLatency returns a boolean if a field has been set.
func (o *ConstTrafficType) HasMaxLatency() bool {
	if o != nil && !IsNil(o.MaxLatency) {
		return true
	}

	return false
}

// SetMaxLatency gets a reference to the given int32 and assigns it to the MaxLatency field.
func (o *ConstTrafficType) SetMaxLatency(v int32) {
	o.MaxLatency = &v
}

// GetMaxLoss returns the MaxLoss field value if set, zero value otherwise.
func (o *ConstTrafficType) GetMaxLoss() int32 {
	if o == nil || IsNil(o.MaxLoss) {
		var ret int32
		return ret
	}
	return *o.MaxLoss
}

// GetMaxLossOk returns a tuple with the MaxLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstTrafficType) GetMaxLossOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLoss) {
		return nil, false
	}
	return o.MaxLoss, true
}

// HasMaxLoss returns a boolean if a field has been set.
func (o *ConstTrafficType) HasMaxLoss() bool {
	if o != nil && !IsNil(o.MaxLoss) {
		return true
	}

	return false
}

// SetMaxLoss gets a reference to the given int32 and assigns it to the MaxLoss field.
func (o *ConstTrafficType) SetMaxLoss(v int32) {
	o.MaxLoss = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConstTrafficType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstTrafficType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConstTrafficType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConstTrafficType) SetName(v string) {
	o.Name = &v
}

// GetTrafficClass returns the TrafficClass field value if set, zero value otherwise.
func (o *ConstTrafficType) GetTrafficClass() string {
	if o == nil || IsNil(o.TrafficClass) {
		var ret string
		return ret
	}
	return *o.TrafficClass
}

// GetTrafficClassOk returns a tuple with the TrafficClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstTrafficType) GetTrafficClassOk() (*string, bool) {
	if o == nil || IsNil(o.TrafficClass) {
		return nil, false
	}
	return o.TrafficClass, true
}

// HasTrafficClass returns a boolean if a field has been set.
func (o *ConstTrafficType) HasTrafficClass() bool {
	if o != nil && !IsNil(o.TrafficClass) {
		return true
	}

	return false
}

// SetTrafficClass gets a reference to the given string and assigns it to the TrafficClass field.
func (o *ConstTrafficType) SetTrafficClass(v string) {
	o.TrafficClass = &v
}

func (o ConstTrafficType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstTrafficType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	if !IsNil(o.FailoverPolicy) {
		toSerialize["failover_policy"] = o.FailoverPolicy
	}
	if !IsNil(o.MaxJitter) {
		toSerialize["max_jitter"] = o.MaxJitter
	}
	if !IsNil(o.MaxLatency) {
		toSerialize["max_latency"] = o.MaxLatency
	}
	if !IsNil(o.MaxLoss) {
		toSerialize["max_loss"] = o.MaxLoss
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TrafficClass) {
		toSerialize["traffic_class"] = o.TrafficClass
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstTrafficType) UnmarshalJSON(data []byte) (err error) {
	varConstTrafficType := _ConstTrafficType{}

	err = json.Unmarshal(data, &varConstTrafficType)

	if err != nil {
		return err
	}

	*o = ConstTrafficType(varConstTrafficType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display")
		delete(additionalProperties, "dscp")
		delete(additionalProperties, "failover_policy")
		delete(additionalProperties, "max_jitter")
		delete(additionalProperties, "max_latency")
		delete(additionalProperties, "max_loss")
		delete(additionalProperties, "name")
		delete(additionalProperties, "traffic_class")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstTrafficType struct {
	value *ConstTrafficType
	isSet bool
}

func (v NullableConstTrafficType) Get() *ConstTrafficType {
	return v.value
}

func (v *NullableConstTrafficType) Set(val *ConstTrafficType) {
	v.value = val
	v.isSet = true
}

func (v NullableConstTrafficType) IsSet() bool {
	return v.isSet
}

func (v *NullableConstTrafficType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstTrafficType(val *ConstTrafficType) *NullableConstTrafficType {
	return &NullableConstTrafficType{value: val, isSet: true}
}

func (v NullableConstTrafficType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstTrafficType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

