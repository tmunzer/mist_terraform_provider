package resource_nactag

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *NactagModel) (mistsdkgo.NacTag, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewNacTag(plan.Name.ValueString(), mistsdkgo.NacTagType(plan.Type.ValueString()))
	data.SetId(plan.Id.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())

	data.SetAllowUsermacOverride(plan.AllowUsermacOverride.ValueBool())
	data.SetEgressVlanNames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.EgressVlanNames))
	data.SetGbpTag(int32(plan.GbpTag.ValueInt64()))
	data.SetMatch(mistsdkgo.NacTagMatch(plan.Match.ValueString()))
	data.SetMatchAll(plan.MatchAll.ValueBool())
	data.SetRadiusAttrs(mist_transform.ListOfStringTerraformToSdk(ctx, plan.RadiusAttrs))
	data.SetRadiusGroup(plan.RadiusGroup.ValueString())
	data.SetRadiusVendorAttrs(mist_transform.ListOfStringTerraformToSdk(ctx, plan.RadiusVendorAttrs))
	data.SetSessionTimeout(int32(plan.SessionTimeout.ValueInt64()))
	data.SetValues(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Values))
	data.SetVlan(plan.Vlan.ValueString())

	return data, diags
}
