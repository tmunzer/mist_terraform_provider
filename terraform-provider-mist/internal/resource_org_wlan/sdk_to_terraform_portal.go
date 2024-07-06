package resource_org_wlan

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func portalSkToTerraform(ctx context.Context, diags *diag.Diagnostics, data models.WlanPortal) PortalValue {
	sponsors_attr := make(map[string]attr.Value)
	for k, v := range data.GetSponsors() {
		sponsors_attr[k] = types.StringValue(string(v))
	}
	sponsors, e := types.MapValueFrom(ctx, types.StringType, sponsors_attr)
	diags.Append(e...)

	plan_attr := map[string]attr.Value{
		"amazon_client_id":               types.StringValue(data.GetAmazonClientId()),
		"amazon_client_secret":           types.StringValue(data.GetAmazonClientSecret()),
		"amazon_email_domains":           mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAmazonEmailDomains()),
		"amazon_enabled":                 types.BoolValue(data.GetAmazonEnabled()),
		"amazon_expire":                  types.Float64Value(data.GetAmazonExpire()),
		"auth":                           types.StringValue(string(data.GetAuth())),
		"azure_client_id":                types.StringValue(data.GetAzureClientId()),
		"azure_client_secret":            types.StringValue(data.GetAzureClientSecret()),
		"azure_enabled":                  types.BoolValue(data.GetAzureEnabled()),
		"azure_expire":                   types.Float64Value(data.GetAzureExpire()),
		"azure_tenant_id":                types.StringValue(data.GetAzureTenantId()),
		"broadnet_password":              types.StringValue(data.GetBroadnetPassword()),
		"broadnet_sid":                   types.StringValue(data.GetBroadnetSid()),
		"broadnet_user_id":               types.StringValue(data.GetBroadnetUserId()),
		"bypass_when_cloud_down":         types.BoolValue(data.GetBypassWhenCloudDown()),
		"clickatell_api_key":             types.StringValue(data.GetClickatellApiKey()),
		"cross_site":                     types.BoolValue(data.GetCrossSite()),
		"email_enabled":                  types.BoolValue(data.GetEmailEnabled()),
		"enabled":                        types.BoolValue(data.GetEnabled()),
		"expire":                         types.Float64Value(data.GetExpire()),
		"external_portal_url":            types.StringValue(data.GetExternalPortalUrl()),
		"facebook_client_id":             types.StringValue(data.GetFacebookClientId()),
		"facebook_client_secret":         types.StringValue(data.GetFacebookClientSecret()),
		"facebook_email_domains":         mist_transform.ListOfStringSdkToTerraform(ctx, data.GetFacebookEmailDomains()),
		"facebook_enabled":               types.BoolValue(data.GetFacebookEnabled()),
		"facebook_expire":                types.Float64Value(data.GetFacebookExpire()),
		"forward":                        types.BoolValue(data.GetForward()),
		"forward_url":                    types.StringValue(data.GetForwardUrl()),
		"google_client_id":               types.StringValue(data.GetGoogleClientId()),
		"google_client_secret":           types.StringValue(data.GetGoogleClientSecret()),
		"google_email_domains":           mist_transform.ListOfStringSdkToTerraform(ctx, data.GetGoogleEmailDomains()),
		"google_enabled":                 types.BoolValue(data.GetGoogleEnabled()),
		"google_expire":                  types.Float64Value(data.GetGoogleExpire()),
		"gupshup_password":               types.StringValue(data.GetGupshupPassword()),
		"gupshup_userid":                 types.StringValue(data.GetGupshupUserid()),
		"microsoft_client_id":            types.StringValue(data.GetMicrosoftClientId()),
		"microsoft_client_secret":        types.StringValue(data.GetMicrosoftClientSecret()),
		"microsoft_email_domains":        mist_transform.ListOfStringSdkToTerraform(ctx, data.GetMicrosoftEmailDomains()),
		"microsoft_enabled":              types.BoolValue(data.GetMicrosoftEnabled()),
		"microsoft_expire":               types.Float64Value(data.GetMicrosoftExpire()),
		"passphrase_enabled":             types.BoolValue(data.GetPassphraseEnabled()),
		"passphrase_expire":              types.Float64Value(data.GetPassphraseExpire()),
		"password":                       types.StringValue(data.GetPassword()),
		"portal_allowed_hostnames":       types.StringValue(data.GetPortalAllowedHostnames()),
		"portal_allowed_subnets":         types.StringValue(data.GetPortalAllowedSubnets()),
		"portal_api_secret":              types.StringValue(data.GetPortalApiSecret()),
		"portal_denied_hostnames":        types.StringValue(data.GetPortalDeniedHostnames()),
		"portal_image":                   types.StringValue(data.GetPortalImage()),
		"portal_sso_url":                 types.StringValue(data.GetPortalSsoUrl()),
		"predefined_sponsors_enabled":    types.BoolValue(data.GetPredefinedSponsorsEnabled()),
		"privacy":                        types.BoolValue(data.GetPrivacy()),
		"puzzel_password":                types.StringValue(data.GetPuzzelPassword()),
		"puzzel_service_id":              types.StringValue(data.GetPuzzelServiceId()),
		"puzzel_username":                types.StringValue(data.GetPuzzelUsername()),
		"sms_enabled":                    types.BoolValue(data.GetSmsEnabled()),
		"sms_expire":                     types.Float64Value(data.GetSmsExpire()),
		"sms_message_format":             types.StringValue(data.GetSmsMessageFormat()),
		"sms_provider":                   types.StringValue(string(data.GetSmsProvider())),
		"sponsor_auto_approve":           types.BoolValue(data.GetSponsorAutoApprove()),
		"sponsor_email_domains":          mist_transform.ListOfStringSdkToTerraform(ctx, data.GetSponsorEmailDomains()),
		"sponsor_enabled":                types.BoolValue(data.GetSponsorEnabled()),
		"sponsor_expire":                 types.Float64Value(data.GetSponsorExpire()),
		"sponsor_link_validity_duration": types.Int64Value(int64(data.GetSponsorLinkValidityDuration())),
		"sponsor_notify_all":             types.BoolValue(data.GetSponsorNotifyAll()),
		"sponsor_status_notify":          types.BoolValue(data.GetSponsorStatusNotify()),
		"sponsors":                       sponsors,
		"sso_default_role":               types.StringValue(data.GetSsoDefaultRole()),
		"sso_forced_role":                types.StringValue(data.GetSsoForcedRole()),
		"sso_idp_cert":                   types.StringValue(data.GetSsoIdpCert()),
		"sso_idp_sign_algo":              types.StringValue(data.GetSsoIdpSignAlgo()),
		"sso_idp_sso_url":                types.StringValue(data.GetSsoIdpSsoUrl()),
		"sso_issuer":                     types.StringValue(data.GetSsoIssuer()),
		"sso_nameid_format":              types.StringValue(string(data.GetSsoNameidFormat())),
		"telstra_client_id":              types.StringValue(data.GetTelstraClientId()),
		"telstra_client_secret":          types.StringValue(data.GetTelstraClientSecret()),
		"thumbnail":                      types.StringValue(data.GetThumbnail()),
		"twilio_auth_token":              types.StringValue(data.GetTwilioAuthToken()),
		"twilio_phone_number":            types.StringValue(data.GetTwilioPhoneNumber()),
		"twilio_sid":                     types.StringValue(data.GetTwilioSid()),
	}

	r, e := NewPortalValue(PortalValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
