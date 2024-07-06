package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func switchMatchingRulesPortMirroringTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.SwitchPortMirroringProperty {

	data := make(map[string]models.SwitchPortMirroringProperty)
	for k, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(PortMirroringValue)
		item_obj := models.SwitchPortMirroringProperty{}
		item_obj.InputNetworksIngress = mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputNetworksIngress)
		item_obj.InputPortIdsEgress = mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputPortIdsEgress)
		item_obj.InputPortIdsIngress = mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputPortIdsIngress)
		item_obj.OutputPortId = models.ToPointer(plan_obj.OutputPortId.ValueString())
		data[k] = item_obj
	}
	return data
}
func switchMatchingRulesPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.JunosPortConfig {

	data := make(map[string]models.JunosPortConfig)
	for k, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(PortConfigValue)
		item_obj := models.JunosPortConfig{}
		item_obj.Usage = plan_obj.Usage.ValueString()
		item_obj.AeDisableLacp = models.ToPointer(plan_obj.AeDisableLacp.ValueBool())
		item_obj.AeIdx = models.ToPointer(int(plan_obj.AeIdx.ValueInt64()))
		item_obj.AeLacpSlow = models.ToPointer(plan_obj.AeLacpSlow.ValueBool())
		item_obj.Aggregated = models.ToPointer(plan_obj.Aggregated.ValueBool())
		item_obj.Critical = models.ToPointer(plan_obj.Critical.ValueBool())
		item_obj.Description = models.ToPointer(plan_obj.Description.ValueString())
		item_obj.DisableAutoneg = models.ToPointer(plan_obj.DisableAutoneg.ValueBool())
		item_obj.Duplex = models.ToPointer(models.JunosPortConfigDuplexEnum(plan_obj.Duplex.ValueString()))
		item_obj.DynamicUsage.SetValue(models.ToPointer(plan_obj.DynamicUsage.ValueString()))
		item_obj.Esilag = models.ToPointer(plan_obj.Esilag.ValueBool())
		item_obj.Mtu = models.ToPointer(int(plan_obj.Mtu.ValueInt64()))
		item_obj.NoLocalOverwrite = models.ToPointer(plan_obj.NoLocalOverwrite.ValueBool())
		item_obj.PoeDisabled = models.ToPointer(plan_obj.PoeDisabled.ValueBool())
		item_obj.Speed = models.ToPointer(models.JunosPortConfigSpeedEnum(plan_obj.Speed.ValueString()))
		data[k] = item_obj
	}
	return data
}
func switchMatchingRulesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SwitchMatchingRule {
	tflog.Debug(ctx, "switchMatchingRulesTerraformToSdk")

	var data []models.SwitchMatchingRule
	for _, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(MatchingRulesValue)
		item_obj := models.SwitchMatchingRule{}

		item_obj.AdditionalConfigCmds = mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.AdditionalConfigCmds)
		item_obj.MatchRole = models.ToPointer(plan_obj.MatchRole.ValueString())
		item_obj.Name = models.ToPointer(plan_obj.Name.ValueString())
		item_obj.PortConfig = switchMatchingRulesPortConfigTerraformToSdk(ctx, diags, plan_obj.PortConfig)
		item_obj.PortMirroring = switchMatchingRulesPortMirroringTerraformToSdk(ctx, diags, plan_obj.PortMirroring)

		match := map[string]interface{}{}
		if !plan_obj.MatchType.IsUnknown() && !plan_obj.MatchType.IsNull() {
			match_type := plan_obj.MatchType.ValueString()
			match_value := plan_obj.MatchValue.ValueString()
			match[match_type] = match_value
		}
		item_obj.AdditionalProperties = match

		data = append(data, item_obj)
	}
	return data
}

func switchMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SwitchMatchingValue) *models.SwitchMatching {

	data := models.SwitchMatching{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		data.Enable = models.ToPointer(d.Enable.ValueBool())
		data.Rules = switchMatchingRulesTerraformToSdk(ctx, diags, d.MatchingRules)

		return &data
	}

}
