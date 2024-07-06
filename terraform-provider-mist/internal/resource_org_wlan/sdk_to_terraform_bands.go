package resource_org_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bandsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.Dot11Band) basetypes.ListValue {

	var data_list []attr.Value
	for _, v := range data {
		data_list = append(data_list, types.StringValue(string(v)))
	}
	r, e := types.ListValueFrom(ctx, basetypes.StringType{}, data_list)
	diags.Append(e...)

	return r

}
