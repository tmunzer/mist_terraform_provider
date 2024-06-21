package provider

import (
	"context"
	"fmt"
	"strings"
	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	mist_error "terraform-provider-mistapi/internal/provider/utils"
	"terraform-provider-mistapi/internal/resource_networktemplate"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &networktemplateResource{}
	_ resource.ResourceWithConfigure = &networktemplateResource{}
)

func NewNetworkTemplate() resource.Resource {
	return &networktemplateResource{}
}

type networktemplateResource struct {
	client *mistsdkgo.APIClient
}

func (r *networktemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist NetworkTemplate client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*mistsdkgo.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistsdkgo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *networktemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networktemplate"
}

func (r *networktemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_networktemplate.NetworktemplateResourceSchema(ctx)
}

func (r *networktemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting NetworkTemplate Create")
	var plan resource_networktemplate.NetworktemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	networktemplate, orgId := processNetworktemplatePlan(ctx, &plan)
	data, _, err := r.client.OrgsNetworkTemplatesAPI.CreateOrgNetworkTemplate(ctx, orgId).NetworkTemplate(networktemplate).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating NetworkTemplate",
			"Could not create NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state := processNetworktemplateData(ctx, data)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *networktemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_networktemplate.NetworktemplateModel
	tflog.Info(ctx, "Starting NetworkTemplate Read: networktemplate_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsNetworkTemplatesAPI.GetOrgNetworkTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting NetworkTemplate",
			"Could not get NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}
	state = processNetworktemplateData(ctx, data)
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *networktemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_networktemplate.NetworktemplateModel
	tflog.Info(ctx, "Starting NetworkTemplate Update")

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	networktemplateId := state.Id.ValueString()
	networktemplate, OrgId := processNetworktemplatePlan(ctx, &plan)
	tflog.Info(ctx, "Starting NetworkTemplate Update for NetworkTemplate "+networktemplateId)
	data, _, err := r.client.OrgsNetworkTemplatesAPI.UpdateOrgNetworkTemplates(ctx, OrgId, networktemplateId).NetworkTemplate(networktemplate).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating NetworkTemplate",
			"Could not create NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	test := processNetworktemplateData(ctx, data)
	diags = resp.State.Set(ctx, test)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *networktemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_networktemplate.NetworktemplateModel
	tflog.Info(ctx, "Starting NetworkTemplate Delete: networktemplate_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.OrgsNetworkTemplatesAPI.DeleteOrgNetworkTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating NetworkTemplate",
			"Could not delete NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}
}

func processNetworktemplateData(ctx context.Context, data *mistsdkgo.NetworkTemplate) resource_networktemplate.NetworktemplateModel {
	var state resource_networktemplate.NetworktemplateModel

	tflog.Debug(ctx, "fgdiosgfdjriogjfdiosgjdfio0", map[string]interface{}{})
	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())

	addictionalConfigCmds := convertDataListOfString(ctx, data.GetAdditionalConfigCmds())
	state.AdditionalConfigCmds = addictionalConfigCmds

	dnsServers := convertDataListOfString(ctx, data.GetDnsServers())
	state.DnsServers = dnsServers

	dnsSuffix := convertDataListOfString(ctx, data.GetDnsSuffix())
	state.DnsSuffix = dnsSuffix

	ntpServers := convertDataListOfString(ctx, data.GetNtpServers())
	state.NtpServers = ntpServers

	tflog.Debug(ctx, "fgdiosgfdjriogjfdiosgjdfio1", map[string]interface{}{})
	// NETWORKS
	var net_type = map[string]attr.Type{
		"vlan_id": basetypes.Int64Type{},
		"subnet":  basetypes.StringType{},
	}
	var l attr.Type = resource_networktemplate.NetworksValue{}.Type(ctx)
	net_value := make(map[string]attr.Value)
	for k, v := range data.GetNetworks() {

		var net_state = map[string]attr.Value{
			"vlan_id": types.Int64Value(int64(v.GetVlanId())),
			"subnet":  types.StringValue(v.GetSubnet()),
		}
		n, _ := resource_networktemplate.NewNetworksValue(net_type, net_state)
		net_value[k] = n
	}
	state.Networks, _ = types.MapValueFrom(ctx, l, net_value)

	tflog.Debug(ctx, "fgdiosgfdjriogjfdiosgjdfio2", map[string]interface{}{})

	////////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	// //RADIUS

	// // RADIUS Acct
	var rc_srv_state_type = map[string]attr.Type{
		"host":            basetypes.StringType{},
		"port":            basetypes.Int64Type{},
		"secret":          basetypes.StringType{},
		"keywrap_enabled": basetypes.BoolType{},
		"keywrap_format":  basetypes.StringType{},
		"keywrap_kek":     basetypes.StringType{},
		"keywrap_mack":    basetypes.StringType{},
	}

	// // RADIUS Acct
	// var rc_acct_obj []resource_networktemplate.RadiusAcctServersValue
	// for _, acct_data := range data.RadiusConfig.GetAcctServers() {
	// 	var rc_acct_state_value = map[string]attr.Value{
	// 		"host":            types.StringValue(acct_data.GetHost()),
	// 		"port":            types.Int64Value(int64(acct_data.GetPort())),
	// 		"secret":          types.StringValue(acct_data.GetSecret()),
	// 		"keywrap_enabled": types.BoolValue(acct_data.GetKeywrapEnabled()),
	// 		"keywrap_format":  types.StringValue(acct_data.GetKeywrapFormat()),
	// 		"keywrap_kek":     types.StringValue(acct_data.GetKeywrapKek()),
	// 		"keywrap_mack":    types.StringValue(acct_data.GetKeywrapMack()),
	// 	}
	// 	test, e := resource_networktemplate.NewRadiusAcctServersValue(resource_networktemplate.RadiusAuthServersValue{}.AttributeTypes(ctx), rc_acct_state_value)
	// 	if e != nil {
	// 		mist_error.LogError(ctx, e, "NewRadiusAcctServersValue")
	// 	}
	// 	rc_acct_obj = append(rc_acct_obj, test)
	// }
	// rc_acct_state, e := types.ListValueFrom(ctx, resource_networktemplate.RadiusAcctServersType{}, rc_acct_obj)
	// if e != nil {
	// 	mist_error.LogError(ctx, e, "NewRadiusAcctServersValue.ListValueFrom")
	// }
	// // RADIUS Acct)
	var rc_acct_items []attr.Value
	var rc_acct_type attr.Type = resource_networktemplate.RadiusAcctServersType{}
	for _, item := range data.RadiusConfig.GetAcctServers() {
		var rc_acct_state_value = resource_networktemplate.RadiusAcctServersValue{
			Host:           types.StringValue(item.GetHost()),
			Port:           types.Int64Value(int64(item.GetPort())),
			Secret:         types.StringValue(item.GetSecret()),
			KeywrapEnabled: types.BoolValue(item.GetKeywrapEnabled()),
			KeywrapFormat:  types.StringValue(item.GetKeywrapFormat()),
			KeywrapKek:     types.StringValue(item.GetKeywrapKek()),
			KeywrapMack:    types.StringValue(item.GetKeywrapMack()),
		}
		rc_acct_items = append(rc_acct_items, rc_acct_state_value)

	}
	rc_acct_state, e := types.ListValueFrom(ctx, rc_acct_type, rc_acct_items)
	if e != nil {
		mist_error.LogError(ctx, e, "NewRadiusAcctServersValue.ListValueFrom")
	}
	// // RADIUS Auth
	var rc_auth_items []attr.Value
	rc_auth_type := resource_networktemplate.RadiusAuthServersType{}
	for _, item := range data.RadiusConfig.GetAuthServers() {
		rc_auth_state_value := make(map[string]attr.Value)
		rc_auth_state_value["host"] = types.StringValue(item.GetHost())
		rc_auth_state_value["port"] = types.Int64Value(int64(item.GetPort()))
		rc_auth_state_value["secret"] = types.StringValue(item.GetSecret())
		rc_auth_state_value["keywrap_enabled"] = types.BoolValue(item.GetKeywrapEnabled())
		rc_auth_state_value["keywrap_format"] = types.StringValue(item.GetKeywrapFormat())
		rc_auth_state_value["keywrap_kek"] = types.StringValue(item.GetKeywrapKek())
		rc_auth_state_value["keywrap_mack"] = types.StringValue(item.GetKeywrapMack())
		test, _ := resource_networktemplate.NewRadiusAuthServersValue(rc_srv_state_type, rc_auth_state_value)
		rc_auth_items = append(rc_auth_items, test)

	}
	rc_auth_state, e := types.ListValue(rc_auth_type, rc_auth_items)
	if e != nil {
		mist_error.LogError(ctx, e, "NewRadiusAuthServersValue.ListValueFrom")
	}
	// // RADIUS Auth
	var rc_state_value = map[string]attr.Value{
		"acct_interim_interval": types.Int64Value(int64(data.RadiusConfig.GetAcctInterimInterval())),
		"auth_servers_retries":  types.Int64Value(int64(data.RadiusConfig.GetAuthServersRetries())),
		"auth_servers_timeout":  types.Int64Value(int64(data.RadiusConfig.GetAuthServersTimeout())),
		"coa_enabled":           types.BoolValue(data.RadiusConfig.GetCoaEnabled()),
		"coa_port":              types.Int64Value(int64(data.RadiusConfig.GetCoaPort())),
		"network":               types.StringValue(data.RadiusConfig.GetNetwork()),
		"source_ip":             types.StringValue(data.RadiusConfig.GetSourceIp()),
		"radius_acct_servers":   rc_acct_state,
		"radius_auth_servers":   rc_auth_state,
	}

	rc_state, e := resource_networktemplate.NewRadiusConfigValue(resource_networktemplate.RadiusConfigValue{}.AttributeTypes(ctx), rc_state_value)
	if e != nil {
		mist_error.LogError(ctx, e, "NewRadiusConfigValue")
	}
	state.RadiusConfig = rc_state

	return state
}

func processNetworktemplatePlan(ctx context.Context, plan *resource_networktemplate.NetworktemplateModel) (mistsdkgo.NetworkTemplate, string) {
	data := *mistsdkgo.NewNetworkTemplate()
	data.SetName(plan.Name.ValueString())

	addictionalConfigCmds := convertPlanListOfString(ctx, plan.AdditionalConfigCmds)
	data.SetAdditionalConfigCmds(addictionalConfigCmds)

	dnsServers := convertPlanListOfString(ctx, plan.DnsServers)
	data.SetDnsServers(dnsServers)

	dnsSuffix := convertPlanListOfString(ctx, plan.DnsSuffix)
	data.SetDnsSuffix(dnsSuffix)

	ntpServers := convertPlanListOfString(ctx, plan.NtpServers)
	data.SetNtpServers(ntpServers)

	// NETWORKS
	net_data := make(map[string]mistsdkgo.NetworkTemplateNetwork)
	for vlan_name, vlan_data_attr := range plan.Networks.Elements() {
		var vlan_data_interface interface{} = vlan_data_attr
		vlan_plan := vlan_data_interface.(resource_networktemplate.NetworksValue)
		var vlan_id int32 = int32(vlan_plan.VlanId.ValueInt64())
		var subnet string = vlan_plan.Subnet.ValueString()
		new_net := *mistsdkgo.NewNetworkTemplateNetwork()
		new_net.SetVlanId(vlan_id)
		new_net.SetSubnet(subnet)
		net_data[vlan_name] = new_net
	}
	data.SetNetworks(net_data)

	// RADIUS
	rc_plan := plan.RadiusConfig
	rc_data := mistsdkgo.NewJunosRadiusConfig()
	rc_data.SetAcctInterimInterval(int32(rc_plan.AcctInterimInterval.ValueInt64()))
	rc_data.SetAuthServersRetries(int32(rc_plan.AuthServersRetries.ValueInt64()))
	rc_data.SetAuthServersTimeout(int32(rc_plan.AuthServersTimeout.ValueInt64()))
	rc_data.SetCoaEnabled(rc_plan.CoaEnabled.ValueBool())
	rc_data.SetCoaPort(int32(rc_plan.CoaPort.ValueInt64()))
	rc_data.SetNetwork(rc_plan.Network.ValueString())
	rc_data.SetSourceIp(rc_plan.SourceIp.ValueString())
	// RADIUS Acct
	var rc_acct_data []mistsdkgo.RadiusAcctServer
	for _, acct_plan_attr := range rc_plan.RadiusAcctServers.Elements() {
		var acct_plan_interface interface{} = acct_plan_attr
		acct_plan := acct_plan_interface.(resource_networktemplate.RadiusAcctServersValue)
		acct_data := mistsdkgo.NewRadiusAcctServer(acct_plan.Host.ValueString(), int32(acct_plan.Port.ValueInt64()), acct_plan.Secret.ValueString())
		rc_acct_data = append(rc_acct_data, *acct_data)
	}
	rc_data.SetAcctServers(rc_acct_data)
	// RADIUS Auth
	var rc_auth_data []mistsdkgo.RadiusAuthServer
	for _, auth_plan_attr := range rc_plan.RadiusAuthServers.Elements() {
		var auth_plan_interface interface{} = auth_plan_attr
		auth_plan := auth_plan_interface.(resource_networktemplate.RadiusAuthServersValue)
		auth_data := mistsdkgo.NewRadiusAuthServer(auth_plan.Host.ValueString(), int32(auth_plan.Port.ValueInt64()), auth_plan.Secret.ValueString())
		auth_data.SetKeywrapEnabled(auth_plan.KeywrapEnabled.ValueBool())
		auth_data.SetKeywrapFormat(auth_plan.KeywrapFormat.ValueString())
		auth_data.SetKeywrapKek(auth_plan.KeywrapKek.ValueString())
		auth_data.SetKeywrapMack(auth_plan.KeywrapMack.ValueString())
		rc_auth_data = append(rc_auth_data, *auth_data)
	}
	rc_data.SetAuthServers(rc_auth_data)
	data.SetRadiusConfig(*rc_data)

	var orgId = plan.OrgId.ValueString()
	return data, orgId
}

func convertPlanListOfString(ctx context.Context, list basetypes.ListValue) []string {
	var items []string
	for _, item := range list.Elements() {
		items = append(items, strings.ReplaceAll(item.String(), "\"", ""))
	}
	return items
}

func convertDataListOfString(ctx context.Context, data []string) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	for _, item := range data {
		value := strings.ReplaceAll(item, "\"", "")
		items = append(items, types.StringValue(value))
	}
	list, _ := types.ListValue(items_type, items)
	return list
}
