variable "pivotal_om_target" { type = "string" }
variable "pivotal_om_username" { type = "string" }
variable "pivotal_om_password" { type = "string" }

provider "pivotal_om" {
    target =  "${var.pivotal_om_target}"
    username =  "${var.pivotal_om_username}"
    password =  "${var.pivotal_om_password}"
    skip_authentication_setup = true
}