package resource_org_wlan

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func hotspot20TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Hotspot20Value) mistapigo.WlanHotspot20 {

	var operators []mistapigo.WlanHotspot20OperatorsItem
	for _, v := range plan.Operators.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		op, _ := mistapigo.NewWlanHotspot20OperatorsItemFromValue(v_plan.ValueString())
		operators = append(operators, *op)
	}

	data := *mistapigo.NewWlanHotspot20()
	data.SetDomainName(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DomainName))
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetNaiRealms(mist_transform.ListOfStringTerraformToSdk(ctx, plan.NaiRealms))
	data.SetOperators(operators)
	data.SetRcoi(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Rcoi))
	data.SetVenueName(plan.VenueName.ValueString())

	return data
}
