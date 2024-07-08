package resource_org_networktemplate

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
		if !plan.Clients.IsNull() && !plan.Clients.IsUnknown() {
			data.Clients = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Clients)
		}
		if !plan.ClientListName.IsNull() && !plan.ClientListName.IsUnknown() {
			data.ClientListName = plan.ClientListName.ValueStringPointer()
		}

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
		if !plan.Categories.IsNull() && !plan.Categories.IsUnknown() {
			data.Categories = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Categories)
		}
		if !plan.GroupName.IsNull() && !plan.GroupName.IsUnknown() {
			data.GroupName = plan.GroupName.ValueStringPointer()
		}
		if !plan.Targets.IsNull() && !plan.Targets.IsUnknown() {
			data.Targets = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Targets)
		}
		if !plan.Version.IsNull() && !plan.Version.IsUnknown() {
			data.Version = models.ToPointer(models.SnmpConfigTrapVerionEnum(plan.Version.ValueString()))
		}

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
		if !plan.Authorization.IsNull() && !plan.Authorization.IsUnknown() {
			data.Authorization = plan.Authorization.ValueStringPointer()
		}
		if !plan.ClientListName.IsNull() && !plan.ClientListName.IsUnknown() {
			data.ClientListName = plan.ClientListName.ValueStringPointer()
		}
		if !plan.CommunityName.IsNull() && !plan.CommunityName.IsUnknown() {
			data.CommunityName = plan.CommunityName.ValueStringPointer()
		}
		if !plan.View.IsNull() && !plan.View.IsUnknown() {
			data.View = plan.View.ValueStringPointer()
		}

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
		if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
			data.Name = plan.Name.ValueStringPointer()
		}
		if !plan.Tag.IsNull() && !plan.Tag.IsUnknown() {
			data.Tag = plan.Tag.ValueStringPointer()
		}
		if !plan.Tag.IsNull() && !plan.Tag.IsUnknown() {
			data.Type = models.ToPointer(models.Snmpv3ConfigNotifyTypeEnum(plan.Tag.ValueString()))
		}

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

		if !plan.Include.IsNull() && !plan.Include.IsUnknown() {
			data.Include = plan.Include.ValueBoolPointer()
		}
		if !plan.Oid.IsNull() && !plan.Oid.IsUnknown() {
			data.Oid = plan.Oid.ValueStringPointer()
		}

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

		if !plan.Snmpv3Contents.IsNull() && !plan.Snmpv3Contents.IsUnknown() {
			data.Contents = snmpConfigV3NotifyFilterContentTerraformToSdk(ctx, diags, plan.Snmpv3Contents)
		}
		if !plan.ProfileName.IsNull() && !plan.ProfileName.IsUnknown() {
			data.ProfileName = plan.ProfileName.ValueStringPointer()
		}

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

		if !plan.Address.IsNull() && !plan.Address.IsUnknown() {
			data.Address = plan.Address.ValueStringPointer()
		}
		if !plan.AddressMask.IsNull() && !plan.AddressMask.IsUnknown() {
			data.AddressMask = plan.AddressMask.ValueStringPointer()
		}
		if !plan.Port.IsNull() && !plan.Port.IsUnknown() {
			data.Port = models.ToPointer(int(plan.Port.ValueInt64()))
		}
		if !plan.TagList.IsNull() && !plan.TagList.IsUnknown() {
			data.TagList = plan.TagList.ValueStringPointer()
		}
		if !plan.TargetAddressName.IsNull() && !plan.TargetAddressName.IsUnknown() {
			data.TargetAddressName = plan.TargetAddressName.ValueStringPointer()
		}
		if !plan.TargetParameters.IsNull() && !plan.TargetParameters.IsUnknown() {
			data.TargetParameters = plan.TargetParameters.ValueStringPointer()
		}

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

		if !plan.MessageProcessingModel.IsNull() && !plan.MessageProcessingModel.IsUnknown() {
			data.MessageProcessingModel = models.ToPointer(models.Snmpv3ConfigTargetParamMessProcessModelEnum(plan.MessageProcessingModel.ValueString()))
		}
		if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
			data.Name = plan.Name.ValueStringPointer()
		}
		if !plan.NotifyFilter.IsNull() && !plan.NotifyFilter.IsUnknown() {
			data.NotifyFilter = plan.NotifyFilter.ValueStringPointer()
		}
		if !plan.SecurityLevel.IsNull() && !plan.SecurityLevel.IsUnknown() {
			data.SecurityLevel = models.ToPointer(models.Snmpv3ConfigTargetParamSecurityLevelEnum(plan.SecurityLevel.ValueString()))
		}
		if !plan.SecurityModel.IsNull() && !plan.SecurityModel.IsUnknown() {
			data.SecurityModel = models.ToPointer(models.Snmpv3ConfigTargetParamSecurityModelEnum(plan.SecurityModel.ValueString()))
		}
		if !plan.SecurityName.IsNull() && !plan.SecurityName.IsUnknown() {
			data.SecurityName = plan.SecurityName.ValueStringPointer()
		}

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

		if !plan.AuthenticationPassword.IsNull() && !plan.AuthenticationPassword.IsUnknown() {
			data.AuthenticationPassword = plan.AuthenticationPassword.ValueStringPointer()
		}
		if !plan.AuthenticationType.IsNull() && !plan.AuthenticationType.IsUnknown() {
			data.AuthenticationType = models.ToPointer(models.SnmpUsmpUserAuthenticationTypeEnum(plan.AuthenticationType.ValueString()))
		}
		if !plan.EncryptionPassword.IsNull() && !plan.EncryptionPassword.IsUnknown() {
			data.EncryptionPassword = plan.EncryptionPassword.ValueStringPointer()
		}
		if !plan.EncryptionType.IsNull() && !plan.EncryptionType.IsUnknown() {
			data.EncryptionType = models.ToPointer(models.SnmpUsmpUserEncryptionTypeEnum(plan.EncryptionType.ValueString()))
		}
		if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
			data.Name = plan.Name.ValueStringPointer()
		}

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

		if !plan.EngineType.IsNull() && !plan.EngineType.IsUnknown() {
			data.EngineType = models.ToPointer(models.SnmpUsmEngineTypeEnum(plan.EngineType.ValueString()))
		}
		if !plan.Engineid.IsNull() && !plan.Engineid.IsUnknown() {
			data.EngineId = plan.Engineid.ValueStringPointer()
		}
		if !plan.Snmpv3Users.IsNull() && !plan.Snmpv3Users.IsUnknown() {
			data.Users = snmpConfigV3UsmUsersTerraformToSdk(ctx, diags, plan.Snmpv3Users)
		}

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

		if !plan.ContextPrefix.IsNull() && !plan.ContextPrefix.IsUnknown() {
			data.ContextPrefix = plan.ContextPrefix.ValueStringPointer()
		}
		if !plan.NotifyView.IsNull() && !plan.NotifyView.IsUnknown() {
			data.NotifyView = plan.NotifyView.ValueStringPointer()
		}
		if !plan.ReadView.IsNull() && !plan.ReadView.IsUnknown() {
			data.ReadView = plan.ReadView.ValueStringPointer()
		}
		if !plan.SecurityLevel.IsNull() && !plan.SecurityLevel.IsUnknown() {
			data.SecurityLevel = models.ToPointer(models.SnmpVacmAccessItemPrefixListItemLevelEnum(plan.SecurityLevel.ValueString()))
		}
		if !plan.SecurityModel.IsNull() && !plan.SecurityModel.IsUnknown() {
			data.SecurityModel = models.ToPointer(models.SnmpVacmAccessItemPrefixListItemModelEnum(plan.SecurityModel.ValueString()))
		}
		if !plan.ContextPrefix.IsNull() && !plan.ContextPrefix.IsUnknown() {
			data.ContextPrefix = plan.PrefixListType.ValueStringPointer()
		}
		if !plan.WriteView.IsNull() && !plan.WriteView.IsUnknown() {
			data.WriteView = plan.WriteView.ValueStringPointer()
		}

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

		if !plan.GroupName.IsNull() && !plan.GroupName.IsUnknown() {
			data.GroupName = plan.GroupName.ValueStringPointer()
		}
		if !plan.PrefixList.IsNull() && !plan.PrefixList.IsUnknown() {
			data.PrefixList = prefix_list
		}

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

		if !plan.Group.IsNull() && !plan.Group.IsUnknown() {
			data.Group = plan.Group.ValueStringPointer()
		}
		if !plan.SecurityName.IsNull() && !plan.SecurityName.IsUnknown() {
			data.SecurityName = plan.SecurityName.ValueStringPointer()
		}

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

		if !plan.SecurityModel.IsNull() && !plan.SecurityModel.IsUnknown() {
			data.SecurityModel = models.ToPointer(models.SnmpVacmSecurityModelEnum(plan.SecurityModel.ValueString()))
		}
		if !plan.Snmpv3VacmContent.IsNull() && !plan.Snmpv3VacmContent.IsUnknown() {
			data.Content = snmpConfigV3VacmSecurityToGroupContentTerraformToSdk(ctx, diags, plan.Snmpv3VacmContent)
		}

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

		if !plan.Access.IsNull() && !plan.Access.IsUnknown() {
			data.Access = snmpConfigV3VacmAccessTerraformToSdk(ctx, diags, plan.Access)
		}
		if !plan.SecurityToGroup.IsNull() && !plan.SecurityToGroup.IsUnknown() {
			data.SecurityToGroup = snmpConfigV3VacmSecurityToGroupTerraformToSdk(ctx, diags, plan.SecurityToGroup)
		}

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

		if !plan.Notify.IsNull() && !plan.Notify.IsUnknown() {
			data.Notify = snmpConfigV3NotifyTerraformToSdk(ctx, diags, plan.Notify)
		}
		if !plan.NotifyFilter.IsNull() && !plan.NotifyFilter.IsUnknown() {
			data.NotifyFilter = snmpConfigV3NotifyFilterTerraformToSdk(ctx, diags, plan.NotifyFilter)
		}
		if !plan.TargetAddress.IsNull() && !plan.TargetAddress.IsUnknown() {
			data.TargetAddress = snmpConfigV3TargetAddressTerraformToSdk(ctx, diags, plan.TargetAddress)
		}
		if !plan.TargetParameters.IsNull() && !plan.TargetParameters.IsUnknown() {
			data.TargetParameters = snmpConfigV3TargetParametersTerraformToSdk(ctx, diags, plan.TargetParameters)
		}
		if !plan.Usm.IsNull() && !plan.Usm.IsUnknown() {
			data.Usm = snmpConfigV3UsmTerraformToSdk(ctx, diags, plan.Usm)
		}
		if !plan.Vacm.IsNull() && !plan.Vacm.IsUnknown() {
			data.Vacm = snmpConfigV3VacmTerraformToSdk(ctx, diags, plan.Vacm)
		}
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

		if !plan.Include.IsNull() && !plan.Include.IsUnknown() {
			data.Include = plan.Include.ValueBoolPointer()
		}
		if !plan.Oid.IsNull() && !plan.Oid.IsUnknown() {
			data.Oid = plan.Oid.ValueStringPointer()
		}
		if !plan.ViewName.IsNull() && !plan.ViewName.IsUnknown() {
			data.ViewName = plan.ViewName.ValueStringPointer()
		}

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
	if !d.ClientList.IsNull() && !d.ClientList.IsUnknown() {
		data.ClientList = client_list
	}
	if !d.Contact.IsNull() && !d.Contact.IsUnknown() {
		data.Contact = d.Contact.ValueStringPointer()
	}
	if !d.Description.IsNull() && !d.Description.IsUnknown() {
		data.Description = d.Description.ValueStringPointer()
	}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.EngineId.IsNull() && !d.EngineId.IsUnknown() {
		data.EngineId = models.ToPointer(models.SnmpConfigEngineIdEnum(d.EngineId.ValueString()))
	}
	if !d.Location.IsNull() && !d.Location.IsUnknown() {
		data.Location = d.Location.ValueStringPointer()
	}
	if !d.Name.IsNull() && !d.Name.IsUnknown() {
		data.Name = d.Name.ValueStringPointer()
	}
	if !d.Network.IsNull() && !d.Network.IsUnknown() {
		data.Network = d.Network.ValueStringPointer()
	}
	if !d.TrapGroups.IsNull() && !d.TrapGroups.IsUnknown() {
		data.TrapGroups = trap_groups
	}
	if !d.V2cConfig.IsNull() && !d.V2cConfig.IsUnknown() {
		data.V2cConfig = v2c_config
	}
	if !d.V3Config.IsNull() && !d.V3Config.IsUnknown() {
		data.V3Config = v3_config
	}
	if !d.Views.IsNull() && !d.Views.IsUnknown() {
		data.Views = views
	}

	return &data
}
