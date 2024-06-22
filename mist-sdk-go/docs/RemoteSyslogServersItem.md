# RemoteSyslogServersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to [**[]RemoteSyslogContentItem**](RemoteSyslogContentItem.md) |  | [optional] 
**ExplicitPriority** | Pointer to **bool** |  | [optional] 
**Facility** | Pointer to [**RemoteSyslogFacility**](RemoteSyslogFacility.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Match** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] [default to 514]
**Protocol** | Pointer to [**RemoteSyslogProtocol**](RemoteSyslogProtocol.md) |  | [optional] [default to REMOTESYSLOGPROTOCOL_UDP]
**RoutingInstance** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to [**RemoteSyslogSeverity**](RemoteSyslogSeverity.md) |  | [optional] 
**SourceAddress** | Pointer to **string** | if source_address is configured, will use the vlan firstly otherwise use source_ip | [optional] 
**StructuredData** | Pointer to **bool** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewRemoteSyslogServersItem

`func NewRemoteSyslogServersItem() *RemoteSyslogServersItem`

NewRemoteSyslogServersItem instantiates a new RemoteSyslogServersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogServersItemWithDefaults

`func NewRemoteSyslogServersItemWithDefaults() *RemoteSyslogServersItem`

NewRemoteSyslogServersItemWithDefaults instantiates a new RemoteSyslogServersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *RemoteSyslogServersItem) GetContents() []RemoteSyslogContentItem`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *RemoteSyslogServersItem) GetContentsOk() (*[]RemoteSyslogContentItem, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *RemoteSyslogServersItem) SetContents(v []RemoteSyslogContentItem)`

SetContents sets Contents field to given value.

### HasContents

`func (o *RemoteSyslogServersItem) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetExplicitPriority

`func (o *RemoteSyslogServersItem) GetExplicitPriority() bool`

GetExplicitPriority returns the ExplicitPriority field if non-nil, zero value otherwise.

### GetExplicitPriorityOk

`func (o *RemoteSyslogServersItem) GetExplicitPriorityOk() (*bool, bool)`

GetExplicitPriorityOk returns a tuple with the ExplicitPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPriority

`func (o *RemoteSyslogServersItem) SetExplicitPriority(v bool)`

SetExplicitPriority sets ExplicitPriority field to given value.

### HasExplicitPriority

`func (o *RemoteSyslogServersItem) HasExplicitPriority() bool`

HasExplicitPriority returns a boolean if a field has been set.

### GetFacility

`func (o *RemoteSyslogServersItem) GetFacility() RemoteSyslogFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *RemoteSyslogServersItem) GetFacilityOk() (*RemoteSyslogFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *RemoteSyslogServersItem) SetFacility(v RemoteSyslogFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *RemoteSyslogServersItem) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHost

`func (o *RemoteSyslogServersItem) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RemoteSyslogServersItem) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RemoteSyslogServersItem) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RemoteSyslogServersItem) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMatch

`func (o *RemoteSyslogServersItem) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *RemoteSyslogServersItem) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *RemoteSyslogServersItem) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *RemoteSyslogServersItem) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetPort

`func (o *RemoteSyslogServersItem) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RemoteSyslogServersItem) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RemoteSyslogServersItem) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RemoteSyslogServersItem) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *RemoteSyslogServersItem) GetProtocol() RemoteSyslogProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RemoteSyslogServersItem) GetProtocolOk() (*RemoteSyslogProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RemoteSyslogServersItem) SetProtocol(v RemoteSyslogProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RemoteSyslogServersItem) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRoutingInstance

`func (o *RemoteSyslogServersItem) GetRoutingInstance() string`

GetRoutingInstance returns the RoutingInstance field if non-nil, zero value otherwise.

### GetRoutingInstanceOk

`func (o *RemoteSyslogServersItem) GetRoutingInstanceOk() (*string, bool)`

GetRoutingInstanceOk returns a tuple with the RoutingInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingInstance

`func (o *RemoteSyslogServersItem) SetRoutingInstance(v string)`

SetRoutingInstance sets RoutingInstance field to given value.

### HasRoutingInstance

`func (o *RemoteSyslogServersItem) HasRoutingInstance() bool`

HasRoutingInstance returns a boolean if a field has been set.

### GetSeverity

`func (o *RemoteSyslogServersItem) GetSeverity() RemoteSyslogSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RemoteSyslogServersItem) GetSeverityOk() (*RemoteSyslogSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RemoteSyslogServersItem) SetSeverity(v RemoteSyslogSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RemoteSyslogServersItem) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSourceAddress

`func (o *RemoteSyslogServersItem) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *RemoteSyslogServersItem) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *RemoteSyslogServersItem) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *RemoteSyslogServersItem) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### GetStructuredData

`func (o *RemoteSyslogServersItem) GetStructuredData() bool`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *RemoteSyslogServersItem) GetStructuredDataOk() (*bool, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *RemoteSyslogServersItem) SetStructuredData(v bool)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *RemoteSyslogServersItem) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.

### GetTag

`func (o *RemoteSyslogServersItem) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RemoteSyslogServersItem) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RemoteSyslogServersItem) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RemoteSyslogServersItem) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


