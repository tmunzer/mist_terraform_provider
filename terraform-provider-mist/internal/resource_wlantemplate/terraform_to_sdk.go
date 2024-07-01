package resource_wlantemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *WlantemplateModel) (mistsdkgo.Template, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewTemplate(plan.Name.ValueString())
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
