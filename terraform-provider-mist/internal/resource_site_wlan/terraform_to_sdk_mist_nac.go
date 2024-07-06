package resource_site_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"mistapi/models"
)

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) models.WlanMistNac {
	data := models.NewWlanMistNac()
	data.SetEnabled(d.Enabled.ValueBool())
	return *data
}
