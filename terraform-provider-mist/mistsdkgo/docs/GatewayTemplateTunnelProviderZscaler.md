# GatewayTemplateTunnelProviderZscaler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AupAcceptanceRequired** | Pointer to **bool** |  | [optional] [default to true]
**AupExpire** | Pointer to **int32** | days before AUP is requested again | [optional] [default to 1]
**AupSslProxy** | Pointer to **bool** | proxy HTTPs traffic, requiring Zscaler cert to be installed in browser | [optional] [default to false]
**DownloadMbps** | Pointer to **int32** | the download bandwidth cap of the link, in Mbps | [optional] 
**EnableAup** | Pointer to **bool** | if &#x60;use_xff&#x60;&#x3D;&#x3D;&#x60;true&#x60;, display Acceptable Use Policy (AUP) | [optional] [default to false]
**EnableCaution** | Pointer to **bool** | when &#x60;enforce_authentication&#x60;&#x3D;&#x3D;&#x60;false&#x60;, display caution notification for non-authenticated users | [optional] [default to false]
**EnforceAuthentication** | Pointer to **bool** |  | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 
**SubLocations** | Pointer to [**[]GatewayTemplateTunnelProviderZscalerSubLocationsInner**](GatewayTemplateTunnelProviderZscalerSubLocationsInner.md) | if &#x60;use_xff&#x60;&#x3D;&#x3D;&#x60;true&#x60; | [optional] 
**UploadMbps** | Pointer to **int32** | the download bandwidth cap of the link, in Mbps | [optional] 
**UseXff** | Pointer to **bool** | location uses proxy chaining to forward traffic | [optional] 

## Methods

### NewGatewayTemplateTunnelProviderZscaler

`func NewGatewayTemplateTunnelProviderZscaler() *GatewayTemplateTunnelProviderZscaler`

NewGatewayTemplateTunnelProviderZscaler instantiates a new GatewayTemplateTunnelProviderZscaler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelProviderZscalerWithDefaults

`func NewGatewayTemplateTunnelProviderZscalerWithDefaults() *GatewayTemplateTunnelProviderZscaler`

NewGatewayTemplateTunnelProviderZscalerWithDefaults instantiates a new GatewayTemplateTunnelProviderZscaler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAupAcceptanceRequired

`func (o *GatewayTemplateTunnelProviderZscaler) GetAupAcceptanceRequired() bool`

GetAupAcceptanceRequired returns the AupAcceptanceRequired field if non-nil, zero value otherwise.

### GetAupAcceptanceRequiredOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetAupAcceptanceRequiredOk() (*bool, bool)`

GetAupAcceptanceRequiredOk returns a tuple with the AupAcceptanceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupAcceptanceRequired

`func (o *GatewayTemplateTunnelProviderZscaler) SetAupAcceptanceRequired(v bool)`

SetAupAcceptanceRequired sets AupAcceptanceRequired field to given value.

### HasAupAcceptanceRequired

`func (o *GatewayTemplateTunnelProviderZscaler) HasAupAcceptanceRequired() bool`

HasAupAcceptanceRequired returns a boolean if a field has been set.

### GetAupExpire

`func (o *GatewayTemplateTunnelProviderZscaler) GetAupExpire() int32`

GetAupExpire returns the AupExpire field if non-nil, zero value otherwise.

### GetAupExpireOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetAupExpireOk() (*int32, bool)`

GetAupExpireOk returns a tuple with the AupExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupExpire

`func (o *GatewayTemplateTunnelProviderZscaler) SetAupExpire(v int32)`

SetAupExpire sets AupExpire field to given value.

### HasAupExpire

`func (o *GatewayTemplateTunnelProviderZscaler) HasAupExpire() bool`

HasAupExpire returns a boolean if a field has been set.

### GetAupSslProxy

`func (o *GatewayTemplateTunnelProviderZscaler) GetAupSslProxy() bool`

GetAupSslProxy returns the AupSslProxy field if non-nil, zero value otherwise.

### GetAupSslProxyOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetAupSslProxyOk() (*bool, bool)`

GetAupSslProxyOk returns a tuple with the AupSslProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupSslProxy

`func (o *GatewayTemplateTunnelProviderZscaler) SetAupSslProxy(v bool)`

SetAupSslProxy sets AupSslProxy field to given value.

### HasAupSslProxy

`func (o *GatewayTemplateTunnelProviderZscaler) HasAupSslProxy() bool`

HasAupSslProxy returns a boolean if a field has been set.

### GetDownloadMbps

`func (o *GatewayTemplateTunnelProviderZscaler) GetDownloadMbps() int32`

GetDownloadMbps returns the DownloadMbps field if non-nil, zero value otherwise.

### GetDownloadMbpsOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetDownloadMbpsOk() (*int32, bool)`

GetDownloadMbpsOk returns a tuple with the DownloadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMbps

`func (o *GatewayTemplateTunnelProviderZscaler) SetDownloadMbps(v int32)`

SetDownloadMbps sets DownloadMbps field to given value.

### HasDownloadMbps

`func (o *GatewayTemplateTunnelProviderZscaler) HasDownloadMbps() bool`

HasDownloadMbps returns a boolean if a field has been set.

### GetEnableAup

`func (o *GatewayTemplateTunnelProviderZscaler) GetEnableAup() bool`

GetEnableAup returns the EnableAup field if non-nil, zero value otherwise.

### GetEnableAupOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetEnableAupOk() (*bool, bool)`

GetEnableAupOk returns a tuple with the EnableAup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAup

`func (o *GatewayTemplateTunnelProviderZscaler) SetEnableAup(v bool)`

SetEnableAup sets EnableAup field to given value.

### HasEnableAup

`func (o *GatewayTemplateTunnelProviderZscaler) HasEnableAup() bool`

HasEnableAup returns a boolean if a field has been set.

### GetEnableCaution

`func (o *GatewayTemplateTunnelProviderZscaler) GetEnableCaution() bool`

GetEnableCaution returns the EnableCaution field if non-nil, zero value otherwise.

### GetEnableCautionOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetEnableCautionOk() (*bool, bool)`

GetEnableCautionOk returns a tuple with the EnableCaution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaution

`func (o *GatewayTemplateTunnelProviderZscaler) SetEnableCaution(v bool)`

SetEnableCaution sets EnableCaution field to given value.

### HasEnableCaution

`func (o *GatewayTemplateTunnelProviderZscaler) HasEnableCaution() bool`

HasEnableCaution returns a boolean if a field has been set.

### GetEnforceAuthentication

`func (o *GatewayTemplateTunnelProviderZscaler) GetEnforceAuthentication() bool`

GetEnforceAuthentication returns the EnforceAuthentication field if non-nil, zero value otherwise.

### GetEnforceAuthenticationOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetEnforceAuthenticationOk() (*bool, bool)`

GetEnforceAuthenticationOk returns a tuple with the EnforceAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceAuthentication

`func (o *GatewayTemplateTunnelProviderZscaler) SetEnforceAuthentication(v bool)`

SetEnforceAuthentication sets EnforceAuthentication field to given value.

### HasEnforceAuthentication

`func (o *GatewayTemplateTunnelProviderZscaler) HasEnforceAuthentication() bool`

HasEnforceAuthentication returns a boolean if a field has been set.

### GetName

`func (o *GatewayTemplateTunnelProviderZscaler) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayTemplateTunnelProviderZscaler) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayTemplateTunnelProviderZscaler) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubLocations

`func (o *GatewayTemplateTunnelProviderZscaler) GetSubLocations() []GatewayTemplateTunnelProviderZscalerSubLocationsInner`

GetSubLocations returns the SubLocations field if non-nil, zero value otherwise.

### GetSubLocationsOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetSubLocationsOk() (*[]GatewayTemplateTunnelProviderZscalerSubLocationsInner, bool)`

GetSubLocationsOk returns a tuple with the SubLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLocations

`func (o *GatewayTemplateTunnelProviderZscaler) SetSubLocations(v []GatewayTemplateTunnelProviderZscalerSubLocationsInner)`

SetSubLocations sets SubLocations field to given value.

### HasSubLocations

`func (o *GatewayTemplateTunnelProviderZscaler) HasSubLocations() bool`

HasSubLocations returns a boolean if a field has been set.

### GetUploadMbps

`func (o *GatewayTemplateTunnelProviderZscaler) GetUploadMbps() int32`

GetUploadMbps returns the UploadMbps field if non-nil, zero value otherwise.

### GetUploadMbpsOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetUploadMbpsOk() (*int32, bool)`

GetUploadMbpsOk returns a tuple with the UploadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadMbps

`func (o *GatewayTemplateTunnelProviderZscaler) SetUploadMbps(v int32)`

SetUploadMbps sets UploadMbps field to given value.

### HasUploadMbps

`func (o *GatewayTemplateTunnelProviderZscaler) HasUploadMbps() bool`

HasUploadMbps returns a boolean if a field has been set.

### GetUseXff

`func (o *GatewayTemplateTunnelProviderZscaler) GetUseXff() bool`

GetUseXff returns the UseXff field if non-nil, zero value otherwise.

### GetUseXffOk

`func (o *GatewayTemplateTunnelProviderZscaler) GetUseXffOk() (*bool, bool)`

GetUseXffOk returns a tuple with the UseXff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseXff

`func (o *GatewayTemplateTunnelProviderZscaler) SetUseXff(v bool)`

SetUseXff sets UseXff field to given value.

### HasUseXff

`func (o *GatewayTemplateTunnelProviderZscaler) HasUseXff() bool`

HasUseXff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


