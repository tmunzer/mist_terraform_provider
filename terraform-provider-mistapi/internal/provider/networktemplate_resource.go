package provider

import (
	"context"
	"fmt"
	"strings"
	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
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

	state = processNetworktemplateData(ctx, data)

	diags = resp.State.Set(ctx, state)
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

	var t = map[string]attr.Type{
		"vlan_id": basetypes.Int64Type{},
		"subnet":  basetypes.StringType{},
	}
	var l attr.Type = resource_networktemplate.NetworksValue{}.Type(ctx)
	m := make(map[string]attr.Value)
	for k, v := range data.GetNetworks() {

		var net = map[string]attr.Value{
			"vlan_id": types.Int64Value(int64(v.GetVlanId())),
			"subnet":  types.StringValue(v.GetSubnet()),
		}
		n, _ := resource_networktemplate.NewNetworksValue(t, net)
		m[k] = n
	}
	state.Networks, _ = types.MapValueFrom(ctx, l, m)

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

	networks := plan.Networks.Elements()
	m := make(map[string]mistsdkgo.NetworkTemplateNetwork)
	for vlan_name, vlan_data_attr := range networks {
		var vlan_data_interface interface{} = vlan_data_attr
		vlan_data := vlan_data_interface.(resource_networktemplate.NetworksValue)
		var vlan_id int32 = int32(vlan_data.VlanId.ValueInt64())
		var subnet string = strings.ReplaceAll(vlan_data.Subnet.String(), "\"", "")
		d := *mistsdkgo.NewNetworkTemplateNetwork()
		d.SetVlanId(vlan_id)
		d.SetSubnet(subnet)
		m[vlan_name] = d
	}
	data.SetNetworks(m)

	var orgId = plan.OrgId.ValueString()
	return data, orgId
}

func convertPlanListOfString(ctx context.Context, list basetypes.ListValue) []string {
	var items []string
	for _, item := range list.Elements() {
		items = append(items, item.String())
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
