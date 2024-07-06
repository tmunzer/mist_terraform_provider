package resource_site_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bandsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []models.Dot11Band {

	var data_list []models.Dot11Band
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		data, e := models.NewDot11BandFromValue(v_plan.ValueString())
		if e != nil {
			diags.AddError("bandsTerraformToSdk", e.Error())
		} else {
			data_list = append(data_list, *data)
		}
	}

	return data_list
}
