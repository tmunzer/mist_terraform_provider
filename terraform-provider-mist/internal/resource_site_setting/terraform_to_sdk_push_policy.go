package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	hours "terraform-provider-mist/internal/commons/hours"

	"mistapi/models"
)

func pushPolicyPushWindowConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SiteSettingConfigPushPolicyPushWindow {
	tflog.Debug(ctx, "pushPolicyPushWindowConfigTerraformToSdk")
	data := models.NewSiteSettingConfigPushPolicyPushWindow()

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

func pushPolicyConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ConfigPushPolicyValue) *models.SiteSettingConfigPushPolicy {
	tflog.Debug(ctx, "pushPolicyConfigTerraformToSdk")
	data := models.NewSiteSettingConfigPushPolicy()

	data.SetNoPush(d.NoPush.ValueBool())

	push_window := pushPolicyPushWindowConfigTerraformToSdk(ctx, diags, d.PushWindow)
	data.SetPushWindow(*push_window)

	return data
}
