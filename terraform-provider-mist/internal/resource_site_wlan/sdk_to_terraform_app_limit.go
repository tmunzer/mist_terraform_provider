package resource_site_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func appLimitSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data models.WlanAppLimit) AppLimitValue {

	app_limit_attr := make(map[string]attr.Value)
	for k, v := range data.GetApps() {
		app_limit_attr[k] = types.Int64Value(int64(v))
	}
	app_limit, e := types.MapValueFrom(ctx, types.Int64Type, app_limit_attr)
	diags.Append(e...)

	wxtag_limit_attr := make(map[string]attr.Value)
	for k, v := range data.GetWxtagIds() {
		app_limit_attr[k] = types.Int64Value(int64(v))
	}
	wxtag_limit, e := types.MapValueFrom(ctx, types.Int64Type, wxtag_limit_attr)
	diags.Append(e...)

	plan_attr := map[string]attr.Value{
		"apps":      app_limit,
		"enabled":   types.BoolValue(data.GetEnabled()),
		"wxtag_ids": wxtag_limit,
	}
	state, e := NewAppLimitValue(AppLimitValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return state

}
