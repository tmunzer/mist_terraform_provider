package resource_wlantemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.Template) (WlantemplateModel, diag.Diagnostics) {
	var state WlantemplateModel
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
