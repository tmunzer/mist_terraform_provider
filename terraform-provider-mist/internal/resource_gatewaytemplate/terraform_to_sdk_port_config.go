package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func gatewayPortVpnPathTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.GatewayPortVpnPath {
	tflog.Debug(ctx, "gatewayPortVpnPathTerraformToSdk")
	data_map := make(map[string]mistsdkgo.GatewayPortVpnPath)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VpnPathsValue)
		data := *mistsdkgo.NewGatewayPortVpnPath()
		data.SetBfdProfile(mistsdkgo.GatewayPortVpnPathBfdProfile(plan.BfdProfile.ValueString()))
		data.SetBfdUseTunnelMode(plan.BfdUseTunnelMode.ValueBool())
		data.SetRole(mistsdkgo.GatewayPortVpnPathRole(plan.Role.ValueString()))

		traffic_shaping := gatewayPortTrafficShapingTerraformToSdk(ctx, diags, plan.TrafficShaping)
		data.SetTrafficShaping(traffic_shaping)

		data_map[k] = data
	}
	return data_map
}

func gatewayPortTrafficShapingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.GatewayTrafficShaping {
	tflog.Debug(ctx, "gatewayPortTrafficShapingTerraformToSdk")
	data := *mistsdkgo.NewGatewayTrafficShaping()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewTrafficShapingValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetClassPercentages(mist_transform.ListOfIntTerraformToSdk(ctx, plan.ClassPercentages))
		data.SetEnabled(plan.Enabled.ValueBool())
		return data
	}
}

func gatewayIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.GatewayIpConfig {
	tflog.Debug(ctx, "gatewayIpConfigTerraformToSdk")
	data := mistsdkgo.NewGatewayIpConfig()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		plan := NewIpConfigValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetDns(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Dns))
		data.SetDnsSuffix(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix))
		data.SetGateway(plan.Gateway.ValueString())
		data.SetIp(plan.Ip.ValueString())
		data.SetNetmask(plan.Netmask.ValueString())
		data.SetNetwork(plan.Network.ValueString())
		data.SetPoserPassword(plan.PoserPassword.ValueString())
		data.SetPppoeUsername(plan.PppoeUsername.ValueString())
		data.SetPppoeAuth(mistsdkgo.GatewayWanPpoeAuth(plan.PppoeAuth.ValueString()))
		data.SetType(mistsdkgo.GatewayWanType(plan.IpConfigType.ValueString()))
		return *data
	}
}

func portConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.GatewayPortConfig {
	tflog.Debug(ctx, "portConfigTerraformToSdk")
	data_map := make(map[string]mistsdkgo.GatewayPortConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PortConfigValue)

		data := *mistsdkgo.NewGatewayPortConfig(mistsdkgo.GatewayPortUsage(plan.Usage.ValueString()))
		data.SetDescription(plan.Description.ValueString())
		data.SetDisableAutoneg(plan.DisableAutoneg.ValueBool())
		data.SetDisabled(plan.Disabled.ValueBool())
		data.SetDslType(mistsdkgo.GatewayPortDslType(plan.DslType.ValueString()))
		data.SetDslVci(int32(plan.DslVci.ValueInt64()))
		data.SetDslVpi(int32(plan.DslVpi.ValueInt64()))
		data.SetDuplex(mistsdkgo.GatewayPortDuplex(plan.Duplex.ValueString()))

		t, _ := plan.IpConfig.ToObjectValue(ctx)
		ip_config := gatewayIpConfigTerraformToSdk(ctx, diags, t)
		data.SetIpConfig(ip_config)

		data.SetLteApn(plan.LteApn.ValueString())
		data.SetLteAuth(mistsdkgo.GatewayPortLteAuth(plan.LteAuth.ValueString()))
		data.SetLteBackup(plan.LteBackup.ValueBool())
		data.SetLtePassword(plan.LtePassword.ValueString())
		data.SetLteUsername(plan.LteUsername.ValueString())
		data.SetMtu(int32(plan.Mtu.ValueInt64()))
		data.SetName(plan.Name.ValueString())
		data.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks))
		data.SetOuterVlanId(int32(plan.OuterVlanId.ValueInt64()))
		data.SetPoeDisabled(plan.PoeDisabled.ValueBool())
		data.SetPortNetwork(plan.PortNetwork.ValueString())
		data.SetPreserveDscp(plan.PreserveDscp.ValueBool())
		data.SetRedundant(plan.Redundant.ValueBool())
		data.SetRethIdx(int32(plan.RethIdx.ValueInt64()))
		data.SetRethNode(plan.RethNode.ValueString())
		data.SetSpeed(plan.Speed.ValueString())
		data.SetSsrNoVirtualMac(plan.SsrNoVirtualMac.ValueBool())
		data.SetSvrPortRange(plan.SvrPortRange.ValueString())

		traffic_shaping := gatewayPortTrafficShapingTerraformToSdk(ctx, diags, plan.TrafficShaping)
		data.SetTrafficShaping(traffic_shaping)

		data.SetVlanId(int32(plan.VlanId.ValueInt64()))

		vpn_paths := gatewayPortVpnPathTerraformToSdk(ctx, diags, plan.VpnPaths)
		data.SetVpnPaths(vpn_paths)

		data.SetWanArpPolicer(mistsdkgo.GatewayPortWanArpPolicer(plan.WanArpPolicer.ValueString()))
		data.SetWanExtIp(plan.WanExtIp.ValueString())
		data.SetWanType(mistsdkgo.GatewayPortWanType(plan.WanType.ValueString()))

		data_map[k] = data
	}
	return data_map
}
