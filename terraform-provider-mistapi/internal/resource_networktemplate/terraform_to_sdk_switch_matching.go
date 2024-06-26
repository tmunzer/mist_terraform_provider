package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mistapi/internal/provider/utils/transform"
)

func switchMatchingRulesPortMirroringTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.SwitchPortMirroring {

	data := make(map[string]mistsdkgo.SwitchPortMirroring)
	for k, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(PortMirroringValue)
		item_obj := mistsdkgo.NewSwitchPortMirroring()
		item_obj.SetInputNetworksIngress(mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputNetworksIngress))
		item_obj.SetInputPortIdsEgress(mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputPortIdsEgress))
		item_obj.SetInputPortIdsIngress(mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputPortIdsIngress))
		item_obj.SetOutputPortId(plan_obj.OutputPortId.ValueString())
		data[k] = *item_obj
	}
	return data
}
func switchMatchingRulesPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.JunosPortConfig {

	data := make(map[string]mistsdkgo.JunosPortConfig)
	for k, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(PortConfigValue)
		item_obj := mistsdkgo.NewJunosPortConfig(plan_obj.Usage.ValueString())
		item_obj.SetAeDisableLacp(plan_obj.AeDisableLacp.ValueBool())
		item_obj.SetAeIdx(int32(plan_obj.AeIdx.ValueInt64()))
		item_obj.SetAeLacpSlow(plan_obj.AeLacpSlow.ValueBool())
		item_obj.SetAggregated(plan_obj.Aggregated.ValueBool())
		item_obj.SetCritical(plan_obj.Critical.ValueBool())
		item_obj.SetDescription(plan_obj.Description.ValueString())
		item_obj.SetDisableAutoneg(plan_obj.DisableAutoneg.ValueBool())
		item_obj.SetDynamicUsage(plan_obj.DynamicUsage.ValueString())
		item_obj.SetEsilag(plan_obj.Esilag.ValueBool())
		item_obj.SetMtu(int32(plan_obj.Mtu.ValueInt64()))
		item_obj.SetNoLocalOverwrite(plan_obj.NoLocalOverwrite.ValueBool())
		item_obj.SetPoeDisabled(plan_obj.PoeDisabled.ValueBool())
		item_obj.SetSpeed(mistsdkgo.JunosPortConfigSpeed(plan_obj.Speed.ValueString()))
		data[k] = *item_obj
	}
	return data
}
func switchMatchingRulesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.SwitchMatchingRule {

	var data []mistsdkgo.SwitchMatchingRule
	for _, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(RulesValue)
		item_obj := mistsdkgo.NewSwitchMatchingRule()
		switch_matching_rule_port_config := switchMatchingRulesPortConfigTerraformToSdk(ctx, diags, plan_obj.PortConfig)
		switch_matching_rule_port_mirroring := switchMatchingRulesPortMirroringTerraformToSdk(ctx, diags, plan_obj.PortMirroring)

		item_obj.SetAdditionalConfigCmds(mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.AdditionalConfigCmds))
		item_obj.SetMatchRole(plan_obj.MatchRole.ValueString())
		item_obj.SetName(plan_obj.Name.ValueString())
		item_obj.SetPortConfig(switch_matching_rule_port_config)
		item_obj.SetPortMirroring(switch_matching_rule_port_mirroring)

		data = append(data, *item_obj)
	}
	return data
}

func switchMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SwitchMatchingValue) mistsdkgo.SwitchMatching {

	data := mistsdkgo.NewSwitchMatching()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		switch_matching_rules := switchMatchingRulesTerraformToSdk(ctx, diags, d.Rules)

		data.SetEnable(d.Enable.ValueBool())
		data.SetRules(switch_matching_rules)

		return *data
	}

}
