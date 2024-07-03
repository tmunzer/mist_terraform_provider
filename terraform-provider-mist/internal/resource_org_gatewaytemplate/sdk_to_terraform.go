package resource_org_gatewaytemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.GatewayTemplate) (OrgGatewaytemplateModel, diag.Diagnostics) {
	var state OrgGatewaytemplateModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())

	state.AdditionalConfigCmds = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAdditionalConfigCmds())

	state.BgpConfig = bgpConfigSdkToTerraform(ctx, &diags, data.GetBgpConfig())

	state.DhcpdConfig = dhcpdConfigSdkToTerraform(ctx, &diags, data.GetDhcpdConfig())

	if data.DnsOverride != nil {
		state.DnsOverride = types.BoolValue(*data.DnsOverride)
	}

	state.DnsServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.DnsServers)

	state.DnsSuffix = mist_transform.ListOfStringSdkToTerraform(ctx, data.DnsSuffix)

	state.ExtraRoutes = extraRoutesSdkToTerraform(ctx, &diags, data.GetExtraRoutes())

	state.IdpProfiles = idpProfileSdkToTerraform(ctx, &diags, data.GetIdpProfiles())

	state.IpConfigs = ipConfigsSdkToTerraform(ctx, &diags, data.GetIpConfigs())

	state.Networks = NetworksSdkToTerraform(ctx, &diags, data.GetNetworks())

	if data.NtpOverride != nil {
		state.NtpOverride = types.BoolValue(*data.NtpOverride)
	}

	state.NtpServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.NtpServers)

	state.PathPreferences = pathPreferencesSdkToTerraform(ctx, &diags, data.GetPathPreferences())

	state.PortConfig = portConfigSdkToTerraform(ctx, &diags, data.GetPortConfig())

	state.RouterId = types.StringValue(data.GetRouterId())

	state.RoutingPolicies = routingPolociesSdkToTerraform(ctx, &diags, data.GetRoutingPolicies())

	state.ServicePolicies = servicePoliciesSdkToTerraform(ctx, &diags, data.GetServicePolicies())

	state.TunnelConfigs = tunnelConfigsSdkToTerraform(ctx, &diags, data.GetTunnelConfigs())

	state.TunnelProviderOptions = tunnelProviderSdkToTerraform(ctx, &diags, data.GetTunnelProviderOptions())

	state.Type = types.StringValue(string(data.GetType()))

	return state, diags
}
