package resource_org_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"mistapi/models"
)

func qosTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d QosValue) models.WlanQos {
	data := models.NewWlanQos()
	data.SetClass(models.WlanQosClass(d.Class.ValueString()))
	data.SetOverwrite(d.Overwrite.ValueBool())

	return *data
}
