package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func appLimitTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AppLimitValue) mistsdkgo.WlanAppLimit {

	data := *mistsdkgo.NewWlanAppLimit()

	app_lmimit := make(map[string]int32)
	for k, v := range plan.Apps.Elements() {
		var v_interface interface{} = v
		app_lmimit[k] = int32(v_interface.(int64))
	}

	wxtags_limit := make(map[string]int32)
	for k, v := range plan.WxtagIds.Elements() {
		var v_interface interface{} = v
		wxtags_limit[k] = int32(v_interface.(int64))
	}

	data.SetApps(app_lmimit)
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetWxtagIds(wxtags_limit)

	return data
}