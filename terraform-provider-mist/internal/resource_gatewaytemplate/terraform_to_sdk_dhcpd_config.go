package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func dhcpdConfigFixedBindingsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.DhcpdConfigFixedBinding {
	tflog.Debug(ctx, "dhcpdConfigFixedBindingsTerraformToSdk")
	data_map := make(map[string]mistsdkgo.DhcpdConfigFixedBinding)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(FixedBindingsValue)

		data := *mistsdkgo.NewDhcpdConfigFixedBinding()
		data.SetIp(plan.Ip.ValueString())
		data.SetName(plan.Name.ValueString())

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.DhcpdConfigOption {
	tflog.Debug(ctx, "dhcpdConfigOptionsTerraformToSdk")
	data_map := make(map[string]mistsdkgo.DhcpdConfigOption)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OptionsValue)

		data := *mistsdkgo.NewDhcpdConfigOption()
		data.SetType(mistsdkgo.DhcpdConfigOptionType(plan.OptionsType.ValueString()))
		data.SetValue(plan.Value.ValueString())

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigVendorOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.DhcpdConfigVendorOption {
	tflog.Debug(ctx, "dhcpdConfigVendorOptionsTerraformToSdk")
	data_map := make(map[string]mistsdkgo.DhcpdConfigVendorOption)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VendorEncapulatedValue)

		data := *mistsdkgo.NewDhcpdConfigVendorOption()
		data.SetType(mistsdkgo.DhcpdConfigVendorOptionType(plan.VendorEncapulatedType.ValueString()))
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

		data := *mistsdkgo.NewDhcpdConfig()
		data.SetDnsServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers))
		data.SetDnsSuffix(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix))
		data.SetFixedBindings(fixed_bindings)
		data.SetGateway(plan.Gateway.ValueString())
		data.SetIpEnd(plan.IpEnd4.ValueString())
		data.SetIpEnd6(plan.IpEnd6.ValueString())
		data.SetIpStart(plan.IpStart4.ValueString())
		data.SetIpStart6(plan.IpStart6.ValueString())
		data.SetLeaseTime(int32(plan.LeaseTime.ValueInt64()))
		data.SetOptions(options)
		data.SetServerIdOverride(plan.ServerIdOverride.ValueBool())
		data.SetServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Servers4))
		data.SetServers6(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Servers6))
		data.SetType(mistsdkgo.DhcpdConfigType(plan.Type4.ValueString()))
		data.SetType6(mistsdkgo.DhcpdConfigType6(plan.Type6.ValueString()))
		data.SetVendorEncapulated(vendor_encapulated)

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpdConfigValue) mistsdkgo.DhcpdConfigs {
	tflog.Debug(ctx, "dhcpdConfigTerraformToSdk")

	data := *mistsdkgo.NewDhcpdConfigs()

	data.SetEnabled(d.Enabled.ValueBool())
	data.AdditionalProperties = dhcpdConfigConfigsTerraformToSdk(ctx, diags, d.Config)

	return data
}
