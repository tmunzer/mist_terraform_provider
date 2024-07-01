package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"
)

func TerraformToSdk(ctx context.Context, plan *NetworktemplateModel) (mistsdkgo.NetworkTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := *mistsdkgo.NewNetworkTemplate()
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

	extraRoutes := extraRoutesTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
	data.SetExtraRoutes(extraRoutes)

	mistNac := mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
	data.SetMistNac(mistNac)

	ntpServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)
	data.SetNtpServers(ntpServers)

	networks := NetworksTerraformToSdk(ctx, &diags, plan.Networks)
	data.SetNetworks(networks)

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
