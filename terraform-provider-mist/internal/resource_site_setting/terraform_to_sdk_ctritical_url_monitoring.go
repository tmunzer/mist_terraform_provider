package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func criticalUrlMonitoringMonitorsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.SiteSettingCriticalUrlMonitoringMonitor {
	tflog.Debug(ctx, "criticalUrlMonitoringMonitorsTerraformToSdk")
	var data_list []mistsdkgo.SiteSettingCriticalUrlMonitoringMonitor
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(MonitorsValue)
		data := mistsdkgo.NewSiteSettingCriticalUrlMonitoringMonitor()
		data.SetUrl(plan.Url.ValueString())
		data.SetVlanId(int32(plan.VlanId.ValueInt64()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func criticalUrlMonitoringTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d CriticalUrlMonitoringValue) mistsdkgo.SiteSettingCriticalUrlMonitoring {
	tflog.Debug(ctx, "criticalUrlMonitoringTerraformToSdk")
	data := mistsdkgo.NewSiteSettingCriticalUrlMonitoring()

	data.SetEnabled(d.Enabled.ValueBool())

	monitors := criticalUrlMonitoringMonitorsTerraformToSdk(ctx, diags, d.Monitors)
	data.SetMonitors(monitors)

	return *data
}
