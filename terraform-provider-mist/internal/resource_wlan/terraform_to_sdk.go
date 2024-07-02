package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/commons/radius_acct"
	"terraform-provider-mist/internal/commons/radius_auth"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *WlanModel) (mistsdkgo.Wlan, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewWlan(plan.Ssid.ValueString())
	data.SetId(plan.Id.ValueString())
	data.SetSiteId(plan.SiteId.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())
	data.SetTemplateId(plan.TemplateId.ValueString())

	data.SetAcctImmediateUpdate(plan.AcctImmediateUpdate.ValueBool())
	data.SetAcctInterimInterval(int32(plan.AcctInterimInterval.ValueInt64()))

	acct_servers := radius_acct.RadiusAcctServersTerraformToSdk(ctx, &diags, plan.AcctServers)
	data.SetAcctServers(acct_servers)

	airwatch := airwatchTerraformToSdk(ctx, &diags, plan.Airwatch)
	data.SetAirwatch(airwatch)

	data.SetAllowIpv6Ndp(plan.AllowIpv6Ndp.ValueBool())
	data.SetAllowMdns(plan.AllowMdns.ValueBool())
	data.SetAllowMdns(plan.AllowMdns.ValueBool())
	data.SetApIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.ApIds))

	app_limit := appLimitTerraformToSdk(ctx, &diags, plan.AppLimit)
	data.SetAppLimit(app_limit)

	app_qos := appQosTerraformToSdk(ctx, &diags, plan.AppQos)
	data.SetAppQos(app_qos)

	data.SetApplyTo(mistsdkgo.WlanApplyTo(plan.ApplyTo.ValueString()))
	data.SetArpFilter(plan.ArpFilter.ValueBool())

	auth := authTerraformToSdk(ctx, &diags, plan.Auth)
	data.SetAuth(auth)

	data.SetAuthServerSelection(mistsdkgo.WlanAuthServerSelection(plan.AuthServerSelection.ValueString()))

	auth_servers := radius_auth.RadiusAuthServersTerraformToSdk(ctx, &diags, plan.AuthServers)
	data.SetAuthServers(auth_servers)

	data.SetAuthServersNasId(plan.AuthServersNasId.ValueString())
	data.SetAuthServersNasIp(plan.AuthServersNasIp.ValueString())
	data.SetAuthServersRetries(int32(plan.AuthServersRetries.ValueInt64()))
	data.SetAuthServersTimeout(int32(plan.AuthServersTimeout.ValueInt64()))
	data.SetBandSteer(plan.BandSteer.ValueBool())
	data.SetBandSteerForceBand5(plan.BandSteerForceBand5.ValueBool())

	bands := bandsTerraformToSdk(ctx, &diags, plan.Bands)
	data.SetBands(bands)

	data.SetBlockBlacklistClients(plan.BlockBlacklistClients.ValueBool())

	bonjour := bonjourTerraformToSdk(ctx, &diags, plan.Bonjour)
	data.SetBonjour(bonjour)

	cisco_cwa := ciscoCwaTerraformToSdk(ctx, &diags, plan.CiscoCwa)
	data.SetCiscoCwa(cisco_cwa)

	data.SetClientLimitDown(int32(plan.ClientLimitDown.ValueInt64()))
	data.SetClientLimitDownEnabled(plan.ClientLimitDownEnabled.ValueBool())
	data.SetClientLimitUp(int32(plan.ClientLimitUp.ValueInt64()))
	data.SetClientLimitUpEnabled(plan.ClientLimitUpEnabled.ValueBool())

	coa_servers := coaServerTerraformToSdk(ctx, &diags, plan.CoaServers)
	data.SetCoaServers(coa_servers)

	data.SetDisable11ax(plan.Disable11ax.ValueBool())
	data.SetDisableHtVhtRates(plan.DisableHtVhtRates.ValueBool())
	data.SetDisableUapsd(plan.DisableUapsd.ValueBool())
	data.SetDisableV1RoamNotify(plan.DisableV1RoamNotify.ValueBool())
	data.SetDisableV2RoamNotify(plan.DisableV2RoamNotify.ValueBool())
	data.SetDisableWmm(plan.DisableWmm.ValueBool())

	dns_server_rewrite := dnsServerRewriteTerraformToSdk(ctx, &diags, plan.DnsServerRewrite)
	data.SetDnsServerRewrite(dns_server_rewrite)

	data.SetDtim(int32(plan.Dtim.ValueInt64()))

	dynamic_psk := dynamicPskTerraformToSdk(ctx, &diags, plan.DynamicPsk)
	data.SetDynamicPsk(dynamic_psk)

	dynamic_vlan := dynamicVlanTerraformToSdk(ctx, &diags, plan.DynamicVlan)
	data.SetDynamicVlan(dynamic_vlan)

	data.SetEnableLocalKeycaching(plan.EnableLocalKeycaching.ValueBool())
	data.SetEnableWirelessBridging(plan.EnableWirelessBridging.ValueBool())
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetFastDot1xTimers(plan.FastDot1xTimers.ValueBool())
	data.SetHideSsid(plan.HideSsid.ValueBool())
	data.SetHostnameIe(plan.HostnameIe.ValueBool())

	hotspot20 := hotspot20TerraformToSdk(ctx, &diags, plan.Hotspot20)
	data.SetHotspot20(hotspot20)

	inject_dhcp_option_82 := injectDhcpOption82TerraformToSdk(ctx, &diags, plan.InjectDhcpOption82)
	data.SetInjectDhcpOption82(inject_dhcp_option_82)

	data.SetInterface(mistsdkgo.WlanInterface(plan.Interface.ValueString()))
	data.SetIsolation(plan.Isolation.ValueBool())
	data.SetL2Isolation(plan.L2Isolation.ValueBool())
	data.SetLegacyOverds(plan.LegacyOverds.ValueBool())
	data.SetLimitBcast(plan.LimitBcast.ValueBool())
	data.SetLimitProbeResponse(plan.LimitProbeResponse.ValueBool())
	data.SetMaxIdletime(int32(plan.MaxIdletime.ValueInt64()))

	mist_nac := mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
	data.SetMistNac(mist_nac)

	data.SetMxtunnelIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.MxtunnelIds))
	data.SetMxtunnelName(mist_transform.ListOfStringTerraformToSdk(ctx, plan.MxtunnelName))
	data.SetNoStaticDns(plan.NoStaticDns.ValueBool())
	data.SetNoStaticIp(plan.NoStaticIp.ValueBool())

	portal := portalTerraformToSdk(ctx, &diags, plan.Portal)
	data.SetPortal(portal)

	data.SetPortalAllowedHostnames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.PortalAllowedHostnames))
	data.SetPortalAllowedSubnets(mist_transform.ListOfStringTerraformToSdk(ctx, plan.PortalAllowedSubnets))
	data.SetPortalDeniedHostnames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.PortalDeniedHostnames))
	data.SetPortalSsoUrl(plan.PortalSsoUrl.ValueString())

	qos := qosTerraformToSdk(ctx, &diags, plan.Qos)
	data.SetQos(qos)

	radesc := radsecTerraformToSdk(ctx, &diags, plan.Radsec)
	data.SetRadsec(radesc)

	data.SetRoamMode(mistsdkgo.WlanRoamMode(plan.RoamMode.ValueString()))

	schedule := scheduleTerraformToSdk(ctx, &diags, plan.Schedule)
	data.SetSchedule(schedule)

	data.SetSleExcluded(plan.SleExcluded.ValueBool())
	data.SetUseEapolV1(plan.UseEapolV1.ValueBool())
	data.SetVlanEnabled(plan.VlanEnabled.ValueBool())
	data.SetVlanId(int32(plan.VlanId.ValueInt64()))

	vlan_ids := vlanIdsTerraformToSdk(ctx, &diags, plan.VlanIds)
	data.SetVlanIds(vlan_ids)
	data.SetVlanPooling(plan.VlanPooling.ValueBool())
	data.SetWlanLimitDown(int32(plan.WlanLimitDown.ValueInt64()))
	data.SetWlanLimitDownEnabled(plan.WlanLimitDownEnabled.ValueBool())
	data.SetWlanLimitUp(int32(plan.WlanLimitUp.ValueInt64()))
	data.SetWlanLimitUpEnabled(plan.WlanLimitUpEnabled.ValueBool())
	data.SetWxtagIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.WxtagIds))
	data.SetWxtunnelId(plan.WxtunnelId.ValueString())
	data.SetWxtunnelRemoteId(plan.WxtunnelRemoteId.ValueString())

	return data, diags
}
