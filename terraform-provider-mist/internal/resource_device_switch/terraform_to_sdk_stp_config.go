package resource_device_switch

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func stpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d StpConfigValue) mistapigo.SwitchStpConfig {
	tflog.Debug(ctx, "stpConfigTerraformToSdk")

	data := *mistapigo.NewSwitchStpConfig()

	data.SetType(mistapigo.SwitchStpConfigType(d.StpConfigType.ValueString()))

	return data
}
