package resource_wxtag

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func specsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []mistsdkgo.WxlanTagSpec {
	var data_list []mistsdkgo.WxlanTagSpec
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		p := v_interface.(SpecsValue)
		data := mistsdkgo.NewWxlanTagSpec()
		data.SetPortRange(p.PortRange.ValueString())
		data.SetProtocol(p.Protocol.ValueString())
		data.SetSubnets(mist_transform.ListOfStringTerraformToSdk(ctx, p.Subnets))

		data_list = append(data_list, *data)
	}
	return data_list

}
