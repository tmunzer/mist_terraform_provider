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

// checks if the NacPortalSso type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NacPortalSso{}

// NacPortalSso struct for NacPortalSso
type NacPortalSso struct {
	IdpCert *string `json:"idp_cert,omitempty"`
	IdpSignAlgo *string `json:"idp_sign_algo,omitempty"`
	IdpSsoUrl *string `json:"idp_sso_url,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	NameidFormat *string `json:"nameid_format,omitempty"`
	SsoRoleMatching []NacPortalSsoRoleMatching `json:"sso_role_matching,omitempty"`
	// if it's desired to inject a role into Cert's Subject (so it can be used later on in policy)
	UseSsoRoleForCert *bool `json:"use_sso_role_for_cert,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NacPortalSso NacPortalSso

// NewNacPortalSso instantiates a new NacPortalSso object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNacPortalSso() *NacPortalSso {
	this := NacPortalSso{}
	return &this
}

// NewNacPortalSsoWithDefaults instantiates a new NacPortalSso object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNacPortalSsoWithDefaults() *NacPortalSso {
	this := NacPortalSso{}
	return &this
}

// GetIdpCert returns the IdpCert field value if set, zero value otherwise.
func (o *NacPortalSso) GetIdpCert() string {
	if o == nil || IsNil(o.IdpCert) {
		var ret string
		return ret
	}
	return *o.IdpCert
}

// GetIdpCertOk returns a tuple with the IdpCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacPortalSso) GetIdpCertOk() (*string, bool) {
	if o == nil || IsNil(o.IdpCert) {
		return nil, false
	}
	return o.IdpCert, true
}

// HasIdpCert returns a boolean if a field has been set.
func (o *NacPortalSso) HasIdpCert() bool {
	if o != nil && !IsNil(o.IdpCert) {
		return true
	}

	return false
}

// SetIdpCert gets a reference to the given string and assigns it to the IdpCert field.
func (o *NacPortalSso) SetIdpCert(v string) {
	o.IdpCert = &v
}

// GetIdpSignAlgo returns the IdpSignAlgo field value if set, zero value otherwise.
func (o *NacPortalSso) GetIdpSignAlgo() string {
	if o == nil || IsNil(o.IdpSignAlgo) {
		var ret string
		return ret
	}
	return *o.IdpSignAlgo
}

// GetIdpSignAlgoOk returns a tuple with the IdpSignAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacPortalSso) GetIdpSignAlgoOk() (*string, bool) {
	if o == nil || IsNil(o.IdpSignAlgo) {
		return nil, false
	}
	return o.IdpSignAlgo, true
}

// HasIdpSignAlgo returns a boolean if a field has been set.
func (o *NacPortalSso) HasIdpSignAlgo() bool {
	if o != nil && !IsNil(o.IdpSignAlgo) {
		return true
	}

	return false
}

// SetIdpSignAlgo gets a reference to the given string and assigns it to the IdpSignAlgo field.
func (o *NacPortalSso) SetIdpSignAlgo(v string) {
	o.IdpSignAlgo = &v
}

// GetIdpSsoUrl returns the IdpSsoUrl field value if set, zero value otherwise.
func (o *NacPortalSso) GetIdpSsoUrl() string {
	if o == nil || IsNil(o.IdpSsoUrl) {
		var ret string
		return ret
	}
	return *o.IdpSsoUrl
}

// GetIdpSsoUrlOk returns a tuple with the IdpSsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacPortalSso) GetIdpSsoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IdpSsoUrl) {
		return nil, false
	}
	return o.IdpSsoUrl, true
}

// HasIdpSsoUrl returns a boolean if a field has been set.
func (o *NacPortalSso) HasIdpSsoUrl() bool {
	if o != nil && !IsNil(o.IdpSsoUrl) {
		return true
	}

	return false
}

// SetIdpSsoUrl gets a reference to the given string and assigns it to the IdpSsoUrl field.
func (o *NacPortalSso) SetIdpSsoUrl(v string) {
	o.IdpSsoUrl = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *NacPortalSso) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacPortalSso) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *NacPortalSso) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *NacPortalSso) SetIssuer(v string) {
	o.Issuer = &v
}

// GetNameidFormat returns the NameidFormat field value if set, zero value otherwise.
func (o *NacPortalSso) GetNameidFormat() string {
	if o == nil || IsNil(o.NameidFormat) {
		var ret string
		return ret
	}
	return *o.NameidFormat
}

// GetNameidFormatOk returns a tuple with the NameidFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacPortalSso) GetNameidFormatOk() (*string, bool) {
	if o == nil || IsNil(o.NameidFormat) {
		return nil, false
	}
	return o.NameidFormat, true
}

// HasNameidFormat returns a boolean if a field has been set.
func (o *NacPortalSso) HasNameidFormat() bool {
	if o != nil && !IsNil(o.NameidFormat) {
		return true
	}

	return false
}

// SetNameidFormat gets a reference to the given string and assigns it to the NameidFormat field.
func (o *NacPortalSso) SetNameidFormat(v string) {
	o.NameidFormat = &v
}

// GetSsoRoleMatching returns the SsoRoleMatching field value if set, zero value otherwise.
func (o *NacPortalSso) GetSsoRoleMatching() []NacPortalSsoRoleMatching {
	if o == nil || IsNil(o.SsoRoleMatching) {
		var ret []NacPortalSsoRoleMatching
		return ret
	}
	return o.SsoRoleMatching
}

// GetSsoRoleMatchingOk returns a tuple with the SsoRoleMatching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacPortalSso) GetSsoRoleMatchingOk() ([]NacPortalSsoRoleMatching, bool) {
	if o == nil || IsNil(o.SsoRoleMatching) {
		return nil, false
	}
	return o.SsoRoleMatching, true
}

// HasSsoRoleMatching returns a boolean if a field has been set.
func (o *NacPortalSso) HasSsoRoleMatching() bool {
	if o != nil && !IsNil(o.SsoRoleMatching) {
		return true
	}

	return false
}

// SetSsoRoleMatching gets a reference to the given []NacPortalSsoRoleMatching and assigns it to the SsoRoleMatching field.
func (o *NacPortalSso) SetSsoRoleMatching(v []NacPortalSsoRoleMatching) {
	o.SsoRoleMatching = v
}

// GetUseSsoRoleForCert returns the UseSsoRoleForCert field value if set, zero value otherwise.
func (o *NacPortalSso) GetUseSsoRoleForCert() bool {
	if o == nil || IsNil(o.UseSsoRoleForCert) {
		var ret bool
		return ret
	}
	return *o.UseSsoRoleForCert
}

// GetUseSsoRoleForCertOk returns a tuple with the UseSsoRoleForCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacPortalSso) GetUseSsoRoleForCertOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSsoRoleForCert) {
		return nil, false
	}
	return o.UseSsoRoleForCert, true
}

// HasUseSsoRoleForCert returns a boolean if a field has been set.
func (o *NacPortalSso) HasUseSsoRoleForCert() bool {
	if o != nil && !IsNil(o.UseSsoRoleForCert) {
		return true
	}

	return false
}

// SetUseSsoRoleForCert gets a reference to the given bool and assigns it to the UseSsoRoleForCert field.
func (o *NacPortalSso) SetUseSsoRoleForCert(v bool) {
	o.UseSsoRoleForCert = &v
}

func (o NacPortalSso) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NacPortalSso) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdpCert) {
		toSerialize["idp_cert"] = o.IdpCert
	}
	if !IsNil(o.IdpSignAlgo) {
		toSerialize["idp_sign_algo"] = o.IdpSignAlgo
	}
	if !IsNil(o.IdpSsoUrl) {
		toSerialize["idp_sso_url"] = o.IdpSsoUrl
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.NameidFormat) {
		toSerialize["nameid_format"] = o.NameidFormat
	}
	if !IsNil(o.SsoRoleMatching) {
		toSerialize["sso_role_matching"] = o.SsoRoleMatching
	}
	if !IsNil(o.UseSsoRoleForCert) {
		toSerialize["use_sso_role_for_cert"] = o.UseSsoRoleForCert
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NacPortalSso) UnmarshalJSON(data []byte) (err error) {
	varNacPortalSso := _NacPortalSso{}

	err = json.Unmarshal(data, &varNacPortalSso)

	if err != nil {
		return err
	}

	*o = NacPortalSso(varNacPortalSso)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "idp_cert")
		delete(additionalProperties, "idp_sign_algo")
		delete(additionalProperties, "idp_sso_url")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "nameid_format")
		delete(additionalProperties, "sso_role_matching")
		delete(additionalProperties, "use_sso_role_for_cert")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNacPortalSso struct {
	value *NacPortalSso
	isSet bool
}

func (v NullableNacPortalSso) Get() *NacPortalSso {
	return v.value
}

func (v *NullableNacPortalSso) Set(val *NacPortalSso) {
	v.value = val
	v.isSet = true
}

func (v NullableNacPortalSso) IsSet() bool {
	return v.isSet
}

func (v *NullableNacPortalSso) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNacPortalSso(val *NacPortalSso) *NullableNacPortalSso {
	return &NullableNacPortalSso{value: val, isSet: true}
}

func (v NullableNacPortalSso) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNacPortalSso) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

