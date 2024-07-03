package resource_org_sitegroup

import (
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgSitegroupModel) (mistsdkgo.Sitegroup, string, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewSitegroup(plan.Name.ValueString())
	var orgId = plan.OrgId.ValueString()
	return data, orgId, diags
}
