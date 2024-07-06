package resource_org_wlan

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func hotspot20SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data models.WlanHotspot20) Hotspot20Value {

	var operators_list []attr.Value
	for _, v := range data.GetOperators() {
		operators_list = append(operators_list, types.StringValue(string(v)))
	}
	operators, e := types.ListValue(basetypes.StringType{}, operators_list)
	diags.Append(e...)

	plan_attr := map[string]attr.Value{
		"domain_name": mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDomainName()),
		"enabled":     types.BoolValue(data.GetEnabled()),
		"nai_realms":  mist_transform.ListOfStringSdkToTerraform(ctx, data.GetNaiRealms()),
		"operators":   operators,
		"rcoi":        mist_transform.ListOfStringSdkToTerraform(ctx, data.GetRcoi()),
		"venue_name":  types.StringValue(data.GetVenueName()),
	}
	r, e := NewHotspot20Value(Hotspot20Value{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
