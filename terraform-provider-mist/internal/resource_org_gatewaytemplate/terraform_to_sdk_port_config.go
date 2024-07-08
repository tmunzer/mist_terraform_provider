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
		if plan.BfdProfile.IsNull() && !plan.BfdProfile.IsUnknown() {
			data.BfdProfile = models.ToPointer(models.GatewayPortVpnPathBfdProfileEnum(plan.BfdProfile.ValueString()))
		}
		if plan.BfdUseTunnelMode.IsNull() && !plan.BfdUseTunnelMode.IsUnknown() {
			data.BfdUseTunnelMode = models.ToPointer(plan.BfdUseTunnelMode.ValueBool())
		}
		if plan.Role.IsNull() && !plan.Role.IsUnknown() {
			data.Role = models.ToPointer(models.GatewayPortVpnPathRoleEnum(plan.Role.ValueString()))
		}

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
		if plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
			data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
		}
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
		if plan.Gateway.IsNull() && !plan.Gateway.IsUnknown() {
			data.Gateway = models.ToPointer(plan.Gateway.ValueString())
		}
		if plan.Ip.IsNull() && !plan.Ip.IsUnknown() {
			data.Ip = models.ToPointer(plan.Ip.ValueString())
		}
		if plan.Netmask.IsNull() && !plan.Netmask.IsUnknown() {
			data.Netmask = models.ToPointer(plan.Netmask.ValueString())
		}
		if plan.Network.IsNull() && !plan.Network.IsUnknown() {
			data.Network = models.ToPointer(plan.Network.ValueString())
		}
		if plan.PoserPassword.IsNull() && !plan.PoserPassword.IsUnknown() {
			data.PoserPassword = models.ToPointer(plan.PoserPassword.ValueString())
		}
		if plan.PppoeUsername.IsNull() && !plan.PppoeUsername.IsUnknown() {
			data.PppoeUsername = models.ToPointer(plan.PppoeUsername.ValueString())
		}
		if plan.PppoeAuth.IsNull() && !plan.PppoeAuth.IsUnknown() {
			data.PppoeAuth = models.ToPointer(models.GatewayWanPpoeAuthEnum(plan.PppoeAuth.ValueString()))
		}
		if plan.IpConfigType.IsNull() && !plan.IpConfigType.IsUnknown() {
			data.Type = models.ToPointer(models.GatewayWanTypeEnum(plan.IpConfigType.ValueString()))
		}
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
		if plan.Description.IsNull() && !plan.Description.IsUnknown() {
			data.Description = models.ToPointer(plan.Description.ValueString())
		}
		if plan.DisableAutoneg.IsNull() && !plan.DisableAutoneg.IsUnknown() {
			data.DisableAutoneg = models.ToPointer(plan.DisableAutoneg.ValueBool())
		}
		if plan.Disabled.IsNull() && !plan.Disabled.IsUnknown() {
			data.Disabled = models.ToPointer(plan.Disabled.ValueBool())
		}
		if plan.DslType.IsNull() && !plan.DslType.IsUnknown() {
			data.DslType = models.ToPointer(models.GatewayPortDslTypeEnum(plan.DslType.ValueString()))
		}
		if plan.DslVci.IsNull() && !plan.DslVci.IsUnknown() {
			data.DslVci = models.ToPointer(int(plan.DslVci.ValueInt64()))
		}
		if plan.DslVpi.IsNull() && !plan.DslVpi.IsUnknown() {
			data.DslVpi = models.ToPointer(int(plan.DslVpi.ValueInt64()))
		}
		if plan.Duplex.IsNull() && !plan.Duplex.IsUnknown() {
			data.Duplex = models.ToPointer(models.GatewayPortDuplexEnum(plan.Duplex.ValueString()))
		}

		t, _ := plan.IpConfig.ToObjectValue(ctx)
		ip_config := gatewayIpConfigTerraformToSdk(ctx, diags, t)
		data.IpConfig = ip_config

		if plan.LteApn.IsNull() && !plan.LteApn.IsUnknown() {
			data.LteApn = models.ToPointer(plan.LteApn.ValueString())
		}
		if plan.LteAuth.IsNull() && !plan.LteAuth.IsUnknown() {
			data.LteAuth = models.ToPointer(models.GatewayPortLteAuthEnum(plan.LteAuth.ValueString()))
		}
		if plan.LteBackup.IsNull() && !plan.LteBackup.IsUnknown() {
			data.LteBackup = models.ToPointer(plan.LteBackup.ValueBool())
		}
		if plan.LtePassword.IsNull() && !plan.LtePassword.IsUnknown() {
			data.LtePassword = models.ToPointer(plan.LtePassword.ValueString())
		}
		if plan.LteUsername.IsNull() && !plan.LteUsername.IsUnknown() {
			data.LteUsername = models.ToPointer(plan.LteUsername.ValueString())
		}
		if plan.Mtu.IsNull() && !plan.Mtu.IsUnknown() {
			data.Mtu = models.ToPointer(int(plan.Mtu.ValueInt64()))
		}
		if plan.Name.IsNull() && !plan.Name.IsUnknown() {
			data.Name = models.ToPointer(plan.Name.ValueString())
		}
		data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks)
		if plan.OuterVlanId.IsNull() && !plan.OuterVlanId.IsUnknown() && plan.OuterVlanId.ValueInt64() != 0 {
			data.OuterVlanId = models.ToPointer(int(plan.OuterVlanId.ValueInt64()))
		}
		if plan.PoeDisabled.IsNull() && !plan.PoeDisabled.IsUnknown() {
			data.PoeDisabled = models.ToPointer(plan.PoeDisabled.ValueBool())
		}
		if plan.PortNetwork.IsNull() && !plan.PortNetwork.IsUnknown() {
			data.PortNetwork = models.ToPointer(plan.PortNetwork.ValueString())
		}
		if plan.PreserveDscp.IsNull() && !plan.PreserveDscp.IsUnknown() {
			data.PreserveDscp = models.ToPointer(plan.PreserveDscp.ValueBool())
		}
		if plan.Redundant.IsNull() && !plan.Redundant.IsUnknown() {
			data.Redundant = models.ToPointer(plan.Redundant.ValueBool())
		}
		if plan.RethIdx.IsNull() && !plan.RethIdx.IsUnknown() {
			data.RethIdx = models.ToPointer(int(plan.RethIdx.ValueInt64()))
		}
		if plan.RethNode.IsNull() && !plan.RethNode.IsUnknown() {
			data.RethNode = models.ToPointer(plan.RethNode.ValueString())
		}
		if plan.Speed.IsNull() && !plan.Speed.IsUnknown() {
			data.Speed = models.ToPointer(plan.Speed.ValueString())
		}
		if plan.SsrNoVirtualMac.IsNull() && !plan.SsrNoVirtualMac.IsUnknown() {
			data.SsrNoVirtualMac = models.ToPointer(plan.SsrNoVirtualMac.ValueBool())
		}
		if plan.SvrPortRange.IsNull() && !plan.SvrPortRange.IsUnknown() {
			data.SvrPortRange = models.ToPointer(plan.SvrPortRange.ValueString())
		}

		traffic_shaping := gatewayPortTrafficShapingTerraformToSdk(ctx, diags, plan.TrafficShaping)
		data.TrafficShaping = traffic_shaping

		if plan.VlanId.IsNull() && !plan.VlanId.IsUnknown() {
			data.VlanId = models.ToPointer(int(plan.VlanId.ValueInt64()))
		}

		vpn_paths := gatewayPortVpnPathTerraformToSdk(ctx, diags, plan.VpnPaths)
		data.VpnPaths = vpn_paths

		if plan.WanArpPolicer.IsNull() && !plan.WanArpPolicer.IsUnknown() {
			data.WanArpPolicer = models.ToPointer(models.GatewayPortWanArpPolicerEnum(plan.WanArpPolicer.ValueString()))
		}
		if plan.WanExtIp.IsNull() && !plan.WanExtIp.IsUnknown() {
			data.WanExtIp = models.ToPointer(plan.WanExtIp.ValueString())
		}
		data.WanSourceNat = portConfigWanSourceNatTerraformToSdk(ctx, diags, plan.WanSourceNat)
		data.WanType = models.ToPointer(models.GatewayPortWanTypeEnum(plan.WanType.ValueString()))

		tflog.Error(ctx, "-------------------in", map[string]interface{}{"in": plan, "out": data})
		data_map[k] = data
	}
	return data_map
}
