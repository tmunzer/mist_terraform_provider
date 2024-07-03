package resource_site_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func qosTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d QosValue) mistapigo.WlanQos {
	data := mistapigo.NewWlanQos()
	data.SetClass(mistapigo.WlanQosClass(d.Class.ValueString()))
	data.SetOverwrite(d.Overwrite.ValueBool())

	return *data
}
