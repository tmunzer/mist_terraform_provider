package resource_site_setting

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func zoneOccupancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SiteZoneOccupancyAlert) ZoneOccupancyAlertValue {
	tflog.Debug(ctx, "zoneOccupancySdkToTerraform")

	r_attr_type := ZoneOccupancyAlertValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"email_notifiers": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetEmailNotifiers()),
		"enabled":         types.BoolValue(d.GetEnabled()),
		"threshold":       types.Int64Value(int64(d.GetThreshold())),
	}
	r, e := NewZoneOccupancyAlertValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
