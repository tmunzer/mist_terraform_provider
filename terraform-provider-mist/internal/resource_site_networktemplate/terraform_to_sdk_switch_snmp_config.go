package resource_site_networktemplate

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
		data := models.SnmpConfigClientList{}
		data.Clients = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Clients)
		data.ClientListName = plan.ClientListName.ValueStringPointer()

		data_list = append(data_list, data)
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
		data := models.SnmpConfigTrapGroup{}
		data.Categories = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Categories)
		data.GroupName = plan.GroupName.ValueStringPointer()
		data.Targets = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Targets)
		data.Version = models.ToPointer(models.SnmpConfigTrapVerionEnum(plan.Version.ValueString()))

		data_list = append(data_list, data)
	}

	return data_list
}

// //////////////////////////////////
// ////////// V2c
func snmpConfigV2cTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpConfigV2CConfig {
	var data_list []models.SnmpConfigV2CConfig
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(V2cConfigValue)
		data := models.SnmpConfigV2CConfig{}
		data.Authorization = plan.Authorization.ValueStringPointer()
		data.ClientListName = plan.ClientListName.ValueStringPointer()
		data.CommunityName = plan.CommunityName.ValueStringPointer()
		data.View = plan.View.ValueStringPointer()

		data_list = append(data_list, data)
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
		data := models.Snmpv3ConfigNotifyItems{}
		data.Name = plan.Name.ValueStringPointer()
		data.Tag = plan.Tag.ValueStringPointer()
		data.Type = models.ToPointer(models.Snmpv3ConfigNotifyTypeEnum(plan.Tag.ValueString()))

		data_list = append(data_list, data)
	}

	return data_list
}

func snmpConfigV3NotifyFilterContentTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Snmpv3ConfigNotifyFilterItemContent {
	var data_list []models.Snmpv3ConfigNotifyFilterItemContent
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(Snmpv3ContentsValue)
		data := models.Snmpv3ConfigNotifyFilterItemContent{}

		data.Include = plan.Include.ValueBoolPointer()
		data.Oid = plan.Oid.ValueStringPointer()

		data_list = append(data_list, data)
	}

	return data_list
}

func snmpConfigV3NotifyFilterTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Snmpv3ConfigNotifyFilterItem {
	var data_list []models.Snmpv3ConfigNotifyFilterItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NotifyFilterValue)
		data := models.Snmpv3ConfigNotifyFilterItem{}

		content := snmpConfigV3NotifyFilterContentTerraformToSdk(ctx, diags, plan.Snmpv3Contents)

		data.Contents = content
		data.ProfileName = plan.ProfileName.ValueStringPointer()

		data_list = append(data_list, data)
	}

	return data_list
}

// V3 TARGETS
func snmpConfigV3TargetAddressTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Snmpv3ConfigTargetAddressItem {
	var data_list []models.Snmpv3ConfigTargetAddressItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TargetAddressValue)
		data := models.Snmpv3ConfigTargetAddressItem{}

		data.Address = plan.Address.ValueStringPointer()
		data.AddressMask = plan.AddressMask.ValueStringPointer()
		data.Port = models.ToPointer(int(plan.Port.ValueInt64()))
		data.TagList = plan.TagList.ValueStringPointer()
		data.TargetAddressName = plan.TargetAddressName.ValueStringPointer()
		data.TargetParameters = plan.TargetParameters.ValueStringPointer()

		data_list = append(data_list, data)
	}

	return data_list
}

func snmpConfigV3TargetParametersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Snmpv3ConfigTargetParam {
	var data_list []models.Snmpv3ConfigTargetParam
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TargetParametersValue)
		data := models.Snmpv3ConfigTargetParam{}

		data.MessageProcessingModel = models.ToPointer(models.Snmpv3ConfigTargetParamMessProcessModelEnum(plan.MessageProcessingModel.ValueString()))
		data.Name = plan.Name.ValueStringPointer()
		data.NotifyFilter = plan.NotifyFilter.ValueStringPointer()
		data.SecurityLevel = models.ToPointer(models.Snmpv3ConfigTargetParamSecurityLevelEnum(plan.SecurityLevel.ValueString()))
		data.SecurityModel = models.ToPointer(models.Snmpv3ConfigTargetParamSecurityModelEnum(plan.SecurityModel.ValueString()))
		data.SecurityName = plan.SecurityName.ValueStringPointer()

		data_list = append(data_list, data)
	}

	return data_list
}

// V3 USM
func snmpConfigV3UsmUsersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpUsmpUser {
	var data_list []models.SnmpUsmpUser
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(Snmpv3UsersValue)
		data := models.SnmpUsmpUser{}

		data.AuthenticationPassword = plan.AuthenticationPassword.ValueStringPointer()
		data.AuthenticationType = models.ToPointer(models.SnmpUsmpUserAuthenticationTypeEnum(plan.AuthenticationType.ValueString()))
		data.EncryptionPassword = plan.EncryptionPassword.ValueStringPointer()
		data.EncryptionType = models.ToPointer(models.SnmpUsmpUserEncryptionTypeEnum(plan.EncryptionType.ValueString()))
		data.Name = plan.Name.ValueStringPointer()

		data_list = append(data_list, data)
	}

	return data_list
}
func snmpConfigV3UsmTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SnmpUsm {
	data := models.SnmpUsm{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(UsmValue)

		data.EngineType = models.ToPointer(models.SnmpUsmEngineTypeEnum(plan.EngineType.ValueString()))
		data.EngineId = plan.Engineid.ValueStringPointer()
		data.Users = snmpConfigV3UsmUsersTerraformToSdk(ctx, diags, plan.Snmpv3Users)

		return &data
	}
}

// V3 VACM ACCESS
func snmpConfigV3VacmAccessPrefixTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpVacmAccessItemPrefixListItem {
	var data_list []models.SnmpVacmAccessItemPrefixListItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PrefixListValue)
		data := models.SnmpVacmAccessItemPrefixListItem{}

		data.ContextPrefix = plan.ContextPrefix.ValueStringPointer()
		data.NotifyView = plan.NotifyView.ValueStringPointer()
		data.ReadView = plan.ReadView.ValueStringPointer()
		data.SecurityLevel = models.ToPointer(models.SnmpVacmAccessItemPrefixListItemLevelEnum(plan.SecurityLevel.ValueString()))
		data.SecurityModel = models.ToPointer(models.SnmpVacmAccessItemPrefixListItemModelEnum(plan.SecurityModel.ValueString()))
		data.ContextPrefix = plan.PrefixListType.ValueStringPointer()
		data.WriteView = plan.WriteView.ValueStringPointer()

		data_list = append(data_list, data)
	}

	return data_list
}
func snmpConfigV3VacmAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpVacmAccessItem {
	var data_list []models.SnmpVacmAccessItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(AccessValue)
		data := models.SnmpVacmAccessItem{}

		prefix_list := snmpConfigV3VacmAccessPrefixTerraformToSdk(ctx, diags, plan.PrefixList)

		data.GroupName = plan.GroupName.ValueStringPointer()
		data.PrefixList = prefix_list

		data_list = append(data_list, data)
	}

	return data_list
}

// V3 VACM SEC TO GROUP
func snmpConfigV3VacmSecurityToGroupContentTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpVacmSecurityToGroupContentItem {
	var data_list []models.SnmpVacmSecurityToGroupContentItem
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(Snmpv3VacmContentValue)
		data := models.SnmpVacmSecurityToGroupContentItem{}

		data.Group = plan.Group.ValueStringPointer()
		data.SecurityName = plan.SecurityName.ValueStringPointer()

		data_list = append(data_list, data)
	}

	return data_list
}
func snmpConfigV3VacmSecurityToGroupTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SnmpVacmSecurityToGroup {
	data := models.SnmpVacmSecurityToGroup{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(SecurityToGroupValue)

		data.SecurityModel = models.ToPointer(models.SnmpVacmSecurityModelEnum(plan.SecurityModel.ValueString()))
		data.Content = snmpConfigV3VacmSecurityToGroupContentTerraformToSdk(ctx, diags, plan.Snmpv3VacmContent)

		return &data
	}
}
func snmpConfigV3VacmTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SnmpVacm {
	data := models.SnmpVacm{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(VacmValue)

		data.Access = snmpConfigV3VacmAccessTerraformToSdk(ctx, diags, plan.Access)
		data.SecurityToGroup = snmpConfigV3VacmSecurityToGroupTerraformToSdk(ctx, diags, plan.SecurityToGroup)

		return &data
	}
}

// V3 MAIN
func snmpConfigV3TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.Snmpv3Config {
	data := models.Snmpv3Config{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(V3ConfigValue)

		data.Notify = snmpConfigV3NotifyTerraformToSdk(ctx, diags, plan.Notify)
		data.NotifyFilter = snmpConfigV3NotifyFilterTerraformToSdk(ctx, diags, plan.NotifyFilter)
		data.TargetAddress = snmpConfigV3TargetAddressTerraformToSdk(ctx, diags, plan.TargetAddress)
		data.TargetParameters = snmpConfigV3TargetParametersTerraformToSdk(ctx, diags, plan.TargetParameters)
		data.Usm = snmpConfigV3UsmTerraformToSdk(ctx, diags, plan.Usm)
		data.Vacm = snmpConfigV3VacmTerraformToSdk(ctx, diags, plan.Vacm)
		return &data
	}
}

// //////////////////////////////////////////////
// ////////// VIEWS
func snmpConfigViewsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpConfigView {
	var data_list []models.SnmpConfigView
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ViewsValue)
		data := models.SnmpConfigView{}

		data.Include = plan.Include.ValueBoolPointer()
		data.Oid = plan.Oid.ValueStringPointer()
		data.ViewName = plan.ViewName.ValueStringPointer()

		data_list = append(data_list, data)
	}

	return data_list
}

// //////////////////////////////////////////////
// ////////// MAIN
func snmpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SnmpConfigValue) *models.SnmpConfig {

	client_list := snmpConfigClientListTerraformToSdk(ctx, diags, d.ClientList)
	trap_groups := snmpConfigTrapGroupsTerraformToSdk(ctx, diags, d.TrapGroups)
	v2c_config := snmpConfigV2cTerraformToSdk(ctx, diags, d.V2cConfig)
	v3_config := snmpConfigV3TerraformToSdk(ctx, diags, d.V3Config)
	views := snmpConfigViewsTerraformToSdk(ctx, diags, d.Views)

	data := models.SnmpConfig{}
	data.ClientList = client_list
	data.Contact = d.Contact.ValueStringPointer()
	data.Description = d.Description.ValueStringPointer()
	data.Enabled = d.Enabled.ValueBoolPointer()
	data.EngineId = models.ToPointer(models.SnmpConfigEngineIdEnum(d.EngineId.ValueString()))
	data.Location = d.Location.ValueStringPointer()
	data.Name = d.Name.ValueStringPointer()
	data.Network = d.Network.ValueStringPointer()
	data.TrapGroups = trap_groups
	data.V2cConfig = v2c_config
	data.V3Config = v3_config
	data.Views = views

	return &data
}
