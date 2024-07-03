package resource_site_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *SiteWlanModel) (mistsdkgo.Wlan, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewWlan(plan.Ssid.ValueString())
	data.SetId(plan.Id.ValueString())
	data.SetSiteId(plan.SiteId.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())

	unset := make(map[string]interface{})

	if plan.AcctImmediateUpdate.IsNull() || plan.AcctImmediateUpdate.IsUnknown() {
		unset["-acct_immediate_update"] = ""
	} else {
		data.SetAcctImmediateUpdate(plan.AcctImmediateUpdate.ValueBool())
	}

	if plan.AcctInterimInterval.IsNull() || plan.AcctInterimInterval.IsUnknown() {
		unset["-acct_interim_interval"] = ""
	} else {
		data.SetAcctInterimInterval(int32(plan.AcctInterimInterval.ValueInt64()))
	}

	if plan.AcctServers.IsNull() || plan.AcctServers.IsUnknown() {
		unset["-acct_servers"] = ""
	} else {
		acct_servers := radiusAcctServersTerraformToSdk(ctx, &diags, plan.AcctServers)
		data.SetAcctServers(acct_servers)
	}

	if plan.Airwatch.IsNull() || plan.Airwatch.IsUnknown() {
		unset["-airwatch"] = ""
	} else {
		airwatch := airwatchTerraformToSdk(ctx, &diags, plan.Airwatch)
		data.SetAirwatch(airwatch)
	}

	if plan.AllowIpv6Ndp.IsNull() || plan.AllowIpv6Ndp.IsUnknown() {
		unset["-allow_ipv6_ndp"] = ""
	} else {
		data.SetAllowIpv6Ndp(plan.AllowIpv6Ndp.ValueBool())
	}

	if plan.AllowMdns.IsNull() || plan.AllowMdns.IsUnknown() {
		unset["-allow_mdns"] = ""
	} else {
		data.SetAllowMdns(plan.AllowMdns.ValueBool())
	}

	if plan.AllowMdns.IsNull() || plan.AllowMdns.IsUnknown() {
		unset["-allow_ssdp"] = ""
	} else {
		data.SetAllowMdns(plan.AllowMdns.ValueBool())
	}

	if plan.ApIds.IsNull() || plan.ApIds.IsUnknown() {
		unset["-ap_ids"] = ""
	} else {
		data.SetApIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.ApIds))
	}

	if plan.AppLimit.IsNull() || plan.AppLimit.IsUnknown() {
		unset["-app_limit"] = ""
	} else {
		app_limit := appLimitTerraformToSdk(ctx, &diags, plan.AppLimit)
		data.SetAppLimit(app_limit)
	}

	if plan.AppQos.IsNull() || plan.AppQos.IsUnknown() {
		unset["-app_qos"] = ""
	} else {
		app_qos := appQosTerraformToSdk(ctx, &diags, plan.AppQos)
		data.SetAppQos(app_qos)
	}

	if plan.ApplyTo.IsNull() || plan.ApplyTo.IsUnknown() {
		unset["-apply_to"] = ""
	} else {
		data.SetApplyTo(mistsdkgo.WlanApplyTo(plan.ApplyTo.ValueString()))
	}

	if plan.ArpFilter.IsNull() || plan.ArpFilter.IsUnknown() {
		unset["-arp_filter"] = ""
	} else {
		data.SetArpFilter(plan.ArpFilter.ValueBool())
	}

	if plan.Auth.IsNull() || plan.Auth.IsUnknown() {
		unset["-auth"] = ""
	} else {
		auth := authTerraformToSdk(ctx, &diags, plan.Auth)
		data.SetAuth(auth)
	}

	if plan.AuthServerSelection.IsNull() || plan.AuthServerSelection.IsUnknown() {
		unset["-auth_server_selection"] = ""
	} else {
		data.SetAuthServerSelection(mistsdkgo.WlanAuthServerSelection(plan.AuthServerSelection.ValueString()))
	}

	if plan.AuthServers.IsNull() || plan.AuthServers.IsUnknown() {
		unset["-auth_servers"] = ""
	} else {
		auth_servers := radiusAuthServersTerraformToSdk(ctx, &diags, plan.AuthServers)
		data.SetAuthServers(auth_servers)
	}

	if plan.AuthServersNasId.IsNull() || plan.AuthServersNasId.IsUnknown() {
		unset["-auth_servers_nas_id"] = ""
	} else {
		data.SetAuthServersNasId(plan.AuthServersNasId.ValueString())
	}

	if plan.AuthServersNasIp.IsNull() || plan.AuthServersNasIp.IsUnknown() {
		unset["-auth_servers_nas_ip"] = ""
	} else {
		data.SetAuthServersNasIp(plan.AuthServersNasIp.ValueString())
	}

	if plan.AuthServersRetries.IsNull() || plan.AuthServersRetries.IsUnknown() {
		unset["-auth_servers_retries"] = ""
	} else {
		data.SetAuthServersRetries(int32(plan.AuthServersRetries.ValueInt64()))
	}

	if plan.AuthServersTimeout.IsNull() || plan.AuthServersTimeout.IsUnknown() {
		unset["-auth_servers_timeout"] = ""
	} else {
		data.SetAuthServersTimeout(int32(plan.AuthServersTimeout.ValueInt64()))
	}

	if plan.BandSteer.IsNull() || plan.BandSteer.IsUnknown() {
		unset["-band_steer"] = ""
	} else {
		data.SetBandSteer(plan.BandSteer.ValueBool())
	}

	if plan.BandSteerForceBand5.IsNull() || plan.BandSteerForceBand5.IsUnknown() {
		unset["-band_steer_force_band5"] = ""
	} else {
		data.SetBandSteerForceBand5(plan.BandSteerForceBand5.ValueBool())
	}

	if plan.Bands.IsNull() || plan.Bands.IsUnknown() {
		unset["-bands"] = ""
	} else {
		bands := bandsTerraformToSdk(ctx, &diags, plan.Bands)
		data.SetBands(bands)
	}

	if plan.BlockBlacklistClients.IsNull() || plan.BlockBlacklistClients.IsUnknown() {
		unset["-block_blacklist_clients"] = ""
	} else {
		data.SetBlockBlacklistClients(plan.BlockBlacklistClients.ValueBool())
	}

	if plan.Bonjour.IsNull() || plan.Bonjour.IsUnknown() {
		unset["-bonjour"] = ""
	} else {
		bonjour := bonjourTerraformToSdk(ctx, &diags, plan.Bonjour)
		data.SetBonjour(bonjour)
	}

	if plan.CiscoCwa.IsNull() || plan.CiscoCwa.IsUnknown() {
		unset["-cisco_cwa"] = ""
	} else {
		cisco_cwa := ciscoCwaTerraformToSdk(ctx, &diags, plan.CiscoCwa)
		data.SetCiscoCwa(cisco_cwa)
	}

	if plan.ClientLimitDown.IsNull() || plan.ClientLimitDown.IsUnknown() {
		unset["-client_limit_down"] = ""
	} else {
		data.SetClientLimitDown(int32(plan.ClientLimitDown.ValueInt64()))
	}

	if plan.ClientLimitDownEnabled.IsNull() || plan.ClientLimitDownEnabled.IsUnknown() {
		unset["-client_limit_down_enabled"] = ""
	} else {
		data.SetClientLimitDownEnabled(plan.ClientLimitDownEnabled.ValueBool())
	}

	if plan.ClientLimitUp.IsNull() || plan.ClientLimitUp.IsUnknown() {
		unset["-client_limit_up"] = ""
	} else {
		data.SetClientLimitUp(int32(plan.ClientLimitUp.ValueInt64()))
	}

	if plan.ClientLimitUpEnabled.IsNull() || plan.ClientLimitUpEnabled.IsUnknown() {
		unset["-client_limit_up_enabled"] = ""
	} else {
		data.SetClientLimitUpEnabled(plan.ClientLimitUpEnabled.ValueBool())
	}

	if plan.CoaServers.IsNull() || plan.CoaServers.IsUnknown() {
		unset["-coa_servers"] = ""
	} else {
		coa_servers := coaServerTerraformToSdk(ctx, &diags, plan.CoaServers)
		data.SetCoaServers(coa_servers)
	}

	if plan.Disable11ax.IsNull() || plan.Disable11ax.IsUnknown() {
		unset["-disable_11ax"] = ""
	} else {
		data.SetDisable11ax(plan.Disable11ax.ValueBool())
	}

	if plan.DisableHtVhtRates.IsNull() || plan.DisableHtVhtRates.IsUnknown() {
		unset["-disable_ht_vht_rates"] = ""
	} else {
		data.SetDisableHtVhtRates(plan.DisableHtVhtRates.ValueBool())
	}

	if plan.DisableUapsd.IsNull() || plan.DisableUapsd.IsUnknown() {
		unset["-disable_uapsd"] = ""
	} else {
		data.SetDisableUapsd(plan.DisableUapsd.ValueBool())
	}

	if plan.DisableV1RoamNotify.IsNull() || plan.DisableV1RoamNotify.IsUnknown() {
		unset["-disable_v1_roam_notify"] = ""
	} else {
		data.SetDisableV1RoamNotify(plan.DisableV1RoamNotify.ValueBool())
	}

	if plan.DisableV2RoamNotify.IsNull() || plan.DisableV2RoamNotify.IsUnknown() {
		unset["-disable_v2_roam_notify"] = ""
	} else {
		data.SetDisableV2RoamNotify(plan.DisableV2RoamNotify.ValueBool())
	}

	if plan.DisableWmm.IsNull() || plan.DisableWmm.IsUnknown() {
		unset["-disable_wmm"] = ""
	} else {
		data.SetDisableWmm(plan.DisableWmm.ValueBool())
	}

	if plan.DnsServerRewrite.IsNull() || plan.DnsServerRewrite.IsUnknown() {
		unset["-dns_server_rewrite"] = ""
	} else {
		dns_server_rewrite := dnsServerRewriteTerraformToSdk(ctx, &diags, plan.DnsServerRewrite)
		data.SetDnsServerRewrite(dns_server_rewrite)
	}

	if plan.Dtim.IsNull() || plan.Dtim.IsUnknown() {
		unset["-dtim"] = ""
	} else {
		data.SetDtim(int32(plan.Dtim.ValueInt64()))
	}

	if plan.DynamicPsk.IsNull() || plan.DynamicPsk.IsUnknown() {
		unset["-dynamic_psk"] = ""
	} else {
		dynamic_psk := dynamicPskTerraformToSdk(ctx, &diags, plan.DynamicPsk)
		data.SetDynamicPsk(dynamic_psk)
	}

	if plan.DynamicVlan.IsNull() || plan.DynamicVlan.IsUnknown() {
		unset["-dynamic_vlan"] = ""
	} else {
		dynamic_vlan := dynamicVlanTerraformToSdk(ctx, &diags, plan.DynamicVlan)
		data.SetDynamicVlan(dynamic_vlan)
	}

	if plan.EnableLocalKeycaching.IsNull() || plan.EnableLocalKeycaching.IsUnknown() {
		unset["-enable_local_keycaching"] = ""
	} else {
		data.SetEnableLocalKeycaching(plan.EnableLocalKeycaching.ValueBool())
	}

	if plan.EnableWirelessBridging.IsNull() || plan.EnableWirelessBridging.IsUnknown() {
		unset["-enable_wireless_bridging"] = ""
	} else {
		data.SetEnableWirelessBridging(plan.EnableWirelessBridging.ValueBool())
	}

	if plan.EnableWirelessBridgingDhcpTracking.IsNull() || plan.EnableWirelessBridgingDhcpTracking.IsUnknown() {
		unset["-enable_wireless_bridging_dhcp_tracking"] = ""
	} else {
		data.SetEnableWirelessBridgingDhcpTracking(plan.EnableWirelessBridgingDhcpTracking.ValueBool())
	}

	if plan.Enabled.IsNull() || plan.Enabled.IsUnknown() {
		unset["-enabled"] = ""
	} else {
		data.SetEnabled(plan.Enabled.ValueBool())
	}

	if plan.FastDot1xTimers.IsNull() || plan.FastDot1xTimers.IsUnknown() {
		unset["-fast_dot1x_timers"] = ""
	} else {
		data.SetFastDot1xTimers(plan.FastDot1xTimers.ValueBool())
	}

	if plan.HideSsid.IsNull() || plan.HideSsid.IsUnknown() {
		unset["-hide_ssid"] = ""
	} else {
		data.SetHideSsid(plan.HideSsid.ValueBool())
	}

	if plan.HostnameIe.IsNull() || plan.HostnameIe.IsUnknown() {
		unset["-hostname_ie"] = ""
	} else {
		data.SetHostnameIe(plan.HostnameIe.ValueBool())
	}

	if plan.Hotspot20.IsNull() || plan.Hotspot20.IsUnknown() {
		unset["-hotspot20"] = ""
	} else {
		hotspot20 := hotspot20TerraformToSdk(ctx, &diags, plan.Hotspot20)
		data.SetHotspot20(hotspot20)
	}

	if plan.InjectDhcpOption82.IsNull() || plan.InjectDhcpOption82.IsUnknown() {
		unset["-inject_dhcp_option_82"] = ""
	} else {
		inject_dhcp_option_82 := injectDhcpOption82TerraformToSdk(ctx, &diags, plan.InjectDhcpOption82)
		data.SetInjectDhcpOption82(inject_dhcp_option_82)
	}

	if plan.Interface.IsNull() || plan.Interface.IsUnknown() {
		unset["-interface"] = ""
	} else {
		data.SetInterface(mistsdkgo.WlanInterface(plan.Interface.ValueString()))
	}

	if plan.Isolation.IsNull() || plan.Isolation.IsUnknown() {
		unset["-isolation"] = ""
	} else {
		data.SetIsolation(plan.Isolation.ValueBool())
	}

	if plan.L2Isolation.IsNull() || plan.L2Isolation.IsUnknown() {
		unset["-l2_isolation"] = ""
	} else {
		data.SetL2Isolation(plan.L2Isolation.ValueBool())
	}

	if plan.LegacyOverds.IsNull() || plan.LegacyOverds.IsUnknown() {
		unset["-legacy_overds"] = ""
	} else {
		data.SetLegacyOverds(plan.LegacyOverds.ValueBool())
	}

	if plan.LimitBcast.IsNull() || plan.LimitBcast.IsUnknown() {
		unset["-limit_bcast"] = ""
	} else {
		data.SetLimitBcast(plan.LimitBcast.ValueBool())
	}

	if plan.LimitProbeResponse.IsNull() || plan.LimitProbeResponse.IsUnknown() {
		unset["-limit_probe_response"] = ""
	} else {
		data.SetLimitProbeResponse(plan.LimitProbeResponse.ValueBool())
	}

	if plan.MaxIdletime.IsNull() || plan.MaxIdletime.IsUnknown() {
		unset["-max_idletime"] = ""
	} else {
		data.SetMaxIdletime(int32(plan.MaxIdletime.ValueInt64()))
	}

	if plan.MistNac.IsNull() || plan.MistNac.IsUnknown() {
		unset["-mist_nac"] = ""
	} else {
		mist_nac := mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
		data.SetMistNac(mist_nac)
	}

	if plan.MxtunnelIds.IsNull() || plan.MxtunnelIds.IsUnknown() {
		unset["-mxtunnel_ids"] = ""
	} else {
		data.SetMxtunnelIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.MxtunnelIds))
	}

	if plan.MxtunnelName.IsNull() || plan.MxtunnelName.IsUnknown() {
		unset["-mxtunnel_name"] = ""
	} else {
		data.SetMxtunnelName(mist_transform.ListOfStringTerraformToSdk(ctx, plan.MxtunnelName))
	}

	if plan.NoStaticDns.IsNull() || plan.NoStaticDns.IsUnknown() {
		unset["-no_static_dns"] = ""
	} else {
		data.SetNoStaticDns(plan.NoStaticDns.ValueBool())
	}

	if plan.NoStaticIp.IsNull() || plan.NoStaticIp.IsUnknown() {
		unset["-no_static_ip"] = ""
	} else {
		data.SetNoStaticIp(plan.NoStaticIp.ValueBool())
	}

	if plan.Portal.IsNull() || plan.Portal.IsUnknown() {
		unset["-portal"] = ""
	} else {
		portal := portalTerraformToSdk(ctx, &diags, plan.Portal)
		data.SetPortal(portal)
	}

	if plan.PortalAllowedHostnames.IsNull() || plan.PortalAllowedHostnames.IsUnknown() {
		unset["-portal_allowed_hostnames"] = ""
	} else {
		data.SetPortalAllowedHostnames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.PortalAllowedHostnames))
	}

	if plan.PortalAllowedSubnets.IsNull() || plan.PortalAllowedSubnets.IsUnknown() {
		unset["-portal_allowed_subnets"] = ""
	} else {
		data.SetPortalAllowedSubnets(mist_transform.ListOfStringTerraformToSdk(ctx, plan.PortalAllowedSubnets))
	}

	if plan.PortalDeniedHostnames.IsNull() || plan.PortalDeniedHostnames.IsUnknown() {
		unset["-portal_denied_hostnames"] = ""
	} else {
		data.SetPortalDeniedHostnames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.PortalDeniedHostnames))
	}

	if plan.Qos.IsNull() || plan.Qos.IsUnknown() {
		unset["-qos"] = ""
	} else {
		qos := qosTerraformToSdk(ctx, &diags, plan.Qos)
		data.SetQos(qos)
	}

	if plan.Radsec.IsNull() || plan.Radsec.IsUnknown() {
		unset["-radsec"] = ""
	} else {
		radesc := radsecTerraformToSdk(ctx, &diags, plan.Radsec)
		data.SetRadsec(radesc)
	}

	if plan.RoamMode.IsNull() || plan.RoamMode.IsUnknown() {
		unset["-roam_mode"] = ""
	} else {
		data.SetRoamMode(mistsdkgo.WlanRoamMode(plan.RoamMode.ValueString()))
	}

	if plan.Schedule.IsNull() || plan.Schedule.IsUnknown() {
		unset["-schedule"] = ""
	} else {
		schedule := scheduleTerraformToSdk(ctx, &diags, plan.Schedule)
		data.SetSchedule(schedule)
	}

	if plan.SleExcluded.IsNull() || plan.SleExcluded.IsUnknown() {
		unset["-sle_excluded"] = ""
	} else {
		data.SetSleExcluded(plan.SleExcluded.ValueBool())
	}

	if plan.UseEapolV1.IsNull() || plan.UseEapolV1.IsUnknown() {
		unset["-use_eapol_v1"] = ""
	} else {
		data.SetUseEapolV1(plan.UseEapolV1.ValueBool())
	}

	if plan.VlanEnabled.IsNull() || plan.VlanEnabled.IsUnknown() {
		unset["-vlan_enabled"] = ""
	} else {
		data.SetVlanEnabled(plan.VlanEnabled.ValueBool())
	}

	if plan.VlanId.IsNull() || plan.VlanId.IsUnknown() {
		unset["-vlan_id"] = ""
	} else {
		data.SetVlanId(int32(plan.VlanId.ValueInt64()))
	}

	if plan.VlanIds.IsNull() || plan.VlanIds.IsUnknown() {
		unset["-vlan_ids"] = ""
	} else {
		vlan_ids := vlanIdsTerraformToSdk(ctx, &diags, plan.VlanIds)
		data.SetVlanIds(vlan_ids)
	}

	if plan.VlanPooling.IsNull() || plan.VlanPooling.IsUnknown() {
		unset["-vlan_pooling"] = ""
	} else {
		data.SetVlanPooling(plan.VlanPooling.ValueBool())
	}

	if plan.WlanLimitDown.IsNull() || plan.WlanLimitDown.IsUnknown() {
		unset["-wlan_limit_down"] = ""
	} else {
		data.SetWlanLimitDown(int32(plan.WlanLimitDown.ValueInt64()))
	}

	if plan.WlanLimitDownEnabled.IsNull() || plan.WlanLimitDownEnabled.IsUnknown() {
		unset["-wlan_limit_down_enabled"] = ""
	} else {
		data.SetWlanLimitDownEnabled(plan.WlanLimitDownEnabled.ValueBool())
	}

	if plan.WlanLimitUp.IsNull() || plan.WlanLimitUp.IsUnknown() {
		unset["-wlan_limit_up"] = ""
	} else {
		data.SetWlanLimitUp(int32(plan.WlanLimitUp.ValueInt64()))
	}

	if plan.WlanLimitUpEnabled.IsNull() || plan.WlanLimitUpEnabled.IsUnknown() {
		unset["-wlan_limit_up_enabled"] = ""
	} else {
		data.SetWlanLimitUpEnabled(plan.WlanLimitUpEnabled.ValueBool())
	}

	if plan.WxtagIds.IsNull() || plan.WxtagIds.IsUnknown() {
		unset["-wxtag_ids"] = ""
	} else {
		data.SetWxtagIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.WxtagIds))
	}

	if plan.WxtunnelId.IsNull() || plan.WxtunnelId.IsUnknown() {
		unset["-wxtunnel_id"] = ""
	} else {
		data.SetWxtunnelId(plan.WxtunnelId.ValueString())
	}

	if plan.WxtunnelRemoteId.IsNull() || plan.WxtunnelRemoteId.IsUnknown() {
		unset["-wxtunnel_remote_id"] = ""
	} else {
		data.SetWxtunnelRemoteId(plan.WxtunnelRemoteId.ValueString())
	}

	data.AdditionalProperties = unset

	return data, diags
}
