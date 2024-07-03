package resource_site_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"terraform-provider-mist/internal/commons/hours"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func scheduleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ScheduleValue) mistapigo.WlanSchedule {
	data := mistapigo.NewWlanSchedule()

	data.SetEnabled(d.Enabled.ValueBool())

	hours := hours.HoursTerraformToSdk(ctx, diags, d.Hours)
	data.SetHours(hours)

	return *data
}
