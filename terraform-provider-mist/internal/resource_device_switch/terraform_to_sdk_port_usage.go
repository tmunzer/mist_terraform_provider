package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func portUsageScTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SwitchPortUsageStormControl {
	data := models.NewSwitchPortUsageStormControl()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		// var sc_attr_interface interface{} = d
		// sc_attr_value := sc_attr_interface.(StormControlValue)
		sc_attr_value := NewStormControlValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetNoMulticast(sc_attr_value.NoMulticast.ValueBool())
		data.SetNoRegisteredMulticast(sc_attr_value.NoRegisteredMulticast.ValueBool())
		data.SetNoUnknownUnicast(sc_attr_value.NoUnknownUnicast.ValueBool())
		data.SetPercentage(int32(sc_attr_value.Percentage.ValueInt64()))
	}
	return *data
}
func portUsageRulesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SwitchPortUsageDynamicRule {

	var data []models.SwitchPortUsageDynamicRule
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_data := v_interface.(RulesValue)
		rule_src, _ := models.NewSwitchPortUsageDynamicRuleSrcFromValue(v_data.Src.ValueString())
		rule := models.NewSwitchPortUsageDynamicRule(*rule_src)
		rule.SetEquals(v_data.Equals.ValueString())
		rule.SetEqualsAny(mist_transform.ListOfStringTerraformToSdk(ctx, v_data.EqualsAny))
		rule.SetExpression(v_data.Expression.ValueString())
		rule.SetUsage(v_data.Usage.ValueString())
		data = append(data, *rule)
	}
	return data
}
func portUsageTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.SwitchPortUsage {
	data := make(map[string]models.SwitchPortUsage)
	for pu_name, pu_attr := range d.Elements() {
		var pu_attr_interface interface{} = pu_attr
		pu_attr_value := pu_attr_interface.(PortUsagesValue)
		new_pu := models.NewSwitchPortUsage()
		new_pu.SetAllNetworks(pu_attr_value.AllNetworks.ValueBool())
		new_pu.SetAllowDhcpd(pu_attr_value.AllowDhcpd.ValueBool())
		new_pu.SetAllowMultipleSupplicants(pu_attr_value.AllowMultipleSupplicants.ValueBool())
		new_pu.SetBypassAuthWhenServerDown(pu_attr_value.BypassAuthWhenServerDown.ValueBool())
		new_pu.SetBypassAuthWhenServerDownForUnkonwnClient(pu_attr_value.BypassAuthWhenServerDownForUnkonwnClient.ValueBool())
		new_pu.SetDescription(pu_attr_value.Description.ValueString())
		new_pu.SetDisableAutoneg(pu_attr_value.DisableAutoneg.ValueBool())
		new_pu.SetDisabled(pu_attr_value.Disabled.ValueBool())
		new_pu.SetDuplex(models.SwitchPortUsageDuplex(pu_attr_value.Duplex.ValueString()))
		new_pu.SetDynamicVlanNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.DynamicVlanNetworks))
		new_pu.SetEnableMacAuth(pu_attr_value.EnableMacAuth.ValueBool())
		new_pu.SetEnableQos(pu_attr_value.EnableQos.ValueBool())
		new_pu.SetGuestNetwork(pu_attr_value.GuestNetwork.ValueString())
		new_pu.SetInterSwitchLink(pu_attr_value.InterSwitchLink.ValueBool())
		new_pu.SetMacAuthOnly(pu_attr_value.MacAuthOnly.ValueBool())
		new_pu.SetMacAuthProtocol(models.SwitchPortUsageMacAuthProtocol(pu_attr_value.MacAuthProtocol.ValueString()))
		new_pu.SetMacLimit(int32(pu_attr_value.MacLimit.ValueInt64()))
		new_pu.SetMode(models.SwitchPortUsageMode(pu_attr_value.Mode.ValueString()))
		new_pu.SetMtu(int32(pu_attr_value.Mtu.ValueInt64()))
		new_pu.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.Networks))
		new_pu.SetPersistMac(pu_attr_value.PersistMac.ValueBool())
		new_pu.SetPoeDisabled(pu_attr_value.PoeDisabled.ValueBool())
		new_pu.SetPortAuth(pu_attr_value.PortAuth.ValueString())
		new_pu.SetPortNetwork(pu_attr_value.PortNetwork.ValueString())
		new_pu.SetReauthInterval(int32(pu_attr_value.ReauthInterval.ValueInt64()))
		new_pu.SetRejectedNetwork(pu_attr_value.RejectedNetwork.ValueString())
		rules := portUsageRulesTerraformToSdk(ctx, diags, pu_attr_value.Rules)
		new_pu.SetRules(rules)
		new_pu.SetResetDefaultWhen(models.SwitchPortUsageDynamicResetDefaultWhen(pu_attr_value.ResetDefaultWhen.ValueString()))
		new_pu.SetSpeed(pu_attr_value.Speed.ValueString())
		storm_control := portUsageScTerraformToSdk(ctx, diags, pu_attr_value.StormControl)
		new_pu.SetStormControl(storm_control)
		new_pu.SetStpEdge(pu_attr_value.StpEdge.ValueBool())
		new_pu.SetVoipNetwork(pu_attr_value.VoipNetwork.ValueString())

		data[pu_name] = *new_pu
	}
	return data
}