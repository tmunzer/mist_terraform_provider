package resource_site_setting

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func synthteticTestVlansSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.SyntheticTestProperties) basetypes.ListValue {
	tflog.Debug(ctx, "synthteticTestVlansSdkToTerraform")
	var data_list = []VlansValue{}
	for _, v := range d {
		data_map_attr_type := VlansValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"custom_test_urls": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetCustomTestUrls()),
			"disabled":         types.BoolValue(v.GetDisabled()),
			"vlan_ids":         mist_transform.ListOfIntSdkToTerraform(ctx, v.GetVlanIds()),
		}
		data, e := NewVlansValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := VlansValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func synthteticTestSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SyntheticTestConfig) SyntheticTestValue {
	tflog.Debug(ctx, "synthteticTestSdkToTerraform")

	r_attr_type := SyntheticTestValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"disabled": types.BoolValue(d.GetDisabled()),
		"vlans":    synthteticTestVlansSdkToTerraform(ctx, diags, d.GetVlans()),
	}
	r, e := NewSyntheticTestValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
