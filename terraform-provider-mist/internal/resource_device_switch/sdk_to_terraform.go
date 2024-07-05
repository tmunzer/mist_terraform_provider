package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func SdkToTerraform(ctx context.Context, mist_device *mistapigo.MistDevice) (DeviceSwitchModel, diag.Diagnostics) {
	var state DeviceSwitchModel
	var diags diag.Diagnostics

	data := mist_device.DeviceSwitch

	// a, _ := mist_device.MarshalJSON()
	// data := mistapigo.NewDeviceSwitch()
	// data.UnmarshalJSON(a)

	state.Id = types.StringValue(data.GetId())
	state.DeviceprofileId = types.StringValue(data.GetDeviceprofileId())
	state.MapId = types.StringValue(data.GetMapId())
	state.SiteId = types.StringValue(data.GetSiteId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())
	state.Notes = types.StringValue(data.GetNotes())

	state.AclPolicies = aclPoliciesSdkToTerraform(ctx, &diags, data.GetAclPolicies())

	state.AclTags = aclTagsSdkToTerraform(ctx, &diags, data.GetAclTags())

	state.AdditionalConfigCmds = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAdditionalConfigCmds())

	if data.HasDhcpSnooping() {
		state.DhcpSnooping = dhcpSnoopingSdkToTerraform(ctx, &diags, data.GetDhcpSnooping())
	}

	state.DhcpdConfig = dhcpdConfigSdkToTerraform(ctx, &diags, data.GetDhcpdConfig())

	state.DisableAutoConfig = types.BoolValue(data.GetDisableAutoConfig())

	state.DnsServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsServers())

	state.DnsSuffix = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsSuffix())

	state.ExtraRoutes = extraRoutesSdkToTerraform(ctx, &diags, data.GetExtraRoutes())

	state.ExtraRoutes6 = extraRoutes6SdkToTerraform(ctx, &diags, data.GetExtraRoutes6())

	state.Image1Url = types.StringValue(data.GetImage1Url())
	state.Image2Url = types.StringValue(data.GetImage2Url())
	state.Image3Url = types.StringValue(data.GetImage3Url())

	state.IpConfig = ipConfigSdkToTerraform(ctx, &diags, data.GetIpConfig())

	state.Managed = types.BoolValue(data.GetManaged())

	state.MistNac = mistNacSdkToTerraform(ctx, &diags, data.GetMistNac())

	state.Networks = NetworksSdkToTerraform(ctx, &diags, data.GetNetworks())

	state.Notes = types.StringValue(data.GetNotes())

	state.NtpServers = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetNtpServers())

	state.OobIpConfig = oobIpConfigSdkToTerraform(ctx, &diags, data.GetOobIpConfig())

	state.OspfConfig = ospfConfigSdkToTerraform(ctx, &diags, data.GetOspfConfig())

	state.OtherIpConfigs = otherIpConfigSdkToTerraform(ctx, &diags, data.GetOtherIpConfigs())

	state.PortConfig = portConfigSdkToTerraform(ctx, &diags, data.GetPortConfig())

	state.PortMirroring = portMirroringSdkToTerraform(ctx, &diags, data.GetPortMirroring())

	state.PortUsages = portUsagesSdkToTerraform(ctx, &diags, data.GetPortUsages())

	state.RadiusConfig = radiusConfigSdkToTerraform(ctx, &diags, data.GetRadiusConfig())

	state.RemoteSyslog = remoteSyslogSdkToTerraform(ctx, &diags, data.GetRemoteSyslog())

	state.Role = types.StringValue(data.GetRole())

	state.RouterId = types.StringValue(data.GetRouterId())

	state.SnmpConfig = snmpConfigSdkToTerraform(ctx, &diags, data.GetSnmpConfig())

	state.StpConfig = stpConfigSdkToTerraform(ctx, &diags, data.GetStpConfig())

	state.SwitchMgmt = switchMgmtSdkToTerraform(ctx, &diags, data.GetSwitchMgmt())

	state.UseRouterIdAsSourceIp = types.BoolValue(data.GetUseRouterIdAsSourceIp())

	state.Vars = varsSdkToTerraform(ctx, &diags, data.GetVars())

	state.VirtualChassis = virtualChassisSdkToTerraform(ctx, &diags, data.GetVirtualChassis())

	state.VrfConfig = vrfConfigSdkToTerraform(ctx, &diags, data.GetVrfConfig())

	state.VrfInstances = vrfInstancesSdkToTerraform(ctx, &diags, data.GetVrfInstances())

	state.VrrpConfig = vrrpSdkToTerraform(ctx, &diags, data.GetVrrpConfig())

	state.X = types.Float64Value(float64(data.GetX()))
	state.Y = types.Float64Value(float64(data.GetY()))

	return state, diags
}
