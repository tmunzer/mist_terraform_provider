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

	////////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//RADIUS

	// // RADIUS Acct
	rc_srv_state_type := make(map[string]attr.Type)
	rc_srv_state_type["host"] = basetypes.StringType{}
	rc_srv_state_type["port"] = basetypes.Int64Type{}
	rc_srv_state_type["secret"] = basetypes.StringType{}
	rc_srv_state_type["keywrap_enabled"] = basetypes.BoolType{}
	rc_srv_state_type["keywrap_format"] = basetypes.StringType{}
	rc_srv_state_type["keywrap_kek"] = basetypes.StringType{}
	rc_srv_state_type["keywrap_mack"] = basetypes.StringType{}

	// // RADIUS Acct
	var rc_acct_list_type attr.Type = resource_networktemplate.AcctServersValue{}.Type(ctx)
	var rc_acct_list_value []attr.Value
	if len(data.RadiusConfig.AcctServers) > 0 {
		for _, srv_data := range data.RadiusConfig.GetAcctServers() {
			rc_srv_state_value := make(map[string]attr.Value)
			rc_srv_state_value["host"] = types.StringValue(srv_data.GetHost())
			rc_srv_state_value["port"] = types.Int64Value(int64(srv_data.GetPort()))
			rc_srv_state_value["secret"] = types.StringValue(srv_data.GetSecret())
			rc_srv_state_value["keywrap_enabled"] = types.BoolValue(srv_data.GetKeywrapEnabled())
			rc_srv_state_value["keywrap_format"] = types.StringValue(srv_data.GetKeywrapFormat())
			rc_srv_state_value["keywrap_kek"] = types.StringValue(srv_data.GetKeywrapKek())
			rc_srv_state_value["keywrap_mack"] = types.StringValue(srv_data.GetKeywrapMack())
			test, e := resource_networktemplate.NewAcctServersValue(rc_srv_state_type, rc_srv_state_value)
			if e != nil {
				mist_error.LogError(ctx, e, "NewRadiusAcctServersValue.NewRadiusAcctServersValue")
			}
			rc_acct_list_value = append(rc_acct_list_value, test)

		}
	}
	rc_acct_state, e := types.ListValueFrom(ctx, rc_acct_list_type, rc_acct_list_value)
	if e != nil {
		mist_error.LogError(ctx, e, "NewRadiusAcctServersValue.ListValueFrom")
	}

	// // RADIUS Auth

	var rc_auth_list_type attr.Type = resource_networktemplate.AuthServersValue{}.Type(ctx)
	var rc_auth_list_value []attr.Value
	if len(data.RadiusConfig.AuthServers) > 0 {
		for _, srv_data := range data.RadiusConfig.GetAuthServers() {
			rc_srv_state_value := make(map[string]attr.Value)
			rc_srv_state_value["host"] = types.StringValue(srv_data.GetHost())
			rc_srv_state_value["port"] = types.Int64Value(int64(srv_data.GetPort()))
			rc_srv_state_value["secret"] = types.StringValue(srv_data.GetSecret())
			rc_srv_state_value["keywrap_enabled"] = types.BoolValue(srv_data.GetKeywrapEnabled())
			rc_srv_state_value["keywrap_format"] = types.StringValue(srv_data.GetKeywrapFormat())
			rc_srv_state_value["keywrap_kek"] = types.StringValue(srv_data.GetKeywrapKek())
			rc_srv_state_value["keywrap_mack"] = types.StringValue(srv_data.GetKeywrapMack())
			test, e := resource_networktemplate.NewAuthServersValue(rc_srv_state_type, rc_srv_state_value)
			if e != nil {
				mist_error.LogError(ctx, e, "NewRadiusAuthServersValue.NewRadiusAuthServersValue")
			}
			rc_auth_list_value = append(rc_auth_list_value, test)
		}
	}
	rc_auth_state, e := types.ListValueFrom(ctx, rc_auth_list_type, rc_auth_list_value)
	if e != nil {
		mist_error.LogError(ctx, e, "NewRadiusAuthServersValue.ListValueFrom")
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	attrTypes := make(map[string]attr.Type)
	attrTypes["acct_interim_interval"] = basetypes.Int64Type{}
	attrTypes["auth_servers_retries"] = basetypes.Int64Type{}
	attrTypes["auth_servers_timeout"] = basetypes.Int64Type{}
	attrTypes["coa_enabled"] = basetypes.BoolType{}
	attrTypes["coa_port"] = basetypes.Int64Type{}
	attrTypes["network"] = basetypes.StringType{}
	attrTypes["acct_servers"] = basetypes.ListType{ElemType: resource_networktemplate.AcctServersValue{}.Type(ctx)}
	attrTypes["auth_servers"] = basetypes.ListType{ElemType: resource_networktemplate.AuthServersValue{}.Type(ctx)}
	attrTypes["source_ip"] = basetypes.StringType{}

	rc_state_value := make(map[string]attr.Value)
	rc_state_value["acct_interim_interval"] = types.Int64Value(int64(data.RadiusConfig.GetAcctInterimInterval()))
	rc_state_value["auth_servers_retries"] = types.Int64Value(int64(data.RadiusConfig.GetAuthServersRetries()))
	rc_state_value["auth_servers_timeout"] = types.Int64Value(int64(data.RadiusConfig.GetAuthServersTimeout()))
	rc_state_value["coa_enabled"] = types.BoolValue(data.RadiusConfig.GetCoaEnabled())
	rc_state_value["coa_port"] = types.Int64Value(int64(data.RadiusConfig.GetCoaPort()))
	rc_state_value["network"] = types.StringValue(data.RadiusConfig.GetNetwork())
	rc_state_value["source_ip"] = types.StringValue(data.RadiusConfig.GetSourceIp())
	rc_state_value["acct_servers"] = rc_acct_state
	rc_state_value["auth_servers"] = rc_auth_state

	rc_state, e := resource_networktemplate.NewRadiusConfigValue(attrTypes, rc_state_value)
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
	////////////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////////
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
	for _, acct_plan_attr := range rc_plan.AcctServers.Elements() {
		var acct_plan_interface interface{} = acct_plan_attr
		acct_plan := acct_plan_interface.(resource_networktemplate.AcctServersValue)
		acct_data := mistsdkgo.NewRadiusAcctServer(acct_plan.Host.ValueString(), int32(acct_plan.Port.ValueInt64()), acct_plan.Secret.ValueString())
		acct_data.SetKeywrapEnabled(acct_plan.KeywrapEnabled.ValueBool())
		acct_data.SetKeywrapFormat(acct_plan.KeywrapFormat.ValueString())
		acct_data.SetKeywrapKek(acct_plan.KeywrapKek.ValueString())
		acct_data.SetKeywrapMack(acct_plan.KeywrapMack.ValueString())
		rc_acct_data = append(rc_acct_data, *acct_data)
	}
	rc_data.SetAcctServers(rc_acct_data)
	// RADIUS Auth
	var rc_auth_data []mistsdkgo.RadiusAuthServer
	for _, auth_plan_attr := range rc_plan.AuthServers.Elements() {
		var auth_plan_interface interface{} = auth_plan_attr
		auth_plan := auth_plan_interface.(resource_networktemplate.AuthServersValue)
		auth_data := mistsdkgo.NewRadiusAuthServer(auth_plan.Host.ValueString(), int32(auth_plan.Port.ValueInt64()), auth_plan.Secret.ValueString())
		auth_data.SetKeywrapEnabled(auth_plan.KeywrapEnabled.ValueBool())
		auth_data.SetKeywrapFormat(auth_plan.KeywrapFormat.ValueString())
		auth_data.SetKeywrapKek(auth_plan.KeywrapKek.ValueString())
		auth_data.SetKeywrapMack(auth_plan.KeywrapMack.ValueString())
		rc_auth_data = append(rc_auth_data, *auth_data)
	}
	rc_data.SetAuthServers(rc_auth_data)
	data.SetRadiusConfig(*rc_data)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// PORT USAGE
	// pu_plan := plan.PortUsages
	// pu_data := make(map[string]mistsdkgo.JunosPortUsages)
	// for _, pu := range pu_plan.Elements(){

	// 	pu_data := mistsdkgo.NewJunosPortUsages()
	// 	pu_data.SetAllNetworks(pu)
	// }
	// data.SetPortUsages(pu_data)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////////
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
