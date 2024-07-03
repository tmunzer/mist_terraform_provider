package resource_org_wlan

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func appLimitTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AppLimitValue) mistapigo.WlanAppLimit {

	data := *mistapigo.NewWlanAppLimit()

	app_limit := make(map[string]int32)
	for k, v := range plan.Apps.Elements() {
		var v_interface interface{} = v
		app_limit[k] = int32(v_interface.(int64))
	}

	wxtags_limit := make(map[string]int32)
	for k, v := range plan.WxtagIds.Elements() {
		var v_interface interface{} = v
		wxtags_limit[k] = int32(v_interface.(int64))
	}

	data.SetApps(app_limit)
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetWxtagIds(wxtags_limit)

	return data
}
