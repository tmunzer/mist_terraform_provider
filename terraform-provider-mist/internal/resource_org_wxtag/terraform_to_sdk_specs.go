package resource_org_wxtag

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func specsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []mistapigo.WxlanTagSpec {
	var data_list []mistapigo.WxlanTagSpec
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		p := v_interface.(SpecsValue)
		data := mistapigo.NewWxlanTagSpec()
		data.SetPortRange(p.PortRange.ValueString())
		data.SetProtocol(p.Protocol.ValueString())
		data.SetSubnets(mist_transform.ListOfStringTerraformToSdk(ctx, p.Subnets))

		data_list = append(data_list, *data)
	}
	return data_list

}
