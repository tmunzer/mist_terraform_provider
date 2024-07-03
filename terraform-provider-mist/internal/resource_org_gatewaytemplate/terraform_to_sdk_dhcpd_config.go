package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func dhcpdConfigFixedBindingsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.DhcpdConfigFixedBinding {
	tflog.Debug(ctx, "dhcpdConfigFixedBindingsTerraformToSdk")
	data_map := make(map[string]mistapigo.DhcpdConfigFixedBinding)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(FixedBindingsValue)

		data := *mistapigo.NewDhcpdConfigFixedBinding()
		data.SetIp(plan.Ip.ValueString())
		data.SetName(plan.Name.ValueString())

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.DhcpdConfigOption {
	tflog.Debug(ctx, "dhcpdConfigOptionsTerraformToSdk")
	data_map := make(map[string]mistapigo.DhcpdConfigOption)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OptionsValue)

		data := *mistapigo.NewDhcpdConfigOption()
		data.SetType(mistapigo.DhcpdConfigOptionType(plan.OptionsType.ValueString()))
		data.SetValue(plan.Value.ValueString())

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigVendorOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.DhcpdConfigVendorOption {
	tflog.Debug(ctx, "dhcpdConfigVendorOptionsTerraformToSdk")
	data_map := make(map[string]mistapigo.DhcpdConfigVendorOption)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VendorEncapulatedValue)

		data := *mistapigo.NewDhcpdConfigVendorOption()
		data.SetType(mistapigo.DhcpdConfigVendorOptionType(plan.VendorEncapulatedType.ValueString()))
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

		data := *mistapigo.NewDhcpdConfig()
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
		data.SetType(mistapigo.DhcpdConfigType(plan.Type4.ValueString()))
		data.SetType6(mistapigo.DhcpdConfigType6(plan.Type6.ValueString()))
		data.SetVendorEncapulated(vendor_encapulated)

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpdConfigValue) mistapigo.DhcpdConfigs {
	tflog.Debug(ctx, "dhcpdConfigTerraformToSdk")

	data := *mistapigo.NewDhcpdConfigs()

	data.SetEnabled(d.Enabled.ValueBool())
	data.AdditionalProperties = dhcpdConfigConfigsTerraformToSdk(ctx, diags, d.Config)

	return data
}
