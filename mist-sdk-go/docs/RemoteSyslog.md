# RemoteSyslog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteSyslogArchive** | Pointer to [**RemoteSyslogRemoteSyslogArchive**](RemoteSyslogRemoteSyslogArchive.md) |  | [optional] 
**RemoteSyslogConsole** | Pointer to [**RemoteSyslogRemoteSyslogConsole**](RemoteSyslogRemoteSyslogConsole.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**RemoteSyslogFiles** | Pointer to [**[]SyslogFileConfig**](SyslogFileConfig.md) |  | [optional] 
**Network** | Pointer to **string** | if source_address is configured, will use the vlan firstly otherwise use source_ip | [optional] 
**SendToAllServers** | Pointer to **bool** |  | [optional] [default to false]
**RemoteSyslogServers** | Pointer to [**[]SyslogServer**](SyslogServer.md) |  | [optional] 
**TimeFormat** | Pointer to **string** |  | [optional] 
**RemoteSyslogUsers** | Pointer to [**[]RemoteSyslogUser**](RemoteSyslogUser.md) |  | [optional] 

## Methods

### NewRemoteSyslog

`func NewRemoteSyslog() *RemoteSyslog`

NewRemoteSyslog instantiates a new RemoteSyslog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogWithDefaults

`func NewRemoteSyslogWithDefaults() *RemoteSyslog`

NewRemoteSyslogWithDefaults instantiates a new RemoteSyslog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteSyslogArchive

`func (o *RemoteSyslog) GetRemoteSyslogArchive() RemoteSyslogRemoteSyslogArchive`

GetRemoteSyslogArchive returns the RemoteSyslogArchive field if non-nil, zero value otherwise.

### GetRemoteSyslogArchiveOk

`func (o *RemoteSyslog) GetRemoteSyslogArchiveOk() (*RemoteSyslogRemoteSyslogArchive, bool)`

GetRemoteSyslogArchiveOk returns a tuple with the RemoteSyslogArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSyslogArchive

`func (o *RemoteSyslog) SetRemoteSyslogArchive(v RemoteSyslogRemoteSyslogArchive)`

SetRemoteSyslogArchive sets RemoteSyslogArchive field to given value.

### HasRemoteSyslogArchive

`func (o *RemoteSyslog) HasRemoteSyslogArchive() bool`

HasRemoteSyslogArchive returns a boolean if a field has been set.

### GetRemoteSyslogConsole

`func (o *RemoteSyslog) GetRemoteSyslogConsole() RemoteSyslogRemoteSyslogConsole`

GetRemoteSyslogConsole returns the RemoteSyslogConsole field if non-nil, zero value otherwise.

### GetRemoteSyslogConsoleOk

`func (o *RemoteSyslog) GetRemoteSyslogConsoleOk() (*RemoteSyslogRemoteSyslogConsole, bool)`

GetRemoteSyslogConsoleOk returns a tuple with the RemoteSyslogConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSyslogConsole

`func (o *RemoteSyslog) SetRemoteSyslogConsole(v RemoteSyslogRemoteSyslogConsole)`

SetRemoteSyslogConsole sets RemoteSyslogConsole field to given value.

### HasRemoteSyslogConsole

`func (o *RemoteSyslog) HasRemoteSyslogConsole() bool`

HasRemoteSyslogConsole returns a boolean if a field has been set.

### GetEnabled

`func (o *RemoteSyslog) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RemoteSyslog) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RemoteSyslog) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RemoteSyslog) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRemoteSyslogFiles

`func (o *RemoteSyslog) GetRemoteSyslogFiles() []SyslogFileConfig`

GetRemoteSyslogFiles returns the RemoteSyslogFiles field if non-nil, zero value otherwise.

### GetRemoteSyslogFilesOk

`func (o *RemoteSyslog) GetRemoteSyslogFilesOk() (*[]SyslogFileConfig, bool)`

GetRemoteSyslogFilesOk returns a tuple with the RemoteSyslogFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSyslogFiles

`func (o *RemoteSyslog) SetRemoteSyslogFiles(v []SyslogFileConfig)`

SetRemoteSyslogFiles sets RemoteSyslogFiles field to given value.

### HasRemoteSyslogFiles

`func (o *RemoteSyslog) HasRemoteSyslogFiles() bool`

HasRemoteSyslogFiles returns a boolean if a field has been set.

### GetNetwork

`func (o *RemoteSyslog) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RemoteSyslog) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RemoteSyslog) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RemoteSyslog) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSendToAllServers

`func (o *RemoteSyslog) GetSendToAllServers() bool`

GetSendToAllServers returns the SendToAllServers field if non-nil, zero value otherwise.

### GetSendToAllServersOk

`func (o *RemoteSyslog) GetSendToAllServersOk() (*bool, bool)`

GetSendToAllServersOk returns a tuple with the SendToAllServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToAllServers

`func (o *RemoteSyslog) SetSendToAllServers(v bool)`

SetSendToAllServers sets SendToAllServers field to given value.

### HasSendToAllServers

`func (o *RemoteSyslog) HasSendToAllServers() bool`

HasSendToAllServers returns a boolean if a field has been set.

### GetRemoteSyslogServers

`func (o *RemoteSyslog) GetRemoteSyslogServers() []SyslogServer`

GetRemoteSyslogServers returns the RemoteSyslogServers field if non-nil, zero value otherwise.

### GetRemoteSyslogServersOk

`func (o *RemoteSyslog) GetRemoteSyslogServersOk() (*[]SyslogServer, bool)`

GetRemoteSyslogServersOk returns a tuple with the RemoteSyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSyslogServers

`func (o *RemoteSyslog) SetRemoteSyslogServers(v []SyslogServer)`

SetRemoteSyslogServers sets RemoteSyslogServers field to given value.

### HasRemoteSyslogServers

`func (o *RemoteSyslog) HasRemoteSyslogServers() bool`

HasRemoteSyslogServers returns a boolean if a field has been set.

### GetTimeFormat

`func (o *RemoteSyslog) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *RemoteSyslog) GetTimeFormatOk() (*string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *RemoteSyslog) SetTimeFormat(v string)`

SetTimeFormat sets TimeFormat field to given value.

### HasTimeFormat

`func (o *RemoteSyslog) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### GetRemoteSyslogUsers

`func (o *RemoteSyslog) GetRemoteSyslogUsers() []RemoteSyslogUser`

GetRemoteSyslogUsers returns the RemoteSyslogUsers field if non-nil, zero value otherwise.

### GetRemoteSyslogUsersOk

`func (o *RemoteSyslog) GetRemoteSyslogUsersOk() (*[]RemoteSyslogUser, bool)`

GetRemoteSyslogUsersOk returns a tuple with the RemoteSyslogUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSyslogUsers

`func (o *RemoteSyslog) SetRemoteSyslogUsers(v []RemoteSyslogUser)`

SetRemoteSyslogUsers sets RemoteSyslogUsers field to given value.

### HasRemoteSyslogUsers

`func (o *RemoteSyslog) HasRemoteSyslogUsers() bool`

HasRemoteSyslogUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


