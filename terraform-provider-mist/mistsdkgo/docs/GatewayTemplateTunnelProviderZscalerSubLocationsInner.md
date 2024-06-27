# GatewayTemplateTunnelProviderZscalerSubLocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AupAcceptanceRequired** | Pointer to **bool** |  | [optional] [default to true]
**AupExpire** | Pointer to **int32** | days before AUP is requested again | [optional] [default to 1]
**AupSslProxy** | Pointer to **bool** | proxy HTTPs traffic, requiring Zscaler cert to be installed in browser | [optional] [default to false]
**DownloadMbps** | Pointer to **int32** | the download bandwidth cap of the link, in Mbps | [optional] 
**EnableAup** | Pointer to **bool** | if &#x60;use_xff&#x60;&#x3D;&#x3D;&#x60;true&#x60;, display Acceptable Use Policy (AUP) | [optional] 
**EnableCaution** | Pointer to **bool** | when &#x60;enforce_authentication&#x60;&#x3D;&#x3D;&#x60;false&#x60;, display caution notification for non-authenticated users | [optional] [default to false]
**EnforceAuthentication** | Pointer to **bool** |  | [optional] [default to false]
**Subnets** | Pointer to **[]string** |  | [optional] 
**UploadMbps** | Pointer to **int32** | the download bandwidth cap of the link, in Mbps | [optional] 

## Methods

### NewGatewayTemplateTunnelProviderZscalerSubLocationsInner

`func NewGatewayTemplateTunnelProviderZscalerSubLocationsInner() *GatewayTemplateTunnelProviderZscalerSubLocationsInner`

NewGatewayTemplateTunnelProviderZscalerSubLocationsInner instantiates a new GatewayTemplateTunnelProviderZscalerSubLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelProviderZscalerSubLocationsInnerWithDefaults

`func NewGatewayTemplateTunnelProviderZscalerSubLocationsInnerWithDefaults() *GatewayTemplateTunnelProviderZscalerSubLocationsInner`

NewGatewayTemplateTunnelProviderZscalerSubLocationsInnerWithDefaults instantiates a new GatewayTemplateTunnelProviderZscalerSubLocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAupAcceptanceRequired

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetAupAcceptanceRequired() bool`

GetAupAcceptanceRequired returns the AupAcceptanceRequired field if non-nil, zero value otherwise.

### GetAupAcceptanceRequiredOk

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetAupAcceptanceRequiredOk() (*bool, bool)`

GetAupAcceptanceRequiredOk returns a tuple with the AupAcceptanceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupAcceptanceRequired

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) SetAupAcceptanceRequired(v bool)`

SetAupAcceptanceRequired sets AupAcceptanceRequired field to given value.

### HasAupAcceptanceRequired

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) HasAupAcceptanceRequired() bool`

HasAupAcceptanceRequired returns a boolean if a field has been set.

### GetAupExpire

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetAupExpire() int32`

GetAupExpire returns the AupExpire field if non-nil, zero value otherwise.

### GetAupExpireOk

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetAupExpireOk() (*int32, bool)`

GetAupExpireOk returns a tuple with the AupExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupExpire

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) SetAupExpire(v int32)`

SetAupExpire sets AupExpire field to given value.

### HasAupExpire

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) HasAupExpire() bool`

HasAupExpire returns a boolean if a field has been set.

### GetAupSslProxy

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetAupSslProxy() bool`

GetAupSslProxy returns the AupSslProxy field if non-nil, zero value otherwise.

### GetAupSslProxyOk

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetAupSslProxyOk() (*bool, bool)`

GetAupSslProxyOk returns a tuple with the AupSslProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupSslProxy

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) SetAupSslProxy(v bool)`

SetAupSslProxy sets AupSslProxy field to given value.

### HasAupSslProxy

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) HasAupSslProxy() bool`

HasAupSslProxy returns a boolean if a field has been set.

### GetDownloadMbps

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetDownloadMbps() int32`

GetDownloadMbps returns the DownloadMbps field if non-nil, zero value otherwise.

### GetDownloadMbpsOk

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetDownloadMbpsOk() (*int32, bool)`

GetDownloadMbpsOk returns a tuple with the DownloadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMbps

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) SetDownloadMbps(v int32)`

SetDownloadMbps sets DownloadMbps field to given value.

### HasDownloadMbps

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) HasDownloadMbps() bool`

HasDownloadMbps returns a boolean if a field has been set.

### GetEnableAup

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetEnableAup() bool`

GetEnableAup returns the EnableAup field if non-nil, zero value otherwise.

### GetEnableAupOk

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetEnableAupOk() (*bool, bool)`

GetEnableAupOk returns a tuple with the EnableAup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAup

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) SetEnableAup(v bool)`

SetEnableAup sets EnableAup field to given value.

### HasEnableAup

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) HasEnableAup() bool`

HasEnableAup returns a boolean if a field has been set.

### GetEnableCaution

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetEnableCaution() bool`

GetEnableCaution returns the EnableCaution field if non-nil, zero value otherwise.

### GetEnableCautionOk

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetEnableCautionOk() (*bool, bool)`

GetEnableCautionOk returns a tuple with the EnableCaution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaution

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) SetEnableCaution(v bool)`

SetEnableCaution sets EnableCaution field to given value.

### HasEnableCaution

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) HasEnableCaution() bool`

HasEnableCaution returns a boolean if a field has been set.

### GetEnforceAuthentication

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetEnforceAuthentication() bool`

GetEnforceAuthentication returns the EnforceAuthentication field if non-nil, zero value otherwise.

### GetEnforceAuthenticationOk

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetEnforceAuthenticationOk() (*bool, bool)`

GetEnforceAuthenticationOk returns a tuple with the EnforceAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceAuthentication

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) SetEnforceAuthentication(v bool)`

SetEnforceAuthentication sets EnforceAuthentication field to given value.

### HasEnforceAuthentication

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) HasEnforceAuthentication() bool`

HasEnforceAuthentication returns a boolean if a field has been set.

### GetSubnets

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetUploadMbps

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetUploadMbps() int32`

GetUploadMbps returns the UploadMbps field if non-nil, zero value otherwise.

### GetUploadMbpsOk

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) GetUploadMbpsOk() (*int32, bool)`

GetUploadMbpsOk returns a tuple with the UploadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadMbps

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) SetUploadMbps(v int32)`

SetUploadMbps sets UploadMbps field to given value.

### HasUploadMbps

`func (o *GatewayTemplateTunnelProviderZscalerSubLocationsInner) HasUploadMbps() bool`

HasUploadMbps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


