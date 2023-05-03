################################
# Create Storage Bucket - Main #
################################

# Define the bucket module
module "bucket" {
  source = "./bucket"

  # Pass variables to the bucket module
  gcp_project   = var.gcp_project
  bucket-name   = var.bucket-name
  gcp_region    = var.gcp_region
  gcp_zone      = var.gcp_zone
  storage-class = var.storage-class 
}

