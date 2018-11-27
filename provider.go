package main

import (
	"fmt"	
	"time"
	"errors"
	"strconv"
	"github.com/patrickhuber/terraform-provider-pivotal-om/aws"
	"github.com/patrickhuber/terraform-provider-pivotal-om/gcp"
	"github.com/patrickhuber/terraform-provider-pivotal-om/azure"
	"github.com/patrickhuber/terraform-provider-pivotal-om/vsphere"
	"github.com/patrickhuber/terraform-provider-pivotal-om/openstack"
	"github.com/patrickhuber/terraform-provider-pivotal-om/director"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/network"
)

// Provider creates a provider for pivotal-om https://www.terraform.io/docs/extend/writing-custom-providers.html#the-provider-schema
func Provider() *schema.Provider {
	return &schema.Provider{	
		Schema: providerSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"pivotal_om_aws_iaas_configuration" : aws.ResourceIaasConfiguration(),
			"pivotal_om_aws_availability_zone": aws.ResourceAvailabilityZone(),
			"pivotal_om_gcp_iaas_configuration": gcp.ResourceIaasConfiguration(),
			"pivotal_om_gcp_availability_zone": gcp.ResourceAvailabilityZone(),
			"pivotal_om_azure_iaas_configuration": azure.ResourceIaasConfiguration(),
			"pivotal_om_vsphere_iaas_configuration": vsphere.ResourceIaasConfiguration(),
			"pivotal_om_vsphere_availability_zone": vsphere.ResourceAvailabilityZone(),
			"pivotal_om_openstack_iaas_configuration": openstack.ResourceIaasConfiguration(),
			"pivotal_om_openstack_availability_zone": openstack.ResourceAvailabilityZone(),
			"pivotal_om_director": director.ResourceDirector(),
			"pivotal_om_network": director.ResourceNetwork(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(resourceData *schema.ResourceData) (interface{}, error){	
	provider, err := mapProvider(resourceData)
	if err != nil{
		return nil, err
	}	

	if provider.skipAuthenticationSetup {
		return provider, nil
	}

	service, err := newOMService(provider)
	if err != nil{
		return nil, err
	}

	ensureAvailabilityOutput, err := service.EnsureAvailability(api.EnsureAvailabilityInput{})
	if err != nil{
		return nil, err
	}
	
	if ensureAvailabilityOutput.Status == api.EnsureAvailabilityStatusUnknown {
		return nil, errors.New("could not determine initial configuration status: received unexpected status")
	}

	if ensureAvailabilityOutput.Status != api.EnsureAvailabilityStatusUnstarted {		
		return provider, nil
	}
	
	setupInput, err := mapSetupInput(provider)
	if err != nil{
		return nil, err
	}

	_, err = service.Setup(*setupInput)
	if err != nil{
		return nil, err
	}

	for ensureAvailabilityOutput.Status != api.EnsureAvailabilityStatusComplete {
		ensureAvailabilityOutput, err = service.EnsureAvailability(api.EnsureAvailabilityInput{})
		if err != nil {
			return nil, fmt.Errorf("could not determine final configuration status: %s", err)
		}
	}

	return provider, nil
}

func mapSetupInput(provider *provider) (*api.SetupInput, error){
	setupInput := &api.SetupInput{
		AdminUserName: provider.username,
		AdminPassword: provider.password,
		AdminPasswordConfirmation: provider.password,
		DecryptionPassphrase: provider.decryptionPassphrase,
		EULAAccepted: strconv.FormatBool(provider.eulaAccepted),
		HTTPSProxyURL: provider.httpsProxyURL,
		HTTPProxyURL: provider.httpProxyURL,
		NoProxy: provider.noProxy,
	}
	setupInput.IdentityProvider = provider.authenticationType
	if provider.samlAuthenticationOptions != nil{				
		setupInput.IDPMetadata = provider.samlAuthenticationOptions.idpMetadata
		setupInput.BoshIDPMetadata = provider.samlAuthenticationOptions.boshIDPMetadata
		setupInput.RBACAdminGroup = provider.samlAuthenticationOptions.rbacAdminGroup
		setupInput.RBACGroupsAttribute = provider.samlAuthenticationOptions.rbacGroupsAttribute
	}
	if provider.ldapAuthenticationOptions != nil{
		
	}
	return setupInput, nil
}

func newOMService(provider *provider) (*api.Api, error) {
	client, err := newOAuthClient(provider)
	if err != nil{
		return nil, err
	}	

	service := api.New(api.ApiInput{
		Client: client,		
	})

	return &service, nil
}

func newOAuthClient(provider *provider) (*network.OAuthClient, error){	
	clientID := ""
	clientSecret := ""
	skipSSLValidation := true
	includeCookies := true
	requestTimeout := time.Duration(1800) * time.Second
	connectTimeout := time.Duration(5) * time.Second

	client, err :=  network.NewOAuthClient(
		provider.target, 
		provider.username, 
		provider.password, 
		clientID, 
		clientSecret, 
		skipSSLValidation, 
		includeCookies, 
		requestTimeout, 
		connectTimeout)

	if err != nil {
		return nil, err
	}
	return &client, nil
}

func providerSchema() map[string]*schema.Schema{
	return map[string]*schema.Schema{
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
	}	
}