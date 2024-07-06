package resource_org_gatewaytemplate

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.GatewayTemplate) (OrgGatewaytemplateModel, diag.Diagnostics) {
	var state OrgGatewaytemplateModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Name = types.StringValue(data.Name)

	state.AdditionalConfigCmds = mist_transform.ListOfStringSdkToTerraform(ctx, data.AdditionalConfigCmds)

	state.BgpConfig = bgpConfigSdkToTerraform(ctx, &diags, data.BgpConfig)

	state.DhcpdConfig = dhcpdConfigSdkToTerraform(ctx, &diags, *data.DhcpdConfig)

	if data.DnsOverride != nil {
		state.DnsOverride = types.BoolValue(*data.DnsOverride)
	}

	state.DnsServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.DnsServers)

	state.DnsSuffix = mist_transform.ListOfStringSdkToTerraform(ctx, data.DnsSuffix)

	state.ExtraRoutes = extraRoutesSdkToTerraform(ctx, &diags, data.ExtraRoutes)

	state.IdpProfiles = idpProfileSdkToTerraform(ctx, &diags, data.IdpProfiles)

	state.IpConfigs = ipConfigsSdkToTerraform(ctx, &diags, data.IpConfigs)

	state.Networks = NetworksSdkToTerraform(ctx, &diags, data.Networks)

	if data.NtpOverride != nil {
		state.NtpOverride = types.BoolValue(*data.NtpOverride)
	}

	state.NtpServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.NtpServers)

	state.PathPreferences = pathPreferencesSdkToTerraform(ctx, &diags, data.PathPreferences)

	state.PortConfig = portConfigSdkToTerraform(ctx, &diags, data.PortConfig)

	state.RouterId = types.StringValue(*data.RouterId)

	state.RoutingPolicies = routingPolociesSdkToTerraform(ctx, &diags, data.RoutingPolicies)

	state.ServicePolicies = servicePoliciesSdkToTerraform(ctx, &diags, data.ServicePolicies)

	state.TunnelConfigs = tunnelConfigsSdkToTerraform(ctx, &diags, data.TunnelConfigs)

	state.TunnelProviderOptions = tunnelProviderSdkToTerraform(ctx, &diags, data.TunnelProviderOptions)

	state.Type = types.StringValue(string(*data.Type))

	return state, diags
}
