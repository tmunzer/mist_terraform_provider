package resource_device_switch

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func vrrpGroupsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.VrrpConfigGroup {
	data_map := make(map[string]mistapigo.VrrpConfigGroup)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(GroupsValue)
		data := mistapigo.NewVrrpConfigGroup()
		data.SetPriority(int32(plan.Priority.ValueInt64()))

		data_map[k] = *data
	}
	return data_map
}

func vrrpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VrrpConfigValue) mistapigo.VrrpConfig {
	tflog.Debug(ctx, "vrrpTerraformToSdk")

	data := *mistapigo.NewVrrpConfig()

	groups := vrrpGroupsTerraformToSdk(ctx, diags, d.Groups)

	data.SetEnabled(d.Enabled.ValueBool())
	data.SetGroups(groups)

	return data
}
