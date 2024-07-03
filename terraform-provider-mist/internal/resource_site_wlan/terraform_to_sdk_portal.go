package resource_site_wlan

import (
	"context"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func portalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan PortalValue) mistsdkgo.WlanPortal {

	sponsors := make(map[string]string)
	for k, v := range plan.Sponsors.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		sponsors[k] = v_plan.ValueString()
	}

	data := *mistsdkgo.NewWlanPortal()
	data.SetAmazonClientId(plan.AmazonClientId.ValueString())

	data.SetAmazonClientId(plan.AmazonClientId.ValueString())
	data.SetAmazonClientSecret(plan.AmazonClientSecret.ValueString())
	data.SetAmazonEmailDomains(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AmazonEmailDomains))
	data.SetAmazonEnabled(plan.AmazonEnabled.ValueBool())
	data.SetAmazonExpire(plan.AmazonExpire.ValueFloat64())
	data.SetAuth(mistsdkgo.WlanPortalAuth(plan.Auth.ValueString()))
	data.SetAzureClientId(plan.AzureClientId.ValueString())
	data.SetAzureClientSecret(plan.AzureClientSecret.ValueString())
	data.SetAzureEnabled(plan.AzureEnabled.ValueBool())
	data.SetAzureExpire(plan.AzureExpire.ValueFloat64())
	data.SetAzureTenantId(plan.AzureTenantId.ValueString())
	data.SetBroadnetPassword(plan.BroadnetPassword.ValueString())
	data.SetBroadnetSid(plan.BroadnetSid.ValueString())
	data.SetBroadnetUserId(plan.BroadnetUserId.ValueString())
	data.SetBypassWhenCloudDown(plan.BypassWhenCloudDown.ValueBool())
	data.SetClickatellApiKey(plan.ClickatellApiKey.ValueString())
	data.SetCrossSite(plan.CrossSite.ValueBool())
	data.SetEmailEnabled(plan.EmailEnabled.ValueBool())
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetExpire(plan.Expire.ValueFloat64())
	data.SetExternalPortalUrl(plan.ExternalPortalUrl.ValueString())
	data.SetFacebookClientId(plan.FacebookClientId.ValueString())
	data.SetFacebookClientSecret(plan.FacebookClientSecret.ValueString())
	data.SetFacebookEmailDomains(mist_transform.ListOfStringTerraformToSdk(ctx, plan.FacebookEmailDomains))
	data.SetFacebookEnabled(plan.FacebookEnabled.ValueBool())
	data.SetFacebookExpire(plan.FacebookExpire.ValueFloat64())
	data.SetForward(plan.Forward.ValueBool())
	data.SetForwardUrl(plan.ForwardUrl.ValueString())
	data.SetGoogleClientId(plan.GoogleClientId.ValueString())
	data.SetGoogleClientSecret(plan.GoogleClientSecret.ValueString())
	data.SetGoogleEmailDomains(mist_transform.ListOfStringTerraformToSdk(ctx, plan.GoogleEmailDomains))
	data.SetGoogleEnabled(plan.GoogleEnabled.ValueBool())
	data.SetGoogleExpire(plan.GoogleExpire.ValueFloat64())
	data.SetGupshupPassword(plan.GupshupPassword.ValueString())
	data.SetGupshupUserid(plan.GupshupUserid.ValueString())
	data.SetMicrosoftClientId(plan.MicrosoftClientId.ValueString())
	data.SetMicrosoftClientSecret(plan.MicrosoftClientSecret.ValueString())
	data.SetMicrosoftEmailDomains(mist_transform.ListOfStringTerraformToSdk(ctx, plan.MicrosoftEmailDomains))
	data.SetMicrosoftEnabled(plan.MicrosoftEnabled.ValueBool())
	data.SetMicrosoftExpire(plan.MicrosoftExpire.ValueFloat64())
	data.SetPassphraseEnabled(plan.PassphraseEnabled.ValueBool())
	data.SetPassphraseExpire(plan.PassphraseExpire.ValueFloat64())
	data.SetPassword(plan.Password.ValueString())
	data.SetPortalAllowedHostnames(plan.PortalAllowedHostnames.ValueString())
	data.SetPortalAllowedSubnets(plan.PortalAllowedSubnets.ValueString())
	data.SetPortalApiSecret(plan.PortalApiSecret.ValueString())
	data.SetPortalDeniedHostnames(plan.PortalDeniedHostnames.ValueString())
	data.SetPortalImage(plan.PortalImage.ValueString())
	data.SetPortalSsoUrl(plan.PortalSsoUrl.ValueString())
	data.SetPredefinedSponsorsEnabled(plan.PredefinedSponsorsEnabled.ValueBool())
	data.SetPrivacy(plan.Privacy.ValueBool())
	data.SetPuzzelPassword(plan.PuzzelPassword.ValueString())
	data.SetPuzzelServiceId(plan.PuzzelServiceId.ValueString())
	data.SetPuzzelUsername(plan.PuzzelUsername.ValueString())
	data.SetSmsEnabled(plan.SmsEnabled.ValueBool())
	data.SetSmsExpire(plan.SmsExpire.ValueFloat64())
	data.SetSmsMessageFormat(plan.SmsMessageFormat.ValueString())
	data.SetSmsProvider(mistsdkgo.WlanPortalSmsProvider(plan.SmsProvider.ValueString()))
	data.SetSponsorAutoApprove(plan.SponsorAutoApprove.ValueBool())
	data.SetSponsorEmailDomains(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SponsorEmailDomains))
	data.SetSponsorEnabled(plan.SponsorEnabled.ValueBool())
	data.SetSponsorExpire(plan.SponsorExpire.ValueFloat64())
	data.SetSponsorLinkValidityDuration(int32(plan.SponsorLinkValidityDuration.ValueInt64()))
	data.SetSponsorNotifyAll(plan.SponsorNotifyAll.ValueBool())
	data.SetSponsorStatusNotify(plan.SponsorStatusNotify.ValueBool())
	data.SetSponsors(sponsors)
	data.SetSsoDefaultRole(plan.SsoDefaultRole.ValueString())
	data.SetSsoForcedRole(plan.SsoForcedRole.ValueString())
	data.SetSsoIdpCert(plan.SsoIdpCert.ValueString())
	data.SetSsoIdpSignAlgo(plan.SsoIdpSignAlgo.ValueString())
	data.SetSsoIdpSsoUrl(plan.SsoIdpSsoUrl.ValueString())
	data.SetSsoIssuer(plan.SsoIssuer.ValueString())
	data.SetSsoNameidFormat(mistsdkgo.WlanPortalSsoNameidFormat(plan.SsoNameidFormat.ValueString()))
	data.SetTelstraClientId(plan.TelstraClientId.ValueString())
	data.SetTelstraClientSecret(plan.TelstraClientSecret.ValueString())
	data.SetThumbnail(plan.Thumbnail.ValueString())
	data.SetTwilioAuthToken(plan.TwilioAuthToken.ValueString())
	data.SetTwilioPhoneNumber(plan.TwilioPhoneNumber.ValueString())
	data.SetTwilioSid(plan.TwilioSid.ValueString())

	return data
}
