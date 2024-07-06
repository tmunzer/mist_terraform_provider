package resource_site_setting

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func occupancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteOccupancyAnalytics) OccupancyValue {
	tflog.Debug(ctx, "occupancySdkToTerraform")

	r_attr_type := OccupancyValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"assets_enabled":              types.BoolValue(d.GetAssetsEnabled()),
		"clients_enabled":             types.BoolValue(d.GetClientsEnabled()),
		"min_duration":                types.Int64Value(int64(d.GetMinDuration())),
		"sdkclients_enabled":          types.BoolValue(d.GetSdkclientsEnabled()),
		"unconnected_clients_enabled": types.BoolValue(d.GetUnconnectedClientsEnabled()),
	}
	r, e := NewOccupancyValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
