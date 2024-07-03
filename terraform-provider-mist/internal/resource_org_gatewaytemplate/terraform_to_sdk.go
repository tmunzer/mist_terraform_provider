package resource_org_gatewaytemplate

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgGatewaytemplateModel) (mistapigo.GatewayTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data := *mistapigo.NewGatewayTemplate(plan.Name.ValueString())
	data.SetId(plan.Id.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())

	additional_config_cmds := mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)
	data.SetAdditionalConfigCmds(additional_config_cmds)

	if plan.BgpConfig.IsNull() || plan.BgpConfig.IsUnknown() {
		unset["-bgp_config"] = ""
	} else {
		bgp_config := bgpConfigTerraformToSdk(ctx, &diags, plan.BgpConfig)
		data.SetBgpConfig(bgp_config)
	}

	if plan.DhcpdConfig.IsNull() || plan.DhcpdConfig.IsUnknown() {
		unset["-dhcpd_config"] = ""
	} else {
		dhcpd_config := dhcpdConfigTerraformToSdk(ctx, &diags, plan.DhcpdConfig)
		data.SetDhcpdConfig(dhcpd_config)
	}

	data.SetDnsOverride(plan.DnsOverride.ValueBool())

	data.SetDnsServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers))

	data.SetDnsSuffix(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix))

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		extra_routes := extraRouteTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
		data.SetExtraRoutes(extra_routes)
	}

	if plan.IdpProfiles.IsNull() || plan.IdpProfiles.IsUnknown() {
		unset["-idp_profiles"] = ""
	} else {
		idp_profiles := idpProfileTerraformToSdk(ctx, &diags, plan.IdpProfiles)
		data.SetIdpProfiles(idp_profiles)
	}

	if plan.IpConfigs.IsNull() || plan.IpConfigs.IsUnknown() {
		unset["-ip_configs"] = ""
	} else {
		ip_configs := ipConfigsTerraformToSdk(ctx, &diags, plan.IpConfigs)
		data.SetIpConfigs(ip_configs)
	}

	if plan.Networks.IsNull() || plan.Networks.IsUnknown() {
		unset["-networks"] = ""
	} else {
		networks := networksTerraformToSdk(ctx, &diags, plan.Networks)
		data.SetNetworks(networks)
	}

	data.SetNtpOverride(plan.NtpOverride.ValueBool())

	data.SetNtpServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers))

	if plan.PathPreferences.IsNull() || plan.PathPreferences.IsUnknown() {
		unset["-path_preferences"] = ""
	} else {
		path_preferences := pathPreferencesTerraformToSdk(ctx, &diags, plan.PathPreferences)
		data.SetPathPreferences(path_preferences)
	}

	if plan.PortConfig.IsNull() || plan.PortConfig.IsUnknown() {
		unset["-port_config"] = ""
	} else {
		port_config := portConfigTerraformToSdk(ctx, &diags, plan.PortConfig)
		data.SetPortConfig(port_config)
	}

	data.SetRouterId(plan.RouterId.ValueString())

	if plan.RoutingPolicies.IsNull() || plan.RoutingPolicies.IsUnknown() {
		unset["-routing_policies"] = ""
	} else {
		routing_policies := routingPoliciesTerraformToSdk(ctx, &diags, plan.RoutingPolicies)
		data.SetRoutingPolicies(routing_policies)
	}

	if plan.ServicePolicies.IsNull() || plan.ServicePolicies.IsUnknown() {
		unset["-service_policies"] = ""
	} else {
		service_policies := servicePoliciesTerraformToSdk(ctx, &diags, plan.ServicePolicies)
		data.SetServicePolicies(service_policies)
	}

	if plan.TunnelConfigs.IsNull() || plan.TunnelConfigs.IsUnknown() {
		unset["-tunnel_configs"] = ""
	} else {
		tunnel_configs := tunnelConfigsTerraformToSdk(ctx, &diags, plan.TunnelConfigs)
		data.SetTunnelConfigs(tunnel_configs)
	}

	if plan.TunnelProviderOptions.IsNull() || plan.TunnelProviderOptions.IsUnknown() {
		unset["-tunnel_provider_options"] = ""
	} else {
		tunnel_provider_options := tunnelProviderOptionsTerraformToSdk(ctx, &diags, plan.TunnelProviderOptions)
		data.SetTunnelProviderOptions(tunnel_provider_options)
	}

	data.SetType(mistapigo.GatewayTemplateType(plan.Type.ValueString()))

	return data, diags
}
