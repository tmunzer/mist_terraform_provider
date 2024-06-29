package resource_rftemplate

import (
	"context"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func modelSpecificTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.RfTemplateModelSpecificProperty {

	data_map := make(map[string]mistsdkgo.RfTemplateModelSpecificProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ModelSpecificValue)
		data := *mistsdkgo.NewRfTemplateModelSpecificProperty()

		data.SetAntGain24(int32(*plan.AntGain24.ValueInt64Pointer()))
		data.SetAntGain5(int32(plan.AntGain5.ValueInt64()))
		data.SetAntGain6(int32(plan.AntGain6.ValueInt64()))

		plan_band24, _ := NewBand24Value(plan.Band24.AttributeTypes(ctx), plan.Band24.Attributes())
		band24 := band24TerraformToSdk(ctx, diags, plan_band24)
		data.SetBand24(band24)

		data.SetBand24Usage(mistsdkgo.ApRadioBand24Usage(plan.Band24Usage.ValueString()))

		plan_band5, _ := NewBand5Value(plan.Band5.AttributeTypes(ctx), plan.Band5.Attributes())
		band5 := band5TerraformToSdk(ctx, diags, plan_band5)
		data.SetBand5(band5)

		plan_band6, _ := NewBand6Value(plan.Band6.AttributeTypes(ctx), plan.Band6.Attributes())
		band6 := band6TerraformToSdk(ctx, diags, plan_band6)
		data.SetBand6(band6)

		data_map[k] = data
	}
	return data_map
}
