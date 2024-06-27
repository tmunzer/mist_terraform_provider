# SnmpVacmAccessRopertyPrefixListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextPrefix** | Pointer to **string** | only required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;context_prefix&#x60; | [optional] 
**NotifyView** | Pointer to **string** | refer to view name | [optional] 
**ReadView** | Pointer to **string** | refer to view name | [optional] 
**SecurityLevel** | Pointer to **string** |  | [optional] 
**SecurityModel** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**WriteView** | Pointer to **string** | refer to view name | [optional] 

## Methods

### NewSnmpVacmAccessRopertyPrefixListInner

`func NewSnmpVacmAccessRopertyPrefixListInner() *SnmpVacmAccessRopertyPrefixListInner`

NewSnmpVacmAccessRopertyPrefixListInner instantiates a new SnmpVacmAccessRopertyPrefixListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpVacmAccessRopertyPrefixListInnerWithDefaults

`func NewSnmpVacmAccessRopertyPrefixListInnerWithDefaults() *SnmpVacmAccessRopertyPrefixListInner`

NewSnmpVacmAccessRopertyPrefixListInnerWithDefaults instantiates a new SnmpVacmAccessRopertyPrefixListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextPrefix

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetContextPrefix() string`

GetContextPrefix returns the ContextPrefix field if non-nil, zero value otherwise.

### GetContextPrefixOk

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetContextPrefixOk() (*string, bool)`

GetContextPrefixOk returns a tuple with the ContextPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextPrefix

`func (o *SnmpVacmAccessRopertyPrefixListInner) SetContextPrefix(v string)`

SetContextPrefix sets ContextPrefix field to given value.

### HasContextPrefix

`func (o *SnmpVacmAccessRopertyPrefixListInner) HasContextPrefix() bool`

HasContextPrefix returns a boolean if a field has been set.

### GetNotifyView

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetNotifyView() string`

GetNotifyView returns the NotifyView field if non-nil, zero value otherwise.

### GetNotifyViewOk

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetNotifyViewOk() (*string, bool)`

GetNotifyViewOk returns a tuple with the NotifyView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyView

`func (o *SnmpVacmAccessRopertyPrefixListInner) SetNotifyView(v string)`

SetNotifyView sets NotifyView field to given value.

### HasNotifyView

`func (o *SnmpVacmAccessRopertyPrefixListInner) HasNotifyView() bool`

HasNotifyView returns a boolean if a field has been set.

### GetReadView

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetReadView() string`

GetReadView returns the ReadView field if non-nil, zero value otherwise.

### GetReadViewOk

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetReadViewOk() (*string, bool)`

GetReadViewOk returns a tuple with the ReadView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadView

`func (o *SnmpVacmAccessRopertyPrefixListInner) SetReadView(v string)`

SetReadView sets ReadView field to given value.

### HasReadView

`func (o *SnmpVacmAccessRopertyPrefixListInner) HasReadView() bool`

HasReadView returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *SnmpVacmAccessRopertyPrefixListInner) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *SnmpVacmAccessRopertyPrefixListInner) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetSecurityModel

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetSecurityModel() string`

GetSecurityModel returns the SecurityModel field if non-nil, zero value otherwise.

### GetSecurityModelOk

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetSecurityModelOk() (*string, bool)`

GetSecurityModelOk returns a tuple with the SecurityModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityModel

`func (o *SnmpVacmAccessRopertyPrefixListInner) SetSecurityModel(v string)`

SetSecurityModel sets SecurityModel field to given value.

### HasSecurityModel

`func (o *SnmpVacmAccessRopertyPrefixListInner) HasSecurityModel() bool`

HasSecurityModel returns a boolean if a field has been set.

### GetType

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnmpVacmAccessRopertyPrefixListInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SnmpVacmAccessRopertyPrefixListInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWriteView

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetWriteView() string`

GetWriteView returns the WriteView field if non-nil, zero value otherwise.

### GetWriteViewOk

`func (o *SnmpVacmAccessRopertyPrefixListInner) GetWriteViewOk() (*string, bool)`

GetWriteViewOk returns a tuple with the WriteView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteView

`func (o *SnmpVacmAccessRopertyPrefixListInner) SetWriteView(v string)`

SetWriteView sets WriteView field to given value.

### HasWriteView

`func (o *SnmpVacmAccessRopertyPrefixListInner) HasWriteView() bool`

HasWriteView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


