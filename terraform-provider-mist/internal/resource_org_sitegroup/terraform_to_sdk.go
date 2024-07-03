package resource_org_sitegroup

import (
	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgSitegroupModel) (mistapigo.Sitegroup, string, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistapigo.NewSitegroup(plan.Name.ValueString())
	var orgId = plan.OrgId.ValueString()
	return data, orgId, diags
}
