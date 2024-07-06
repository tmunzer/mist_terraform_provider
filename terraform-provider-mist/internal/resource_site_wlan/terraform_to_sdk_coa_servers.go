package resource_site_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func coaServerTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []models.CoaServer {

	var data_list []models.CoaServer
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(CoaServersValue)
		data := models.NewCoaServer(v_plan.Ip.ValueString(), v_plan.Secret.ValueString())
		data.SetDisableEventTimestampCheck(v_plan.DisableEventTimestampCheck.ValueBool())
		data.SetEnabled(v_plan.Enabled.ValueBool())
		data.SetPort(int32(v_plan.Port.ValueInt64()))

		data_list = append(data_list, *data)
	}

	return data_list
}
