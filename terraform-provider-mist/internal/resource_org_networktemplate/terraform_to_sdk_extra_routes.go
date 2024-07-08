package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"mistapi/models"
)

func extraRouteNextQualifiedTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.ExtraRoutePropertiesNextQualifiedProperties {
	data := make(map[string]models.ExtraRoutePropertiesNextQualifiedProperties)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(NextQualifiedValue)
		v_data := models.ExtraRoutePropertiesNextQualifiedProperties{}
		if !v_plan.Metric.IsNull() && !v_plan.Metric.IsUnknown() {
			v_data.Metric = models.NewOptional(models.ToPointer(int(v_plan.Metric.ValueInt64())))
		}
		if !v_plan.Preference.IsNull() && !v_plan.Preference.IsUnknown() {
			v_data.Preference = models.NewOptional(models.ToPointer(int(v_plan.Preference.ValueInt64())))
		}
		data[k] = v_data
	}
	return data
}
func extraRoutesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.ExtraRouteProperties {
	data := make(map[string]models.ExtraRouteProperties)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(ExtraRoutesValue)

		v_data := models.ExtraRouteProperties{}
		if !v_plan.Discard.IsNull() && !v_plan.Discard.IsUnknown() {
			v_data.Discard = models.ToPointer(v_plan.Discard.ValueBool())
		}
		if !v_plan.Metric.IsNull() && !v_plan.Metric.IsUnknown() {
			v_data.Metric = models.NewOptional(models.ToPointer(int(v_plan.Metric.ValueInt64())))
		}
		v_data.NextQualified = extraRouteNextQualifiedTerraformToSdk(ctx, diags, v_plan.NextQualified)
		if !v_plan.NoResolve.IsNull() && !v_plan.NoResolve.IsUnknown() {
			v_data.NoResolve = models.ToPointer(v_plan.NoResolve.ValueBool())
		}
		if !v_plan.Preference.IsNull() && !v_plan.Preference.IsUnknown() {
			v_data.Preference = models.NewOptional(models.ToPointer(int(v_plan.Preference.ValueInt64())))
		}
		if !v_plan.Via.IsNull() && !v_plan.Via.IsUnknown() {
			v_data.Via = models.ToPointer(v_plan.Via.ValueString())
		}
		data[k] = v_data
	}
	return data
}
