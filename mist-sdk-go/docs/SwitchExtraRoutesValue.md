# SwitchExtraRoutesValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discard** | Pointer to **bool** | this takes precedence | [optional] [default to false]
**Metric** | Pointer to **NullableInt32** |  | [optional] 
**NextQualified** | Pointer to [**map[string]SwitchExtraRoutesValueNextQualifiedValue**](SwitchExtraRoutesValueNextQualifiedValue.md) |  | [optional] 
**NoResolve** | Pointer to **bool** |  | [optional] [default to false]
**Preference** | Pointer to **NullableInt32** |  | [optional] 
**Via** | Pointer to **string** | next-hop IP Address | [optional] 

## Methods

### NewSwitchExtraRoutesValue

`func NewSwitchExtraRoutesValue() *SwitchExtraRoutesValue`

NewSwitchExtraRoutesValue instantiates a new SwitchExtraRoutesValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchExtraRoutesValueWithDefaults

`func NewSwitchExtraRoutesValueWithDefaults() *SwitchExtraRoutesValue`

NewSwitchExtraRoutesValueWithDefaults instantiates a new SwitchExtraRoutesValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscard

`func (o *SwitchExtraRoutesValue) GetDiscard() bool`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *SwitchExtraRoutesValue) GetDiscardOk() (*bool, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *SwitchExtraRoutesValue) SetDiscard(v bool)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *SwitchExtraRoutesValue) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetMetric

`func (o *SwitchExtraRoutesValue) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SwitchExtraRoutesValue) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SwitchExtraRoutesValue) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *SwitchExtraRoutesValue) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *SwitchExtraRoutesValue) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *SwitchExtraRoutesValue) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetNextQualified

`func (o *SwitchExtraRoutesValue) GetNextQualified() map[string]SwitchExtraRoutesValueNextQualifiedValue`

GetNextQualified returns the NextQualified field if non-nil, zero value otherwise.

### GetNextQualifiedOk

`func (o *SwitchExtraRoutesValue) GetNextQualifiedOk() (*map[string]SwitchExtraRoutesValueNextQualifiedValue, bool)`

GetNextQualifiedOk returns a tuple with the NextQualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextQualified

`func (o *SwitchExtraRoutesValue) SetNextQualified(v map[string]SwitchExtraRoutesValueNextQualifiedValue)`

SetNextQualified sets NextQualified field to given value.

### HasNextQualified

`func (o *SwitchExtraRoutesValue) HasNextQualified() bool`

HasNextQualified returns a boolean if a field has been set.

### GetNoResolve

`func (o *SwitchExtraRoutesValue) GetNoResolve() bool`

GetNoResolve returns the NoResolve field if non-nil, zero value otherwise.

### GetNoResolveOk

`func (o *SwitchExtraRoutesValue) GetNoResolveOk() (*bool, bool)`

GetNoResolveOk returns a tuple with the NoResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoResolve

`func (o *SwitchExtraRoutesValue) SetNoResolve(v bool)`

SetNoResolve sets NoResolve field to given value.

### HasNoResolve

`func (o *SwitchExtraRoutesValue) HasNoResolve() bool`

HasNoResolve returns a boolean if a field has been set.

### GetPreference

`func (o *SwitchExtraRoutesValue) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *SwitchExtraRoutesValue) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *SwitchExtraRoutesValue) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *SwitchExtraRoutesValue) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### SetPreferenceNil

`func (o *SwitchExtraRoutesValue) SetPreferenceNil(b bool)`

 SetPreferenceNil sets the value for Preference to be an explicit nil

### UnsetPreference
`func (o *SwitchExtraRoutesValue) UnsetPreference()`

UnsetPreference ensures that no value is present for Preference, not even an explicit nil
### GetVia

`func (o *SwitchExtraRoutesValue) GetVia() string`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *SwitchExtraRoutesValue) GetViaOk() (*string, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *SwitchExtraRoutesValue) SetVia(v string)`

SetVia sets Via field to given value.

### HasVia

`func (o *SwitchExtraRoutesValue) HasVia() bool`

HasVia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


