package resource_org_wlantemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgWlantemplateModel) (mistapigo.Template, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistapigo.NewTemplate(plan.Name.ValueString())
	data.SetId(plan.Id.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())

	applies := appliesTerraformToSdk(ctx, &diags, plan.Applies)
	data.SetApplies(applies)

	data.SetDeviceprofileIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DeviceprofileIds))

	exceptions := exceptionsTerraformToSdk(ctx, &diags, plan.Exceptions)
	data.SetExceptions(exceptions)

	data.SetFilterByDeviceprofile(plan.FilterByDeviceprofile.ValueBool())

	return data, diags
}
