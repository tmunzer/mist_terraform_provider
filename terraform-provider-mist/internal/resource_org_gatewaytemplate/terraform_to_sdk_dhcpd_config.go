package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func dhcpdConfigFixedBindingsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.DhcpdConfigFixedBinding {
	tflog.Debug(ctx, "dhcpdConfigFixedBindingsTerraformToSdk")
	data_map := make(map[string]models.DhcpdConfigFixedBinding)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(FixedBindingsValue)

		data := *&models.DhcpdConfigFixedBinding{}
		data.Ip = models.ToPointer(plan.Ip.ValueString())
		data.Name = models.ToPointer(plan.Name.ValueString())

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.DhcpdConfigOption {
	tflog.Debug(ctx, "dhcpdConfigOptionsTerraformToSdk")
	data_map := make(map[string]models.DhcpdConfigOption)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OptionsValue)

		data := models.DhcpdConfigOption{}
		data.Type = models.ToPointer(models.DhcpdConfigOptionTypeEnum(plan.OptionsType.ValueString()))
		data.Value = models.ToPointer(plan.Value.ValueString())

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigVendorOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.DhcpdConfigVendorOption {
	tflog.Debug(ctx, "dhcpdConfigVendorOptionsTerraformToSdk")
	data_map := make(map[string]models.DhcpdConfigVendorOption)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VendorEncapulatedValue)

		data := models.DhcpdConfigVendorOption{}
		data.Type = models.ToPointer(models.DhcpdConfigVendorOptionTypeEnum(plan.VendorEncapulatedType.ValueString()))
		data.Value = models.ToPointer(plan.Value.ValueString())

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]interface{} {
	tflog.Debug(ctx, "dhcpdConfigConfigsTerraformToSdk")
	data_map := make(map[string]interface{})
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ConfigValue)

		fixed_bindings := dhcpdConfigFixedBindingsTerraformToSdk(ctx, diags, plan.FixedBindings)
		options := dhcpdConfigOptionsTerraformToSdk(ctx, diags, plan.Options)
		vendor_encapulated := dhcpdConfigVendorOptionsTerraformToSdk(ctx, diags, plan.VendorEncapulated)

		data := models.DhcpdConfig{}
		data.DnsServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)
		data.DnsSuffix = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
		data.FixedBindings = fixed_bindings
		data.Gateway = models.ToPointer(plan.Gateway.ValueString())
		data.IpEnd = models.ToPointer(plan.IpEnd4.ValueString())
		data.IpEnd6 = models.ToPointer(plan.IpEnd6.ValueString())
		data.IpStart = models.ToPointer(plan.IpStart4.ValueString())
		data.IpStart6 = models.ToPointer(plan.IpStart6.ValueString())
		data.LeaseTime = models.ToPointer(int(plan.LeaseTime.ValueInt64()))
		data.Options = options
		data.ServerIdOverride = models.ToPointer(plan.ServerIdOverride.ValueBool())
		data.Servers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Servers4)
		data.Servers6 = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Servers6)
		data.Type = models.ToPointer(models.DhcpdConfigTypeEnum(plan.Type4.ValueString()))
		data.Type6 = models.ToPointer(models.DhcpdConfigTypeEnum(plan.Type6.ValueString()))
		data.VendorEncapulated = vendor_encapulated

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpdConfigValue) models.DhcpdConfigs {
	tflog.Debug(ctx, "dhcpdConfigTerraformToSdk")

	data := models.DhcpdConfigs{}

	data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	data.AdditionalProperties = dhcpdConfigConfigsTerraformToSdk(ctx, diags, d.Config)

	return data
}
