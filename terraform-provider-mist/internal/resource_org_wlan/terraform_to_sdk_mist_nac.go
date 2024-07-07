package resource_org_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"mistapi/models"
)

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) *models.WlanMistNac {
	data := models.WlanMistNac{}
	data.Enabled = d.Enabled.ValueBoolPointer()
	return &data
}
