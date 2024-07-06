package resource_org_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func airwatchTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AirwatchValue) models.WlanAirwatch {

	data := *models.NewWlanAirwatch()

	data.SetApiKey(plan.ApiKey.ValueString())
	data.SetConsoleUrl(plan.ConsoleUrl.ValueString())
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetPassword(plan.Password.ValueString())
	data.SetUsername(plan.Username.ValueString())

	return data
}
