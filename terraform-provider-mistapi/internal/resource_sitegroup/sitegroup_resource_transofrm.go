package resource_sitegroup

import (
	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(data *mistsdkgo.Sitegroup) (SitegroupModel, diag.Diagnostics) {
	var state SitegroupModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())
	return state, diags
}

func TerraformToSdk(plan *SitegroupModel) (mistsdkgo.Sitegroup, string, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewSitegroup(plan.Name.ValueString())
	var orgId = plan.OrgId.ValueString()
	return data, orgId, diags
}
