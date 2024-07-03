package resource_org_wlan

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bandsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []mistapigo.Dot11Band {

	var data_list []mistapigo.Dot11Band
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		data, e := mistapigo.NewDot11BandFromValue(v_plan.ValueString())
		if e != nil {
			diags.AddError("bandsTerraformToSdk", e.Error())
		} else {
			data_list = append(data_list, *data)
		}
	}

	return data_list
}
