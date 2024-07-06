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

	state.SiteId = types.StringValue(data.GetSiteId())
	state.OrgId = types.StringValue(data.GetOrgId())

	state.Analytic = analyticSdkToTerraform(ctx, &diags, data.GetAnalytic())

	state.ApUpdownThreshold = types.Int64Value(int64(data.GetApUpdownThreshold()))

	state.AutoUpgrade = autoUpgradeSdkToTerraform(ctx, &diags, data.GetAutoUpgrade())

	state.BleConfig = bleConfigsSdkToTerraform(ctx, &diags, data.GetBleConfig())

	state.BlacklistUrl = types.StringValue(data.GetBlacklistUrl())

	state.ConfigAutoRevert = types.BoolValue(data.GetConfigAutoRevert())

	state.ConfigPushPolicy = configPushPolicySdkToTerraform(ctx, &diags, data.GetConfigPushPolicy())

	state.CriticalUrlMonitoring = criticalUrlMonitoringSdkToTerraform(ctx, &diags, data.GetCriticalUrlMonitoring())

	state.DeviceUpdownThreshold = types.Int64Value(int64(data.GetDeviceUpdownThreshold()))

	state.DisabledSystemDefinedPortUsages = mist_list.ListOfStringSdkToTerraform(ctx, data.GetDisabledSystemDefinedPortUsages())

	state.Engagement = engagementSdkToTerraform(ctx, &diags, data.GetEngagement())

	state.GatewayUpdownThreshold = types.Int64Value(int64(data.GetGatewayUpdownThreshold()))

	state.Led = ledSdkToTerraform(ctx, &diags, data.GetLed())

	state.Occupancy = occupancySdkToTerraform(ctx, &diags, data.GetOccupancy())

	state.PersistConfigOnDevice = types.BoolValue(data.GetPersistConfigOnDevice())

	state.Proxy = proxySdkToTerraform(ctx, &diags, data.GetProxy())

	state.ReportGatt = types.BoolValue(data.GetReportGatt())

	state.Rogue = rogueSdkToTerraform(ctx, &diags, data.GetRogue())

	state.SimpleAlert = simpleAlertSdkToTerraform(ctx, &diags, data.GetSimpleAlert())

	state.Skyatp = skyAtpSdkToTerraform(ctx, &diags, data.GetSkyatp())

	state.SrxApp = srxAppSdkToTerraform(ctx, &diags, data.GetSrxApp())

	state.SshKeys = mist_list.ListOfStringSdkToTerraform(ctx, data.GetSshKeys())

	state.Ssr = ssrSdkToTerraform(ctx, &diags, data.GetSsr())

	state.SwitchUpdownThreshold = types.Int64Value(int64(data.GetSwitchUpdownThreshold()))

	state.SyntheticTest = synthteticTestSdkToTerraform(ctx, &diags, data.GetSyntheticTest())

	state.TrackAnonymousDevices = types.BoolValue(data.GetTrackAnonymousDevices())

	state.Vars = varsSdkToTerraform(ctx, &diags, data.GetVars())

	state.Vna = vnaSdkToTerraform(ctx, &diags, data.GetVna())

	state.WanVna = wanVnaSdkToTerraform(ctx, &diags, data.GetWanVna())

	state.WatchedStationUrl = types.StringValue(data.GetWatchedStationUrl())

	state.WhitelistUrl = types.StringValue(data.GetWhitelistUrl())

	state.Wids = widsSdkToTerraform(ctx, &diags, data.GetWids())

	state.Wifi = wifiSdkToTerraform(ctx, &diags, data.GetWifi())

	state.WiredVna = wiredVnaSdkToTerraform(ctx, &diags, data.GetWiredVna())

	state.ZoneOccupancyAlert = zoneOccupancySdkToTerraform(ctx, &diags, data.GetZoneOccupancyAlert())

	return state, diags
}
