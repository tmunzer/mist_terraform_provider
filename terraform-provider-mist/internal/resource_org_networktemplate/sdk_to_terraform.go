package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func SdkToTerraform(ctx context.Context, data *mistapigo.NetworkTemplate) (OrgNetworktemplateModel, diag.Diagnostics) {
	var state OrgNetworktemplateModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())

	state.AclPolicies = aclPoliciesSdkToTerraform(ctx, &diags, data.GetAclPolicies())

	state.AclTags = aclTagsSdkToTerraform(ctx, &diags, data.GetAclTags())

	state.AdditionalConfigCmds = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAdditionalConfigCmds())

	state.DhcpSnooping = dhcpSnoopingSdkToTerraform(ctx, &diags, data.GetDhcpSnooping())

	state.DnsServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsServers())

	state.DnsSuffix = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsSuffix())

	state.ExtraRoutes = extraRoutesSdkToTerraform(ctx, &diags, data.GetExtraRoutes())

	state.ExtraRoutes6 = extraRoutes6SdkToTerraform(ctx, &diags, data.GetExtraRoutes6())

	state.MistNac = mistNacSdkToTerraform(ctx, &diags, data.GetMistNac())

	state.NtpServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetNtpServers())

	state.Networks = NetworksSdkToTerraform(ctx, &diags, data.GetNetworks())

	state.PortMirroring = portMirroringSdkToTerraform(ctx, &diags, data.GetPortMirroring())

	state.PortUsages = portUsagesSdkToTerraform(ctx, &diags, data.GetPortUsages())

	state.RemoteSyslog = remoteSyslogSdkToTerraform(ctx, &diags, data.GetRemoteSyslog())

	state.RadiusConfig = radiusConfigSdkToTerraform(ctx, &diags, data.GetRadiusConfig())

	state.SwitchMatching = switchMatchingSdkToTerraform(ctx, &diags, data.GetSwitchMatching())

	state.SwitchMgmt = switchMgmtSdkToTerraform(ctx, &diags, data.GetSwitchMgmt())

	state.VrfConfig = vrfConfigSdkToTerraform(ctx, &diags, data.GetVrfConfig())

	state.VrfInstances = vrfInstancesSdkToTerraform(ctx, &diags, data.GetVrfInstances())
	return state, diags
}
