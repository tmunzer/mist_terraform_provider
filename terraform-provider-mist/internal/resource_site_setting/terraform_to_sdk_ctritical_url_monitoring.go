package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func criticalUrlMonitoringMonitorsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SiteSettingCriticalUrlMonitoringMonitor {
	tflog.Debug(ctx, "criticalUrlMonitoringMonitorsTerraformToSdk")
	var data_list []models.SiteSettingCriticalUrlMonitoringMonitor
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(MonitorsValue)
		data := models.NewSiteSettingCriticalUrlMonitoringMonitor()
		data.SetUrl(plan.Url.ValueString())
		data.SetVlanId(int32(plan.VlanId.ValueInt64()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func criticalUrlMonitoringTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d CriticalUrlMonitoringValue) models.SiteSettingCriticalUrlMonitoring {
	tflog.Debug(ctx, "criticalUrlMonitoringTerraformToSdk")
	data := models.NewSiteSettingCriticalUrlMonitoring()

	data.SetEnabled(d.Enabled.ValueBool())

	monitors := criticalUrlMonitoringMonitorsTerraformToSdk(ctx, diags, d.Monitors)
	data.SetMonitors(monitors)

	return *data
}
