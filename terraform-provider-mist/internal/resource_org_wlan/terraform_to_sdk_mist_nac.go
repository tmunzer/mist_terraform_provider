package resource_org_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) mistapigo.WlanMistNac {
	data := mistapigo.NewWlanMistNac()
	data.SetEnabled(d.Enabled.ValueBool())
	return *data
}
