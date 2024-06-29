package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	mist_transform "terraform-provider-mist/internal/provider/utils/transform"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *GatewaytemplateModel) (mistsdkgo.GatewayTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewGatewayTemplate(plan.Name.ValueString())
	data.SetId(plan.Id.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())

	additional_config_cmds := mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)
	data.SetAdditionalConfigCmds(additional_config_cmds)

	bgp_config := bgpConfigTerraformToSdk(ctx, &diags, plan.BgpConfig)
	data.SetBgpConfig(bgp_config)

	dhcpd_config := dhcpdConfigTerraformToSdk(ctx, &diags, plan.DhcpdConfig)
	data.SetDhcpdConfig(dhcpd_config)

	data.SetDnsOverride(plan.DnsOverride.ValueBool())

	data.SetDnsServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers))

	data.SetDnsSuffix(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix))

	extra_routes := extraRouteTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
	data.SetExtraRoutes(extra_routes)

	idp_profiles := idpProfileTerraformToSdk(ctx, &diags, plan.IdpProfiles)
	data.SetIdpProfiles(idp_profiles)

	ip_configs := ipConfigsTerraformToSdk(ctx, &diags, plan.IpConfigs)
	data.SetIpConfigs(ip_configs)

	networks := networksTerraformToSdk(ctx, &diags, plan.Networks)
	data.SetNetworks(networks)

	data.SetNtpOverride(plan.NtpOverride.ValueBool())

	data.SetNtpServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers))

	path_preferences := pathPreferencesTerraformToSdk(ctx, &diags, plan.PathPreferences)
	data.SetPathPreferences(path_preferences)

	port_config := portConfigTerraformToSdk(ctx, &diags, plan.PortConfig)
	data.SetPortConfig(port_config)

	data.SetRouterId(plan.RouterId.ValueString())

	routing_policies := routingPoliciesTerraformToSdk(ctx, &diags, plan.RoutingPolicies)
	data.SetRoutingPolicies(routing_policies)

	service_policies := servicePoliciesTerraformToSdk(ctx, &diags, plan.ServicePolicies)
	data.SetServicePolicies(service_policies)

	tunnel_configs := tunnelConfigsTerraformToSdk(ctx, &diags, plan.TunnelConfigs)
	data.SetTunnelConfigs(tunnel_configs)

	tunnel_provider_options := tunnelProviderOptionsTerraformToSdk(ctx, &diags, plan.TunnelProviderOptions)
	data.SetTunnelProviderOptions(tunnel_provider_options)

	data.SetType(mistsdkgo.GatewayTemplateType(plan.Type.ValueString()))

	return data, diags
}
