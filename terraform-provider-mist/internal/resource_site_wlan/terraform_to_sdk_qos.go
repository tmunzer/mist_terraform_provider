package resource_site_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"mistapi/models"
)

func qosTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d QosValue) *models.WlanQos {
	data := models.WlanQos{}
	data.Class = models.ToPointer(models.WlanQosClassEnum(string(d.Class.ValueString())))
	data.Overwrite = d.Overwrite.ValueBoolPointer()

	return &data
}