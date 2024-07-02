package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicVlanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistsdkgo.WlanDynamicVlan) DynamicVlanValue {

	var local_vlan_ids_list []attr.Value
	for _, v := range data.GetLocalVlanIds() {
		local_vlan_ids_list = append(local_vlan_ids_list, types.Int64Value(int64(*v)))
	}
	local_vlan_ids, e := types.ListValue(basetypes.Int64Type{}, local_vlan_ids_list)
	diags.Append(e...)

	vlans_attr := make(map[string]attr.Value)
	for k, v := range data.GetVlans() {
		vlans_attr[k] = types.StringValue(string(v))
	}
	vlans, e := types.MapValueFrom(ctx, basetypes.StringType{}, vlans_attr)
	diags.Append(e...)

	plan_attr := map[string]attr.Value{
		"default_vlan_id": types.Int64Value(int64(data.GetDefaultVlanId())),
		"enabled":         types.BoolValue(data.GetEnabled()),
		"local_vlan_ids":  local_vlan_ids,
		"type":            types.StringValue(string(data.GetType())),
		"vlans":           vlans,
	}
	r, e := NewDynamicVlanValue(DynamicVlanValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
