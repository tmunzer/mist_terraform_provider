package resource_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func qosTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d QosValue) mistsdkgo.WlanQos {
	data := mistsdkgo.NewWlanQos()
	data.SetClass(mistsdkgo.WlanQosClass(d.Class.ValueString()))
	data.SetOverwrite(d.Overwrite.ValueBool())

	return *data
}
