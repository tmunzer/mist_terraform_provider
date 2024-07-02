package provider

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	envHost     = "MIST_HOST"
	envApitoken = "MIST_API_TOKEN"
)

var _ provider.Provider = (*mistProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &mistProvider{}
	}
}

type mistProvider struct {
	version string
}
type mistProviderData struct {
	client *mistsdkgo.APIClient
}
type mistProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Apitoken types.String `tfsdk:"apitoken"`
}

func (p *mistProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

	resp.Schema = schema.Schema{
		MarkdownDescription: "The Mist Provider allows Terraform to manage Juniper Mist.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "URL of the Mist Cloud, e.g. `api.mist.com`\n." +
					"The preferred approach is to pass the credentials as environment variables `",
				Optional: true,
			},
			"apitoken": schema.StringAttribute{
				MarkdownDescription: "The Mist API Token",
				Optional:            true,
			},
		},
	}
}

func (p *mistProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config mistProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Mist API Host",
			"The provider cannot create the Mist API client as there is an unknown configuration value for the Mist API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the MIST_HOST environment variable.",
		)
	}
	if config.Apitoken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("apitoken"),
			"Unknown Mist API API Token",
			"The provider cannot create the Mist API client as there is an unknown configuration value for the Mist API Token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the MIST_APITOKEN environment variable.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("MIST_HOST")
	apitoken := os.Getenv("MIST_APITOKEN")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}
	if !config.Apitoken.IsNull() {
		apitoken = config.Apitoken.ValueString()
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Mist API Host",
			"The provider cannot create the Mist API client as there is a missing or empty value for the Mist API host. "+
				"Set the host value in the configuration or use the MIST_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if apitoken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("apitoken"),
			"Missing Mist API Token",
			"The provider cannot create the Mist API client as there is a missing or empty value for the Mist API Token. "+
				"Set the host value in the configuration or use the MIST_APITOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	configuration := mistsdkgo.NewConfiguration()
	configuration.Host = host
	configuration.AddDefaultHeader("Authorization", "Token "+apitoken)
	client := mistsdkgo.NewAPIClient(configuration)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *mistProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "mist"
	resp.Version = p.version
}

func (p *mistProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *mistProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOrgResource,
		NewOrgSiteGroupResource,
		NewOrgNetworkTemplate,
		NewOrgServiceResource,
		NewOrgNetworkResource,
		NewOrgGatewayTemplate,
		NewOrgNacTag,
		NewOrgNacRule,
		NewOrgRfTemplate,
		NewOrgWlanTemplate,
		NewOrgWlan,
		NewOrgWxTag,
		NewOrgWxRule,
		NewSiteResource,
		NewSiteSettingResource,
<<<<<<< Updated upstream
		NewNetworkTemplate,
		NewServiceResource,
		NewNetworkResource,
		NewGatewayTemplate,
		NewNacTag,
		NewNacRule,
		NewRfTemplate,
		NewWlanTemplate,
		NewWlan,
=======
		NewSiteWlan,
>>>>>>> Stashed changes
	}
}
