package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func extraRoute6NextQualifiedTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.ExtraRoute6PropertiesNextQualifiedProperties {
	data := make(map[string]mistapigo.ExtraRoute6PropertiesNextQualifiedProperties)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(NextQualifiedValue)
		v_data := *mistapigo.NewExtraRoute6PropertiesNextQualifiedProperties()
		v_data.SetMetric(int32(v_plan.Metric.ValueInt64()))
		v_data.SetPreference(int32(v_plan.Preference.ValueInt64()))
		data[k] = v_data
	}
	return data
}
func extraRoutes6TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.ExtraRoute6Properties {
	data := make(map[string]mistapigo.ExtraRoute6Properties)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(ExtraRoutesValue)
		next_qualified := extraRoute6NextQualifiedTerraformToSdk(ctx, diags, v_plan.NextQualified)

		v_data := *mistapigo.NewExtraRoute6Properties()
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
