# SnmpV3Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notify** | Pointer to [**[]SnmpV3ConfigNotifyInner**](SnmpV3ConfigNotifyInner.md) |  | [optional] 
**NotifyFilter** | Pointer to [**[]SnmpV3ConfigNotifyFilterInner**](SnmpV3ConfigNotifyFilterInner.md) |  | [optional] 
**TargetAddress** | Pointer to [**[]SnmpV3ConfigTargetAddressInner**](SnmpV3ConfigTargetAddressInner.md) |  | [optional] 
**TargetParameters** | Pointer to [**[]SnmpV3ConfigTargetParametersInner**](SnmpV3ConfigTargetParametersInner.md) |  | [optional] 
**Usm** | Pointer to [**SnmpUsm**](SnmpUsm.md) |  | [optional] 
**Vacm** | Pointer to [**SnmpVacm**](SnmpVacm.md) |  | [optional] 

## Methods

### NewSnmpV3Config

`func NewSnmpV3Config() *SnmpV3Config`

NewSnmpV3Config instantiates a new SnmpV3Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpV3ConfigWithDefaults

`func NewSnmpV3ConfigWithDefaults() *SnmpV3Config`

NewSnmpV3ConfigWithDefaults instantiates a new SnmpV3Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotify

`func (o *SnmpV3Config) GetNotify() []SnmpV3ConfigNotifyInner`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *SnmpV3Config) GetNotifyOk() (*[]SnmpV3ConfigNotifyInner, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *SnmpV3Config) SetNotify(v []SnmpV3ConfigNotifyInner)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *SnmpV3Config) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyFilter

`func (o *SnmpV3Config) GetNotifyFilter() []SnmpV3ConfigNotifyFilterInner`

GetNotifyFilter returns the NotifyFilter field if non-nil, zero value otherwise.

### GetNotifyFilterOk

`func (o *SnmpV3Config) GetNotifyFilterOk() (*[]SnmpV3ConfigNotifyFilterInner, bool)`

GetNotifyFilterOk returns a tuple with the NotifyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFilter

`func (o *SnmpV3Config) SetNotifyFilter(v []SnmpV3ConfigNotifyFilterInner)`

SetNotifyFilter sets NotifyFilter field to given value.

### HasNotifyFilter

`func (o *SnmpV3Config) HasNotifyFilter() bool`

HasNotifyFilter returns a boolean if a field has been set.

### GetTargetAddress

`func (o *SnmpV3Config) GetTargetAddress() []SnmpV3ConfigTargetAddressInner`

GetTargetAddress returns the TargetAddress field if non-nil, zero value otherwise.

### GetTargetAddressOk

`func (o *SnmpV3Config) GetTargetAddressOk() (*[]SnmpV3ConfigTargetAddressInner, bool)`

GetTargetAddressOk returns a tuple with the TargetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAddress

`func (o *SnmpV3Config) SetTargetAddress(v []SnmpV3ConfigTargetAddressInner)`

SetTargetAddress sets TargetAddress field to given value.

### HasTargetAddress

`func (o *SnmpV3Config) HasTargetAddress() bool`

HasTargetAddress returns a boolean if a field has been set.

### GetTargetParameters

`func (o *SnmpV3Config) GetTargetParameters() []SnmpV3ConfigTargetParametersInner`

GetTargetParameters returns the TargetParameters field if non-nil, zero value otherwise.

### GetTargetParametersOk

`func (o *SnmpV3Config) GetTargetParametersOk() (*[]SnmpV3ConfigTargetParametersInner, bool)`

GetTargetParametersOk returns a tuple with the TargetParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetParameters

`func (o *SnmpV3Config) SetTargetParameters(v []SnmpV3ConfigTargetParametersInner)`

SetTargetParameters sets TargetParameters field to given value.

### HasTargetParameters

`func (o *SnmpV3Config) HasTargetParameters() bool`

HasTargetParameters returns a boolean if a field has been set.

### GetUsm

`func (o *SnmpV3Config) GetUsm() SnmpUsm`

GetUsm returns the Usm field if non-nil, zero value otherwise.

### GetUsmOk

`func (o *SnmpV3Config) GetUsmOk() (*SnmpUsm, bool)`

GetUsmOk returns a tuple with the Usm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsm

`func (o *SnmpV3Config) SetUsm(v SnmpUsm)`

SetUsm sets Usm field to given value.

### HasUsm

`func (o *SnmpV3Config) HasUsm() bool`

HasUsm returns a boolean if a field has been set.

### GetVacm

`func (o *SnmpV3Config) GetVacm() SnmpVacm`

GetVacm returns the Vacm field if non-nil, zero value otherwise.

### GetVacmOk

`func (o *SnmpV3Config) GetVacmOk() (*SnmpVacm, bool)`

GetVacmOk returns a tuple with the Vacm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacm

`func (o *SnmpV3Config) SetVacm(v SnmpVacm)`

SetVacm sets Vacm field to given value.

### HasVacm

`func (o *SnmpV3Config) HasVacm() bool`

HasVacm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


