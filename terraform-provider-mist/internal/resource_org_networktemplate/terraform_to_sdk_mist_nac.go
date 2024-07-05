package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) mistapigo.SwitchMistNac {
	data := mistapigo.NewSwitchMistNac()
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetwork(d.Network.ValueString())
	return *data
}
