package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	hours "terraform-provider-mist/internal/commons/hours"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func pushPolicyPushWindowConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *mistapigo.SiteSettingConfigPushPolicyPushWindow {
	tflog.Debug(ctx, "pushPolicyPushWindowConfigTerraformToSdk")
	data := mistapigo.NewSiteSettingConfigPushPolicyPushWindow()

	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		vd, e := NewPushWindowValue(PushWindowValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		data.SetEnabled(vd.Enabled.ValueBool())

		hours := hours.HoursTerraformToSdk(ctx, diags, vd.Hours)
		data.SetHours(hours)

		return data
	}
}

func pushPolicyConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ConfigPushPolicyValue) *mistapigo.SiteSettingConfigPushPolicy {
	tflog.Debug(ctx, "pushPolicyConfigTerraformToSdk")
	data := mistapigo.NewSiteSettingConfigPushPolicy()

	data.SetNoPush(d.NoPush.ValueBool())

	push_window := pushPolicyPushWindowConfigTerraformToSdk(ctx, diags, d.PushWindow)
	data.SetPushWindow(*push_window)

	return data
}
