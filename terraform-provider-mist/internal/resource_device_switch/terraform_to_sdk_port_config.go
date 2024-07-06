package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"mistapi/models"
)

func portConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.JunosPortConfig {

	data := make(map[string]models.JunosPortConfig)
	for k, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(PortConfigValue)
		item_obj := models.NewJunosPortConfig(plan_obj.Usage.ValueString())
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
		item_obj.SetSpeed(models.JunosPortConfigSpeed(plan_obj.Speed.ValueString()))
		data[k] = *item_obj
	}
	return data
}
