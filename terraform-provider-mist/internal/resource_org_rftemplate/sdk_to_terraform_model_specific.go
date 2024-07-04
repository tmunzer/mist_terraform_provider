package resource_org_rftemplate

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func modelSpecificSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistapigo.RfTemplateModelSpecificProperty) basetypes.MapValue {
	state_value_map_attr_type := ModelSpecificValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"ant_gain_24":        types.Int64Value(int64(v.GetAntGain24())),
			"ant_gain_5":         types.Int64Value(int64(v.GetAntGain5())),
			"ant_gain_6":         types.Int64Value(int64(v.GetAntGain6())),
			"band_24":            band24SdkToTerraform(ctx, diags, v.GetBand24()),
			"band_24_usage":      types.StringValue(string(v.GetBand24Usage())),
			"band_5":             band5SdkToTerraform(ctx, diags, v.GetBand5()),
			"band_5_on_24_radio": band5SdkToTerraform(ctx, diags, v.GetBand5On24Radio()),
			"band_6":             band6SdkToTerraform(ctx, diags, v.GetBand6()),
		}
		n, e := NewModelSpecificValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := ModelSpecificValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
