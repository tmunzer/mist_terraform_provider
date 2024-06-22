# RemoteSyslogUsersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to [**[]RemoteSyslogContentItem**](RemoteSyslogContentItem.md) |  | [optional] 
**Match** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewRemoteSyslogUsersItem

`func NewRemoteSyslogUsersItem() *RemoteSyslogUsersItem`

NewRemoteSyslogUsersItem instantiates a new RemoteSyslogUsersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogUsersItemWithDefaults

`func NewRemoteSyslogUsersItemWithDefaults() *RemoteSyslogUsersItem`

NewRemoteSyslogUsersItemWithDefaults instantiates a new RemoteSyslogUsersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *RemoteSyslogUsersItem) GetContents() []RemoteSyslogContentItem`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *RemoteSyslogUsersItem) GetContentsOk() (*[]RemoteSyslogContentItem, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *RemoteSyslogUsersItem) SetContents(v []RemoteSyslogContentItem)`

SetContents sets Contents field to given value.

### HasContents

`func (o *RemoteSyslogUsersItem) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetMatch

`func (o *RemoteSyslogUsersItem) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *RemoteSyslogUsersItem) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *RemoteSyslogUsersItem) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *RemoteSyslogUsersItem) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetUser

`func (o *RemoteSyslogUsersItem) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RemoteSyslogUsersItem) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RemoteSyslogUsersItem) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RemoteSyslogUsersItem) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


