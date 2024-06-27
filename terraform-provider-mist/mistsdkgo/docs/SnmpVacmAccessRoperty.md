# SnmpVacmAccessRoperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** |  | [optional] 
**PrefixList** | Pointer to [**[]SnmpVacmAccessRopertyPrefixListInner**](SnmpVacmAccessRopertyPrefixListInner.md) |  | [optional] 

## Methods

### NewSnmpVacmAccessRoperty

`func NewSnmpVacmAccessRoperty() *SnmpVacmAccessRoperty`

NewSnmpVacmAccessRoperty instantiates a new SnmpVacmAccessRoperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpVacmAccessRopertyWithDefaults

`func NewSnmpVacmAccessRopertyWithDefaults() *SnmpVacmAccessRoperty`

NewSnmpVacmAccessRopertyWithDefaults instantiates a new SnmpVacmAccessRoperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *SnmpVacmAccessRoperty) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *SnmpVacmAccessRoperty) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *SnmpVacmAccessRoperty) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *SnmpVacmAccessRoperty) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetPrefixList

`func (o *SnmpVacmAccessRoperty) GetPrefixList() []SnmpVacmAccessRopertyPrefixListInner`

GetPrefixList returns the PrefixList field if non-nil, zero value otherwise.

### GetPrefixListOk

`func (o *SnmpVacmAccessRoperty) GetPrefixListOk() (*[]SnmpVacmAccessRopertyPrefixListInner, bool)`

GetPrefixListOk returns a tuple with the PrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixList

`func (o *SnmpVacmAccessRoperty) SetPrefixList(v []SnmpVacmAccessRopertyPrefixListInner)`

SetPrefixList sets PrefixList field to given value.

### HasPrefixList

`func (o *SnmpVacmAccessRoperty) HasPrefixList() bool`

HasPrefixList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


