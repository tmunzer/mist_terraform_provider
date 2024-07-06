package resource_device_switch

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

		data := *models.NewDhcpdConfigFixedBinding()
		data.SetIp(plan.Ip.ValueString())
		data.SetName(plan.Name.ValueString())

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

		data := *models.NewDhcpdConfigOption()
		data.SetType(models.DhcpdConfigOptionType(plan.OptionsType.ValueString()))
		data.SetValue(plan.Value.ValueString())

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

		data := *models.NewDhcpdConfigVendorOption()
		data.SetType(models.DhcpdConfigVendorOptionType(plan.VendorEncapulatedType.ValueString()))
		data.SetValue(plan.Value.ValueString())

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

		data := *models.NewDhcpdConfig()
		data.SetDnsServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers))
		data.SetDnsSuffix(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix))
		data.SetFixedBindings(fixed_bindings)
		data.SetGateway(plan.Gateway.ValueString())
		data.SetIpEnd(plan.IpEnd.ValueString())
		data.SetIpEnd6(plan.IpEnd6.ValueString())
		data.SetIpStart(plan.IpStart.ValueString())
		data.SetIpStart6(plan.IpStart6.ValueString())
		data.SetLeaseTime(int32(plan.LeaseTime.ValueInt64()))
		data.SetOptions(options)
		data.SetServerIdOverride(plan.ServerIdOverride.ValueBool())
		data.SetServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Servers))
		data.SetServers6(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Servers6))
		data.SetType(models.DhcpdConfigType(plan.ConfigType.ValueString()))
		data.SetType6(models.DhcpdConfigType6(plan.Type6.ValueString()))
		data.SetVendorEncapulated(vendor_encapulated)

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpdConfigValue) models.DhcpdConfigs {
	tflog.Debug(ctx, "dhcpdConfigTerraformToSdk")

	data := *models.NewDhcpdConfigs()

	data.SetEnabled(d.Enabled.ValueBool())
	data.AdditionalProperties = dhcpdConfigConfigsTerraformToSdk(ctx, diags, d.Config)

	return data
}
