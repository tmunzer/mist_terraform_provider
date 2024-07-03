package resource_org_wlantemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *mistapigo.Template) (OrgWlantemplateModel, diag.Diagnostics) {
	var state OrgWlantemplateModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())

	state.Applies = appliesSdkToTerraform(ctx, &diags, data.GetApplies())
	state.DeviceprofileIds = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDeviceprofileIds())
	state.Exceptions = exceptionsSdkToTerraform(ctx, &diags, data.GetExceptions())
	state.FilterByDeviceprofile = types.BoolValue(data.GetFilterByDeviceprofile())

	return state, diags
}
