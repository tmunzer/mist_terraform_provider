# DynamicPortUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**IpConfigType**](IpConfigType.md) |  | [default to IPCONFIGTYPE_DYNAMIC]
**ResetDefaultWhen** | Pointer to [**DynamicPortUsageResetDefaultWhen**](DynamicPortUsageResetDefaultWhen.md) |  | [optional] [default to DYNAMICPORTUSAGERESETDEFAULTWHEN_LINK_DOWN]
**Rules** | Pointer to [**[]DynamicPortUsageRule**](DynamicPortUsageRule.md) |  | [optional] 

## Methods

### NewDynamicPortUsage

`func NewDynamicPortUsage(mode IpConfigType, ) *DynamicPortUsage`

NewDynamicPortUsage instantiates a new DynamicPortUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicPortUsageWithDefaults

`func NewDynamicPortUsageWithDefaults() *DynamicPortUsage`

NewDynamicPortUsageWithDefaults instantiates a new DynamicPortUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *DynamicPortUsage) GetMode() IpConfigType`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DynamicPortUsage) GetModeOk() (*IpConfigType, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DynamicPortUsage) SetMode(v IpConfigType)`

SetMode sets Mode field to given value.


### GetResetDefaultWhen

`func (o *DynamicPortUsage) GetResetDefaultWhen() DynamicPortUsageResetDefaultWhen`

GetResetDefaultWhen returns the ResetDefaultWhen field if non-nil, zero value otherwise.

### GetResetDefaultWhenOk

`func (o *DynamicPortUsage) GetResetDefaultWhenOk() (*DynamicPortUsageResetDefaultWhen, bool)`

GetResetDefaultWhenOk returns a tuple with the ResetDefaultWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetDefaultWhen

`func (o *DynamicPortUsage) SetResetDefaultWhen(v DynamicPortUsageResetDefaultWhen)`

SetResetDefaultWhen sets ResetDefaultWhen field to given value.

### HasResetDefaultWhen

`func (o *DynamicPortUsage) HasResetDefaultWhen() bool`

HasResetDefaultWhen returns a boolean if a field has been set.

### GetRules

`func (o *DynamicPortUsage) GetRules() []DynamicPortUsageRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DynamicPortUsage) GetRulesOk() (*[]DynamicPortUsageRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DynamicPortUsage) SetRules(v []DynamicPortUsageRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *DynamicPortUsage) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


