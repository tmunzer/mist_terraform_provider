package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

// //////////////////////////////////
// ////////// CLIENTS
func snmpConfigClientListTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.SnmpConfigClientList {
	var data_list []mistapigo.SnmpConfigClientList
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ClientListValue)
		data := mistapigo.NewSnmpConfigClientList()
		data.SetClients(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Clients))
		data.SetClientListName(plan.ClientListName.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

// //////////////////////////////////
// ////////// TRAPS
func snmpConfigTrapGroupsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.SnmpConfigTrapGroup {
	var data_list []mistapigo.SnmpConfigTrapGroup
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TrapGroupsValue)
		data := mistapigo.NewSnmpConfigTrapGroup()
		data.SetCategories(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Categories))
		data.SetGroupName(plan.GroupName.ValueString())
		data.SetTargets(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Targets))
		data.SetVersion(mistapigo.SnmpConfigTrapVerion(plan.Version.ValueString()))

		data_list = append(data_list, *data)
	}

	return data_list
}

// //////////////////////////////////
// ////////// V2c
func snmpConfigV2cTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.SnmpConfigV2cConfig {
	var data_list []mistapigo.SnmpConfigV2cConfig
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(V2cConfigValue)
		data := mistapigo.NewSnmpConfigV2cConfig()
		data.SetAuthorization(plan.Authorization.ValueString())
		data.SetClientListName(plan.ClientListName.ValueString())
		data.SetCommunityName(plan.CommunityName.ValueString())
		data.SetView(plan.View.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

// //////////////////////////////////
// ////////// V3
// V3 NOTIFY
func snmpConfigV3NotifyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.Snmpv3ConfigNotifyItems {
	var data_list []mistapigo.Snmpv3ConfigNotifyItems
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NotifyValue)
		data := mistapigo.NewSnmpv3ConfigNotifyItems()
		data.SetName(plan.Name.ValueString())
		data.SetTag(plan.Tag.ValueString())
		data.SetType(mistapigo.Snmpv3ConfigNotifyType(plan.Tag.ValueString()))

		data_list = append(data_list, *data)
	}

	return data_list
}

func snmpConfigV3NotifyFilterContentTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.Snmpv3ConfigNotifyFilterItemContent {
	var data_list []mistapigo.Snmpv3ConfigNotifyFilterItemContent
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(Snmpv3ContentsValue)
		data := mistapigo.NewSnmpv3ConfigNotifyFilterItemContent()

		data.SetInclude(plan.Include.ValueBool())
		data.SetOid(plan.Oid.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

func snmpConfigV3NotifyFilterTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.Snmpv3ConfigNotifyFilterItem {
	var data_list []mistapigo.Snmpv3ConfigNotifyFilterItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NotifyFilterValue)
		data := mistapigo.NewSnmpv3ConfigNotifyFilterItem()

		content := snmpConfigV3NotifyFilterContentTerraformToSdk(ctx, diags, plan.Snmpv3Contents)

		data.SetContents(content)
		data.SetProfileName(plan.ProfileName.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

// V3 TARGETS
func snmpConfigV3TargetAddressTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.Snmpv3ConfigTargetAddressItem {
	var data_list []mistapigo.Snmpv3ConfigTargetAddressItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TargetAddressValue)
		data := mistapigo.NewSnmpv3ConfigTargetAddressItem()

		data.SetAddress(plan.Address.ValueString())
		data.SetAddressMask(plan.AddressMask.ValueString())
		data.SetPort(int32(plan.Port.ValueInt64()))
		data.SetTagList(plan.TagList.ValueString())
		data.SetTargetAddressName(plan.TargetAddressName.ValueString())
		data.SetTargetParameters(plan.TargetParameters.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

func snmpConfigV3TargetParametersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.Snmpv3ConfigTargetParam {
	var data_list []mistapigo.Snmpv3ConfigTargetParam
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TargetParametersValue)
		data := mistapigo.NewSnmpv3ConfigTargetParam()

		data.SetMessageProcessingModel(mistapigo.Snmpv3ConfigTargetParamMessProcessModel(plan.MessageProcessingModel.ValueString()))
		data.SetName(plan.Name.ValueString())
		data.SetNotifyFilter(plan.NotifyFilter.ValueString())
		data.SetSecurityLevel(mistapigo.Snmpv3ConfigTargetParamSecurityLevel(plan.SecurityLevel.ValueString()))
		data.SetSecurityModel(mistapigo.Snmpv3ConfigTargetParamSecurityModel(plan.SecurityModel.ValueString()))
		data.SetSecurityName(plan.SecurityName.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

// V3 USM
func snmpConfigV3UsmUsersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.SnmpUsmpUser {
	var data_list []mistapigo.SnmpUsmpUser
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(Snmpv3UsersValue)
		data := mistapigo.NewSnmpUsmpUser()

		data.SetAuthenticationPassword(plan.AuthenticationPassword.ValueString())
		data.SetAuthenticationType(mistapigo.SnmpUsmpUserAuthenticationType(plan.AuthenticationType.ValueString()))
		data.SetEncryptionPassword(plan.EncryptionPassword.ValueString())
		data.SetEncryptionType(mistapigo.SnmpUsmpUserEncryptionType(plan.EncryptionType.ValueString()))
		data.SetName(plan.Name.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}
func snmpConfigV3UsmTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.SnmpUsm {
	data := *mistapigo.NewSnmpUsm()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(UsmValue)

		users := snmpConfigV3UsmUsersTerraformToSdk(ctx, diags, plan.Snmpv3Users)

		data.SetEngineType(mistapigo.SnmpUsmEngineType(plan.EngineType.ValueString()))
		data.SetEngineId(plan.Engineid.ValueString())
		data.SetUsers(users)

		return data
	}
}

// V3 VACM ACCESS
func snmpConfigV3VacmAccessPrefixTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.SnmpVacmAccessItemPrefixListItem {
	var data_list []mistapigo.SnmpVacmAccessItemPrefixListItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PrefixListValue)
		data := mistapigo.NewSnmpVacmAccessItemPrefixListItem()

		data.SetContextPrefix(plan.ContextPrefix.ValueString())
		data.SetNotifyView(plan.NotifyView.ValueString())
		data.SetReadView(plan.ReadView.ValueString())
		data.SetSecurityLevel(mistapigo.SnmpVacmAccessItemPrefixListItemLevel(plan.SecurityLevel.ValueString()))
		data.SetSecurityModel(mistapigo.SnmpVacmAccessItemPrefixListItemModel(plan.SecurityModel.ValueString()))
		data.SetContextPrefix(plan.PrefixListType.ValueString())
		data.SetWriteView(plan.WriteView.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}
func snmpConfigV3VacmAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.SnmpVacmAccessItem {
	var data_list []mistapigo.SnmpVacmAccessItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(AccessValue)
		data := mistapigo.NewSnmpVacmAccessItem()

		prefix_list := snmpConfigV3VacmAccessPrefixTerraformToSdk(ctx, diags, plan.PrefixList)

		data.SetGroupName(plan.GroupName.ValueString())
		data.SetPrefixList(prefix_list)

		data_list = append(data_list, *data)
	}

	return data_list
}

// V3 VACM SEC TO GROUP
func snmpConfigV3VacmSecurityToGroupContentTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.SnmpVacmSecurityToGroupContentItem {
	var data_list []mistapigo.SnmpVacmSecurityToGroupContentItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(Snmpv3VacmContentValue)
		data := mistapigo.NewSnmpVacmSecurityToGroupContentItem()

		data.SetGroup(plan.Group.ValueString())
		data.SetSecurityName(plan.SecurityName.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}
func snmpConfigV3VacmSecurityToGroupTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.SnmpVacmSecurityToGroup {
	data := *mistapigo.NewSnmpVacmSecurityToGroup()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(SecurityToGroupValue)

		content := snmpConfigV3VacmSecurityToGroupContentTerraformToSdk(ctx, diags, plan.Snmpv3VacmContent)

		data.SetSecurityModel(mistapigo.SnmpVacmSecurityModel(plan.SecurityModel.ValueString()))
		data.SetContent(content)

		return data
	}
}
func snmpConfigV3VacmTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.SnmpVacm {
	data := *mistapigo.NewSnmpVacm()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(VacmValue)

		access := snmpConfigV3VacmAccessTerraformToSdk(ctx, diags, plan.Access)
		security_to_group := snmpConfigV3VacmSecurityToGroupTerraformToSdk(ctx, diags, plan.SecurityToGroup)

		data.SetAccess(access)
		data.SetSecurityToGroup(security_to_group)

		return data
	}
}

// V3 MAIN
func snmpConfigV3TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.Snmpv3Config {
	data := *mistapigo.NewSnmpv3Config()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(V3ConfigValue)

		notify := snmpConfigV3NotifyTerraformToSdk(ctx, diags, plan.Notify)
		notify_filter := snmpConfigV3NotifyFilterTerraformToSdk(ctx, diags, plan.NotifyFilter)
		target_address := snmpConfigV3TargetAddressTerraformToSdk(ctx, diags, plan.TargetAddress)
		target_parameters := snmpConfigV3TargetParametersTerraformToSdk(ctx, diags, plan.TargetParameters)
		usm := snmpConfigV3UsmTerraformToSdk(ctx, diags, plan.Usm)
		vacm := snmpConfigV3VacmTerraformToSdk(ctx, diags, plan.Vacm)

		data.SetNotify(notify)
		data.SetNotifyFilter(notify_filter)
		data.SetTargetAddress(target_address)
		data.SetTargetParameters(target_parameters)
		data.SetUsm(usm)
		data.SetVacm(vacm)
		return data
	}
}

// //////////////////////////////////////////////
// ////////// VIEWS
func snmpConfigViewsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.SnmpConfigView {
	var data_list []mistapigo.SnmpConfigView
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ViewsValue)
		data := mistapigo.NewSnmpConfigView()

		data.SetInclude(plan.Include.ValueBool())
		data.SetOid(plan.Oid.ValueString())
		data.SetViewName(plan.ViewName.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

// //////////////////////////////////////////////
// ////////// MAIN
func snmpConfigSyslogTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SnmpConfigValue) mistapigo.SnmpConfig {

	client_list := snmpConfigClientListTerraformToSdk(ctx, diags, d.ClientList)
	trap_groups := snmpConfigTrapGroupsTerraformToSdk(ctx, diags, d.TrapGroups)
	v2c_config := snmpConfigV2cTerraformToSdk(ctx, diags, d.V2cConfig)
	v3_config := snmpConfigV3TerraformToSdk(ctx, diags, d.V3Config)
	views := snmpConfigViewsTerraformToSdk(ctx, diags, d.Views)

	data := mistapigo.NewSnmpConfig()
	data.SetClientList(client_list)
	data.SetContact(d.Contact.ValueString())
	data.SetDescription(d.Description.ValueString())
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetEngineId(mistapigo.SnmpConfigEngineId(d.EngineId.ValueString()))
	data.SetLocation(d.Location.ValueString())
	data.SetName(d.Name.ValueString())
	data.SetNetwork(d.Network.ValueString())
	data.SetTrapGroups(trap_groups)
	data.SetV2cConfig(v2c_config)
	data.SetV3Config(v3_config)
	data.SetViews(views)

	return *data
}
