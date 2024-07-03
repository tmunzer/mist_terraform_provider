package resource_site_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/commons/hours"
)

func scheduleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ScheduleValue) mistsdkgo.WlanSchedule {
	data := mistsdkgo.NewWlanSchedule()

	data.SetEnabled(d.Enabled.ValueBool())

	hours := hours.HoursTerraformToSdk(ctx, diags, d.Hours)
	data.SetHours(hours)

	return *data
}
