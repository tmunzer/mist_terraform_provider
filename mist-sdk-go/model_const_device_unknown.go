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

// checks if the ConstDeviceUnknown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstDeviceUnknown{}

// ConstDeviceUnknown struct for ConstDeviceUnknown
type ConstDeviceUnknown struct {
	ApType *string `json:"ap_type,omitempty"`
	Description *string `json:"description,omitempty"`
	Display *string `json:"display,omitempty"`
	HasExtio *bool `json:"has_extio,omitempty"`
	HasVble *bool `json:"has_vble,omitempty"`
	HasWifiBand24 *bool `json:"has_wifi_band24,omitempty"`
	HasWifiBand5 *bool `json:"has_wifi_band5,omitempty"`
	Model *string `json:"model,omitempty"`
	Unmanaged *bool `json:"unmanaged,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstDeviceUnknown ConstDeviceUnknown

// NewConstDeviceUnknown instantiates a new ConstDeviceUnknown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstDeviceUnknown() *ConstDeviceUnknown {
	this := ConstDeviceUnknown{}
	return &this
}

// NewConstDeviceUnknownWithDefaults instantiates a new ConstDeviceUnknown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstDeviceUnknownWithDefaults() *ConstDeviceUnknown {
	this := ConstDeviceUnknown{}
	return &this
}

// GetApType returns the ApType field value if set, zero value otherwise.
func (o *ConstDeviceUnknown) GetApType() string {
	if o == nil || IsNil(o.ApType) {
		var ret string
		return ret
	}
	return *o.ApType
}

// GetApTypeOk returns a tuple with the ApType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceUnknown) GetApTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ApType) {
		return nil, false
	}
	return o.ApType, true
}

// HasApType returns a boolean if a field has been set.
func (o *ConstDeviceUnknown) HasApType() bool {
	if o != nil && !IsNil(o.ApType) {
		return true
	}

	return false
}

// SetApType gets a reference to the given string and assigns it to the ApType field.
func (o *ConstDeviceUnknown) SetApType(v string) {
	o.ApType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConstDeviceUnknown) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceUnknown) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConstDeviceUnknown) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConstDeviceUnknown) SetDescription(v string) {
	o.Description = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ConstDeviceUnknown) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceUnknown) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ConstDeviceUnknown) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *ConstDeviceUnknown) SetDisplay(v string) {
	o.Display = &v
}

// GetHasExtio returns the HasExtio field value if set, zero value otherwise.
func (o *ConstDeviceUnknown) GetHasExtio() bool {
	if o == nil || IsNil(o.HasExtio) {
		var ret bool
		return ret
	}
	return *o.HasExtio
}

// GetHasExtioOk returns a tuple with the HasExtio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceUnknown) GetHasExtioOk() (*bool, bool) {
	if o == nil || IsNil(o.HasExtio) {
		return nil, false
	}
	return o.HasExtio, true
}

// HasHasExtio returns a boolean if a field has been set.
func (o *ConstDeviceUnknown) HasHasExtio() bool {
	if o != nil && !IsNil(o.HasExtio) {
		return true
	}

	return false
}

// SetHasExtio gets a reference to the given bool and assigns it to the HasExtio field.
func (o *ConstDeviceUnknown) SetHasExtio(v bool) {
	o.HasExtio = &v
}

// GetHasVble returns the HasVble field value if set, zero value otherwise.
func (o *ConstDeviceUnknown) GetHasVble() bool {
	if o == nil || IsNil(o.HasVble) {
		var ret bool
		return ret
	}
	return *o.HasVble
}

// GetHasVbleOk returns a tuple with the HasVble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceUnknown) GetHasVbleOk() (*bool, bool) {
	if o == nil || IsNil(o.HasVble) {
		return nil, false
	}
	return o.HasVble, true
}

// HasHasVble returns a boolean if a field has been set.
func (o *ConstDeviceUnknown) HasHasVble() bool {
	if o != nil && !IsNil(o.HasVble) {
		return true
	}

	return false
}

// SetHasVble gets a reference to the given bool and assigns it to the HasVble field.
func (o *ConstDeviceUnknown) SetHasVble(v bool) {
	o.HasVble = &v
}

// GetHasWifiBand24 returns the HasWifiBand24 field value if set, zero value otherwise.
func (o *ConstDeviceUnknown) GetHasWifiBand24() bool {
	if o == nil || IsNil(o.HasWifiBand24) {
		var ret bool
		return ret
	}
	return *o.HasWifiBand24
}

// GetHasWifiBand24Ok returns a tuple with the HasWifiBand24 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceUnknown) GetHasWifiBand24Ok() (*bool, bool) {
	if o == nil || IsNil(o.HasWifiBand24) {
		return nil, false
	}
	return o.HasWifiBand24, true
}

// HasHasWifiBand24 returns a boolean if a field has been set.
func (o *ConstDeviceUnknown) HasHasWifiBand24() bool {
	if o != nil && !IsNil(o.HasWifiBand24) {
		return true
	}

	return false
}

// SetHasWifiBand24 gets a reference to the given bool and assigns it to the HasWifiBand24 field.
func (o *ConstDeviceUnknown) SetHasWifiBand24(v bool) {
	o.HasWifiBand24 = &v
}

// GetHasWifiBand5 returns the HasWifiBand5 field value if set, zero value otherwise.
func (o *ConstDeviceUnknown) GetHasWifiBand5() bool {
	if o == nil || IsNil(o.HasWifiBand5) {
		var ret bool
		return ret
	}
	return *o.HasWifiBand5
}

// GetHasWifiBand5Ok returns a tuple with the HasWifiBand5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceUnknown) GetHasWifiBand5Ok() (*bool, bool) {
	if o == nil || IsNil(o.HasWifiBand5) {
		return nil, false
	}
	return o.HasWifiBand5, true
}

// HasHasWifiBand5 returns a boolean if a field has been set.
func (o *ConstDeviceUnknown) HasHasWifiBand5() bool {
	if o != nil && !IsNil(o.HasWifiBand5) {
		return true
	}

	return false
}

// SetHasWifiBand5 gets a reference to the given bool and assigns it to the HasWifiBand5 field.
func (o *ConstDeviceUnknown) SetHasWifiBand5(v bool) {
	o.HasWifiBand5 = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ConstDeviceUnknown) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceUnknown) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ConstDeviceUnknown) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ConstDeviceUnknown) SetModel(v string) {
	o.Model = &v
}

// GetUnmanaged returns the Unmanaged field value if set, zero value otherwise.
func (o *ConstDeviceUnknown) GetUnmanaged() bool {
	if o == nil || IsNil(o.Unmanaged) {
		var ret bool
		return ret
	}
	return *o.Unmanaged
}

// GetUnmanagedOk returns a tuple with the Unmanaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceUnknown) GetUnmanagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Unmanaged) {
		return nil, false
	}
	return o.Unmanaged, true
}

// HasUnmanaged returns a boolean if a field has been set.
func (o *ConstDeviceUnknown) HasUnmanaged() bool {
	if o != nil && !IsNil(o.Unmanaged) {
		return true
	}

	return false
}

// SetUnmanaged gets a reference to the given bool and assigns it to the Unmanaged field.
func (o *ConstDeviceUnknown) SetUnmanaged(v bool) {
	o.Unmanaged = &v
}

func (o ConstDeviceUnknown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstDeviceUnknown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApType) {
		toSerialize["ap_type"] = o.ApType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.HasExtio) {
		toSerialize["has_extio"] = o.HasExtio
	}
	if !IsNil(o.HasVble) {
		toSerialize["has_vble"] = o.HasVble
	}
	if !IsNil(o.HasWifiBand24) {
		toSerialize["has_wifi_band24"] = o.HasWifiBand24
	}
	if !IsNil(o.HasWifiBand5) {
		toSerialize["has_wifi_band5"] = o.HasWifiBand5
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Unmanaged) {
		toSerialize["unmanaged"] = o.Unmanaged
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstDeviceUnknown) UnmarshalJSON(data []byte) (err error) {
	varConstDeviceUnknown := _ConstDeviceUnknown{}

	err = json.Unmarshal(data, &varConstDeviceUnknown)

	if err != nil {
		return err
	}

	*o = ConstDeviceUnknown(varConstDeviceUnknown)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "display")
		delete(additionalProperties, "has_extio")
		delete(additionalProperties, "has_vble")
		delete(additionalProperties, "has_wifi_band24")
		delete(additionalProperties, "has_wifi_band5")
		delete(additionalProperties, "model")
		delete(additionalProperties, "unmanaged")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstDeviceUnknown struct {
	value *ConstDeviceUnknown
	isSet bool
}

func (v NullableConstDeviceUnknown) Get() *ConstDeviceUnknown {
	return v.value
}

func (v *NullableConstDeviceUnknown) Set(val *ConstDeviceUnknown) {
	v.value = val
	v.isSet = true
}

func (v NullableConstDeviceUnknown) IsSet() bool {
	return v.isSet
}

func (v *NullableConstDeviceUnknown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstDeviceUnknown(val *ConstDeviceUnknown) *NullableConstDeviceUnknown {
	return &NullableConstDeviceUnknown{value: val, isSet: true}
}

func (v NullableConstDeviceUnknown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstDeviceUnknown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

