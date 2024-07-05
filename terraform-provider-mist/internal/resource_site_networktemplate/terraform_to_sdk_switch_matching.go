package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func switchMatchingRulesPortMirroringTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.SwitchPortMirroringProperty {

	data := make(map[string]mistapigo.SwitchPortMirroringProperty)
	for k, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(PortMirroringValue)
		item_obj := mistapigo.NewSwitchPortMirroringProperty()
		item_obj.SetInputNetworksIngress(mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputNetworksIngress))
		item_obj.SetInputPortIdsEgress(mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputPortIdsEgress))
		item_obj.SetInputPortIdsIngress(mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputPortIdsIngress))
		item_obj.SetOutputPortId(plan_obj.OutputPortId.ValueString())
		data[k] = *item_obj
	}
	return data
}
func switchMatchingRulesPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.JunosPortConfig {

	data := make(map[string]mistapigo.JunosPortConfig)
	for k, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(PortConfigValue)
		item_obj := mistapigo.NewJunosPortConfig(plan_obj.Usage.ValueString())
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
		item_obj.SetSpeed(mistapigo.JunosPortConfigSpeed(plan_obj.Speed.ValueString()))
		data[k] = *item_obj
	}
	return data
}
func switchMatchingRulesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.SwitchMatchingRule {
	tflog.Debug(ctx, "switchMatchingRulesTerraformToSdk")

	var data []mistapigo.SwitchMatchingRule
	for _, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(MatchingRulesValue)
		item_obj := mistapigo.NewSwitchMatchingRule()
		switch_matching_rule_port_config := switchMatchingRulesPortConfigTerraformToSdk(ctx, diags, plan_obj.PortConfig)
		switch_matching_rule_port_mirroring := switchMatchingRulesPortMirroringTerraformToSdk(ctx, diags, plan_obj.PortMirroring)

		item_obj.SetAdditionalConfigCmds(mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.AdditionalConfigCmds))
		item_obj.SetMatchRole(plan_obj.MatchRole.ValueString())
		item_obj.SetName(plan_obj.Name.ValueString())
		item_obj.SetPortConfig(switch_matching_rule_port_config)
		item_obj.SetPortMirroring(switch_matching_rule_port_mirroring)

		match := map[string]interface{}{}
		if !plan_obj.MatchType.IsUnknown() && !plan_obj.MatchType.IsNull() {
			match_type := plan_obj.MatchType.ValueString()
			match_value := plan_obj.MatchValue.ValueString()
			match[match_type] = match_value
		}
		item_obj.AdditionalProperties = match

		data = append(data, *item_obj)
	}
	return data
}

func switchMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SwitchMatchingValue) mistapigo.SwitchMatching {

	data := mistapigo.NewSwitchMatching()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		switch_matching_rules := switchMatchingRulesTerraformToSdk(ctx, diags, d.MatchingRules)

		data.SetEnable(d.Enable.ValueBool())
		data.SetRules(switch_matching_rules)

		return *data
	}

}
