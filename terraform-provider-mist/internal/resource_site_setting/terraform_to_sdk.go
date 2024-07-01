package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"
	"terraform-provider-mist/internal/resource_networktemplate"
)

func TerraformToSdk(ctx context.Context, plan *SiteSettingModel) (mistsdkgo.SiteSetting, diag.Diagnostics) {
	data := *mistsdkgo.NewSiteSetting()
	var diags diag.Diagnostics

	data.SetSiteId(plan.SiteId.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())

	analytic := analyticTerraformToSdk(ctx, &diags, plan.Analytic)
	data.SetAnalytic(*analytic)
	data.SetApUpdownThreshold(int32(plan.ApUpdownThreshold.ValueInt64()))

	auto_upgrade := siteSettingAutoUpgradeTerraformToSdk(ctx, &diags, plan.AutoUpgrade)
	data.SetAutoUpgrade(*auto_upgrade)

	ble_config := siteSettingBleConfigTerraformToSdk(ctx, &diags, plan.BleConfig)
	data.SetBleConfig(*ble_config)

	data.SetConfigAutoRevert(plan.ConfigAutoRevert.ValueBool())

	push_policy := pushPolicyConfigTerraformToSdk(ctx, &diags, plan.ConfigPushPolicy)
	data.SetConfigPushPolicy(*push_policy)

	critical_url_monitoring := criticalUrlMonitoringTerraformToSdk(ctx, &diags, plan.CriticalUrlMonitoring)
	data.SetCriticalUrlMonitoring(critical_url_monitoring)

	data.SetDeviceUpdownThreshold(int32(plan.DeviceUpdownThreshold.ValueInt64()))
	data.SetDisabledSystemDefinedPortUsages(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DisabledSystemDefinedPortUsages))
	data.SetDnsServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers))
	data.SetDnsSuffix(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix))

	engagement := engagementTerraformToSdk(ctx, &diags, plan.Engagement)
	data.SetEngagement(*engagement)

	data.SetGatewayUpdownThreshold(int32(plan.GatewayUpdownThreshold.ValueInt64()))

	led := ledTerraformToSdk(ctx, &diags, plan.Led)
	data.SetLed(led)

	network := resource_networktemplate.NetworksTerraformToSdk(ctx, &diags, plan.Networks)
	data.SetNetworks(network)

	data.SetNtpServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers))

	occupancy := occupancyTerraformToSdk(ctx, &diags, plan.Occupancy)
	data.SetOccupancy(occupancy)

	data.SetPersistConfigOnDevice(plan.PersistConfigOnDevice.ValueBool())

	proxy := proxyTerraformToSdk(ctx, &diags, plan.Proxy)
	data.SetProxy(proxy)

	data.SetReportGatt(plan.ReportGatt.ValueBool())

	rogue := rogueTerraformToSdk(ctx, &diags, plan.Rogue)
	data.SetRogue(rogue)

	simple_alert := simpleAlertTerraformToSdk(ctx, &diags, plan.SimpleAlert)
	data.SetSimpleAlert(simple_alert)

	sky_atp := skyAtpTerraformToSdk(ctx, &diags, plan.Skyatp)
	data.SetSkyatp(sky_atp)

	srx_app := srxAppTerraformToSdk(ctx, &diags, plan.SrxApp)
	data.SetSrxApp(srx_app)

	data.SetSshKeys(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SshKeys))

	ssr := ssrTerraformToSdk(ctx, &diags, plan.Ssr)
	data.SetSsr(ssr)

	data.SetSwitchUpdownThreshold(int32(plan.SwitchUpdownThreshold.ValueInt64()))

	synthetic_test := syntheticTestTerraformToSdk(ctx, &diags, plan.SyntheticTest)
	data.SetSyntheticTest(synthetic_test)

	data.SetTrackAnonymousDevices(plan.TrackAnonymousDevices.ValueBool())

	vars := varsTerraformToSdk(ctx, &diags, plan.Vars)
	data.SetVars(vars)

	vna := vnaTerraformToSdk(ctx, &diags, plan.Vna)
	data.SetVna(vna)

	wan_vna := wanVnaTerraformToSdk(ctx, &diags, plan.WanVna)
	data.SetWanVna(wan_vna)

	wids := widsTerraformToSdk(ctx, &diags, plan.Wids)
	data.SetWids(wids)

	wifi := wifiTerraformToSdk(ctx, &diags, plan.Wifi)
	data.SetWifi(wifi)

	wired_van := wiredVnaTerraformToSdk(ctx, &diags, plan.WiredVna)
	data.SetWiredVna(wired_van)

	zone_occupancy_alert := zoneOccupancyTerraformToSdk(ctx, &diags, plan.ZoneOccupancyAlert)
	data.SetZoneOccupancyAlert(zone_occupancy_alert)

	return data, diags
}
