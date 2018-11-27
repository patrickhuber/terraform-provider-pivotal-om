package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

type provider struct {
	target                        string
	skipAuthenticationSetup       bool
	authenticationType            string
	username                      string
	password                      string
	decryptionPassphrase          string
	eulaAccepted                  bool
	httpProxyURL                  string
	httpsProxyURL                 string
	noProxy                       string
	internalAuthenticationOptions *internalAuthenticationOptions
	ldapAuthenticationOptions     *ldapAuthenticationOptions
	samlAuthenticationOptions     *samlAuthenticationOptions
}

type internalAuthenticationOptions struct {
	username string
	password string
}

type ldapAuthenticationOptions struct {
	serverURL         string
	serverSSLCert     string
	username          string
	password          string
	emailAttribute    string
	groupSearchFilter string
	groupSearchBase   string
	userSearchFilter  string
	userSearchBase    string
	referrals         string
}

type samlAuthenticationOptions struct {
	idpMetadata         string
	boshIDPMetadata     string
	rbacAdminGroup      string
	rbacGroupsAttribute string
}

func mapProvider(resourceData *schema.ResourceData) (*provider, error) {
	var err error
	p := &provider{}

	if p.username, err = getRequiredString(resourceData, "username"); err != nil {
		return nil, err
	}

	if p.password, err = getRequiredString(resourceData, "password"); err != nil {
		return nil, err
	}

	if p.target, err = getRequiredString(resourceData, "target"); err != nil {
		return nil, err
	}

	// if the flag is not found (default is skip) or the value to skip is true then skip the rest of the properties
	var ok bool
	if p.skipAuthenticationSetup, ok = getBoolOk(resourceData, "skip_authentication_setup"); !ok || p.skipAuthenticationSetup {
		return p, nil
	}

	if p.authenticationType, err = getRequiredString(resourceData, "authentication_type"); err != nil {
		return nil, err
	}

	if p.decryptionPassphrase, err = getRequiredString(resourceData, "decryption_passphrase"); err != nil {
		return nil, err
	}

	if p.eulaAccepted, err = getRequiredBool(resourceData, "eula_accepted"); err != nil {
		return nil, err
	}

	if p.httpProxyURL, err = getRequiredString(resourceData, "http_proxy"); err != nil {
		return nil, err
	}

	if p.httpsProxyURL, err = getRequiredString(resourceData, "https_proxy"); err != nil {
		return nil, err
	}

	if p.noProxy, err = getRequiredString(resourceData, "no_proxy"); err != nil {
		return nil, err
	}

	switch p.authenticationType {
	case "internal":
		if p.internalAuthenticationOptions, err = mapInternalAuthenticationOptions(resourceData); err != nil {
			return nil, err
		}
	case "ldap":
		if p.ldapAuthenticationOptions, err = mapLdapAuthenticationOptions(resourceData); err != nil {
			return nil, err
		}
	case "saml":
		if p.samlAuthenticationOptions, err = mapSamlAuthenticationOptions(resourceData); err != nil {
			return nil, err
		}
	}

	return p, nil
}

func mapInternalAuthenticationOptions(resourceData *schema.ResourceData) (*internalAuthenticationOptions, error) {
	ops := &internalAuthenticationOptions{}
	var err error
	if ops.password, err = getRequiredString(resourceData, "internal_authentication_options.username"); err != nil {
		return nil, err
	}
	if ops.username, err = getRequiredString(resourceData, "internal_authentication_options.password"); err != nil {
		return nil, err
	}
	return ops, nil
}

func mapLdapAuthenticationOptions(resourceData *schema.ResourceData) (*ldapAuthenticationOptions, error) {
	ops := &ldapAuthenticationOptions{}
	var err error
	if ops.password, err = getRequiredString(resourceData, "ldap_authentication_options.username"); err != nil {
		return nil, err
	}
	if ops.username, err = getRequiredString(resourceData, "ldap_authentication_options.password"); err != nil {
		return nil, err
	}
	return ops, nil
}

func mapSamlAuthenticationOptions(resourceData *schema.ResourceData) (*samlAuthenticationOptions, error) {
	ops := &samlAuthenticationOptions{}
	var err error
	if ops.boshIDPMetadata, err = getRequiredString(resourceData, "saml_authentication_options.bosh_idp_metadata"); err != nil {
		return nil, err
	}
	if ops.idpMetadata, err = getRequiredString(resourceData, "saml_authentication_options.idp_metadata"); err != nil {
		return nil, err
	}
	if ops.rbacAdminGroup, err = getRequiredString(resourceData, "rbac_admin_group"); err != nil {
		return nil, err
	}
	if ops.rbacGroupsAttribute, err = getRequiredString(resourceData, "rbac_groups_attribute"); err != nil {
		return nil, err
	}
	return ops, nil
}

func getRequiredBool(resourceData *schema.ResourceData, key string) (bool, error) {
	value, ok := getBoolOk(resourceData, key)
	if !ok {
		return false, fmt.Errorf("required field '%s' was not supplied", key)
	}
	return value, nil
}

func getBoolOk(resourceData *schema.ResourceData, key string) (bool, bool) {
	value, ok := resourceData.GetOk(key)
	if !ok {
		return false, ok
	}
	return value.(bool), true
}

func getRequiredString(resourceData *schema.ResourceData, key string) (string, error) {
	value, ok := getStringOk(resourceData, key)
	if !ok {
		return "", fmt.Errorf("required field '%s' was not supplied", key)
	}
	return value, nil
}

func getStringOk(resourceData *schema.ResourceData, key string) (string, bool) {
	value, ok := resourceData.GetOk(key)
	if !ok {
		return "", ok
	}
	return value.(string), true
}
