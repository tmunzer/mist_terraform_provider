package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"mistapi/models"
)

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) *models.SwitchMistNac {
	data := models.SwitchMistNac{}
	data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	data.Network = models.ToPointer(d.Network.ValueString())
	return &data
}
