# SyslogServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to [**[]RemoteSyslogContent**](RemoteSyslogContent.md) |  | [optional] 
**ExplicitPriority** | Pointer to **bool** |  | [optional] 
**Facility** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Match** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] [default to 514]
**Protocol** | Pointer to **string** |  | [optional] [default to "udp"]
**RoutingInstance** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**SourceAddress** | Pointer to **string** | if source_address is configured, will use the vlan firstly otherwise use source_ip | [optional] 
**StructuredData** | Pointer to **bool** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewSyslogServer

`func NewSyslogServer() *SyslogServer`

NewSyslogServer instantiates a new SyslogServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogServerWithDefaults

`func NewSyslogServerWithDefaults() *SyslogServer`

NewSyslogServerWithDefaults instantiates a new SyslogServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *SyslogServer) GetContents() []RemoteSyslogContent`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *SyslogServer) GetContentsOk() (*[]RemoteSyslogContent, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *SyslogServer) SetContents(v []RemoteSyslogContent)`

SetContents sets Contents field to given value.

### HasContents

`func (o *SyslogServer) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetExplicitPriority

`func (o *SyslogServer) GetExplicitPriority() bool`

GetExplicitPriority returns the ExplicitPriority field if non-nil, zero value otherwise.

### GetExplicitPriorityOk

`func (o *SyslogServer) GetExplicitPriorityOk() (*bool, bool)`

GetExplicitPriorityOk returns a tuple with the ExplicitPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPriority

`func (o *SyslogServer) SetExplicitPriority(v bool)`

SetExplicitPriority sets ExplicitPriority field to given value.

### HasExplicitPriority

`func (o *SyslogServer) HasExplicitPriority() bool`

HasExplicitPriority returns a boolean if a field has been set.

### GetFacility

`func (o *SyslogServer) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *SyslogServer) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *SyslogServer) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *SyslogServer) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHost

`func (o *SyslogServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SyslogServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SyslogServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SyslogServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMatch

`func (o *SyslogServer) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *SyslogServer) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *SyslogServer) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *SyslogServer) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetPort

`func (o *SyslogServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SyslogServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SyslogServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SyslogServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *SyslogServer) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SyslogServer) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SyslogServer) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SyslogServer) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRoutingInstance

`func (o *SyslogServer) GetRoutingInstance() string`

GetRoutingInstance returns the RoutingInstance field if non-nil, zero value otherwise.

### GetRoutingInstanceOk

`func (o *SyslogServer) GetRoutingInstanceOk() (*string, bool)`

GetRoutingInstanceOk returns a tuple with the RoutingInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingInstance

`func (o *SyslogServer) SetRoutingInstance(v string)`

SetRoutingInstance sets RoutingInstance field to given value.

### HasRoutingInstance

`func (o *SyslogServer) HasRoutingInstance() bool`

HasRoutingInstance returns a boolean if a field has been set.

### GetSeverity

`func (o *SyslogServer) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SyslogServer) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SyslogServer) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SyslogServer) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSourceAddress

`func (o *SyslogServer) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *SyslogServer) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *SyslogServer) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *SyslogServer) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### GetStructuredData

`func (o *SyslogServer) GetStructuredData() bool`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *SyslogServer) GetStructuredDataOk() (*bool, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *SyslogServer) SetStructuredData(v bool)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *SyslogServer) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.

### GetTag

`func (o *SyslogServer) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SyslogServer) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SyslogServer) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *SyslogServer) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

