package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mistapi/internal/provider/utils/transform"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.NetworkTemplate) (NetworktemplateModel, diag.Diagnostics) {
	var state NetworktemplateModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())

	state.AclPolicies = aclPoliciesSdkToTerraform(ctx, &diags, data.GetAclPolicies())

	state.AdditionalConfigCmds = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAdditionalConfigCmds())

	state.DhcpSnooping = dhcpSnoopingSdkToTerraform(ctx, &diags, data.GetDhcpSnooping())

	state.DnsServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsServers())

	state.DnsSuffix = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsSuffix())

	state.MistNac = mistNacSdkToTerraform(ctx, &diags, data.GetMistNac())

	state.NtpServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetNtpServers())

	state.Networks = networksSdkToTerraform(ctx, &diags, data.GetNetworks())

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
