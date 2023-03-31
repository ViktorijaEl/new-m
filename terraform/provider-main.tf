#####################
## Provider - Main ##
#####################

terraform {
  required_version = ">= 1.2.0"
   backend "gcs" {  
      bucket = "dev-github-terraform"
      prefix = "terraform/state"
   }
 }

provider "google" {
  project = var.gcp_project
  region  = var.gcp_region
}
