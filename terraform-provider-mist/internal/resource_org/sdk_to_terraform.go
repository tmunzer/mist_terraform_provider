package resource_org

import (
	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(data models.Org) (OrgModel, diag.Diagnostics) {
	var state OrgModel
	var diags diag.Diagnostics

	var alarmtemplate_id basetypes.StringValue
	var allow_mist basetypes.BoolValue
	var msp_id basetypes.StringValue
	var msp_name basetypes.StringValue
	var session_expiry basetypes.Int64Value
	state.Id = types.StringValue(data.Id.String())

	if data.AlarmtemplateId.IsValueSet() {
		alarmtemplate_id = types.StringValue(data.AlarmtemplateId.Value().String())
	}
	if data.AllowMist != nil {
		allow_mist = types.BoolValue(*data.AllowMist)
	}
	if data.MspId != nil {
		msp_id = types.StringValue(data.MspId.String())
	}
	if data.MspName != nil {
		msp_name = types.StringPointerValue(data.MspName)
	}
	if data.SessionExpiry != nil {
		session_expiry = types.Int64Value(*data.SessionExpiry)
	}

	state.AlarmtemplateId = alarmtemplate_id
	state.AllowMist = allow_mist
	state.MspId = msp_id
	state.MspName = msp_name
	state.Name = types.StringValue(data.Name)
	state.SessionExpiry = session_expiry

	var items []attr.Value
	for _, item := range data.OrggroupIds {
		tmp := types.StringValue(item.String())
		items = append(items, tmp)
	}
	state.OrggroupIds, _ = types.ListValue(types.StringType, items)
	return state, diags
}
