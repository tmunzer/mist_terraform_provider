package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func airwatchTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AirwatchValue) mistsdkgo.WlanAirwatch {

	data := *mistsdkgo.NewWlanAirwatch()

	data.SetApiKey(plan.ApiKey.ValueString())
	data.SetConsoleUrl(plan.ConsoleUrl.ValueString())
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetPassword(plan.Password.ValueString())
	data.SetUsername(plan.Username.ValueString())

	return data
}
