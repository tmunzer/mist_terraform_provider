# SnmpV3ConfigNotifyFilterInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to [**[]SnmpV3ConfigNotifyFilterInnerContentsInner**](SnmpV3ConfigNotifyFilterInnerContentsInner.md) |  | [optional] 
**ProfileName** | Pointer to **string** |  | [optional] 

## Methods

### NewSnmpV3ConfigNotifyFilterInner

`func NewSnmpV3ConfigNotifyFilterInner() *SnmpV3ConfigNotifyFilterInner`

NewSnmpV3ConfigNotifyFilterInner instantiates a new SnmpV3ConfigNotifyFilterInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpV3ConfigNotifyFilterInnerWithDefaults

`func NewSnmpV3ConfigNotifyFilterInnerWithDefaults() *SnmpV3ConfigNotifyFilterInner`

NewSnmpV3ConfigNotifyFilterInnerWithDefaults instantiates a new SnmpV3ConfigNotifyFilterInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *SnmpV3ConfigNotifyFilterInner) GetContents() []SnmpV3ConfigNotifyFilterInnerContentsInner`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *SnmpV3ConfigNotifyFilterInner) GetContentsOk() (*[]SnmpV3ConfigNotifyFilterInnerContentsInner, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *SnmpV3ConfigNotifyFilterInner) SetContents(v []SnmpV3ConfigNotifyFilterInnerContentsInner)`

SetContents sets Contents field to given value.

### HasContents

`func (o *SnmpV3ConfigNotifyFilterInner) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetProfileName

`func (o *SnmpV3ConfigNotifyFilterInner) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *SnmpV3ConfigNotifyFilterInner) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *SnmpV3ConfigNotifyFilterInner) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *SnmpV3ConfigNotifyFilterInner) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


