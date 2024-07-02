package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bonjourServicesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data map[string]mistsdkgo.WlanBonjourServiceProperties) basetypes.MapValue {

	map_attr_values := make(map[string]attr.Value)
	for k, v := range data {
		var v_item = map[string]attr.Value{
			"disable_local": types.BoolValue(v.GetDisableLocal()),
			"radius_groups": mist_transform.ListOfStringSdkToTerraform(ctx, v.RadiusGroups),
			"scope":         types.StringValue(string(v.GetScope())),
		}
		v_obj, e := NewServicesValue(ServicesValue{}.AttributeTypes(ctx), v_item)
		diags.Append(e...)
		map_attr_values[k] = v_obj
	}
	r, e := types.MapValueFrom(ctx, ServicesValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return r
}

func bonjourSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistsdkgo.WlanBonjour) BonjourValue {
	var additional_vlan_ids_list []attr.Value
	for _, v := range data.GetAdditionalVlanIds() {
		additional_vlan_ids_list = append(additional_vlan_ids_list, types.Int64Value(int64(v)))
	}
	additional_vlan_ids, e := types.ListValue(basetypes.Int64Type{}, additional_vlan_ids_list)
	diags.Append(e...)

	plan_attr := map[string]attr.Value{
		"additional_vlan_ids": additional_vlan_ids,
		"enabled":             types.BoolValue(data.GetEnabled()),
		"services":            bonjourServicesSdkToTerraform(ctx, diags, data.GetServices()),
	}
	r, e := NewBonjourValue(BonjourValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
