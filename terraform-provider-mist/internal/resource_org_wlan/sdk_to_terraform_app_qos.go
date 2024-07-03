package resource_org_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appQosAppsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data map[string]mistsdkgo.WlanAppQosAppsProperties) basetypes.MapValue {

	map_attr_values := make(map[string]attr.Value)
	for k, v := range data {
		var v_item = map[string]attr.Value{
			"dscp":       types.Int64Value(int64(v.GetDscp())),
			"dst_subnet": types.StringValue(v.GetDstSubnet()),
			"src_subnet": types.StringValue(v.GetSrcSubnet()),
		}
		v_obj, e := NewAppsValue(AppsValue{}.AttributeTypes(ctx), v_item)
		diags.Append(e...)
		map_attr_values[k] = v_obj
	}
	r, e := types.MapValueFrom(ctx, AppsValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return r
}

func appQosOthersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistsdkgo.WlanAppQosOthersItem) basetypes.ListValue {

	var data_list = []OthersValue{}
	for _, v := range data {
		data_map_value := map[string]attr.Value{
			"dscp":        types.Int64Value(int64(v.GetDscp())),
			"dst_subnet":  types.StringValue(v.GetDstSubnet()),
			"port_ranges": types.StringValue(v.GetPortRanges()),
			"protocol":    types.StringValue(v.GetProtocol()),
			"src_subnet":  types.StringValue(v.GetSrcSubnet()),
		}
		data, e := NewOthersValue(OthersValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, OthersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func appQosSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistsdkgo.WlanAppQos) AppQosValue {

	plan_attr := map[string]attr.Value{
		"apps":    appQosAppsSdkToTerraform(ctx, diags, data.GetApps()),
		"enabled": types.BoolValue(data.GetEnabled()),
		"others":  appQosOthersSdkToTerraform(ctx, diags, data.GetOthers()),
	}
	r, e := NewAppQosValue(AppQosValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
