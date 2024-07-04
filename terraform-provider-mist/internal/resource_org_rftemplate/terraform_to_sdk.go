package resource_org_rftemplate

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgRftemplateModel) (mistapigo.RfTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistapigo.NewRfTemplate(plan.Name.ValueString())

	data.SetId(plan.Id.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())

	data.SetAntGain24(int32(*plan.AntGain24.ValueInt64Pointer()))
	data.SetAntGain5(int32(plan.AntGain5.ValueInt64()))
	data.SetAntGain6(int32(plan.AntGain6.ValueInt64()))

	band24 := band24TerraformToSdk(ctx, &diags, plan.Band24)
	data.SetBand24(band24)

	data.SetBand24Usage(mistapigo.RadioBand24Usage(plan.Band24Usage.ValueString()))

	band5 := band5TerraformToSdk(ctx, &diags, plan.Band5)
	data.SetBand5(band5)

	band6 := band6TerraformToSdk(ctx, &diags, plan.Band6)
	data.SetBand6(band6)

	data.SetCountryCode(plan.CountryCode.ValueString())

	model_specific := modelSpecificTerraformToSdk(ctx, &diags, plan.ModelSpecific)
	data.SetModelSpecific(model_specific)

	data.SetScanningEnabled(plan.ScanningEnabled.ValueBool())

	return data, diags
}
