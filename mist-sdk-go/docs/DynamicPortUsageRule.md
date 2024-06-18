# DynamicPortUsageRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Equals** | Pointer to **string** |  | [optional] 
**EqualsAny** | Pointer to **[]string** | use &#x60;equals_any&#x60; to match any item in a list | [optional] 
**Expression** | Pointer to **string** | \&quot;[0:3]\&quot;:\&quot;abcdef\&quot; -&gt; \&quot;abc\&quot; \&quot;split(.)[1]\&quot;: \&quot;a.b.c\&quot; -&gt; \&quot;b\&quot; \&quot;split(-)[1][0:3]: \&quot;a1234-b5678-c90\&quot; -&gt; \&quot;b56\&quot; | [optional] 
**Src** | **string** |  | 
**Usage** | Pointer to **string** | &#x60;port_usage&#x60; name | [optional] 

## Methods

### NewDynamicPortUsageRule

`func NewDynamicPortUsageRule(src string, ) *DynamicPortUsageRule`

NewDynamicPortUsageRule instantiates a new DynamicPortUsageRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicPortUsageRuleWithDefaults

`func NewDynamicPortUsageRuleWithDefaults() *DynamicPortUsageRule`

NewDynamicPortUsageRuleWithDefaults instantiates a new DynamicPortUsageRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquals

`func (o *DynamicPortUsageRule) GetEquals() string`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *DynamicPortUsageRule) GetEqualsOk() (*string, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *DynamicPortUsageRule) SetEquals(v string)`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *DynamicPortUsageRule) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### GetEqualsAny

`func (o *DynamicPortUsageRule) GetEqualsAny() []string`

GetEqualsAny returns the EqualsAny field if non-nil, zero value otherwise.

### GetEqualsAnyOk

`func (o *DynamicPortUsageRule) GetEqualsAnyOk() (*[]string, bool)`

GetEqualsAnyOk returns a tuple with the EqualsAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualsAny

`func (o *DynamicPortUsageRule) SetEqualsAny(v []string)`

SetEqualsAny sets EqualsAny field to given value.

### HasEqualsAny

`func (o *DynamicPortUsageRule) HasEqualsAny() bool`

HasEqualsAny returns a boolean if a field has been set.

### GetExpression

`func (o *DynamicPortUsageRule) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *DynamicPortUsageRule) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *DynamicPortUsageRule) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *DynamicPortUsageRule) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetSrc

`func (o *DynamicPortUsageRule) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *DynamicPortUsageRule) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *DynamicPortUsageRule) SetSrc(v string)`

SetSrc sets Src field to given value.


### GetUsage

`func (o *DynamicPortUsageRule) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DynamicPortUsageRule) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DynamicPortUsageRule) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *DynamicPortUsageRule) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


