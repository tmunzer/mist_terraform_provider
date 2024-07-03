package resource_org

import (
	"math/big"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(data *mistapigo.Org) (OrgModel, diag.Diagnostics) {
	var state OrgModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.AlarmtemplateId = types.StringValue(data.GetAlarmtemplateId())
	state.AllowMist = types.BoolValue(data.GetAllowMist())
	state.MspLogoUrl = types.StringValue(data.GetMspLogoUrl())
	state.MspId = types.StringValue(data.GetMspId())
	state.MspName = types.StringValue(data.GetMspName())
	state.Name = types.StringValue(data.GetName())
	state.SessionExpiry = types.NumberValue(big.NewFloat(float64(data.GetSessionExpiry())))
	var items []attr.Value
	for _, item := range data.GetOrggroupIds() {
		tmp := types.StringValue(item)
		items = append(items, tmp)
	}
	state.OrggroupIds, _ = types.ListValue(types.StringType, items)
	return state, diags
}
