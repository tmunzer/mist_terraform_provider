package resource_site_wxtag

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func specsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistsdkgo.WxlanTagSpec) basetypes.ListValue {

	var data_list = []SpecsValue{}
	for _, v := range data {
		data_map_value := map[string]attr.Value{
			"port_range": types.StringValue(v.GetPortRange()),
			"protocol":   types.StringValue(v.GetProtocol()),
			"subnets":    mist_transform.ListOfStringSdkToTerraform(ctx, v.GetSubnets()),
		}
		data, e := NewSpecsValue(SpecsValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, SpecsValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
