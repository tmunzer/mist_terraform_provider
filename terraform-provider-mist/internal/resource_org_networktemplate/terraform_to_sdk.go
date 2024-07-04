package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func TerraformToSdk(ctx context.Context, plan *OrgNetworktemplateModel) (mistapigo.NetworkTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data := *mistapigo.NewNetworkTemplate()
	data.SetName(plan.Name.ValueString())

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

	dnsServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)
	data.SetDnsServers(dnsServers)

	dnsSuffix := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
	data.SetDnsSuffix(dnsSuffix)

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		extraRoutes := extraRoutesTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
		data.SetExtraRoutes(extraRoutes)
	}

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

	if plan.PortMirrorings.IsNull() || plan.PortMirrorings.IsUnknown() {
		unset["-port_mirrorings"] = ""
	} else {
		port_mirroring := portMirroringTerraformToSdk(ctx, &diags, plan.PortMirrorings)
		data.SetPortMirrorings(port_mirroring)
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

	if plan.SnmpConfig.IsNull() || plan.SnmpConfig.IsUnknown() {
		unset["-snmp_config"] = ""
	} else {
		snmp_config := snmpConfigSyslogTerraformToSdk(ctx, &diags, plan.SnmpConfig)
		data.SetSnmpConfig(snmp_config)
	}

	if plan.SwitchMatching.IsNull() || plan.SwitchMatching.IsUnknown() {
		unset["-switch_matching"] = ""
	} else {
		switch_matching := switchMatchingTerraformToSdk(ctx, &diags, plan.SwitchMatching)
		data.SetSwitchMatching(switch_matching)
	}

	if plan.SwitchMgmt.IsNull() || plan.SwitchMgmt.IsUnknown() {
		unset["-switch_mgmt"] = ""
	} else {
		switch_mgmt := switchMgmtTerraformToSdk(ctx, &diags, plan.SwitchMgmt)
		data.SetSwitchMgmt(switch_mgmt)
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

	data.AdditionalProperties = unset
	return data, diags
}
