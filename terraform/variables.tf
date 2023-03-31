#####################################
# Create Storage Bucket - Variables #
#####################################

variable "bucket-name" {
  type        = string
  description = "The name of the Google Storage Bucket to create"
}

variable "storage-class" {
  type        = string
  description = "The storage class of the Google Storage Bucket to create"
}



variable "vm-name" {
  type        = string
  description = ""
}

variable "machine_type" {} 

variable "gcp_zone" {}

variable "image" {}

variable "network" {}
