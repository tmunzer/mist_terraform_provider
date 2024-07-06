package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

// //////////////////////////////////
// ////////// CLIENTS
func snmpConfigClientListTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpConfigClientList {
	var data_list []models.SnmpConfigClientList
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ClientListValue)
		data := models.NewSnmpConfigClientList()
		data.SetClients(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Clients))
		data.SetClientListName(plan.ClientListName.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

// //////////////////////////////////
// ////////// TRAPS
func snmpConfigTrapGroupsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpConfigTrapGroup {
	var data_list []models.SnmpConfigTrapGroup
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TrapGroupsValue)
		data := models.NewSnmpConfigTrapGroup()
		data.SetCategories(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Categories))
		data.SetGroupName(plan.GroupName.ValueString())
		data.SetTargets(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Targets))
		data.SetVersion(models.SnmpConfigTrapVerion(plan.Version.ValueString()))

		data_list = append(data_list, *data)
	}

	return data_list
}

// //////////////////////////////////
// ////////// V2c
func snmpConfigV2cTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpConfigV2cConfig {
	var data_list []models.SnmpConfigV2cConfig
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(V2cConfigValue)
		data := models.NewSnmpConfigV2cConfig()
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
func snmpConfigV3NotifyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Snmpv3ConfigNotifyItems {
	var data_list []models.Snmpv3ConfigNotifyItems
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NotifyValue)
		data := models.NewSnmpv3ConfigNotifyItems()
		data.SetName(plan.Name.ValueString())
		data.SetTag(plan.Tag.ValueString())
		data.SetType(models.Snmpv3ConfigNotifyType(plan.Tag.ValueString()))

		data_list = append(data_list, *data)
	}

	return data_list
}

func snmpConfigV3NotifyFilterContentTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Snmpv3ConfigNotifyFilterItemContent {
	var data_list []models.Snmpv3ConfigNotifyFilterItemContent
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(Snmpv3ContentsValue)
		data := models.NewSnmpv3ConfigNotifyFilterItemContent()

		data.SetInclude(plan.Include.ValueBool())
		data.SetOid(plan.Oid.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

func snmpConfigV3NotifyFilterTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Snmpv3ConfigNotifyFilterItem {
	var data_list []models.Snmpv3ConfigNotifyFilterItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NotifyFilterValue)
		data := models.NewSnmpv3ConfigNotifyFilterItem()

		content := snmpConfigV3NotifyFilterContentTerraformToSdk(ctx, diags, plan.Snmpv3Contents)

		data.SetContents(content)
		data.SetProfileName(plan.ProfileName.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

// V3 TARGETS
func snmpConfigV3TargetAddressTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Snmpv3ConfigTargetAddressItem {
	var data_list []models.Snmpv3ConfigTargetAddressItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TargetAddressValue)
		data := models.NewSnmpv3ConfigTargetAddressItem()

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

func snmpConfigV3TargetParametersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Snmpv3ConfigTargetParam {
	var data_list []models.Snmpv3ConfigTargetParam
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TargetParametersValue)
		data := models.NewSnmpv3ConfigTargetParam()

		data.SetMessageProcessingModel(models.Snmpv3ConfigTargetParamMessProcessModel(plan.MessageProcessingModel.ValueString()))
		data.SetName(plan.Name.ValueString())
		data.SetNotifyFilter(plan.NotifyFilter.ValueString())
		data.SetSecurityLevel(models.Snmpv3ConfigTargetParamSecurityLevel(plan.SecurityLevel.ValueString()))
		data.SetSecurityModel(models.Snmpv3ConfigTargetParamSecurityModel(plan.SecurityModel.ValueString()))
		data.SetSecurityName(plan.SecurityName.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

// V3 USM
func snmpConfigV3UsmUsersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpUsmpUser {
	var data_list []models.SnmpUsmpUser
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(Snmpv3UsersValue)
		data := models.NewSnmpUsmpUser()

		data.SetAuthenticationPassword(plan.AuthenticationPassword.ValueString())
		data.SetAuthenticationType(models.SnmpUsmpUserAuthenticationType(plan.AuthenticationType.ValueString()))
		data.SetEncryptionPassword(plan.EncryptionPassword.ValueString())
		data.SetEncryptionType(models.SnmpUsmpUserEncryptionType(plan.EncryptionType.ValueString()))
		data.SetName(plan.Name.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}
func snmpConfigV3UsmTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SnmpUsm {
	data := *models.NewSnmpUsm()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(UsmValue)

		users := snmpConfigV3UsmUsersTerraformToSdk(ctx, diags, plan.Snmpv3Users)

		data.SetEngineType(models.SnmpUsmEngineType(plan.EngineType.ValueString()))
		data.SetEngineId(plan.Engineid.ValueString())
		data.SetUsers(users)

		return data
	}
}

// V3 VACM ACCESS
func snmpConfigV3VacmAccessPrefixTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpVacmAccessItemPrefixListItem {
	var data_list []models.SnmpVacmAccessItemPrefixListItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PrefixListValue)
		data := models.NewSnmpVacmAccessItemPrefixListItem()

		data.SetContextPrefix(plan.ContextPrefix.ValueString())
		data.SetNotifyView(plan.NotifyView.ValueString())
		data.SetReadView(plan.ReadView.ValueString())
		data.SetSecurityLevel(models.SnmpVacmAccessItemPrefixListItemLevel(plan.SecurityLevel.ValueString()))
		data.SetSecurityModel(models.SnmpVacmAccessItemPrefixListItemModel(plan.SecurityModel.ValueString()))
		data.SetContextPrefix(plan.PrefixListType.ValueString())
		data.SetWriteView(plan.WriteView.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}
func snmpConfigV3VacmAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpVacmAccessItem {
	var data_list []models.SnmpVacmAccessItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(AccessValue)
		data := models.NewSnmpVacmAccessItem()

		prefix_list := snmpConfigV3VacmAccessPrefixTerraformToSdk(ctx, diags, plan.PrefixList)

		data.SetGroupName(plan.GroupName.ValueString())
		data.SetPrefixList(prefix_list)

		data_list = append(data_list, *data)
	}

	return data_list
}

// V3 VACM SEC TO GROUP
func snmpConfigV3VacmSecurityToGroupContentTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpVacmSecurityToGroupContentItem {
	var data_list []models.SnmpVacmSecurityToGroupContentItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(Snmpv3VacmContentValue)
		data := models.NewSnmpVacmSecurityToGroupContentItem()

		data.SetGroup(plan.Group.ValueString())
		data.SetSecurityName(plan.SecurityName.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}
func snmpConfigV3VacmSecurityToGroupTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SnmpVacmSecurityToGroup {
	data := *models.NewSnmpVacmSecurityToGroup()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(SecurityToGroupValue)

		content := snmpConfigV3VacmSecurityToGroupContentTerraformToSdk(ctx, diags, plan.Snmpv3VacmContent)

		data.SetSecurityModel(models.SnmpVacmSecurityModel(plan.SecurityModel.ValueString()))
		data.SetContent(content)

		return data
	}
}
func snmpConfigV3VacmTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SnmpVacm {
	data := *models.NewSnmpVacm()
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
func snmpConfigV3TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.Snmpv3Config {
	data := *models.NewSnmpv3Config()
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
func snmpConfigViewsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpConfigView {
	var data_list []models.SnmpConfigView
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ViewsValue)
		data := models.NewSnmpConfigView()

		data.SetInclude(plan.Include.ValueBool())
		data.SetOid(plan.Oid.ValueString())
		data.SetViewName(plan.ViewName.ValueString())

		data_list = append(data_list, *data)
	}

	return data_list
}

// //////////////////////////////////////////////
// ////////// MAIN
func snmpConfigSyslogTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SnmpConfigValue) models.SnmpConfig {

	client_list := snmpConfigClientListTerraformToSdk(ctx, diags, d.ClientList)
	trap_groups := snmpConfigTrapGroupsTerraformToSdk(ctx, diags, d.TrapGroups)
	v2c_config := snmpConfigV2cTerraformToSdk(ctx, diags, d.V2cConfig)
	v3_config := snmpConfigV3TerraformToSdk(ctx, diags, d.V3Config)
	views := snmpConfigViewsTerraformToSdk(ctx, diags, d.Views)

	data := models.NewSnmpConfig()
	data.SetClientList(client_list)
	data.SetContact(d.Contact.ValueString())
	data.SetDescription(d.Description.ValueString())
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetEngineId(models.SnmpConfigEngineId(d.EngineId.ValueString()))
	data.SetLocation(d.Location.ValueString())
	data.SetName(d.Name.ValueString())
	data.SetNetwork(d.Network.ValueString())
	data.SetTrapGroups(trap_groups)
	data.SetV2cConfig(v2c_config)
	data.SetV3Config(v3_config)
	data.SetViews(views)

	return *data
}
