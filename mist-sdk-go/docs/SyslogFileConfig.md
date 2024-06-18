# SyslogFileConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archive** | Pointer to [**SyslogFileConfigArchive**](SyslogFileConfigArchive.md) |  | [optional] 
**Contents** | Pointer to [**[]RemoteSyslogContent**](RemoteSyslogContent.md) |  | [optional] 
**ExplicitPriority** | Pointer to **bool** |  | [optional] 
**File** | Pointer to **string** |  | [optional] 
**Match** | Pointer to **string** |  | [optional] 
**StructuredData** | Pointer to **bool** |  | [optional] 

## Methods

### NewSyslogFileConfig

`func NewSyslogFileConfig() *SyslogFileConfig`

NewSyslogFileConfig instantiates a new SyslogFileConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogFileConfigWithDefaults

`func NewSyslogFileConfigWithDefaults() *SyslogFileConfig`

NewSyslogFileConfigWithDefaults instantiates a new SyslogFileConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *SyslogFileConfig) GetArchive() SyslogFileConfigArchive`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *SyslogFileConfig) GetArchiveOk() (*SyslogFileConfigArchive, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *SyslogFileConfig) SetArchive(v SyslogFileConfigArchive)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *SyslogFileConfig) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetContents

`func (o *SyslogFileConfig) GetContents() []RemoteSyslogContent`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *SyslogFileConfig) GetContentsOk() (*[]RemoteSyslogContent, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *SyslogFileConfig) SetContents(v []RemoteSyslogContent)`

SetContents sets Contents field to given value.

### HasContents

`func (o *SyslogFileConfig) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetExplicitPriority

`func (o *SyslogFileConfig) GetExplicitPriority() bool`

GetExplicitPriority returns the ExplicitPriority field if non-nil, zero value otherwise.

### GetExplicitPriorityOk

`func (o *SyslogFileConfig) GetExplicitPriorityOk() (*bool, bool)`

GetExplicitPriorityOk returns a tuple with the ExplicitPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPriority

`func (o *SyslogFileConfig) SetExplicitPriority(v bool)`

SetExplicitPriority sets ExplicitPriority field to given value.

### HasExplicitPriority

`func (o *SyslogFileConfig) HasExplicitPriority() bool`

HasExplicitPriority returns a boolean if a field has been set.

### GetFile

`func (o *SyslogFileConfig) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SyslogFileConfig) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SyslogFileConfig) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *SyslogFileConfig) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetMatch

`func (o *SyslogFileConfig) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *SyslogFileConfig) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *SyslogFileConfig) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *SyslogFileConfig) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetStructuredData

`func (o *SyslogFileConfig) GetStructuredData() bool`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *SyslogFileConfig) GetStructuredDataOk() (*bool, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *SyslogFileConfig) SetStructuredData(v bool)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *SyslogFileConfig) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


