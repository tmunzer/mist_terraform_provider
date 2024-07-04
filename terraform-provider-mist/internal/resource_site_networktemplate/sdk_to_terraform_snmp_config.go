package resource_site_networktemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

// //////////////////////////////////
// ////////// CLIENTS
func snmpClientListSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.SnmpConfigClientList) basetypes.ListValue {
	var data_list = []ClientListValue{}
	for _, v := range data {
		data_map_value := map[string]attr.Value{
			"client_list_name": types.StringValue(v.GetClientListName()),
			"clients":          mist_transform.ListOfStringSdkToTerraform(ctx, v.GetClients()),
		}
		data, e := NewClientListValue(ClientListValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, ClientListValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// TRAP GROUPS
func snmpTrapGroupsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.SnmpConfigTrapGroup) basetypes.ListValue {
	var data_list = []TrapGroupsValue{}
	for _, v := range data {
		data_map_value := map[string]attr.Value{
			"categories": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetCategories()),
			"group_name": types.StringValue(v.GetGroupName()),
			"targets":    mist_transform.ListOfStringSdkToTerraform(ctx, v.GetTargets()),
			"version":    types.StringValue(string(v.GetVersion())),
		}
		data, e := NewTrapGroupsValue(TrapGroupsValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, TrapGroupsValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// V2
func snmpV2cSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.SnmpConfigV2cConfig) basetypes.ListValue {
	var data_list = []V2cConfigValue{}
	for _, v := range data {
		data_map_value := map[string]attr.Value{
			"authorization":    types.StringValue(v.GetAuthorization()),
			"client_list_name": types.StringValue(v.GetClientListName()),
			"community_name":   types.StringValue(v.GetCommunityName()),
			"view":             types.StringValue(v.GetView()),
		}
		data, e := NewV2cConfigValue(V2cConfigValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, V2cConfigValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// V3
// NOTIFY
func snmpV3NotifySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.Snmpv3ConfigNotifyItems) basetypes.ListValue {
	var data_list = []NotifyValue{}
	for _, v := range data {
		data_map_value := map[string]attr.Value{
			"name": types.StringValue(v.GetName()),
			"tag":  types.StringValue(v.GetTag()),
			"type": types.StringValue(string(v.GetType())),
		}
		data, e := NewNotifyValue(NotifyValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, NotifyValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// NOTIFY Filter
func snmpV3NotifyFilterContentSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.Snmpv3ConfigNotifyFilterItemContent) basetypes.ListValue {
	var data_list = []Snmpv3ContentsValue{}
	for _, v := range data {
		data_map_value := map[string]attr.Value{
			"include": types.BoolValue(v.GetInclude()),
			"oid":     types.StringValue(v.GetOid()),
		}
		data, e := NewSnmpv3ContentsValue(Snmpv3ContentsValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, Snmpv3ContentsValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
func snmpV3NotifyFilterSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.Snmpv3ConfigNotifyFilterItem) basetypes.ListValue {
	var data_list = []NotifyFilterValue{}
	for _, v := range data {

		contents := snmpV3NotifyFilterContentSdkToTerraform(ctx, diags, v.GetContents())

		data_map_value := map[string]attr.Value{
			"profile_name": types.StringValue(v.GetProfileName()),
			"contents":     contents,
		}
		data, e := NewNotifyFilterValue(NotifyFilterValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, NotifyFilterValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// TARGET ADDRESS
func snmpV3TargetAddressSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.Snmpv3ConfigTargetAddressItem) basetypes.ListValue {
	var data_list = []TargetAddressValue{}
	for _, v := range data {

		data_map_value := map[string]attr.Value{
			"address":             types.StringValue(v.GetAddress()),
			"address_mask":        types.StringValue(v.GetAddressMask()),
			"port":                types.Int64Value(int64(v.GetPort())),
			"tag_list":            types.StringValue(v.GetTagList()),
			"target_address_name": types.StringValue(v.GetTargetAddressName()),
			"target_parameters":   types.StringValue(v.GetTargetParameters()),
		}
		data, e := NewTargetAddressValue(TargetAddressValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, TargetAddressValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// TARGET PARAMETERS
func snmpV3TargetParametersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.Snmpv3ConfigTargetParam) basetypes.ListValue {
	var data_list = []TargetParametersValue{}
	for _, v := range data {

		data_map_value := map[string]attr.Value{
			"message_processing_model": types.StringValue(string(v.GetMessageProcessingModel())),
			"name":                     types.StringValue(v.GetName()),
			"notify_filter":            types.StringValue(v.GetNotifyFilter()),
			"security_level":           types.StringValue(string(v.GetSecurityLevel())),
			"security_model":           types.StringValue(string(v.GetSecurityModel())),
			"security_name":            types.StringValue(v.GetSecurityName()),
		}
		data, e := NewTargetParametersValue(TargetParametersValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, TargetParametersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// USM
func snmpV3UsmUsersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.SnmpUsmpUser) basetypes.ListValue {
	var data_list = []Snmpv3UsersValue{}
	for _, v := range data {

		data_map_value := map[string]attr.Value{
			"authentication_password": types.StringValue(v.GetAuthenticationPassword()),
			"authentication_type":     types.StringValue(string(v.GetAuthenticationType())),
			"encryption_password":     types.StringValue(v.GetEncryptionPassword()),
			"encryption_type":         types.StringValue(string(v.GetEncryptionType())),
			"name":                    types.StringValue(v.GetName()),
		}
		data, e := NewSnmpv3UsersValue(Snmpv3UsersValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, Snmpv3UsersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
func snmpV3UsmSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SnmpUsm) basetypes.ObjectValue {
	users := snmpV3UsmUsersSdkToTerraform(ctx, diags, d.GetUsers())

	r_attr_type := UsmValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"engine_type": types.StringValue(string(d.GetEngineType())),
		"engineid":    types.StringValue(d.GetEngineId()),
		"users":       users,
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

// VACM ACCESS
func snmpV3VacmAccessPrefixListSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.SnmpVacmAccessItemPrefixListItem) basetypes.ListValue {
	var data_list = []PrefixListValue{}
	for _, v := range data {

		data_map_value := map[string]attr.Value{
			"context_prefix": types.StringValue(v.GetContextPrefix()),
			"notify_view":    types.StringValue(v.GetNotifyView()),
			"read_view":      types.StringValue(v.GetReadView()),
			"security_level": types.StringValue(string(v.GetSecurityLevel())),
			"security_model": types.StringValue(string(v.GetSecurityModel())),
			"type":           types.StringValue(string(v.GetType())),
			"write_view":     types.StringValue(v.GetWriteView()),
		}
		data, e := NewPrefixListValue(PrefixListValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, PrefixListValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
func snmpV3VacmAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.SnmpVacmAccessItem) basetypes.ListValue {
	var data_list = []AccessValue{}
	for _, v := range data {

		prefix_list := snmpV3VacmAccessPrefixListSdkToTerraform(ctx, diags, v.GetPrefixList())

		data_map_value := map[string]attr.Value{
			"group_name":  types.StringValue(v.GetGroupName()),
			"prefix_list": prefix_list,
		}
		data, e := NewAccessValue(AccessValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, AccessValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// VACM SECURITY TO GROUP
func snmpV3VacmSecurityToGroupContentSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.SnmpVacmSecurityToGroupContentItem) basetypes.ListValue {
	var data_list = []Snmpv3VacmContentValue{}
	for _, v := range data {

		data_map_value := map[string]attr.Value{
			"group":         types.StringValue(v.GetGroup()),
			"security_name": types.StringValue(v.GetSecurityName()),
		}
		data, e := NewSnmpv3VacmContentValue(Snmpv3VacmContentValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, Snmpv3VacmContentValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
func snmpV3VacmSecurityToGroupSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SnmpVacmSecurityToGroup) basetypes.ObjectValue {
	content := snmpV3VacmSecurityToGroupContentSdkToTerraform(ctx, diags, d.GetContent())

	r_attr_type := SecurityToGroupValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"security_model": types.StringValue(string(d.GetSecurityModel())),
		"content":        content,
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

// VACM
func snmpV3VacmSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SnmpVacm) basetypes.ObjectValue {
	access := snmpV3VacmAccessSdkToTerraform(ctx, diags, d.GetAccess())
	security_to_group := snmpV3VacmSecurityToGroupSdkToTerraform(ctx, diags, d.GetSecurityToGroup())

	r_attr_type := VacmValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"access":            access,
		"security_to_group": security_to_group,
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

// V3
func snmpV3SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistapigo.Snmpv3Config) basetypes.ObjectValue {
	notify := snmpV3NotifySdkToTerraform(ctx, diags, data.GetNotify())
	notify_filter := snmpV3NotifyFilterSdkToTerraform(ctx, diags, data.GetNotifyFilter())
	target_address := snmpV3TargetAddressSdkToTerraform(ctx, diags, data.GetTargetAddress())
	target_parameters := snmpV3TargetParametersSdkToTerraform(ctx, diags, data.GetTargetParameters())
	usm := snmpV3UsmSdkToTerraform(ctx, diags, data.GetUsm())
	vacm := snmpV3VacmSdkToTerraform(ctx, diags, data.GetVacm())

	r_attr_type := V3ConfigValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"notify":            notify,
		"notify_filter":     notify_filter,
		"target_address":    target_address,
		"target_parameters": target_parameters,
		"usm":               usm,
		"vacm":              vacm,
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

// ////////////////////////////////
// VIEWS
func snmpConfigViewsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.SnmpConfigView) basetypes.ListValue {
	var data_list = []ViewsValue{}
	for _, v := range data {

		data_map_value := map[string]attr.Value{
			"include":   types.BoolValue(v.GetInclude()),
			"oid":       types.StringValue(v.GetOid()),
			"view_name": types.StringValue(v.GetViewName()),
		}
		data, e := NewViewsValue(ViewsValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, ViewsValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// MAIN
func snmpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SnmpConfig) SnmpConfigValue {

	client_list := snmpClientListSdkToTerraform(ctx, diags, d.GetClientList())
	trap_groups := snmpTrapGroupsSdkToTerraform(ctx, diags, d.GetTrapGroups())
	v2c_config := snmpV2cSdkToTerraform(ctx, diags, d.GetV2cConfig())
	v3_config := snmpV3SdkToTerraform(ctx, diags, d.GetV3Config())
	views := snmpConfigViewsSdkToTerraform(ctx, diags, d.GetViews())

	data_map_attr_type := SnmpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"client_list": client_list,
		"contact":     types.StringValue(d.GetContact()),
		"description": types.StringValue(d.GetDescription()),
		"enabled":     types.BoolValue(d.GetEnabled()),
		"engine_id":   types.StringValue(string(d.GetEngineId())),
		"location":    types.StringValue(d.GetLocation()),
		"name":        types.StringValue(d.GetName()),
		"network":     types.StringValue(d.GetNetwork()),
		"trap_groups": trap_groups,
		"v2c_config":  v2c_config,
		"v3_config":   v3_config,
		"views":       views,
	}

	state_result, e := NewSnmpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return state_result
}
