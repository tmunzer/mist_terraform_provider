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

// checks if the GatewayRoutingPolicyItemMatchingVpnPathSla type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayRoutingPolicyItemMatchingVpnPathSla{}

// GatewayRoutingPolicyItemMatchingVpnPathSla struct for GatewayRoutingPolicyItemMatchingVpnPathSla
type GatewayRoutingPolicyItemMatchingVpnPathSla struct {
	MaxJitter NullableInt32 `json:"max_jitter,omitempty"`
	MaxLatency NullableInt32 `json:"max_latency,omitempty"`
	MaxLoss NullableInt32 `json:"max_loss,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayRoutingPolicyItemMatchingVpnPathSla GatewayRoutingPolicyItemMatchingVpnPathSla

// NewGatewayRoutingPolicyItemMatchingVpnPathSla instantiates a new GatewayRoutingPolicyItemMatchingVpnPathSla object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayRoutingPolicyItemMatchingVpnPathSla() *GatewayRoutingPolicyItemMatchingVpnPathSla {
	this := GatewayRoutingPolicyItemMatchingVpnPathSla{}
	return &this
}

// NewGatewayRoutingPolicyItemMatchingVpnPathSlaWithDefaults instantiates a new GatewayRoutingPolicyItemMatchingVpnPathSla object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayRoutingPolicyItemMatchingVpnPathSlaWithDefaults() *GatewayRoutingPolicyItemMatchingVpnPathSla {
	this := GatewayRoutingPolicyItemMatchingVpnPathSla{}
	return &this
}

// GetMaxJitter returns the MaxJitter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) GetMaxJitter() int32 {
	if o == nil || IsNil(o.MaxJitter.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxJitter.Get()
}

// GetMaxJitterOk returns a tuple with the MaxJitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) GetMaxJitterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxJitter.Get(), o.MaxJitter.IsSet()
}

// HasMaxJitter returns a boolean if a field has been set.
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) HasMaxJitter() bool {
	if o != nil && o.MaxJitter.IsSet() {
		return true
	}

	return false
}

// SetMaxJitter gets a reference to the given NullableInt32 and assigns it to the MaxJitter field.
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) SetMaxJitter(v int32) {
	o.MaxJitter.Set(&v)
}
// SetMaxJitterNil sets the value for MaxJitter to be an explicit nil
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) SetMaxJitterNil() {
	o.MaxJitter.Set(nil)
}

// UnsetMaxJitter ensures that no value is present for MaxJitter, not even an explicit nil
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) UnsetMaxJitter() {
	o.MaxJitter.Unset()
}

// GetMaxLatency returns the MaxLatency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) GetMaxLatency() int32 {
	if o == nil || IsNil(o.MaxLatency.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxLatency.Get()
}

// GetMaxLatencyOk returns a tuple with the MaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) GetMaxLatencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLatency.Get(), o.MaxLatency.IsSet()
}

// HasMaxLatency returns a boolean if a field has been set.
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) HasMaxLatency() bool {
	if o != nil && o.MaxLatency.IsSet() {
		return true
	}

	return false
}

// SetMaxLatency gets a reference to the given NullableInt32 and assigns it to the MaxLatency field.
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) SetMaxLatency(v int32) {
	o.MaxLatency.Set(&v)
}
// SetMaxLatencyNil sets the value for MaxLatency to be an explicit nil
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) SetMaxLatencyNil() {
	o.MaxLatency.Set(nil)
}

// UnsetMaxLatency ensures that no value is present for MaxLatency, not even an explicit nil
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) UnsetMaxLatency() {
	o.MaxLatency.Unset()
}

// GetMaxLoss returns the MaxLoss field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) GetMaxLoss() int32 {
	if o == nil || IsNil(o.MaxLoss.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxLoss.Get()
}

// GetMaxLossOk returns a tuple with the MaxLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) GetMaxLossOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLoss.Get(), o.MaxLoss.IsSet()
}

// HasMaxLoss returns a boolean if a field has been set.
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) HasMaxLoss() bool {
	if o != nil && o.MaxLoss.IsSet() {
		return true
	}

	return false
}

// SetMaxLoss gets a reference to the given NullableInt32 and assigns it to the MaxLoss field.
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) SetMaxLoss(v int32) {
	o.MaxLoss.Set(&v)
}
// SetMaxLossNil sets the value for MaxLoss to be an explicit nil
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) SetMaxLossNil() {
	o.MaxLoss.Set(nil)
}

// UnsetMaxLoss ensures that no value is present for MaxLoss, not even an explicit nil
func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) UnsetMaxLoss() {
	o.MaxLoss.Unset()
}

func (o GatewayRoutingPolicyItemMatchingVpnPathSla) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayRoutingPolicyItemMatchingVpnPathSla) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxJitter.IsSet() {
		toSerialize["max_jitter"] = o.MaxJitter.Get()
	}
	if o.MaxLatency.IsSet() {
		toSerialize["max_latency"] = o.MaxLatency.Get()
	}
	if o.MaxLoss.IsSet() {
		toSerialize["max_loss"] = o.MaxLoss.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayRoutingPolicyItemMatchingVpnPathSla) UnmarshalJSON(data []byte) (err error) {
	varGatewayRoutingPolicyItemMatchingVpnPathSla := _GatewayRoutingPolicyItemMatchingVpnPathSla{}

	err = json.Unmarshal(data, &varGatewayRoutingPolicyItemMatchingVpnPathSla)

	if err != nil {
		return err
	}

	*o = GatewayRoutingPolicyItemMatchingVpnPathSla(varGatewayRoutingPolicyItemMatchingVpnPathSla)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "max_jitter")
		delete(additionalProperties, "max_latency")
		delete(additionalProperties, "max_loss")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayRoutingPolicyItemMatchingVpnPathSla struct {
	value *GatewayRoutingPolicyItemMatchingVpnPathSla
	isSet bool
}

func (v NullableGatewayRoutingPolicyItemMatchingVpnPathSla) Get() *GatewayRoutingPolicyItemMatchingVpnPathSla {
	return v.value
}

func (v *NullableGatewayRoutingPolicyItemMatchingVpnPathSla) Set(val *GatewayRoutingPolicyItemMatchingVpnPathSla) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayRoutingPolicyItemMatchingVpnPathSla) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayRoutingPolicyItemMatchingVpnPathSla) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayRoutingPolicyItemMatchingVpnPathSla(val *GatewayRoutingPolicyItemMatchingVpnPathSla) *NullableGatewayRoutingPolicyItemMatchingVpnPathSla {
	return &NullableGatewayRoutingPolicyItemMatchingVpnPathSla{value: val, isSet: true}
}

func (v NullableGatewayRoutingPolicyItemMatchingVpnPathSla) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayRoutingPolicyItemMatchingVpnPathSla) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

