package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func SdkToTerraform(ctx context.Context, data *models.SiteSetting) (SiteNetworktemplateModel, diag.Diagnostics) {
	var state SiteNetworktemplateModel
	var diags diag.Diagnostics

	state.AclPolicies = aclPoliciesSdkToTerraform(ctx, &diags, data.AclPolicies)

	state.AclTags = aclTagsSdkToTerraform(ctx, &diags, data.AclTags)

	state.AdditionalConfigCmds = mist_transform.ListOfStringSdkToTerraform(ctx, data.AdditionalConfigCmds)

	state.DhcpSnooping = dhcpSnoopingSdkToTerraform(ctx, &diags, data.DhcpSnooping)

	state.DnsServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.DnsServers)

	state.DnsSuffix = mist_transform.ListOfStringSdkToTerraform(ctx, data.DnsSuffix)

	state.ExtraRoutes = extraRoutesSdkToTerraform(ctx, &diags, data.ExtraRoutes)

	state.ExtraRoutes6 = extraRoutes6SdkToTerraform(ctx, &diags, data.ExtraRoutes6)

	state.MistNac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)

	state.NtpServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.NtpServers)

	state.Networks = NetworksSdkToTerraform(ctx, &diags, data.Networks)

	state.PortMirroring = portMirroringSdkToTerraform(ctx, &diags, data.PortMirroring)

	state.PortUsages = portUsagesSdkToTerraform(ctx, &diags, data.PortUsages)

	state.RemoteSyslog = remoteSyslogSdkToTerraform(ctx, &diags, data.RemoteSyslog)

	state.RadiusConfig = radiusConfigSdkToTerraform(ctx, &diags, data.RadiusConfig)

	state.SnmpConfig = snmpConfigSdkToTerraform(ctx, &diags, data.SnmpConfig)

	state.SwitchMatching = switchMatchingSdkToTerraform(ctx, &diags, data.SwitchMatching)

	state.SwitchMgmt = switchMgmtSdkToTerraform(ctx, &diags, data.SwitchMgmt)

	//state.VrfConfig = vrfConfigSdkToTerraform(ctx, &diags, data.VrfConfig)

	state.VrfInstances = vrfInstancesSdkToTerraform(ctx, &diags, data.VrfInstances)
	return state, diags
}
