package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mistapi/internal/provider/utils/transform"
)

func TerraformToSdk(ctx context.Context, plan *NetworktemplateModel) (mistsdkgo.NetworkTemplate, string, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := *mistsdkgo.NewNetworkTemplate()
	data.SetName(plan.Name.ValueString())

	addictionalConfigCmds := mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)
	data.SetAdditionalConfigCmds(addictionalConfigCmds)

	dnsServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)
	data.SetDnsServers(dnsServers)

	dnsSuffix := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
	data.SetDnsSuffix(dnsSuffix)

	dhcpSnooping := dhcpSnoopingTerraformToSdk(ctx, &diags, plan.DhcpSnooping)
	data.SetDhcpSnooping(dhcpSnooping)

	mistNac := mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
	data.SetMistNac(mistNac)

	ntpServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)
	data.SetNtpServers(ntpServers)

	networks := networksTerraformToSdk(ctx, &diags, plan.Networks)
	data.SetNetworks(networks)

	port_usage := portUsageTerraformToSdk(ctx, &diags, plan.PortUsages)
	data.SetPortUsages(port_usage)

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

	var orgId = plan.OrgId.ValueString()
	return data, orgId, diags
}
