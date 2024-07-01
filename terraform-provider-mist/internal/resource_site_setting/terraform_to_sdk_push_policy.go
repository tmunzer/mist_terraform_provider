package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func pushPolicyPushWindowConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *mistsdkgo.SiteSettingConfigPushPolicyPushWindow {
	data := mistsdkgo.NewSiteSettingConfigPushPolicyPushWindow()

	v := NewPushWindowValueMust(d.AttributeTypes(ctx), d.Attributes())
	data.SetEnabled(v.Enabled.ValueBool())

	hours := HoursConfigTerraformToSdk(ctx, diags, v.Hours)
	data.SetHours(hours)

	return data
}

func pushPolicyConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ConfigPushPolicyValue) *mistsdkgo.SiteSettingConfigPushPolicy {
	data := mistsdkgo.NewSiteSettingConfigPushPolicy()

	data.SetNoPush(d.NoPush.ValueBool())

	push_window := pushPolicyPushWindowConfigTerraformToSdk(ctx, diags, d.PushWindow)
	data.SetPushWindow(*push_window)

	return data
}
