package resource_sitegroup

import (
	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *SitegroupModel) (mistsdkgo.Sitegroup, string, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewSitegroup(plan.Name.ValueString())
	var orgId = plan.OrgId.ValueString()
	return data, orgId, diags
}