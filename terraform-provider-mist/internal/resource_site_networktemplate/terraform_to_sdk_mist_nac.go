package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) mistsdkgo.NetworkTemplateMistNac {
	data := mistsdkgo.NewNetworkTemplateMistNac()
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetwork(d.Network.ValueString())
	return *data
}
