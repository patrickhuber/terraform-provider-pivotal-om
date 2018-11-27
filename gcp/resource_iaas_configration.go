package gcp

import "github.com/hashicorp/terraform/helper/schema"

func ResourceIaasConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,				
			},
			"service_account_key": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"region": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"resource_prefix" : &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
