package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func TerraformToSdk(ctx context.Context, plan *DeviceSwitchModel) (models.MistDevice, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data := models.NewDeviceSwitch()
	data.SetDeviceprofileId(plan.DeviceprofileId.ValueString())
	data.SetMapId(plan.MapId.ValueString())
	data.SetName(plan.Name.ValueString())
	data.SetNotes(plan.Notes.ValueString())

	if plan.AclPolicies.IsNull() || plan.AclPolicies.IsUnknown() {
		unset["-acl_policies"] = ""
	} else {
		aclPolicies := aclPoliciesTerraformToSdk(ctx, &diags, plan.AclPolicies)
		data.SetAclPolicies(aclPolicies)
	}

	if plan.AclTags.IsNull() || plan.AclTags.IsUnknown() {
		unset["-acl_tags"] = ""
	} else {
		aclTAgs := actTagsTerraformToSdk(ctx, &diags, plan.AclTags)
		data.SetAclTags(aclTAgs)
	}

	addictionalConfigCmds := mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)
	data.SetAdditionalConfigCmds(addictionalConfigCmds)

	if plan.DhcpSnooping.IsNull() || plan.DhcpSnooping.IsUnknown() {
		unset["-dhcp_snooping"] = ""
	} else {
		dhcpSnooping := dhcpSnoopingTerraformToSdk(ctx, &diags, plan.DhcpSnooping)
		data.SetDhcpSnooping(dhcpSnooping)
	}

	if plan.DhcpdConfig.IsNull() || plan.DhcpdConfig.IsUnknown() {
		unset["-dhcpd_config"] = ""
	} else {
		dhcpd_config := dhcpdConfigTerraformToSdk(ctx, &diags, plan.DhcpdConfig)
		data.SetDhcpdConfig(dhcpd_config)
	}

	data.SetDisableAutoConfig(plan.DisableAutoConfig.ValueBool())

	dnsServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)
	data.SetDnsServers(dnsServers)

	dnsSuffix := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
	data.SetDnsSuffix(dnsSuffix)

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		extra_routes := extraRoutesTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
		data.SetExtraRoutes(extra_routes)
	}

	if plan.ExtraRoutes6.IsNull() || plan.ExtraRoutes6.IsUnknown() {
		unset["-extra_routes6"] = ""
	} else {
		extra_routes6 := extraRoutes6TerraformToSdk(ctx, &diags, plan.ExtraRoutes6)
		data.SetExtraRoutes6(extra_routes6)
	}

	if plan.IpConfig.IsNull() || plan.IpConfig.IsUnknown() {
		unset["-ip_config"] = ""
	} else {
		ip_config := ipConfigTerraformToSdk(ctx, &diags, plan.IpConfig)
		data.SetIpConfig(ip_config)
	}

	data.SetManaged(plan.Managed.ValueBool())

	if plan.MistNac.IsNull() || plan.MistNac.IsUnknown() {
		unset["-mist_nac"] = ""
	} else {
		mistNac := mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
		data.SetMistNac(mistNac)
	}

	if plan.Networks.IsNull() || plan.Networks.IsUnknown() {
		unset["-networks"] = ""
	} else {
		networks := NetworksTerraformToSdk(ctx, &diags, plan.Networks)
		data.SetNetworks(networks)
	}

	ntpServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)
	data.SetNtpServers(ntpServers)

	if plan.OobIpConfig.IsNull() || plan.OobIpConfig.IsUnknown() {
		unset["-oob_ip_config"] = ""
	} else {
		oob_ip_config := oobIpConfigTerraformToSdk(ctx, &diags, plan.OobIpConfig)
		data.SetOobIpConfig(oob_ip_config)
	}

	if plan.OspfConfig.IsNull() || plan.OspfConfig.IsUnknown() {
		unset["-ospf_config"] = ""
	} else {
		ospf_config := ospfConfigTerraformToSdk(ctx, &diags, plan.OspfConfig)
		data.SetOspfConfig(ospf_config)
	}

	if plan.OtherIpConfigs.IsNull() || plan.OtherIpConfigs.IsUnknown() {
		unset["-other_ip_configs"] = ""
	} else {
		other_ip_configs := otherIpConfigTerraformToSdk(ctx, &diags, plan.OtherIpConfigs)
		data.SetOtherIpConfigs(other_ip_configs)
	}

	if plan.PortConfig.IsNull() || plan.PortConfig.IsUnknown() {
		unset["-port_config"] = ""
	} else {
		port_config := portConfigTerraformToSdk(ctx, &diags, plan.PortConfig)
		data.SetPortConfig(port_config)
	}

	if plan.PortMirroring.IsNull() || plan.PortMirroring.IsUnknown() {
		unset["-port_mirrorings"] = ""
	} else {
		port_mirroring := portMirroringTerraformToSdk(ctx, &diags, plan.PortMirroring)
		data.SetPortMirroring(port_mirroring)
	}

	if plan.PortUsages.IsNull() || plan.PortUsages.IsUnknown() {
		unset["-port_usages"] = ""
	} else {
		port_usage := portUsageTerraformToSdk(ctx, &diags, plan.PortUsages)
		data.SetPortUsages(port_usage)
	}

	if plan.RadiusConfig.IsNull() || plan.RadiusConfig.IsUnknown() {
		unset["-radius_config"] = ""
	} else {
		radius_config := radiusConfigTerraformToSdk(ctx, &diags, plan.RadiusConfig)
		data.SetRadiusConfig(radius_config)
	}

	if plan.RemoteSyslog.IsNull() || plan.RemoteSyslog.IsUnknown() {
		unset["-remote_syslog"] = ""
	} else {
		remote_syslog := remoteSyslogTerraformToSdk(ctx, &diags, plan.RemoteSyslog)
		data.SetRemoteSyslog(remote_syslog)
	}

	data.SetRole(plan.Role.ValueString())
	data.SetRouterId(plan.RouterId.ValueString())

	if plan.SnmpConfig.IsNull() || plan.SnmpConfig.IsUnknown() {
		unset["-snmp_config"] = ""
	} else {
		snmp_config := snmpConfigSyslogTerraformToSdk(ctx, &diags, plan.SnmpConfig)
		data.SetSnmpConfig(snmp_config)
	}

	if plan.StpConfig.IsNull() || plan.StpConfig.IsUnknown() {
		unset["-stp_config"] = ""
	} else {
		stp_config := stpConfigTerraformToSdk(ctx, &diags, plan.StpConfig)
		data.SetStpConfig(stp_config)
	}

	if plan.SwitchMgmt.IsNull() || plan.SwitchMgmt.IsUnknown() {
		unset["-switch_mgmt"] = ""
	} else {
		switch_mgmt := switchMgmtTerraformToSdk(ctx, &diags, plan.SwitchMgmt)
		data.SetSwitchMgmt(switch_mgmt)
	}

	data.SetUseRouterIdAsSourceIp(plan.UseRouterIdAsSourceIp.ValueBool())

	if plan.Vars.IsNull() || plan.Vars.IsUnknown() {
		unset["-vars"] = ""
	} else {
		vars := varsTerraformToSdk(ctx, &diags, plan.Vars)
		data.SetVars(vars)
	}

	if plan.VirtualChassis.IsNull() || plan.VirtualChassis.IsUnknown() {
		unset["-virtual_chassis"] = ""
	} else {
		virtual_chassis := virtualChassisTerraformToSdk(ctx, &diags, plan.VirtualChassis)
		data.SetVirtualChassis(virtual_chassis)
	}

	if plan.VrfConfig.IsNull() || plan.VrfConfig.IsUnknown() {
		unset["-vrf_config"] = ""
	} else {
		vrfConfig := vrfConfigTerraformToSdk(ctx, &diags, plan.VrfConfig)
		data.SetVrfConfig(vrfConfig)
	}

	if plan.VrfInstances.IsNull() || plan.VrfInstances.IsUnknown() {
		unset["-vrf_instances"] = ""
	} else {
		vrfInstance := vrfInstancesTerraformToSdk(ctx, &diags, plan.VrfInstances)
		data.SetVrfInstances(vrfInstance)
	}

	if plan.VrrpConfig.IsNull() || plan.VrrpConfig.IsUnknown() {
		unset["-vrrp_config"] = ""
	} else {
		vrrp_config := vrrpTerraformToSdk(ctx, &diags, plan.VrrpConfig)
		data.SetVrrpConfig(vrrp_config)
	}

	data.SetX(plan.X.ValueFloat64())
	data.SetY(plan.Y.ValueFloat64())

	data.AdditionalProperties = unset

	var mist_device models.MistDevice
	mist_device.DeviceSwitch = data

	return mist_device, diags
}
