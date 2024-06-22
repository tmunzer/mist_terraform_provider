package resource_org

import (
	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgModel) (mistsdkgo.Org, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewOrg(plan.Name.ValueString())
	return data, diags
}
