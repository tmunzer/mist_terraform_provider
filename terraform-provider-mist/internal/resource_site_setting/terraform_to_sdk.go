package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func TerraformToSdk(ctx context.Context, plan *SiteSettingModel) (mistsdkgo.SiteSetting, diag.Diagnostics) {
	data := *mistsdkgo.NewSiteSetting()
	var diags diag.Diagnostics

	data.SetSiteId(plan.SiteId.ValueString())

	analytic := analyticTerraformToSdk(ctx, &diags, plan.Analytic)
	data.SetAnalytic(*analytic)
	data.SetApUpdownThreshold(int32(plan.ApUpdownThreshold.ValueInt64()))

	if !plan.AutoUpgrade.IsNull() && !plan.AutoUpgrade.IsUnknown() {
		auto_upgrade := siteSettingAutoUpgradeTerraformToSdk(ctx, &diags, plan.AutoUpgrade)
		data.SetAutoUpgrade(*auto_upgrade)
	} else {
		data.SetAutoUpgrade(*mistsdkgo.NewSiteSettingAutoUpgradeWithDefaults())
	}

	ble_config := siteSettingBleConfigTerraformToSdk(ctx, &diags, plan.BleConfig)
	data.SetBleConfig(*ble_config)

	data.SetConfigAutoRevert(plan.ConfigAutoRevert.ValueBool())

	if !plan.ConfigPushPolicy.IsNull() && !plan.ConfigPushPolicy.IsUnknown() {
		push_policy := pushPolicyConfigTerraformToSdk(ctx, &diags, plan.ConfigPushPolicy)
		data.SetConfigPushPolicy(*push_policy)
	}

	if !plan.CriticalUrlMonitoring.IsNull() && !plan.CriticalUrlMonitoring.IsUnknown() {
		critical_url_monitoring := criticalUrlMonitoringTerraformToSdk(ctx, &diags, plan.CriticalUrlMonitoring)
		data.SetCriticalUrlMonitoring(critical_url_monitoring)
	}

	data.SetDeviceUpdownThreshold(int32(plan.DeviceUpdownThreshold.ValueInt64()))
	data.SetDisabledSystemDefinedPortUsages(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DisabledSystemDefinedPortUsages))

	if !plan.Engagement.IsNull() && !plan.Engagement.IsUnknown() {
		engagement := engagementTerraformToSdk(ctx, &diags, plan.Engagement)
		data.SetEngagement(*engagement)

	}

	if !plan.GatewayUpdownThreshold.IsNull() && !plan.GatewayUpdownThreshold.IsUnknown() {
		data.SetGatewayUpdownThreshold(int32(plan.GatewayUpdownThreshold.ValueInt64()))

	}

	if !plan.Led.IsNull() && !plan.Led.IsUnknown() {
		led := ledTerraformToSdk(ctx, &diags, plan.Led)
		data.SetLed(led)
	} else {
		data.SetLed(*mistsdkgo.NewApLedWithDefaults())
	}

	if !plan.Occupancy.IsNull() && !plan.Occupancy.IsUnknown() {
		occupancy := occupancyTerraformToSdk(ctx, &diags, plan.Occupancy)
		data.SetOccupancy(occupancy)
	}

	data.SetPersistConfigOnDevice(plan.PersistConfigOnDevice.ValueBool())

	if !plan.Proxy.IsNull() && !plan.Proxy.IsUnknown() {
		proxy := proxyTerraformToSdk(ctx, &diags, plan.Proxy)
		data.SetProxy(proxy)

	}

	data.SetReportGatt(plan.ReportGatt.ValueBool())

	if !plan.Rogue.IsNull() && !plan.Rogue.IsUnknown() {
		rogue := rogueTerraformToSdk(ctx, &diags, plan.Rogue)
		data.SetRogue(rogue)

	}

	if !plan.SimpleAlert.IsNull() && !plan.SimpleAlert.IsUnknown() {
		simple_alert := simpleAlertTerraformToSdk(ctx, &diags, plan.SimpleAlert)
		data.SetSimpleAlert(simple_alert)

	}

	if !plan.Skyatp.IsNull() && !plan.Skyatp.IsUnknown() {
		sky_atp := skyAtpTerraformToSdk(ctx, &diags, plan.Skyatp)
		data.SetSkyatp(sky_atp)

	}

	if !plan.SrxApp.IsNull() && !plan.SrxApp.IsUnknown() {
		srx_app := srxAppTerraformToSdk(ctx, &diags, plan.SrxApp)
		data.SetSrxApp(srx_app)

	}

	data.SetSshKeys(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SshKeys))

	if !plan.Ssr.IsNull() && !plan.Ssr.IsUnknown() {
		ssr := ssrTerraformToSdk(ctx, &diags, plan.Ssr)
		data.SetSsr(ssr)

	}

	data.SetSwitchUpdownThreshold(int32(plan.SwitchUpdownThreshold.ValueInt64()))

	if !plan.SyntheticTest.IsNull() && !plan.SyntheticTest.IsUnknown() {
		synthetic_test := syntheticTestTerraformToSdk(ctx, &diags, plan.SyntheticTest)
		data.SetSyntheticTest(synthetic_test)

	}

	data.SetTrackAnonymousDevices(plan.TrackAnonymousDevices.ValueBool())

	if !plan.Vars.IsNull() && !plan.Vars.IsUnknown() {
		vars := varsTerraformToSdk(ctx, &diags, plan.Vars)
		data.SetVars(vars)

	}

	if !plan.Vna.IsNull() && !plan.Vna.IsUnknown() {
		vna := vnaTerraformToSdk(ctx, &diags, plan.Vna)
		data.SetVna(vna)

	}

	if !plan.WanVna.IsNull() && !plan.WanVna.IsUnknown() {
		wan_vna := wanVnaTerraformToSdk(ctx, &diags, plan.WanVna)
		data.SetWanVna(wan_vna)
	}

	if !plan.Wids.IsNull() && !plan.Wids.IsUnknown() {
		wids := widsTerraformToSdk(ctx, &diags, plan.Wids)
		data.SetWids(wids)
	}

	if !plan.Wifi.IsNull() && !plan.Wifi.IsUnknown() {
		wifi := wifiTerraformToSdk(ctx, &diags, plan.Wifi)
		data.SetWifi(wifi)

	}

	if !plan.WiredVna.IsNull() && !plan.WiredVna.IsUnknown() {
		wired_van := wiredVnaTerraformToSdk(ctx, &diags, plan.WiredVna)
		data.SetWiredVna(wired_van)

	}

	if !plan.ZoneOccupancyAlert.IsNull() && !plan.ZoneOccupancyAlert.IsUnknown() {
		zone_occupancy_alert := zoneOccupancyTerraformToSdk(ctx, &diags, plan.ZoneOccupancyAlert)
		data.SetZoneOccupancyAlert(zone_occupancy_alert)
	}
	return data, diags
}
