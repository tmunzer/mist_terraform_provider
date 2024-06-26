# WebhookOccupancyAlertsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEvents** | Pointer to [**[]WebhookOccupancyAlertsEventAlertEventsItems**](WebhookOccupancyAlertsEventAlertEventsItems.md) | list of occupancy alerts for non-compliance zones within the site detected around the same time | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**SiteId** | **string** |  | [readonly] 
**SiteName** | **string** |  | [readonly] 

## Methods

### NewWebhookOccupancyAlertsEvent

`func NewWebhookOccupancyAlertsEvent(siteId string, siteName string, ) *WebhookOccupancyAlertsEvent`

NewWebhookOccupancyAlertsEvent instantiates a new WebhookOccupancyAlertsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOccupancyAlertsEventWithDefaults

`func NewWebhookOccupancyAlertsEventWithDefaults() *WebhookOccupancyAlertsEvent`

NewWebhookOccupancyAlertsEventWithDefaults instantiates a new WebhookOccupancyAlertsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEvents

`func (o *WebhookOccupancyAlertsEvent) GetAlertEvents() []WebhookOccupancyAlertsEventAlertEventsItems`

GetAlertEvents returns the AlertEvents field if non-nil, zero value otherwise.

### GetAlertEventsOk

`func (o *WebhookOccupancyAlertsEvent) GetAlertEventsOk() (*[]WebhookOccupancyAlertsEventAlertEventsItems, bool)`

GetAlertEventsOk returns a tuple with the AlertEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEvents

`func (o *WebhookOccupancyAlertsEvent) SetAlertEvents(v []WebhookOccupancyAlertsEventAlertEventsItems)`

SetAlertEvents sets AlertEvents field to given value.

### HasAlertEvents

`func (o *WebhookOccupancyAlertsEvent) HasAlertEvents() bool`

HasAlertEvents returns a boolean if a field has been set.

### GetForSite

`func (o *WebhookOccupancyAlertsEvent) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *WebhookOccupancyAlertsEvent) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *WebhookOccupancyAlertsEvent) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *WebhookOccupancyAlertsEvent) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookOccupancyAlertsEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookOccupancyAlertsEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookOccupancyAlertsEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSiteName

`func (o *WebhookOccupancyAlertsEvent) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *WebhookOccupancyAlertsEvent) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *WebhookOccupancyAlertsEvent) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


