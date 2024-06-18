# SwitchExtraRoutes6Value

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

### NewSwitchExtraRoutes6Value

`func NewSwitchExtraRoutes6Value() *SwitchExtraRoutes6Value`

NewSwitchExtraRoutes6Value instantiates a new SwitchExtraRoutes6Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchExtraRoutes6ValueWithDefaults

`func NewSwitchExtraRoutes6ValueWithDefaults() *SwitchExtraRoutes6Value`

NewSwitchExtraRoutes6ValueWithDefaults instantiates a new SwitchExtraRoutes6Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscard

`func (o *SwitchExtraRoutes6Value) GetDiscard() bool`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *SwitchExtraRoutes6Value) GetDiscardOk() (*bool, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *SwitchExtraRoutes6Value) SetDiscard(v bool)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *SwitchExtraRoutes6Value) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetMetric

`func (o *SwitchExtraRoutes6Value) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SwitchExtraRoutes6Value) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SwitchExtraRoutes6Value) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *SwitchExtraRoutes6Value) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *SwitchExtraRoutes6Value) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *SwitchExtraRoutes6Value) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetNextQualified

`func (o *SwitchExtraRoutes6Value) GetNextQualified() map[string]SwitchExtraRoutesValueNextQualifiedValue`

GetNextQualified returns the NextQualified field if non-nil, zero value otherwise.

### GetNextQualifiedOk

`func (o *SwitchExtraRoutes6Value) GetNextQualifiedOk() (*map[string]SwitchExtraRoutesValueNextQualifiedValue, bool)`

GetNextQualifiedOk returns a tuple with the NextQualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextQualified

`func (o *SwitchExtraRoutes6Value) SetNextQualified(v map[string]SwitchExtraRoutesValueNextQualifiedValue)`

SetNextQualified sets NextQualified field to given value.

### HasNextQualified

`func (o *SwitchExtraRoutes6Value) HasNextQualified() bool`

HasNextQualified returns a boolean if a field has been set.

### GetNoResolve

`func (o *SwitchExtraRoutes6Value) GetNoResolve() bool`

GetNoResolve returns the NoResolve field if non-nil, zero value otherwise.

### GetNoResolveOk

`func (o *SwitchExtraRoutes6Value) GetNoResolveOk() (*bool, bool)`

GetNoResolveOk returns a tuple with the NoResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoResolve

`func (o *SwitchExtraRoutes6Value) SetNoResolve(v bool)`

SetNoResolve sets NoResolve field to given value.

### HasNoResolve

`func (o *SwitchExtraRoutes6Value) HasNoResolve() bool`

HasNoResolve returns a boolean if a field has been set.

### GetPreference

`func (o *SwitchExtraRoutes6Value) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *SwitchExtraRoutes6Value) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *SwitchExtraRoutes6Value) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *SwitchExtraRoutes6Value) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### SetPreferenceNil

`func (o *SwitchExtraRoutes6Value) SetPreferenceNil(b bool)`

 SetPreferenceNil sets the value for Preference to be an explicit nil

### UnsetPreference
`func (o *SwitchExtraRoutes6Value) UnsetPreference()`

UnsetPreference ensures that no value is present for Preference, not even an explicit nil
### GetVia

`func (o *SwitchExtraRoutes6Value) GetVia() string`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *SwitchExtraRoutes6Value) GetViaOk() (*string, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *SwitchExtraRoutes6Value) SetVia(v string)`

SetVia sets Via field to given value.

### HasVia

`func (o *SwitchExtraRoutes6Value) HasVia() bool`

HasVia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


