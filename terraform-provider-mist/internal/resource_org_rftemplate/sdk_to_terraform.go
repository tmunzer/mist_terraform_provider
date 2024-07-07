package resource_org_rftemplate

import (
	"context"
	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data models.RfTemplate) (OrgRftemplateModel, diag.Diagnostics) {
	var state OrgRftemplateModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Name = types.StringValue(data.Name)

	state.AntGain24 = types.Int64Value(int64(*data.AntGain24))

	state.AntGain5 = types.Int64Value(int64(*data.AntGain5))

	state.AntGain6 = types.Int64Value(int64(*data.AntGain6))

	state.Band24 = band24SdkToTerraform(ctx, &diags, data.Band24)

	state.Band24Usage = types.StringValue(string(*data.Band24Usage))

	state.Band5 = band5SdkToTerraform(ctx, &diags, data.Band5)

	state.Band6 = band6SdkToTerraform(ctx, &diags, data.Band6)

	state.CountryCode = types.StringValue(*data.CountryCode)

	state.ModelSpecific = modelSpecificSdkToTerraform(ctx, &diags, data.ModelSpecific)

	return state, diags
}
