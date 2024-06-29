package resource_rftemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.RfTemplate) (RftemplateModel, diag.Diagnostics) {
	var state RftemplateModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())

	state.AntGain24 = types.Int64Value(int64(data.GetAntGain24()))

	state.AntGain5 = types.Int64Value(int64(data.GetAntGain5()))

	state.AntGain6 = types.Int64Value(int64(data.GetAntGain6()))

	state.Band24 = band24SdkToTerraform(ctx, &diags, data.GetBand24())

	state.Band24Usage = types.StringValue(string(data.GetBand24Usage()))

	state.Band5 = band5SdkToTerraform(ctx, &diags, data.GetBand5())

	state.Band6 = band6SdkToTerraform(ctx, &diags, data.GetBand6())

	state.CountryCode = types.StringValue(data.GetCountryCode())

	state.ModelSpecific = modelSpecificSdkToTerraform(ctx, &diags, data.GetModelSpecific())

	return state, diags
}
