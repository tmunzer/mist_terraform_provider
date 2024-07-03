package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func autoUpgradeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SiteSettingAutoUpgrade) AutoUpgradeValue {
	tflog.Debug(ctx, "autoUpgradeSdkToTerraform")

	custom_versions_map_value := make(map[string]string)
	for k, v := range d.GetCustomVersions() {
		custom_versions_map_value[k] = v
	}
	custom_versions, e := types.MapValueFrom(ctx, types.StringType, custom_versions_map_value)
	diags.Append(e...)

	r_attr_type := AutoUpgradeValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"custom_versions": custom_versions,
		"day_of_week":     types.StringValue(string(d.GetDayOfWeek())),
		"enabled":         types.BoolValue(d.GetEnabled()),
		"time_of_day":     types.StringValue(d.GetTimeOfDay()),
		"version":         types.StringValue(string(d.GetVersion())),
	}
	r, e := NewAutoUpgradeValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
