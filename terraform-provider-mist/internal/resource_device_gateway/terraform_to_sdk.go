package resource_device_gateway

import (
	"context"
	"mistapi/models"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *DeviceGatewayModel) (models.MistDevice, diag.Diagnostics) {
	data := models.DeviceGateway{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	deviceprofile_id, e := uuid.Parse(plan.DeviceprofileId.ValueString())
	if e == nil {
		data.DeviceprofileId = models.ToPointer(deviceprofile_id)
	} else {
		unset["deviceprofile_id"] = nil
	}
	map_id, e := uuid.Parse(plan.MapId.ValueString())
	if e == nil {
		data.MapId = &map_id
	} else {
		unset["map_id"] = nil
	}

	data.Name = plan.Name.ValueStringPointer()
	data.Notes = plan.Notes.ValueStringPointer()

	if plan.AdditionalConfigCmds.IsNull() || plan.AdditionalConfigCmds.IsUnknown() {
		data.AdditionalConfigCmds = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)
	} else {
		unset["-additional_config_cmds"] = ""
	}

	if plan.BgpConfig.IsNull() || plan.BgpConfig.IsUnknown() {
		unset["-bgp_config"] = ""
	} else {
		bgp_config := bgpConfigTerraformToSdk(ctx, &diags, plan.BgpConfig)
		data.BgpConfig = bgp_config
	}

	if plan.DhcpdConfig.IsNull() || plan.DhcpdConfig.IsUnknown() {
		unset["-dhcpd_config"] = ""
	} else {
		dhcpd_config := dhcpdConfigTerraformToSdk(ctx, &diags, plan.DhcpdConfig)
		data.DhcpdConfig = &dhcpd_config
	}

	data.DnsServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)

	data.DnsSuffix = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		extra_routes := extraRouteTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
		data.ExtraRoutes = extra_routes
	}

	if plan.IdpProfiles.IsNull() || plan.IdpProfiles.IsUnknown() {
		unset["-idp_profiles"] = ""
	} else {
		idp_profiles := idpProfileTerraformToSdk(ctx, &diags, plan.IdpProfiles)
		data.IdpProfiles = idp_profiles
	}

	if plan.IpConfigs.IsNull() || plan.IpConfigs.IsUnknown() {
		unset["-ip_configs"] = ""
	} else {
		ip_configs := ipConfigsTerraformToSdk(ctx, &diags, plan.IpConfigs)
		data.IpConfigs = ip_configs
	}

	if plan.Networks.IsNull() || plan.Networks.IsUnknown() {
		unset["-networks"] = ""
	} else {
		networks := networksTerraformToSdk(ctx, &diags, plan.Networks)
		data.Networks = networks
	}

	data.NtpServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)

	if plan.OobIpConfig.IsNull() || plan.OobIpConfig.IsUnknown() {
		unset["-oob_ip_config"] = ""
	} else {
		oob_ip_config := oobIpConfigTerraformToSdk(ctx, &diags, plan.OobIpConfig)
		data.OobIpConfig = oob_ip_config
	}

	if plan.PathPreferences.IsNull() || plan.PathPreferences.IsUnknown() {
		unset["-path_preferences"] = ""
	} else {
		path_preferences := pathPreferencesTerraformToSdk(ctx, &diags, plan.PathPreferences)
		data.PathPreferences = path_preferences
	}

	if plan.PortConfig.IsNull() || plan.PortConfig.IsUnknown() {
		unset["-port_config"] = ""
	} else {
		port_config := portConfigTerraformToSdk(ctx, &diags, plan.PortConfig)
		data.PortConfig = port_config
	}

	data.RouterId = plan.RouterId.ValueStringPointer()

	if plan.RoutingPolicies.IsNull() || plan.RoutingPolicies.IsUnknown() {
		unset["-routing_policies"] = ""
	} else {
		routing_policies := routingPoliciesTerraformToSdk(ctx, &diags, plan.RoutingPolicies)
		data.RoutingPolicies = routing_policies
	}

	if plan.ServicePolicies.IsNull() || plan.ServicePolicies.IsUnknown() {
		unset["-service_policies"] = ""
	} else {
		service_policies := servicePoliciesTerraformToSdk(ctx, &diags, plan.ServicePolicies)
		data.ServicePolicies = service_policies
	}

	if plan.TunnelConfigs.IsNull() || plan.TunnelConfigs.IsUnknown() {
		unset["-tunnel_configs"] = ""
	} else {
		tunnel_configs := tunnelConfigsTerraformToSdk(ctx, &diags, plan.TunnelConfigs)
		data.TunnelConfigs = tunnel_configs
	}

	if plan.TunnelProviderOptions.IsNull() || plan.TunnelProviderOptions.IsUnknown() {
		unset["-tunnel_provider_options"] = ""
	} else {
		tunnel_provider_options := tunnelProviderOptionsTerraformToSdk(ctx, &diags, plan.TunnelProviderOptions)
		data.TunnelProviderOptions = &tunnel_provider_options
	}

	if !plan.Vars.IsNull() && !plan.Vars.IsUnknown() {
		data.Vars = varsTerraformToSdk(ctx, &diags, plan.Vars)
	} else {
		unset["-vars"] = ""
	}

	if !plan.X.IsNull() && !plan.X.IsUnknown() {
		data.X = plan.X.ValueFloat64Pointer()
	} else {
		unset["-x"] = ""
	}
	if !plan.Y.IsNull() && !plan.Y.IsUnknown() {
		data.Y = plan.Y.ValueFloat64Pointer()
	} else {
		unset["-y"] = ""
	}

	data.AdditionalProperties = unset

	mist_device := models.MistDeviceContainer.FromDeviceGateway(data)
	return mist_device, diags
}
