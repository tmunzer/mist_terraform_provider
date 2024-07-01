package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bandsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []mistsdkgo.Dot11Band {

	var data_list []mistsdkgo.Dot11Band
	for _, v := range plan.Elements() {
		v_plan, _ := mistsdkgo.NewDot11BandFromValue(v.String())
		data_list = append(data_list, *v_plan)
	}

	return data_list
}
