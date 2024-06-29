package resource_nactag

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/utils/transform"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.NacTag) (NactagModel, diag.Diagnostics) {
	var state NactagModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())

	state.AllowUsermacOverride = types.BoolValue(data.GetAllowUsermacOverride())
	state.EgressVlanNames = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetEgressVlanNames())
	state.GbpTag = types.Int64Value(int64(data.GetGbpTag()))
	state.Match = types.StringValue(string(data.GetMatch()))
	state.MatchAll = types.BoolValue(data.GetMatchAll())
	state.RadiusAttrs = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetRadiusAttrs())
	state.RadiusGroup = types.StringValue(data.GetRadiusGroup())
	state.RadiusVendorAttrs = mist_transform.ListOfStringSdkToTerraform(ctx, data.RadiusVendorAttrs)
	state.SessionTimeout = types.Int64Value(int64(*data.SessionTimeout))
	state.Type = types.StringValue(string(data.GetType()))
	state.Values = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetValues())
	state.Vlan = types.StringValue(data.GetVlan())

	return state, diags
}
