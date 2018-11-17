package main

import "github.com/hashicorp/terraform/helper/schema"

func resourceInternalUserStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceInternalUserStoreCreate,
		Read:   resourceInternalUserStoreRead,
		Update: resourceInternalUserStorerUpdate,
		Delete: resourceInternalUserStoreDelete,

		Schema: map[string]*schema.Schema{
			"decryption_passphrase": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
			"no_proxyh": &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func resourceInternalUserStoreCreate(d *schema.ResourceData, m interface{}) error {
	return resourceInternalUserStoreRead(d, m)
}

func resourceInternalUserStoreRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInternalUserStorerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceInternalUserStoreRead(d, m)
}

func resourceInternalUserStoreDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
