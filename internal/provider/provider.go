package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const Name = "kubevip"

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		return &schema.Provider{
			Schema:               providerSchema(),
			ResourcesMap:         providerResources(),
			DataSourcesMap:       providerDataSources(),
			ConfigureContextFunc: configure(),
		}
	}
}

type Provider struct {
}

func init() {
	// Set descriptions to support markdown syntax
	schema.DescriptionKind = schema.StringMarkdown
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		resourceARPManifest: buildResourceARPManifest(),
	}
}

func providerDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}

func configure() schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return &Provider{}, nil
	}
}