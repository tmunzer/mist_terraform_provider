package resource_org

import (
	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgModel) (mistapigo.Org, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistapigo.NewOrg(plan.Name.ValueString())
	return data, diags
}
