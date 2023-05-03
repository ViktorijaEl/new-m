# Create a GCS Bucket
resource "google_storage_bucket" "tf-bucket" {
  project       = var.gcp_project
  name          = var.bucket-name
  location      = var.gcp_region
  force_destroy = true
  public_access_prevention = "enforced"
  uniform_bucket_level_access = true
  storage_class = var.storage-class  

  versioning {
    enabled = true
  }

  logging {
    log_bucket = "dev-github-terraform"
  }
}
