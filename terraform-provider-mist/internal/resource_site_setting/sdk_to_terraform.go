package resource_site_setting

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_list "terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.SiteSetting) (SiteSettingModel, diag.Diagnostics) {
	var state SiteSettingModel
	var diags diag.Diagnostics

	state.SiteId = types.StringValue(data.SiteId.String())
	state.OrgId = types.StringValue(data.OrgId.String())

	// state.Analytic = analyticSdkToTerraform(ctx, &diags, data.Analytic)

	state.ApUpdownThreshold = types.Int64Value(int64(*data.ApUpdownThreshold.Value()))

	state.AutoUpgrade = autoUpgradeSdkToTerraform(ctx, &diags, *data.AutoUpgrade)

	state.BleConfig = bleConfigsSdkToTerraform(ctx, &diags, data.BleConfig)

	state.BlacklistUrl = types.StringValue(*data.BlacklistUrl)

	state.ConfigAutoRevert = types.BoolValue(*data.ConfigAutoRevert)

	state.ConfigPushPolicy = configPushPolicySdkToTerraform(ctx, &diags, data.ConfigPushPolicy)

	state.CriticalUrlMonitoring = criticalUrlMonitoringSdkToTerraform(ctx, &diags, data.CriticalUrlMonitoring)

	state.DeviceUpdownThreshold = types.Int64Value(int64(*data.DeviceUpdownThreshold))

	state.DisabledSystemDefinedPortUsages = mist_list.ListOfStringSdkToTerraform(ctx, data.DisabledSystemDefinedPortUsages)

	state.Engagement = engagementSdkToTerraform(ctx, &diags, data.Engagement)

	state.GatewayUpdownThreshold = types.Int64Value(int64(*data.GatewayUpdownThreshold.Value()))

	state.Led = ledSdkToTerraform(ctx, &diags, data.Led)

	state.Occupancy = occupancySdkToTerraform(ctx, &diags, data.Occupancy)

	state.PersistConfigOnDevice = types.BoolValue(*data.PersistConfigOnDevice)

	state.Proxy = proxySdkToTerraform(ctx, &diags, data.Proxy)

	state.ReportGatt = types.BoolValue(*data.ReportGatt)

	state.Rogue = rogueSdkToTerraform(ctx, &diags, data.Rogue)

	state.SimpleAlert = simpleAlertSdkToTerraform(ctx, &diags, data.SimpleAlert)

	state.Skyatp = skyAtpSdkToTerraform(ctx, &diags, data.Skyatp)

	// state.SrxApp = srxAppSdkToTerraform(ctx, &diags, data.SrxApp)

	state.SshKeys = mist_list.ListOfStringSdkToTerraform(ctx, data.SshKeys)

	state.Ssr = ssrSdkToTerraform(ctx, &diags, data.Ssr)

	state.SwitchUpdownThreshold = types.Int64Value(int64(*data.SwitchUpdownThreshold.Value()))

	state.SyntheticTest = synthteticTestSdkToTerraform(ctx, &diags, data.Synthetictest)

	state.TrackAnonymousDevices = types.BoolValue(*data.TrackAnonymousDevices)

	state.Vars = varsSdkToTerraform(ctx, &diags, data.Vars)

	// state.Vna = vnaSdkToTerraform(ctx, &diags, data.Vna)

	// state.WanVna = wanVnaSdkToTerraform(ctx, &diags, data.WanVna)

	state.WatchedStationUrl = types.StringValue(*data.WatchedStationUrl)

	state.WhitelistUrl = types.StringValue(*data.WhitelistUrl)

	state.Wids = widsSdkToTerraform(ctx, &diags, data.Wids)

	state.Wifi = wifiSdkToTerraform(ctx, &diags, data.Wifi)

	// state.WiredVna = wiredVnaSdkToTerraform(ctx, &diags, data.WiredVna)

	state.ZoneOccupancyAlert = zoneOccupancySdkToTerraform(ctx, &diags, *data.ZoneOccupancyAlert)

	return state, diags
}
