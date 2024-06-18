# JunosAclPolicyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**AllowDeny**](AllowDeny.md) |  | [optional] [default to ALLOWDENY_ALLOW]
**DstTag** | Pointer to **string** |  | [optional] 

## Methods

### NewJunosAclPolicyAction

`func NewJunosAclPolicyAction() *JunosAclPolicyAction`

NewJunosAclPolicyAction instantiates a new JunosAclPolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosAclPolicyActionWithDefaults

`func NewJunosAclPolicyActionWithDefaults() *JunosAclPolicyAction`

NewJunosAclPolicyActionWithDefaults instantiates a new JunosAclPolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *JunosAclPolicyAction) GetAction() AllowDeny`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *JunosAclPolicyAction) GetActionOk() (*AllowDeny, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *JunosAclPolicyAction) SetAction(v AllowDeny)`

SetAction sets Action field to given value.

### HasAction

`func (o *JunosAclPolicyAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDstTag

`func (o *JunosAclPolicyAction) GetDstTag() string`

GetDstTag returns the DstTag field if non-nil, zero value otherwise.

### GetDstTagOk

`func (o *JunosAclPolicyAction) GetDstTagOk() (*string, bool)`

GetDstTagOk returns a tuple with the DstTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstTag

`func (o *JunosAclPolicyAction) SetDstTag(v string)`

SetDstTag sets DstTag field to given value.

### HasDstTag

`func (o *JunosAclPolicyAction) HasDstTag() bool`

HasDstTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


