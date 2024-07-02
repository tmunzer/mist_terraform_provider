package resource_wxtag

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.WxlanTag) (WxtagModel, diag.Diagnostics) {
	var state WxtagModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.SiteId = types.StringValue(data.GetSiteId())

	state.Name = types.StringValue(data.GetName())
	state.LastIps = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetLastIps())
	state.Mac = types.StringValue(data.GetMac())
	state.Match = types.StringValue(string(data.GetMatch()))
	state.Op = types.StringValue(string(data.GetOp()))
	state.ResourceMac = types.StringValue(data.GetResourceMac())
	state.Services = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetServices())

	specs := specsSdkToTerraform(ctx, &diags, data.GetSpecs())
	state.Specs = specs

	state.Subnet = types.StringValue(data.GetSubnet())
	state.Type = types.StringValue(string(data.GetType()))
	state.Values = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetValues())
	state.VlanId = types.Int64Value(int64(data.GetVlanId()))

	return state, diags
}
