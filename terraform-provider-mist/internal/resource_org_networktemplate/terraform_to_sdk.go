package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func TerraformToSdk(ctx context.Context, plan *OrgNetworktemplateModel) (mistapigo.NetworkTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := *mistapigo.NewNetworkTemplate()
	data.SetName(plan.Name.ValueString())

	aclPolicies := aclPoliciesTerraformToSdk(ctx, &diags, plan.AclPolicies)
	data.SetAclPolicies(aclPolicies)

	aclTAgs := actTagsTerraformToSdk(ctx, &diags, plan.AclTags)
	data.SetAclTags(aclTAgs)

	addictionalConfigCmds := mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)
	data.SetAdditionalConfigCmds(addictionalConfigCmds)

	dnsServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)
	data.SetDnsServers(dnsServers)

	dnsSuffix := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
	data.SetDnsSuffix(dnsSuffix)

	dhcpSnooping := dhcpSnoopingTerraformToSdk(ctx, &diags, plan.DhcpSnooping)
	data.SetDhcpSnooping(dhcpSnooping)

<<<<<<< Updated upstream
	extraRoutes := extraRoutesTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
	data.SetExtraRoutes(extraRoutes)
=======
	if plan.ExtraRoutes6.IsNull() || plan.ExtraRoutes6.IsUnknown() {
		unset["-extra_routes6"] = ""
	} else {
		extra_routes6 := extraRoutes6TerraformToSdk(ctx, &diags, plan.ExtraRoutes6)
		data.SetExtraRoutes6(extra_routes6)
	}

	if plan.MistNac.IsNull() || plan.MistNac.IsUnknown() {
		unset["-mist_nac"] = ""
	} else {
		mistNac := mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
		data.SetMistNac(mistNac)
	}
>>>>>>> Stashed changes

	mistNac := mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
	data.SetMistNac(mistNac)

	ntpServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)
	data.SetNtpServers(ntpServers)

<<<<<<< Updated upstream
	networks := NetworksTerraformToSdk(ctx, &diags, plan.Networks)
	data.SetNetworks(networks)
=======
	if plan.PortMirroring.IsNull() || plan.PortMirroring.IsUnknown() {
		unset["-port_mirrorings"] = ""
	} else {
		port_mirroring := portMirroringTerraformToSdk(ctx, &diags, plan.PortMirroring)
		data.SetPortMirroring(port_mirroring)
	}
>>>>>>> Stashed changes

	port_usage := portUsageTerraformToSdk(ctx, &diags, plan.PortUsages)
	data.SetPortUsages(port_usage)

	port_mirroring := portMirroringTerraformToSdk(ctx, &diags, plan.PortMirrorings)
	data.SetPortMirrorings(port_mirroring)

	radius_config := radiusConfigTerraformToSdk(ctx, &diags, plan.RadiusConfig)
	data.SetRadiusConfig(radius_config)

	remote_syslog := remoteSyslogTerraformToSdk(ctx, &diags, plan.RemoteSyslog)
	data.SetRemoteSyslog(remote_syslog)

	switch_matching := switchMatchingTerraformToSdk(ctx, &diags, plan.SwitchMatching)
	data.SetSwitchMatching(switch_matching)

	switch_mgmt := switchMgmtTerraformToSdk(ctx, &diags, plan.SwitchMgmt)
	data.SetSwitchMgmt(switch_mgmt)

	vrfConfig := vrfConfigTerraformToSdk(ctx, &diags, plan.VrfConfig)
	data.SetVrfConfig(vrfConfig)

	vrfInstance := vrfInstancesTerraformToSdk(ctx, &diags, plan.VrfInstances)
	data.SetVrfInstances(vrfInstance)

	return data, diags
}
