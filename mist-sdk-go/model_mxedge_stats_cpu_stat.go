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

// checks if the MxedgeStatsCpuStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeStatsCpuStat{}

// MxedgeStatsCpuStat CPU/core stats list
type MxedgeStatsCpuStat struct {
	Cpus *map[string]CpuStat `json:"cpus,omitempty"`
	// percentage of Idle, Idle/(Idle + Busy) since last sampling
	Idle *int32 `json:"idle,omitempty"`
	// percentage of Interrupt, (Irq + SoftIrq)/(Idle + Busy) since last sampling
	Interrupt *int32 `json:"interrupt,omitempty"`
	// percentage of System, System/(Idle + Busy) since last sampling
	System *int32 `json:"system,omitempty"`
	// percentage of load, Busy/(Idle + Busy) since last sampling
	Usage *int32 `json:"usage,omitempty"`
	// percentage of User, User/(Idle + Busy) since last sampling
	User *int32 `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeStatsCpuStat MxedgeStatsCpuStat

// NewMxedgeStatsCpuStat instantiates a new MxedgeStatsCpuStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeStatsCpuStat() *MxedgeStatsCpuStat {
	this := MxedgeStatsCpuStat{}
	return &this
}

// NewMxedgeStatsCpuStatWithDefaults instantiates a new MxedgeStatsCpuStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeStatsCpuStatWithDefaults() *MxedgeStatsCpuStat {
	this := MxedgeStatsCpuStat{}
	return &this
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *MxedgeStatsCpuStat) GetCpus() map[string]CpuStat {
	if o == nil || IsNil(o.Cpus) {
		var ret map[string]CpuStat
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsCpuStat) GetCpusOk() (*map[string]CpuStat, bool) {
	if o == nil || IsNil(o.Cpus) {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *MxedgeStatsCpuStat) HasCpus() bool {
	if o != nil && !IsNil(o.Cpus) {
		return true
	}

	return false
}

// SetCpus gets a reference to the given map[string]CpuStat and assigns it to the Cpus field.
func (o *MxedgeStatsCpuStat) SetCpus(v map[string]CpuStat) {
	o.Cpus = &v
}

// GetIdle returns the Idle field value if set, zero value otherwise.
func (o *MxedgeStatsCpuStat) GetIdle() int32 {
	if o == nil || IsNil(o.Idle) {
		var ret int32
		return ret
	}
	return *o.Idle
}

// GetIdleOk returns a tuple with the Idle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsCpuStat) GetIdleOk() (*int32, bool) {
	if o == nil || IsNil(o.Idle) {
		return nil, false
	}
	return o.Idle, true
}

// HasIdle returns a boolean if a field has been set.
func (o *MxedgeStatsCpuStat) HasIdle() bool {
	if o != nil && !IsNil(o.Idle) {
		return true
	}

	return false
}

// SetIdle gets a reference to the given int32 and assigns it to the Idle field.
func (o *MxedgeStatsCpuStat) SetIdle(v int32) {
	o.Idle = &v
}

// GetInterrupt returns the Interrupt field value if set, zero value otherwise.
func (o *MxedgeStatsCpuStat) GetInterrupt() int32 {
	if o == nil || IsNil(o.Interrupt) {
		var ret int32
		return ret
	}
	return *o.Interrupt
}

// GetInterruptOk returns a tuple with the Interrupt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsCpuStat) GetInterruptOk() (*int32, bool) {
	if o == nil || IsNil(o.Interrupt) {
		return nil, false
	}
	return o.Interrupt, true
}

// HasInterrupt returns a boolean if a field has been set.
func (o *MxedgeStatsCpuStat) HasInterrupt() bool {
	if o != nil && !IsNil(o.Interrupt) {
		return true
	}

	return false
}

// SetInterrupt gets a reference to the given int32 and assigns it to the Interrupt field.
func (o *MxedgeStatsCpuStat) SetInterrupt(v int32) {
	o.Interrupt = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *MxedgeStatsCpuStat) GetSystem() int32 {
	if o == nil || IsNil(o.System) {
		var ret int32
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsCpuStat) GetSystemOk() (*int32, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *MxedgeStatsCpuStat) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given int32 and assigns it to the System field.
func (o *MxedgeStatsCpuStat) SetSystem(v int32) {
	o.System = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *MxedgeStatsCpuStat) GetUsage() int32 {
	if o == nil || IsNil(o.Usage) {
		var ret int32
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsCpuStat) GetUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *MxedgeStatsCpuStat) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int32 and assigns it to the Usage field.
func (o *MxedgeStatsCpuStat) SetUsage(v int32) {
	o.Usage = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MxedgeStatsCpuStat) GetUser() int32 {
	if o == nil || IsNil(o.User) {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsCpuStat) GetUserOk() (*int32, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MxedgeStatsCpuStat) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *MxedgeStatsCpuStat) SetUser(v int32) {
	o.User = &v
}

func (o MxedgeStatsCpuStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeStatsCpuStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpus) {
		toSerialize["cpus"] = o.Cpus
	}
	if !IsNil(o.Idle) {
		toSerialize["idle"] = o.Idle
	}
	if !IsNil(o.Interrupt) {
		toSerialize["interrupt"] = o.Interrupt
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeStatsCpuStat) UnmarshalJSON(data []byte) (err error) {
	varMxedgeStatsCpuStat := _MxedgeStatsCpuStat{}

	err = json.Unmarshal(data, &varMxedgeStatsCpuStat)

	if err != nil {
		return err
	}

	*o = MxedgeStatsCpuStat(varMxedgeStatsCpuStat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cpus")
		delete(additionalProperties, "idle")
		delete(additionalProperties, "interrupt")
		delete(additionalProperties, "system")
		delete(additionalProperties, "usage")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeStatsCpuStat struct {
	value *MxedgeStatsCpuStat
	isSet bool
}

func (v NullableMxedgeStatsCpuStat) Get() *MxedgeStatsCpuStat {
	return v.value
}

func (v *NullableMxedgeStatsCpuStat) Set(val *MxedgeStatsCpuStat) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeStatsCpuStat) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeStatsCpuStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeStatsCpuStat(val *MxedgeStatsCpuStat) *NullableMxedgeStatsCpuStat {
	return &NullableMxedgeStatsCpuStat{value: val, isSet: true}
}

func (v NullableMxedgeStatsCpuStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeStatsCpuStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

