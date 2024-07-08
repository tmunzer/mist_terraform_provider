package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func portUsageScTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SwitchPortUsageStormControl {
	data := models.SwitchPortUsageStormControl{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var sc_attr_interface interface{} = d
		v_plan := sc_attr_interface.(StormControlValue)
		if !v_plan.NoMulticast.IsNull() && !v_plan.NoMulticast.IsUnknown() {
			data.NoMulticast = models.ToPointer(v_plan.NoMulticast.ValueBool())
		}
		if !v_plan.NoRegisteredMulticast.IsNull() && !v_plan.NoRegisteredMulticast.IsUnknown() {
			data.NoRegisteredMulticast = models.ToPointer(v_plan.NoRegisteredMulticast.ValueBool())
		}
		if !v_plan.NoUnknownUnicast.IsNull() && !v_plan.NoUnknownUnicast.IsUnknown() {
			data.NoUnknownUnicast = models.ToPointer(v_plan.NoUnknownUnicast.ValueBool())
		}
		if !v_plan.Percentage.IsNull() && !v_plan.Percentage.IsUnknown() {
			data.Percentage = models.ToPointer(int(v_plan.Percentage.ValueInt64()))
		}
	}
	return data
}
func portUsageRulesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SwitchPortUsageDynamicRule {

	var data []models.SwitchPortUsageDynamicRule
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(RulesValue)
		rule := models.SwitchPortUsageDynamicRule{}
		if !v_plan.Equals.IsNull() && !v_plan.Equals.IsUnknown() {
			rule.Equals = models.ToPointer(v_plan.Equals.ValueString())
		}
		if !v_plan.EqualsAny.IsNull() && !v_plan.EqualsAny.IsUnknown() {
			rule.EqualsAny = mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.EqualsAny)
		}
		if !v_plan.Expression.IsNull() && !v_plan.Expression.IsUnknown() {
			rule.Expression = models.ToPointer(v_plan.Expression.ValueString())
		}
		if !v_plan.Usage.IsNull() && !v_plan.Usage.IsUnknown() {
			rule.Usage = models.ToPointer(v_plan.Usage.ValueString())
		}
		if !v_plan.Src.IsNull() && !v_plan.Src.IsUnknown() {
			rule.Src = models.SwitchPortUsageDynamicRuleSrcEnum(v_plan.Src.ValueString())
		}
		data = append(data, rule)
	}
	return data
}
func portUsageTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.SwitchPortUsage {
	data := make(map[string]models.SwitchPortUsage)
	for pu_name, pu_attr := range d.Elements() {
		var pu_attr_interface interface{} = pu_attr
		pu_attr_value := pu_attr_interface.(PortUsagesValue)

		new_pu := models.SwitchPortUsage{}
		if !pu_attr_value.AllNetworks.IsNull() && !pu_attr_value.AllNetworks.IsUnknown() {
			new_pu.AllNetworks = models.ToPointer(pu_attr_value.AllNetworks.ValueBool())
		}
		if !pu_attr_value.AllowDhcpd.IsNull() && !pu_attr_value.AllowDhcpd.IsUnknown() {
			new_pu.AllowDhcpd = models.ToPointer(pu_attr_value.AllowDhcpd.ValueBool())
		}
		if !pu_attr_value.AllowMultipleSupplicants.IsNull() && !pu_attr_value.AllowMultipleSupplicants.IsUnknown() {
			new_pu.AllowMultipleSupplicants = models.ToPointer(pu_attr_value.AllowMultipleSupplicants.ValueBool())
		}
		if !pu_attr_value.BypassAuthWhenServerDown.IsNull() && !pu_attr_value.BypassAuthWhenServerDown.IsUnknown() {
			new_pu.BypassAuthWhenServerDown = models.ToPointer(pu_attr_value.BypassAuthWhenServerDown.ValueBool())
		}
		if !pu_attr_value.BypassAuthWhenServerDownForUnkonwnClient.IsNull() && !pu_attr_value.BypassAuthWhenServerDownForUnkonwnClient.IsUnknown() {
			new_pu.BypassAuthWhenServerDownForUnkonwnClient = models.ToPointer(pu_attr_value.BypassAuthWhenServerDownForUnkonwnClient.ValueBool())
		}
		if !pu_attr_value.Description.IsNull() && !pu_attr_value.Description.IsUnknown() {
			new_pu.Description = models.ToPointer(pu_attr_value.Description.ValueString())
		}
		if !pu_attr_value.DisableAutoneg.IsNull() && !pu_attr_value.DisableAutoneg.IsUnknown() {
			new_pu.DisableAutoneg = models.ToPointer(pu_attr_value.DisableAutoneg.ValueBool())
		}
		if !pu_attr_value.Disabled.IsNull() && !pu_attr_value.Disabled.IsUnknown() {
			new_pu.Disabled = models.ToPointer(pu_attr_value.Disabled.ValueBool())
		}
		if !pu_attr_value.Duplex.IsNull() && !pu_attr_value.Duplex.IsUnknown() {
			new_pu.Duplex = models.ToPointer(models.SwitchPortUsageDuplexEnum(pu_attr_value.Duplex.ValueString()))
		}
		if !pu_attr_value.DynamicVlanNetworks.IsNull() && !pu_attr_value.DynamicVlanNetworks.IsUnknown() {
			new_pu.DynamicVlanNetworks = mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.DynamicVlanNetworks)
		}
		if !pu_attr_value.EnableMacAuth.IsNull() && !pu_attr_value.EnableMacAuth.IsUnknown() {
			new_pu.EnableMacAuth = models.ToPointer(pu_attr_value.EnableMacAuth.ValueBool())
		}
		if !pu_attr_value.EnableQos.IsNull() && !pu_attr_value.EnableQos.IsUnknown() {
			new_pu.EnableQos = models.ToPointer(pu_attr_value.EnableQos.ValueBool())
		}
		if !pu_attr_value.GuestNetwork.IsNull() && !pu_attr_value.GuestNetwork.IsUnknown() {
			new_pu.GuestNetwork = models.NewOptional(models.ToPointer(pu_attr_value.GuestNetwork.ValueString()))
		}
		if !pu_attr_value.InterSwitchLink.IsNull() && !pu_attr_value.InterSwitchLink.IsUnknown() {
			new_pu.InterSwitchLink = models.ToPointer(pu_attr_value.InterSwitchLink.ValueBool())
		}
		if !pu_attr_value.MacAuthOnly.IsNull() && !pu_attr_value.MacAuthOnly.IsUnknown() {
			new_pu.MacAuthOnly = models.ToPointer(pu_attr_value.MacAuthOnly.ValueBool())
		}
		if !pu_attr_value.MacAuthProtocol.IsNull() && !pu_attr_value.MacAuthProtocol.IsUnknown() {
			new_pu.MacAuthProtocol = models.ToPointer(models.SwitchPortUsageMacAuthProtocolEnum(pu_attr_value.MacAuthProtocol.ValueString()))
		}
		if !pu_attr_value.MacLimit.IsNull() && !pu_attr_value.MacLimit.IsUnknown() {
			new_pu.MacLimit = models.ToPointer(int(pu_attr_value.MacLimit.ValueInt64()))
		}
		if !pu_attr_value.Mode.IsNull() && !pu_attr_value.Mode.IsUnknown() {
			new_pu.Mode = models.ToPointer(models.SwitchPortUsageModeEnum(pu_attr_value.Mode.ValueString()))
		}
		if !pu_attr_value.Mtu.IsNull() && !pu_attr_value.Mtu.IsUnknown() {
			new_pu.Mtu = models.ToPointer(int(pu_attr_value.Mtu.ValueInt64()))
		}
		if !pu_attr_value.Networks.IsNull() && !pu_attr_value.Networks.IsUnknown() {
			new_pu.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.Networks)
		}
		if !pu_attr_value.PersistMac.IsNull() && !pu_attr_value.PersistMac.IsUnknown() {
			new_pu.PersistMac = models.ToPointer(pu_attr_value.PersistMac.ValueBool())
		}
		if !pu_attr_value.PoeDisabled.IsNull() && !pu_attr_value.PoeDisabled.IsUnknown() {
			new_pu.PoeDisabled = models.ToPointer(pu_attr_value.PoeDisabled.ValueBool())
		}
		new_pu.PortAuth = models.ToPointer(pu_attr_value.PortAuth.ValueString())
		if !pu_attr_value.PortNetwork.IsNull() && !pu_attr_value.PortNetwork.IsUnknown() {
			new_pu.PortNetwork = models.ToPointer(pu_attr_value.PortNetwork.ValueString())
		}
		if !pu_attr_value.ReauthInterval.IsNull() && !pu_attr_value.ReauthInterval.IsUnknown() {
			new_pu.ReauthInterval = models.ToPointer(int(pu_attr_value.ReauthInterval.ValueInt64()))
		}
		if !pu_attr_value.RejectedNetwork.IsNull() && !pu_attr_value.RejectedNetwork.IsUnknown() {
			new_pu.RejectedNetwork = models.NewOptional(models.ToPointer(pu_attr_value.RejectedNetwork.ValueString()))
		}
		if !pu_attr_value.Rules.IsNull() && !pu_attr_value.Rules.IsUnknown() {
			new_pu.Rules = portUsageRulesTerraformToSdk(ctx, diags, pu_attr_value.Rules)
		}
		if !pu_attr_value.ResetDefaultWhen.IsNull() && !pu_attr_value.ResetDefaultWhen.IsUnknown() {
			new_pu.ResetDefaultWhen = models.ToPointer(models.SwitchPortUsageDynamicResetDefaultWhenEnum(pu_attr_value.ResetDefaultWhen.ValueString()))
		}
		if !pu_attr_value.Speed.IsNull() && !pu_attr_value.Speed.IsUnknown() {
			new_pu.Speed = models.ToPointer(pu_attr_value.Speed.ValueString())
		}
		if !pu_attr_value.StormControl.IsNull() && !pu_attr_value.StormControl.IsUnknown() {
			storm_control := portUsageScTerraformToSdk(ctx, diags, pu_attr_value.StormControl)
			new_pu.StormControl = models.ToPointer(storm_control)
		}
		if !pu_attr_value.StpEdge.IsNull() && !pu_attr_value.StpEdge.IsUnknown() {
			new_pu.StpEdge = models.ToPointer(pu_attr_value.StpEdge.ValueBool())
		}
		if !pu_attr_value.VoipNetwork.IsNull() && !pu_attr_value.VoipNetwork.IsUnknown() {
			new_pu.VoipNetwork = models.ToPointer(pu_attr_value.VoipNetwork.ValueString())
		}

		data[pu_name] = new_pu
	}
	return data
}
