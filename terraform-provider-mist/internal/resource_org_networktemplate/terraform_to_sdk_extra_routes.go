package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func extraRouteNextQualifiedTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.ExtraRoutePropertiesNextQualifiedProperties {
	data := make(map[string]mistsdkgo.ExtraRoutePropertiesNextQualifiedProperties)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(NextQualifiedValue)
		v_data := *mistsdkgo.NewExtraRoutePropertiesNextQualifiedProperties()
		v_data.SetMetric(int32(v_plan.Metric.ValueInt64()))
		v_data.SetPreference(int32(v_plan.Preference.ValueInt64()))
		data[k] = v_data
	}
	return data
}
func extraRoutesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.ExtraRouteProperties {
	data := make(map[string]mistsdkgo.ExtraRouteProperties)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(ExtraRoutesValue)
		next_qualified := extraRouteNextQualifiedTerraformToSdk(ctx, diags, v_plan.NextQualified)

		v_data := *mistsdkgo.NewExtraRouteProperties()
		v_data.SetDiscard(v_plan.Discard.ValueBool())
		v_data.SetMetric(int32(v_plan.Metric.ValueInt64()))
		v_data.SetNextQualified(next_qualified)
		v_data.SetNoResolve(v_plan.NoResolve.ValueBool())
		v_data.SetPreference(int32(v_plan.Preference.ValueInt64()))
		v_data.SetVia(v_plan.Via.ValueString())
		data[k] = v_data
	}
	return data
}
