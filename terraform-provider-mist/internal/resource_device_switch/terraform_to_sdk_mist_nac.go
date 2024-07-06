package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"mistapi/models"
)

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) models.SwitchMistNac {
	data := models.NewSwitchMistNac()
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetwork(d.Network.ValueString())
	return *data
}
