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

// checks if the JunosEvpnOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JunosEvpnOptions{}

// JunosEvpnOptions EVPN Options
type JunosEvpnOptions struct {
	// optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server-id-overrides
	AutoLoopbackSubnet *string `json:"auto_loopback_subnet,omitempty"`
	// optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server-id-overrides
	AutoLoopbackSubnet6 *string `json:"auto_loopback_subnet6,omitempty"`
	// optional, this generates router_id automatically, if specified, `router_id_prefix` is ignored
	AutoRouterIdSubnet *string `json:"auto_router_id_subnet,omitempty"`
	// optional, this generates router_id automatically, if specified, `router_id_prefix` is ignored
	AutoRouterIdSubnet6 *string `json:"auto_router_id_subnet6,omitempty"`
	// optional, for ERB or CLOS, you can either use esilag to upstream routers or to also be the virtual-gateway when `routed_at` != `core`, whether to do virtual-gateway at core as well
	CoreAsBorder *bool `json:"core_as_border,omitempty"`
	Overlay *JunosEvpnOptionsOverlay `json:"overlay,omitempty"`
	// by default, JUNOS uses 00-00-5e-00-01-01 as the virtual-gateway-address's v4-mac if enabled, 00-00-5e-00-XX-YY will be used (where XX=vlan_id/256, YY=vlan_id%256)
	PerVlanVgaV4Mac *bool `json:"per_vlan_vga_v4_mac,omitempty"`
	RoutedAt *JunosEvpnOptionsRoutedAt `json:"routed_at,omitempty"`
	Underlay *JunosEvpnOptionsUnderlay `json:"underlay,omitempty"`
	// optional, for EX9200 only to seggregate virtual-switches
	VsInstances *map[string]JunosEvpnOptionsVsInstancesValue `json:"vs_instances,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JunosEvpnOptions JunosEvpnOptions

// NewJunosEvpnOptions instantiates a new JunosEvpnOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJunosEvpnOptions() *JunosEvpnOptions {
	this := JunosEvpnOptions{}
	var autoRouterIdSubnet6 string = "fd31:5700:1::/64"
	this.AutoRouterIdSubnet6 = &autoRouterIdSubnet6
	var coreAsBorder bool = false
	this.CoreAsBorder = &coreAsBorder
	var perVlanVgaV4Mac bool = false
	this.PerVlanVgaV4Mac = &perVlanVgaV4Mac
	var routedAt JunosEvpnOptionsRoutedAt = JUNOSEVPNOPTIONSROUTEDAT_EDGE
	this.RoutedAt = &routedAt
	return &this
}

// NewJunosEvpnOptionsWithDefaults instantiates a new JunosEvpnOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJunosEvpnOptionsWithDefaults() *JunosEvpnOptions {
	this := JunosEvpnOptions{}
	var autoRouterIdSubnet6 string = "fd31:5700:1::/64"
	this.AutoRouterIdSubnet6 = &autoRouterIdSubnet6
	var coreAsBorder bool = false
	this.CoreAsBorder = &coreAsBorder
	var perVlanVgaV4Mac bool = false
	this.PerVlanVgaV4Mac = &perVlanVgaV4Mac
	var routedAt JunosEvpnOptionsRoutedAt = JUNOSEVPNOPTIONSROUTEDAT_EDGE
	this.RoutedAt = &routedAt
	return &this
}

// GetAutoLoopbackSubnet returns the AutoLoopbackSubnet field value if set, zero value otherwise.
func (o *JunosEvpnOptions) GetAutoLoopbackSubnet() string {
	if o == nil || IsNil(o.AutoLoopbackSubnet) {
		var ret string
		return ret
	}
	return *o.AutoLoopbackSubnet
}

// GetAutoLoopbackSubnetOk returns a tuple with the AutoLoopbackSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnOptions) GetAutoLoopbackSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.AutoLoopbackSubnet) {
		return nil, false
	}
	return o.AutoLoopbackSubnet, true
}

// HasAutoLoopbackSubnet returns a boolean if a field has been set.
func (o *JunosEvpnOptions) HasAutoLoopbackSubnet() bool {
	if o != nil && !IsNil(o.AutoLoopbackSubnet) {
		return true
	}

	return false
}

// SetAutoLoopbackSubnet gets a reference to the given string and assigns it to the AutoLoopbackSubnet field.
func (o *JunosEvpnOptions) SetAutoLoopbackSubnet(v string) {
	o.AutoLoopbackSubnet = &v
}

// GetAutoLoopbackSubnet6 returns the AutoLoopbackSubnet6 field value if set, zero value otherwise.
func (o *JunosEvpnOptions) GetAutoLoopbackSubnet6() string {
	if o == nil || IsNil(o.AutoLoopbackSubnet6) {
		var ret string
		return ret
	}
	return *o.AutoLoopbackSubnet6
}

// GetAutoLoopbackSubnet6Ok returns a tuple with the AutoLoopbackSubnet6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnOptions) GetAutoLoopbackSubnet6Ok() (*string, bool) {
	if o == nil || IsNil(o.AutoLoopbackSubnet6) {
		return nil, false
	}
	return o.AutoLoopbackSubnet6, true
}

// HasAutoLoopbackSubnet6 returns a boolean if a field has been set.
func (o *JunosEvpnOptions) HasAutoLoopbackSubnet6() bool {
	if o != nil && !IsNil(o.AutoLoopbackSubnet6) {
		return true
	}

	return false
}

// SetAutoLoopbackSubnet6 gets a reference to the given string and assigns it to the AutoLoopbackSubnet6 field.
func (o *JunosEvpnOptions) SetAutoLoopbackSubnet6(v string) {
	o.AutoLoopbackSubnet6 = &v
}

// GetAutoRouterIdSubnet returns the AutoRouterIdSubnet field value if set, zero value otherwise.
func (o *JunosEvpnOptions) GetAutoRouterIdSubnet() string {
	if o == nil || IsNil(o.AutoRouterIdSubnet) {
		var ret string
		return ret
	}
	return *o.AutoRouterIdSubnet
}

// GetAutoRouterIdSubnetOk returns a tuple with the AutoRouterIdSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnOptions) GetAutoRouterIdSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.AutoRouterIdSubnet) {
		return nil, false
	}
	return o.AutoRouterIdSubnet, true
}

// HasAutoRouterIdSubnet returns a boolean if a field has been set.
func (o *JunosEvpnOptions) HasAutoRouterIdSubnet() bool {
	if o != nil && !IsNil(o.AutoRouterIdSubnet) {
		return true
	}

	return false
}

// SetAutoRouterIdSubnet gets a reference to the given string and assigns it to the AutoRouterIdSubnet field.
func (o *JunosEvpnOptions) SetAutoRouterIdSubnet(v string) {
	o.AutoRouterIdSubnet = &v
}

// GetAutoRouterIdSubnet6 returns the AutoRouterIdSubnet6 field value if set, zero value otherwise.
func (o *JunosEvpnOptions) GetAutoRouterIdSubnet6() string {
	if o == nil || IsNil(o.AutoRouterIdSubnet6) {
		var ret string
		return ret
	}
	return *o.AutoRouterIdSubnet6
}

// GetAutoRouterIdSubnet6Ok returns a tuple with the AutoRouterIdSubnet6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnOptions) GetAutoRouterIdSubnet6Ok() (*string, bool) {
	if o == nil || IsNil(o.AutoRouterIdSubnet6) {
		return nil, false
	}
	return o.AutoRouterIdSubnet6, true
}

// HasAutoRouterIdSubnet6 returns a boolean if a field has been set.
func (o *JunosEvpnOptions) HasAutoRouterIdSubnet6() bool {
	if o != nil && !IsNil(o.AutoRouterIdSubnet6) {
		return true
	}

	return false
}

// SetAutoRouterIdSubnet6 gets a reference to the given string and assigns it to the AutoRouterIdSubnet6 field.
func (o *JunosEvpnOptions) SetAutoRouterIdSubnet6(v string) {
	o.AutoRouterIdSubnet6 = &v
}

// GetCoreAsBorder returns the CoreAsBorder field value if set, zero value otherwise.
func (o *JunosEvpnOptions) GetCoreAsBorder() bool {
	if o == nil || IsNil(o.CoreAsBorder) {
		var ret bool
		return ret
	}
	return *o.CoreAsBorder
}

// GetCoreAsBorderOk returns a tuple with the CoreAsBorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnOptions) GetCoreAsBorderOk() (*bool, bool) {
	if o == nil || IsNil(o.CoreAsBorder) {
		return nil, false
	}
	return o.CoreAsBorder, true
}

// HasCoreAsBorder returns a boolean if a field has been set.
func (o *JunosEvpnOptions) HasCoreAsBorder() bool {
	if o != nil && !IsNil(o.CoreAsBorder) {
		return true
	}

	return false
}

// SetCoreAsBorder gets a reference to the given bool and assigns it to the CoreAsBorder field.
func (o *JunosEvpnOptions) SetCoreAsBorder(v bool) {
	o.CoreAsBorder = &v
}

// GetOverlay returns the Overlay field value if set, zero value otherwise.
func (o *JunosEvpnOptions) GetOverlay() JunosEvpnOptionsOverlay {
	if o == nil || IsNil(o.Overlay) {
		var ret JunosEvpnOptionsOverlay
		return ret
	}
	return *o.Overlay
}

// GetOverlayOk returns a tuple with the Overlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnOptions) GetOverlayOk() (*JunosEvpnOptionsOverlay, bool) {
	if o == nil || IsNil(o.Overlay) {
		return nil, false
	}
	return o.Overlay, true
}

// HasOverlay returns a boolean if a field has been set.
func (o *JunosEvpnOptions) HasOverlay() bool {
	if o != nil && !IsNil(o.Overlay) {
		return true
	}

	return false
}

// SetOverlay gets a reference to the given JunosEvpnOptionsOverlay and assigns it to the Overlay field.
func (o *JunosEvpnOptions) SetOverlay(v JunosEvpnOptionsOverlay) {
	o.Overlay = &v
}

// GetPerVlanVgaV4Mac returns the PerVlanVgaV4Mac field value if set, zero value otherwise.
func (o *JunosEvpnOptions) GetPerVlanVgaV4Mac() bool {
	if o == nil || IsNil(o.PerVlanVgaV4Mac) {
		var ret bool
		return ret
	}
	return *o.PerVlanVgaV4Mac
}

// GetPerVlanVgaV4MacOk returns a tuple with the PerVlanVgaV4Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnOptions) GetPerVlanVgaV4MacOk() (*bool, bool) {
	if o == nil || IsNil(o.PerVlanVgaV4Mac) {
		return nil, false
	}
	return o.PerVlanVgaV4Mac, true
}

// HasPerVlanVgaV4Mac returns a boolean if a field has been set.
func (o *JunosEvpnOptions) HasPerVlanVgaV4Mac() bool {
	if o != nil && !IsNil(o.PerVlanVgaV4Mac) {
		return true
	}

	return false
}

// SetPerVlanVgaV4Mac gets a reference to the given bool and assigns it to the PerVlanVgaV4Mac field.
func (o *JunosEvpnOptions) SetPerVlanVgaV4Mac(v bool) {
	o.PerVlanVgaV4Mac = &v
}

// GetRoutedAt returns the RoutedAt field value if set, zero value otherwise.
func (o *JunosEvpnOptions) GetRoutedAt() JunosEvpnOptionsRoutedAt {
	if o == nil || IsNil(o.RoutedAt) {
		var ret JunosEvpnOptionsRoutedAt
		return ret
	}
	return *o.RoutedAt
}

// GetRoutedAtOk returns a tuple with the RoutedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnOptions) GetRoutedAtOk() (*JunosEvpnOptionsRoutedAt, bool) {
	if o == nil || IsNil(o.RoutedAt) {
		return nil, false
	}
	return o.RoutedAt, true
}

// HasRoutedAt returns a boolean if a field has been set.
func (o *JunosEvpnOptions) HasRoutedAt() bool {
	if o != nil && !IsNil(o.RoutedAt) {
		return true
	}

	return false
}

// SetRoutedAt gets a reference to the given JunosEvpnOptionsRoutedAt and assigns it to the RoutedAt field.
func (o *JunosEvpnOptions) SetRoutedAt(v JunosEvpnOptionsRoutedAt) {
	o.RoutedAt = &v
}

// GetUnderlay returns the Underlay field value if set, zero value otherwise.
func (o *JunosEvpnOptions) GetUnderlay() JunosEvpnOptionsUnderlay {
	if o == nil || IsNil(o.Underlay) {
		var ret JunosEvpnOptionsUnderlay
		return ret
	}
	return *o.Underlay
}

// GetUnderlayOk returns a tuple with the Underlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnOptions) GetUnderlayOk() (*JunosEvpnOptionsUnderlay, bool) {
	if o == nil || IsNil(o.Underlay) {
		return nil, false
	}
	return o.Underlay, true
}

// HasUnderlay returns a boolean if a field has been set.
func (o *JunosEvpnOptions) HasUnderlay() bool {
	if o != nil && !IsNil(o.Underlay) {
		return true
	}

	return false
}

// SetUnderlay gets a reference to the given JunosEvpnOptionsUnderlay and assigns it to the Underlay field.
func (o *JunosEvpnOptions) SetUnderlay(v JunosEvpnOptionsUnderlay) {
	o.Underlay = &v
}

// GetVsInstances returns the VsInstances field value if set, zero value otherwise.
func (o *JunosEvpnOptions) GetVsInstances() map[string]JunosEvpnOptionsVsInstancesValue {
	if o == nil || IsNil(o.VsInstances) {
		var ret map[string]JunosEvpnOptionsVsInstancesValue
		return ret
	}
	return *o.VsInstances
}

// GetVsInstancesOk returns a tuple with the VsInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnOptions) GetVsInstancesOk() (*map[string]JunosEvpnOptionsVsInstancesValue, bool) {
	if o == nil || IsNil(o.VsInstances) {
		return nil, false
	}
	return o.VsInstances, true
}

// HasVsInstances returns a boolean if a field has been set.
func (o *JunosEvpnOptions) HasVsInstances() bool {
	if o != nil && !IsNil(o.VsInstances) {
		return true
	}

	return false
}

// SetVsInstances gets a reference to the given map[string]JunosEvpnOptionsVsInstancesValue and assigns it to the VsInstances field.
func (o *JunosEvpnOptions) SetVsInstances(v map[string]JunosEvpnOptionsVsInstancesValue) {
	o.VsInstances = &v
}

func (o JunosEvpnOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JunosEvpnOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoLoopbackSubnet) {
		toSerialize["auto_loopback_subnet"] = o.AutoLoopbackSubnet
	}
	if !IsNil(o.AutoLoopbackSubnet6) {
		toSerialize["auto_loopback_subnet6"] = o.AutoLoopbackSubnet6
	}
	if !IsNil(o.AutoRouterIdSubnet) {
		toSerialize["auto_router_id_subnet"] = o.AutoRouterIdSubnet
	}
	if !IsNil(o.AutoRouterIdSubnet6) {
		toSerialize["auto_router_id_subnet6"] = o.AutoRouterIdSubnet6
	}
	if !IsNil(o.CoreAsBorder) {
		toSerialize["core_as_border"] = o.CoreAsBorder
	}
	if !IsNil(o.Overlay) {
		toSerialize["overlay"] = o.Overlay
	}
	if !IsNil(o.PerVlanVgaV4Mac) {
		toSerialize["per_vlan_vga_v4_mac"] = o.PerVlanVgaV4Mac
	}
	if !IsNil(o.RoutedAt) {
		toSerialize["routed_at"] = o.RoutedAt
	}
	if !IsNil(o.Underlay) {
		toSerialize["underlay"] = o.Underlay
	}
	if !IsNil(o.VsInstances) {
		toSerialize["vs_instances"] = o.VsInstances
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JunosEvpnOptions) UnmarshalJSON(data []byte) (err error) {
	varJunosEvpnOptions := _JunosEvpnOptions{}

	err = json.Unmarshal(data, &varJunosEvpnOptions)

	if err != nil {
		return err
	}

	*o = JunosEvpnOptions(varJunosEvpnOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auto_loopback_subnet")
		delete(additionalProperties, "auto_loopback_subnet6")
		delete(additionalProperties, "auto_router_id_subnet")
		delete(additionalProperties, "auto_router_id_subnet6")
		delete(additionalProperties, "core_as_border")
		delete(additionalProperties, "overlay")
		delete(additionalProperties, "per_vlan_vga_v4_mac")
		delete(additionalProperties, "routed_at")
		delete(additionalProperties, "underlay")
		delete(additionalProperties, "vs_instances")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJunosEvpnOptions struct {
	value *JunosEvpnOptions
	isSet bool
}

func (v NullableJunosEvpnOptions) Get() *JunosEvpnOptions {
	return v.value
}

func (v *NullableJunosEvpnOptions) Set(val *JunosEvpnOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableJunosEvpnOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableJunosEvpnOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJunosEvpnOptions(val *JunosEvpnOptions) *NullableJunosEvpnOptions {
	return &NullableJunosEvpnOptions{value: val, isSet: true}
}

func (v NullableJunosEvpnOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJunosEvpnOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

