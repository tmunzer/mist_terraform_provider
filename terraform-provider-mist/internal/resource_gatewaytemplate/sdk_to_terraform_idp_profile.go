package resource_gatewaytemplate

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"
)

func idpProfileOverwriteMatchingSeveritiesSdkToTerraform(ctx context.Context, data []mistsdkgo.IdpProfileMatchingSeverityValue) basetypes.ListValue {
	tflog.Debug(ctx, "idpProfileOverwriteMatchingSeveritiesSdkToTerraform")
	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	for _, item := range data {
		value := strings.ReplaceAll(string(item), "\"", "")
		items = append(items, types.StringValue(value))
	}
	list, _ := types.ListValue(items_type, items)
	return list
}

func idpProfileOverwriteMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.IdpProfileMatching) basetypes.MapValue {
	tflog.Debug(ctx, "idpProfileOverwriteMatchingSdkToTerraform")
	data_map_type := IpdProfileOverwriteMatchingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"attack_name": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetAttackName()),
		"dst_subnet":  mist_transform.ListOfStringSdkToTerraform(ctx, d.GetDstSubnet()),
		"severity":    idpProfileOverwriteMatchingSeveritiesSdkToTerraform(ctx, d.GetSeverity()),
	}

	data := NewIpdProfileOverwriteMatchingValueMust(data_map_type, data_map_value)
	r, e := types.MapValueFrom(ctx, IpdProfileOverwriteMatchingValue{}.Type(ctx), data)
	diags.Append(e...)
	return r
}

func idpProfileOverwritesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.IdpProfileOverwrite) basetypes.ListValue {
	tflog.Debug(ctx, "idpProfileOverwritesSdkToTerraform")
	var data_list = []OverwritesValue{}

	for _, v := range d {
		data_map_attr_type := OverwritesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"action":   types.StringValue(string(v.GetAction())),
			"matching": idpProfileOverwriteMatchingSdkToTerraform(ctx, diags, v.GetMatching()),
		}

		data, e := NewOverwritesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := OverwritesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func idpProfileSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.IdpProfile) basetypes.MapValue {
	tflog.Debug(ctx, "idpProfileSdkToTerraform")
	port_usage_type := IdpProfilesValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		var port_usage_state = map[string]attr.Value{
			"base_profile": types.StringValue(string(v.GetBaseProfile())),
			"name":         types.StringValue(v.GetName()),
			"overwrites":   idpProfileOverwritesSdkToTerraform(ctx, diags, v.GetOverwrites()),
		}
		port_usage_object, e := NewIdpProfilesValue(port_usage_type, port_usage_state)
		diags.Append(e...)
		state_value_map[k] = port_usage_object
	}
	state_type := IdpProfilesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}