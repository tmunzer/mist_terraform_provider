package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	hours "terraform-provider-mist/internal/commons/hours"

	"mistapi/models"
)

func pushPolicyPushWindowConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.WlanSchedule {
	tflog.Debug(ctx, "pushPolicyPushWindowConfigTerraformToSdk")
	data := models.WlanSchedule{}

	if !d.IsNull() && !d.IsUnknown() {
		vd, e := NewPushWindowValue(PushWindowValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		data.Enabled = vd.Enabled.ValueBoolPointer()

		hours := hours.HoursTerraformToSdk(ctx, diags, vd.Hours)
		data.Hours = hours

	}
	return &data
}

func pushPolicyConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ConfigPushPolicyValue) *models.SiteSettingConfigPushPolicy {
	tflog.Debug(ctx, "pushPolicyConfigTerraformToSdk")
	data := models.SiteSettingConfigPushPolicy{}

	if !d.IsNull() && !d.IsUnknown() {
		data.NoPush = d.NoPush.ValueBoolPointer()
	}

	push_window := pushPolicyPushWindowConfigTerraformToSdk(ctx, diags, d.PushWindow)
	data.PushWindow = push_window

	return &data
}