package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func gatewayPortVpnPathTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayPortVpnPath {
	tflog.Debug(ctx, "gatewayPortVpnPathTerraformToSdk")
	data_map := make(map[string]models.GatewayPortVpnPath)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VpnPathsValue)
		data := models.GatewayPortVpnPath{}
		data.BfdProfile = models.ToPointer(models.GatewayPortVpnPathBfdProfileEnum(plan.BfdProfile.ValueString()))
		data.BfdUseTunnelMode = models.ToPointer(plan.BfdUseTunnelMode.ValueBool())
		data.Role = models.ToPointer(models.GatewayPortVpnPathRoleEnum(plan.Role.ValueString()))

		traffic_shaping := gatewayPortTrafficShapingTerraformToSdk(ctx, diags, plan.TrafficShaping)
		data.TrafficShaping = traffic_shaping

		data_map[k] = data
	}
	return data_map
}

func gatewayPortTrafficShapingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.GatewayTrafficShaping {
	tflog.Debug(ctx, "gatewayPortTrafficShapingTerraformToSdk")
	data := models.GatewayTrafficShaping{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewTrafficShapingValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.ClassPercentages = mist_transform.ListOfIntTerraformToSdk(ctx, plan.ClassPercentages)
		data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
		return &data
	}
}

func gatewayIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.GatewayIpConfig {
	tflog.Debug(ctx, "gatewayIpConfigTerraformToSdk")
	data := models.GatewayIpConfig{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewIpConfigValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Dns = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Dns)
		data.DnsSuffix = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
		data.Gateway = models.ToPointer(plan.Gateway.ValueString())
		data.Ip = models.ToPointer(plan.Ip.ValueString())
		data.Netmask = models.ToPointer(plan.Netmask.ValueString())
		data.Network = models.ToPointer(plan.Network.ValueString())
		data.PoserPassword = models.ToPointer(plan.PoserPassword.ValueString())
		data.PppoeUsername = models.ToPointer(plan.PppoeUsername.ValueString())
		data.PppoeAuth = models.ToPointer(models.GatewayWanPpoeAuthEnum(plan.PppoeAuth.ValueString()))
		data.Type = models.ToPointer(models.GatewayWanTypeEnum(plan.IpConfigType.ValueString()))
		return &data
	}
}

func portConfigWanSourceNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.GatewayPortWanSourceNat {
	data := models.GatewayPortWanSourceNat{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewWanSourceNatValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Disabled = plan.Disabled.ValueBoolPointer()
		data.NatPool = plan.NatPool.ValueStringPointer()
		return &data
	}
}
func portConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayPortConfig {
	tflog.Debug(ctx, "portConfigTerraformToSdk")
	data_map := make(map[string]models.GatewayPortConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PortConfigValue)

		data := models.GatewayPortConfig{}
		data.Usage = models.GatewayPortUsageEnum(plan.Usage.ValueString())
		data.Description = models.ToPointer(plan.Description.ValueString())
		data.DisableAutoneg = models.ToPointer(plan.DisableAutoneg.ValueBool())
		data.Disabled = models.ToPointer(plan.Disabled.ValueBool())
		data.DslType = models.ToPointer(models.GatewayPortDslTypeEnum(plan.DslType.ValueString()))
		data.DslVci = models.ToPointer(int(plan.DslVci.ValueInt64()))
		data.DslVpi = models.ToPointer(int(plan.DslVpi.ValueInt64()))
		data.Duplex = models.ToPointer(models.GatewayPortDuplexEnum(plan.Duplex.ValueString()))

		t, _ := plan.IpConfig.ToObjectValue(ctx)
		ip_config := gatewayIpConfigTerraformToSdk(ctx, diags, t)
		data.IpConfig = ip_config

		data.LteApn = models.ToPointer(plan.LteApn.ValueString())
		data.LteAuth = models.ToPointer(models.GatewayPortLteAuthEnum(plan.LteAuth.ValueString()))
		data.LteBackup = models.ToPointer(plan.LteBackup.ValueBool())
		data.LtePassword = models.ToPointer(plan.LtePassword.ValueString())
		data.LteUsername = models.ToPointer(plan.LteUsername.ValueString())
		data.Mtu = models.ToPointer(int(plan.Mtu.ValueInt64()))
		data.Name = models.ToPointer(plan.Name.ValueString())
		data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks)
		data.OuterVlanId = models.ToPointer(int(plan.OuterVlanId.ValueInt64()))
		data.PoeDisabled = models.ToPointer(plan.PoeDisabled.ValueBool())
		data.PortNetwork = models.ToPointer(plan.PortNetwork.ValueString())
		data.PreserveDscp = models.ToPointer(plan.PreserveDscp.ValueBool())
		data.Redundant = models.ToPointer(plan.Redundant.ValueBool())
		data.RethIdx = models.ToPointer(int(plan.RethIdx.ValueInt64()))
		data.RethNode = models.ToPointer(plan.RethNode.ValueString())
		data.Speed = models.ToPointer(plan.Speed.ValueString())
		data.SsrNoVirtualMac = models.ToPointer(plan.SsrNoVirtualMac.ValueBool())
		data.SvrPortRange = models.ToPointer(plan.SvrPortRange.ValueString())

		traffic_shaping := gatewayPortTrafficShapingTerraformToSdk(ctx, diags, plan.TrafficShaping)
		data.TrafficShaping = traffic_shaping

		data.VlanId = models.ToPointer(int(plan.VlanId.ValueInt64()))

		vpn_paths := gatewayPortVpnPathTerraformToSdk(ctx, diags, plan.VpnPaths)
		data.VpnPaths = vpn_paths

		data.WanArpPolicer = models.ToPointer(models.GatewayPortWanArpPolicerEnum(plan.WanArpPolicer.ValueString()))
		data.WanExtIp = models.ToPointer(plan.WanExtIp.ValueString())
		data.WanSourceNat = portConfigWanSourceNatTerraformToSdk(ctx, diags, plan.WanSourceNat)
		data.WanType = models.ToPointer(models.GatewayPortWanTypeEnum(plan.WanType.ValueString()))

		data_map[k] = data
	}
	return data_map
}
