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
		sc_attr_value := sc_attr_interface.(StormControlValue)
		data.NoMulticast = models.ToPointer(sc_attr_value.NoMulticast.ValueBool())
		data.NoRegisteredMulticast = models.ToPointer(sc_attr_value.NoRegisteredMulticast.ValueBool())
		data.NoUnknownUnicast = models.ToPointer(sc_attr_value.NoUnknownUnicast.ValueBool())
		data.Percentage = models.ToPointer(int(sc_attr_value.Percentage.ValueInt64()))
	}
	return data
}
func portUsageRulesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SwitchPortUsageDynamicRule {

	var data []models.SwitchPortUsageDynamicRule
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_data := v_interface.(RulesValue)
		rule := models.SwitchPortUsageDynamicRule{} //(*rule_src)
		rule.Equals = models.ToPointer(v_data.Equals.ValueString())
		rule.EqualsAny = mist_transform.ListOfStringTerraformToSdk(ctx, v_data.EqualsAny)
		rule.Expression = models.ToPointer(v_data.Expression.ValueString())
		rule.Usage = models.ToPointer(v_data.Usage.ValueString())
		rule.Src = models.SwitchPortUsageDynamicRuleSrcEnum(v_data.Src.ValueString())
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
		new_pu.AllNetworks = models.ToPointer(pu_attr_value.AllNetworks.ValueBool())
		new_pu.AllowDhcpd = models.ToPointer(pu_attr_value.AllowDhcpd.ValueBool())
		new_pu.AllowMultipleSupplicants = models.ToPointer(pu_attr_value.AllowMultipleSupplicants.ValueBool())
		new_pu.BypassAuthWhenServerDown = models.ToPointer(pu_attr_value.BypassAuthWhenServerDown.ValueBool())
		new_pu.BypassAuthWhenServerDownForUnkonwnClient = models.ToPointer(pu_attr_value.BypassAuthWhenServerDownForUnkonwnClient.ValueBool())
		new_pu.Description = models.ToPointer(pu_attr_value.Description.ValueString())
		new_pu.DisableAutoneg = models.ToPointer(pu_attr_value.DisableAutoneg.ValueBool())
		new_pu.Disabled = models.ToPointer(pu_attr_value.Disabled.ValueBool())
		new_pu.Duplex = models.ToPointer(models.SwitchPortUsageDuplexEnum(pu_attr_value.Duplex.ValueString()))
		new_pu.DynamicVlanNetworks = mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.DynamicVlanNetworks)
		new_pu.EnableMacAuth = models.ToPointer(pu_attr_value.EnableMacAuth.ValueBool())
		new_pu.EnableQos = models.ToPointer(pu_attr_value.EnableQos.ValueBool())
		new_pu.GuestNetwork.SetValue(models.ToPointer(pu_attr_value.GuestNetwork.ValueString()))
		new_pu.InterSwitchLink = models.ToPointer(pu_attr_value.InterSwitchLink.ValueBool())
		new_pu.MacAuthOnly = models.ToPointer(pu_attr_value.MacAuthOnly.ValueBool())
		new_pu.MacAuthProtocol = models.ToPointer(models.SwitchPortUsageMacAuthProtocolEnum(pu_attr_value.MacAuthProtocol.ValueString()))
		new_pu.MacLimit = models.ToPointer(int(pu_attr_value.MacLimit.ValueInt64()))
		new_pu.Mode = models.ToPointer(models.SwitchPortUsageModeEnum(pu_attr_value.Mode.ValueString()))
		new_pu.Mtu = models.ToPointer(int(pu_attr_value.Mtu.ValueInt64()))
		new_pu.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.Networks)
		new_pu.PersistMac = models.ToPointer(pu_attr_value.PersistMac.ValueBool())
		new_pu.PoeDisabled = models.ToPointer(pu_attr_value.PoeDisabled.ValueBool())
		new_pu.PortAuth = models.ToPointer(pu_attr_value.PortAuth.ValueString())
		new_pu.PortNetwork = models.ToPointer(pu_attr_value.PortNetwork.ValueString())
		new_pu.ReauthInterval = models.ToPointer(int(pu_attr_value.ReauthInterval.ValueInt64()))
		new_pu.RejectedNetwork.SetValue(models.ToPointer(pu_attr_value.RejectedNetwork.ValueString()))
		new_pu.Rules = portUsageRulesTerraformToSdk(ctx, diags, pu_attr_value.Rules)
		new_pu.ResetDefaultWhen = models.ToPointer(models.SwitchPortUsageDynamicResetDefaultWhenEnum(pu_attr_value.ResetDefaultWhen.ValueString()))
		new_pu.Speed = models.ToPointer(pu_attr_value.Speed.ValueString())
		storm_control := portUsageScTerraformToSdk(ctx, diags, pu_attr_value.StormControl)
		new_pu.StormControl = models.ToPointer(storm_control)
		new_pu.StpEdge = models.ToPointer(pu_attr_value.StpEdge.ValueBool())
		new_pu.VoipNetwork = models.ToPointer(pu_attr_value.VoipNetwork.ValueString())

		data[pu_name] = new_pu
	}
	return data
}
