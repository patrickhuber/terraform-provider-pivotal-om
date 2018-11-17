package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider creates a provider for pivotal-om https://www.terraform.io/docs/extend/writing-custom-providers.html#the-provider-schema
func Provider() *schema.Provider {
	return &schema.Provider{	
		Schema: map[string]*schema.Schema{
			"target": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("OM_TARGET", nil),
				Description: "location of the Ops Manager VM",
			},
			"username": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("OM_USERNAME", nil),
				Description: "admin username for the Ops Manager VM",
			},
			"password": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("OM_PASSWORD", nil),
				Description: "admin password for the Ops Manager VM",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"internal_user_store": resourceInternalUserStore(),
		},
	}
}
