package resource_site_setting

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func criticalUrlMonitoringMonitorSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.SiteSettingCriticalUrlMonitoringMonitor) basetypes.ListValue {
	tflog.Debug(ctx, "criticalUrlMonitoringMonitorSdkToTerraform")
	var data_list = []MonitorsValue{}
	for _, v := range d {
		data_map_attr_type := MonitorsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"url":     types.StringValue(v.GetUrl()),
			"vlan_id": types.Int64Value(int64(v.GetVlanId())),
		}
		data, e := NewMonitorsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := MonitorsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func criticalUrlMonitoringSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteSettingCriticalUrlMonitoring) CriticalUrlMonitoringValue {
	tflog.Debug(ctx, "criticalUrlMonitoringSdkToTerraform")

	r_attr_type := CriticalUrlMonitoringValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled":  types.BoolValue(d.GetEnabled()),
		"monitors": criticalUrlMonitoringMonitorSdkToTerraform(ctx, diags, d.GetMonitors()),
	}
	r, e := NewCriticalUrlMonitoringValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
