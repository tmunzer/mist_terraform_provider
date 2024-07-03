package resource_org_nactag

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgNactagModel) (mistapigo.NacTag, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistapigo.NewNacTag(plan.Name.ValueString(), mistapigo.NacTagType(plan.Type.ValueString()))
	data.SetId(plan.Id.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())

	data.SetAllowUsermacOverride(plan.AllowUsermacOverride.ValueBool())
	data.SetEgressVlanNames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.EgressVlanNames))
	data.SetGbpTag(int32(plan.GbpTag.ValueInt64()))
	data.SetMatch(mistapigo.NacTagMatch(plan.Match.ValueString()))
	data.SetMatchAll(plan.MatchAll.ValueBool())
	data.SetRadiusAttrs(mist_transform.ListOfStringTerraformToSdk(ctx, plan.RadiusAttrs))
	data.SetRadiusGroup(plan.RadiusGroup.ValueString())
	data.SetRadiusVendorAttrs(mist_transform.ListOfStringTerraformToSdk(ctx, plan.RadiusVendorAttrs))
	data.SetSessionTimeout(int32(plan.SessionTimeout.ValueInt64()))
	data.SetValues(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Values))
	data.SetVlan(plan.Vlan.ValueString())

	return data, diags
}
