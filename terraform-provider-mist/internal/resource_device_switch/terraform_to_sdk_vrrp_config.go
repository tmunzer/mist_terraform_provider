package resource_device_switch

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func vrrpGroupsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.VrrpConfigGroup {
	data_map := make(map[string]models.VrrpConfigGroup)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(GroupsValue)
		data := models.NewVrrpConfigGroup()
		data.SetPriority(int32(plan.Priority.ValueInt64()))

		data_map[k] = *data
	}
	return data_map
}

func vrrpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VrrpConfigValue) models.VrrpConfig {
	tflog.Debug(ctx, "vrrpTerraformToSdk")

	data := *models.NewVrrpConfig()

	groups := vrrpGroupsTerraformToSdk(ctx, diags, d.Groups)

	data.SetEnabled(d.Enabled.ValueBool())
	data.SetGroups(groups)

	return data
}
