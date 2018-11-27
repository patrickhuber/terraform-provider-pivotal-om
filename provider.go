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
			"skip_authentication_setup":{
				Type: schema.TypeBool,
				Optional: true,
				Description: "skips authentication setup",
			},
			"authentication_type": {
				Type: schema.TypeString,
				Required: true,
				Description: "authentication type (internal|saml|ldap)",
			},			
			"internal_authentication_options":{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"password": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"saml_authentication_options":{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bosh_idp_metadata": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"idp_metadata": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"rbac_admin_group": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"rbac_groups_attribute": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"ldap_authentication_options":{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_url": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},						
						"server_ssl_cert": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},						
						"username": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"password": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"email_attribute": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"group_search_filter": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
						},
						"group_search_base": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
						},
						"user_search_filter": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
						},
						"user_search_base": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
						},
						"referrals" : &schema.Schema{
							Type: schema.TypeString,
							Required: true,
							Default: "follow",
						},
					},
				},
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
			"decryption_passphrase" : {
				Type: schema.TypeString,
				Required: false,				
			},
			"eula_accepted": &schema.Schema{
				Type:    schema.TypeBool,
				Default: true,
			},
			"http_proxy": &schema.Schema{
				Type: schema.TypeString,
			},
			"https_proxy": &schema.Schema{
				Type: schema.TypeString,
			},
			"no_proxy": &schema.Schema{
				Type: schema.TypeString,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
		},
	}
}
