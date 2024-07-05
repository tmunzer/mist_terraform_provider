package resource_device_switch

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ospfConfigAreasTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.OspfConfigArea {
	data_map := make(map[string]mistapigo.OspfConfigArea)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(AreasValue)
		data := *mistapigo.NewOspfConfigArea()
		data.SetNoSummary(plan.NoSummary.ValueBool())
		data_map[k] = data
	}
	return data_map
}

func ospfConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OspfConfigValue) mistapigo.OspfConfig {
	tflog.Debug(ctx, "ospfConfigTerraformToSdk")

	data := *mistapigo.NewOspfConfig()

	areas := ospfConfigAreasTerraformToSdk(ctx, diags, d.Areas)

	data.SetAreas(areas)
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetReferenceBandwidth(d.ReferenceBandwidth.ValueString())

	return data
}
