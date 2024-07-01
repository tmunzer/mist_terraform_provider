package resource_nacrule

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func matchingPortTypesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.NacRuleMatchingPortType) basetypes.ListValue {
	list_attr_types := types.StringType
	var list_attr_values []attr.Value
	for _, v := range d {
		v_string := types.StringValue(string(v))
		list_attr_values = append(list_attr_values, v_string)
	}

	r, e := types.ListValueFrom(ctx, list_attr_types, list_attr_values)
	diags.Append(e...)
	return r
}

func matchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.NacRuleMatching) MatchingValue {

	data_map_attr_type := MatchingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"auth_type":     types.StringValue(string(d.GetAuthType())),
		"nactags":       mist_transform.ListOfStringSdkToTerraform(ctx, d.GetNactags()),
		"port_types":    matchingPortTypesSdkToTerraform(ctx, diags, d.GetPortTypes()),
		"site_ids":      mist_transform.ListOfStringSdkToTerraform(ctx, d.GetSiteIds()),
		"sitegroup_ids": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetSitegroupIds()),
		"vendor":        mist_transform.ListOfStringSdkToTerraform(ctx, d.GetVendor()),
	}
	data, e := NewMatchingValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func notMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.NacRuleMatching) NotMatchingValue {

	data_map_attr_type := NotMatchingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"auth_type":     types.StringValue(string(d.GetAuthType())),
		"nactags":       mist_transform.ListOfStringSdkToTerraform(ctx, d.GetNactags()),
		"port_types":    matchingPortTypesSdkToTerraform(ctx, diags, d.GetPortTypes()),
		"site_ids":      mist_transform.ListOfStringSdkToTerraform(ctx, d.GetSiteIds()),
		"sitegroup_ids": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetSitegroupIds()),
		"vendor":        mist_transform.ListOfStringSdkToTerraform(ctx, d.GetVendor()),
	}
	data, e := NewNotMatchingValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
