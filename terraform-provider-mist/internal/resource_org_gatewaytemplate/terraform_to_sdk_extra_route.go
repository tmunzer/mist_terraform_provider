package resource_org_gatewaytemplate

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func extraRouteTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.GatewayExtraRoute {
	tflog.Debug(ctx, "extraRouteTerraformToSdk")
	data_map := make(map[string]mistapigo.GatewayExtraRoute)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ExtraRoutesValue)

		data := *mistapigo.NewGatewayExtraRoute()
		data.SetVia(plan.Via.ValueString())

		data_map[k] = data
	}
	return data_map
}
