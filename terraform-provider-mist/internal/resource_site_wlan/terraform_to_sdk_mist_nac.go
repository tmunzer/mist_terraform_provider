package resource_site_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) mistsdkgo.WlanMistNac {
	data := mistsdkgo.NewWlanMistNac()
	data.SetEnabled(d.Enabled.ValueBool())
	return *data
}
