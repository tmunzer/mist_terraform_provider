package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.SiteSetting) (SiteNetworktemplateModel, diag.Diagnostics) {
	var state SiteNetworktemplateModel
	var diags diag.Diagnostics

	state.SiteId = types.StringValue(data.GetSiteId())

	state.AclPolicies = aclPoliciesSdkToTerraform(ctx, &diags, data.GetAclPolicies())

	state.AclTags = aclTagsSdkToTerraform(ctx, &diags, data.GetAclTags())

	state.AdditionalConfigCmds = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAdditionalConfigCmds())

	if data.HasDhcpSnooping() {
		state.DhcpSnooping = dhcpSnoopingSdkToTerraform(ctx, &diags, data.GetDhcpSnooping())
	}

	state.DnsServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsServers())

	state.DnsSuffix = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsSuffix())

	state.ExtraRoutes = extraRoutesSdkToTerraform(ctx, &diags, data.GetExtraRoutes())

	state.MistNac = mistNacSdkToTerraform(ctx, &diags, data.GetMistNac())

	state.NtpServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetNtpServers())

	state.Networks = NetworksSdkToTerraform(ctx, &diags, data.GetNetworks())

	state.PortMirrorings = portMirroringSdkToTerraform(ctx, &diags, data.GetPortMirrorings())

	state.PortUsages = portUsagesSdkToTerraform(ctx, &diags, data.GetPortUsages())

	state.RemoteSyslog = remoteSyslogSdkToTerraform(ctx, &diags, data.GetRemoteSyslog())

	state.RadiusConfig = radiusConfigSdkToTerraform(ctx, &diags, data.GetRadiusConfig())

	state.SwitchMatching = switchMatchingSdkToTerraform(ctx, &diags, data.GetSwitchMatching())

	state.SwitchMgmt = switchMgmtSdkToTerraform(ctx, &diags, data.GetSwitchMgmt())

	state.VrfConfig = vrfConfigSdkToTerraform(ctx, &diags, data.GetVrfConfig())

	state.VrfInstances = vrfInstancesSdkToTerraform(ctx, &diags, data.GetVrfInstances())
	return state, diags
}
