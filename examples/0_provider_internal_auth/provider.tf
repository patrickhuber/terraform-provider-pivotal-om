variable "pivotal_om_target" { type = "string" }
variable "pivotal_om_username" { type = "string" }
variable "pivotal_om_password" { type = "string" }
variable "pivotal_om_decryption_passphrase" { type = "string" }

provider "pivotal_om" {
    target =  "${var.pivotal_om_target}"
    username =  "${var.pivotal_om_username}"
    password =  "${var.pivotal_om_password}"
    decryption_passphrase = "$var.pivotal_om_decryption_passphrase}"    

    authentication_type = "internal" # internal | ldap | saml
    internal_authentication_options = {
        username = "${var.pivotal_om_username}"
        password = "${var.pivotal_om_password}"
    }
}