package resource_site_wlan

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func coaServerTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []mistapigo.CoaServer {

	var data_list []mistapigo.CoaServer
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(CoaServersValue)
		data := mistapigo.NewCoaServer(v_plan.Ip.ValueString(), v_plan.Secret.ValueString())
		data.SetDisableEventTimestampCheck(v_plan.DisableEventTimestampCheck.ValueBool())
		data.SetEnabled(v_plan.Enabled.ValueBool())
		data.SetPort(int32(v_plan.Port.ValueInt64()))

		data_list = append(data_list, *data)
	}

	return data_list
}
