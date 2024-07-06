package resource_org_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"terraform-provider-mist/internal/commons/hours"

	"mistapi/models"
)

func scheduleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ScheduleValue) models.WlanSchedule {
	data := models.NewWlanSchedule()

	data.SetEnabled(d.Enabled.ValueBool())

	hours := hours.HoursTerraformToSdk(ctx, diags, d.Hours)
	data.SetHours(hours)

	return *data
}
