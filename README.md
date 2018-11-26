# Terraform Provider Pivotal Ops Manager

Provides a terraform provider for pivotal ops manager

## Getting Started

### Install the plugin

[install](docs/INSTALL.md)

### Create the provider

```HCL
variable "pivotal_om_target" { type = "string" }
variable "pivotal_om_username" { type = "string" }
variable "pivotal_om_password" { type = "string" }
variable "pivotal_om_decryption_passphrase" { type = "string" }

provider "pivotal_om" {
    target =  "${var.pivotal_om_target}"
    username =  "${var.pivotal_om_username}"
    password =  "${var.pivotal_om_password}"
}
```

### Create the user store resource

```HCL
resource "pivotal_om_internal_user_store" "user_store" {
    decryption_passphrase =  "${var.pivotal_om_decryption_passphrase}"    
}
```

### Configure IaaS

#### aws

```HCL

variable "aws_ssh_private_key" { type = "string" }
variable "aws_access_key_id" {}
variable "aws_access_key_secret" {}

resource "pivotal_om_aws_iaas_configuration" "iaas_configuration" {
    director_configuration_id = "${pivotal_om_director.director_configuration.id}"
    access_key_id =  ""
    access_key_secret =  ""
    iam_instance_profile =  ""
    vpc_id =  ""
    security_group_id =  ""
    key_pair_name =  "pcf-ops-manager-key"
    ssh_private_key =  "${var.aws_ssh_private_key}"
    region =  "us-east-1"
    encrypt_ebs_volumes =  false
}

resource "pivotal_om_aws_availability_zone" "az1"{
    name = "us-east-1a"
    zone = "us-east-1a"    
    iaas_configuration_id = "${pivotal_om_aws_iaas_configuration.iaas_configuration}"
}

resource "pivotal_om_aws_availability_zone" "az2"{
    name = "us-east-1b"
    zone = "us-east-1b"
    iaas_configuration_id = "${pivotal_om_aws_iaas_configuration.iaas_configuration}"
}

resource "pivotal_om_aws_availability_zone" "az3"{
    name = "us-east-1c"
    zone = "us-east-1c"
    iaas_configuration_id = "${pivotal_om_aws_iaas_configuration.iaas_configuration}"
}
```

#### gcp

```
resource "pivotal_om_gcp_iaas_configuration" "iaas_configuration"{
    director_configuration_id = "${pivotal_om_director.director_configuration.id}"
    project_id = ""
    resource_prefix = "" # is this needed?
    service_account_key = ""
    region = "us_central1"
}

resource "pivotal_om_gcp_az" "az1" {
    name = "us-central1-a"
    zone = "us-central1-a"    
    iaas_configuration_id = "${pivotal_om_gcp_iaas_configuration.iaas_configuration}"
}


resource "pivotal_om_gcp_az" "az2" {
    name = "us-central1-b"
    zone = "us-central1-b"
    iaas_configuration_id = "${pivotal_om_gcp_iaas_configuration.iaas_configuration}"
}


resource "pivotal_om_gcp_az" "az3" {
    name = "us-central1-c"
    zone = "us-central1-c"
    iaas_configuration_id = "${pivotal_om_gcp_iaas_configuration.iaas_configuration}"
}
```

#### azure

```
resource "pivotal_om_azure_iaas_configuration" "iaas_configuration" {
    director_configuration_id = "${pivotal_om_director.director_configuration.id}"
    subscription_id = ""
    tenant_id = ""
    client_id = ""
    client_secret = ""
    resource_group_name = ""
    bosh_storage_account_name = ""
    deployments_storage_account_name = ""
    default_security_group = ""
    ssh_public_key = ""
    ssh_private_key = ""
    cloud_storage_type = ""
    storage_account_type = ""
    environment = "AzureCloud"
}
```

#### vsphere

```
resource "pivotal_om_vsphere_iaas_configuration" "iaas_configuration"{    
    director_configuration_id = "${pivotal_om_director.director_configuration.id}"
    vcenter_host = ""
    vcenter_username = ""
    vcenter_password = ""
    datacenter = ""
    disk_type = ""
    ephemeral_datastores_string = ""
    persistent_datastores_string = ""
    bosh_vm_folder = ""
    bosh_template_folder = ""
    bosh_disk_path = ""
    ssl_verification_enabled = false
    nsx_networking_enabled = false
}

resource "pivotal_om_vsphere_az" "az1" {
    name = "az1"
    cluster = ""
    resource_pool = "" # optional
    iaas_configuration_id = "${pivotal_om_vsphere_iaas_configuration.iaas_configuration}"
}

resource "pivotal_om_vsphere_az" "az2" {    
    name = "az2"
    cluster = ""
    resource_pool = "" # optional
    iaas_configuration_id = "${pivotal_om_vsphere_iaas_configuration.iaas_configuration}"
}

resource "pivotal_om_vsphere_az" "az3" {    
    name = "az3"
    cluster = ""
    resource_pool = "" # optional
    iaas_configuration_id = "${pivotal_om_vsphere_iaas_configuration.iaas_configuration}"
}
```

#### openstack

```
resource "pivotal_om_openstack_iaas_configuration" "iaas_configuration"{    
    director_configuration_id = "${pivotal_om_director.director_configuration.id}"
    identity_endpoint = ""
    username = ""
    password = ""
    keystone_version = ""
    domain = ""
    tenant = ""
    networking_model = ""
    security_group = ""
    key_pair_name = ""
    ssh_private_key = ""
    region = ""
    ignore_server_availability_zone = true
    dhcp_enabled = false
    api_ssl_cert = ""
}

resource "pivotal_om_openstack_az" "az1" {
    name = "az1"
    zone = "nova"
    iaas_configuration_id = "${pivotal_om_openstack_iaas_configuration.iaas_configuration}"
}

resource "pivotal_om_openstack_az" "az2" {
    name = "az2"
    zone = ""
    iaas_configuration_id = "${pivotal_om_openstack_iaas_configuration.iaas_configuration}"
}

resource "pivotal_om_openstack_az" "az3" {
    name = "az3"
    zone = ""
    iaas_configuration_id = "${pivotal_om_openstack_iaas_configuration.iaas_configuration}"
}
```

#### director

```
resource "pivotal_om_director" "director_configuration"{
    ntp_servers  =  ["0.amazon.pool.ntp.org", "1.amazon.pool.ntp.org", "3.amazon.pool.ntp.org"]
    resurrector_enabled =  false
    post_deploy_scripts_enabled =  false
    recreate_all_vms_enabled =  false
    bosh_deploy_retries_enabled =  false
    keep_unreachable_director_vms =  false

    # ("external" | "internal" )
    database_type = "external"    
    external_database_options {
        host = ""
        port = 3306
        user = "{var.external_database_username}"
        password = "{var.external_database_password}"
        database = "{var.external_database_name}"
    }

    # ("s3" | "internal" | "gcs")
    blobstore_type = "s3"

    s3_blobstore_location{
        s3_endpoint =  ""
        bucket_name =  ""
        access_key =  ""
        secret_key = ""
        signature_version =  4
        region =  "region"
    }
    
    networks_ids = [
        "${pivotal_om_network.infrastructure_network.id}",
        "${pivotal_om_network.deployment_network.id}",
        "${pivotal_om_network.services_network.id}"
    ]

    # optional for azure. put in iaas config?
    availibility_zones = []

    security {
        trusted_certificates = []
        generate_passwords = true
    }

    resource {        
        director{
            instance_type = "m4.large"
        }        
    }

    network_assignment {
        # for azure use `network_name = "${pivotal_om_network.infrastructure_network.name}"`
        singleton_az = "${pivotal_om_aws_availability_zone.az2.name}"        
    }
}

resource "pivotal_om_network" "deployment_network"{
    name = ""
    services_network = false
    subnet {
        iaas_identifier = ""
        cidr = ""
        reserved_ip_ranges = []
        dns = []
        gateway = ""
    }
    # empty or optional for azure
    availability_zones = ["az1", "az2", "az3"]
}

resource "pivotal_om_network" "infrastructure_network"{
    name = ""
    services_network = false
    subnet {
        iaas_identifier = ""
        cidr = ""
        reserved_ip_ranges = []
        dns = []
        gateway = ""
    }
    # empty or optional for azure
    availability_zones = ["az1", "az2", "az3"]
}

resource "pivotal_om_network" "services_network"{
    name = ""
    services_network = false
    subnet {
        iaas_identifier = ""
        cidr = ""
        reserved_ip_ranges = []
        dns = []
        gateway = ""
    }
    # empty or optional for azure
    availability_zones = ["az1", "az2", "az3"]
}
```