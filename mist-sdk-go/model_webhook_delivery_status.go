/*
Mist API

> Version: **2406.1.10** > > Date: **July 1, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.11
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistsdkgo

import (
	"encoding/json"
	"fmt"
)

// WebhookDeliveryStatus webhook delivery status
type WebhookDeliveryStatus string

// List of webhook_delivery_status
const (
	WEBHOOKDELIVERYSTATUS_EMPTY WebhookDeliveryStatus = ""
	WEBHOOKDELIVERYSTATUS_SUCCESS WebhookDeliveryStatus = "success"
	WEBHOOKDELIVERYSTATUS_FAILURE WebhookDeliveryStatus = "failure"
)

// All allowed values of WebhookDeliveryStatus enum
var AllowedWebhookDeliveryStatusEnumValues = []WebhookDeliveryStatus{
	"",
	"success",
	"failure",
}

func (v *WebhookDeliveryStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookDeliveryStatus(value)
	for _, existing := range AllowedWebhookDeliveryStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookDeliveryStatus", value)
}

// NewWebhookDeliveryStatusFromValue returns a pointer to a valid WebhookDeliveryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookDeliveryStatusFromValue(v string) (*WebhookDeliveryStatus, error) {
	ev := WebhookDeliveryStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookDeliveryStatus: valid values are %v", v, AllowedWebhookDeliveryStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookDeliveryStatus) IsValid() bool {
	for _, existing := range AllowedWebhookDeliveryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to webhook_delivery_status value
func (v WebhookDeliveryStatus) Ptr() *WebhookDeliveryStatus {
	return &v
}

type NullableWebhookDeliveryStatus struct {
	value *WebhookDeliveryStatus
	isSet bool
}

func (v NullableWebhookDeliveryStatus) Get() *WebhookDeliveryStatus {
	return v.value
}

func (v *NullableWebhookDeliveryStatus) Set(val *WebhookDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDeliveryStatus(val *WebhookDeliveryStatus) *NullableWebhookDeliveryStatus {
	return &NullableWebhookDeliveryStatus{value: val, isSet: true}
}

func (v NullableWebhookDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
