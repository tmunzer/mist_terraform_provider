package resource_org_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func extraRouteTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.GatewayExtraRoute {
	tflog.Debug(ctx, "extraRouteTerraformToSdk")
	data_map := make(map[string]mistsdkgo.GatewayExtraRoute)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ExtraRoutesValue)

		data := *mistsdkgo.NewGatewayExtraRoute()
		data.SetVia(plan.Via.ValueString())

		data_map[k] = data
	}
	return data_map
}
